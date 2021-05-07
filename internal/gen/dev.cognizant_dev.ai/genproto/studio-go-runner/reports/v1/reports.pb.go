// Copyright © 2020 Cognizant Digital Business, Evolutionary AI. All rights reserved. Issued under the Apache 2.0 license.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: reports.proto

package reports

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LogSeverity int32

const (
	// (0) The log entry has no assigned severity level.
	LogSeverity_Default LogSeverity = 0
	// (100) Trace information.
	LogSeverity_Trace LogSeverity = 100
	// (200) Debug information.
	LogSeverity_Debug LogSeverity = 200
	// (300) Routine information, such as ongoing status or performance.
	LogSeverity_Info LogSeverity = 300
	// (400) Normal but significant events, such as start up, shut down, or
	// a configuration change.
	LogSeverity_Warning LogSeverity = 400
	// (500) Error events are likely to cause problems.
	LogSeverity_Error LogSeverity = 500
	// (600) One or more systems are unusable.
	LogSeverity_Fatal LogSeverity = 600
)

// Enum value maps for LogSeverity.
var (
	LogSeverity_name = map[int32]string{
		0:   "Default",
		100: "Trace",
		200: "Debug",
		300: "Info",
		400: "Warning",
		500: "Error",
		600: "Fatal",
	}
	LogSeverity_value = map[string]int32{
		"Default": 0,
		"Trace":   100,
		"Debug":   200,
		"Info":    300,
		"Warning": 400,
		"Error":   500,
		"Fatal":   600,
	}
)

func (x LogSeverity) Enum() *LogSeverity {
	p := new(LogSeverity)
	*p = x
	return p
}

func (x LogSeverity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogSeverity) Descriptor() protoreflect.EnumDescriptor {
	return file_reports_proto_enumTypes[0].Descriptor()
}

func (LogSeverity) Type() protoreflect.EnumType {
	return &file_reports_proto_enumTypes[0]
}

func (x LogSeverity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogSeverity.Descriptor instead.
func (LogSeverity) EnumDescriptor() ([]byte, []int) {
	return file_reports_proto_rawDescGZIP(), []int{0}
}

type TaskState int32

const (
	// (0) The Task is in an intialization phase and has not started
	TaskState_Prestart TaskState = 0
	// (1) The task is in a starting state, optional transitional state
	TaskState_Started TaskState = 1
	// (2) The task is stopping, optional transitional state
	TaskState_Stopping TaskState = 2
	// (20) Terminal state indicating the task failed
	TaskState_Failed TaskState = 20
	// (21) Terminal state indicating the task completed successfully
	TaskState_Success TaskState = 21
)

// Enum value maps for TaskState.
var (
	TaskState_name = map[int32]string{
		0:  "Prestart",
		1:  "Started",
		2:  "Stopping",
		20: "Failed",
		21: "Success",
	}
	TaskState_value = map[string]int32{
		"Prestart": 0,
		"Started":  1,
		"Stopping": 2,
		"Failed":   20,
		"Success":  21,
	}
)

func (x TaskState) Enum() *TaskState {
	p := new(TaskState)
	*p = x
	return p
}

func (x TaskState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskState) Descriptor() protoreflect.EnumDescriptor {
	return file_reports_proto_enumTypes[1].Descriptor()
}

func (TaskState) Type() protoreflect.EnumType {
	return &file_reports_proto_enumTypes[1]
}

func (x TaskState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskState.Descriptor instead.
func (TaskState) EnumDescriptor() ([]byte, []int) {
	return file_reports_proto_rawDescGZIP(), []int{1}
}

// Queue message format. All messages conform to this format
type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Timestamp of the time when the report message was emitted
	Time *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	// The unique ID of this experiment assigned by the experimenter.
	// If the report has no associated or known experiment id this
	// field will not be present.
	ExperimentId *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=experiment_id,json=experimentId,proto3" json:"experiment_id,omitempty"`
	// A unique ID that was generated by the runners attempt to run the experiment.
	// This value will change between attempts. If the report has no associated or
	// known experiment id this field will not be present.
	UniqueId *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
	// A unique ID denoting the host, pod, or node that this experiment attempt
	// is being performed on.
	ExecutorId *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=executor_id,json=executorId,proto3" json:"executor_id,omitempty"`
	// Types that are assignable to Payload:
	//	*Report_ProtoAny
	//	*Report_Text
	//	*Report_Logging
	//	*Report_Progress
	Payload isReport_Payload `protobuf_oneof:"payload"`
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reports_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_reports_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_reports_proto_rawDescGZIP(), []int{0}
}

func (x *Report) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Report) GetExperimentId() *wrapperspb.StringValue {
	if x != nil {
		return x.ExperimentId
	}
	return nil
}

func (x *Report) GetUniqueId() *wrapperspb.StringValue {
	if x != nil {
		return x.UniqueId
	}
	return nil
}

func (x *Report) GetExecutorId() *wrapperspb.StringValue {
	if x != nil {
		return x.ExecutorId
	}
	return nil
}

func (m *Report) GetPayload() isReport_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Report) GetProtoAny() *anypb.Any {
	if x, ok := x.GetPayload().(*Report_ProtoAny); ok {
		return x.ProtoAny
	}
	return nil
}

func (x *Report) GetText() *wrapperspb.StringValue {
	if x, ok := x.GetPayload().(*Report_Text); ok {
		return x.Text
	}
	return nil
}

func (x *Report) GetLogging() *LogEntry {
	if x, ok := x.GetPayload().(*Report_Logging); ok {
		return x.Logging
	}
	return nil
}

func (x *Report) GetProgress() *Progress {
	if x, ok := x.GetPayload().(*Report_Progress); ok {
		return x.Progress
	}
	return nil
}

type isReport_Payload interface {
	isReport_Payload()
}

type Report_ProtoAny struct {
	ProtoAny *anypb.Any `protobuf:"bytes,10,opt,name=proto_any,json=protoAny,proto3,oneof"`
}

type Report_Text struct {
	// The log entry payload, represented as a Unicode string (UTF-8).
	Text *wrapperspb.StringValue `protobuf:"bytes,11,opt,name=text,proto3,oneof"`
}

type Report_Logging struct {
	// A structured log message from the runner
	Logging *LogEntry `protobuf:"bytes,12,opt,name=logging,proto3,oneof"`
}

type Report_Progress struct {
	// A progress message generated by the experiment
	Progress *Progress `protobuf:"bytes,13,opt,name=progress,proto3,oneof"`
}

func (*Report_ProtoAny) isReport_Payload() {}

func (*Report_Text) isReport_Payload() {}

func (*Report_Logging) isReport_Payload() {}

func (*Report_Progress) isReport_Payload() {}

// Structured log message
type LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Timestamp of the time when the log entry was generated
	Time *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	// Severity code for the log entry
	Severity LogSeverity `protobuf:"varint,2,opt,name=severity,proto3,enum=dev.studio_go_runner.reports.LogSeverity" json:"severity,omitempty"`
	// Mesage string
	Message *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// Key value pairs of context information
	Fields map[string]string `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *LogEntry) Reset() {
	*x = LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reports_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntry) ProtoMessage() {}

func (x *LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_reports_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntry.ProtoReflect.Descriptor instead.
func (*LogEntry) Descriptor() ([]byte, []int) {
	return file_reports_proto_rawDescGZIP(), []int{1}
}

func (x *LogEntry) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *LogEntry) GetSeverity() LogSeverity {
	if x != nil {
		return x.Severity
	}
	return LogSeverity_Default
}

func (x *LogEntry) GetMessage() *wrapperspb.StringValue {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *LogEntry) GetFields() map[string]string {
	if x != nil {
		return x.Fields
	}
	return nil
}

// JSON messages emitted by the experiment application code.  These messages
// should conform to the format documented in [docs/metadata.md](docs/metadata.md#JSON-Document)
type Progress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Timestamp of the time when the json payload was generated by the experiment
	Time *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	// The valid json payload emitted by the experiment, or empty if the message is generated
	// by the runner itself
	Json *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=json,proto3" json:"json,omitempty"`
	// This field is used to indicated the state of the task to which this Progress message relates
	State TaskState       `protobuf:"varint,3,opt,name=state,proto3,enum=dev.studio_go_runner.reports.TaskState" json:"state,omitempty"`
	Error *Progress_Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Progress) Reset() {
	*x = Progress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reports_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Progress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Progress) ProtoMessage() {}

func (x *Progress) ProtoReflect() protoreflect.Message {
	mi := &file_reports_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Progress.ProtoReflect.Descriptor instead.
func (*Progress) Descriptor() ([]byte, []int) {
	return file_reports_proto_rawDescGZIP(), []int{2}
}

func (x *Progress) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Progress) GetJson() *wrapperspb.StringValue {
	if x != nil {
		return x.Json
	}
	return nil
}

func (x *Progress) GetState() TaskState {
	if x != nil {
		return x.State
	}
	return TaskState_Prestart
}

func (x *Progress) GetError() *Progress_Error {
	if x != nil {
		return x.Error
	}
	return nil
}

// This field is used to send any available failure information such as an error code
type Progress_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The text message associated with an error
	Msg *wrapperspb.StringValue `protobuf:"bytes,20,opt,name=msg,proto3" json:"msg,omitempty"`
	// If an error code is available it will be included in the code value
	Code int32 `protobuf:"varint,21,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Progress_Error) Reset() {
	*x = Progress_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reports_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Progress_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Progress_Error) ProtoMessage() {}

func (x *Progress_Error) ProtoReflect() protoreflect.Message {
	mi := &file_reports_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Progress_Error.ProtoReflect.Descriptor instead.
func (*Progress_Error) Descriptor() ([]byte, []int) {
	return file_reports_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Progress_Error) GetMsg() *wrapperspb.StringValue {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *Progress_Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_reports_proto protoreflect.FileDescriptor

var file_reports_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1c, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x67, 0x6f, 0x5f, 0x72,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x03, 0x0a, 0x06, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65,
	0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x49,
	0x64, 0x12, 0x33, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x6e, 0x79, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x41, 0x6e, 0x79, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x42, 0x0a, 0x07, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x67, 0x6f, 0x5f, 0x72, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x48, 0x00, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x44,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x67, 0x6f,
	0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0xc0, 0x02, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x2e, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x08,
	0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x67, 0x6f, 0x5f, 0x72,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x4c, 0x6f,
	0x67, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x67, 0x6f, 0x5f, 0x72, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xbc, 0x02, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x30, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6a, 0x73, 0x6f,
	0x6e, 0x12, 0x3d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x27, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x67, 0x6f,
	0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x42, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x67, 0x6f, 0x5f,
	0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x1a, 0x4b, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2e, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x2a, 0x62, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x0b, 0x0a, 0x07, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x54, 0x72, 0x61, 0x63, 0x65, 0x10, 0x64, 0x12, 0x0a, 0x0a, 0x05, 0x44, 0x65, 0x62, 0x75,
	0x67, 0x10, 0xc8, 0x01, 0x12, 0x09, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0xac, 0x02, 0x12,
	0x0c, 0x0a, 0x07, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x90, 0x03, 0x12, 0x0a, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0xf4, 0x03, 0x12, 0x0a, 0x0a, 0x05, 0x46, 0x61, 0x74,
	0x61, 0x6c, 0x10, 0xd8, 0x04, 0x2a, 0x4d, 0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x53, 0x74, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x14, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x10, 0x15, 0x42, 0x43, 0x5a, 0x41, 0x64, 0x65, 0x76, 0x2e, 0x63, 0x6f, 0x67, 0x6e,
	0x69, 0x7a, 0x61, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x76, 0x2e, 0x61, 0x69, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2d, 0x67, 0x6f, 0x2d,
	0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_reports_proto_rawDescOnce sync.Once
	file_reports_proto_rawDescData = file_reports_proto_rawDesc
)

func file_reports_proto_rawDescGZIP() []byte {
	file_reports_proto_rawDescOnce.Do(func() {
		file_reports_proto_rawDescData = protoimpl.X.CompressGZIP(file_reports_proto_rawDescData)
	})
	return file_reports_proto_rawDescData
}

var file_reports_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_reports_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_reports_proto_goTypes = []interface{}{
	(LogSeverity)(0),               // 0: dev.studio_go_runner.reports.LogSeverity
	(TaskState)(0),                 // 1: dev.studio_go_runner.reports.TaskState
	(*Report)(nil),                 // 2: dev.studio_go_runner.reports.Report
	(*LogEntry)(nil),               // 3: dev.studio_go_runner.reports.LogEntry
	(*Progress)(nil),               // 4: dev.studio_go_runner.reports.Progress
	nil,                            // 5: dev.studio_go_runner.reports.LogEntry.FieldsEntry
	(*Progress_Error)(nil),         // 6: dev.studio_go_runner.reports.Progress.Error
	(*timestamppb.Timestamp)(nil),  // 7: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil), // 8: google.protobuf.StringValue
	(*anypb.Any)(nil),              // 9: google.protobuf.Any
}
var file_reports_proto_depIdxs = []int32{
	7,  // 0: dev.studio_go_runner.reports.Report.time:type_name -> google.protobuf.Timestamp
	8,  // 1: dev.studio_go_runner.reports.Report.experiment_id:type_name -> google.protobuf.StringValue
	8,  // 2: dev.studio_go_runner.reports.Report.unique_id:type_name -> google.protobuf.StringValue
	8,  // 3: dev.studio_go_runner.reports.Report.executor_id:type_name -> google.protobuf.StringValue
	9,  // 4: dev.studio_go_runner.reports.Report.proto_any:type_name -> google.protobuf.Any
	8,  // 5: dev.studio_go_runner.reports.Report.text:type_name -> google.protobuf.StringValue
	3,  // 6: dev.studio_go_runner.reports.Report.logging:type_name -> dev.studio_go_runner.reports.LogEntry
	4,  // 7: dev.studio_go_runner.reports.Report.progress:type_name -> dev.studio_go_runner.reports.Progress
	7,  // 8: dev.studio_go_runner.reports.LogEntry.time:type_name -> google.protobuf.Timestamp
	0,  // 9: dev.studio_go_runner.reports.LogEntry.severity:type_name -> dev.studio_go_runner.reports.LogSeverity
	8,  // 10: dev.studio_go_runner.reports.LogEntry.message:type_name -> google.protobuf.StringValue
	5,  // 11: dev.studio_go_runner.reports.LogEntry.fields:type_name -> dev.studio_go_runner.reports.LogEntry.FieldsEntry
	7,  // 12: dev.studio_go_runner.reports.Progress.time:type_name -> google.protobuf.Timestamp
	8,  // 13: dev.studio_go_runner.reports.Progress.json:type_name -> google.protobuf.StringValue
	1,  // 14: dev.studio_go_runner.reports.Progress.state:type_name -> dev.studio_go_runner.reports.TaskState
	6,  // 15: dev.studio_go_runner.reports.Progress.error:type_name -> dev.studio_go_runner.reports.Progress.Error
	8,  // 16: dev.studio_go_runner.reports.Progress.Error.msg:type_name -> google.protobuf.StringValue
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_reports_proto_init() }
func file_reports_proto_init() {
	if File_reports_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reports_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Report); i {
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
		file_reports_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntry); i {
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
		file_reports_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Progress); i {
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
		file_reports_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Progress_Error); i {
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
	file_reports_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Report_ProtoAny)(nil),
		(*Report_Text)(nil),
		(*Report_Logging)(nil),
		(*Report_Progress)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_reports_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_reports_proto_goTypes,
		DependencyIndexes: file_reports_proto_depIdxs,
		EnumInfos:         file_reports_proto_enumTypes,
		MessageInfos:      file_reports_proto_msgTypes,
	}.Build()
	File_reports_proto = out.File
	file_reports_proto_rawDesc = nil
	file_reports_proto_goTypes = nil
	file_reports_proto_depIdxs = nil
}
