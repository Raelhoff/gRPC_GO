package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/Raelhoff/GoProto/protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/reflection"
)

type server struct{}

func main() {

	// NewServer creates a gRPC server which has no service registered and has not started
	// to accept requests yet.
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// We are making use of the function that compiled proto made for us to register
	// our GRPC server so that the clients can make use of the functions tide to our
	// server remotely via the GRPC server (like MakeTransaction function)

	// The first argument is the grpc server instance
	// The second argument is the service who's methods we want to expose (in our case)
	// we have put it in this program only
	pb.RegisterLoraTransactionServer(s, &server{})
	pb.RegisterListLoraTransactionServer(s, &server{})

	// Serve accepts incoming connections on the listener lis, creating a new ServerTransport
	// and service goroutine for each. The service goroutines read gRPC requests and then
	// call the registered handlers to reply to them.
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

// [ctx] is used by the goroutines to interact with GRPC
// [in] is the type of TransactionRequest
/*
	This function signature matches the service that we mentioned in the protobuf
*/
func (s *server) MakeTransaction(ctx context.Context, in *pb.LoraRequest) (*pb.LoraResponse, error) {
	/*
		IdDevice:     1234,
			Id:           456,
			Input1:       0,
			Input2:       0,
			Output:       0,
			AlarmBattery: true,
			AlarmPower:   true,
			SensorError:  true,
			Sensors: []*pb.Device_Sensor{
				{Type: 1, Value: 30},
				{Type: 2, Value: 15},
			},
			LastUpdated: timestamppb.Now(),
		}
	*/
	// Business logic will come here
	fmt.Println("IdDevice: ", in.IdDevice)
	fmt.Println("id: ", in.Id)
	fmt.Println("Input1: ", in.Input1)
	fmt.Println("Input2: ", in.Input2)
	fmt.Println("Output: ", in.Output)
	fmt.Println("AlarmBattery: ", in.AlarmBattery)
	fmt.Println("AlarmPower: ", in.AlarmPower)
	fmt.Println("SensorError: ", in.SensorError)
	fmt.Println("Sensors: ", in.Sensors)
	fmt.Println("LastUpdated: ", in.LastUpdated)

	// Returning a response of type Transaction Response
	return &pb.LoraResponse{Confirmation: true}, nil
}

func (s *server) MakeListTransaction(ctx context.Context, in *pb.ListLoraRequest) (*pb.ListLoraResponse, error) {
	/*
		IdDevice:     1234,
			Id:           456,
			Input1:       0,
			Input2:       0,
			Output:       0,
			AlarmBattery: true,
			AlarmPower:   true,
			SensorError:  true,
			Sensors: []*pb.Device_Sensor{
				{Type: 1, Value: 30},
				{Type: 2, Value: 15},
			},
			LastUpdated: timestamppb.Now(),
		}
	*/
	// Business logic will come here

	fmt.Println("\n\n ------ List: ", len(in.ListDevice))
	for i := 0; i < len(in.ListDevice); i++ {
		fmt.Println("IdDevice: ", in.ListDevice[i].IdDevice)
		fmt.Println("id: ", in.ListDevice[i].Id)
		fmt.Println("Input1: ", in.ListDevice[i].Input1)
		fmt.Println("Input2: ", in.ListDevice[i].Input2)
		fmt.Println("Output: ", in.ListDevice[i].Output)
		fmt.Println("AlarmBattery: ", in.ListDevice[i].AlarmBattery)
		fmt.Println("AlarmPower: ", in.ListDevice[i].AlarmPower)
		fmt.Println("SensorError: ", in.ListDevice[i].SensorError)
		fmt.Println("Sensors: ", in.ListDevice[i].Sensors)
		fmt.Println("LastUpdated: ", in.ListDevice[i].LastUpdated)

	}

	// Returning a response of type Transaction Response
	return &pb.ListLoraResponse{Confirmation: true}, nil
}
