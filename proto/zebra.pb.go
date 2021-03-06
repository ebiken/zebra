// Code generated by protoc-gen-go.
// source: zebra.proto
// DO NOT EDIT!

/*
Package zebra is a generated protocol buffer package.

It is generated from these files:
	zebra.proto

It has these top-level messages:
	InterfaceRequest
	InterfaceUpdate
	RouterIdRequest
	RouterIdUpdate
	RedistRequest
	HwAddr
	Prefix
	Nexthop
	Address
	Route
*/
package zebra

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AFI int32

const (
	AFI_AFI_IP  AFI = 0
	AFI_AFI_IP6 AFI = 1
	AFI_AFI_MAX AFI = 2
)

var AFI_name = map[int32]string{
	0: "AFI_IP",
	1: "AFI_IP6",
	2: "AFI_MAX",
}
var AFI_value = map[string]int32{
	"AFI_IP":  0,
	"AFI_IP6": 1,
	"AFI_MAX": 2,
}

func (x AFI) String() string {
	return proto.EnumName(AFI_name, int32(x))
}
func (AFI) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RouteType int32

const (
	RouteType_RIB_UNKNOWN   RouteType = 0
	RouteType_RIB_KERNEL    RouteType = 1
	RouteType_RIB_CONNECTED RouteType = 2
	RouteType_RIB_STATIC    RouteType = 3
	RouteType_RIB_RIP       RouteType = 4
	RouteType_RIB_OSPF      RouteType = 5
	RouteType_RIB_ISIS      RouteType = 6
	RouteType_RIB_BGP       RouteType = 7
)

var RouteType_name = map[int32]string{
	0: "RIB_UNKNOWN",
	1: "RIB_KERNEL",
	2: "RIB_CONNECTED",
	3: "RIB_STATIC",
	4: "RIB_RIP",
	5: "RIB_OSPF",
	6: "RIB_ISIS",
	7: "RIB_BGP",
}
var RouteType_value = map[string]int32{
	"RIB_UNKNOWN":   0,
	"RIB_KERNEL":    1,
	"RIB_CONNECTED": 2,
	"RIB_STATIC":    3,
	"RIB_RIP":       4,
	"RIB_OSPF":      5,
	"RIB_ISIS":      6,
	"RIB_BGP":       7,
}

func (x RouteType) String() string {
	return proto.EnumName(RouteType_name, int32(x))
}
func (RouteType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type RouteSubType int32

const (
	RouteSubType_RIB_SUB_OSPF_IA         RouteSubType = 0
	RouteSubType_RIB_SUB_OSPF_NSSA_1     RouteSubType = 1
	RouteSubType_RIB_SUB_OSPF_NSSA_2     RouteSubType = 2
	RouteSubType_RIB_SUB_OSPF_EXTERNAL_1 RouteSubType = 3
	RouteSubType_RIB_SUB_OSPF_EXTERNAL_2 RouteSubType = 4
	RouteSubType_RIB_SUB_BGP_IBGP        RouteSubType = 5
	RouteSubType_RIB_SUB_BGP_EBGP        RouteSubType = 6
	RouteSubType_RIB_SUB_BGP_CONFED      RouteSubType = 7
	RouteSubType_RIB_SUB_ISIS_L1         RouteSubType = 8
	RouteSubType_RIB_SUB_ISIS_L2         RouteSubType = 9
	RouteSubType_RIB_SUB_ISIS_IA         RouteSubType = 10
)

var RouteSubType_name = map[int32]string{
	0:  "RIB_SUB_OSPF_IA",
	1:  "RIB_SUB_OSPF_NSSA_1",
	2:  "RIB_SUB_OSPF_NSSA_2",
	3:  "RIB_SUB_OSPF_EXTERNAL_1",
	4:  "RIB_SUB_OSPF_EXTERNAL_2",
	5:  "RIB_SUB_BGP_IBGP",
	6:  "RIB_SUB_BGP_EBGP",
	7:  "RIB_SUB_BGP_CONFED",
	8:  "RIB_SUB_ISIS_L1",
	9:  "RIB_SUB_ISIS_L2",
	10: "RIB_SUB_ISIS_IA",
}
var RouteSubType_value = map[string]int32{
	"RIB_SUB_OSPF_IA":         0,
	"RIB_SUB_OSPF_NSSA_1":     1,
	"RIB_SUB_OSPF_NSSA_2":     2,
	"RIB_SUB_OSPF_EXTERNAL_1": 3,
	"RIB_SUB_OSPF_EXTERNAL_2": 4,
	"RIB_SUB_BGP_IBGP":        5,
	"RIB_SUB_BGP_EBGP":        6,
	"RIB_SUB_BGP_CONFED":      7,
	"RIB_SUB_ISIS_L1":         8,
	"RIB_SUB_ISIS_L2":         9,
	"RIB_SUB_ISIS_IA":         10,
}

func (x RouteSubType) String() string {
	return proto.EnumName(RouteSubType_name, int32(x))
}
func (RouteSubType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Op int32

const (
	Op_NoOperation              Op = 0
	Op_InterfaceSubscribe       Op = 1
	Op_InterfaceUnsubscribe     Op = 2
	Op_RouterIdSubscribe        Op = 3
	Op_RouterIdUnsubscribe      Op = 4
	Op_RedistSubscribe          Op = 5
	Op_RedistUnsubscribe        Op = 6
	Op_RedistDefaultSubscribe   Op = 7
	Op_RedistDefaultUnsubscribe Op = 8
	Op_RouteAdd                 Op = 9
	Op_RouteDelete              Op = 10
	Op_InterfaceAdd             Op = 11
	Op_InterfaceDelete          Op = 12
	Op_InterfaceAddrAdd         Op = 13
	Op_InterfaceAddrDelete      Op = 14
	Op_InterfaceUp              Op = 15
	Op_InterfaceDown            Op = 16
	Op_InterfaceFlagChange      Op = 17
	Op_InterfaceNameChange      Op = 18
	Op_InterfaceMtuChange       Op = 19
)

var Op_name = map[int32]string{
	0:  "NoOperation",
	1:  "InterfaceSubscribe",
	2:  "InterfaceUnsubscribe",
	3:  "RouterIdSubscribe",
	4:  "RouterIdUnsubscribe",
	5:  "RedistSubscribe",
	6:  "RedistUnsubscribe",
	7:  "RedistDefaultSubscribe",
	8:  "RedistDefaultUnsubscribe",
	9:  "RouteAdd",
	10: "RouteDelete",
	11: "InterfaceAdd",
	12: "InterfaceDelete",
	13: "InterfaceAddrAdd",
	14: "InterfaceAddrDelete",
	15: "InterfaceUp",
	16: "InterfaceDown",
	17: "InterfaceFlagChange",
	18: "InterfaceNameChange",
	19: "InterfaceMtuChange",
}
var Op_value = map[string]int32{
	"NoOperation":              0,
	"InterfaceSubscribe":       1,
	"InterfaceUnsubscribe":     2,
	"RouterIdSubscribe":        3,
	"RouterIdUnsubscribe":      4,
	"RedistSubscribe":          5,
	"RedistUnsubscribe":        6,
	"RedistDefaultSubscribe":   7,
	"RedistDefaultUnsubscribe": 8,
	"RouteAdd":                 9,
	"RouteDelete":              10,
	"InterfaceAdd":             11,
	"InterfaceDelete":          12,
	"InterfaceAddrAdd":         13,
	"InterfaceAddrDelete":      14,
	"InterfaceUp":              15,
	"InterfaceDown":            16,
	"InterfaceFlagChange":      17,
	"InterfaceNameChange":      18,
	"InterfaceMtuChange":       19,
}

func (x Op) String() string {
	return proto.EnumName(Op_name, int32(x))
}
func (Op) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type InterfaceRequest struct {
	Op    Op     `protobuf:"varint,1,opt,name=op,enum=zebra.Op" json:"op,omitempty"`
	VrfId uint32 `protobuf:"varint,2,opt,name=vrf_id,json=vrfId" json:"vrf_id,omitempty"`
}

func (m *InterfaceRequest) Reset()                    { *m = InterfaceRequest{} }
func (m *InterfaceRequest) String() string            { return proto.CompactTextString(m) }
func (*InterfaceRequest) ProtoMessage()               {}
func (*InterfaceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InterfaceRequest) GetOp() Op {
	if m != nil {
		return m.Op
	}
	return Op_NoOperation
}

func (m *InterfaceRequest) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

type InterfaceUpdate struct {
	Op       Op         `protobuf:"varint,1,opt,name=op,enum=zebra.Op" json:"op,omitempty"`
	VrfId    uint32     `protobuf:"varint,2,opt,name=vrf_id,json=vrfId" json:"vrf_id,omitempty"`
	Name     string     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Index    uint32     `protobuf:"varint,4,opt,name=index" json:"index,omitempty"`
	Flags    uint32     `protobuf:"varint,5,opt,name=flags" json:"flags,omitempty"`
	Mtu      uint32     `protobuf:"varint,6,opt,name=mtu" json:"mtu,omitempty"`
	Metric   uint32     `protobuf:"varint,7,opt,name=metric" json:"metric,omitempty"`
	HwAddr   *HwAddr    `protobuf:"bytes,8,opt,name=hw_addr,json=hwAddr" json:"hw_addr,omitempty"`
	AddrIpv4 []*Address `protobuf:"bytes,9,rep,name=addr_ipv4,json=addrIpv4" json:"addr_ipv4,omitempty"`
	AddrIpv6 []*Address `protobuf:"bytes,10,rep,name=addr_ipv6,json=addrIpv6" json:"addr_ipv6,omitempty"`
}

func (m *InterfaceUpdate) Reset()                    { *m = InterfaceUpdate{} }
func (m *InterfaceUpdate) String() string            { return proto.CompactTextString(m) }
func (*InterfaceUpdate) ProtoMessage()               {}
func (*InterfaceUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *InterfaceUpdate) GetOp() Op {
	if m != nil {
		return m.Op
	}
	return Op_NoOperation
}

func (m *InterfaceUpdate) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *InterfaceUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InterfaceUpdate) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *InterfaceUpdate) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *InterfaceUpdate) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *InterfaceUpdate) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *InterfaceUpdate) GetHwAddr() *HwAddr {
	if m != nil {
		return m.HwAddr
	}
	return nil
}

func (m *InterfaceUpdate) GetAddrIpv4() []*Address {
	if m != nil {
		return m.AddrIpv4
	}
	return nil
}

func (m *InterfaceUpdate) GetAddrIpv6() []*Address {
	if m != nil {
		return m.AddrIpv6
	}
	return nil
}

type RouterIdRequest struct {
	Op    Op     `protobuf:"varint,1,opt,name=op,enum=zebra.Op" json:"op,omitempty"`
	VrfId uint32 `protobuf:"varint,2,opt,name=vrf_id,json=vrfId" json:"vrf_id,omitempty"`
}

func (m *RouterIdRequest) Reset()                    { *m = RouterIdRequest{} }
func (m *RouterIdRequest) String() string            { return proto.CompactTextString(m) }
func (*RouterIdRequest) ProtoMessage()               {}
func (*RouterIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RouterIdRequest) GetOp() Op {
	if m != nil {
		return m.Op
	}
	return Op_NoOperation
}

func (m *RouterIdRequest) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

type RouterIdUpdate struct {
	VrfId    uint32 `protobuf:"varint,1,opt,name=vrf_id,json=vrfId" json:"vrf_id,omitempty"`
	RouterId []byte `protobuf:"bytes,2,opt,name=router_id,json=routerId,proto3" json:"router_id,omitempty"`
}

func (m *RouterIdUpdate) Reset()                    { *m = RouterIdUpdate{} }
func (m *RouterIdUpdate) String() string            { return proto.CompactTextString(m) }
func (*RouterIdUpdate) ProtoMessage()               {}
func (*RouterIdUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RouterIdUpdate) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *RouterIdUpdate) GetRouterId() []byte {
	if m != nil {
		return m.RouterId
	}
	return nil
}

type RedistRequest struct {
	Op     Op        `protobuf:"varint,1,opt,name=op,enum=zebra.Op" json:"op,omitempty"`
	Afi    AFI       `protobuf:"varint,2,opt,name=afi,enum=zebra.AFI" json:"afi,omitempty"`
	AllVrf bool      `protobuf:"varint,3,opt,name=all_vrf,json=allVrf" json:"all_vrf,omitempty"`
	VrfId  uint32    `protobuf:"varint,4,opt,name=vrf_id,json=vrfId" json:"vrf_id,omitempty"`
	Type   RouteType `protobuf:"varint,5,opt,name=type,enum=zebra.RouteType" json:"type,omitempty"`
}

func (m *RedistRequest) Reset()                    { *m = RedistRequest{} }
func (m *RedistRequest) String() string            { return proto.CompactTextString(m) }
func (*RedistRequest) ProtoMessage()               {}
func (*RedistRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RedistRequest) GetOp() Op {
	if m != nil {
		return m.Op
	}
	return Op_NoOperation
}

func (m *RedistRequest) GetAfi() AFI {
	if m != nil {
		return m.Afi
	}
	return AFI_AFI_IP
}

func (m *RedistRequest) GetAllVrf() bool {
	if m != nil {
		return m.AllVrf
	}
	return false
}

func (m *RedistRequest) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *RedistRequest) GetType() RouteType {
	if m != nil {
		return m.Type
	}
	return RouteType_RIB_UNKNOWN
}

type HwAddr struct {
	Addr []byte `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (m *HwAddr) Reset()                    { *m = HwAddr{} }
func (m *HwAddr) String() string            { return proto.CompactTextString(m) }
func (*HwAddr) ProtoMessage()               {}
func (*HwAddr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *HwAddr) GetAddr() []byte {
	if m != nil {
		return m.Addr
	}
	return nil
}

type Prefix struct {
	Addr   []byte `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Length uint32 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
}

func (m *Prefix) Reset()                    { *m = Prefix{} }
func (m *Prefix) String() string            { return proto.CompactTextString(m) }
func (*Prefix) ProtoMessage()               {}
func (*Prefix) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Prefix) GetAddr() []byte {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *Prefix) GetLength() uint32 {
	if m != nil {
		return m.Length
	}
	return 0
}

type Nexthop struct {
	Addr    []byte `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Ifindex uint32 `protobuf:"varint,2,opt,name=ifindex" json:"ifindex,omitempty"`
}

func (m *Nexthop) Reset()                    { *m = Nexthop{} }
func (m *Nexthop) String() string            { return proto.CompactTextString(m) }
func (*Nexthop) ProtoMessage()               {}
func (*Nexthop) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Nexthop) GetAddr() []byte {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *Nexthop) GetIfindex() uint32 {
	if m != nil {
		return m.Ifindex
	}
	return 0
}

type Address struct {
	Addr  *Prefix `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Flags uint32  `protobuf:"varint,2,opt,name=flags" json:"flags,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Address) GetAddr() *Prefix {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *Address) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

type Route struct {
	Op       Op           `protobuf:"varint,1,opt,name=op,enum=zebra.Op" json:"op,omitempty"`
	VrfId    uint32       `protobuf:"varint,2,opt,name=vrf_id,json=vrfId" json:"vrf_id,omitempty"`
	Prefix   *Prefix      `protobuf:"bytes,3,opt,name=prefix" json:"prefix,omitempty"`
	Type     RouteType    `protobuf:"varint,4,opt,name=type,enum=zebra.RouteType" json:"type,omitempty"`
	SubType  RouteSubType `protobuf:"varint,5,opt,name=sub_type,json=subType,enum=zebra.RouteSubType" json:"sub_type,omitempty"`
	Distance uint32       `protobuf:"varint,6,opt,name=distance" json:"distance,omitempty"`
	Metric   uint32       `protobuf:"varint,7,opt,name=metric" json:"metric,omitempty"`
	Tag      uint32       `protobuf:"varint,8,opt,name=tag" json:"tag,omitempty"`
	Color    []string     `protobuf:"bytes,9,rep,name=color" json:"color,omitempty"`
	Nexthops []*Nexthop   `protobuf:"bytes,10,rep,name=nexthops" json:"nexthops,omitempty"`
	Aux      []byte       `protobuf:"bytes,11,opt,name=aux,proto3" json:"aux,omitempty"`
}

func (m *Route) Reset()                    { *m = Route{} }
func (m *Route) String() string            { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()               {}
func (*Route) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Route) GetOp() Op {
	if m != nil {
		return m.Op
	}
	return Op_NoOperation
}

func (m *Route) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *Route) GetPrefix() *Prefix {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *Route) GetType() RouteType {
	if m != nil {
		return m.Type
	}
	return RouteType_RIB_UNKNOWN
}

func (m *Route) GetSubType() RouteSubType {
	if m != nil {
		return m.SubType
	}
	return RouteSubType_RIB_SUB_OSPF_IA
}

func (m *Route) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *Route) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Route) GetTag() uint32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *Route) GetColor() []string {
	if m != nil {
		return m.Color
	}
	return nil
}

func (m *Route) GetNexthops() []*Nexthop {
	if m != nil {
		return m.Nexthops
	}
	return nil
}

func (m *Route) GetAux() []byte {
	if m != nil {
		return m.Aux
	}
	return nil
}

func init() {
	proto.RegisterType((*InterfaceRequest)(nil), "zebra.InterfaceRequest")
	proto.RegisterType((*InterfaceUpdate)(nil), "zebra.InterfaceUpdate")
	proto.RegisterType((*RouterIdRequest)(nil), "zebra.RouterIdRequest")
	proto.RegisterType((*RouterIdUpdate)(nil), "zebra.RouterIdUpdate")
	proto.RegisterType((*RedistRequest)(nil), "zebra.RedistRequest")
	proto.RegisterType((*HwAddr)(nil), "zebra.HwAddr")
	proto.RegisterType((*Prefix)(nil), "zebra.Prefix")
	proto.RegisterType((*Nexthop)(nil), "zebra.Nexthop")
	proto.RegisterType((*Address)(nil), "zebra.Address")
	proto.RegisterType((*Route)(nil), "zebra.Route")
	proto.RegisterEnum("zebra.AFI", AFI_name, AFI_value)
	proto.RegisterEnum("zebra.RouteType", RouteType_name, RouteType_value)
	proto.RegisterEnum("zebra.RouteSubType", RouteSubType_name, RouteSubType_value)
	proto.RegisterEnum("zebra.Op", Op_name, Op_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Zebra service

type ZebraClient interface {
	InterfaceService(ctx context.Context, opts ...grpc.CallOption) (Zebra_InterfaceServiceClient, error)
	RouterIdService(ctx context.Context, opts ...grpc.CallOption) (Zebra_RouterIdServiceClient, error)
	RedistService(ctx context.Context, opts ...grpc.CallOption) (Zebra_RedistServiceClient, error)
	RouteService(ctx context.Context, opts ...grpc.CallOption) (Zebra_RouteServiceClient, error)
}

type zebraClient struct {
	cc *grpc.ClientConn
}

func NewZebraClient(cc *grpc.ClientConn) ZebraClient {
	return &zebraClient{cc}
}

func (c *zebraClient) InterfaceService(ctx context.Context, opts ...grpc.CallOption) (Zebra_InterfaceServiceClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Zebra_serviceDesc.Streams[0], c.cc, "/zebra.Zebra/InterfaceService", opts...)
	if err != nil {
		return nil, err
	}
	x := &zebraInterfaceServiceClient{stream}
	return x, nil
}

type Zebra_InterfaceServiceClient interface {
	Send(*InterfaceRequest) error
	Recv() (*InterfaceUpdate, error)
	grpc.ClientStream
}

type zebraInterfaceServiceClient struct {
	grpc.ClientStream
}

func (x *zebraInterfaceServiceClient) Send(m *InterfaceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *zebraInterfaceServiceClient) Recv() (*InterfaceUpdate, error) {
	m := new(InterfaceUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *zebraClient) RouterIdService(ctx context.Context, opts ...grpc.CallOption) (Zebra_RouterIdServiceClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Zebra_serviceDesc.Streams[1], c.cc, "/zebra.Zebra/RouterIdService", opts...)
	if err != nil {
		return nil, err
	}
	x := &zebraRouterIdServiceClient{stream}
	return x, nil
}

type Zebra_RouterIdServiceClient interface {
	Send(*RouterIdRequest) error
	Recv() (*RouterIdUpdate, error)
	grpc.ClientStream
}

type zebraRouterIdServiceClient struct {
	grpc.ClientStream
}

func (x *zebraRouterIdServiceClient) Send(m *RouterIdRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *zebraRouterIdServiceClient) Recv() (*RouterIdUpdate, error) {
	m := new(RouterIdUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *zebraClient) RedistService(ctx context.Context, opts ...grpc.CallOption) (Zebra_RedistServiceClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Zebra_serviceDesc.Streams[2], c.cc, "/zebra.Zebra/RedistService", opts...)
	if err != nil {
		return nil, err
	}
	x := &zebraRedistServiceClient{stream}
	return x, nil
}

type Zebra_RedistServiceClient interface {
	Send(*RedistRequest) error
	Recv() (*Route, error)
	grpc.ClientStream
}

type zebraRedistServiceClient struct {
	grpc.ClientStream
}

func (x *zebraRedistServiceClient) Send(m *RedistRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *zebraRedistServiceClient) Recv() (*Route, error) {
	m := new(Route)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *zebraClient) RouteService(ctx context.Context, opts ...grpc.CallOption) (Zebra_RouteServiceClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Zebra_serviceDesc.Streams[3], c.cc, "/zebra.Zebra/RouteService", opts...)
	if err != nil {
		return nil, err
	}
	x := &zebraRouteServiceClient{stream}
	return x, nil
}

type Zebra_RouteServiceClient interface {
	Send(*Route) error
	Recv() (*Route, error)
	grpc.ClientStream
}

type zebraRouteServiceClient struct {
	grpc.ClientStream
}

func (x *zebraRouteServiceClient) Send(m *Route) error {
	return x.ClientStream.SendMsg(m)
}

func (x *zebraRouteServiceClient) Recv() (*Route, error) {
	m := new(Route)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Zebra service

type ZebraServer interface {
	InterfaceService(Zebra_InterfaceServiceServer) error
	RouterIdService(Zebra_RouterIdServiceServer) error
	RedistService(Zebra_RedistServiceServer) error
	RouteService(Zebra_RouteServiceServer) error
}

func RegisterZebraServer(s *grpc.Server, srv ZebraServer) {
	s.RegisterService(&_Zebra_serviceDesc, srv)
}

func _Zebra_InterfaceService_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ZebraServer).InterfaceService(&zebraInterfaceServiceServer{stream})
}

type Zebra_InterfaceServiceServer interface {
	Send(*InterfaceUpdate) error
	Recv() (*InterfaceRequest, error)
	grpc.ServerStream
}

type zebraInterfaceServiceServer struct {
	grpc.ServerStream
}

func (x *zebraInterfaceServiceServer) Send(m *InterfaceUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func (x *zebraInterfaceServiceServer) Recv() (*InterfaceRequest, error) {
	m := new(InterfaceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Zebra_RouterIdService_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ZebraServer).RouterIdService(&zebraRouterIdServiceServer{stream})
}

type Zebra_RouterIdServiceServer interface {
	Send(*RouterIdUpdate) error
	Recv() (*RouterIdRequest, error)
	grpc.ServerStream
}

type zebraRouterIdServiceServer struct {
	grpc.ServerStream
}

func (x *zebraRouterIdServiceServer) Send(m *RouterIdUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func (x *zebraRouterIdServiceServer) Recv() (*RouterIdRequest, error) {
	m := new(RouterIdRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Zebra_RedistService_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ZebraServer).RedistService(&zebraRedistServiceServer{stream})
}

type Zebra_RedistServiceServer interface {
	Send(*Route) error
	Recv() (*RedistRequest, error)
	grpc.ServerStream
}

type zebraRedistServiceServer struct {
	grpc.ServerStream
}

func (x *zebraRedistServiceServer) Send(m *Route) error {
	return x.ServerStream.SendMsg(m)
}

func (x *zebraRedistServiceServer) Recv() (*RedistRequest, error) {
	m := new(RedistRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Zebra_RouteService_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ZebraServer).RouteService(&zebraRouteServiceServer{stream})
}

type Zebra_RouteServiceServer interface {
	Send(*Route) error
	Recv() (*Route, error)
	grpc.ServerStream
}

type zebraRouteServiceServer struct {
	grpc.ServerStream
}

func (x *zebraRouteServiceServer) Send(m *Route) error {
	return x.ServerStream.SendMsg(m)
}

func (x *zebraRouteServiceServer) Recv() (*Route, error) {
	m := new(Route)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Zebra_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zebra.Zebra",
	HandlerType: (*ZebraServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "InterfaceService",
			Handler:       _Zebra_InterfaceService_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RouterIdService",
			Handler:       _Zebra_RouterIdService_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RedistService",
			Handler:       _Zebra_RedistService_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RouteService",
			Handler:       _Zebra_RouteService_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "zebra.proto",
}

func init() { proto.RegisterFile("zebra.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1042 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x36, 0x29, 0x89, 0x94, 0x46, 0x3f, 0xde, 0xac, 0xff, 0x58, 0xc7, 0x07, 0x57, 0x68, 0x0b,
	0xc1, 0x01, 0x8c, 0x58, 0x35, 0x5c, 0xf4, 0x28, 0xeb, 0x27, 0x25, 0xe2, 0x50, 0x02, 0x65, 0xb7,
	0x41, 0x2f, 0x04, 0x25, 0x2e, 0x6d, 0x02, 0x34, 0xc9, 0x92, 0x94, 0xec, 0xf4, 0x5e, 0xf4, 0xde,
	0x07, 0xe8, 0x2b, 0xf4, 0x15, 0xfa, 0x08, 0x7d, 0xa4, 0x62, 0x97, 0x4b, 0x6a, 0x25, 0x3b, 0x41,
	0x9b, 0x93, 0x76, 0xbe, 0xf9, 0x66, 0x76, 0x76, 0x7e, 0x38, 0x82, 0xfa, 0xaf, 0x64, 0x16, 0xdb,
	0xa7, 0x51, 0x1c, 0xa6, 0x21, 0xae, 0x30, 0xa1, 0x3d, 0x00, 0xa4, 0x07, 0x29, 0x89, 0x5d, 0x7b,
	0x4e, 0x4c, 0xf2, 0xcb, 0x82, 0x24, 0x29, 0xfe, 0x02, 0xe4, 0x30, 0xd2, 0xa4, 0x63, 0xa9, 0xd3,
	0xea, 0xd6, 0x4e, 0x33, 0xa3, 0x71, 0x64, 0xca, 0x61, 0x84, 0xf7, 0x40, 0x59, 0xc6, 0xae, 0xe5,
	0x39, 0x9a, 0x7c, 0x2c, 0x75, 0x9a, 0x66, 0x65, 0x19, 0xbb, 0xba, 0xd3, 0xfe, 0x4b, 0x86, 0xed,
	0xc2, 0xcd, 0x4d, 0xe4, 0xd8, 0x29, 0xf9, 0xff, 0x5e, 0x30, 0x86, 0x72, 0x60, 0xdf, 0x13, 0xad,
	0x74, 0x2c, 0x75, 0x6a, 0x26, 0x3b, 0xe3, 0x5d, 0xa8, 0x78, 0x81, 0x43, 0x1e, 0xb5, 0x72, 0xc6,
	0x64, 0x02, 0x45, 0x5d, 0xdf, 0xbe, 0x4d, 0xb4, 0x4a, 0x86, 0x32, 0x01, 0x23, 0x28, 0xdd, 0xa7,
	0x0b, 0x4d, 0x61, 0x18, 0x3d, 0xe2, 0x7d, 0x50, 0xee, 0x49, 0x1a, 0x7b, 0x73, 0x4d, 0x65, 0x20,
	0x97, 0xf0, 0x37, 0xa0, 0xde, 0x3d, 0x58, 0xb6, 0xe3, 0xc4, 0x5a, 0xf5, 0x58, 0xea, 0xd4, 0xbb,
	0x4d, 0x1e, 0xe0, 0x0f, 0x0f, 0x3d, 0xc7, 0x89, 0x4d, 0xe5, 0x8e, 0xfd, 0xe2, 0x57, 0x50, 0xa3,
	0x24, 0xcb, 0x8b, 0x96, 0xe7, 0x5a, 0xed, 0xb8, 0xd4, 0xa9, 0x77, 0x5b, 0x9c, 0x49, 0xf5, 0x24,
	0x49, 0xcc, 0x2a, 0x25, 0xe8, 0xd1, 0xf2, 0x5c, 0x24, 0x5f, 0x68, 0xf0, 0x49, 0xf2, 0x45, 0xbb,
	0x0f, 0xdb, 0x66, 0xb8, 0x48, 0x49, 0xac, 0x3b, 0x9f, 0x9f, 0xf6, 0x01, 0xb4, 0x72, 0x27, 0x3c,
	0xe9, 0x2b, 0xa2, 0x24, 0x66, 0xf6, 0x25, 0xd4, 0x62, 0x46, 0xcc, 0x5d, 0x34, 0xcc, 0x6a, 0xcc,
	0x2d, 0xdb, 0x7f, 0x4a, 0xd0, 0x34, 0x89, 0xe3, 0x25, 0xe9, 0x7f, 0x88, 0xe4, 0x08, 0x4a, 0xb6,
	0xeb, 0x31, 0x1f, 0xad, 0x2e, 0xe4, 0xcf, 0x1b, 0xe9, 0x26, 0x85, 0xf1, 0x01, 0xa8, 0xb6, 0xef,
	0x5b, 0xcb, 0xd8, 0x65, 0x45, 0xac, 0x9a, 0x8a, 0xed, 0xfb, 0x3f, 0xc6, 0xae, 0x10, 0x57, 0x59,
	0x8c, 0xeb, 0x2b, 0x28, 0xa7, 0x1f, 0x22, 0xc2, 0xca, 0xd8, 0xea, 0x22, 0xee, 0x8e, 0xbd, 0xe9,
	0xfa, 0x43, 0x44, 0x4c, 0xa6, 0x6d, 0x1f, 0x81, 0x92, 0xd5, 0x85, 0x76, 0x08, 0x2b, 0x9a, 0xc4,
	0x9e, 0xc0, 0xce, 0xed, 0x73, 0x50, 0x26, 0x31, 0x71, 0xbd, 0xc7, 0xe7, 0xb4, 0xb4, 0x03, 0x7c,
	0x12, 0xdc, 0xa6, 0x77, 0x3c, 0x73, 0x5c, 0x6a, 0x7f, 0x07, 0xaa, 0x41, 0x1e, 0xd3, 0xbb, 0x30,
	0x7a, 0xd6, 0x4c, 0x03, 0xd5, 0x73, 0xb3, 0xc6, 0xcb, 0xec, 0x72, 0xb1, 0x7d, 0x09, 0x2a, 0xaf,
	0x26, 0xfe, 0x52, 0x30, 0x5c, 0xb5, 0x50, 0x16, 0x0c, 0xf7, 0x53, 0x34, 0xaa, 0x2c, 0x34, 0x6a,
	0xfb, 0x1f, 0x19, 0x2a, 0xec, 0x91, 0x9f, 0x31, 0x24, 0x5f, 0x83, 0x12, 0xb1, 0x1b, 0x58, 0x86,
	0x9f, 0x5c, 0xcb, 0x95, 0x45, 0x66, 0xcb, 0x9f, 0xca, 0x2c, 0x3e, 0x85, 0x6a, 0xb2, 0x98, 0x59,
	0x42, 0x0d, 0x76, 0x44, 0xe6, 0x74, 0x31, 0x63, 0x64, 0x35, 0xc9, 0x0e, 0xf8, 0x10, 0xaa, 0xb4,
	0x4f, 0xec, 0x60, 0x4e, 0xf8, 0x98, 0x15, 0xf2, 0x47, 0x67, 0x0d, 0x41, 0x29, 0xb5, 0x6f, 0xd9,
	0x9c, 0x35, 0x4d, 0x7a, 0xa4, 0x49, 0x99, 0x87, 0x7e, 0x18, 0xb3, 0x89, 0xaa, 0x99, 0x99, 0x80,
	0x4f, 0xa0, 0x1a, 0x64, 0x15, 0x49, 0x36, 0xa6, 0x87, 0x17, 0xca, 0x2c, 0xf4, 0xd4, 0xa7, 0xbd,
	0x78, 0xd4, 0xea, 0xac, 0x62, 0xf4, 0x78, 0xf2, 0x0a, 0x4a, 0xbd, 0x91, 0x8e, 0x01, 0x94, 0xde,
	0x48, 0xb7, 0xf4, 0x09, 0xda, 0xc2, 0x75, 0x50, 0xb3, 0xf3, 0x05, 0x92, 0x72, 0xe1, 0x5d, 0xef,
	0x3d, 0x92, 0x4f, 0x7e, 0x93, 0xa0, 0x56, 0xa4, 0x02, 0x6f, 0x43, 0xdd, 0xd4, 0x2f, 0xad, 0x1b,
	0xe3, 0xad, 0x31, 0xfe, 0xc9, 0x40, 0x5b, 0xb8, 0x05, 0x40, 0x81, 0xb7, 0x43, 0xd3, 0x18, 0x5e,
	0x21, 0x09, 0xbf, 0x80, 0x26, 0x95, 0xfb, 0x63, 0xc3, 0x18, 0xf6, 0xaf, 0x87, 0x03, 0x24, 0xe7,
	0x94, 0xe9, 0x75, 0xef, 0x5a, 0xef, 0xa3, 0x12, 0x75, 0x4f, 0x65, 0x53, 0x9f, 0xa0, 0x32, 0x6e,
	0x40, 0x95, 0x0a, 0xe3, 0xe9, 0x64, 0x84, 0x2a, 0xb9, 0xa4, 0x4f, 0xf5, 0x29, 0x52, 0x72, 0xe2,
	0xe5, 0x9b, 0x09, 0x52, 0x4f, 0xfe, 0x90, 0xa1, 0x21, 0x26, 0x1a, 0xef, 0xc0, 0x36, 0x73, 0x7b,
	0x93, 0x59, 0x5b, 0x7a, 0x0f, 0x6d, 0xe1, 0x03, 0xd8, 0x59, 0x03, 0x8d, 0xe9, 0xb4, 0x67, 0x9d,
	0x21, 0xe9, 0x79, 0x45, 0x17, 0xc9, 0xf8, 0x25, 0x1c, 0xac, 0x29, 0x86, 0xef, 0xaf, 0x87, 0xa6,
	0xd1, 0xbb, 0xb2, 0xce, 0x50, 0xe9, 0xe3, 0xca, 0x2e, 0x2a, 0xe3, 0x5d, 0x40, 0xb9, 0xf2, 0xf2,
	0xcd, 0xc4, 0xd2, 0x69, 0x9c, 0x95, 0x4d, 0x74, 0x48, 0x51, 0x05, 0xef, 0x03, 0x16, 0xd1, 0xfe,
	0xd8, 0x18, 0x0d, 0x07, 0x48, 0x15, 0x1f, 0x41, 0x1f, 0x6d, 0x5d, 0x9d, 0xa1, 0xea, 0x53, 0xb0,
	0x8b, 0x6a, 0x4f, 0x40, 0xbd, 0x87, 0xe0, 0xe4, 0xef, 0x12, 0xc8, 0xe3, 0x88, 0x56, 0xc5, 0x08,
	0xc7, 0x11, 0x89, 0xed, 0xd4, 0x0b, 0x03, 0xb4, 0x45, 0xaf, 0x2b, 0x56, 0xcc, 0x74, 0x31, 0x4b,
	0xe6, 0xb1, 0x37, 0x23, 0x48, 0xc2, 0x1a, 0xec, 0xae, 0x56, 0x4f, 0x90, 0x14, 0x1a, 0x19, 0xef,
	0xc1, 0x8b, 0xfc, 0xf3, 0xb8, 0x32, 0x28, 0xb1, 0xb4, 0xe5, 0x5f, 0x4d, 0x81, 0x5f, 0x66, 0xe1,
	0xb0, 0xef, 0xe0, 0x8a, 0x5d, 0x61, 0x4e, 0x18, 0x28, 0x72, 0x15, 0x7c, 0x08, 0xfb, 0x19, 0x3c,
	0x20, 0xae, 0xbd, 0xf0, 0x05, 0x13, 0x15, 0x1f, 0x81, 0xb6, 0xa6, 0x13, 0x2d, 0xab, 0xac, 0x1f,
	0xe8, 0xf5, 0x3d, 0xc7, 0x41, 0x35, 0xd6, 0x7c, 0x54, 0x1a, 0x10, 0x9f, 0xa4, 0x04, 0x01, 0x46,
	0xd0, 0x28, 0x9e, 0x43, 0x29, 0x75, 0x1a, 0x56, 0x81, 0x70, 0x5a, 0x83, 0x96, 0x44, 0xa4, 0xc5,
	0x94, 0xda, 0xa4, 0x4f, 0x5b, 0x43, 0x39, 0xbd, 0x45, 0xaf, 0x11, 0xf6, 0x33, 0xda, 0xa6, 0x3d,
	0xbd, 0x72, 0x1a, 0x3e, 0x04, 0x08, 0xad, 0x19, 0x8f, 0x7c, 0xfb, 0xb6, 0x7f, 0x67, 0x07, 0xb7,
	0x04, 0xbd, 0x58, 0x53, 0x18, 0xf6, 0x3d, 0xe1, 0x0a, 0xbc, 0x56, 0x92, 0x77, 0xe9, 0x82, 0xe3,
	0x3b, 0xdd, 0xdf, 0x65, 0xa8, 0xfc, 0x4c, 0x47, 0x17, 0xeb, 0x42, 0x98, 0x53, 0x12, 0x2f, 0xbd,
	0x39, 0xc1, 0x07, 0x7c, 0xac, 0x37, 0xff, 0x77, 0x1c, 0xee, 0x6f, 0x2a, 0xb2, 0xa5, 0xd6, 0xde,
	0xea, 0x48, 0xaf, 0x25, 0x3c, 0x5a, 0x6d, 0xcc, 0xdc, 0xd3, 0xbe, 0xf8, 0xb1, 0x5a, 0x6d, 0xd2,
	0xc3, 0xbd, 0x0d, 0x7c, 0xcd, 0xcf, 0xf7, 0xf9, 0xb6, 0xcb, 0xbd, 0xec, 0xe6, 0x6c, 0x71, 0x07,
	0x1e, 0x36, 0x44, 0x1f, 0xdc, 0xf4, 0x75, 0x3e, 0xae, 0xdc, 0x72, 0x8d, 0xf3, 0x9c, 0xc5, 0x4c,
	0x61, 0x7f, 0xb6, 0xbe, 0xfd, 0x37, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xc3, 0xd4, 0xf5, 0x7b, 0x09,
	0x00, 0x00,
}
