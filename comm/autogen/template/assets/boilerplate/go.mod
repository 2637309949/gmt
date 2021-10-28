module {{.name}}-service

go 1.13

require (
	comm v0.0.0-00010101000000-000000000000
	github.com/micro/go-micro/v2 v2.9.1
	github.com/shopspring/decimal v1.2.0
	github.com/xormplus/xorm v0.0.0-20210822100304-4e1d4fcc1e67
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d // indirect
	golang.org/x/sys v0.0.0-20211013075003-97ac67df715c // indirect
	proto v0.0.0-00010101000000-000000000000
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace (
	comm => ../comm
	proto => ../proto
)
