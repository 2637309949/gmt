### Install micro env
```shell
cd script && sh install.sh
```

### Run helloworld
```shell
cd helloworld/ && go run srv/main.go
cd helloworld/ && go run api/main.go
```

### Test

```shell
curl 127.0.0.1:8080/helloworld/handler/articleAdd  --header \
"Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiMTYwMSIsImV4cCI6MTYzNDIwMjMzNiwiaWF0IjoxNjMxNjEwMzM2LCJpc3MiOiJiYWlkdS5jb20ifQ.1QrmhtcsDhVQ39BQGlOSH7Ckxt2hkHYK9943iquf7BA"
```
