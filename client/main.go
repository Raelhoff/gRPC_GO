package main

import (
    "fmt"
    "fiber-mongo-api/configs" //add this
    "fiber-mongo-api/routes" //add this
    "github.com/gofiber/fiber/v2" 
//xy	"golang.org/x/net/context"    
//    pb "github.com/Raelhoff/gRPC_GO/protofiles"
  //  "google.golang.org/grpc"
//	"google.golang.org/protobuf/types/known/timestamppb"
)


func main() {
    app := fiber.New()

    //run database
    configs.ConnectDB()

    // run gRPC
    configs.ConnectGRPC()
    
  //  var conn * grpc.ClientConn  = configs.ReturnClientGRPC()
   // client := pb.NewLoraTransactionClient(conn)
    
    routes.UserRoute(app) //add this
    routes.DeviceRoute(app)

    app.Listen(":3033")
    fmt.Println("Http Server - Listen: 3033")

}
