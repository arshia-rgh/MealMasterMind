### Meal Planning (Golang)

#### gRpc:

use this command to add newest generated gRPC code

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    grpc/user/user.proto
```