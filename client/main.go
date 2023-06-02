package main

import (
	"log"

	pb "github.com/Raelhoff/GoProto/protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// grpc server address
const address = "localhost:8000"

func main() {
	// Set up connection with the grpc server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while making connection, %v", err)
	}

	// Create a client instance
	c := pb.NewLoraTransactionClient(conn)

	// Lets invoke the remote function from client on the server
	c.MakeTransaction(
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
