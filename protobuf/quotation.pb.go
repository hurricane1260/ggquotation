// Code generated by protoc-gen-go. DO NOT EDIT.
// source: quotation.proto

package protobuf

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

// MDGW发布快照行情消息定义
type Head3Xxx11 struct {
	MsgType              int32    `protobuf:"varint,1,opt,name=MsgType,proto3" json:"MsgType,omitempty"`
	OrigTime             int64    `protobuf:"varint,2,opt,name=OrigTime,proto3" json:"OrigTime,omitempty"`
	ChannelNo            int32    `protobuf:"varint,3,opt,name=ChannelNo,proto3" json:"ChannelNo,omitempty"`
	MDStreamID           string   `protobuf:"bytes,4,opt,name=MDStreamID,proto3" json:"MDStreamID,omitempty"`
	SecurityID           string   `protobuf:"bytes,5,opt,name=SecurityID,proto3" json:"SecurityID,omitempty"`
	SecurityIDSource     string   `protobuf:"bytes,6,opt,name=SecurityIDSource,proto3" json:"SecurityIDSource,omitempty"`
	TradingPhaseCode     []byte   `protobuf:"bytes,7,opt,name=TradingPhaseCode,proto3" json:"TradingPhaseCode,omitempty"`
	PrevClosePx          int64    `protobuf:"varint,8,opt,name=PrevClosePx,proto3" json:"PrevClosePx,omitempty"`
	NumTrades            int64    `protobuf:"varint,9,opt,name=NumTrades,proto3" json:"NumTrades,omitempty"`
	TotalVolumeTrade     int64    `protobuf:"varint,10,opt,name=TotalVolumeTrade,proto3" json:"TotalVolumeTrade,omitempty"`
	TotalValueTrade      int64    `protobuf:"varint,11,opt,name=TotalValueTrade,proto3" json:"TotalValueTrade,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Head3Xxx11) Reset()         { *m = Head3Xxx11{} }
func (m *Head3Xxx11) String() string { return proto.CompactTextString(m) }
func (*Head3Xxx11) ProtoMessage()    {}
func (*Head3Xxx11) Descriptor() ([]byte, []int) {
	return fileDescriptor_d39eb187a325c5cc, []int{0}
}

func (m *Head3Xxx11) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Head3Xxx11.Unmarshal(m, b)
}
func (m *Head3Xxx11) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Head3Xxx11.Marshal(b, m, deterministic)
}
func (m *Head3Xxx11) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Head3Xxx11.Merge(m, src)
}
func (m *Head3Xxx11) XXX_Size() int {
	return xxx_messageInfo_Head3Xxx11.Size(m)
}
func (m *Head3Xxx11) XXX_DiscardUnknown() {
	xxx_messageInfo_Head3Xxx11.DiscardUnknown(m)
}

var xxx_messageInfo_Head3Xxx11 proto.InternalMessageInfo

func (m *Head3Xxx11) GetMsgType() int32 {
	if m != nil {
		return m.MsgType
	}
	return 0
}

func (m *Head3Xxx11) GetOrigTime() int64 {
	if m != nil {
		return m.OrigTime
	}
	return 0
}

func (m *Head3Xxx11) GetChannelNo() int32 {
	if m != nil {
		return m.ChannelNo
	}
	return 0
}

func (m *Head3Xxx11) GetMDStreamID() string {
	if m != nil {
		return m.MDStreamID
	}
	return ""
}

func (m *Head3Xxx11) GetSecurityID() string {
	if m != nil {
		return m.SecurityID
	}
	return ""
}

func (m *Head3Xxx11) GetSecurityIDSource() string {
	if m != nil {
		return m.SecurityIDSource
	}
	return ""
}

func (m *Head3Xxx11) GetTradingPhaseCode() []byte {
	if m != nil {
		return m.TradingPhaseCode
	}
	return nil
}

func (m *Head3Xxx11) GetPrevClosePx() int64 {
	if m != nil {
		return m.PrevClosePx
	}
	return 0
}

func (m *Head3Xxx11) GetNumTrades() int64 {
	if m != nil {
		return m.NumTrades
	}
	return 0
}

func (m *Head3Xxx11) GetTotalVolumeTrade() int64 {
	if m != nil {
		return m.TotalVolumeTrade
	}
	return 0
}

func (m *Head3Xxx11) GetTotalValueTrade() int64 {
	if m != nil {
		return m.TotalValueTrade
	}
	return 0
}

//集中竞价交易业务行情快照扩展字段
type MDEntry300111 struct {
	MDEntryType          string   `protobuf:"bytes,1,opt,name=MDEntryType,proto3" json:"MDEntryType,omitempty"`
	MDEntryPx            int64    `protobuf:"varint,2,opt,name=MDEntryPx,proto3" json:"MDEntryPx,omitempty"`
	MDEntrySize          int64    `protobuf:"varint,3,opt,name=MDEntrySize,proto3" json:"MDEntrySize,omitempty"`
	MDPriceLevel         int32    `protobuf:"varint,4,opt,name=MDPriceLevel,proto3" json:"MDPriceLevel,omitempty"`
	NumberOfOrders       int64    `protobuf:"varint,5,opt,name=NumberOfOrders,proto3" json:"NumberOfOrders,omitempty"`
	NoOrders             uint32   `protobuf:"varint,6,opt,name=NoOrders,proto3" json:"NoOrders,omitempty"`
	OrderQty             int64    `protobuf:"varint,7,opt,name=OrderQty,proto3" json:"OrderQty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MDEntry300111) Reset()         { *m = MDEntry300111{} }
func (m *MDEntry300111) String() string { return proto.CompactTextString(m) }
func (*MDEntry300111) ProtoMessage()    {}
func (*MDEntry300111) Descriptor() ([]byte, []int) {
	return fileDescriptor_d39eb187a325c5cc, []int{1}
}

func (m *MDEntry300111) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MDEntry300111.Unmarshal(m, b)
}
func (m *MDEntry300111) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MDEntry300111.Marshal(b, m, deterministic)
}
func (m *MDEntry300111) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MDEntry300111.Merge(m, src)
}
func (m *MDEntry300111) XXX_Size() int {
	return xxx_messageInfo_MDEntry300111.Size(m)
}
func (m *MDEntry300111) XXX_DiscardUnknown() {
	xxx_messageInfo_MDEntry300111.DiscardUnknown(m)
}

var xxx_messageInfo_MDEntry300111 proto.InternalMessageInfo

func (m *MDEntry300111) GetMDEntryType() string {
	if m != nil {
		return m.MDEntryType
	}
	return ""
}

func (m *MDEntry300111) GetMDEntryPx() int64 {
	if m != nil {
		return m.MDEntryPx
	}
	return 0
}

func (m *MDEntry300111) GetMDEntrySize() int64 {
	if m != nil {
		return m.MDEntrySize
	}
	return 0
}

func (m *MDEntry300111) GetMDPriceLevel() int32 {
	if m != nil {
		return m.MDPriceLevel
	}
	return 0
}

func (m *MDEntry300111) GetNumberOfOrders() int64 {
	if m != nil {
		return m.NumberOfOrders
	}
	return 0
}

func (m *MDEntry300111) GetNoOrders() uint32 {
	if m != nil {
		return m.NoOrders
	}
	return 0
}

func (m *MDEntry300111) GetOrderQty() int64 {
	if m != nil {
		return m.OrderQty
	}
	return 0
}

type Quotation300111 struct {
	Head                 *Head3Xxx11      `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	NoMDEntries          int32            `protobuf:"varint,2,opt,name=NoMDEntries,proto3" json:"NoMDEntries,omitempty"`
	MDEntry              []*MDEntry300111 `protobuf:"bytes,3,rep,name=MDEntry,proto3" json:"MDEntry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Quotation300111) Reset()         { *m = Quotation300111{} }
func (m *Quotation300111) String() string { return proto.CompactTextString(m) }
func (*Quotation300111) ProtoMessage()    {}
func (*Quotation300111) Descriptor() ([]byte, []int) {
	return fileDescriptor_d39eb187a325c5cc, []int{2}
}

func (m *Quotation300111) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Quotation300111.Unmarshal(m, b)
}
func (m *Quotation300111) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Quotation300111.Marshal(b, m, deterministic)
}
func (m *Quotation300111) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quotation300111.Merge(m, src)
}
func (m *Quotation300111) XXX_Size() int {
	return xxx_messageInfo_Quotation300111.Size(m)
}
func (m *Quotation300111) XXX_DiscardUnknown() {
	xxx_messageInfo_Quotation300111.DiscardUnknown(m)
}

var xxx_messageInfo_Quotation300111 proto.InternalMessageInfo

func (m *Quotation300111) GetHead() *Head3Xxx11 {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *Quotation300111) GetNoMDEntries() int32 {
	if m != nil {
		return m.NoMDEntries
	}
	return 0
}

func (m *Quotation300111) GetMDEntry() []*MDEntry300111 {
	if m != nil {
		return m.MDEntry
	}
	return nil
}

//指数行情快照扩展字段
type MDEntry309011 struct {
	MDEntryType          int32    `protobuf:"varint,1,opt,name=MDEntryType,proto3" json:"MDEntryType,omitempty"`
	MDEntryPx            int32    `protobuf:"varint,2,opt,name=MDEntryPx,proto3" json:"MDEntryPx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MDEntry309011) Reset()         { *m = MDEntry309011{} }
func (m *MDEntry309011) String() string { return proto.CompactTextString(m) }
func (*MDEntry309011) ProtoMessage()    {}
func (*MDEntry309011) Descriptor() ([]byte, []int) {
	return fileDescriptor_d39eb187a325c5cc, []int{3}
}

func (m *MDEntry309011) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MDEntry309011.Unmarshal(m, b)
}
func (m *MDEntry309011) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MDEntry309011.Marshal(b, m, deterministic)
}
func (m *MDEntry309011) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MDEntry309011.Merge(m, src)
}
func (m *MDEntry309011) XXX_Size() int {
	return xxx_messageInfo_MDEntry309011.Size(m)
}
func (m *MDEntry309011) XXX_DiscardUnknown() {
	xxx_messageInfo_MDEntry309011.DiscardUnknown(m)
}

var xxx_messageInfo_MDEntry309011 proto.InternalMessageInfo

func (m *MDEntry309011) GetMDEntryType() int32 {
	if m != nil {
		return m.MDEntryType
	}
	return 0
}

func (m *MDEntry309011) GetMDEntryPx() int32 {
	if m != nil {
		return m.MDEntryPx
	}
	return 0
}

type Quotation309011 struct {
	Head                 *Head3Xxx11      `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	NoMDEntries          int32            `protobuf:"varint,2,opt,name=NoMDEntries,proto3" json:"NoMDEntries,omitempty"`
	MDEntry              []*MDEntry309011 `protobuf:"bytes,3,rep,name=MDEntry,proto3" json:"MDEntry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Quotation309011) Reset()         { *m = Quotation309011{} }
func (m *Quotation309011) String() string { return proto.CompactTextString(m) }
func (*Quotation309011) ProtoMessage()    {}
func (*Quotation309011) Descriptor() ([]byte, []int) {
	return fileDescriptor_d39eb187a325c5cc, []int{4}
}

func (m *Quotation309011) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Quotation309011.Unmarshal(m, b)
}
func (m *Quotation309011) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Quotation309011.Marshal(b, m, deterministic)
}
func (m *Quotation309011) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quotation309011.Merge(m, src)
}
func (m *Quotation309011) XXX_Size() int {
	return xxx_messageInfo_Quotation309011.Size(m)
}
func (m *Quotation309011) XXX_DiscardUnknown() {
	xxx_messageInfo_Quotation309011.DiscardUnknown(m)
}

var xxx_messageInfo_Quotation309011 proto.InternalMessageInfo

func (m *Quotation309011) GetHead() *Head3Xxx11 {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *Quotation309011) GetNoMDEntries() int32 {
	if m != nil {
		return m.NoMDEntries
	}
	return 0
}

func (m *Quotation309011) GetMDEntry() []*MDEntry309011 {
	if m != nil {
		return m.MDEntry
	}
	return nil
}

// MDGW发布逐笔委托行情消息定义
type Head30Xx92 struct {
	MsgType              int32    `protobuf:"varint,1,opt,name=MsgType,proto3" json:"MsgType,omitempty"`
	ChannelNo            int64    `protobuf:"varint,2,opt,name=ChannelNo,proto3" json:"ChannelNo,omitempty"`
	ApplSeqNum           int64    `protobuf:"varint,3,opt,name=ApplSeqNum,proto3" json:"ApplSeqNum,omitempty"`
	MDStreamID           int64    `protobuf:"varint,4,opt,name=MDStreamID,proto3" json:"MDStreamID,omitempty"`
	SecurityID           string   `protobuf:"bytes,5,opt,name=SecurityID,proto3" json:"SecurityID,omitempty"`
	SecurityIDSource     string   `protobuf:"bytes,6,opt,name=SecurityIDSource,proto3" json:"SecurityIDSource,omitempty"`
	Price                string   `protobuf:"bytes,7,opt,name=Price,proto3" json:"Price,omitempty"`
	OrderQty             float64  `protobuf:"fixed64,8,opt,name=OrderQty,proto3" json:"OrderQty,omitempty"`
	Side                 int64    `protobuf:"varint,9,opt,name=Side,proto3" json:"Side,omitempty"`
	TransactTime         int64    `protobuf:"varint,10,opt,name=TransactTime,proto3" json:"TransactTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Head30Xx92) Reset()         { *m = Head30Xx92{} }
func (m *Head30Xx92) String() string { return proto.CompactTextString(m) }
func (*Head30Xx92) ProtoMessage()    {}
func (*Head30Xx92) Descriptor() ([]byte, []int) {
	return fileDescriptor_d39eb187a325c5cc, []int{5}
}

func (m *Head30Xx92) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Head30Xx92.Unmarshal(m, b)
}
func (m *Head30Xx92) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Head30Xx92.Marshal(b, m, deterministic)
}
func (m *Head30Xx92) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Head30Xx92.Merge(m, src)
}
func (m *Head30Xx92) XXX_Size() int {
	return xxx_messageInfo_Head30Xx92.Size(m)
}
func (m *Head30Xx92) XXX_DiscardUnknown() {
	xxx_messageInfo_Head30Xx92.DiscardUnknown(m)
}

var xxx_messageInfo_Head30Xx92 proto.InternalMessageInfo

func (m *Head30Xx92) GetMsgType() int32 {
	if m != nil {
		return m.MsgType
	}
	return 0
}

func (m *Head30Xx92) GetChannelNo() int64 {
	if m != nil {
		return m.ChannelNo
	}
	return 0
}

func (m *Head30Xx92) GetApplSeqNum() int64 {
	if m != nil {
		return m.ApplSeqNum
	}
	return 0
}

func (m *Head30Xx92) GetMDStreamID() int64 {
	if m != nil {
		return m.MDStreamID
	}
	return 0
}

func (m *Head30Xx92) GetSecurityID() string {
	if m != nil {
		return m.SecurityID
	}
	return ""
}

func (m *Head30Xx92) GetSecurityIDSource() string {
	if m != nil {
		return m.SecurityIDSource
	}
	return ""
}

func (m *Head30Xx92) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Head30Xx92) GetOrderQty() float64 {
	if m != nil {
		return m.OrderQty
	}
	return 0
}

func (m *Head30Xx92) GetSide() int64 {
	if m != nil {
		return m.Side
	}
	return 0
}

func (m *Head30Xx92) GetTransactTime() int64 {
	if m != nil {
		return m.TransactTime
	}
	return 0
}

type Quotation300192 struct {
	Head                 *Head30Xx92 `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	OrdType              int32       `protobuf:"varint,2,opt,name=OrdType,proto3" json:"OrdType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Quotation300192) Reset()         { *m = Quotation300192{} }
func (m *Quotation300192) String() string { return proto.CompactTextString(m) }
func (*Quotation300192) ProtoMessage()    {}
func (*Quotation300192) Descriptor() ([]byte, []int) {
	return fileDescriptor_d39eb187a325c5cc, []int{6}
}

func (m *Quotation300192) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Quotation300192.Unmarshal(m, b)
}
func (m *Quotation300192) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Quotation300192.Marshal(b, m, deterministic)
}
func (m *Quotation300192) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quotation300192.Merge(m, src)
}
func (m *Quotation300192) XXX_Size() int {
	return xxx_messageInfo_Quotation300192.Size(m)
}
func (m *Quotation300192) XXX_DiscardUnknown() {
	xxx_messageInfo_Quotation300192.DiscardUnknown(m)
}

var xxx_messageInfo_Quotation300192 proto.InternalMessageInfo

func (m *Quotation300192) GetHead() *Head30Xx92 {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *Quotation300192) GetOrdType() int32 {
	if m != nil {
		return m.OrdType
	}
	return 0
}

// MDGW发布逐笔成交行情消息定义
type Head30Xx91 struct {
	MsgType              int32    `protobuf:"varint,1,opt,name=MsgType,proto3" json:"MsgType,omitempty"`
	ChannelNo            int64    `protobuf:"varint,2,opt,name=ChannelNo,proto3" json:"ChannelNo,omitempty"`
	ApplSeqNum           int64    `protobuf:"varint,3,opt,name=ApplSeqNum,proto3" json:"ApplSeqNum,omitempty"`
	MDStreamID           int64    `protobuf:"varint,4,opt,name=MDStreamID,proto3" json:"MDStreamID,omitempty"`
	BidApplSeqNum        string   `protobuf:"bytes,5,opt,name=BidApplSeqNum,proto3" json:"BidApplSeqNum,omitempty"`
	OfferApplSeqNum      string   `protobuf:"bytes,6,opt,name=OfferApplSeqNum,proto3" json:"OfferApplSeqNum,omitempty"`
	SecurityID           string   `protobuf:"bytes,7,opt,name=SecurityID,proto3" json:"SecurityID,omitempty"`
	SecurityIDSource     float64  `protobuf:"fixed64,8,opt,name=SecurityIDSource,proto3" json:"SecurityIDSource,omitempty"`
	LastPx               int64    `protobuf:"varint,9,opt,name=LastPx,proto3" json:"LastPx,omitempty"`
	LastQty              int64    `protobuf:"varint,10,opt,name=LastQty,proto3" json:"LastQty,omitempty"`
	ExecType             int64    `protobuf:"varint,11,opt,name=ExecType,proto3" json:"ExecType,omitempty"`
	TransactTime         int64    `protobuf:"varint,12,opt,name=TransactTime,proto3" json:"TransactTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Head30Xx91) Reset()         { *m = Head30Xx91{} }
func (m *Head30Xx91) String() string { return proto.CompactTextString(m) }
func (*Head30Xx91) ProtoMessage()    {}
func (*Head30Xx91) Descriptor() ([]byte, []int) {
	return fileDescriptor_d39eb187a325c5cc, []int{7}
}

func (m *Head30Xx91) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Head30Xx91.Unmarshal(m, b)
}
func (m *Head30Xx91) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Head30Xx91.Marshal(b, m, deterministic)
}
func (m *Head30Xx91) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Head30Xx91.Merge(m, src)
}
func (m *Head30Xx91) XXX_Size() int {
	return xxx_messageInfo_Head30Xx91.Size(m)
}
func (m *Head30Xx91) XXX_DiscardUnknown() {
	xxx_messageInfo_Head30Xx91.DiscardUnknown(m)
}

var xxx_messageInfo_Head30Xx91 proto.InternalMessageInfo

func (m *Head30Xx91) GetMsgType() int32 {
	if m != nil {
		return m.MsgType
	}
	return 0
}

func (m *Head30Xx91) GetChannelNo() int64 {
	if m != nil {
		return m.ChannelNo
	}
	return 0
}

func (m *Head30Xx91) GetApplSeqNum() int64 {
	if m != nil {
		return m.ApplSeqNum
	}
	return 0
}

func (m *Head30Xx91) GetMDStreamID() int64 {
	if m != nil {
		return m.MDStreamID
	}
	return 0
}

func (m *Head30Xx91) GetBidApplSeqNum() string {
	if m != nil {
		return m.BidApplSeqNum
	}
	return ""
}

func (m *Head30Xx91) GetOfferApplSeqNum() string {
	if m != nil {
		return m.OfferApplSeqNum
	}
	return ""
}

func (m *Head30Xx91) GetSecurityID() string {
	if m != nil {
		return m.SecurityID
	}
	return ""
}

func (m *Head30Xx91) GetSecurityIDSource() float64 {
	if m != nil {
		return m.SecurityIDSource
	}
	return 0
}

func (m *Head30Xx91) GetLastPx() int64 {
	if m != nil {
		return m.LastPx
	}
	return 0
}

func (m *Head30Xx91) GetLastQty() int64 {
	if m != nil {
		return m.LastQty
	}
	return 0
}

func (m *Head30Xx91) GetExecType() int64 {
	if m != nil {
		return m.ExecType
	}
	return 0
}

func (m *Head30Xx91) GetTransactTime() int64 {
	if m != nil {
		return m.TransactTime
	}
	return 0
}

type Quotation300191 struct {
	Head                 *Head30Xx92 `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Quotation300191) Reset()         { *m = Quotation300191{} }
func (m *Quotation300191) String() string { return proto.CompactTextString(m) }
func (*Quotation300191) ProtoMessage()    {}
func (*Quotation300191) Descriptor() ([]byte, []int) {
	return fileDescriptor_d39eb187a325c5cc, []int{8}
}

func (m *Quotation300191) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Quotation300191.Unmarshal(m, b)
}
func (m *Quotation300191) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Quotation300191.Marshal(b, m, deterministic)
}
func (m *Quotation300191) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quotation300191.Merge(m, src)
}
func (m *Quotation300191) XXX_Size() int {
	return xxx_messageInfo_Quotation300191.Size(m)
}
func (m *Quotation300191) XXX_DiscardUnknown() {
	xxx_messageInfo_Quotation300191.DiscardUnknown(m)
}

var xxx_messageInfo_Quotation300191 proto.InternalMessageInfo

func (m *Quotation300191) GetHead() *Head30Xx92 {
	if m != nil {
		return m.Head
	}
	return nil
}

func init() {
	proto.RegisterType((*Head3Xxx11)(nil), "protobuf.Head3xxx11")
	proto.RegisterType((*MDEntry300111)(nil), "protobuf.MDEntry300111")
	proto.RegisterType((*Quotation300111)(nil), "protobuf.Quotation300111")
	proto.RegisterType((*MDEntry309011)(nil), "protobuf.MDEntry309011")
	proto.RegisterType((*Quotation309011)(nil), "protobuf.Quotation309011")
	proto.RegisterType((*Head30Xx92)(nil), "protobuf.Head30xx92")
	proto.RegisterType((*Quotation300192)(nil), "protobuf.Quotation300192")
	proto.RegisterType((*Head30Xx91)(nil), "protobuf.Head30xx91")
	proto.RegisterType((*Quotation300191)(nil), "protobuf.Quotation300191")
}

func init() { proto.RegisterFile("quotation.proto", fileDescriptor_d39eb187a325c5cc) }

var fileDescriptor_d39eb187a325c5cc = []byte{
	// 665 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0x4d, 0x6b, 0xdb, 0x4c,
	0x10, 0xc7, 0x51, 0x14, 0xc5, 0xf1, 0x3a, 0x79, 0xf2, 0xb0, 0x84, 0x76, 0x29, 0xa5, 0x18, 0x51,
	0x8a, 0xe8, 0x21, 0x58, 0xce, 0x29, 0xf4, 0xd4, 0xc6, 0x81, 0x06, 0xe2, 0x97, 0xac, 0xdd, 0xde,
	0x15, 0x6b, 0x9c, 0x08, 0x64, 0xcb, 0x59, 0x49, 0x41, 0xee, 0x87, 0x28, 0xa5, 0x1f, 0xa7, 0xdf,
	0xab, 0xf7, 0xb2, 0xb3, 0x2b, 0xeb, 0xc5, 0xa6, 0xc9, 0xa5, 0xed, 0x49, 0x9a, 0xff, 0xce, 0xae,
	0x34, 0xf3, 0xfb, 0xcf, 0x92, 0xa3, 0xfb, 0x34, 0x4a, 0xbc, 0x24, 0x88, 0x16, 0x27, 0x4b, 0x11,
	0x25, 0x11, 0xdd, 0xc7, 0xc7, 0x4d, 0x3a, 0xb3, 0xbf, 0x9a, 0x84, 0x7c, 0x04, 0xcf, 0x3f, 0xcd,
	0xb2, 0xcc, 0x75, 0x29, 0x23, 0x8d, 0x7e, 0x7c, 0x3b, 0x59, 0x2d, 0x81, 0x19, 0x6d, 0xc3, 0xb1,
	0x78, 0x1e, 0xd2, 0x17, 0x64, 0x7f, 0x28, 0x82, 0xdb, 0x49, 0x30, 0x07, 0xb6, 0xd3, 0x36, 0x1c,
	0x93, 0xaf, 0x63, 0xfa, 0x92, 0x34, 0xcf, 0xef, 0xbc, 0xc5, 0x02, 0xc2, 0x41, 0xc4, 0x4c, 0xdc,
	0x57, 0x08, 0xf4, 0x15, 0x21, 0xfd, 0xde, 0x38, 0x11, 0xe0, 0xcd, 0x2f, 0x7b, 0x6c, 0xb7, 0x6d,
	0x38, 0x4d, 0x5e, 0x52, 0xe4, 0xfa, 0x18, 0xa6, 0xa9, 0x08, 0x92, 0xd5, 0x65, 0x8f, 0x59, 0x6a,
	0xbd, 0x50, 0xe8, 0x5b, 0xf2, 0x7f, 0x11, 0x8d, 0xa3, 0x54, 0x4c, 0x81, 0xed, 0x61, 0xd6, 0x86,
	0x2e, 0x73, 0x27, 0xc2, 0xf3, 0x83, 0xc5, 0xed, 0xe8, 0xce, 0x8b, 0xe1, 0x3c, 0xf2, 0x81, 0x35,
	0xda, 0x86, 0x73, 0xc0, 0x37, 0x74, 0xda, 0x26, 0xad, 0x91, 0x80, 0x87, 0xf3, 0x30, 0x8a, 0x61,
	0x94, 0xb1, 0x7d, 0x2c, 0xaa, 0x2c, 0xc9, 0xba, 0x06, 0xe9, 0x5c, 0x6e, 0x84, 0x98, 0x35, 0x71,
	0xbd, 0x10, 0xf0, 0x5b, 0x51, 0xe2, 0x85, 0x9f, 0xa3, 0x30, 0x9d, 0x03, 0x8a, 0x8c, 0x60, 0xd2,
	0x86, 0x4e, 0x1d, 0x72, 0xa4, 0x34, 0x2f, 0x4c, 0x75, 0x6a, 0x0b, 0x53, 0xeb, 0xb2, 0xfd, 0xd3,
	0x20, 0x87, 0xfd, 0xde, 0xc5, 0x22, 0x11, 0xab, 0xd3, 0x4e, 0xc7, 0x75, 0x5d, 0xf9, 0x9f, 0x5a,
	0x58, 0x73, 0x69, 0xf2, 0xb2, 0x24, 0xff, 0x53, 0x87, 0xa3, 0x4c, 0xc3, 0x29, 0x84, 0xd2, 0xfe,
	0x71, 0xf0, 0x05, 0x90, 0x8f, 0xc9, 0xcb, 0x12, 0xb5, 0xc9, 0x41, 0xbf, 0x37, 0x12, 0xc1, 0x14,
	0xae, 0xe0, 0x01, 0x42, 0x64, 0x64, 0xf1, 0x8a, 0x46, 0xdf, 0x90, 0xff, 0x06, 0xe9, 0xfc, 0x06,
	0xc4, 0x70, 0x36, 0x14, 0x3e, 0x88, 0x18, 0x49, 0x99, 0xbc, 0xa6, 0x4a, 0x9f, 0x0c, 0x22, 0x9d,
	0x21, 0x29, 0x1d, 0xf2, 0x75, 0xac, 0x3c, 0xe4, 0x83, 0xb8, 0x4e, 0x56, 0x48, 0x05, 0x3d, 0xa4,
	0x62, 0xfb, 0x9b, 0x41, 0x8e, 0xae, 0x73, 0x9b, 0xea, 0xca, 0x1d, 0xb2, 0x2b, 0xbd, 0x89, 0x25,
	0xb7, 0xba, 0xc7, 0x27, 0xb9, 0x6b, 0x4f, 0x0a, 0xc7, 0x72, 0xcc, 0x90, 0x35, 0x0e, 0x22, 0x55,
	0x52, 0x00, 0x31, 0xf6, 0xc0, 0xe2, 0x65, 0x89, 0xba, 0xa4, 0xa1, 0x4b, 0x66, 0x66, 0xdb, 0x74,
	0x5a, 0xdd, 0xe7, 0xc5, 0x71, 0x95, 0x7e, 0xf3, 0x3c, 0xcf, 0x1e, 0x96, 0x48, 0x9c, 0x75, 0xb6,
	0x93, 0xb0, 0x1e, 0x21, 0x61, 0x95, 0x48, 0xd4, 0x6b, 0xc4, 0x33, 0xff, 0x7a, 0x8d, 0xf2, 0xab,
	0x45, 0x8d, 0x3f, 0x76, 0xf4, 0xfc, 0x77, 0xb2, 0xec, 0xac, 0xfb, 0x9b, 0xf9, 0xaf, 0xcc, 0xb8,
	0xf6, 0x58, 0x65, 0xc6, 0xdf, 0x2f, 0x97, 0xe1, 0x18, 0xee, 0x07, 0xe9, 0x5c, 0x5b, 0xac, 0xa4,
	0x6c, 0xb9, 0x03, 0xcc, 0x3f, 0x76, 0x07, 0x1c, 0x13, 0x0b, 0x7d, 0x8b, 0x16, 0x6b, 0x72, 0x15,
	0x54, 0xbc, 0x27, 0x47, 0xdd, 0x28, 0xbc, 0x47, 0x29, 0xd9, 0x1d, 0x07, 0x3e, 0xe8, 0x11, 0xc7,
	0x77, 0x39, 0x13, 0x13, 0xe1, 0x2d, 0x62, 0x6f, 0x9a, 0xe0, 0x9d, 0xa7, 0x26, 0xbb, 0xa2, 0xd9,
	0x9f, 0x6a, 0x96, 0x3d, 0xeb, 0x3e, 0x82, 0x13, 0x9b, 0xac, 0x71, 0x32, 0xd2, 0x18, 0x0a, 0x1f,
	0x5b, 0xad, 0x50, 0xe6, 0xa1, 0xfd, 0xdd, 0x2c, 0x31, 0x71, 0xff, 0x19, 0x93, 0xd7, 0xe4, 0xf0,
	0x43, 0xe0, 0x97, 0x8e, 0x50, 0x58, 0xaa, 0xa2, 0xbc, 0xd9, 0x86, 0xb3, 0x19, 0x88, 0x52, 0x9e,
	0x02, 0x53, 0x97, 0x6b, 0x8c, 0x1b, 0x4f, 0x62, 0xac, 0x48, 0x6d, 0x32, 0x7e, 0x46, 0xf6, 0xae,
	0xbc, 0x38, 0x19, 0x65, 0x9a, 0x99, 0x8e, 0x64, 0xaf, 0xe4, 0x9b, 0x84, 0xac, 0x80, 0xe5, 0xa1,
	0xe4, 0x7f, 0x91, 0xc1, 0x14, 0xdb, 0xa8, 0xae, 0xde, 0x75, 0xbc, 0xc1, 0xfa, 0x60, 0x0b, 0xeb,
	0x77, 0x75, 0xd6, 0xee, 0xd3, 0x59, 0xdf, 0xec, 0xe1, 0xd2, 0xe9, 0xaf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xb2, 0xb0, 0x2b, 0x72, 0x89, 0x07, 0x00, 0x00,
}