// Code generated by protoc-gen-go.
// source: server/pkg/metrics/metrics.proto
// DO NOT EDIT!

/*
Package metrics is a generated protocol buffer package.

It is generated from these files:
	server/pkg/metrics/metrics.proto

It has these top-level messages:
	Metrics
*/
package metrics

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Metrics struct {
	ID        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	PodID     string `protobuf:"bytes,2,opt,name=pod_id,json=podId" json:"pod_id,omitempty"`
	Nodes     int64  `protobuf:"varint,3,opt,name=nodes" json:"nodes,omitempty"`
	Version   string `protobuf:"bytes,4,opt,name=version" json:"version,omitempty"`
	Repos     int64  `protobuf:"varint,5,opt,name=repos" json:"repos,omitempty"`
	Commits   int64  `protobuf:"varint,6,opt,name=commits" json:"commits,omitempty"`
	Files     int64  `protobuf:"varint,7,opt,name=files" json:"files,omitempty"`
	Bytes     int64  `protobuf:"varint,8,opt,name=bytes" json:"bytes,omitempty"`
	Jobs      int64  `protobuf:"varint,9,opt,name=jobs" json:"jobs,omitempty"`
	Pipelines int64  `protobuf:"varint,10,opt,name=pipelines" json:"pipelines,omitempty"`
}

func (m *Metrics) Reset()                    { *m = Metrics{} }
func (m *Metrics) String() string            { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()               {}
func (*Metrics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Metrics)(nil), "Metrics")
}

var fileDescriptor0 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0x8f, 0x3b, 0x8a, 0xc3, 0x30,
	0x10, 0x40, 0xf1, 0x7f, 0x3d, 0xc5, 0x16, 0x62, 0x17, 0x54, 0x6c, 0x61, 0xb6, 0x4a, 0x15, 0x17,
	0x39, 0x45, 0x8a, 0x34, 0xbe, 0x40, 0xc0, 0x96, 0x12, 0x94, 0xd8, 0x1e, 0x21, 0x89, 0x40, 0x2e,
	0x9d, 0x33, 0x64, 0x34, 0xc2, 0xa4, 0xd2, 0xbc, 0xa7, 0x87, 0xc4, 0x40, 0xe7, 0xb5, 0x7b, 0x68,
	0xd7, 0xdb, 0xfb, 0xb5, 0x5f, 0x74, 0x70, 0x66, 0xf2, 0xdb, 0xb9, 0xb7, 0x0e, 0x03, 0xfe, 0xbf,
	0x32, 0x68, 0x4e, 0xc9, 0x88, 0x6f, 0xc8, 0x8d, 0x92, 0x59, 0x97, 0xed, 0xda, 0x81, 0x26, 0xf1,
	0x0b, 0xb5, 0x45, 0x75, 0x26, 0x97, 0xb3, 0xab, 0x88, 0x8e, 0x4a, 0xfc, 0x40, 0xb5, 0xa2, 0xd2,
	0x5e, 0x16, 0x64, 0x8b, 0x21, 0x81, 0x90, 0xd0, 0xd0, 0x4f, 0xde, 0xe0, 0x2a, 0x4b, 0xae, 0x37,
	0x8c, 0xbd, 0xd3, 0x16, 0xbd, 0xac, 0x52, 0xcf, 0x10, 0xfb, 0x09, 0x97, 0xc5, 0x04, 0x2f, 0x6b,
	0xf6, 0x1b, 0xc6, 0xfe, 0x62, 0x66, 0x7a, 0xbf, 0x49, 0x3d, 0x43, 0xb4, 0xe3, 0x33, 0x90, 0xfd,
	0x4a, 0x96, 0x41, 0x08, 0x28, 0x6f, 0x38, 0x7a, 0xd9, 0xb2, 0xe4, 0x59, 0xfc, 0x41, 0x6b, 0x8d,
	0xd5, 0xb3, 0x59, 0xa9, 0x06, 0xbe, 0xf8, 0x88, 0xb1, 0xe6, 0xbd, 0x0f, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x98, 0x2c, 0xab, 0xbc, 0x1b, 0x01, 0x00, 0x00,
}
