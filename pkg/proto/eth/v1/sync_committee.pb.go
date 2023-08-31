// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: pkg/proto/eth/v1/sync_committee.proto

// Note: largely inspired by
// https://github.com/prysmaticlabs/prysm/tree/develop/proto/eth/v1

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
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

type SyncCommitteeContribution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot              uint64 `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	SubcommitteeIndex uint64 `protobuf:"varint,2,opt,name=subcommittee_index,proto3" json:"subcommittee_index,omitempty"`
	AggregationBits   string `protobuf:"bytes,3,opt,name=aggregation_bits,proto3" json:"aggregation_bits,omitempty"`
	Signature         string `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	BeaconBlockRoot   string `protobuf:"bytes,5,opt,name=beacon_block_root,proto3" json:"beacon_block_root,omitempty"`
}

func (x *SyncCommitteeContribution) Reset() {
	*x = SyncCommitteeContribution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_eth_v1_sync_committee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncCommitteeContribution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncCommitteeContribution) ProtoMessage() {}

func (x *SyncCommitteeContribution) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_eth_v1_sync_committee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncCommitteeContribution.ProtoReflect.Descriptor instead.
func (*SyncCommitteeContribution) Descriptor() ([]byte, []int) {
	return file_pkg_proto_eth_v1_sync_committee_proto_rawDescGZIP(), []int{0}
}

func (x *SyncCommitteeContribution) GetSlot() uint64 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *SyncCommitteeContribution) GetSubcommitteeIndex() uint64 {
	if x != nil {
		return x.SubcommitteeIndex
	}
	return 0
}

func (x *SyncCommitteeContribution) GetAggregationBits() string {
	if x != nil {
		return x.AggregationBits
	}
	return ""
}

func (x *SyncCommitteeContribution) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *SyncCommitteeContribution) GetBeaconBlockRoot() string {
	if x != nil {
		return x.BeaconBlockRoot
	}
	return ""
}

type SyncCommitteeContributionV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot              *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=slot,proto3" json:"slot,omitempty"`
	SubcommitteeIndex *wrapperspb.UInt64Value `protobuf:"bytes,2,opt,name=subcommittee_index,proto3" json:"subcommittee_index,omitempty"`
	AggregationBits   string                  `protobuf:"bytes,3,opt,name=aggregation_bits,proto3" json:"aggregation_bits,omitempty"`
	Signature         string                  `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	BeaconBlockRoot   string                  `protobuf:"bytes,5,opt,name=beacon_block_root,proto3" json:"beacon_block_root,omitempty"`
}

func (x *SyncCommitteeContributionV2) Reset() {
	*x = SyncCommitteeContributionV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_eth_v1_sync_committee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncCommitteeContributionV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncCommitteeContributionV2) ProtoMessage() {}

func (x *SyncCommitteeContributionV2) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_eth_v1_sync_committee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncCommitteeContributionV2.ProtoReflect.Descriptor instead.
func (*SyncCommitteeContributionV2) Descriptor() ([]byte, []int) {
	return file_pkg_proto_eth_v1_sync_committee_proto_rawDescGZIP(), []int{1}
}

func (x *SyncCommitteeContributionV2) GetSlot() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Slot
	}
	return nil
}

func (x *SyncCommitteeContributionV2) GetSubcommitteeIndex() *wrapperspb.UInt64Value {
	if x != nil {
		return x.SubcommitteeIndex
	}
	return nil
}

func (x *SyncCommitteeContributionV2) GetAggregationBits() string {
	if x != nil {
		return x.AggregationBits
	}
	return ""
}

func (x *SyncCommitteeContributionV2) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *SyncCommitteeContributionV2) GetBeaconBlockRoot() string {
	if x != nil {
		return x.BeaconBlockRoot
	}
	return ""
}

var File_pkg_proto_eth_v1_sync_committee_proto protoreflect.FileDescriptor

var file_pkg_proto_eth_v1_sync_committee_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x78, 0x61, 0x74, 0x75, 0x2e, 0x65, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x01, 0x0a, 0x19, 0x53, 0x79, 0x6e, 0x63, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x73, 0x75, 0x62, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x73, 0x75, 0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x62, 0x69, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x62,
	0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x6f, 0x6f, 0x74,
	0x22, 0x95, 0x02, 0x0a, 0x1b, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x32,
	0x12, 0x30, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x73, 0x6c,
	0x6f, 0x74, 0x12, 0x4c, 0x0a, 0x12, 0x73, 0x75, 0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x73, 0x75,
	0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x2a, 0x0a, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x62, 0x69, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x62, 0x65,
	0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x74, 0x68, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x6f,
	0x70, 0x73, 0x2f, 0x78, 0x61, 0x74, 0x75, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_eth_v1_sync_committee_proto_rawDescOnce sync.Once
	file_pkg_proto_eth_v1_sync_committee_proto_rawDescData = file_pkg_proto_eth_v1_sync_committee_proto_rawDesc
)

func file_pkg_proto_eth_v1_sync_committee_proto_rawDescGZIP() []byte {
	file_pkg_proto_eth_v1_sync_committee_proto_rawDescOnce.Do(func() {
		file_pkg_proto_eth_v1_sync_committee_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_eth_v1_sync_committee_proto_rawDescData)
	})
	return file_pkg_proto_eth_v1_sync_committee_proto_rawDescData
}

var file_pkg_proto_eth_v1_sync_committee_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_proto_eth_v1_sync_committee_proto_goTypes = []interface{}{
	(*SyncCommitteeContribution)(nil),   // 0: xatu.eth.v1.SyncCommitteeContribution
	(*SyncCommitteeContributionV2)(nil), // 1: xatu.eth.v1.SyncCommitteeContributionV2
	(*wrapperspb.UInt64Value)(nil),      // 2: google.protobuf.UInt64Value
}
var file_pkg_proto_eth_v1_sync_committee_proto_depIdxs = []int32{
	2, // 0: xatu.eth.v1.SyncCommitteeContributionV2.slot:type_name -> google.protobuf.UInt64Value
	2, // 1: xatu.eth.v1.SyncCommitteeContributionV2.subcommittee_index:type_name -> google.protobuf.UInt64Value
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_proto_eth_v1_sync_committee_proto_init() }
func file_pkg_proto_eth_v1_sync_committee_proto_init() {
	if File_pkg_proto_eth_v1_sync_committee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_eth_v1_sync_committee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncCommitteeContribution); i {
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
		file_pkg_proto_eth_v1_sync_committee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncCommitteeContributionV2); i {
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
			RawDescriptor: file_pkg_proto_eth_v1_sync_committee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_eth_v1_sync_committee_proto_goTypes,
		DependencyIndexes: file_pkg_proto_eth_v1_sync_committee_proto_depIdxs,
		MessageInfos:      file_pkg_proto_eth_v1_sync_committee_proto_msgTypes,
	}.Build()
	File_pkg_proto_eth_v1_sync_committee_proto = out.File
	file_pkg_proto_eth_v1_sync_committee_proto_rawDesc = nil
	file_pkg_proto_eth_v1_sync_committee_proto_goTypes = nil
	file_pkg_proto_eth_v1_sync_committee_proto_depIdxs = nil
}
