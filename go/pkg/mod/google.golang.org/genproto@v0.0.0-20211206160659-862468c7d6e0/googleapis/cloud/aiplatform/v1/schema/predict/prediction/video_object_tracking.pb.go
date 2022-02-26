// Copyright 2021 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/aiplatform/v1/schema/predict/prediction/video_object_tracking.proto

package prediction

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Prediction output format for Video Object Tracking.
type VideoObjectTrackingPredictionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource ID of the AnnotationSpec that had been identified.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The display name of the AnnotationSpec that had been identified.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The beginning, inclusive, of the video's time segment in which the
	// object instance has been detected. Expressed as a number of seconds as
	// measured from the start of the video, with fractions up to a microsecond
	// precision, and with "s" appended at the end.
	TimeSegmentStart *durationpb.Duration `protobuf:"bytes,3,opt,name=time_segment_start,json=timeSegmentStart,proto3" json:"time_segment_start,omitempty"`
	// The end, inclusive, of the video's time segment in which the
	// object instance has been detected. Expressed as a number of seconds as
	// measured from the start of the video, with fractions up to a microsecond
	// precision, and with "s" appended at the end.
	TimeSegmentEnd *durationpb.Duration `protobuf:"bytes,4,opt,name=time_segment_end,json=timeSegmentEnd,proto3" json:"time_segment_end,omitempty"`
	// The Model's confidence in correction of this prediction, higher
	// value means higher confidence.
	Confidence *wrapperspb.FloatValue `protobuf:"bytes,5,opt,name=confidence,proto3" json:"confidence,omitempty"`
	// All of the frames of the video in which a single object instance has been
	// detected. The bounding boxes in the frames identify the same object.
	Frames []*VideoObjectTrackingPredictionResult_Frame `protobuf:"bytes,6,rep,name=frames,proto3" json:"frames,omitempty"`
}

func (x *VideoObjectTrackingPredictionResult) Reset() {
	*x = VideoObjectTrackingPredictionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoObjectTrackingPredictionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoObjectTrackingPredictionResult) ProtoMessage() {}

func (x *VideoObjectTrackingPredictionResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoObjectTrackingPredictionResult.ProtoReflect.Descriptor instead.
func (*VideoObjectTrackingPredictionResult) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_rawDescGZIP(), []int{0}
}

func (x *VideoObjectTrackingPredictionResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VideoObjectTrackingPredictionResult) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *VideoObjectTrackingPredictionResult) GetTimeSegmentStart() *durationpb.Duration {
	if x != nil {
		return x.TimeSegmentStart
	}
	return nil
}

func (x *VideoObjectTrackingPredictionResult) GetTimeSegmentEnd() *durationpb.Duration {
	if x != nil {
		return x.TimeSegmentEnd
	}
	return nil
}

func (x *VideoObjectTrackingPredictionResult) GetConfidence() *wrapperspb.FloatValue {
	if x != nil {
		return x.Confidence
	}
	return nil
}

func (x *VideoObjectTrackingPredictionResult) GetFrames() []*VideoObjectTrackingPredictionResult_Frame {
	if x != nil {
		return x.Frames
	}
	return nil
}

// The fields `xMin`, `xMax`, `yMin`, and `yMax` refer to a bounding box,
// i.e. the rectangle over the video frame pinpointing the found
// AnnotationSpec. The coordinates are relative to the frame size, and the
// point 0,0 is in the top left of the frame.
type VideoObjectTrackingPredictionResult_Frame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A time (frame) of a video in which the object has been detected.
	// Expressed as a number of seconds as measured from the
	// start of the video, with fractions up to a microsecond precision, and
	// with "s" appended at the end.
	TimeOffset *durationpb.Duration `protobuf:"bytes,1,opt,name=time_offset,json=timeOffset,proto3" json:"time_offset,omitempty"`
	// The leftmost coordinate of the bounding box.
	XMin *wrapperspb.FloatValue `protobuf:"bytes,2,opt,name=x_min,json=xMin,proto3" json:"x_min,omitempty"`
	// The rightmost coordinate of the bounding box.
	XMax *wrapperspb.FloatValue `protobuf:"bytes,3,opt,name=x_max,json=xMax,proto3" json:"x_max,omitempty"`
	// The topmost coordinate of the bounding box.
	YMin *wrapperspb.FloatValue `protobuf:"bytes,4,opt,name=y_min,json=yMin,proto3" json:"y_min,omitempty"`
	// The bottommost coordinate of the bounding box.
	YMax *wrapperspb.FloatValue `protobuf:"bytes,5,opt,name=y_max,json=yMax,proto3" json:"y_max,omitempty"`
}

func (x *VideoObjectTrackingPredictionResult_Frame) Reset() {
	*x = VideoObjectTrackingPredictionResult_Frame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoObjectTrackingPredictionResult_Frame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoObjectTrackingPredictionResult_Frame) ProtoMessage() {}

func (x *VideoObjectTrackingPredictionResult_Frame) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoObjectTrackingPredictionResult_Frame.ProtoReflect.Descriptor instead.
func (*VideoObjectTrackingPredictionResult_Frame) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_rawDescGZIP(), []int{0, 0}
}

func (x *VideoObjectTrackingPredictionResult_Frame) GetTimeOffset() *durationpb.Duration {
	if x != nil {
		return x.TimeOffset
	}
	return nil
}

func (x *VideoObjectTrackingPredictionResult_Frame) GetXMin() *wrapperspb.FloatValue {
	if x != nil {
		return x.XMin
	}
	return nil
}

func (x *VideoObjectTrackingPredictionResult_Frame) GetXMax() *wrapperspb.FloatValue {
	if x != nil {
		return x.XMax
	}
	return nil
}

func (x *VideoObjectTrackingPredictionResult_Frame) GetYMin() *wrapperspb.FloatValue {
	if x != nil {
		return x.YMin
	}
	return nil
}

func (x *VideoObjectTrackingPredictionResult_Frame) GetYMax() *wrapperspb.FloatValue {
	if x != nil {
		return x.YMax
	}
	return nil
}

var File_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_rawDesc = []byte{
	0x0a, 0x50, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x05, 0x0a, 0x23, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x47, 0x0a, 0x12, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x43, 0x0a, 0x10, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0e, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x64, 0x12,
	0x3b, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x77, 0x0a, 0x06,
	0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x5f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x06, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x73, 0x1a, 0x8b, 0x02, 0x0a, 0x05, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12,
	0x3a, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0a, 0x74, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x78,
	0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f,
	0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x78, 0x4d, 0x69, 0x6e, 0x12, 0x30, 0x0a,
	0x05, 0x78, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x78, 0x4d, 0x61, 0x78, 0x12,
	0x30, 0x0a, 0x05, 0x79, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x79, 0x4d, 0x69,
	0x6e, 0x12, 0x30, 0x0a, 0x05, 0x79, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x79,
	0x4d, 0x61, 0x78, 0x42, 0xf1, 0x02, 0x0a, 0x38, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x28, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x3b, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0xaa, 0x02, 0x34, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0xca, 0x02, 0x34, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31,
	0x5c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5c, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x5c,
	0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0xea, 0x02, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x3a, 0x3a, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x3a, 0x3a, 0x50, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_rawDescData = file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_goTypes = []interface{}{
	(*VideoObjectTrackingPredictionResult)(nil),       // 0: google.cloud.aiplatform.v1.schema.predict.prediction.VideoObjectTrackingPredictionResult
	(*VideoObjectTrackingPredictionResult_Frame)(nil), // 1: google.cloud.aiplatform.v1.schema.predict.prediction.VideoObjectTrackingPredictionResult.Frame
	(*durationpb.Duration)(nil),                       // 2: google.protobuf.Duration
	(*wrapperspb.FloatValue)(nil),                     // 3: google.protobuf.FloatValue
}
var file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_depIdxs = []int32{
	2, // 0: google.cloud.aiplatform.v1.schema.predict.prediction.VideoObjectTrackingPredictionResult.time_segment_start:type_name -> google.protobuf.Duration
	2, // 1: google.cloud.aiplatform.v1.schema.predict.prediction.VideoObjectTrackingPredictionResult.time_segment_end:type_name -> google.protobuf.Duration
	3, // 2: google.cloud.aiplatform.v1.schema.predict.prediction.VideoObjectTrackingPredictionResult.confidence:type_name -> google.protobuf.FloatValue
	1, // 3: google.cloud.aiplatform.v1.schema.predict.prediction.VideoObjectTrackingPredictionResult.frames:type_name -> google.cloud.aiplatform.v1.schema.predict.prediction.VideoObjectTrackingPredictionResult.Frame
	2, // 4: google.cloud.aiplatform.v1.schema.predict.prediction.VideoObjectTrackingPredictionResult.Frame.time_offset:type_name -> google.protobuf.Duration
	3, // 5: google.cloud.aiplatform.v1.schema.predict.prediction.VideoObjectTrackingPredictionResult.Frame.x_min:type_name -> google.protobuf.FloatValue
	3, // 6: google.cloud.aiplatform.v1.schema.predict.prediction.VideoObjectTrackingPredictionResult.Frame.x_max:type_name -> google.protobuf.FloatValue
	3, // 7: google.cloud.aiplatform.v1.schema.predict.prediction.VideoObjectTrackingPredictionResult.Frame.y_min:type_name -> google.protobuf.FloatValue
	3, // 8: google.cloud.aiplatform.v1.schema.predict.prediction.VideoObjectTrackingPredictionResult.Frame.y_max:type_name -> google.protobuf.FloatValue
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() {
	file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_init()
}
func file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_init() {
	if File_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoObjectTrackingPredictionResult); i {
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
		file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoObjectTrackingPredictionResult_Frame); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto = out.File
	file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_schema_predict_prediction_video_object_tracking_proto_depIdxs = nil
}
