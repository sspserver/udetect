module github.com/sspserver/udetect/examples/server

go 1.16

require (
	github.com/demdxx/goconfig v0.0.0-20191123141255-c40c2d9e90f5
	github.com/gravitational/log v0.0.0-20200127200505-fdffa14162b0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0
	github.com/sspserver/udetect v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2
	go.uber.org/zap v1.19.1
	google.golang.org/grpc v1.41.0
	launchpad.net/gocheck v0.0.0-20140225173054-000000000087 // indirect
)

replace github.com/sspserver/udetect => ../../
