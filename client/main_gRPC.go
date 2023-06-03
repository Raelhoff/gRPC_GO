package main

import (
	"log"
	"fmt"

	pb "github.com/Raelhoff/gRPC_GO/protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)


func main() {
	// Set up connection with the grpc server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while making connection, %v", err)
	}

	// Create a client instance
	c := pb.NewLoraTransactionClient(conn)

	// Lets invoke the remote function from client on the server
	tx, err := c.MakeTransaction(
		context.Background(),
		&pb.LoraRequest{
			IdDevice:     17081990,
			Id:           1234,
			Input1:       0,
			Input2:       0,
			Output:       0,
			AlarmBattery: false,
			AlarmPower:   false,
			SensorError:  false,
			Sensors: []*pb.LoraRequest_Sensor{
				{Type: 1, Value: 30},
				{Type: 2, Value: 15},
			},
			LastUpdated: timestamppb.Now(),
		},
	)

 	if err != nil {
        	log.Fatalf("Error, %v", err)
   	 }
	
	log.Fatalf("Response: %v", tx)
        fmt.Println("Msg: ", tx.Msg)
        fmt.Println("Confirmation: ", tx.Confirmation)

	// Create a client instance
	d := pb.NewListLoraTransactionClient(conn)

	d.MakeListTransaction(
		context.Background(),
		&pb.ListLoraRequest{
			ListDevice: []*pb.LoraRequest{
				{
					IdDevice:     10011990,
					Id:           1234,
					Input1:       0,
					Input2:       0,
					Output:       0,
					AlarmBattery: false,
					AlarmPower:   false,
					SensorError:  false,
					Sensors: []*pb.LoraRequest_Sensor{
						{Type: 1, Value: 30},
						{Type: 2, Value: 15},
					},
					LastUpdated: timestamppb.Now(),
				}, {
					IdDevice:     15011990,
					Id:           1232334,
					Input1:       0,
					Input2:       0,
					Output:       0,
					AlarmBattery: false,
					AlarmPower:   false,
					SensorError:  true,
					Sensors: []*pb.LoraRequest_Sensor{
						{Type: 1, Value: 30},
						{Type: 2, Value: 15},
					},
					LastUpdated: timestamppb.Now(),
				},
			},
		},
	)
}