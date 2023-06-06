package main

import (
	"fmt"
	"log"
	"net"
	"time"
        "strconv"  
	pb "github.com/Raelhoff/gRPC_GO/protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/reflection"
       	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
        "os"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/google/uuid"
)

type server struct{}


func main() {
/*       	
wallet, err := gateway.NewFileSystemWallet("./wallets")
	if err != nil {
		fmt.Printf("Failed to create wallet: %s\n", err)
		os.Exit(1)
	}

	if !wallet.Exists("Admin") {
		fmt.Println("Failed to get Admin from wallet")
		os.Exit(1)
	}

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile("./connection.json")),
		gateway.WithIdentity(wallet, "Admin"),
	)

	if err != nil {
		fmt.Printf("Failed to connect: %v", err)
	}

	if gw == nil {
		fmt.Println("Failed to create gateway")
	}

*/
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

func String(n int32) string {
    buf := [11]byte{}
    pos := len(buf)
    i := int64(n)
    signed := i < 0
    if signed {
        i = -i
    }
    for {
        pos--
        buf[pos], i = '0'+byte(i%10), i/10
        if i == 0 {
            if signed {
                pos--
                buf[pos] = '-'
            }
            return string(buf[pos:])
        }
    }
}


func ReadWallet() *gateway.Wallet  {

      w, err := gateway.NewFileSystemWallet("./wallets")
     if err != nil {
                fmt.Printf("Failed to create wallet: %s\n", err)
                os.Exit(1)
     }

     if !w.Exists("Admin") {
                fmt.Println("Failed to get Admin from wallet")
                os.Exit(1)
     }

    fmt.Println("ReadWallet to wallets")
    return w
}

//Client instance
var wallet *gateway.Wallet = ReadWallet()                                                                                                                           

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


// wallet, err := gateway.NewFileSystemWallet("./wallets")
  //      if err != nil {
    //            fmt.Printf("Failed to create wallet: %s\n", err)
     //           os.Exit(1)
      //  }

        //if !wallet.Exists("Admin") {
          //      fmt.Println("Failed to get Admin from wallet")
            //    os.Exit(1)
   //     }

        gw, err := gateway.Connect(
                gateway.WithConfig(config.FromFile("./connection.json")),
                gateway.WithIdentity(wallet, "Admin"),
        )

        if err != nil {
                fmt.Printf("Failed to connect: %v", err)
        }

        if gw == nil {
                fmt.Println("Failed to create gateway")
        }


	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		fmt.Printf("Failed to get network: %v", err)
	}

	//var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	contract := network.GetContract("smart")
	uuid.SetRand(nil)

// invoke
	fmt.Println("\n\n --- invoke ---")
	var temp float32 = float32(in.Sensors[0].Value)
	var humid float32 = float32(in.Sensors[1].Value)
	fmt.Println(String (int32(temp)))
//	fmt.Println(humid)


//	var stemp string = strconv.FormatFloat(fmt.Sprintf(temp), 'E', -1, 32)

	t := time.Now()
	fmt.Println(t.String())

	resultInvoke, err := contract.SubmitTransaction("put", String(in.Id), String( in.Input1),  String( in.Input2),  String( in.Output), strconv.FormatBool(in.AlarmBattery), strconv.FormatBool(in.AlarmPower), strconv.FormatBool(in.SensorError), String (int32(temp)), String (int32(humid)), t.String(), "0")
	if err != nil {
		fmt.Printf("Failed to commit transaction: %v", err)
	} else {
		fmt.Println("Commit is successful - invoke")
		fmt.Println(string(resultInvoke))
	}

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
