### Quick start

- Install autogen
```shell
go get -u github.com/2637309949/gmt/comm/autogen
```

- Init micro repo
```shell
autogen init t5-xk
```

- New service
```shell
cd t5-xk && autogen new t5-001 -s "root:gmt@tcp(127.0.0.1:3306)/test" -t "order"
```

```shell
$ ls --all
.  ..  comm  helloworld-service  proto  t5-001-service  workspace.code-workspace
```

### Document
- [编译环境](./docs/编译环境.md)
- [安装微服](./docs/安装微服.md)
- [编译网关](./docs/编译网关.md)
- [注册中心](./docs/注册中心.md)
- [全文索引](./docs/全文索引.md)
- [链路跟踪](./docs/链路跟踪.md)
- [消息订阅](./docs/消息订阅.md)
- [监控系统](./docs/监控系统.md)
- [数据平台](./docs/数据平台.md)
- [日志中心](./docs/日志中心.md)
- [运行应用](./docs/运行应用.md)
- [版本控制](./docs/版本控制.md)
- [版本迭代](./docs/版本迭代.md)
- [容器平台](./docs/容器平台.md)