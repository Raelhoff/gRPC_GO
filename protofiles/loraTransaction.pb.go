// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: loraTransaction.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoraRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdDevice     int32                  `protobuf:"varint,1,opt,name=idDevice,proto3" json:"idDevice,omitempty"`
	Id           int32                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Input1       int32                  `protobuf:"varint,3,opt,name=input1,proto3" json:"input1,omitempty"`
	Input2       int32                  `protobuf:"varint,4,opt,name=input2,proto3" json:"input2,omitempty"`
	Output       int32                  `protobuf:"varint,5,opt,name=output,proto3" json:"output,omitempty"`
	AlarmBattery bool                   `protobuf:"varint,6,opt,name=alarm_battery,json=alarmBattery,proto3" json:"alarm_battery,omitempty"`
	AlarmPower   bool                   `protobuf:"varint,7,opt,name=alarm_power,json=alarmPower,proto3" json:"alarm_power,omitempty"`
	SensorError  bool                   `protobuf:"varint,8,opt,name=sensor_error,json=sensorError,proto3" json:"sensor_error,omitempty"`
	Sensors      []*LoraRequest_Sensor  `protobuf:"bytes,9,rep,name=sensors,proto3" json:"sensors,omitempty"`
	LastUpdated  *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
}

func (x *LoraRequest) Reset() {
	*x = LoraRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loraTransaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoraRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoraRequest) ProtoMessage() {}

func (x *LoraRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loraTransaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoraRequest.ProtoReflect.Descriptor instead.
func (*LoraRequest) Descriptor() ([]byte, []int) {
	return file_loraTransaction_proto_rawDescGZIP(), []int{0}
}

func (x *LoraRequest) GetIdDevice() int32 {
	if x != nil {
		return x.IdDevice
	}
	return 0
}

func (x *LoraRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LoraRequest) GetInput1() int32 {
	if x != nil {
		return x.Input1
	}
	return 0
}

func (x *LoraRequest) GetInput2() int32 {
	if x != nil {
		return x.Input2
	}
	return 0
}

func (x *LoraRequest) GetOutput() int32 {
	if x != nil {
		return x.Output
	}
	return 0
}

func (x *LoraRequest) GetAlarmBattery() bool {
	if x != nil {
		return x.AlarmBattery
	}
	return false
}

func (x *LoraRequest) GetAlarmPower() bool {
	if x != nil {
		return x.AlarmPower
	}
	return false
}

func (x *LoraRequest) GetSensorError() bool {
	if x != nil {
		return x.SensorError
	}
	return false
}

func (x *LoraRequest) GetSensors() []*LoraRequest_Sensor {
	if x != nil {
		return x.Sensors
	}
	return nil
}

func (x *LoraRequest) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

type LoraResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg          string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Confirmation bool   `protobuf:"varint,2,opt,name=confirmation,proto3" json:"confirmation,omitempty"`
}

func (x *LoraResponse) Reset() {
	*x = LoraResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loraTransaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoraResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoraResponse) ProtoMessage() {}

func (x *LoraResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loraTransaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoraResponse.ProtoReflect.Descriptor instead.
func (*LoraResponse) Descriptor() ([]byte, []int) {
	return file_loraTransaction_proto_rawDescGZIP(), []int{1}
}

func (x *LoraResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *LoraResponse) GetConfirmation() bool {
	if x != nil {
		return x.Confirmation
	}
	return false
}

type ListLoraRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListDevice []*LoraRequest `protobuf:"bytes,1,rep,name=listDevice,proto3" json:"listDevice,omitempty"`
}

func (x *ListLoraRequest) Reset() {
	*x = ListLoraRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loraTransaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLoraRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLoraRequest) ProtoMessage() {}

func (x *ListLoraRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loraTransaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLoraRequest.ProtoReflect.Descriptor instead.
func (*ListLoraRequest) Descriptor() ([]byte, []int) {
	return file_loraTransaction_proto_rawDescGZIP(), []int{2}
}

func (x *ListLoraRequest) GetListDevice() []*LoraRequest {
	if x != nil {
		return x.ListDevice
	}
	return nil
}

type ListLoraResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg          string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Confirmation bool   `protobuf:"varint,2,opt,name=confirmation,proto3" json:"confirmation,omitempty"`
}

func (x *ListLoraResponse) Reset() {
	*x = ListLoraResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loraTransaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLoraResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLoraResponse) ProtoMessage() {}

func (x *ListLoraResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loraTransaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLoraResponse.ProtoReflect.Descriptor instead.
func (*ListLoraResponse) Descriptor() ([]byte, []int) {
	return file_loraTransaction_proto_rawDescGZIP(), []int{3}
}

func (x *ListLoraResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ListLoraResponse) GetConfirmation() bool {
	if x != nil {
		return x.Confirmation
	}
	return false
}

type LoraRequest_Sensor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  int32   `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Value float32 `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *LoraRequest_Sensor) Reset() {
	*x = LoraRequest_Sensor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loraTransaction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoraRequest_Sensor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoraRequest_Sensor) ProtoMessage() {}

func (x *LoraRequest_Sensor) ProtoReflect() protoreflect.Message {
	mi := &file_loraTransaction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoraRequest_Sensor.ProtoReflect.Descriptor instead.
func (*LoraRequest_Sensor) Descriptor() ([]byte, []int) {
	return file_loraTransaction_proto_rawDescGZIP(), []int{0, 0}
}

func (x *LoraRequest_Sensor) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *LoraRequest_Sensor) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_loraTransaction_proto protoreflect.FileDescriptor

var file_loraTransaction_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6c, 0x6f, 0x72, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x03, 0x0a, 0x0b, 0x4c, 0x6f, 0x72, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x32,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6c, 0x61, 0x72,
	0x6d, 0x5f, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x42, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x38, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e,
	0x4c, 0x6f, 0x72, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x52, 0x07, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x12, 0x3c, 0x0a, 0x0b, 0x6c,
	0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61,
	0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x32, 0x0a, 0x06, 0x53, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x44, 0x0a,
	0x0c, 0x4c, 0x6f, 0x72, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x4a, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x72, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x72, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22,
	0x48, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x72, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x59, 0x0a, 0x0f, 0x4c, 0x6f, 0x72,
	0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x0f,
	0x4d, 0x61, 0x6b, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x72,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x72, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x32, 0x69, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x72, 0x61,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x52, 0x0a, 0x13, 0x4d,
	0x61, 0x6b, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x72, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4c, 0x6f, 0x72, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loraTransaction_proto_rawDescOnce sync.Once
	file_loraTransaction_proto_rawDescData = file_loraTransaction_proto_rawDesc
)

func file_loraTransaction_proto_rawDescGZIP() []byte {
	file_loraTransaction_proto_rawDescOnce.Do(func() {
		file_loraTransaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_loraTransaction_proto_rawDescData)
	})
	return file_loraTransaction_proto_rawDescData
}

var file_loraTransaction_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_loraTransaction_proto_goTypes = []interface{}{
	(*LoraRequest)(nil),           // 0: protofiles.LoraRequest
	(*LoraResponse)(nil),          // 1: protofiles.LoraResponse
	(*ListLoraRequest)(nil),       // 2: protofiles.ListLoraRequest
	(*ListLoraResponse)(nil),      // 3: protofiles.ListLoraResponse
	(*LoraRequest_Sensor)(nil),    // 4: protofiles.LoraRequest.Sensor
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_loraTransaction_proto_depIdxs = []int32{
	4, // 0: protofiles.LoraRequest.sensors:type_name -> protofiles.LoraRequest.Sensor
	5, // 1: protofiles.LoraRequest.lastUpdated:type_name -> google.protobuf.Timestamp
	0, // 2: protofiles.ListLoraRequest.listDevice:type_name -> protofiles.LoraRequest
	0, // 3: protofiles.LoraTransaction.MakeTransaction:input_type -> protofiles.LoraRequest
	2, // 4: protofiles.ListLoraTransaction.MakeListTransaction:input_type -> protofiles.ListLoraRequest
	1, // 5: protofiles.LoraTransaction.MakeTransaction:output_type -> protofiles.LoraResponse
	3, // 6: protofiles.ListLoraTransaction.MakeListTransaction:output_type -> protofiles.ListLoraResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_loraTransaction_proto_init() }
func file_loraTransaction_proto_init() {
	if File_loraTransaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loraTransaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoraRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_loraTransaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoraResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_loraTransaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLoraRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_loraTransaction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLoraResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_loraTransaction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoraRequest_Sensor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_loraTransaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_loraTransaction_proto_goTypes,
		DependencyIndexes: file_loraTransaction_proto_depIdxs,
		MessageInfos:      file_loraTransaction_proto_msgTypes,
	}.Build()
	File_loraTransaction_proto = out.File
	file_loraTransaction_proto_rawDesc = nil
	file_loraTransaction_proto_goTypes = nil
	file_loraTransaction_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LoraTransactionClient is the client API for LoraTransaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoraTransactionClient interface {
	MakeTransaction(ctx context.Context, in *LoraRequest, opts ...grpc.CallOption) (*LoraResponse, error)
}

type loraTransactionClient struct {
	cc grpc.ClientConnInterface
}

func NewLoraTransactionClient(cc grpc.ClientConnInterface) LoraTransactionClient {
	return &loraTransactionClient{cc}
}

func (c *loraTransactionClient) MakeTransaction(ctx context.Context, in *LoraRequest, opts ...grpc.CallOption) (*LoraResponse, error) {
	out := new(LoraResponse)
	err := c.cc.Invoke(ctx, "/protofiles.LoraTransaction/MakeTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoraTransactionServer is the server API for LoraTransaction service.
type LoraTransactionServer interface {
	MakeTransaction(context.Context, *LoraRequest) (*LoraResponse, error)
}

// UnimplementedLoraTransactionServer can be embedded to have forward compatible implementations.
type UnimplementedLoraTransactionServer struct {
}

func (*UnimplementedLoraTransactionServer) MakeTransaction(context.Context, *LoraRequest) (*LoraResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeTransaction not implemented")
}

func RegisterLoraTransactionServer(s *grpc.Server, srv LoraTransactionServer) {
	s.RegisterService(&_LoraTransaction_serviceDesc, srv)
}

func _LoraTransaction_MakeTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoraRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoraTransactionServer).MakeTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.LoraTransaction/MakeTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoraTransactionServer).MakeTransaction(ctx, req.(*LoraRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoraTransaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protofiles.LoraTransaction",
	HandlerType: (*LoraTransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeTransaction",
			Handler:    _LoraTransaction_MakeTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loraTransaction.proto",
}

// ListLoraTransactionClient is the client API for ListLoraTransaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ListLoraTransactionClient interface {
	MakeListTransaction(ctx context.Context, in *ListLoraRequest, opts ...grpc.CallOption) (*ListLoraResponse, error)
}

type listLoraTransactionClient struct {
	cc grpc.ClientConnInterface
}

func NewListLoraTransactionClient(cc grpc.ClientConnInterface) ListLoraTransactionClient {
	return &listLoraTransactionClient{cc}
}

func (c *listLoraTransactionClient) MakeListTransaction(ctx context.Context, in *ListLoraRequest, opts ...grpc.CallOption) (*ListLoraResponse, error) {
	out := new(ListLoraResponse)
	err := c.cc.Invoke(ctx, "/protofiles.ListLoraTransaction/MakeListTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListLoraTransactionServer is the server API for ListLoraTransaction service.
type ListLoraTransactionServer interface {
	MakeListTransaction(context.Context, *ListLoraRequest) (*ListLoraResponse, error)
}

// UnimplementedListLoraTransactionServer can be embedded to have forward compatible implementations.
type UnimplementedListLoraTransactionServer struct {
}

func (*UnimplementedListLoraTransactionServer) MakeListTransaction(context.Context, *ListLoraRequest) (*ListLoraResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeListTransaction not implemented")
}

func RegisterListLoraTransactionServer(s *grpc.Server, srv ListLoraTransactionServer) {
	s.RegisterService(&_ListLoraTransaction_serviceDesc, srv)
}

func _ListLoraTransaction_MakeListTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLoraRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListLoraTransactionServer).MakeListTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.ListLoraTransaction/MakeListTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListLoraTransactionServer).MakeListTransaction(ctx, req.(*ListLoraRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ListLoraTransaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protofiles.ListLoraTransaction",
	HandlerType: (*ListLoraTransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeListTransaction",
			Handler:    _ListLoraTransaction_MakeListTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loraTransaction.proto",
}
