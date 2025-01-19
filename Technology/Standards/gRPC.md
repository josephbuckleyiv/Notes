
## Proto Definition
Think JSON, GraphQL. A format for declaring data structure -- in this case, transpiled to your language of choice [potentially annoying]. Here is an example:

```
syntax = "proto3"
package notificationproto;
option go_package = "grpcstreaming/notificationproto"

service NotificationService {
  rpc GetNotifications(NotificationRequest) returns (stream Notificaiton) {}
}

message NotificationRequest {
  string user_id string = 1;

message Notifcation {
  string user_id = 1;
  string content = 2;
  int64 created_at = 3;
}
```

Stream of notification objects returned upon receving NotificationRequest --> this is defined by the NotificationService definition.

## Starting a server in Go
```
func main() {
  lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", notificationservice.Address [port]))
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  var opts []grpc.ServerOption

  grppcServer := grpc.NewServer(opts...)

  # pb -> handler
  # Redisclient = grpcStreaming.NewRedisClient()
  # handler = notificationService.NewHandler(redisClient)
  pb.RegiseterRouteGuideServer(grpcServer, newServer())
  grpcServer.Serve(lis)
}
```

## Streaming Architecture
Should be a Client, Server and Publisher. Server
