// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: pkg/api/apipb/api.proto

package apipb

import (
	_ "github.com/slntopp/nocloud/pkg/google/api"
	healthpb "github.com/slntopp/nocloud/pkg/health/healthpb"
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

var File_pkg_api_apipb_api_proto protoreflect.FileDescriptor

var file_pkg_api_apipb_api_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x70, 0x62, 0x2f,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6e, 0x6f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x61, 0x70, 0x69, 0x1a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x70, 0x62, 0x2f, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x6d, 0x0a, 0x0d, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x05, 0x50, 0x72,
	0x6f, 0x62, 0x65, 0x12, 0x1b, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x8c, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d,
	0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x70, 0x69, 0x42, 0x08, 0x41, 0x70, 0x69,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6c, 0x6e, 0x74, 0x6f, 0x70, 0x70, 0x2f, 0x6e, 0x6f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x70,
	0x62, 0xa2, 0x02, 0x03, 0x4e, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x61, 0x70, 0x69, 0xca, 0x02, 0x0a, 0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x70,
	0x69, 0xe2, 0x02, 0x16, 0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x70, 0x69, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x4e, 0x6f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_pkg_api_apipb_api_proto_goTypes = []interface{}{
	(*healthpb.ProbeRequest)(nil),  // 0: nocloudhealth.ProbeRequest
	(*healthpb.ProbeResponse)(nil), // 1: nocloudhealth.ProbeResponse
}
var file_pkg_api_apipb_api_proto_depIdxs = []int32{
	0, // 0: nocloudapi.HealthService.Probe:input_type -> nocloudhealth.ProbeRequest
	1, // 1: nocloudapi.HealthService.Probe:output_type -> nocloudhealth.ProbeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_api_apipb_api_proto_init() }
func file_pkg_api_apipb_api_proto_init() {
	if File_pkg_api_apipb_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_api_apipb_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_apipb_api_proto_goTypes,
		DependencyIndexes: file_pkg_api_apipb_api_proto_depIdxs,
	}.Build()
	File_pkg_api_apipb_api_proto = out.File
	file_pkg_api_apipb_api_proto_rawDesc = nil
	file_pkg_api_apipb_api_proto_goTypes = nil
	file_pkg_api_apipb_api_proto_depIdxs = nil
}
