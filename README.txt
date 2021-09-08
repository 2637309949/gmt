1. install micro env

cd script && sh install.sh

2. run helloworld

cd helloworld/ && go run srv/main.go
cd helloworld/ && go run api/main.go

3. test helloworld

curl "http://localhost:8080/helloworld/handler/call?name=John"