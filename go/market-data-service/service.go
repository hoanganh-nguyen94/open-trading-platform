package main

import (
	"context"
	"fmt"
	"github.com/ettec/open-trading-platform/go/market-data-gateway/actor"
	mdgapi "github.com/ettec/open-trading-platform/go/market-data-gateway/api"
	"github.com/ettec/open-trading-platform/go/model"
	"github.com/ettech/open-trading-platform/go/market-data-service/api"
	"github.com/ettech/open-trading-platform/go/market-data-service/internal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

type service struct {
	partyIdToConnection map[string]actor.ClientConnection
	quoteDistributor    actor.QuoteDistributor
	connMux             sync.Mutex
}

func newService(id string, marketGatewayAddress string, maxReconnectInterval time.Duration) (*service, error) {

	mdcToDistributorChan := make(chan *model.ClobQuote, 1000)

	mdcFn := func(targetAddress string) (mdgapi.MarketDataGatewayClient, internal.GrpcConnection, error) {
		conn, err := grpc.Dial(targetAddress, grpc.WithInsecure(), grpc.WithBackoffMaxDelay(maxReconnectInterval))
		if err != nil {
			return nil, nil, err
		}

		client := mdgapi.NewMarketDataGatewayClient(conn)
		return client, conn, nil
	}

	mdc, err := internal.NewMarketDataGatewayClient(id, marketGatewayAddress, mdcToDistributorChan, mdcFn)

	if err != nil {
		return nil, err
	}

	qd := actor.NewQuoteDistributor(mdc.Subscribe, mdcToDistributorChan)
	s := &service{partyIdToConnection: make(map[string]actor.ClientConnection), quoteDistributor: qd}

	return s, nil
}

func (s *service) getConnection(partyId string) (actor.ClientConnection, bool) {
	s.connMux.Lock()
	defer s.connMux.Unlock()

	con, ok := s.partyIdToConnection[partyId]
	return con, ok
}

func (s *service) addConnection(subscriberId string, out chan<- *model.ClobQuote) actor.ClientConnection {
	s.connMux.Lock()
	defer s.connMux.Unlock()

	if conn, ok := s.partyIdToConnection[subscriberId]; ok {
		log.Printf("connection for client %v already exists, closing existing connection.", subscriberId)
		conn.Close()
		log.Print("connection closed:", subscriberId)
	}

	cc := actor.NewClientConnection(subscriberId, out, s.quoteDistributor, maxSubscriptions)

	s.partyIdToConnection[subscriberId] = cc

	return cc
}

func (s *service) Subscribe(_ context.Context, r *api.MdsSubscribeRequest) (*model.Empty, error) {

	if conn, ok := s.getConnection(r.SubscriberId); ok {

		if err := conn.Subscribe(r.ListingId); err != nil {
			return nil, err
		}

		return &model.Empty{}, nil
	} else {
		return nil, fmt.Errorf("failed to subscribe, no connection exists for subscriber " + r.SubscriberId)
	}

}

func (s *service) Connect(request *api.MdsConnectRequest, stream api.MarketDataService_ConnectServer) error {

	subscriberId := request.GetSubscriberId()

	log.Println("connect request received for subscriber ", subscriberId)

	out := make(chan *model.ClobQuote, 100)

	s.addConnection(subscriberId, out)

	for mdUpdate := range out {
		if err := stream.Send(mdUpdate); err != nil {
			log.Printf("error on connection for subscriber %v, closing connection, error:%v", subscriberId, err)
			break
		}
	}

	return nil
}

const (
	ServiceIdKey   = "SERVICE_ID"
	GatewayAddress = "GATEWAY_ADDRESS"
	ConnectRetrySeconds      = "CONNECT_RETRY_SECONDS"

	// The maximum number of listing subscriptions per connection
	MaxSubscriptionsKey = "MAX_SUBSCRIPTIONS"

)

var maxSubscriptions = 10000

func main() {

	port := "50651"
	fmt.Println("Starting Market Data Service on port:" + port)
	lis, err := net.Listen("tcp", "0.0.0.0:"+port)

	id, ok := os.LookupEnv(ServiceIdKey)
	if !ok {
		log.Panicf("expected %v env var to be set", ServiceIdKey)
	}

	fixSimAddress, ok := os.LookupEnv(GatewayAddress)
	if !ok {
		log.Panicf("expected %v env var to be set", GatewayAddress)
	}

	maxSubsEnv, ok := os.LookupEnv(MaxSubscriptionsKey)
	if ok {
		maxSubscriptions, err = strconv.Atoi(maxSubsEnv)
		if err != nil {
			log.Panicf("cannot parse %v, error: %v", MaxSubscriptionsKey, err)
		}
	}

	connectRetrySecs := getOptionalBootstrapIntEnvVar(ConnectRetrySeconds, 60 )

	if err != nil {
		log.Fatalf("Error while listening : %v", err)
	}

	s := grpc.NewServer()
	mdcService, err := newService(id, fixSimAddress, time.Duration(connectRetrySecs)*time.Second)
	if err != nil {
		log.Panicf("failed to create market data service:%v", err)
	}

	api.RegisterMarketDataServiceServer(s, mdcService)

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}

}


func getOptionalBootstrapIntEnvVar(key string, def int) int {
	strValue, exists := os.LookupEnv(key)
	result := def
	if exists {
		var err error
		result, err = strconv.Atoi(strValue)
		if err != nil {
			log.Panicf("cannot parse %v, error: %v", key, err)
		}
	}

	log.Printf("%v set to %v", key, result)

	return result
}