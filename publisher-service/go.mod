module publisher

go 1.13

require (
	comm v0.0.0-00010101000000-000000000000
	github.com/micro/go-micro/v2 v2.9.1
	github.com/netdata/go-orchestrator v0.0.0-20190905093727-c793edba0e8f
	github.com/shopspring/decimal v1.2.0
	github.com/stretchr/objx v0.2.0 // indirect
	proto v0.0.0-00010101000000-000000000000
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace (
	comm => ../comm
	proto => ../proto
)