### Install micro env

cd script && sh install.sh

### Run helloworld

cd helloworld/ && go run srv/main.go
cd helloworld/ && go run api/main.go

### Test

curl http://localhost:8080/helloworld/handler/call?name=John