# v-stream
video streaming server

# 環境構築
## 依存パッケージのインストール
`protoc`コマンドのインストール
```
$ brew install protobuf
$ which protoc
```

gRPCの利用に必要なGoパッケージのインストール
```
$ go get -u google.golang.org/grpc
$ go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## protoファイルからコードを自動生成する
Protocol Buffersの生成
```
$ protoc -I=./api/proto \
--go_out=:./pkg/pb \     
--go_opt=module=github.com/Be3751/v-stream/pkg/pb \     
api/proto/stream.proto
```

gRPCの生成
```
$ protoc -I=./api/proto \
--go-grpc_out=:./pkg/pb \
--go-grpc_opt=module=github.com/Be3751/v-stream/pkg/pb \
api/proto/stream.proto
```
