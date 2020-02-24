package actor

import (
	"github.com/ettec/open-trading-platform/go/model"
	"testing"
	"time"
)

type testMarketDataClient struct {
	subscribe func(listingId int32) error
	close     func() error
}

func (t *testMarketDataClient) Subscribe(listingId int32) error{
	t.subscribe(listingId)
	return nil
}

func (t *testMarketDataClient) Close() error {
	return t.close()
}

func TestNewMdServerConnection(t *testing.T) {

	connectedCalled := make(chan bool, 10)

	dial := func(target string, source chan<- *model.ClobQuote) (Connection, error) {
		connectedCalled <- true
		return &testMarketDataClient{
			subscribe: nil,
			close:     nil,
		}, nil
	}

	out := make(chan *model.ClobQuote, 100)
	mdConn := NewMdServerConnection("testconnection", out, dial, 0)

	if !<-connectedCalled {
		t.Errorf("expected distConnection method to be called")
	}

	if mdConn.connection == nil {
		t.Errorf("expected the md conn to have a distConnection")
	}

}


func TestSubscribe(t *testing.T) {

	subscribed := make(chan int32, 0)

	dial := func(target string, source chan<- *model.ClobQuote) (Connection, error) {
		return &testMarketDataClient{

			subscribe: func(listingId int32) error {
				subscribed <- listingId
				return nil
			},
			close: nil,
		}, nil
	}

	out := make(chan *model.ClobQuote, 100)
	mdConn := NewMdServerConnection("testconnection", out, dial, 0)

	mdConn.Subscribe(1)
	mdConn.Subscribe(2)



	if v := <-subscribed; v != 1 {
		t.Errorf("expected to receive 1, got %v", v)
	}

	if v :=<-subscribed; v != 2 {
		t.Errorf("expected to receive 2, got %v", v)
	}

}

func TestRefreshesAreForwardedToSink(t *testing.T) {

	var clobSource chan<- *model.ClobQuote

	connected := make(chan bool)
	dial := func(target string, source chan<- *model.ClobQuote) (Connection, error) {
		clobSource = source
		connected <- true
		return &testMarketDataClient{

			subscribe: nil,
			close:     nil,
		}, nil
	}

	out := make(chan *model.ClobQuote, 100)
	NewMdServerConnection("testconnection", out, dial, 0)

	<-connected

	clobSource <- &model.ClobQuote{ListingId: 1}
	clobSource <- &model.ClobQuote{ListingId: 2}

	r1 := <-out
	if r1.ListingId != 1 {
		t.Errorf("unexpected refresh")
	}

	r2 := <-out
	if r2.ListingId != 2 {
		t.Errorf("unexpected refresh")
	}

}

func TestSubscribesSentWhilstNotConnectedAreResentOnConnect(t *testing.T) {

	subscribed := make(chan int32, 20)
	var clobSource chan<- *model.ClobQuote
	connected := make(chan bool)

	newConnFn := func(target string, source chan<- *model.ClobQuote) (Connection, error) {
		clobSource = source
		connected <- true
		return &testMarketDataClient{
			subscribe: func(listingId int32) error {
				subscribed <- listingId
				return nil
			},
			close: nil,
		}, nil
	}

	out := make(chan *model.ClobQuote, 100)
	mdConn := NewMdServerConnection("testconnection", out, newConnFn, 1)

	<-connected
	close(clobSource)

	mdConn.Subscribe(1)
	mdConn.Subscribe(2)

	if <-subscribed != 1 {
		t.Error("expected to receive 1")
	}

	if <-subscribed != 2 {
		t.Error("expected to receive 2")
	}
}

func TestReconnectOccursAfterConnectionFailure(t *testing.T) {

	var clobSource chan<- *model.ClobQuote
	subscriptions := make(chan int32, 10)
	connected := make(chan bool)

	dial := func(target string, source chan<- *model.ClobQuote) (Connection, error) {
		clobSource = source
		connected <- true
		return &testMarketDataClient{

			subscribe: func(listingId int32) error {
				subscriptions <- listingId
				return nil
			},
			close: func() error {
				return nil
			},
		}, nil
	}

	reconnectInterval := 1 * time.Second

	out := make(chan *model.ClobQuote, 100)
	mdConn := NewMdServerConnection("testconnection", out, dial, 0)

	<-connected

	mdConn.Subscribe(1)
	mdConn.Subscribe(2)

	<-subscriptions
	<-subscriptions

	clobSource <- &model.ClobQuote{ListingId: 1}

	r1 := <-out
	if r1.ListingId != 1 {
		t.Errorf("unexpected refresh")
	}

	close(clobSource)

	time.Sleep(reconnectInterval / 2)

	time.Sleep(reconnectInterval)

	<-connected

	if s := <-subscriptions; s != 1 {
		t.Errorf("expected subscribeRequest to be resent for 1")
	}

	if s := <-subscriptions; s != 2 {
		t.Errorf("expected subscribeRequest to be resent for 2")
	}

}
