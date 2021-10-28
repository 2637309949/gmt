module proto

go 1.13

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gohouse/converter v0.0.3 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro/v2 v2.9.1
)

replace (
	comm => ../comm
	proto => ../proto
)
