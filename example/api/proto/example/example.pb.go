// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_api_example

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/micro/go-micro/api/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x85, 0xd1, 0x7a, 0x60, 0x51, 0x21, 0x91,
	0xf4, 0x7c, 0xbd, 0xdc, 0xcc, 0xe4, 0xa2, 0x7c, 0xbd, 0xc4, 0x82, 0x4c, 0x3d, 0xa8, 0x9c, 0x94,
	0x6e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0x58, 0x56, 0x3f, 0x3d,
	0x5f, 0x17, 0xc2, 0x48, 0x2c, 0xc8, 0xd4, 0x87, 0x18, 0x08, 0xd2, 0x00, 0x66, 0x19, 0x99, 0x71,
	0xb1, 0xbb, 0x42, 0x74, 0x0a, 0x69, 0x73, 0xb1, 0x38, 0x27, 0xe6, 0xe4, 0x08, 0xf1, 0xeb, 0xa5,
	0x43, 0x8c, 0x0c, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x91, 0x12, 0x40, 0x08, 0x14, 0x17, 0xe4,
	0xe7, 0x15, 0xa7, 0x2a, 0x31, 0x24, 0xb1, 0x81, 0xb5, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xf1, 0x91, 0x00, 0x79, 0xa2, 0x00, 0x00, 0x00,
}