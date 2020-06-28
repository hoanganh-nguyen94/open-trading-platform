module github.com/ettech/open-trading-platform/go/market-data/market-data-service

go 1.13

require (
	github.com/ettec/otp-common v0.0.0-20200627105436-e95192c9a900
	github.com/ettec/otp-mdcommon v0.0.0-20200628160946-c4fb2dda07ff
	github.com/ettec/otp-model v0.0.2-0.20200627105317-ed67f7c52141
	github.com/golang/protobuf v1.4.0
	github.com/google/uuid v1.1.1
	github.com/prometheus/client_golang v1.6.0
	google.golang.org/grpc v1.25.1
	k8s.io/api v0.17.4
	k8s.io/apimachinery v0.17.4
	k8s.io/client-go v0.17.4
)

replace github.com/ettec/otp-mdcommon v0.0.0-20200628160946-c4fb2dda07ff => ../market-data-common
