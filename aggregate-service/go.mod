module aggregate-service

go 1.13

require (
	comm v0.0.0-00010101000000-000000000000
	proto v0.0.0-00010101000000-000000000000
	github.com/fortytw2/leaktest v1.3.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/olivere/elastic v6.2.37+incompatible
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace (
	comm => ../comm
	proto => ../proto
)
