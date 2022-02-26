// Copyright 2016-18, OpenCensus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: opencensus/proto/stats/v1/stats.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Measure_Type int32

const (
	// Unknown type.
	Measure_TYPE_UNSPECIFIED Measure_Type = 0
	// Indicates an int64 Measure.
	Measure_INT64 Measure_Type = 1
	// Indicates a double Measure.
	Measure_DOUBLE Measure_Type = 2
)

// Enum value maps for Measure_Type.
var (
	Measure_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "INT64",
		2: "DOUBLE",
	}
	Measure_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"INT64":            1,
		"DOUBLE":           2,
	}
)

func (x Measure_Type) Enum() *Measure_Type {
	p := new(Measure_Type)
	*p = x
	return p
}

func (x Measure_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Measure_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_opencensus_proto_stats_v1_stats_proto_enumTypes[0].Descriptor()
}

func (Measure_Type) Type() protoreflect.EnumType {
	return &file_opencensus_proto_stats_v1_stats_proto_enumTypes[0]
}

func (x Measure_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Measure_Type.Descriptor instead.
func (Measure_Type) EnumDescriptor() ([]byte, []int) {
	return file_opencensus_proto_stats_v1_stats_proto_rawDescGZIP(), []int{1, 0}
}

// TODO(bdrutu): Consider if this should be moved to a "tags" directory to match the API structure.
type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opencensus_proto_stats_v1_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_opencensus_proto_stats_v1_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_opencensus_proto_stats_v1_stats_proto_rawDescGZIP(), []int{0}
}

func (x *Tag) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Tag) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Measure .
type Measure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A string by which the measure will be referred to, e.g. "rpc_server_latency". Names MUST be
	// unique within the library.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Describes the measure, e.g. "RPC latency in seconds".
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Describes the unit used for the Measure. Follows the format described by
	// http://unitsofmeasure.org/ucum.html.
	Unit string `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
	// The type used for this Measure.
	Type Measure_Type `protobuf:"varint,4,opt,name=type,proto3,enum=opencensus.proto.stats.v1.Measure_Type" json:"type,omitempty"`
}

func (x *Measure) Reset() {
	*x = Measure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opencensus_proto_stats_v1_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Measure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Measure) ProtoMessage() {}

func (x *Measure) ProtoReflect() protoreflect.Message {
	mi := &file_opencensus_proto_stats_v1_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Measure.ProtoReflect.Descriptor instead.
func (*Measure) Descriptor() ([]byte, []int) {
	return file_opencensus_proto_stats_v1_stats_proto_rawDescGZIP(), []int{1}
}

func (x *Measure) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Measure) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Measure) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Measure) GetType() Measure_Type {
	if x != nil {
		return x.Type
	}
	return Measure_TYPE_UNSPECIFIED
}

type View struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A string by which the View will be referred to, e.g. "rpc_latency". Names MUST be unique
	// within the library.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Describes the view, e.g. "RPC latency distribution"
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The Measure to which this view is applied.
	Measure *Measure `protobuf:"bytes,3,opt,name=measure,proto3" json:"measure,omitempty"`
	// An array of tag keys. These values associated with tags of this name form the basis by which
	// individual stats will be aggregated (one aggregation per unique tag value). If none are
	// provided, then all data is recorded in a single aggregation.
	Columns []string `protobuf:"bytes,4,rep,name=columns,proto3" json:"columns,omitempty"`
	// The description of the aggregation used for this view which describes how data collected are
	// aggregated.
	//
	// Types that are assignable to Aggregation:
	//	*View_CountAggregation
	//	*View_SumAggregation
	//	*View_LastValueAggregation
	//	*View_DistributionAggregation
	Aggregation isView_Aggregation `protobuf_oneof:"aggregation"`
}

func (x *View) Reset() {
	*x = View{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opencensus_proto_stats_v1_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *View) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*View) ProtoMessage() {}

func (x *View) ProtoReflect() protoreflect.Message {
	mi := &file_opencensus_proto_stats_v1_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use View.ProtoReflect.Descriptor instead.
func (*View) Descriptor() ([]byte, []int) {
	return file_opencensus_proto_stats_v1_stats_proto_rawDescGZIP(), []int{2}
}

func (x *View) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *View) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *View) GetMeasure() *Measure {
	if x != nil {
		return x.Measure
	}
	return nil
}

func (x *View) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (m *View) GetAggregation() isView_Aggregation {
	if m != nil {
		return m.Aggregation
	}
	return nil
}

func (x *View) GetCountAggregation() *CountAggregation {
	if x, ok := x.GetAggregation().(*View_CountAggregation); ok {
		return x.CountAggregation
	}
	return nil
}

func (x *View) GetSumAggregation() *SumAggregation {
	if x, ok := x.GetAggregation().(*View_SumAggregation); ok {
		return x.SumAggregation
	}
	return nil
}

func (x *View) GetLastValueAggregation() *LastValueAggregation {
	if x, ok := x.GetAggregation().(*View_LastValueAggregation); ok {
		return x.LastValueAggregation
	}
	return nil
}

func (x *View) GetDistributionAggregation() *DistributionAggregation {
	if x, ok := x.GetAggregation().(*View_DistributionAggregation); ok {
		return x.DistributionAggregation
	}
	return nil
}

type isView_Aggregation interface {
	isView_Aggregation()
}

type View_CountAggregation struct {
	// Counts the number of measurements recorded.
	CountAggregation *CountAggregation `protobuf:"bytes,5,opt,name=count_aggregation,json=countAggregation,proto3,oneof"`
}

type View_SumAggregation struct {
	// Indicates that data collected and aggregated with this Aggregation will be summed up.
	SumAggregation *SumAggregation `protobuf:"bytes,6,opt,name=sum_aggregation,json=sumAggregation,proto3,oneof"`
}

type View_LastValueAggregation struct {
	// Indicates that data collected and aggregated with this Aggregation will represent the last
	// recorded value. This is useful to support Gauges.
	LastValueAggregation *LastValueAggregation `protobuf:"bytes,7,opt,name=last_value_aggregation,json=lastValueAggregation,proto3,oneof"`
}

type View_DistributionAggregation struct {
	// Indicates that the desired Aggregation is a histogram distribution. A distribution
	// Aggregation may contain a histogram of the values in the population. User should define the
	// bucket boundaries for that histogram (see DistributionAggregation).
	DistributionAggregation *DistributionAggregation `protobuf:"bytes,8,opt,name=distribution_aggregation,json=distributionAggregation,proto3,oneof"`
}

func (*View_CountAggregation) isView_Aggregation() {}

func (*View_SumAggregation) isView_Aggregation() {}

func (*View_LastValueAggregation) isView_Aggregation() {}

func (*View_DistributionAggregation) isView_Aggregation() {}

type CountAggregation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CountAggregation) Reset() {
	*x = CountAggregation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opencensus_proto_stats_v1_stats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountAggregation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountAggregation) ProtoMessage() {}

func (x *CountAggregation) ProtoReflect() protoreflect.Message {
	mi := &file_opencensus_proto_stats_v1_stats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountAggregation.ProtoReflect.Descriptor instead.
func (*CountAggregation) Descriptor() ([]byte, []int) {
	return file_opencensus_proto_stats_v1_stats_proto_rawDescGZIP(), []int{3}
}

type SumAggregation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SumAggregation) Reset() {
	*x = SumAggregation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opencensus_proto_stats_v1_stats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumAggregation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumAggregation) ProtoMessage() {}

func (x *SumAggregation) ProtoReflect() protoreflect.Message {
	mi := &file_opencensus_proto_stats_v1_stats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumAggregation.ProtoReflect.Descriptor instead.
func (*SumAggregation) Descriptor() ([]byte, []int) {
	return file_opencensus_proto_stats_v1_stats_proto_rawDescGZIP(), []int{4}
}

type LastValueAggregation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LastValueAggregation) Reset() {
	*x = LastValueAggregation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opencensus_proto_stats_v1_stats_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LastValueAggregation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LastValueAggregation) ProtoMessage() {}

func (x *LastValueAggregation) ProtoReflect() protoreflect.Message {
	mi := &file_opencensus_proto_stats_v1_stats_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LastValueAggregation.ProtoReflect.Descriptor instead.
func (*LastValueAggregation) Descriptor() ([]byte, []int) {
	return file_opencensus_proto_stats_v1_stats_proto_rawDescGZIP(), []int{5}
}

type DistributionAggregation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A Distribution may optionally contain a histogram of the values in the
	// population. The bucket boundaries for that histogram are described by
	// `bucket_bounds`. This defines `size(bucket_bounds) + 1` (= N)
	// buckets. The boundaries for bucket index i are:
	//
	// (-infinity, bucket_bounds[i]) for i == 0
	// [bucket_bounds[i-1], bucket_bounds[i]) for 0 < i < N-2
	// [bucket_bounds[i-1], +infinity) for i == N-1
	//
	// i.e. an underflow bucket (number 0), zero or more finite buckets (1
	// through N - 2, and an overflow bucket (N - 1), with inclusive lower
	// bounds and exclusive upper bounds.
	//
	// If `bucket_bounds` has no elements (zero size), then there is no
	// histogram associated with the Distribution. If `bucket_bounds` has only
	// one element, there are no finite buckets, and that single element is the
	// common boundary of the overflow and underflow buckets. The values must
	// be monotonically increasing.
	BucketBounds []float64 `protobuf:"fixed64,1,rep,packed,name=bucket_bounds,json=bucketBounds,proto3" json:"bucket_bounds,omitempty"`
}

func (x *DistributionAggregation) Reset() {
	*x = DistributionAggregation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opencensus_proto_stats_v1_stats_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributionAggregation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributionAggregation) ProtoMessage() {}

func (x *DistributionAggregation) ProtoReflect() protoreflect.Message {
	mi := &file_opencensus_proto_stats_v1_stats_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributionAggregation.ProtoReflect.Descriptor instead.
func (*DistributionAggregation) Descriptor() ([]byte, []int) {
	return file_opencensus_proto_stats_v1_stats_proto_rawDescGZIP(), []int{6}
}

func (x *DistributionAggregation) GetBucketBounds() []float64 {
	if x != nil {
		return x.BucketBounds
	}
	return nil
}

// Describes a data point to be collected for a Measure.
type Measurement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []*Tag `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	// The name of the measure to which the value is applied.
	MeasureName string `protobuf:"bytes,2,opt,name=measure_name,json=measureName,proto3" json:"measure_name,omitempty"`
	// The recorded value, MUST have the appropriate type to match the Measure.
	//
	// Types that are assignable to Value:
	//	*Measurement_DoubleValue
	//	*Measurement_IntValue
	Value isMeasurement_Value `protobuf_oneof:"value"`
	// The time when this measurement was recorded. If the implementation uses a async buffer to
	// record measurements this may be the time when the measurement was read from the buffer.
	Time *timestamp.Timestamp `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Measurement) Reset() {
	*x = Measurement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opencensus_proto_stats_v1_stats_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Measurement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Measurement) ProtoMessage() {}

func (x *Measurement) ProtoReflect() protoreflect.Message {
	mi := &file_opencensus_proto_stats_v1_stats_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Measurement.ProtoReflect.Descriptor instead.
func (*Measurement) Descriptor() ([]byte, []int) {
	return file_opencensus_proto_stats_v1_stats_proto_rawDescGZIP(), []int{7}
}

func (x *Measurement) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Measurement) GetMeasureName() string {
	if x != nil {
		return x.MeasureName
	}
	return ""
}

func (m *Measurement) GetValue() isMeasurement_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Measurement) GetDoubleValue() float64 {
	if x, ok := x.GetValue().(*Measurement_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (x *Measurement) GetIntValue() int64 {
	if x, ok := x.GetValue().(*Measurement_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *Measurement) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type isMeasurement_Value interface {
	isMeasurement_Value()
}

type Measurement_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Measurement_IntValue struct {
	IntValue int64 `protobuf:"varint,4,opt,name=int_value,json=intValue,proto3,oneof"`
}

func (*Measurement_DoubleValue) isMeasurement_Value() {}

func (*Measurement_IntValue) isMeasurement_Value() {}

var File_opencensus_proto_stats_v1_stats_proto protoreflect.FileDescriptor

var file_opencensus_proto_stats_v1_stats_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0xc5, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x3b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x33, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x22, 0xaf, 0x04, 0x0a, 0x04, 0x56,
	0x69, 0x65, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x07, 0x6d, 0x65, 0x61,
	0x73, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x52, 0x07,
	0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x73, 0x12, 0x5a, 0x0a, 0x11, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x10, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x54, 0x0a,
	0x0f, 0x73, 0x75, 0x6d, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x75, 0x6d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x0e, 0x73, 0x75, 0x6d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x67, 0x0a, 0x16, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x61, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6f, 0x0a, 0x18,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x17, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x0a,
	0x0b, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x12, 0x0a, 0x10,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x10, 0x0a, 0x0e, 0x53, 0x75, 0x6d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x16, 0x0a, 0x14, 0x4c, 0x61, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x41,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3e, 0x0a, 0x17, 0x44, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52, 0x0c, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x22, 0xe1, 0x01, 0x0a, 0x0b, 0x4d,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x63,
	0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x23, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x8c,
	0x01, 0x0a, 0x1c, 0x69, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x42,
	0x0a, 0x53, 0x74, 0x61, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x2d, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x76,
	0x31, 0xea, 0x02, 0x19, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_opencensus_proto_stats_v1_stats_proto_rawDescOnce sync.Once
	file_opencensus_proto_stats_v1_stats_proto_rawDescData = file_opencensus_proto_stats_v1_stats_proto_rawDesc
)

func file_opencensus_proto_stats_v1_stats_proto_rawDescGZIP() []byte {
	file_opencensus_proto_stats_v1_stats_proto_rawDescOnce.Do(func() {
		file_opencensus_proto_stats_v1_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_opencensus_proto_stats_v1_stats_proto_rawDescData)
	})
	return file_opencensus_proto_stats_v1_stats_proto_rawDescData
}

var file_opencensus_proto_stats_v1_stats_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_opencensus_proto_stats_v1_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_opencensus_proto_stats_v1_stats_proto_goTypes = []interface{}{
	(Measure_Type)(0),               // 0: opencensus.proto.stats.v1.Measure.Type
	(*Tag)(nil),                     // 1: opencensus.proto.stats.v1.Tag
	(*Measure)(nil),                 // 2: opencensus.proto.stats.v1.Measure
	(*View)(nil),                    // 3: opencensus.proto.stats.v1.View
	(*CountAggregation)(nil),        // 4: opencensus.proto.stats.v1.CountAggregation
	(*SumAggregation)(nil),          // 5: opencensus.proto.stats.v1.SumAggregation
	(*LastValueAggregation)(nil),    // 6: opencensus.proto.stats.v1.LastValueAggregation
	(*DistributionAggregation)(nil), // 7: opencensus.proto.stats.v1.DistributionAggregation
	(*Measurement)(nil),             // 8: opencensus.proto.stats.v1.Measurement
	(*timestamp.Timestamp)(nil),     // 9: google.protobuf.Timestamp
}
var file_opencensus_proto_stats_v1_stats_proto_depIdxs = []int32{
	0, // 0: opencensus.proto.stats.v1.Measure.type:type_name -> opencensus.proto.stats.v1.Measure.Type
	2, // 1: opencensus.proto.stats.v1.View.measure:type_name -> opencensus.proto.stats.v1.Measure
	4, // 2: opencensus.proto.stats.v1.View.count_aggregation:type_name -> opencensus.proto.stats.v1.CountAggregation
	5, // 3: opencensus.proto.stats.v1.View.sum_aggregation:type_name -> opencensus.proto.stats.v1.SumAggregation
	6, // 4: opencensus.proto.stats.v1.View.last_value_aggregation:type_name -> opencensus.proto.stats.v1.LastValueAggregation
	7, // 5: opencensus.proto.stats.v1.View.distribution_aggregation:type_name -> opencensus.proto.stats.v1.DistributionAggregation
	1, // 6: opencensus.proto.stats.v1.Measurement.tags:type_name -> opencensus.proto.stats.v1.Tag
	9, // 7: opencensus.proto.stats.v1.Measurement.time:type_name -> google.protobuf.Timestamp
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_opencensus_proto_stats_v1_stats_proto_init() }
func file_opencensus_proto_stats_v1_stats_proto_init() {
	if File_opencensus_proto_stats_v1_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_opencensus_proto_stats_v1_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_opencensus_proto_stats_v1_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Measure); i {
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
		file_opencensus_proto_stats_v1_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*View); i {
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
		file_opencensus_proto_stats_v1_stats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountAggregation); i {
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
		file_opencensus_proto_stats_v1_stats_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumAggregation); i {
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
		file_opencensus_proto_stats_v1_stats_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LastValueAggregation); i {
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
		file_opencensus_proto_stats_v1_stats_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributionAggregation); i {
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
		file_opencensus_proto_stats_v1_stats_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Measurement); i {
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
	file_opencensus_proto_stats_v1_stats_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*View_CountAggregation)(nil),
		(*View_SumAggregation)(nil),
		(*View_LastValueAggregation)(nil),
		(*View_DistributionAggregation)(nil),
	}
	file_opencensus_proto_stats_v1_stats_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*Measurement_DoubleValue)(nil),
		(*Measurement_IntValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_opencensus_proto_stats_v1_stats_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_opencensus_proto_stats_v1_stats_proto_goTypes,
		DependencyIndexes: file_opencensus_proto_stats_v1_stats_proto_depIdxs,
		EnumInfos:         file_opencensus_proto_stats_v1_stats_proto_enumTypes,
		MessageInfos:      file_opencensus_proto_stats_v1_stats_proto_msgTypes,
	}.Build()
	File_opencensus_proto_stats_v1_stats_proto = out.File
	file_opencensus_proto_stats_v1_stats_proto_rawDesc = nil
	file_opencensus_proto_stats_v1_stats_proto_goTypes = nil
	file_opencensus_proto_stats_v1_stats_proto_depIdxs = nil
}
