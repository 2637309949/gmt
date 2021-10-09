module helloworld

go 1.13

require (
	comm v0.0.0-00010101000000-000000000000
	proto v0.0.0-00010101000000-000000000000
	github.com/micro/go-micro/v2 v2.9.1
	github.com/shopspring/decimal v1.2.0
	github.com/xormplus/xorm v0.0.0-20210822100304-4e1d4fcc1e67
)

replace (
	comm => github.com/2637309949/gmt/comm v0.0.0-20211009094502-0309d19a1e45
	proto => github.com/2637309949/gmt/proto v0.0.0-20211009094502-0309d19a1e45
)
