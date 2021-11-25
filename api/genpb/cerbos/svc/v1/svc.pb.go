// Copyright 2021 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: cerbos/svc/v1/svc.proto

package svcv1

import (
	v1 "github.com/cerbos/cerbos/api/genpb/cerbos/request/v1"
	v11 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_cerbos_svc_v1_svc_proto protoreflect.FileDescriptor

var file_cerbos_svc_v1_svc_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x73, 0x76, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x65, 0x72, 0x62, 0x6f,
	0x73, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73,
	0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x63, 0x65, 0x72, 0x62, 0x6f,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc6, 0x07, 0x0a, 0x0d, 0x43,
	0x65, 0x72, 0x62, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xf6, 0x01, 0x0a,
	0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65,
	0x74, 0x12, 0x2a, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x87, 0x01, 0x92, 0x41,
	0x6f, 0x12, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x1a, 0x66, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x20,
	0x77, 0x68, 0x65, 0x74, 0x68, 0x65, 0x72, 0x20, 0x61, 0x20, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69,
	0x70, 0x61, 0x6c, 0x20, 0x68, 0x61, 0x73, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x20, 0x6f, 0x6e, 0x20, 0x61, 0x20, 0x73, 0x65, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x3a, 0x01, 0x2a, 0x12, 0x86, 0x02, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x2c, 0x2e, 0x63,
	0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x65, 0x72,
	0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x91, 0x01, 0x92, 0x41, 0x6a,
	0x12, 0x14, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x20, 0x62, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x52, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x61, 0x20,
	0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x27, 0x73, 0x20, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x20, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x20, 0x6f, 0x66, 0x20, 0x68, 0x65, 0x74, 0x65, 0x72, 0x6f, 0x67, 0x65, 0x6e, 0x65,
	0x6f, 0x75, 0x73, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x20, 0x61, 0x6e,
	0x64, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e,
	0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x3a, 0x01, 0x2a, 0x12, 0xc5,
	0x01, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x2e,
	0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x69, 0x92, 0x41, 0x4e,
	0x12, 0x16, 0x47, 0x65, 0x74, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x69, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x34, 0x47, 0x65, 0x74, 0x20, 0x69, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x65, 0x2e, 0x67, 0x2e, 0x20,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0xc7, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x2c, 0x2e,
	0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x65,
	0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0x92, 0x41, 0x2c,
	0x12, 0x14, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x20, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x20, 0x70, 0x6c, 0x61, 0x6e, 0x1a, 0x14, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x20, 0x71, 0x75, 0x65, 0x72, 0x79, 0x20, 0x70, 0x6c, 0x61, 0x6e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1e, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x3a, 0x01, 0x2a,
	0x1a, 0x21, 0x92, 0x41, 0x1e, 0x12, 0x1c, 0x43, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x20, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x20, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x32, 0x8a, 0x05, 0x0a, 0x12, 0x43, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc9, 0x01, 0x0a, 0x11, 0x41,
	0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x2b, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e,
	0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58, 0x92, 0x41,
	0x29, 0x12, 0x16, 0x41, 0x64, 0x64, 0x20, 0x6f, 0x72, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x62, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x42,
	0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26,
	0x22, 0x0d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a,
	0x01, 0x2a, 0x5a, 0x12, 0x1a, 0x0d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0xc8, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2d,
	0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x45,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e,
	0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x45,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f,
	0x92, 0x41, 0x29, 0x12, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x61, 0x75, 0x64, 0x69, 0x74, 0x20,
	0x6c, 0x6f, 0x67, 0x20, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x62, 0x0f, 0x0a, 0x0d, 0x0a,
	0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x6c, 0x6f, 0x67, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x7b, 0x6b, 0x69, 0x6e, 0x64, 0x7d, 0x30,
	0x01, 0x12, 0xb8, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x12, 0x26, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x65, 0x72,
	0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x56, 0x92, 0x41, 0x20, 0x12, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x20,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x62, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x42, 0x61,
	0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x12,
	0x12, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x5a, 0x17, 0x22, 0x12, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x1a, 0x22, 0x92, 0x41,
	0x1f, 0x12, 0x1d, 0x43, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x20, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x32, 0xdb, 0x03, 0x0a, 0x17, 0x43, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x97, 0x01, 0x0a,
	0x12, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x97, 0x01, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x79, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x2e,
	0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x45, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x65,
	0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a,
	0x12, 0x8b, 0x01, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x12, 0x29, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x3a, 0x01, 0x2a, 0x42, 0xe1,
	0x01, 0x0a, 0x15, 0x64, 0x65, 0x76, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x76, 0x63, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x63, 0x65, 0x72, 0x62,
	0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x62, 0x2f, 0x63, 0x65, 0x72,
	0x62, 0x6f, 0x73, 0x2f, 0x73, 0x76, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x76, 0x63, 0x76, 0x31,
	0xaa, 0x02, 0x11, 0x43, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31,
	0x2e, 0x53, 0x76, 0x63, 0x92, 0x41, 0x7b, 0x12, 0x3f, 0x0a, 0x06, 0x43, 0x65, 0x72, 0x62, 0x6f,
	0x73, 0x22, 0x2d, 0x0a, 0x06, 0x43, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x12, 0x12, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x1a,
	0x0f, 0x69, 0x6e, 0x66, 0x6f, 0x40, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x64, 0x65, 0x76,
	0x32, 0x06, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x2a, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a,
	0x11, 0x0a, 0x0f, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x02,
	0x08, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cerbos_svc_v1_svc_proto_goTypes = []interface{}{
	(*v1.CheckResourceSetRequest)(nil),      // 0: cerbos.request.v1.CheckResourceSetRequest
	(*v1.CheckResourceBatchRequest)(nil),    // 1: cerbos.request.v1.CheckResourceBatchRequest
	(*v1.ServerInfoRequest)(nil),            // 2: cerbos.request.v1.ServerInfoRequest
	(*v1.ResourcesQueryPlanRequest)(nil),    // 3: cerbos.request.v1.ResourcesQueryPlanRequest
	(*v1.AddOrUpdatePolicyRequest)(nil),     // 4: cerbos.request.v1.AddOrUpdatePolicyRequest
	(*v1.ListAuditLogEntriesRequest)(nil),   // 5: cerbos.request.v1.ListAuditLogEntriesRequest
	(*v1.ListPoliciesRequest)(nil),          // 6: cerbos.request.v1.ListPoliciesRequest
	(*v1.PlaygroundValidateRequest)(nil),    // 7: cerbos.request.v1.PlaygroundValidateRequest
	(*v1.PlaygroundEvaluateRequest)(nil),    // 8: cerbos.request.v1.PlaygroundEvaluateRequest
	(*v1.PlaygroundProxyRequest)(nil),       // 9: cerbos.request.v1.PlaygroundProxyRequest
	(*v11.CheckResourceSetResponse)(nil),    // 10: cerbos.response.v1.CheckResourceSetResponse
	(*v11.CheckResourceBatchResponse)(nil),  // 11: cerbos.response.v1.CheckResourceBatchResponse
	(*v11.ServerInfoResponse)(nil),          // 12: cerbos.response.v1.ServerInfoResponse
	(*v11.ResourcesQueryPlanResponse)(nil),  // 13: cerbos.response.v1.ResourcesQueryPlanResponse
	(*v11.AddOrUpdatePolicyResponse)(nil),   // 14: cerbos.response.v1.AddOrUpdatePolicyResponse
	(*v11.ListAuditLogEntriesResponse)(nil), // 15: cerbos.response.v1.ListAuditLogEntriesResponse
	(*v11.ListPoliciesResponse)(nil),        // 16: cerbos.response.v1.ListPoliciesResponse
	(*v11.PlaygroundValidateResponse)(nil),  // 17: cerbos.response.v1.PlaygroundValidateResponse
	(*v11.PlaygroundEvaluateResponse)(nil),  // 18: cerbos.response.v1.PlaygroundEvaluateResponse
	(*v11.PlaygroundProxyResponse)(nil),     // 19: cerbos.response.v1.PlaygroundProxyResponse
}
var file_cerbos_svc_v1_svc_proto_depIdxs = []int32{
	0,  // 0: cerbos.svc.v1.CerbosService.CheckResourceSet:input_type -> cerbos.request.v1.CheckResourceSetRequest
	1,  // 1: cerbos.svc.v1.CerbosService.CheckResourceBatch:input_type -> cerbos.request.v1.CheckResourceBatchRequest
	2,  // 2: cerbos.svc.v1.CerbosService.ServerInfo:input_type -> cerbos.request.v1.ServerInfoRequest
	3,  // 3: cerbos.svc.v1.CerbosService.ResourcesQueryPlan:input_type -> cerbos.request.v1.ResourcesQueryPlanRequest
	4,  // 4: cerbos.svc.v1.CerbosAdminService.AddOrUpdatePolicy:input_type -> cerbos.request.v1.AddOrUpdatePolicyRequest
	5,  // 5: cerbos.svc.v1.CerbosAdminService.ListAuditLogEntries:input_type -> cerbos.request.v1.ListAuditLogEntriesRequest
	6,  // 6: cerbos.svc.v1.CerbosAdminService.ListPolicies:input_type -> cerbos.request.v1.ListPoliciesRequest
	7,  // 7: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundValidate:input_type -> cerbos.request.v1.PlaygroundValidateRequest
	8,  // 8: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundEvaluate:input_type -> cerbos.request.v1.PlaygroundEvaluateRequest
	9,  // 9: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundProxy:input_type -> cerbos.request.v1.PlaygroundProxyRequest
	10, // 10: cerbos.svc.v1.CerbosService.CheckResourceSet:output_type -> cerbos.response.v1.CheckResourceSetResponse
	11, // 11: cerbos.svc.v1.CerbosService.CheckResourceBatch:output_type -> cerbos.response.v1.CheckResourceBatchResponse
	12, // 12: cerbos.svc.v1.CerbosService.ServerInfo:output_type -> cerbos.response.v1.ServerInfoResponse
	13, // 13: cerbos.svc.v1.CerbosService.ResourcesQueryPlan:output_type -> cerbos.response.v1.ResourcesQueryPlanResponse
	14, // 14: cerbos.svc.v1.CerbosAdminService.AddOrUpdatePolicy:output_type -> cerbos.response.v1.AddOrUpdatePolicyResponse
	15, // 15: cerbos.svc.v1.CerbosAdminService.ListAuditLogEntries:output_type -> cerbos.response.v1.ListAuditLogEntriesResponse
	16, // 16: cerbos.svc.v1.CerbosAdminService.ListPolicies:output_type -> cerbos.response.v1.ListPoliciesResponse
	17, // 17: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundValidate:output_type -> cerbos.response.v1.PlaygroundValidateResponse
	18, // 18: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundEvaluate:output_type -> cerbos.response.v1.PlaygroundEvaluateResponse
	19, // 19: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundProxy:output_type -> cerbos.response.v1.PlaygroundProxyResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_cerbos_svc_v1_svc_proto_init() }
func file_cerbos_svc_v1_svc_proto_init() {
	if File_cerbos_svc_v1_svc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cerbos_svc_v1_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_cerbos_svc_v1_svc_proto_goTypes,
		DependencyIndexes: file_cerbos_svc_v1_svc_proto_depIdxs,
	}.Build()
	File_cerbos_svc_v1_svc_proto = out.File
	file_cerbos_svc_v1_svc_proto_rawDesc = nil
	file_cerbos_svc_v1_svc_proto_goTypes = nil
	file_cerbos_svc_v1_svc_proto_depIdxs = nil
}
