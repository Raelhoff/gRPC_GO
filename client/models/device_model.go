package models

import "go.mongodb.org/mongo-driver/bson/primitive"
import "time"

/*
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

type Sensors struct {
    Type   int32                `bson:"type" json:"type" ` 
    Value   float32               `bson:"value" json:"value"` 
}

type Device struct {
    InsertedID    primitive.ObjectID `json:"InsertedID"`
    IdDevice int32                `json:"id" validate:"required"`
    Input1   int32                `json:"input1, Input1"` 
    Input2   int32                `json:"input2, Input2"` 
    Output   int32                `json:"output, Output"` 
    Alarm_battery   bool        `json:"alarm_battery, Alarm_battery"` 
    Alarm_power     bool        `json:"alarm_power, Alarm_power" `    
    Sensor_error    bool        `json:"sensor_error, Sensor_error" ` 
    Sensors_obj    []Sensors    `bson:"sensors,Sensors" json:"sensors, Sensors"`
    LastUpdated    time.Time    `bson:"lastUpdated" json:"lastUpdated"`
}
