// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: master.proto

package kpmasterproto

import (
	installment "github.com/djoonta/kpmasterproto/installment"
	occupation "github.com/djoonta/kpmasterproto/occupation"
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

var File_master_proto protoreflect.FileDescriptor

var file_master_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6f, 0x63,
	0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe7, 0x03, 0x0a, 0x11, 0x4f, 0x63,
	0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5b, 0x0a, 0x06, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x44, 0x12, 0x26, 0x2e, 0x6b, 0x70, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6e, 0x64,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x07,
	0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x27, 0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f,
	0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6b, 0x70,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x63, 0x63, 0x75,
	0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x26, 0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x32, 0x75, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x07, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x12, 0x28, 0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74,
	0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x6f, 0x6f, 0x6e, 0x74, 0x61,
	0x2f, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_master_proto_goTypes = []interface{}{
	(*occupation.OccupationFindIDRequest)(nil),    // 0: kpmasterproto.OccupationFindIDRequest
	(*occupation.OccupationFindAllRequest)(nil),   // 1: kpmasterproto.OccupationFindAllRequest
	(*occupation.OccupationCreateRequest)(nil),    // 2: kpmasterproto.OccupationCreateRequest
	(*occupation.OccupationDeleteRequest)(nil),    // 3: kpmasterproto.OccupationDeleteRequest
	(*occupation.OccupationUpdateRequest)(nil),    // 4: kpmasterproto.OccupationUpdateRequest
	(*installment.InstallmentFindAllRequest)(nil), // 5: kpmasterproto.InstallmentFindAllRequest
	(*occupation.OccupationFindIDResponse)(nil),   // 6: kpmasterproto.OccupationFindIDResponse
	(*occupation.OccupationFindAllResponse)(nil),  // 7: kpmasterproto.OccupationFindAllResponse
	(*occupation.OccupationCreateResponse)(nil),   // 8: kpmasterproto.OccupationCreateResponse
	(*occupation.OccupationDeleteResponse)(nil),   // 9: kpmasterproto.OccupationDeleteResponse
	(*occupation.OccupationUpdateResponse)(nil),   // 10: kpmasterproto.OccupationUpdateResponse
	(*installment.InstallmentFindIDResponse)(nil), // 11: kpmasterproto.InstallmentFindIDResponse
}
var file_master_proto_depIdxs = []int32{
	0,  // 0: kpmasterproto.OccupationService.FindID:input_type -> kpmasterproto.OccupationFindIDRequest
	1,  // 1: kpmasterproto.OccupationService.FindAll:input_type -> kpmasterproto.OccupationFindAllRequest
	2,  // 2: kpmasterproto.OccupationService.Create:input_type -> kpmasterproto.OccupationCreateRequest
	3,  // 3: kpmasterproto.OccupationService.Delete:input_type -> kpmasterproto.OccupationDeleteRequest
	4,  // 4: kpmasterproto.OccupationService.Update:input_type -> kpmasterproto.OccupationUpdateRequest
	5,  // 5: kpmasterproto.InstallmentService.FindAll:input_type -> kpmasterproto.InstallmentFindAllRequest
	6,  // 6: kpmasterproto.OccupationService.FindID:output_type -> kpmasterproto.OccupationFindIDResponse
	7,  // 7: kpmasterproto.OccupationService.FindAll:output_type -> kpmasterproto.OccupationFindAllResponse
	8,  // 8: kpmasterproto.OccupationService.Create:output_type -> kpmasterproto.OccupationCreateResponse
	9,  // 9: kpmasterproto.OccupationService.Delete:output_type -> kpmasterproto.OccupationDeleteResponse
	10, // 10: kpmasterproto.OccupationService.Update:output_type -> kpmasterproto.OccupationUpdateResponse
	11, // 11: kpmasterproto.InstallmentService.FindAll:output_type -> kpmasterproto.InstallmentFindIDResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_master_proto_init() }
func file_master_proto_init() {
	if File_master_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_master_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_master_proto_goTypes,
		DependencyIndexes: file_master_proto_depIdxs,
	}.Build()
	File_master_proto = out.File
	file_master_proto_rawDesc = nil
	file_master_proto_goTypes = nil
	file_master_proto_depIdxs = nil
}
