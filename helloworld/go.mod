module helloworld

go 1.13

require (
	github.com/micro/go-micro/v2 v2.9.1
	comm v0.0.0-00010101000000-000000000000
	proto v0.0.0-00010101000000-000000000000
)

replace proto => ../proto
replace comm => ../comm
