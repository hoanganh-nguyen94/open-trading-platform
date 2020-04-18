module github.com/ettec/open-trading-platform/go/common

go 1.13

require (
	github.com/emicklei/go-restful v0.0.0-20170410110728-ff4f55a20633
	github.com/ettec/open-trading-platform/go/model v0.0.0
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	google.golang.org/grpc v1.25.1
	k8s.io/api v0.17.4
	k8s.io/apimachinery v0.17.4
	k8s.io/client-go v0.17.4

)

replace github.com/ettec/open-trading-platform/go/model v0.0.0 => ../model
