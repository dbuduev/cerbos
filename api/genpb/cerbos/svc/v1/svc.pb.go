// Copyright 2021-2025 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: cerbos/svc/v1/svc.proto

package svcv1

import (
	v1 "github.com/cerbos/cerbos/api/genpb/cerbos/request/v1"
	v11 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/visibility"
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
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd6, 0x0a, 0x0a, 0x0d, 0x43, 0x65, 0x72, 0x62,
	0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa7, 0x02, 0x0a, 0x10, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x74, 0x12, 0x2a,
	0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x65, 0x72,
	0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb8, 0x01, 0x92, 0x41, 0x9f, 0x01, 0x12,
	0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x1a, 0x93, 0x01, 0x5b, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x3a, 0x20, 0x55, 0x73, 0x65, 0x20, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x6e, 0x73,
	0x74, 0x65, 0x61, 0x64, 0x5d, 0x20, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x77, 0x68, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x20, 0x61, 0x20, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x20,
	0x68, 0x61, 0x73, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x20,
	0x74, 0x6f, 0x20, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x67,
	0x69, 0x76, 0x65, 0x6e, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x6f, 0x6e, 0x20,
	0x61, 0x20, 0x73, 0x65, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x20, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x58, 0x01, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x12, 0xb6, 0x02, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x2c, 0x2e, 0x63, 0x65, 0x72,
	0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f,
	0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xc1, 0x01, 0x92, 0x41, 0x99, 0x01, 0x12,
	0x14, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x7f, 0x5b, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x3a, 0x20, 0x55, 0x73, 0x65, 0x20, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x6e, 0x73, 0x74, 0x65,
	0x61, 0x64, 0x5d, 0x20, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x61, 0x20, 0x70, 0x72, 0x69, 0x6e,
	0x63, 0x69, 0x70, 0x61, 0x6c, 0x27, 0x73, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x20, 0x62, 0x61, 0x74, 0x63, 0x68, 0x20, 0x6f,
	0x66, 0x20, 0x68, 0x65, 0x74, 0x65, 0x72, 0x6f, 0x67, 0x65, 0x6e, 0x65, 0x6f, 0x75, 0x73, 0x20,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x58, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01,
	0x2a, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x12, 0xf0, 0x01, 0x0a,
	0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12,
	0x28, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63, 0x65, 0x72, 0x62,
	0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x87, 0x01, 0x92, 0x41, 0x65, 0x12, 0x0f, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x52, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x20, 0x61, 0x20, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x27,
	0x73, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x74, 0x6f,
	0x20, 0x61, 0x20, 0x62, 0x61, 0x74, 0x63, 0x68, 0x20, 0x6f, 0x66, 0x20, 0x68, 0x65, 0x74, 0x65,
	0x72, 0x6f, 0x67, 0x65, 0x6e, 0x65, 0x6f, 0x75, 0x73, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12,
	0xc5, 0x01, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24,
	0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x69, 0x92, 0x41,
	0x4e, 0x12, 0x16, 0x47, 0x65, 0x74, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x69, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x34, 0x47, 0x65, 0x74, 0x20, 0x69,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x65, 0x2e, 0x67, 0x2e,
	0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x83, 0x02, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x27, 0x2e, 0x63, 0x65, 0x72, 0x62,
	0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9d, 0x01,
	0x92, 0x41, 0x7c, 0x12, 0x0e, 0x50, 0x6c, 0x61, 0x6e, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x1a, 0x6a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x20, 0x61, 0x20, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x20, 0x70, 0x6c, 0x61, 0x6e, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x6d,
	0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x73, 0x61, 0x74, 0x69, 0x73, 0x66, 0x69, 0x65, 0x64,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x20, 0x61,
	0x20, 0x73, 0x65, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x21, 0x92,
	0x41, 0x1e, 0x12, 0x1c, 0x43, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x20, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x20, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x32, 0xf8, 0x10, 0x0a, 0x12, 0x43, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc9, 0x01, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x4f,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2b, 0x2e,
	0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x65, 0x72,
	0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58, 0x92, 0x41, 0x29, 0x12, 0x16,
	0x41, 0x64, 0x64, 0x20, 0x6f, 0x72, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x62, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01, 0x2a,
	0x5a, 0x12, 0x3a, 0x01, 0x2a, 0x1a, 0x0d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x22, 0x0d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0xb0, 0x01, 0x0a, 0x0f, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x29, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x70,
	0x65, 0x63, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x45, 0x92, 0x41, 0x23, 0x12, 0x10, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x20, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x62, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x69,
	0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x12, 0x9c, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0x92, 0x41, 0x20, 0x12, 0x0d,
	0x4c, 0x69, 0x73, 0x74, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x62, 0x0f, 0x0a,
	0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x8e, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0x23, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f,
	0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x35, 0x92, 0x41, 0x1d, 0x12, 0x0a, 0x47, 0x65, 0x74, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x62, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0xd3, 0x01, 0x0a, 0x0d, 0x44, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x27, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f,
	0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6e, 0x92, 0x41,
	0x21, 0x12, 0x0e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x62, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x44, 0x5a, 0x1a, 0x3a, 0x01, 0x2a, 0x1a, 0x15, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x5a, 0x0f, 0x2a, 0x0d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x15, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x12, 0xbc, 0x01, 0x0a,
	0x0c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x26, 0x2e,
	0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x5a, 0x92, 0x41, 0x20, 0x12, 0x0d, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x62, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x5a, 0x19, 0x3a, 0x01, 0x2a, 0x1a,
	0x14, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x14, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0xc8, 0x01, 0x0a, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x2d, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x4f, 0x92, 0x41, 0x29, 0x12, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x20, 0x6c, 0x6f, 0x67, 0x20, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x62, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x7b, 0x6b,
	0x69, 0x6e, 0x64, 0x7d, 0x30, 0x01, 0x12, 0xc7, 0x01, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x4f, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x2b, 0x2e, 0x63,
	0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x65, 0x72, 0x62,
	0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x56, 0x92, 0x41, 0x27, 0x12, 0x14, 0x41,
	0x64, 0x64, 0x20, 0x6f, 0x72, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x62, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01, 0x2a, 0x5a, 0x12, 0x3a,
	0x01, 0x2a, 0x1a, 0x0d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x22, 0x0d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x12, 0x97, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73,
	0x12, 0x25, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x38, 0x92, 0x41, 0x1f, 0x12, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x73, 0x62, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x12, 0x8e, 0x01, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x23, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f,
	0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x92, 0x41, 0x1d, 0x12, 0x0a, 0x47, 0x65, 0x74, 0x20, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x9a, 0x01, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x26, 0x2e, 0x63,
	0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38,
	0x92, 0x41, 0x20, 0x12, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x62, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x2a, 0x0d, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x9c, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x6c,
	0x6f, 0x61, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x25, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f,
	0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c,
	0x6f, 0x61, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3d, 0x92, 0x41, 0x1f, 0x12, 0x0c, 0x52,
	0x65, 0x6c, 0x6f, 0x61, 0x64, 0x20, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x0f, 0x0a, 0x0d, 0x0a,
	0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x12, 0x13, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x22, 0x92, 0x41, 0x1f, 0x12, 0x1d, 0x43, 0x65,
	0x72, 0x62, 0x6f, 0x73, 0x20, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x32, 0xf7, 0x04, 0x0a, 0x17,
	0x43, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x97, 0x01, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x79,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2c,
	0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63,
	0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6c,
	0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x87, 0x01, 0x0a, 0x0e, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x54, 0x65, 0x73, 0x74, 0x12, 0x28, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6c, 0x61, 0x79,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x12, 0x97, 0x01, 0x0a, 0x12,
	0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x65, 0x12, 0x2c, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2e, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x65, 0x76, 0x61,
	0x6c, 0x75, 0x61, 0x74, 0x65, 0x12, 0x8b, 0x01, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x29, 0x2e, 0x63, 0x65, 0x72, 0x62,
	0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x1a, 0x10, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x49, 0x4e, 0x54,
	0x45, 0x52, 0x4e, 0x41, 0x4c, 0x42, 0xe1, 0x01, 0x92, 0x41, 0x7b, 0x12, 0x3f, 0x0a, 0x06, 0x43,
	0x65, 0x72, 0x62, 0x6f, 0x73, 0x22, 0x2d, 0x0a, 0x06, 0x43, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x12,
	0x12, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e,
	0x64, 0x65, 0x76, 0x1a, 0x0f, 0x69, 0x6e, 0x66, 0x6f, 0x40, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73,
	0x2e, 0x64, 0x65, 0x76, 0x32, 0x06, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x2a, 0x01, 0x02, 0x32,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x5a, 0x11, 0x0a, 0x0f, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x02, 0x08, 0x01, 0x0a, 0x15, 0x64, 0x65, 0x76, 0x2e, 0x63, 0x65, 0x72, 0x62,
	0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x76, 0x63, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73,
	0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x62, 0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x73, 0x76, 0x63, 0x2f, 0x76, 0x31, 0x3b,
	0x73, 0x76, 0x63, 0x76, 0x31, 0xaa, 0x02, 0x11, 0x43, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x41,
	0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x76, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_cerbos_svc_v1_svc_proto_goTypes = []any{
	(*v1.CheckResourceSetRequest)(nil),      // 0: cerbos.request.v1.CheckResourceSetRequest
	(*v1.CheckResourceBatchRequest)(nil),    // 1: cerbos.request.v1.CheckResourceBatchRequest
	(*v1.CheckResourcesRequest)(nil),        // 2: cerbos.request.v1.CheckResourcesRequest
	(*v1.ServerInfoRequest)(nil),            // 3: cerbos.request.v1.ServerInfoRequest
	(*v1.PlanResourcesRequest)(nil),         // 4: cerbos.request.v1.PlanResourcesRequest
	(*v1.AddOrUpdatePolicyRequest)(nil),     // 5: cerbos.request.v1.AddOrUpdatePolicyRequest
	(*v1.InspectPoliciesRequest)(nil),       // 6: cerbos.request.v1.InspectPoliciesRequest
	(*v1.ListPoliciesRequest)(nil),          // 7: cerbos.request.v1.ListPoliciesRequest
	(*v1.GetPolicyRequest)(nil),             // 8: cerbos.request.v1.GetPolicyRequest
	(*v1.DisablePolicyRequest)(nil),         // 9: cerbos.request.v1.DisablePolicyRequest
	(*v1.EnablePolicyRequest)(nil),          // 10: cerbos.request.v1.EnablePolicyRequest
	(*v1.ListAuditLogEntriesRequest)(nil),   // 11: cerbos.request.v1.ListAuditLogEntriesRequest
	(*v1.AddOrUpdateSchemaRequest)(nil),     // 12: cerbos.request.v1.AddOrUpdateSchemaRequest
	(*v1.ListSchemasRequest)(nil),           // 13: cerbos.request.v1.ListSchemasRequest
	(*v1.GetSchemaRequest)(nil),             // 14: cerbos.request.v1.GetSchemaRequest
	(*v1.DeleteSchemaRequest)(nil),          // 15: cerbos.request.v1.DeleteSchemaRequest
	(*v1.ReloadStoreRequest)(nil),           // 16: cerbos.request.v1.ReloadStoreRequest
	(*v1.PlaygroundValidateRequest)(nil),    // 17: cerbos.request.v1.PlaygroundValidateRequest
	(*v1.PlaygroundTestRequest)(nil),        // 18: cerbos.request.v1.PlaygroundTestRequest
	(*v1.PlaygroundEvaluateRequest)(nil),    // 19: cerbos.request.v1.PlaygroundEvaluateRequest
	(*v1.PlaygroundProxyRequest)(nil),       // 20: cerbos.request.v1.PlaygroundProxyRequest
	(*v11.CheckResourceSetResponse)(nil),    // 21: cerbos.response.v1.CheckResourceSetResponse
	(*v11.CheckResourceBatchResponse)(nil),  // 22: cerbos.response.v1.CheckResourceBatchResponse
	(*v11.CheckResourcesResponse)(nil),      // 23: cerbos.response.v1.CheckResourcesResponse
	(*v11.ServerInfoResponse)(nil),          // 24: cerbos.response.v1.ServerInfoResponse
	(*v11.PlanResourcesResponse)(nil),       // 25: cerbos.response.v1.PlanResourcesResponse
	(*v11.AddOrUpdatePolicyResponse)(nil),   // 26: cerbos.response.v1.AddOrUpdatePolicyResponse
	(*v11.InspectPoliciesResponse)(nil),     // 27: cerbos.response.v1.InspectPoliciesResponse
	(*v11.ListPoliciesResponse)(nil),        // 28: cerbos.response.v1.ListPoliciesResponse
	(*v11.GetPolicyResponse)(nil),           // 29: cerbos.response.v1.GetPolicyResponse
	(*v11.DisablePolicyResponse)(nil),       // 30: cerbos.response.v1.DisablePolicyResponse
	(*v11.EnablePolicyResponse)(nil),        // 31: cerbos.response.v1.EnablePolicyResponse
	(*v11.ListAuditLogEntriesResponse)(nil), // 32: cerbos.response.v1.ListAuditLogEntriesResponse
	(*v11.AddOrUpdateSchemaResponse)(nil),   // 33: cerbos.response.v1.AddOrUpdateSchemaResponse
	(*v11.ListSchemasResponse)(nil),         // 34: cerbos.response.v1.ListSchemasResponse
	(*v11.GetSchemaResponse)(nil),           // 35: cerbos.response.v1.GetSchemaResponse
	(*v11.DeleteSchemaResponse)(nil),        // 36: cerbos.response.v1.DeleteSchemaResponse
	(*v11.ReloadStoreResponse)(nil),         // 37: cerbos.response.v1.ReloadStoreResponse
	(*v11.PlaygroundValidateResponse)(nil),  // 38: cerbos.response.v1.PlaygroundValidateResponse
	(*v11.PlaygroundTestResponse)(nil),      // 39: cerbos.response.v1.PlaygroundTestResponse
	(*v11.PlaygroundEvaluateResponse)(nil),  // 40: cerbos.response.v1.PlaygroundEvaluateResponse
	(*v11.PlaygroundProxyResponse)(nil),     // 41: cerbos.response.v1.PlaygroundProxyResponse
}
var file_cerbos_svc_v1_svc_proto_depIdxs = []int32{
	0,  // 0: cerbos.svc.v1.CerbosService.CheckResourceSet:input_type -> cerbos.request.v1.CheckResourceSetRequest
	1,  // 1: cerbos.svc.v1.CerbosService.CheckResourceBatch:input_type -> cerbos.request.v1.CheckResourceBatchRequest
	2,  // 2: cerbos.svc.v1.CerbosService.CheckResources:input_type -> cerbos.request.v1.CheckResourcesRequest
	3,  // 3: cerbos.svc.v1.CerbosService.ServerInfo:input_type -> cerbos.request.v1.ServerInfoRequest
	4,  // 4: cerbos.svc.v1.CerbosService.PlanResources:input_type -> cerbos.request.v1.PlanResourcesRequest
	5,  // 5: cerbos.svc.v1.CerbosAdminService.AddOrUpdatePolicy:input_type -> cerbos.request.v1.AddOrUpdatePolicyRequest
	6,  // 6: cerbos.svc.v1.CerbosAdminService.InspectPolicies:input_type -> cerbos.request.v1.InspectPoliciesRequest
	7,  // 7: cerbos.svc.v1.CerbosAdminService.ListPolicies:input_type -> cerbos.request.v1.ListPoliciesRequest
	8,  // 8: cerbos.svc.v1.CerbosAdminService.GetPolicy:input_type -> cerbos.request.v1.GetPolicyRequest
	9,  // 9: cerbos.svc.v1.CerbosAdminService.DisablePolicy:input_type -> cerbos.request.v1.DisablePolicyRequest
	10, // 10: cerbos.svc.v1.CerbosAdminService.EnablePolicy:input_type -> cerbos.request.v1.EnablePolicyRequest
	11, // 11: cerbos.svc.v1.CerbosAdminService.ListAuditLogEntries:input_type -> cerbos.request.v1.ListAuditLogEntriesRequest
	12, // 12: cerbos.svc.v1.CerbosAdminService.AddOrUpdateSchema:input_type -> cerbos.request.v1.AddOrUpdateSchemaRequest
	13, // 13: cerbos.svc.v1.CerbosAdminService.ListSchemas:input_type -> cerbos.request.v1.ListSchemasRequest
	14, // 14: cerbos.svc.v1.CerbosAdminService.GetSchema:input_type -> cerbos.request.v1.GetSchemaRequest
	15, // 15: cerbos.svc.v1.CerbosAdminService.DeleteSchema:input_type -> cerbos.request.v1.DeleteSchemaRequest
	16, // 16: cerbos.svc.v1.CerbosAdminService.ReloadStore:input_type -> cerbos.request.v1.ReloadStoreRequest
	17, // 17: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundValidate:input_type -> cerbos.request.v1.PlaygroundValidateRequest
	18, // 18: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundTest:input_type -> cerbos.request.v1.PlaygroundTestRequest
	19, // 19: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundEvaluate:input_type -> cerbos.request.v1.PlaygroundEvaluateRequest
	20, // 20: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundProxy:input_type -> cerbos.request.v1.PlaygroundProxyRequest
	21, // 21: cerbos.svc.v1.CerbosService.CheckResourceSet:output_type -> cerbos.response.v1.CheckResourceSetResponse
	22, // 22: cerbos.svc.v1.CerbosService.CheckResourceBatch:output_type -> cerbos.response.v1.CheckResourceBatchResponse
	23, // 23: cerbos.svc.v1.CerbosService.CheckResources:output_type -> cerbos.response.v1.CheckResourcesResponse
	24, // 24: cerbos.svc.v1.CerbosService.ServerInfo:output_type -> cerbos.response.v1.ServerInfoResponse
	25, // 25: cerbos.svc.v1.CerbosService.PlanResources:output_type -> cerbos.response.v1.PlanResourcesResponse
	26, // 26: cerbos.svc.v1.CerbosAdminService.AddOrUpdatePolicy:output_type -> cerbos.response.v1.AddOrUpdatePolicyResponse
	27, // 27: cerbos.svc.v1.CerbosAdminService.InspectPolicies:output_type -> cerbos.response.v1.InspectPoliciesResponse
	28, // 28: cerbos.svc.v1.CerbosAdminService.ListPolicies:output_type -> cerbos.response.v1.ListPoliciesResponse
	29, // 29: cerbos.svc.v1.CerbosAdminService.GetPolicy:output_type -> cerbos.response.v1.GetPolicyResponse
	30, // 30: cerbos.svc.v1.CerbosAdminService.DisablePolicy:output_type -> cerbos.response.v1.DisablePolicyResponse
	31, // 31: cerbos.svc.v1.CerbosAdminService.EnablePolicy:output_type -> cerbos.response.v1.EnablePolicyResponse
	32, // 32: cerbos.svc.v1.CerbosAdminService.ListAuditLogEntries:output_type -> cerbos.response.v1.ListAuditLogEntriesResponse
	33, // 33: cerbos.svc.v1.CerbosAdminService.AddOrUpdateSchema:output_type -> cerbos.response.v1.AddOrUpdateSchemaResponse
	34, // 34: cerbos.svc.v1.CerbosAdminService.ListSchemas:output_type -> cerbos.response.v1.ListSchemasResponse
	35, // 35: cerbos.svc.v1.CerbosAdminService.GetSchema:output_type -> cerbos.response.v1.GetSchemaResponse
	36, // 36: cerbos.svc.v1.CerbosAdminService.DeleteSchema:output_type -> cerbos.response.v1.DeleteSchemaResponse
	37, // 37: cerbos.svc.v1.CerbosAdminService.ReloadStore:output_type -> cerbos.response.v1.ReloadStoreResponse
	38, // 38: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundValidate:output_type -> cerbos.response.v1.PlaygroundValidateResponse
	39, // 39: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundTest:output_type -> cerbos.response.v1.PlaygroundTestResponse
	40, // 40: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundEvaluate:output_type -> cerbos.response.v1.PlaygroundEvaluateResponse
	41, // 41: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundProxy:output_type -> cerbos.response.v1.PlaygroundProxyResponse
	21, // [21:42] is the sub-list for method output_type
	0,  // [0:21] is the sub-list for method input_type
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
