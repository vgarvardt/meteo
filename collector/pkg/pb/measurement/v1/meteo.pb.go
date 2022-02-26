// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: measurement/v1/meteo.proto

package v1

import (
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

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time          *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	MeasurementId string                 `protobuf:"bytes,2,opt,name=measurement_id,json=measurementId,proto3" json:"measurement_id,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_measurement_v1_meteo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_measurement_v1_meteo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_measurement_v1_meteo_proto_rawDescGZIP(), []int{0}
}

func (x *Meta) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Meta) GetMeasurementId() string {
	if x != nil {
		return x.MeasurementId
	}
	return ""
}

type Climate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta                *Meta   `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Humidity            float32 `protobuf:"fixed32,2,opt,name=humidity,proto3" json:"humidity,omitempty"`
	Pressure            float32 `protobuf:"fixed32,3,opt,name=pressure,proto3" json:"pressure,omitempty"`
	TemperatureHumidity float32 `protobuf:"fixed32,4,opt,name=temperature_humidity,json=temperatureHumidity,proto3" json:"temperature_humidity,omitempty"`
	TemperaturePressure float32 `protobuf:"fixed32,5,opt,name=temperature_pressure,json=temperaturePressure,proto3" json:"temperature_pressure,omitempty"`
}

func (x *Climate) Reset() {
	*x = Climate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_measurement_v1_meteo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Climate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Climate) ProtoMessage() {}

func (x *Climate) ProtoReflect() protoreflect.Message {
	mi := &file_measurement_v1_meteo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Climate.ProtoReflect.Descriptor instead.
func (*Climate) Descriptor() ([]byte, []int) {
	return file_measurement_v1_meteo_proto_rawDescGZIP(), []int{1}
}

func (x *Climate) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Climate) GetHumidity() float32 {
	if x != nil {
		return x.Humidity
	}
	return 0
}

func (x *Climate) GetPressure() float32 {
	if x != nil {
		return x.Pressure
	}
	return 0
}

func (x *Climate) GetTemperatureHumidity() float32 {
	if x != nil {
		return x.TemperatureHumidity
	}
	return 0
}

func (x *Climate) GetTemperaturePressure() float32 {
	if x != nil {
		return x.TemperaturePressure
	}
	return 0
}

type System struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta           *Meta          `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	CpuTemperature float32        `protobuf:"fixed32,2,opt,name=cpu_temperature,json=cpuTemperature,proto3" json:"cpu_temperature,omitempty"`
	La             *System_LA     `protobuf:"bytes,3,opt,name=la,proto3" json:"la,omitempty"`
	Memory         *System_Memory `protobuf:"bytes,4,opt,name=memory,proto3" json:"memory,omitempty"`
	Disk           *System_Disk   `protobuf:"bytes,5,opt,name=disk,proto3" json:"disk,omitempty"`
}

func (x *System) Reset() {
	*x = System{}
	if protoimpl.UnsafeEnabled {
		mi := &file_measurement_v1_meteo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *System) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*System) ProtoMessage() {}

func (x *System) ProtoReflect() protoreflect.Message {
	mi := &file_measurement_v1_meteo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use System.ProtoReflect.Descriptor instead.
func (*System) Descriptor() ([]byte, []int) {
	return file_measurement_v1_meteo_proto_rawDescGZIP(), []int{2}
}

func (x *System) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *System) GetCpuTemperature() float32 {
	if x != nil {
		return x.CpuTemperature
	}
	return 0
}

func (x *System) GetLa() *System_LA {
	if x != nil {
		return x.La
	}
	return nil
}

func (x *System) GetMemory() *System_Memory {
	if x != nil {
		return x.Memory
	}
	return nil
}

func (x *System) GetDisk() *System_Disk {
	if x != nil {
		return x.Disk
	}
	return nil
}

type System_LA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min1  float32 `protobuf:"fixed32,1,opt,name=min1,proto3" json:"min1,omitempty"`
	Min5  float32 `protobuf:"fixed32,2,opt,name=min5,proto3" json:"min5,omitempty"`
	Min15 float32 `protobuf:"fixed32,3,opt,name=min15,proto3" json:"min15,omitempty"`
}

func (x *System_LA) Reset() {
	*x = System_LA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_measurement_v1_meteo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *System_LA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*System_LA) ProtoMessage() {}

func (x *System_LA) ProtoReflect() protoreflect.Message {
	mi := &file_measurement_v1_meteo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use System_LA.ProtoReflect.Descriptor instead.
func (*System_LA) Descriptor() ([]byte, []int) {
	return file_measurement_v1_meteo_proto_rawDescGZIP(), []int{2, 0}
}

func (x *System_LA) GetMin1() float32 {
	if x != nil {
		return x.Min1
	}
	return 0
}

func (x *System_LA) GetMin5() float32 {
	if x != nil {
		return x.Min5
	}
	return 0
}

func (x *System_LA) GetMin15() float32 {
	if x != nil {
		return x.Min15
	}
	return 0
}

type System_Memory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalKb     int64 `protobuf:"varint,1,opt,name=total_kb,json=totalKb,proto3" json:"total_kb,omitempty"`
	UsedKb      int64 `protobuf:"varint,2,opt,name=used_kb,json=usedKb,proto3" json:"used_kb,omitempty"`
	FreeKb      int64 `protobuf:"varint,3,opt,name=free_kb,json=freeKb,proto3" json:"free_kb,omitempty"`
	SharedKb    int64 `protobuf:"varint,4,opt,name=shared_kb,json=sharedKb,proto3" json:"shared_kb,omitempty"`
	CacheKb     int64 `protobuf:"varint,5,opt,name=cache_kb,json=cacheKb,proto3" json:"cache_kb,omitempty"`
	AvailableKb int64 `protobuf:"varint,6,opt,name=available_kb,json=availableKb,proto3" json:"available_kb,omitempty"`
}

func (x *System_Memory) Reset() {
	*x = System_Memory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_measurement_v1_meteo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *System_Memory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*System_Memory) ProtoMessage() {}

func (x *System_Memory) ProtoReflect() protoreflect.Message {
	mi := &file_measurement_v1_meteo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use System_Memory.ProtoReflect.Descriptor instead.
func (*System_Memory) Descriptor() ([]byte, []int) {
	return file_measurement_v1_meteo_proto_rawDescGZIP(), []int{2, 1}
}

func (x *System_Memory) GetTotalKb() int64 {
	if x != nil {
		return x.TotalKb
	}
	return 0
}

func (x *System_Memory) GetUsedKb() int64 {
	if x != nil {
		return x.UsedKb
	}
	return 0
}

func (x *System_Memory) GetFreeKb() int64 {
	if x != nil {
		return x.FreeKb
	}
	return 0
}

func (x *System_Memory) GetSharedKb() int64 {
	if x != nil {
		return x.SharedKb
	}
	return 0
}

func (x *System_Memory) GetCacheKb() int64 {
	if x != nil {
		return x.CacheKb
	}
	return 0
}

func (x *System_Memory) GetAvailableKb() int64 {
	if x != nil {
		return x.AvailableKb
	}
	return 0
}

type System_Disk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalKb     int64 `protobuf:"varint,1,opt,name=total_kb,json=totalKb,proto3" json:"total_kb,omitempty"`
	UsedKb      int64 `protobuf:"varint,2,opt,name=used_kb,json=usedKb,proto3" json:"used_kb,omitempty"`
	AvailableKb int64 `protobuf:"varint,3,opt,name=available_kb,json=availableKb,proto3" json:"available_kb,omitempty"`
	UsePrct     int32 `protobuf:"varint,4,opt,name=use_prct,json=usePrct,proto3" json:"use_prct,omitempty"`
}

func (x *System_Disk) Reset() {
	*x = System_Disk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_measurement_v1_meteo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *System_Disk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*System_Disk) ProtoMessage() {}

func (x *System_Disk) ProtoReflect() protoreflect.Message {
	mi := &file_measurement_v1_meteo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use System_Disk.ProtoReflect.Descriptor instead.
func (*System_Disk) Descriptor() ([]byte, []int) {
	return file_measurement_v1_meteo_proto_rawDescGZIP(), []int{2, 2}
}

func (x *System_Disk) GetTotalKb() int64 {
	if x != nil {
		return x.TotalKb
	}
	return 0
}

func (x *System_Disk) GetUsedKb() int64 {
	if x != nil {
		return x.UsedKb
	}
	return 0
}

func (x *System_Disk) GetAvailableKb() int64 {
	if x != nil {
		return x.AvailableKb
	}
	return 0
}

func (x *System_Disk) GetUsePrct() int32 {
	if x != nil {
		return x.UsePrct
	}
	return 0
}

var File_measurement_v1_meteo_proto protoreflect.FileDescriptor

var file_measurement_v1_meteo_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x74, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a,
	0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xd1, 0x01, 0x0a,
	0x07, 0x43, 0x6c, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x08, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x74, 0x65,
	0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x13, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x48, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x31, 0x0a,
	0x14, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x13, 0x74, 0x65, 0x6d,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x50, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65,
	0x22, 0xdf, 0x04, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x28, 0x0a, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x70, 0x75, 0x5f, 0x74, 0x65, 0x6d,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e,
	0x63, 0x70, 0x75, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x29,
	0x0a, 0x02, 0x6c, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65, 0x61,
	0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x4c, 0x41, 0x52, 0x02, 0x6c, 0x61, 0x12, 0x35, 0x0a, 0x06, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x12, 0x2f, 0x0a, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x04, 0x64, 0x69, 0x73,
	0x6b, 0x1a, 0x42, 0x0a, 0x02, 0x4c, 0x41, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x69, 0x6e, 0x31, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6d, 0x69, 0x6e, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x6d,
	0x69, 0x6e, 0x35, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6d, 0x69, 0x6e, 0x35, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x31, 0x35, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x6d, 0x69, 0x6e, 0x31, 0x35, 0x1a, 0xb0, 0x01, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6b, 0x62, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4b, 0x62, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x64, 0x5f, 0x6b, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x64, 0x4b, 0x62, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x6b, 0x62, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x72, 0x65, 0x65, 0x4b, 0x62, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x6b, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4b, 0x62, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x5f, 0x6b, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x4b, 0x62, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x6b, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4b, 0x62, 0x1a, 0x78, 0x0a, 0x04, 0x44, 0x69, 0x73, 0x6b,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6b, 0x62, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4b, 0x62, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x64, 0x5f, 0x6b, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x64, 0x4b, 0x62, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x6b, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4b, 0x62, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x5f, 0x70,
	0x72, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x75, 0x73, 0x65, 0x50, 0x72,
	0x63, 0x74, 0x42, 0x10, 0x5a, 0x0e, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_measurement_v1_meteo_proto_rawDescOnce sync.Once
	file_measurement_v1_meteo_proto_rawDescData = file_measurement_v1_meteo_proto_rawDesc
)

func file_measurement_v1_meteo_proto_rawDescGZIP() []byte {
	file_measurement_v1_meteo_proto_rawDescOnce.Do(func() {
		file_measurement_v1_meteo_proto_rawDescData = protoimpl.X.CompressGZIP(file_measurement_v1_meteo_proto_rawDescData)
	})
	return file_measurement_v1_meteo_proto_rawDescData
}

var file_measurement_v1_meteo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_measurement_v1_meteo_proto_goTypes = []interface{}{
	(*Meta)(nil),                  // 0: measurement.v1.Meta
	(*Climate)(nil),               // 1: measurement.v1.Climate
	(*System)(nil),                // 2: measurement.v1.System
	(*System_LA)(nil),             // 3: measurement.v1.System.LA
	(*System_Memory)(nil),         // 4: measurement.v1.System.Memory
	(*System_Disk)(nil),           // 5: measurement.v1.System.Disk
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_measurement_v1_meteo_proto_depIdxs = []int32{
	6, // 0: measurement.v1.Meta.time:type_name -> google.protobuf.Timestamp
	0, // 1: measurement.v1.Climate.meta:type_name -> measurement.v1.Meta
	0, // 2: measurement.v1.System.meta:type_name -> measurement.v1.Meta
	3, // 3: measurement.v1.System.la:type_name -> measurement.v1.System.LA
	4, // 4: measurement.v1.System.memory:type_name -> measurement.v1.System.Memory
	5, // 5: measurement.v1.System.disk:type_name -> measurement.v1.System.Disk
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_measurement_v1_meteo_proto_init() }
func file_measurement_v1_meteo_proto_init() {
	if File_measurement_v1_meteo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_measurement_v1_meteo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_measurement_v1_meteo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Climate); i {
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
		file_measurement_v1_meteo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*System); i {
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
		file_measurement_v1_meteo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*System_LA); i {
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
		file_measurement_v1_meteo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*System_Memory); i {
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
		file_measurement_v1_meteo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*System_Disk); i {
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
			RawDescriptor: file_measurement_v1_meteo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_measurement_v1_meteo_proto_goTypes,
		DependencyIndexes: file_measurement_v1_meteo_proto_depIdxs,
		MessageInfos:      file_measurement_v1_meteo_proto_msgTypes,
	}.Build()
	File_measurement_v1_meteo_proto = out.File
	file_measurement_v1_meteo_proto_rawDesc = nil
	file_measurement_v1_meteo_proto_goTypes = nil
	file_measurement_v1_meteo_proto_depIdxs = nil
}
