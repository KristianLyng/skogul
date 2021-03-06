// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pfe_port_oc.proto

package telemetry

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type InterfacesPfePort struct {
	Interface            []*InterfacesPfePortInterfaceList `protobuf:"bytes,151,rep,name=interface" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *InterfacesPfePort) Reset()         { *m = InterfacesPfePort{} }
func (m *InterfacesPfePort) String() string { return proto.CompactTextString(m) }
func (*InterfacesPfePort) ProtoMessage()    {}
func (*InterfacesPfePort) Descriptor() ([]byte, []int) {
	return fileDescriptor_87ca96f985bf0515, []int{0}
}

func (m *InterfacesPfePort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfacesPfePort.Unmarshal(m, b)
}
func (m *InterfacesPfePort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfacesPfePort.Marshal(b, m, deterministic)
}
func (m *InterfacesPfePort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfacesPfePort.Merge(m, src)
}
func (m *InterfacesPfePort) XXX_Size() int {
	return xxx_messageInfo_InterfacesPfePort.Size(m)
}
func (m *InterfacesPfePort) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfacesPfePort.DiscardUnknown(m)
}

var xxx_messageInfo_InterfacesPfePort proto.InternalMessageInfo

func (m *InterfacesPfePort) GetInterface() []*InterfacesPfePortInterfaceList {
	if m != nil {
		return m.Interface
	}
	return nil
}

type InterfacesPfePortInterfaceList struct {
	State                *InterfacesPfePortInterfaceListStateType `protobuf:"bytes,151,opt,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *InterfacesPfePortInterfaceList) Reset()         { *m = InterfacesPfePortInterfaceList{} }
func (m *InterfacesPfePortInterfaceList) String() string { return proto.CompactTextString(m) }
func (*InterfacesPfePortInterfaceList) ProtoMessage()    {}
func (*InterfacesPfePortInterfaceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_87ca96f985bf0515, []int{0, 0}
}

func (m *InterfacesPfePortInterfaceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfacesPfePortInterfaceList.Unmarshal(m, b)
}
func (m *InterfacesPfePortInterfaceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfacesPfePortInterfaceList.Marshal(b, m, deterministic)
}
func (m *InterfacesPfePortInterfaceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfacesPfePortInterfaceList.Merge(m, src)
}
func (m *InterfacesPfePortInterfaceList) XXX_Size() int {
	return xxx_messageInfo_InterfacesPfePortInterfaceList.Size(m)
}
func (m *InterfacesPfePortInterfaceList) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfacesPfePortInterfaceList.DiscardUnknown(m)
}

var xxx_messageInfo_InterfacesPfePortInterfaceList proto.InternalMessageInfo

func (m *InterfacesPfePortInterfaceList) GetState() *InterfacesPfePortInterfaceListStateType {
	if m != nil {
		return m.State
	}
	return nil
}

type InterfacesPfePortInterfaceListStateType struct {
	Counters             *InterfacesPfePortInterfaceListStateTypeCountersType `protobuf:"bytes,151,opt,name=counters" json:"counters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                             `json:"-"`
	XXX_unrecognized     []byte                                               `json:"-"`
	XXX_sizecache        int32                                                `json:"-"`
}

func (m *InterfacesPfePortInterfaceListStateType) Reset() {
	*m = InterfacesPfePortInterfaceListStateType{}
}
func (m *InterfacesPfePortInterfaceListStateType) String() string { return proto.CompactTextString(m) }
func (*InterfacesPfePortInterfaceListStateType) ProtoMessage()    {}
func (*InterfacesPfePortInterfaceListStateType) Descriptor() ([]byte, []int) {
	return fileDescriptor_87ca96f985bf0515, []int{0, 0, 0}
}

func (m *InterfacesPfePortInterfaceListStateType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfacesPfePortInterfaceListStateType.Unmarshal(m, b)
}
func (m *InterfacesPfePortInterfaceListStateType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfacesPfePortInterfaceListStateType.Marshal(b, m, deterministic)
}
func (m *InterfacesPfePortInterfaceListStateType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfacesPfePortInterfaceListStateType.Merge(m, src)
}
func (m *InterfacesPfePortInterfaceListStateType) XXX_Size() int {
	return xxx_messageInfo_InterfacesPfePortInterfaceListStateType.Size(m)
}
func (m *InterfacesPfePortInterfaceListStateType) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfacesPfePortInterfaceListStateType.DiscardUnknown(m)
}

var xxx_messageInfo_InterfacesPfePortInterfaceListStateType proto.InternalMessageInfo

func (m *InterfacesPfePortInterfaceListStateType) GetCounters() *InterfacesPfePortInterfaceListStateTypeCountersType {
	if m != nil {
		return m.Counters
	}
	return nil
}

type InterfacesPfePortInterfaceListStateTypeCountersType struct {
	InOctets             *uint64  `protobuf:"varint,51,opt,name=in_octets,json=inOctets" json:"in_octets,omitempty"`
	InUnicastPkts        *uint64  `protobuf:"varint,52,opt,name=in_unicast_pkts,json=inUnicastPkts" json:"in_unicast_pkts,omitempty"`
	InBroadcastPkts      *uint64  `protobuf:"varint,53,opt,name=in_broadcast_pkts,json=inBroadcastPkts" json:"in_broadcast_pkts,omitempty"`
	InMulticastPkts      *uint64  `protobuf:"varint,54,opt,name=in_multicast_pkts,json=inMulticastPkts" json:"in_multicast_pkts,omitempty"`
	InDiscards           *uint64  `protobuf:"varint,55,opt,name=in_discards,json=inDiscards" json:"in_discards,omitempty"`
	InErrors             *uint64  `protobuf:"varint,56,opt,name=in_errors,json=inErrors" json:"in_errors,omitempty"`
	InUnknownProtos      *uint32  `protobuf:"varint,57,opt,name=in_unknown_protos,json=inUnknownProtos" json:"in_unknown_protos,omitempty"`
	OutOctets            *uint64  `protobuf:"varint,58,opt,name=out_octets,json=outOctets" json:"out_octets,omitempty"`
	OutUnicastPkts       *uint64  `protobuf:"varint,59,opt,name=out_unicast_pkts,json=outUnicastPkts" json:"out_unicast_pkts,omitempty"`
	OutBroadcastPkts     *uint64  `protobuf:"varint,60,opt,name=out_broadcast_pkts,json=outBroadcastPkts" json:"out_broadcast_pkts,omitempty"`
	OutMulticastPkts     *uint64  `protobuf:"varint,61,opt,name=out_multicast_pkts,json=outMulticastPkts" json:"out_multicast_pkts,omitempty"`
	OutDiscards          *uint64  `protobuf:"varint,62,opt,name=out_discards,json=outDiscards" json:"out_discards,omitempty"`
	OutErrors            *uint64  `protobuf:"varint,63,opt,name=out_errors,json=outErrors" json:"out_errors,omitempty"`
	LastClear            *string  `protobuf:"bytes,64,opt,name=last_clear,json=lastClear" json:"last_clear,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) Reset() {
	*m = InterfacesPfePortInterfaceListStateTypeCountersType{}
}
func (m *InterfacesPfePortInterfaceListStateTypeCountersType) String() string {
	return proto.CompactTextString(m)
}
func (*InterfacesPfePortInterfaceListStateTypeCountersType) ProtoMessage() {}
func (*InterfacesPfePortInterfaceListStateTypeCountersType) Descriptor() ([]byte, []int) {
	return fileDescriptor_87ca96f985bf0515, []int{0, 0, 0, 0}
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfacesPfePortInterfaceListStateTypeCountersType.Unmarshal(m, b)
}
func (m *InterfacesPfePortInterfaceListStateTypeCountersType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfacesPfePortInterfaceListStateTypeCountersType.Marshal(b, m, deterministic)
}
func (m *InterfacesPfePortInterfaceListStateTypeCountersType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfacesPfePortInterfaceListStateTypeCountersType.Merge(m, src)
}
func (m *InterfacesPfePortInterfaceListStateTypeCountersType) XXX_Size() int {
	return xxx_messageInfo_InterfacesPfePortInterfaceListStateTypeCountersType.Size(m)
}
func (m *InterfacesPfePortInterfaceListStateTypeCountersType) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfacesPfePortInterfaceListStateTypeCountersType.DiscardUnknown(m)
}

var xxx_messageInfo_InterfacesPfePortInterfaceListStateTypeCountersType proto.InternalMessageInfo

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetInOctets() uint64 {
	if m != nil && m.InOctets != nil {
		return *m.InOctets
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetInUnicastPkts() uint64 {
	if m != nil && m.InUnicastPkts != nil {
		return *m.InUnicastPkts
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetInBroadcastPkts() uint64 {
	if m != nil && m.InBroadcastPkts != nil {
		return *m.InBroadcastPkts
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetInMulticastPkts() uint64 {
	if m != nil && m.InMulticastPkts != nil {
		return *m.InMulticastPkts
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetInDiscards() uint64 {
	if m != nil && m.InDiscards != nil {
		return *m.InDiscards
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetInErrors() uint64 {
	if m != nil && m.InErrors != nil {
		return *m.InErrors
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetInUnknownProtos() uint32 {
	if m != nil && m.InUnknownProtos != nil {
		return *m.InUnknownProtos
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetOutOctets() uint64 {
	if m != nil && m.OutOctets != nil {
		return *m.OutOctets
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetOutUnicastPkts() uint64 {
	if m != nil && m.OutUnicastPkts != nil {
		return *m.OutUnicastPkts
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetOutBroadcastPkts() uint64 {
	if m != nil && m.OutBroadcastPkts != nil {
		return *m.OutBroadcastPkts
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetOutMulticastPkts() uint64 {
	if m != nil && m.OutMulticastPkts != nil {
		return *m.OutMulticastPkts
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetOutDiscards() uint64 {
	if m != nil && m.OutDiscards != nil {
		return *m.OutDiscards
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetOutErrors() uint64 {
	if m != nil && m.OutErrors != nil {
		return *m.OutErrors
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetLastClear() string {
	if m != nil && m.LastClear != nil {
		return *m.LastClear
	}
	return ""
}

var E_JnprInterfacesPfePortExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*InterfacesPfePort)(nil),
	Field:         60,
	Name:          "jnpr_interfaces_pfe_port_ext",
	Tag:           "bytes,60,opt,name=jnpr_interfaces_pfe_port_ext",
	Filename:      "pfe_port_oc.proto",
}

func init() {
	proto.RegisterType((*InterfacesPfePort)(nil), "interfaces_pfe_port")
	proto.RegisterType((*InterfacesPfePortInterfaceList)(nil), "interfaces_pfe_port.interface_list")
	proto.RegisterType((*InterfacesPfePortInterfaceListStateType)(nil), "interfaces_pfe_port.interface_list.state_type")
	proto.RegisterType((*InterfacesPfePortInterfaceListStateTypeCountersType)(nil), "interfaces_pfe_port.interface_list.state_type.counters_type")
	proto.RegisterExtension(E_JnprInterfacesPfePortExt)
}

func init() { proto.RegisterFile("pfe_port_oc.proto", fileDescriptor_87ca96f985bf0515) }

var fileDescriptor_87ca96f985bf0515 = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0x65, 0xd1, 0xa0, 0xe6, 0x86, 0xb4, 0x74, 0x8a, 0x84, 0x15, 0xa8, 0x08, 0x20, 0x21,
	0x0b, 0x21, 0x2f, 0xca, 0x7f, 0x09, 0x3f, 0x2a, 0x64, 0x01, 0x12, 0x10, 0x19, 0x75, 0xc1, 0x6a,
	0x64, 0x9c, 0x89, 0x34, 0xc4, 0x99, 0x6b, 0xcd, 0xdc, 0x51, 0xdb, 0x17, 0xe0, 0x15, 0xd8, 0xf0,
	0x28, 0x3c, 0x0c, 0x8f, 0x82, 0x66, 0xfc, 0x17, 0x57, 0x59, 0xc0, 0xd2, 0xe7, 0x7e, 0xe7, 0x5c,
	0xcf, 0xd1, 0x85, 0xbd, 0x62, 0x21, 0x78, 0x81, 0x9a, 0x38, 0x66, 0x71, 0xa1, 0x91, 0x70, 0xb4,
	0x4f, 0x22, 0x17, 0x2b, 0x41, 0xfa, 0x9c, 0x13, 0x16, 0xa5, 0x78, 0xe7, 0xf7, 0x65, 0xd8, 0x97,
	0x8a, 0x84, 0x5e, 0xa4, 0x99, 0x30, 0xbc, 0x76, 0xb1, 0x63, 0xe8, 0x37, 0x72, 0xf8, 0x33, 0x18,
	0x5f, 0x8a, 0x06, 0x87, 0x77, 0xe3, 0x0d, 0x64, 0xab, 0xf1, 0x5c, 0x1a, 0x4a, 0x5a, 0xdb, 0xe8,
	0x4f, 0x0f, 0x76, 0xba, 0x53, 0x36, 0x85, 0x9e, 0xa1, 0x94, 0x7c, 0x64, 0x10, 0x0d, 0x0e, 0xe3,
	0x7f, 0x88, 0x8c, 0xbd, 0x83, 0xd3, 0x79, 0x21, 0x92, 0xd2, 0x3d, 0xfa, 0xd1, 0x03, 0x68, 0x55,
	0xf6, 0x15, 0xb6, 0x33, 0xb4, 0xce, 0x64, 0xea, 0xe0, 0xc9, 0xff, 0x05, 0xc7, 0xb5, 0xbf, 0x5c,
	0xd3, 0xc4, 0x8d, 0x7e, 0x6d, 0xc1, 0xb0, 0x33, 0x63, 0x37, 0x5c, 0x33, 0x1c, 0x33, 0x12, 0x64,
	0xc2, 0x87, 0xe3, 0x20, 0xda, 0x4a, 0xb6, 0xa5, 0xfa, 0xec, 0xbf, 0xd9, 0x3d, 0xd8, 0x95, 0x8a,
	0x5b, 0x25, 0xb3, 0xd4, 0x10, 0x2f, 0x96, 0x64, 0xc2, 0x47, 0x1e, 0x19, 0x4a, 0x75, 0x52, 0xaa,
	0xb3, 0x25, 0x19, 0x76, 0x1f, 0xf6, 0xa4, 0xe2, 0xdf, 0x34, 0xa6, 0xf3, 0x96, 0x7c, 0xec, 0xc9,
	0x5d, 0xa9, 0x8e, 0x6b, 0x7d, 0x8d, 0x5d, 0xd9, 0x9c, 0xd6, 0x52, 0x9f, 0xd4, 0xec, 0xc7, 0x5a,
	0xf7, 0xec, 0x2d, 0x18, 0x48, 0xc5, 0xe7, 0xd2, 0x64, 0xa9, 0x9e, 0x9b, 0xf0, 0xa9, 0xa7, 0x40,
	0xaa, 0x77, 0x95, 0x52, 0xfd, 0xbd, 0xd0, 0x1a, 0xb5, 0x09, 0x9f, 0xd5, 0x7f, 0x3f, 0xf5, 0xdf,
	0xd5, 0x26, 0xab, 0x96, 0x0a, 0x4f, 0x15, 0xf7, 0x07, 0x62, 0xc2, 0xe7, 0xe3, 0x20, 0x1a, 0xba,
	0x4d, 0x27, 0xa5, 0x3e, 0xf3, 0x32, 0x3b, 0x00, 0x40, 0x4b, 0x75, 0x0f, 0x47, 0x3e, 0xa9, 0x8f,
	0x96, 0xaa, 0x22, 0x22, 0xb8, 0xea, 0xc6, 0x9d, 0x26, 0x5e, 0x78, 0x68, 0x07, 0x2d, 0xad, 0x57,
	0xf1, 0x00, 0x98, 0x23, 0x2f, 0x74, 0x31, 0xf1, 0xac, 0xcb, 0xe8, 0x96, 0x51, 0xd1, 0x17, 0xda,
	0x78, 0xd9, 0xd0, 0xdd, 0x3a, 0x6e, 0xc3, 0x15, 0x47, 0x37, 0x7d, 0xbc, 0xf2, 0xdc, 0x00, 0x2d,
	0x35, 0x85, 0x54, 0xef, 0xa8, 0x1a, 0x79, 0xdd, 0xbc, 0xa3, 0xaa, 0xe4, 0x00, 0x20, 0x77, 0x6b,
	0xb2, 0x5c, 0xa4, 0x3a, 0x7c, 0x33, 0x0e, 0xa2, 0x7e, 0xd2, 0x77, 0xca, 0x5b, 0x27, 0x1c, 0xad,
	0xe0, 0xe6, 0x77, 0x55, 0x68, 0xbe, 0xe1, 0xd8, 0xb8, 0x38, 0x23, 0x76, 0x3d, 0xfe, 0x60, 0x95,
	0x2c, 0x84, 0xfe, 0x24, 0xe8, 0x14, 0xf5, 0xd2, 0x7c, 0x11, 0xca, 0xb8, 0x45, 0x13, 0x7f, 0xa5,
	0xd7, 0x36, 0x5d, 0x69, 0x12, 0xba, 0xc8, 0xf7, 0xcd, 0x60, 0xb6, 0x10, 0x33, 0xd4, 0x34, 0x3d,
	0xa3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x05, 0x97, 0x08, 0xd6, 0x03, 0x00, 0x00,
}
