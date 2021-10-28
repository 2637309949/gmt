module subscriber-service

go 1.13

require (
	comm v0.0.0-00010101000000-000000000000
	proto v0.0.0-00010101000000-000000000000
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace (
	comm => ../comm
	proto => ../proto
)