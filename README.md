# symgw

symverse gateway & proxy 제어 역할을 하는 서비스입니다.

## Getting Started

1. go build 혹은 바이너리 파일을 직접 다운 받습니다.

2. cli 명령어를 통해 실행 합니다.

`./symgw --env {envFile}`

### CLI Options

help 명령어를 통해 확인하실수 있습니다.

```
./symgw --help

NAME:
   symgw - symverse gateway server

USAGE:
   symgw [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --env FILE, -E FILE  Load configuration from FILE (default: ".env")
   --help, -h           show help
   --version, -v        print the version

```

### env file option

symgw 실행시 config 파일을 통해 proxy 할 workNode 정보와 bootNode 정보를 입력합니다. 

config 파일은 `yaml` 형태로 저장해야 하며 아래와 같은 예제로 작성해야 합니다.

```
// example config.yaml 

database:
  driver: "leveldb" ( default: leveldb ) // leveldb, redis
cache: 
  interval: 15 ( default: 5 ) // 분단위
  use: false // (default: false
host:
  address: "0.0.0.0" ( default: localhost )
  port: 80 ( default: 8080 )
bootNodes:
  - httpUrl: "http://10.100.1.199:9999" 
workNodes: 
  - httpUrl: "http:///127.0.0.1:8545"
    wsUrl: "http:///127.0.0.1:8546"
  - httpUrl: "http:///10.100.1.244:8545"
    wsUrl: "http:///10.100.1.244:8546"
```

- `database.driver` - api request 결과의 cache를 저장할 공간입니다.
- `cache.interval` - cache가 만료되는 시간입니다. ( 분단위 )
- `host.address` - symgw app 의 http Listen 할 host 입니다.
- `host.port` - symgw app 의 http Listen 할 port 입니다.
- `bootNodes.httpUrl` - bootnode를 proxy 할 노드 rpc addr 입니다.
- `workNodes.httpUrl` - gsym node를 proxy 할 노드 rpc addr 입니다.
- `workNodes.wsUrl` - gsym node를 proxy 할 노드 ws rpc addr 입니다.


## Node Api & Bootnode Api Proxy list

symgw api 목록 입니다.

 `GET` /v1/rpc/nodes - env 파일에 저장된 workNodes의 전체 url 정보를 가져옵니다.
 
 `POST` /v1/rpc/node/:number - env 파일에 저장된 workNodes 로 RPC proxy 합니다.
 
 `POST` /v1/rpc/node/:number/ws - env 파일에 저장된 workNodes 를 WS RPC proxy 합니다.
 
 `GET` /v1/bootnode/nodes - env 파일에 저장된 bootNodes의 url을 통해 API를 통해 `getNode` method 를 호출합니다.
 
 `GET` /v1/bootnode/closeNodes - env 파일에 저장된 bootNodes의 url을 통해 API를 통해 `closeNodes` method 를 호출합니다.


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
