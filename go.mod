module github.com/aurora-scheduler/gorealis/v2

require (
	github.com/apache/thrift v0.13.0
	github.com/pkg/errors v0.9.1
	github.com/samuel/go-zookeeper v0.0.0-20171117190445-471cd4e61d7a
	github.com/stretchr/testify v1.5.0
)

replace github.com/apache/thrift v0.13.0 => github.com/ridv/thrift v0.13.2

go 1.13
