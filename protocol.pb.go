// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol.proto

package udetect

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ConnType int32

const (
	ConnType_Unknown  ConnType = 0
	ConnType_Ethernet ConnType = 1
	ConnType_WIFI     ConnType = 2
	ConnType_Cell     ConnType = 3
	ConnType_Cell2G   ConnType = 4
	ConnType_Cell3G   ConnType = 5
	ConnType_Cell4G   ConnType = 6
	ConnType_Cell5G   ConnType = 7
)

var ConnType_name = map[int32]string{
	0: "Unknown",
	1: "Ethernet",
	2: "WIFI",
	3: "Cell",
	4: "Cell2G",
	5: "Cell3G",
	6: "Cell4G",
	7: "Cell5G",
}

var ConnType_value = map[string]int32{
	"Unknown":  0,
	"Ethernet": 1,
	"WIFI":     2,
	"Cell":     3,
	"Cell2G":   4,
	"Cell3G":   5,
	"Cell4G":   6,
	"Cell5G":   7,
}

func (x ConnType) String() string {
	return proto.EnumName(ConnType_name, int32(x))
}

func (ConnType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{0}
}

type DeviceType int32

const (
	DeviceType_Undefined DeviceType = 0
	DeviceType_Mobile    DeviceType = 1
	DeviceType_PC        DeviceType = 2
	DeviceType_TV        DeviceType = 3
	DeviceType_Phone     DeviceType = 4
	DeviceType_Tablet    DeviceType = 5
	DeviceType_Connected DeviceType = 6
	DeviceType_SetTopBox DeviceType = 7
	DeviceType_Watch     DeviceType = 8
	DeviceType_Glasses   DeviceType = 9
)

var DeviceType_name = map[int32]string{
	0: "Undefined",
	1: "Mobile",
	2: "PC",
	3: "TV",
	4: "Phone",
	5: "Tablet",
	6: "Connected",
	7: "SetTopBox",
	8: "Watch",
	9: "Glasses",
}

var DeviceType_value = map[string]int32{
	"Undefined": 0,
	"Mobile":    1,
	"PC":        2,
	"TV":        3,
	"Phone":     4,
	"Tablet":    5,
	"Connected": 6,
	"SetTopBox": 7,
	"Watch":     8,
	"Glasses":   9,
}

func (x DeviceType) String() string {
	return proto.EnumName(DeviceType_name, int32(x))
}

func (DeviceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{1}
}

type UUID struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UUID) Reset()         { *m = UUID{} }
func (m *UUID) String() string { return proto.CompactTextString(m) }
func (*UUID) ProtoMessage()    {}
func (*UUID) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{0}
}

func (m *UUID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UUID.Unmarshal(m, b)
}
func (m *UUID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UUID.Marshal(b, m, deterministic)
}
func (m *UUID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UUID.Merge(m, src)
}
func (m *UUID) XXX_Size() int {
	return xxx_messageInfo_UUID.Size(m)
}
func (m *UUID) XXX_DiscardUnknown() {
	xxx_messageInfo_UUID.DiscardUnknown(m)
}

var xxx_messageInfo_UUID proto.InternalMessageInfo

func (m *UUID) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Carrier struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Carrier) Reset()         { *m = Carrier{} }
func (m *Carrier) String() string { return proto.CompactTextString(m) }
func (*Carrier) ProtoMessage()    {}
func (*Carrier) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{1}
}

func (m *Carrier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Carrier.Unmarshal(m, b)
}
func (m *Carrier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Carrier.Marshal(b, m, deterministic)
}
func (m *Carrier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Carrier.Merge(m, src)
}
func (m *Carrier) XXX_Size() int {
	return xxx_messageInfo_Carrier.Size(m)
}
func (m *Carrier) XXX_DiscardUnknown() {
	xxx_messageInfo_Carrier.DiscardUnknown(m)
}

var xxx_messageInfo_Carrier proto.InternalMessageInfo

func (m *Carrier) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Carrier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Carrier) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type User struct {
	// Unical User Identificator inside the traking system and provided to the user device
	Uuid *UUID `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Session Identificator inside the traking system and provided to the user device
	Sessid *UUID `protobuf:"bytes,2,opt,name=sessid,proto3" json:"sessid,omitempty"`
	// Minimal possible age of the user
	AgeStart int32 `protobuf:"varint,3,opt,name=age_start,json=ageStart,proto3" json:"age_start,omitempty"`
	// Maximal possible age of the user
	AgeEnd int32 `protobuf:"varint,4,opt,name=age_end,json=ageEnd,proto3" json:"age_end,omitempty"`
	// main keywords interested by user
	Keywords             []string `protobuf:"bytes,5,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{2}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUuid() *UUID {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *User) GetSessid() *UUID {
	if m != nil {
		return m.Sessid
	}
	return nil
}

func (m *User) GetAgeStart() int32 {
	if m != nil {
		return m.AgeStart
	}
	return 0
}

func (m *User) GetAgeEnd() int32 {
	if m != nil {
		return m.AgeEnd
	}
	return 0
}

func (m *User) GetKeywords() []string {
	if m != nil {
		return m.Keywords
	}
	return nil
}

type OS struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the OS; like MacOS, Linux, Unix, etc.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Indicates the version or subversion of the software platform.
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OS) Reset()         { *m = OS{} }
func (m *OS) String() string { return proto.CompactTextString(m) }
func (*OS) ProtoMessage()    {}
func (*OS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{3}
}

func (m *OS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OS.Unmarshal(m, b)
}
func (m *OS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OS.Marshal(b, m, deterministic)
}
func (m *OS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OS.Merge(m, src)
}
func (m *OS) XXX_Size() int {
	return xxx_messageInfo_OS.Size(m)
}
func (m *OS) XXX_DiscardUnknown() {
	xxx_messageInfo_OS.DiscardUnknown(m)
}

var xxx_messageInfo_OS proto.InternalMessageInfo

func (m *OS) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OS) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type Browser struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the browser; like Safari, Chrome, Firefox, etc.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Indicates the version or subversion of the browser.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// Is browser UA was detected as robot
	IsRobot              int32    `protobuf:"varint,4,opt,name=is_robot,json=isRobot,proto3" json:"is_robot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Browser) Reset()         { *m = Browser{} }
func (m *Browser) String() string { return proto.CompactTextString(m) }
func (*Browser) ProtoMessage()    {}
func (*Browser) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{4}
}

func (m *Browser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Browser.Unmarshal(m, b)
}
func (m *Browser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Browser.Marshal(b, m, deterministic)
}
func (m *Browser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Browser.Merge(m, src)
}
func (m *Browser) XXX_Size() int {
	return xxx_messageInfo_Browser.Size(m)
}
func (m *Browser) XXX_DiscardUnknown() {
	xxx_messageInfo_Browser.DiscardUnknown(m)
}

var xxx_messageInfo_Browser proto.InternalMessageInfo

func (m *Browser) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Browser) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Browser) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Browser) GetIsRobot() int32 {
	if m != nil {
		return m.IsRobot
	}
	return 0
}

type Device struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Indicates the name of the company that manufactures the device
	// or primarily sells it, e.g. Samsung.
	Make string `protobuf:"bytes,2,opt,name=make,proto3" json:"make,omitempty"`
	// Indicates the model name or number used primarily by the hardware
	// vendor to identify the device, e.g.SM-T805S. When a model identifier
	// is not available the HardwareName will be used.
	Model string `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	// Indicates the name and version of the operating system the device is using.
	Os *OS `protobuf:"bytes,4,opt,name=os,proto3" json:"os,omitempty"`
	// Browser information model
	Browser *Browser `protobuf:"bytes,5,opt,name=browser,proto3" json:"browser,omitempty"`
	// Connection type of the device
	Connectiontype ConnType `protobuf:"varint,6,opt,name=connectiontype,proto3,enum=udetect.ConnType" json:"connectiontype,omitempty"`
	// Indicates the type of the device based on values set in other properties,
	// such as IsMobile, IsTablet, IsSmartphone, IsSmallScreen etc.
	Devicetype DeviceType `protobuf:"varint,7,opt,name=devicetype,proto3,enum=udetect.DeviceType" json:"devicetype,omitempty"`
	// Physical height of the screen in pixels.
	// Indicates the height of the device's screen in pixels.
	// This property is not applicable for a device that does not have a screen.
	// For devices such as tablets or TV which are predominantly used in landscape mode,
	// the pixel height will be the smaller value compared to the pixel width.
	Height int32 `protobuf:"varint,8,opt,name=height,proto3" json:"height,omitempty"`
	// Physical width of the screen in pixels
	// Indicates the width of the device's screen in pixels.
	// This property is not applicable for a device that does not have a screen.
	// For devices such as tablets or TV which are predominantly used in landscape mode,
	// the pixel width will be the larger value compared to the pixel height.
	Width int32 `protobuf:"varint,9,opt,name=width,proto3" json:"width,omitempty"`
	// Screen size as pixels per linear inch.
	// ((ScreenPixelsWidth / ScreenInchesWidth) + (ScreenPixelsHeight / ScreenInchesHeight)) / 2
	Ppi int32 `protobuf:"varint,10,opt,name=ppi,proto3" json:"ppi,omitempty"`
	// The ratio of physical pixels to device independent pixels.
	PxRatio float32 `protobuf:"fixed32,11,opt,name=px_ratio,json=pxRatio,proto3" json:"px_ratio,omitempty"`
	// Hardware version of the device (e.g., "5S" for iPhone 5S).
	Hwver                string   `protobuf:"bytes,12,opt,name=hwver,proto3" json:"hwver,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{5}
}

func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Device) GetMake() string {
	if m != nil {
		return m.Make
	}
	return ""
}

func (m *Device) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *Device) GetOs() *OS {
	if m != nil {
		return m.Os
	}
	return nil
}

func (m *Device) GetBrowser() *Browser {
	if m != nil {
		return m.Browser
	}
	return nil
}

func (m *Device) GetConnectiontype() ConnType {
	if m != nil {
		return m.Connectiontype
	}
	return ConnType_Unknown
}

func (m *Device) GetDevicetype() DeviceType {
	if m != nil {
		return m.Devicetype
	}
	return DeviceType_Undefined
}

func (m *Device) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Device) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Device) GetPpi() int32 {
	if m != nil {
		return m.Ppi
	}
	return 0
}

func (m *Device) GetPxRatio() float32 {
	if m != nil {
		return m.PxRatio
	}
	return 0
}

func (m *Device) GetHwver() string {
	if m != nil {
		return m.Hwver
	}
	return ""
}

type GeoLocation struct {
	// Internal ID
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// IPv4/6
	Ip string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	// Carrier or ISP derived from the IP address
	Carrier *Carrier `protobuf:"bytes,3,opt,name=carrier,proto3" json:"carrier,omitempty"`
	// Latitude from -90 to 90
	Lat float32 `protobuf:"fixed32,4,opt,name=lat,proto3" json:"lat,omitempty"`
	// Longitude from -180 to 180
	Lon float32 `protobuf:"fixed32,5,opt,name=lon,proto3" json:"lon,omitempty"`
	// Country using ISO 3166-1 Alpha 2
	Country string `protobuf:"bytes,6,opt,name=country,proto3" json:"country,omitempty"`
	// Region using ISO 3166-2
	Region string `protobuf:"bytes,7,opt,name=region,proto3" json:"region,omitempty"`
	// Region of a country using FIPS 10-4
	RegionFIPS104 string `protobuf:"bytes,8,opt,name=regionFIPS104,proto3" json:"regionFIPS104,omitempty"`
	Metro         string `protobuf:"bytes,9,opt,name=metro,proto3" json:"metro,omitempty"`
	City          string `protobuf:"bytes,10,opt,name=city,proto3" json:"city,omitempty"`
	Zip           string `protobuf:"bytes,11,opt,name=zip,proto3" json:"zip,omitempty"`
	// Local time as the number +/- of minutes from UTC
	UtcOffset            int32    `protobuf:"varint,12,opt,name=utc_offset,json=utcOffset,proto3" json:"utc_offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeoLocation) Reset()         { *m = GeoLocation{} }
func (m *GeoLocation) String() string { return proto.CompactTextString(m) }
func (*GeoLocation) ProtoMessage()    {}
func (*GeoLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{6}
}

func (m *GeoLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeoLocation.Unmarshal(m, b)
}
func (m *GeoLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeoLocation.Marshal(b, m, deterministic)
}
func (m *GeoLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeoLocation.Merge(m, src)
}
func (m *GeoLocation) XXX_Size() int {
	return xxx_messageInfo_GeoLocation.Size(m)
}
func (m *GeoLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_GeoLocation.DiscardUnknown(m)
}

var xxx_messageInfo_GeoLocation proto.InternalMessageInfo

func (m *GeoLocation) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GeoLocation) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *GeoLocation) GetCarrier() *Carrier {
	if m != nil {
		return m.Carrier
	}
	return nil
}

func (m *GeoLocation) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *GeoLocation) GetLon() float32 {
	if m != nil {
		return m.Lon
	}
	return 0
}

func (m *GeoLocation) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *GeoLocation) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *GeoLocation) GetRegionFIPS104() string {
	if m != nil {
		return m.RegionFIPS104
	}
	return ""
}

func (m *GeoLocation) GetMetro() string {
	if m != nil {
		return m.Metro
	}
	return ""
}

func (m *GeoLocation) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *GeoLocation) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *GeoLocation) GetUtcOffset() int32 {
	if m != nil {
		return m.UtcOffset
	}
	return 0
}

type Request struct {
	Udid                 string   `protobuf:"bytes,1,opt,name=udid,proto3" json:"udid,omitempty"`
	Uid                  *UUID    `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Sessid               *UUID    `protobuf:"bytes,3,opt,name=sessid,proto3" json:"sessid,omitempty"`
	Ip                   string   `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
	Ua                   string   `protobuf:"bytes,5,opt,name=ua,proto3" json:"ua,omitempty"`
	Url                  string   `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{7}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUdid() string {
	if m != nil {
		return m.Udid
	}
	return ""
}

func (m *Request) GetUid() *UUID {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *Request) GetSessid() *UUID {
	if m != nil {
		return m.Sessid
	}
	return nil
}

func (m *Request) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Request) GetUa() string {
	if m != nil {
		return m.Ua
	}
	return ""
}

func (m *Request) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Response struct {
	// User person information which might be detected
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// Device information including browser and OS
	Device *Device `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
	// Location of the device assumed to be the user’s current location defined by a Geo object
	Geo                  *GeoLocation `protobuf:"bytes,3,opt,name=geo,proto3" json:"geo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{8}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *Response) GetGeo() *GeoLocation {
	if m != nil {
		return m.Geo
	}
	return nil
}

func init() {
	proto.RegisterEnum("udetect.ConnType", ConnType_name, ConnType_value)
	proto.RegisterEnum("udetect.DeviceType", DeviceType_name, DeviceType_value)
	proto.RegisterType((*UUID)(nil), "udetect.UUID")
	proto.RegisterType((*Carrier)(nil), "udetect.Carrier")
	proto.RegisterType((*User)(nil), "udetect.User")
	proto.RegisterType((*OS)(nil), "udetect.OS")
	proto.RegisterType((*Browser)(nil), "udetect.Browser")
	proto.RegisterType((*Device)(nil), "udetect.Device")
	proto.RegisterType((*GeoLocation)(nil), "udetect.GeoLocation")
	proto.RegisterType((*Request)(nil), "udetect.Request")
	proto.RegisterType((*Response)(nil), "udetect.Response")
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor_2bc2336598a3f7e0) }

var fileDescriptor_2bc2336598a3f7e0 = []byte{
	// 944 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x72, 0xdc, 0x44,
	0x10, 0x8e, 0xb4, 0x5a, 0xfd, 0xf4, 0xda, 0xce, 0x64, 0x08, 0x20, 0x9c, 0x50, 0x18, 0x15, 0x3f,
	0x2e, 0x1f, 0x6c, 0x62, 0x87, 0x03, 0xdc, 0xb0, 0x9d, 0x6c, 0xd9, 0x05, 0x65, 0xd7, 0xd8, 0x26,
	0x47, 0x97, 0x56, 0x6a, 0xef, 0xaa, 0x2c, 0xcf, 0x08, 0xcd, 0xc8, 0xeb, 0xe5, 0xc0, 0x81, 0xe2,
	0x0d, 0x38, 0xf0, 0x1a, 0x9c, 0x79, 0x0d, 0x5e, 0x81, 0x23, 0x0f, 0x41, 0xf5, 0x48, 0xbb, 0x71,
	0x62, 0x92, 0xca, 0x69, 0xbf, 0xaf, 0xbb, 0xa7, 0xa7, 0xfb, 0x53, 0xf7, 0x2c, 0xac, 0x54, 0xb5,
	0x32, 0x2a, 0x53, 0xe5, 0xa6, 0x05, 0x3c, 0x68, 0x72, 0x34, 0x98, 0x99, 0xd5, 0xc7, 0x63, 0xa5,
	0xc6, 0x25, 0x6e, 0xa5, 0x55, 0xb1, 0x95, 0x4a, 0xa9, 0x4c, 0x6a, 0x0a, 0x25, 0x75, 0x1b, 0x96,
	0x3c, 0x06, 0xef, 0xec, 0xec, 0x60, 0x9f, 0x3f, 0x84, 0xfe, 0x75, 0x5a, 0x36, 0x18, 0x3b, 0x6b,
	0xce, 0xfa, 0x92, 0x68, 0x49, 0xf2, 0x1d, 0x04, 0x7b, 0x69, 0x5d, 0x17, 0x58, 0xf3, 0x15, 0x70,
	0x8b, 0xdc, 0x7a, 0x3d, 0xe1, 0x16, 0x39, 0xe7, 0xe0, 0xc9, 0xf4, 0x0a, 0x63, 0x77, 0xcd, 0x59,
	0x8f, 0x84, 0xc5, 0x64, 0xcb, 0x54, 0x8e, 0x71, 0xaf, 0xb5, 0x11, 0x4e, 0xfe, 0x72, 0xc0, 0x3b,
	0xd3, 0x58, 0xf3, 0x4f, 0xc1, 0x6b, 0x9a, 0x2e, 0xc5, 0x60, 0x7b, 0x79, 0xb3, 0xab, 0x6f, 0x93,
	0xae, 0x17, 0xd6, 0xc5, 0x3f, 0x07, 0x5f, 0xa3, 0xd6, 0x45, 0x6e, 0xb3, 0xde, 0x09, 0xea, 0x9c,
	0xfc, 0x11, 0x44, 0xe9, 0x18, 0xcf, 0xb5, 0x49, 0x6b, 0x63, 0xef, 0xea, 0x8b, 0x30, 0x1d, 0xe3,
	0x09, 0x71, 0xfe, 0x21, 0x04, 0xe4, 0x44, 0x99, 0xc7, 0x9e, 0x75, 0xf9, 0xe9, 0x18, 0x9f, 0xc9,
	0x9c, 0xaf, 0x42, 0x78, 0x89, 0xb3, 0xa9, 0xaa, 0x73, 0x1d, 0xf7, 0xd7, 0x7a, 0xeb, 0x91, 0x58,
	0xf0, 0x43, 0x2f, 0xf4, 0x59, 0x70, 0xe8, 0x85, 0x01, 0x0b, 0x0f, 0xbd, 0x30, 0x64, 0xd1, 0xa1,
	0x17, 0x46, 0x0c, 0x92, 0x5d, 0x70, 0x8f, 0x4e, 0xde, 0xa9, 0xf5, 0x18, 0x82, 0x6b, 0xac, 0x75,
	0xa1, 0x64, 0xd7, 0xfd, 0x9c, 0x26, 0x23, 0x08, 0x76, 0x6b, 0x35, 0xd5, 0xef, 0xa8, 0xe1, 0x1b,
	0x13, 0xf1, 0x8f, 0x20, 0x2c, 0xf4, 0x79, 0xad, 0x46, 0xca, 0x74, 0xad, 0x05, 0x85, 0x16, 0x44,
	0x93, 0x7f, 0x5d, 0xf0, 0xf7, 0xf1, 0xba, 0xc8, 0xf0, 0xff, 0xee, 0xb8, 0x4a, 0x2f, 0x17, 0x77,
	0x10, 0xa6, 0x8f, 0x7d, 0xa5, 0x72, 0x2c, 0xbb, 0x1b, 0x5a, 0xc2, 0x1f, 0x81, 0xab, 0xb4, 0xcd,
	0x3c, 0xd8, 0x1e, 0x2c, 0x94, 0x3f, 0x3a, 0x11, 0xae, 0xd2, 0x7c, 0x03, 0x82, 0x51, 0xdb, 0x45,
	0xdc, 0xb7, 0x11, 0x6c, 0x11, 0xd1, 0x75, 0x27, 0xe6, 0x01, 0xfc, 0x1b, 0x58, 0xc9, 0x94, 0x94,
	0x98, 0xd1, 0xa0, 0x99, 0x59, 0x85, 0xb1, 0xbf, 0xe6, 0xac, 0xaf, 0x6c, 0x3f, 0x58, 0x1c, 0xd9,
	0x53, 0x52, 0x9e, 0xce, 0x2a, 0x14, 0xaf, 0x05, 0xf2, 0x1d, 0x80, 0xdc, 0xf6, 0x61, 0x8f, 0x05,
	0xf6, 0xd8, 0x7b, 0x8b, 0x63, 0x6d, 0x8b, 0xf6, 0xe0, 0xad, 0x30, 0xfe, 0x01, 0xf8, 0x13, 0x2c,
	0xc6, 0x13, 0x13, 0x87, 0xed, 0x17, 0x6f, 0x19, 0xb5, 0x39, 0x2d, 0x72, 0x33, 0x89, 0x23, 0x6b,
	0x6e, 0x09, 0x67, 0xd0, 0xab, 0xaa, 0x22, 0x06, 0x6b, 0x23, 0x48, 0xc2, 0x56, 0x37, 0xe7, 0x35,
	0xed, 0x45, 0x3c, 0x58, 0x73, 0xd6, 0x5d, 0x11, 0x54, 0x37, 0x82, 0x28, 0xa5, 0x98, 0x4c, 0xaf,
	0xb1, 0x8e, 0x97, 0x5a, 0xa5, 0x2c, 0x49, 0xfe, 0x74, 0x61, 0x30, 0x44, 0xf5, 0xbd, 0xca, 0xec,
	0x2e, 0xdd, 0xd1, 0x9c, 0x78, 0xd5, 0x29, 0xee, 0x16, 0x15, 0x89, 0x97, 0xb5, 0x6b, 0x64, 0x15,
	0xbf, 0x2d, 0x5e, 0xb7, 0x5e, 0x62, 0x1e, 0x40, 0xe5, 0x95, 0x69, 0xfb, 0x81, 0x5d, 0x41, 0xd0,
	0x5a, 0x94, 0xb4, 0xb2, 0x93, 0x45, 0x49, 0x9a, 0x91, 0x4c, 0x35, 0xd2, 0xd4, 0x33, 0xab, 0x6c,
	0x24, 0xe6, 0x94, 0xa4, 0xa8, 0x71, 0x4c, 0xc3, 0x13, 0x58, 0x47, 0xc7, 0xf8, 0x67, 0xb0, 0xdc,
	0xa2, 0xe7, 0x07, 0xc7, 0x27, 0x4f, 0xbe, 0x7a, 0x6a, 0x95, 0x8a, 0xc4, 0xab, 0x46, 0x3b, 0x17,
	0x68, 0x6a, 0x65, 0x05, 0xa3, 0xb9, 0x20, 0x62, 0xb7, 0xba, 0x30, 0x33, 0xab, 0x18, 0x6d, 0x75,
	0x61, 0x66, 0x54, 0xd3, 0xcf, 0x45, 0x65, 0xd5, 0x8a, 0x04, 0x41, 0xfe, 0x31, 0x40, 0x63, 0xb2,
	0x73, 0x75, 0x71, 0xa1, 0xd1, 0x58, 0xb9, 0xfa, 0x22, 0x6a, 0x4c, 0x76, 0x64, 0x0d, 0xc9, 0x1f,
	0x0e, 0x04, 0x02, 0x7f, 0x6a, 0x50, 0x1b, 0x4a, 0xd8, 0xe4, 0x9d, 0x60, 0x91, 0xb0, 0x98, 0x7f,
	0x02, 0xbd, 0xe6, 0x4d, 0x7b, 0xdf, 0x7b, 0xf5, 0x6d, 0xe8, 0xbd, 0xed, 0x6d, 0x68, 0xa5, 0xf7,
	0x16, 0xd2, 0xaf, 0x80, 0xdb, 0xa4, 0x56, 0xbb, 0x48, 0xb8, 0x4d, 0x4a, 0x85, 0x37, 0x75, 0xd9,
	0xc9, 0x46, 0x30, 0xf9, 0x05, 0x42, 0x81, 0xba, 0x52, 0x52, 0xa3, 0x7d, 0xa3, 0x68, 0xc4, 0xef,
	0xbc, 0x51, 0x34, 0xdf, 0xd6, 0xc5, 0xbf, 0x04, 0xbf, 0x1d, 0xbd, 0xae, 0xd6, 0xfb, 0xaf, 0x4d,
	0xa7, 0xe8, 0xdc, 0xfc, 0x0b, 0xe8, 0x8d, 0x51, 0x75, 0xd5, 0x3e, 0x5c, 0x44, 0xdd, 0x9a, 0x1b,
	0x41, 0x01, 0x1b, 0x05, 0x84, 0xf3, 0x75, 0xe0, 0x03, 0x08, 0xce, 0xe4, 0xa5, 0x54, 0x53, 0xc9,
	0xee, 0xf1, 0x25, 0x08, 0x9f, 0x99, 0x09, 0xd6, 0x12, 0x0d, 0x73, 0x78, 0x08, 0xde, 0x8b, 0x83,
	0xe7, 0x07, 0xcc, 0x25, 0xb4, 0x87, 0x65, 0xc9, 0x7a, 0x1c, 0xc0, 0x27, 0xb4, 0x3d, 0x64, 0xde,
	0x1c, 0xef, 0x0c, 0x59, 0x7f, 0x8e, 0x9f, 0x0e, 0x99, 0x3f, 0xc7, 0x5f, 0x0f, 0x59, 0xb0, 0xf1,
	0x9b, 0x03, 0xf0, 0x72, 0x87, 0xf8, 0x32, 0x44, 0x67, 0x32, 0xc7, 0x8b, 0x42, 0x62, 0xce, 0xee,
	0x51, 0xe4, 0x0f, 0x6a, 0x54, 0x94, 0xc8, 0x1c, 0xee, 0x83, 0x7b, 0xbc, 0xc7, 0x5c, 0xfa, 0x3d,
	0xfd, 0x91, 0xf5, 0x78, 0x04, 0xfd, 0xe3, 0x89, 0x92, 0xd8, 0x5e, 0x74, 0x9a, 0x8e, 0x4a, 0x34,
	0xac, 0x4f, 0x19, 0xf6, 0xda, 0x05, 0xc6, 0x9c, 0xf9, 0x44, 0x4f, 0xd0, 0x9c, 0xaa, 0x6a, 0x57,
	0xdd, 0xb0, 0x80, 0x0e, 0xbd, 0x48, 0x4d, 0x36, 0x61, 0x21, 0x35, 0x36, 0x2c, 0x53, 0xad, 0x51,
	0xb3, 0x68, 0xfb, 0x18, 0xc2, 0x7d, 0x2b, 0x86, 0xaa, 0xf9, 0x3e, 0x3d, 0x5c, 0x84, 0xf9, 0xcb,
	0x9d, 0xe8, 0xe6, 0x64, 0xf5, 0xc1, 0x2d, 0x4b, 0xfb, 0x81, 0x92, 0xf7, 0x7f, 0xfd, 0xfb, 0x9f,
	0xdf, 0xdd, 0xfb, 0x09, 0x6c, 0x5d, 0x3f, 0xd9, 0x6a, 0x9d, 0xdf, 0x3a, 0x1b, 0x23, 0xdf, 0xfe,
	0x99, 0xed, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x7c, 0x85, 0x3c, 0x05, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DetectorClient is the client API for Detector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DetectorClient interface {
	Detect(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type detectorClient struct {
	cc *grpc.ClientConn
}

func NewDetectorClient(cc *grpc.ClientConn) DetectorClient {
	return &detectorClient{cc}
}

func (c *detectorClient) Detect(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/udetect.Detector/Detect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DetectorServer is the server API for Detector service.
type DetectorServer interface {
	Detect(context.Context, *Request) (*Response, error)
}

// UnimplementedDetectorServer can be embedded to have forward compatible implementations.
type UnimplementedDetectorServer struct {
}

func (*UnimplementedDetectorServer) Detect(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detect not implemented")
}

func RegisterDetectorServer(s *grpc.Server, srv DetectorServer) {
	s.RegisterService(&_Detector_serviceDesc, srv)
}

func _Detector_Detect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetectorServer).Detect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/udetect.Detector/Detect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetectorServer).Detect(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Detector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "udetect.Detector",
	HandlerType: (*DetectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Detect",
			Handler:    _Detector_Detect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol.proto",
}
