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

replace (
	comm => github.com/2637309949/gmt/comm v0.0.0-20211009094502-0309d19a1e45
	proto => github.com/2637309949/gmt/proto v0.0.0-20211009094502-0309d19a1e45
)
