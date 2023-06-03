package configs

import (
    	"fmt"
	"google.golang.org/grpc"
	pb "github.com/Raelhoff/gRPC_GO/protofiles"
)

// grpc server address
const address = "192.168.0.192:8000"

func ConnectGRPC() *grpc.ClientConn  {
    client, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
		fmt.Println("Error while making connection")
    }

    fmt.Println("Connected to GRPC")
    return client
}

//Client instance
var conn  = ConnectGRPC()

func ReturnClientGRPC() *grpc.ClientConn  {
   return conn
}
//getting database collections
func GetProtofileLoraTransaction()    pb.LoraTransactionClient {
    c := pb.NewLoraTransactionClient(conn)
    if c == nil {
         fmt.Println("Error GetProtofileLoraTransaction")
    }
    return  c
}
