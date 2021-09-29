module helloworld

go 1.13

require (
	comm v0.0.0-00010101000000-000000000000
	github.com/gemnasium/logrus-graylog-hook/v3 v3.0.3 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/shopspring/decimal v1.2.0
	github.com/xormplus/xorm v0.0.0-20210822100304-4e1d4fcc1e67
	proto v0.0.0-00010101000000-000000000000
)

replace proto => ../proto

replace comm => ../comm
