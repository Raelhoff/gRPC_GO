package controllers

import (
    "context"
    "fiber-mongo-api/configs"
    "fiber-mongo-api/models"
    "fiber-mongo-api/responses"
    "net/http"
    "time"
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson" //add this
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson/primitive"
//     "fiber-mongo-api/configs" //add this
    "google.golang.org/grpc"
    pb "github.com/Raelhoff/gRPC_GO/protofiles"
    "fmt"	
  //  "log"
    "google.golang.org/protobuf/types/known/timestamppb"
)

var deviceCollection *mongo.Collection = configs.GetCollection(configs.DB, "devices")
var validateDevice = validator.New()
// comm * grpc.ClientConn  = configs.ReturnClientGRPC()



 
// grpc server address

func CreateDevice(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var device models.Device
    defer cancel()

    //validate the request body
    if err := c.BodyParser(&device); err != nil {
        return c.Status(http.StatusBadRequest).JSON(responses.DeviceResponse{Status: http.StatusBadRequest, Message: "error: BodyParser", Data: &fiber.Map{"data": err.Error()}})
    }

    //use the validator library to validate required fields
    if validationErr := validateDevice.Struct(&device); validationErr != nil {
        return c.Status(http.StatusBadRequest).JSON(responses.DeviceResponse{Status: http.StatusBadRequest, Message: "error: validation", Data: &fiber.Map{"data": validationErr.Error()}})
    }

     var conn * grpc.ClientConn  = configs.ReturnClientGRPC()
     client := pb.NewLoraTransactionClient(conn)
    
    if client  != nil {
        fmt.Println("Client ok!!")
   


    // Lets invoke the remote function from client on the server
   tx, err := client.MakeTransaction(
	context.Background(),
	&pb.LoraRequest{
		IdDevice:     1234,
		Id:           device.IdDevice,
		Input1:       device.Input1,
		Input2:       device.Input2,
		Output:       device.Output,
		AlarmBattery: device.Alarm_battery,
		AlarmPower:   device.Alarm_power,
		SensorError:  device.Sensor_error,
		Sensors:   []*pb.LoraRequest_Sensor{
			{Type: (device.Sensors_obj[0].Type + 1), Value: device.Sensors_obj[0].Value },
                        {Type: (device.Sensors_obj[1].Type + 1), Value: device.Sensors_obj[1].Value },
			}, 
		LastUpdated: timestamppb.Now(),
	},
     )

    if err != nil {
        fmt.Println("Error, %v", err)
    }

   fmt.Println("Msg: ", tx.Msg)
   fmt.Println("Confirmation: ", tx.Confirmation)
}
/*
//defer conn.Close()
   
  
    // fmt.Println("Sensors: ", device.Sensors_obj[0].Type) 

   if err != nil {
       	log.Fatalf("Error, %v", err)
   }
	
   log.Fatalf("Response: %v", tx)
   fmt.Println("Msg: ", tx.Msg)
   fmt.Println("Confirmation: ", tx.Confirmation) 
*/
/*

p.Sensors[0].Type

{ version: 111,
  id: 1459633408,
  timestamp: 1685407036,
  input1: 1,
  input2: 1,
  output: 0,
  alarm_battery: false,
  alarm_power: false,
  sensor_error: false,
  sensors: [ { type: 0, value: 24.870001 }, { type: 1, value: 65.970001 } ] }
*/
//   if tx.Confirmation == false {
    fmt.Println("newDevice")
    newDevice := models.Device{
        InsertedID:       primitive.NewObjectID(),
        IdDevice:     device.IdDevice,
        Input1:    device.Input1,
        Input2:    device.Input2,
        Output:    device.Output,
        Alarm_battery:    device.Alarm_battery,
        Alarm_power:    device.Alarm_power,
        Sensor_error:    device.Sensor_error,
        LastUpdated: time.Now(),
        Sensors_obj:     device.Sensors_obj,
    }

 fmt.Println("deviceCollection")

    result, err := deviceCollection.InsertOne(ctx, newDevice)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(responses.DeviceResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }

     fmt.Println("c.Status")

    return c.Status(http.StatusOK).JSON(responses.DeviceResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": result}})
  // }
  // return c.Status(http.StatusOK).JSON(responses.DeviceResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": nil}})
}

func GetAllDevices(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var devices []models.Device
    defer cancel()

    results, err := deviceCollection.Find(ctx, bson.M{})

    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(responses.DeviceResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }

    //reading from the db in an optimal way
    defer results.Close(ctx)
    for results.Next(ctx) {
        var singleDevice models.Device
        if err = results.Decode(&singleDevice); err != nil {
            return c.Status(http.StatusInternalServerError).JSON(responses.DeviceResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
        }

        devices = append(devices, singleDevice)
    }

    return c.Status(http.StatusOK).JSON(
        responses.DeviceResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": devices}},
    )
}
