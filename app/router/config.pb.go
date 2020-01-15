package router

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	net "v2ray.com/core/common/net"
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

// Type of domain value.
type Domain_Type int32

const (
	// The value is used as is.
	Domain_Plain Domain_Type = 0
	// The value is used as a regular expression.
	Domain_Regex Domain_Type = 1
	// The value is a root domain.
	Domain_Domain Domain_Type = 2
	// The value is a domain.
	Domain_Full Domain_Type = 3
)

var Domain_Type_name = map[int32]string{
	0: "Plain",
	1: "Regex",
	2: "Domain",
	3: "Full",
}

var Domain_Type_value = map[string]int32{
	"Plain":  0,
	"Regex":  1,
	"Domain": 2,
	"Full":   3,
}

func (x Domain_Type) String() string {
	return proto.EnumName(Domain_Type_name, int32(x))
}

func (Domain_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6b1608360690c5fc, []int{0, 0}
}

type Config_DomainStrategy int32

const (
	// Use domain as is.
	Config_AsIs Config_DomainStrategy = 0
	// Always resolve IP for domains.
	Config_UseIp Config_DomainStrategy = 1
	// Resolve to IP if the domain doesn't match any rules.
	Config_IpIfNonMatch Config_DomainStrategy = 2
	// Resolve to IP if any rule requires IP matching.
	Config_IpOnDemand Config_DomainStrategy = 3
)

var Config_DomainStrategy_name = map[int32]string{
	0: "AsIs",
	1: "UseIp",
	2: "IpIfNonMatch",
	3: "IpOnDemand",
}

var Config_DomainStrategy_value = map[string]int32{
	"AsIs":         0,
	"UseIp":        1,
	"IpIfNonMatch": 2,
	"IpOnDemand":   3,
}

func (x Config_DomainStrategy) String() string {
	return proto.EnumName(Config_DomainStrategy_name, int32(x))
}

func (Config_DomainStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6b1608360690c5fc, []int{9, 0}
}

// Domain for routing decision.
type Domain struct {
	// Domain matching type.
	Type Domain_Type `protobuf:"varint,1,opt,name=type,proto3,enum=v2ray.core.app.router.Domain_Type" json:"type,omitempty"`
	// Domain value.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// Attributes of this domain. May be used for filtering.
	Attribute            []*Domain_Attribute `protobuf:"bytes,3,rep,name=attribute,proto3" json:"attribute,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Domain) Reset()         { *m = Domain{} }
func (m *Domain) String() string { return proto.CompactTextString(m) }
func (*Domain) ProtoMessage()    {}
func (*Domain) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b1608360690c5fc, []int{0}
}

func (m *Domain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Domain.Unmarshal(m, b)
}
func (m *Domain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Domain.Marshal(b, m, deterministic)
}
func (m *Domain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Domain.Merge(m, src)
}
func (m *Domain) XXX_Size() int {
	return xxx_messageInfo_Domain.Size(m)
}
func (m *Domain) XXX_DiscardUnknown() {
	xxx_messageInfo_Domain.DiscardUnknown(m)
}

var xxx_messageInfo_Domain proto.InternalMessageInfo

func (m *Domain) GetType() Domain_Type {
	if m != nil {
		return m.Type
	}
	return Domain_Plain
}

func (m *Domain) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Domain) GetAttribute() []*Domain_Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

type Domain_Attribute struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Types that are valid to be assigned to TypedValue:
	//	*Domain_Attribute_BoolValue
	//	*Domain_Attribute_IntValue
	TypedValue           isDomain_Attribute_TypedValue `protobuf_oneof:"typed_value"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *Domain_Attribute) Reset()         { *m = Domain_Attribute{} }
func (m *Domain_Attribute) String() string { return proto.CompactTextString(m) }
func (*Domain_Attribute) ProtoMessage()    {}
func (*Domain_Attribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b1608360690c5fc, []int{0, 0}
}

func (m *Domain_Attribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Domain_Attribute.Unmarshal(m, b)
}
func (m *Domain_Attribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Domain_Attribute.Marshal(b, m, deterministic)
}
func (m *Domain_Attribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Domain_Attribute.Merge(m, src)
}
func (m *Domain_Attribute) XXX_Size() int {
	return xxx_messageInfo_Domain_Attribute.Size(m)
}
func (m *Domain_Attribute) XXX_DiscardUnknown() {
	xxx_messageInfo_Domain_Attribute.DiscardUnknown(m)
}

var xxx_messageInfo_Domain_Attribute proto.InternalMessageInfo

func (m *Domain_Attribute) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type isDomain_Attribute_TypedValue interface {
	isDomain_Attribute_TypedValue()
}

type Domain_Attribute_BoolValue struct {
	BoolValue bool `protobuf:"varint,2,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type Domain_Attribute_IntValue struct {
	IntValue int64 `protobuf:"varint,3,opt,name=int_value,json=intValue,proto3,oneof"`
}

func (*Domain_Attribute_BoolValue) isDomain_Attribute_TypedValue() {}

func (*Domain_Attribute_IntValue) isDomain_Attribute_TypedValue() {}

func (m *Domain_Attribute) GetTypedValue() isDomain_Attribute_TypedValue {
	if m != nil {
		return m.TypedValue
	}
	return nil
}

func (m *Domain_Attribute) GetBoolValue() bool {
	if x, ok := m.GetTypedValue().(*Domain_Attribute_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *Domain_Attribute) GetIntValue() int64 {
	if x, ok := m.GetTypedValue().(*Domain_Attribute_IntValue); ok {
		return x.IntValue
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Domain_Attribute) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Domain_Attribute_BoolValue)(nil),
		(*Domain_Attribute_IntValue)(nil),
	}
}

// IP for routing decision, in CIDR form.
type CIDR struct {
	// IP address, should be either 4 or 16 bytes.
	Ip []byte `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	// Number of leading ones in the network mask.
	Prefix               uint32   `protobuf:"varint,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CIDR) Reset()         { *m = CIDR{} }
func (m *CIDR) String() string { return proto.CompactTextString(m) }
func (*CIDR) ProtoMessage()    {}
func (*CIDR) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b1608360690c5fc, []int{1}
}

func (m *CIDR) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIDR.Unmarshal(m, b)
}
func (m *CIDR) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIDR.Marshal(b, m, deterministic)
}
func (m *CIDR) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIDR.Merge(m, src)
}
func (m *CIDR) XXX_Size() int {
	return xxx_messageInfo_CIDR.Size(m)
}
func (m *CIDR) XXX_DiscardUnknown() {
	xxx_messageInfo_CIDR.DiscardUnknown(m)
}

var xxx_messageInfo_CIDR proto.InternalMessageInfo

func (m *CIDR) GetIp() []byte {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *CIDR) GetPrefix() uint32 {
	if m != nil {
		return m.Prefix
	}
	return 0
}

type GeoIP struct {
	CountryCode          string   `protobuf:"bytes,1,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Cidr                 []*CIDR  `protobuf:"bytes,2,rep,name=cidr,proto3" json:"cidr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeoIP) Reset()         { *m = GeoIP{} }
func (m *GeoIP) String() string { return proto.CompactTextString(m) }
func (*GeoIP) ProtoMessage()    {}
func (*GeoIP) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b1608360690c5fc, []int{2}
}

func (m *GeoIP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeoIP.Unmarshal(m, b)
}
func (m *GeoIP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeoIP.Marshal(b, m, deterministic)
}
func (m *GeoIP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeoIP.Merge(m, src)
}
func (m *GeoIP) XXX_Size() int {
	return xxx_messageInfo_GeoIP.Size(m)
}
func (m *GeoIP) XXX_DiscardUnknown() {
	xxx_messageInfo_GeoIP.DiscardUnknown(m)
}

var xxx_messageInfo_GeoIP proto.InternalMessageInfo

func (m *GeoIP) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func (m *GeoIP) GetCidr() []*CIDR {
	if m != nil {
		return m.Cidr
	}
	return nil
}

type GeoIPList struct {
	Entry                []*GeoIP `protobuf:"bytes,1,rep,name=entry,proto3" json:"entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeoIPList) Reset()         { *m = GeoIPList{} }
func (m *GeoIPList) String() string { return proto.CompactTextString(m) }
func (*GeoIPList) ProtoMessage()    {}
func (*GeoIPList) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b1608360690c5fc, []int{3}
}

func (m *GeoIPList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeoIPList.Unmarshal(m, b)
}
func (m *GeoIPList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeoIPList.Marshal(b, m, deterministic)
}
func (m *GeoIPList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeoIPList.Merge(m, src)
}
func (m *GeoIPList) XXX_Size() int {
	return xxx_messageInfo_GeoIPList.Size(m)
}
func (m *GeoIPList) XXX_DiscardUnknown() {
	xxx_messageInfo_GeoIPList.DiscardUnknown(m)
}

var xxx_messageInfo_GeoIPList proto.InternalMessageInfo

func (m *GeoIPList) GetEntry() []*GeoIP {
	if m != nil {
		return m.Entry
	}
	return nil
}

type GeoSite struct {
	CountryCode          string    `protobuf:"bytes,1,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Domain               []*Domain `protobuf:"bytes,2,rep,name=domain,proto3" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GeoSite) Reset()         { *m = GeoSite{} }
func (m *GeoSite) String() string { return proto.CompactTextString(m) }
func (*GeoSite) ProtoMessage()    {}
func (*GeoSite) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b1608360690c5fc, []int{4}
}

func (m *GeoSite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeoSite.Unmarshal(m, b)
}
func (m *GeoSite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeoSite.Marshal(b, m, deterministic)
}
func (m *GeoSite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeoSite.Merge(m, src)
}
func (m *GeoSite) XXX_Size() int {
	return xxx_messageInfo_GeoSite.Size(m)
}
func (m *GeoSite) XXX_DiscardUnknown() {
	xxx_messageInfo_GeoSite.DiscardUnknown(m)
}

var xxx_messageInfo_GeoSite proto.InternalMessageInfo

func (m *GeoSite) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func (m *GeoSite) GetDomain() []*Domain {
	if m != nil {
		return m.Domain
	}
	return nil
}

type GeoSiteList struct {
	Entry                []*GeoSite `protobuf:"bytes,1,rep,name=entry,proto3" json:"entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GeoSiteList) Reset()         { *m = GeoSiteList{} }
func (m *GeoSiteList) String() string { return proto.CompactTextString(m) }
func (*GeoSiteList) ProtoMessage()    {}
func (*GeoSiteList) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b1608360690c5fc, []int{5}
}

func (m *GeoSiteList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeoSiteList.Unmarshal(m, b)
}
func (m *GeoSiteList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeoSiteList.Marshal(b, m, deterministic)
}
func (m *GeoSiteList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeoSiteList.Merge(m, src)
}
func (m *GeoSiteList) XXX_Size() int {
	return xxx_messageInfo_GeoSiteList.Size(m)
}
func (m *GeoSiteList) XXX_DiscardUnknown() {
	xxx_messageInfo_GeoSiteList.DiscardUnknown(m)
}

var xxx_messageInfo_GeoSiteList proto.InternalMessageInfo

func (m *GeoSiteList) GetEntry() []*GeoSite {
	if m != nil {
		return m.Entry
	}
	return nil
}

type RoutingRule struct {
	// Types that are valid to be assigned to TargetTag:
	//	*RoutingRule_Tag
	//	*RoutingRule_BalancingTag
	TargetTag isRoutingRule_TargetTag `protobuf_oneof:"target_tag"`
	// List of domains for target domain matching.
	Domain []*Domain `protobuf:"bytes,2,rep,name=domain,proto3" json:"domain,omitempty"`
	// List of CIDRs for target IP address matching.
	// Deprecated. Use geoip below.
	Cidr []*CIDR `protobuf:"bytes,3,rep,name=cidr,proto3" json:"cidr,omitempty"` // Deprecated: Do not use.
	// List of GeoIPs for target IP address matching. If this entry exists, the cidr above will have no effect.
	// GeoIP fields with the same country code are supposed to contain exactly same content. They will be merged during runtime.
	// For customized GeoIPs, please leave country code empty.
	Geoip []*GeoIP `protobuf:"bytes,10,rep,name=geoip,proto3" json:"geoip,omitempty"`
	// A range of port [from, to]. If the destination port is in this range, this rule takes effect.
	// Deprecated. Use port_list.
	PortRange *net.PortRange `protobuf:"bytes,4,opt,name=port_range,json=portRange,proto3" json:"port_range,omitempty"` // Deprecated: Do not use.
	// List of ports.
	PortList *net.PortList `protobuf:"bytes,14,opt,name=port_list,json=portList,proto3" json:"port_list,omitempty"`
	// List of networks. Deprecated. Use networks.
	NetworkList *net.NetworkList `protobuf:"bytes,5,opt,name=network_list,json=networkList,proto3" json:"network_list,omitempty"` // Deprecated: Do not use.
	// List of networks for matching.
	Networks []net.Network `protobuf:"varint,13,rep,packed,name=networks,proto3,enum=v2ray.core.common.net.Network" json:"networks,omitempty"`
	// List of CIDRs for source IP address matching.
	SourceCidr []*CIDR `protobuf:"bytes,6,rep,name=source_cidr,json=sourceCidr,proto3" json:"source_cidr,omitempty"` // Deprecated: Do not use.
	// List of GeoIPs for source IP address matching. If this entry exists, the source_cidr above will have no effect.
	SourceGeoip          []*GeoIP `protobuf:"bytes,11,rep,name=source_geoip,json=sourceGeoip,proto3" json:"source_geoip,omitempty"`
	UserEmail            []string `protobuf:"bytes,7,rep,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	InboundTag           []string `protobuf:"bytes,8,rep,name=inbound_tag,json=inboundTag,proto3" json:"inbound_tag,omitempty"`
	Protocol             []string `protobuf:"bytes,9,rep,name=protocol,proto3" json:"protocol,omitempty"`
	Attributes           string   `protobuf:"bytes,15,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoutingRule) Reset()         { *m = RoutingRule{} }
func (m *RoutingRule) String() string { return proto.CompactTextString(m) }
func (*RoutingRule) ProtoMessage()    {}
func (*RoutingRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b1608360690c5fc, []int{6}
}

func (m *RoutingRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutingRule.Unmarshal(m, b)
}
func (m *RoutingRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutingRule.Marshal(b, m, deterministic)
}
func (m *RoutingRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutingRule.Merge(m, src)
}
func (m *RoutingRule) XXX_Size() int {
	return xxx_messageInfo_RoutingRule.Size(m)
}
func (m *RoutingRule) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutingRule.DiscardUnknown(m)
}

var xxx_messageInfo_RoutingRule proto.InternalMessageInfo

type isRoutingRule_TargetTag interface {
	isRoutingRule_TargetTag()
}

type RoutingRule_Tag struct {
	Tag string `protobuf:"bytes,1,opt,name=tag,proto3,oneof"`
}

type RoutingRule_BalancingTag struct {
	BalancingTag string `protobuf:"bytes,12,opt,name=balancing_tag,json=balancingTag,proto3,oneof"`
}

func (*RoutingRule_Tag) isRoutingRule_TargetTag() {}

func (*RoutingRule_BalancingTag) isRoutingRule_TargetTag() {}

func (m *RoutingRule) GetTargetTag() isRoutingRule_TargetTag {
	if m != nil {
		return m.TargetTag
	}
	return nil
}

func (m *RoutingRule) GetTag() string {
	if x, ok := m.GetTargetTag().(*RoutingRule_Tag); ok {
		return x.Tag
	}
	return ""
}

func (m *RoutingRule) GetBalancingTag() string {
	if x, ok := m.GetTargetTag().(*RoutingRule_BalancingTag); ok {
		return x.BalancingTag
	}
	return ""
}

func (m *RoutingRule) GetDomain() []*Domain {
	if m != nil {
		return m.Domain
	}
	return nil
}

// Deprecated: Do not use.
func (m *RoutingRule) GetCidr() []*CIDR {
	if m != nil {
		return m.Cidr
	}
	return nil
}

func (m *RoutingRule) GetGeoip() []*GeoIP {
	if m != nil {
		return m.Geoip
	}
	return nil
}

// Deprecated: Do not use.
func (m *RoutingRule) GetPortRange() *net.PortRange {
	if m != nil {
		return m.PortRange
	}
	return nil
}

func (m *RoutingRule) GetPortList() *net.PortList {
	if m != nil {
		return m.PortList
	}
	return nil
}

// Deprecated: Do not use.
func (m *RoutingRule) GetNetworkList() *net.NetworkList {
	if m != nil {
		return m.NetworkList
	}
	return nil
}

func (m *RoutingRule) GetNetworks() []net.Network {
	if m != nil {
		return m.Networks
	}
	return nil
}

// Deprecated: Do not use.
func (m *RoutingRule) GetSourceCidr() []*CIDR {
	if m != nil {
		return m.SourceCidr
	}
	return nil
}

func (m *RoutingRule) GetSourceGeoip() []*GeoIP {
	if m != nil {
		return m.SourceGeoip
	}
	return nil
}

func (m *RoutingRule) GetUserEmail() []string {
	if m != nil {
		return m.UserEmail
	}
	return nil
}

func (m *RoutingRule) GetInboundTag() []string {
	if m != nil {
		return m.InboundTag
	}
	return nil
}

func (m *RoutingRule) GetProtocol() []string {
	if m != nil {
		return m.Protocol
	}
	return nil
}

func (m *RoutingRule) GetAttributes() string {
	if m != nil {
		return m.Attributes
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RoutingRule) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RoutingRule_Tag)(nil),
		(*RoutingRule_BalancingTag)(nil),
	}
}

type BalancingOptimalStrategyConfig struct {
	Timeout              uint32   `protobuf:"varint,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Interval             uint32   `protobuf:"varint,2,opt,name=interval,proto3" json:"interval,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Count                uint32   `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	GlobalLock           bool     `protobuf:"varint,5,opt,name=global_lock,json=globalLock,proto3" json:"global_lock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BalancingOptimalStrategyConfig) Reset()         { *m = BalancingOptimalStrategyConfig{} }
func (m *BalancingOptimalStrategyConfig) String() string { return proto.CompactTextString(m) }
func (*BalancingOptimalStrategyConfig) ProtoMessage()    {}
func (*BalancingOptimalStrategyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b1608360690c5fc, []int{7}
}

func (m *BalancingOptimalStrategyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalancingOptimalStrategyConfig.Unmarshal(m, b)
}
func (m *BalancingOptimalStrategyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalancingOptimalStrategyConfig.Marshal(b, m, deterministic)
}
func (m *BalancingOptimalStrategyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalancingOptimalStrategyConfig.Merge(m, src)
}
func (m *BalancingOptimalStrategyConfig) XXX_Size() int {
	return xxx_messageInfo_BalancingOptimalStrategyConfig.Size(m)
}
func (m *BalancingOptimalStrategyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BalancingOptimalStrategyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BalancingOptimalStrategyConfig proto.InternalMessageInfo

func (m *BalancingOptimalStrategyConfig) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *BalancingOptimalStrategyConfig) GetInterval() uint32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *BalancingOptimalStrategyConfig) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *BalancingOptimalStrategyConfig) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *BalancingOptimalStrategyConfig) GetGlobalLock() bool {
	if m != nil {
		return m.GlobalLock
	}
	return false
}

type BalancingRule struct {
	Tag                   string                          `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	OutboundSelector      []string                        `protobuf:"bytes,2,rep,name=outbound_selector,json=outboundSelector,proto3" json:"outbound_selector,omitempty"`
	Strategy              string                          `protobuf:"bytes,3,opt,name=strategy,proto3" json:"strategy,omitempty"`
	OptimalStrategyConfig *BalancingOptimalStrategyConfig `protobuf:"bytes,4,opt,name=optimal_strategy_config,json=optimalStrategyConfig,proto3" json:"optimal_strategy_config,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                        `json:"-"`
	XXX_unrecognized      []byte                          `json:"-"`
	XXX_sizecache         int32                           `json:"-"`
}

func (m *BalancingRule) Reset()         { *m = BalancingRule{} }
func (m *BalancingRule) String() string { return proto.CompactTextString(m) }
func (*BalancingRule) ProtoMessage()    {}
func (*BalancingRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b1608360690c5fc, []int{8}
}

func (m *BalancingRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalancingRule.Unmarshal(m, b)
}
func (m *BalancingRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalancingRule.Marshal(b, m, deterministic)
}
func (m *BalancingRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalancingRule.Merge(m, src)
}
func (m *BalancingRule) XXX_Size() int {
	return xxx_messageInfo_BalancingRule.Size(m)
}
func (m *BalancingRule) XXX_DiscardUnknown() {
	xxx_messageInfo_BalancingRule.DiscardUnknown(m)
}

var xxx_messageInfo_BalancingRule proto.InternalMessageInfo

func (m *BalancingRule) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *BalancingRule) GetOutboundSelector() []string {
	if m != nil {
		return m.OutboundSelector
	}
	return nil
}

func (m *BalancingRule) GetStrategy() string {
	if m != nil {
		return m.Strategy
	}
	return ""
}

func (m *BalancingRule) GetOptimalStrategyConfig() *BalancingOptimalStrategyConfig {
	if m != nil {
		return m.OptimalStrategyConfig
	}
	return nil
}

type Config struct {
	DomainStrategy       Config_DomainStrategy `protobuf:"varint,1,opt,name=domain_strategy,json=domainStrategy,proto3,enum=v2ray.core.app.router.Config_DomainStrategy" json:"domain_strategy,omitempty"`
	Rule                 []*RoutingRule        `protobuf:"bytes,2,rep,name=rule,proto3" json:"rule,omitempty"`
	BalancingRule        []*BalancingRule      `protobuf:"bytes,3,rep,name=balancing_rule,json=balancingRule,proto3" json:"balancing_rule,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b1608360690c5fc, []int{9}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetDomainStrategy() Config_DomainStrategy {
	if m != nil {
		return m.DomainStrategy
	}
	return Config_AsIs
}

func (m *Config) GetRule() []*RoutingRule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *Config) GetBalancingRule() []*BalancingRule {
	if m != nil {
		return m.BalancingRule
	}
	return nil
}

func init() {
	proto.RegisterEnum("v2ray.core.app.router.Domain_Type", Domain_Type_name, Domain_Type_value)
	proto.RegisterEnum("v2ray.core.app.router.Config_DomainStrategy", Config_DomainStrategy_name, Config_DomainStrategy_value)
	proto.RegisterType((*Domain)(nil), "v2ray.core.app.router.Domain")
	proto.RegisterType((*Domain_Attribute)(nil), "v2ray.core.app.router.Domain.Attribute")
	proto.RegisterType((*CIDR)(nil), "v2ray.core.app.router.CIDR")
	proto.RegisterType((*GeoIP)(nil), "v2ray.core.app.router.GeoIP")
	proto.RegisterType((*GeoIPList)(nil), "v2ray.core.app.router.GeoIPList")
	proto.RegisterType((*GeoSite)(nil), "v2ray.core.app.router.GeoSite")
	proto.RegisterType((*GeoSiteList)(nil), "v2ray.core.app.router.GeoSiteList")
	proto.RegisterType((*RoutingRule)(nil), "v2ray.core.app.router.RoutingRule")
	proto.RegisterType((*BalancingOptimalStrategyConfig)(nil), "v2ray.core.app.router.BalancingOptimalStrategyConfig")
	proto.RegisterType((*BalancingRule)(nil), "v2ray.core.app.router.BalancingRule")
	proto.RegisterType((*Config)(nil), "v2ray.core.app.router.Config")
}

func init() {
	proto.RegisterFile("v2ray.com/core/app/router/config.proto", fileDescriptor_6b1608360690c5fc)
}

var fileDescriptor_6b1608360690c5fc = []byte{
	// 1026 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x6d, 0x6f, 0x1b, 0xc5,
	0x13, 0xcf, 0xf9, 0x6c, 0xd7, 0x37, 0x7e, 0xe8, 0xfd, 0x57, 0xff, 0xc2, 0x11, 0x48, 0x62, 0x4e,
	0x85, 0x5a, 0x02, 0x9d, 0x25, 0x97, 0xf2, 0x02, 0x81, 0x4a, 0xe2, 0x94, 0xc4, 0xa2, 0xb4, 0xd1,
	0xa6, 0xed, 0x0b, 0x78, 0x61, 0xad, 0xcf, 0x9b, 0x63, 0x95, 0xf3, 0xee, 0x69, 0x6f, 0x2f, 0xd4,
	0x9f, 0x86, 0xf7, 0x48, 0x7c, 0x06, 0x3e, 0x07, 0x9f, 0x06, 0xb4, 0x0f, 0x76, 0x12, 0x88, 0xd3,
	0x88, 0x77, 0x3b, 0xb3, 0xbf, 0x99, 0xfd, 0xcd, 0xcc, 0xce, 0x0c, 0x7c, 0x7a, 0x31, 0x92, 0x64,
	0x99, 0xa4, 0x62, 0x31, 0x4c, 0x85, 0xa4, 0x43, 0x52, 0x14, 0x43, 0x29, 0x2a, 0x45, 0xe5, 0x30,
	0x15, 0xfc, 0x8c, 0x65, 0x49, 0x21, 0x85, 0x12, 0xe8, 0xc1, 0x0a, 0x27, 0x69, 0x42, 0x8a, 0x22,
	0xb1, 0x98, 0xed, 0x87, 0xff, 0x30, 0x4f, 0xc5, 0x62, 0x21, 0xf8, 0x90, 0x53, 0x35, 0x2c, 0x84,
	0x54, 0xd6, 0x78, 0xfb, 0xd1, 0x66, 0x14, 0xa7, 0xea, 0x17, 0x21, 0xcf, 0x2d, 0x30, 0xfe, 0xa3,
	0x06, 0xcd, 0x43, 0xb1, 0x20, 0x8c, 0xa3, 0x2f, 0xa1, 0xae, 0x96, 0x05, 0x8d, 0xbc, 0xbe, 0x37,
	0xe8, 0x8d, 0xe2, 0xe4, 0xc6, 0xf7, 0x13, 0x0b, 0x4e, 0x5e, 0x2d, 0x0b, 0x8a, 0x0d, 0x1e, 0xfd,
	0x1f, 0x1a, 0x17, 0x24, 0xaf, 0x68, 0x54, 0xeb, 0x7b, 0x83, 0x00, 0x5b, 0x01, 0x3d, 0x83, 0x80,
	0x28, 0x25, 0xd9, 0xac, 0x52, 0x34, 0xf2, 0xfb, 0xfe, 0xa0, 0x3d, 0x7a, 0x74, 0xbb, 0xcb, 0xfd,
	0x15, 0x1c, 0x5f, 0x5a, 0x6e, 0xe7, 0x10, 0xac, 0xf5, 0x28, 0x04, 0xff, 0x9c, 0x2e, 0x0d, 0xc1,
	0x00, 0xeb, 0x23, 0xda, 0x03, 0x98, 0x09, 0x91, 0x4f, 0x2f, 0x09, 0xb4, 0x8e, 0xb7, 0x70, 0xa0,
	0x75, 0x6f, 0x0c, 0x8d, 0x1d, 0x08, 0x18, 0x57, 0xee, 0xde, 0xef, 0x7b, 0x03, 0xff, 0x78, 0x0b,
	0xb7, 0x18, 0x57, 0xe6, 0xfa, 0xa0, 0x0b, 0x6d, 0x1d, 0xc3, 0xdc, 0x02, 0xe2, 0x11, 0xd4, 0x75,
	0x60, 0x28, 0x80, 0xc6, 0x49, 0x4e, 0x18, 0x0f, 0xb7, 0xf4, 0x11, 0xd3, 0x8c, 0xbe, 0x0d, 0x3d,
	0x04, 0xab, 0x54, 0x85, 0x35, 0xd4, 0x82, 0xfa, 0x77, 0x55, 0x9e, 0x87, 0x7e, 0x9c, 0x40, 0x7d,
	0x3c, 0x39, 0xc4, 0xa8, 0x07, 0x35, 0x56, 0x18, 0x6e, 0x1d, 0x5c, 0x63, 0x05, 0x7a, 0x0f, 0x9a,
	0x85, 0xa4, 0x67, 0xec, 0xad, 0xa1, 0xd5, 0xc5, 0x4e, 0x8a, 0x7f, 0x82, 0xc6, 0x11, 0x15, 0x93,
	0x13, 0xf4, 0x31, 0x74, 0x52, 0x51, 0x71, 0x25, 0x97, 0xd3, 0x54, 0xcc, 0xa9, 0x0b, 0xab, 0xed,
	0x74, 0x63, 0x31, 0xa7, 0x68, 0x08, 0xf5, 0x94, 0xcd, 0x65, 0x54, 0x33, 0xf9, 0xfb, 0x70, 0x43,
	0xfe, 0xf4, 0xf3, 0xd8, 0x00, 0xe3, 0xa7, 0x10, 0x18, 0xe7, 0xcf, 0x59, 0xa9, 0xd0, 0x08, 0x1a,
	0x54, 0xbb, 0x8a, 0x3c, 0x63, 0xfe, 0xd1, 0x06, 0x73, 0x63, 0x80, 0x2d, 0x34, 0x4e, 0xe1, 0xde,
	0x11, 0x15, 0xa7, 0x4c, 0xd1, 0xbb, 0xf0, 0x7b, 0x02, 0xcd, 0xb9, 0xc9, 0x88, 0x63, 0xb8, 0x73,
	0x6b, 0x85, 0xb1, 0x03, 0xc7, 0x63, 0x68, 0xbb, 0x47, 0x0c, 0xcf, 0x2f, 0xae, 0xf3, 0xdc, 0xdd,
	0xcc, 0x53, 0x9b, 0xac, 0x98, 0xfe, 0xd5, 0x80, 0x36, 0x16, 0x95, 0x62, 0x3c, 0xc3, 0x55, 0x4e,
	0x11, 0x02, 0x5f, 0x91, 0xcc, 0xb2, 0x3c, 0xde, 0xc2, 0x5a, 0x40, 0x9f, 0x40, 0x77, 0x46, 0x72,
	0xc2, 0x53, 0xc6, 0xb3, 0xa9, 0xbe, 0xed, 0xb8, 0xdb, 0xce, 0x5a, 0xfd, 0x8a, 0x64, 0xff, 0x31,
	0x0c, 0xf4, 0xd8, 0x55, 0xc7, 0x7f, 0x67, 0x75, 0x0e, 0x6a, 0x91, 0x67, 0x2b, 0xa4, 0x8b, 0x92,
	0x51, 0xc1, 0x8a, 0x08, 0xee, 0x52, 0x14, 0x03, 0x45, 0x63, 0x00, 0xdd, 0xdb, 0x53, 0x49, 0x78,
	0x46, 0xa3, 0x7a, 0xdf, 0x1b, 0xb4, 0x47, 0xfd, 0xab, 0x86, 0xb6, 0xbd, 0x13, 0x4e, 0x55, 0x72,
	0x22, 0xa4, 0xc2, 0x1a, 0x67, 0xde, 0x0c, 0x8a, 0x95, 0x88, 0xbe, 0x06, 0x23, 0x4c, 0x73, 0x56,
	0xaa, 0xa8, 0x67, 0x7c, 0xec, 0xdd, 0xe2, 0x43, 0x57, 0x06, 0xb7, 0x0a, 0x77, 0x42, 0x13, 0xe8,
	0xb8, 0xc1, 0x61, 0x1d, 0x34, 0x8c, 0x83, 0x78, 0x83, 0x83, 0x17, 0x16, 0xaa, 0x2d, 0x0d, 0x8d,
	0x36, 0xbf, 0x54, 0xa0, 0xaf, 0xa0, 0xe5, 0xc4, 0x32, 0xea, 0xf6, 0xfd, 0x41, 0xef, 0x7a, 0xc5,
	0xff, 0xed, 0x06, 0xaf, 0xf1, 0xe8, 0x5b, 0x68, 0x97, 0xa2, 0x92, 0x29, 0x9d, 0x9a, 0xcc, 0x37,
	0xef, 0x96, 0x79, 0xb0, 0x36, 0x63, 0x9d, 0xff, 0xa7, 0xd0, 0x71, 0x1e, 0x6c, 0x19, 0xda, 0x77,
	0x28, 0x83, 0x7b, 0xf3, 0xc8, 0x14, 0x63, 0x07, 0xa0, 0x2a, 0xa9, 0x9c, 0xd2, 0x05, 0x61, 0x79,
	0x74, 0xaf, 0xef, 0x0f, 0x02, 0x1c, 0x68, 0xcd, 0x33, 0xad, 0x40, 0x7b, 0xd0, 0x66, 0x7c, 0x26,
	0x2a, 0x3e, 0x37, 0x1f, 0xae, 0x65, 0xee, 0xc1, 0xa9, 0xf4, 0x67, 0xdb, 0x86, 0x96, 0x19, 0xbd,
	0xa9, 0xc8, 0xa3, 0xc0, 0xdc, 0xae, 0x65, 0xb4, 0x0b, 0xb0, 0x1e, 0x7d, 0x65, 0x74, 0xdf, 0x34,
	0xdc, 0x15, 0xcd, 0x41, 0x07, 0x40, 0x11, 0x99, 0x51, 0xa5, 0x7d, 0xc7, 0xbf, 0x7a, 0xb0, 0x7b,
	0xb0, 0xfa, 0xc7, 0x2f, 0x0b, 0xc5, 0x16, 0x24, 0x3f, 0x55, 0x92, 0x28, 0x9a, 0x2d, 0xc7, 0x66,
	0x95, 0xa0, 0x08, 0xee, 0x29, 0xb6, 0xa0, 0xa2, 0x52, 0xa6, 0x31, 0xba, 0x78, 0x25, 0x6a, 0x1a,
	0x8c, 0x2b, 0x2a, 0x2f, 0x48, 0xee, 0x06, 0xd4, 0x5a, 0xd6, 0x73, 0xb6, 0x92, 0xb9, 0x19, 0x97,
	0x01, 0xd6, 0x47, 0x3d, 0xe3, 0x4d, 0xdf, 0x9b, 0xcf, 0xd7, 0xc5, 0x56, 0xd0, 0xb1, 0x66, 0xb9,
	0x98, 0x91, 0x7c, 0x9a, 0x8b, 0xf4, 0xdc, 0xfc, 0x89, 0x16, 0x06, 0xab, 0x7a, 0x2e, 0xd2, 0xf3,
	0xf8, 0x4f, 0x0f, 0xba, 0x6b, 0x86, 0xa6, 0x4b, 0xc3, 0x2b, 0x5d, 0x6a, 0x7b, 0xf4, 0x33, 0xf8,
	0x9f, 0xa8, 0x94, 0xcd, 0x58, 0x49, 0x73, 0x9a, 0x2a, 0x61, 0x07, 0x5e, 0x80, 0xc3, 0xd5, 0xc5,
	0xa9, 0xd3, 0x6b, 0xd6, 0xa5, 0x8b, 0xd0, 0xd1, 0x5b, 0xcb, 0x68, 0x01, 0xef, 0x0b, 0x9b, 0x84,
	0xe9, 0x4a, 0x37, 0xb5, 0x1b, 0xd5, 0xb5, 0xcc, 0x93, 0x0d, 0x45, 0xbe, 0x3d, 0x87, 0xf8, 0x81,
	0xb8, 0x49, 0x1d, 0xff, 0x5e, 0x83, 0xa6, 0xcb, 0xf2, 0x6b, 0xb8, 0x6f, 0x47, 0xc2, 0xfa, 0x61,
	0xb7, 0x44, 0x3f, 0xdf, 0xf4, 0x33, 0xed, 0xa2, 0xb7, 0xf3, 0x64, 0xe5, 0x17, 0xf7, 0xe6, 0xd7,
	0x64, 0xbd, 0x90, 0x65, 0x95, 0x53, 0x37, 0x94, 0x36, 0x2d, 0xe4, 0x2b, 0x33, 0x10, 0x1b, 0x3c,
	0xfa, 0x1e, 0x7a, 0x97, 0x53, 0xcf, 0x78, 0xb0, 0x13, 0xea, 0xe1, 0xbb, 0xe2, 0x37, 0x3e, 0x2e,
	0x27, 0xa6, 0x16, 0xe3, 0x23, 0xe8, 0x5d, 0xa7, 0xa9, 0x57, 0xdf, 0x7e, 0x39, 0x29, 0xed, 0x6e,
	0x7c, 0x5d, 0xd2, 0x49, 0x11, 0x7a, 0x28, 0x84, 0xce, 0xa4, 0x98, 0x9c, 0xbd, 0x10, 0xfc, 0x07,
	0xa2, 0xd2, 0x9f, 0xc3, 0x1a, 0xea, 0x01, 0x4c, 0x8a, 0x97, 0xfc, 0x90, 0x2e, 0x08, 0x9f, 0x87,
	0xfe, 0xc1, 0x37, 0xf0, 0x41, 0x2a, 0x16, 0x37, 0x53, 0x38, 0xf1, 0x7e, 0x6c, 0xda, 0xd3, 0x6f,
	0xb5, 0x07, 0x6f, 0x46, 0x98, 0x2c, 0x93, 0xb1, 0x46, 0xec, 0x17, 0x85, 0x89, 0x8f, 0xca, 0x59,
	0xd3, 0x34, 0xc9, 0xe3, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x58, 0x98, 0x85, 0x3f, 0x09,
	0x00, 0x00,
}
