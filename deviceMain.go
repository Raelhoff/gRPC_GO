package main

import (
	"fmt"
	"log"
	"time"

	pb "github.com/Raelhoff/gRPC_GO/protofiles"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {

	// using the profo created struct
	p := &pb.Device{
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

	log.Println("\n\n")
	log.Println(time.Now())
	log.Println("\n\n")

	// Serializing the struct and assigning it to body
	body, _ := proto.Marshal(p)

	// De-serializing the body and saving it to p1 for testing
	p1 := &pb.Device{}

	_ = proto.Unmarshal(body, p1)

	fmt.Println("Original struct loaded from proto file:", p)
	fmt.Println("Marshalled proto data: ", body)
	fmt.Println("Unmarshalled struct: ", p1)

       fmt.Println("Sensors: ", p.Sensors[0].Type)
}
