// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: chat.proto

package chat_server

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EmptyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResp) Reset() {
	*x = EmptyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResp) ProtoMessage() {}

func (x *EmptyResp) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResp.ProtoReflect.Descriptor instead.
func (*EmptyResp) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

type CreateChatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 群名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 1 = 普通群聊
	Type int32 `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	// 需要加入群的用户
	Uids  []int64 `protobuf:"varint,3,rep,packed,name=uids,proto3" json:"uids,omitempty"`
	Token string  `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CreateChatReq) Reset() {
	*x = CreateChatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatReq) ProtoMessage() {}

func (x *CreateChatReq) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatReq.ProtoReflect.Descriptor instead.
func (*CreateChatReq) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *CreateChatReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateChatReq) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *CreateChatReq) GetUids() []int64 {
	if x != nil {
		return x.Uids
	}
	return nil
}

func (x *CreateChatReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CreateChatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 群聊ID
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateChatResp) Reset() {
	*x = CreateChatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatResp) ProtoMessage() {}

func (x *CreateChatResp) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatResp.ProtoReflect.Descriptor instead.
func (*CreateChatResp) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *CreateChatResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sender  int64  `protobuf:"varint,2,opt,name=sender,proto3" json:"sender,omitempty"`
	ChatId  int64  `protobuf:"varint,3,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Seq     int64  `protobuf:"varint,5,opt,name=seq,proto3" json:"seq,omitempty"`
	Type    int64  `protobuf:"varint,6,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{3}
}

func (x *Msg) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Msg) GetSender() int64 {
	if x != nil {
		return x.Sender
	}
	return 0
}

func (x *Msg) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *Msg) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Msg) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *Msg) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

type RecentChatsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextSeq int64  `protobuf:"varint,1,opt,name=next_seq,json=nextSeq,proto3" json:"next_seq,omitempty"`
	Cnt     int64  `protobuf:"varint,2,opt,name=cnt,proto3" json:"cnt,omitempty"`
	Order   string `protobuf:"bytes,3,opt,name=order,proto3" json:"order,omitempty"`
	Token   string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RecentChatsReq) Reset() {
	*x = RecentChatsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecentChatsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecentChatsReq) ProtoMessage() {}

func (x *RecentChatsReq) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecentChatsReq.ProtoReflect.Descriptor instead.
func (*RecentChatsReq) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{4}
}

func (x *RecentChatsReq) GetNextSeq() int64 {
	if x != nil {
		return x.NextSeq
	}
	return 0
}

func (x *RecentChatsReq) GetCnt() int64 {
	if x != nil {
		return x.Cnt
	}
	return 0
}

func (x *RecentChatsReq) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

func (x *RecentChatsReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RecentChatsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type    int32  `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Creator string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
	Ctime   int64  `protobuf:"varint,5,opt,name=ctime,proto3" json:"ctime,omitempty"`
	LastMsg *Msg   `protobuf:"bytes,6,opt,name=last_msg,json=lastMsg,proto3" json:"last_msg,omitempty"`
}

func (x *RecentChatsResp) Reset() {
	*x = RecentChatsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecentChatsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecentChatsResp) ProtoMessage() {}

func (x *RecentChatsResp) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecentChatsResp.ProtoReflect.Descriptor instead.
func (*RecentChatsResp) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{5}
}

func (x *RecentChatsResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RecentChatsResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RecentChatsResp) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *RecentChatsResp) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *RecentChatsResp) GetCtime() int64 {
	if x != nil {
		return x.Ctime
	}
	return 0
}

func (x *RecentChatsResp) GetLastMsg() *Msg {
	if x != nil {
		return x.LastMsg
	}
	return nil
}

type JoinToChatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId int64   `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	Uids   []int64 `protobuf:"varint,2,rep,packed,name=uids,proto3" json:"uids,omitempty"`
	Token  string  `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *JoinToChatReq) Reset() {
	*x = JoinToChatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinToChatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinToChatReq) ProtoMessage() {}

func (x *JoinToChatReq) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinToChatReq.ProtoReflect.Descriptor instead.
func (*JoinToChatReq) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{6}
}

func (x *JoinToChatReq) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *JoinToChatReq) GetUids() []int64 {
	if x != nil {
		return x.Uids
	}
	return nil
}

func (x *JoinToChatReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RemoveMemberFromChatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId int64  `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	Uid    int64  `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Token  string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RemoveMemberFromChatReq) Reset() {
	*x = RemoveMemberFromChatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveMemberFromChatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveMemberFromChatReq) ProtoMessage() {}

func (x *RemoveMemberFromChatReq) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveMemberFromChatReq.ProtoReflect.Descriptor instead.
func (*RemoveMemberFromChatReq) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveMemberFromChatReq) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *RemoveMemberFromChatReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *RemoveMemberFromChatReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetChatMsgListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId  int64  `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	NextSeq int64  `protobuf:"varint,2,opt,name=next_seq,json=nextSeq,proto3" json:"next_seq,omitempty"`
	Cnt     int64  `protobuf:"varint,3,opt,name=cnt,proto3" json:"cnt,omitempty"`
	Token   string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetChatMsgListReq) Reset() {
	*x = GetChatMsgListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatMsgListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatMsgListReq) ProtoMessage() {}

func (x *GetChatMsgListReq) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatMsgListReq.ProtoReflect.Descriptor instead.
func (*GetChatMsgListReq) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{8}
}

func (x *GetChatMsgListReq) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *GetChatMsgListReq) GetNextSeq() int64 {
	if x != nil {
		return x.NextSeq
	}
	return 0
}

func (x *GetChatMsgListReq) GetCnt() int64 {
	if x != nil {
		return x.Cnt
	}
	return 0
}

func (x *GetChatMsgListReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetChatMsgListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msgs    []*Msg `protobuf:"bytes,1,rep,name=msgs,proto3" json:"msgs,omitempty"`
	Cnt     int64  `protobuf:"varint,2,opt,name=cnt,proto3" json:"cnt,omitempty"`
	NextSeq int64  `protobuf:"varint,3,opt,name=next_seq,json=nextSeq,proto3" json:"next_seq,omitempty"`
}

func (x *GetChatMsgListResp) Reset() {
	*x = GetChatMsgListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatMsgListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatMsgListResp) ProtoMessage() {}

func (x *GetChatMsgListResp) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatMsgListResp.ProtoReflect.Descriptor instead.
func (*GetChatMsgListResp) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{9}
}

func (x *GetChatMsgListResp) GetMsgs() []*Msg {
	if x != nil {
		return x.Msgs
	}
	return nil
}

func (x *GetChatMsgListResp) GetCnt() int64 {
	if x != nil {
		return x.Cnt
	}
	return 0
}

func (x *GetChatMsgListResp) GetNextSeq() int64 {
	if x != nil {
		return x.NextSeq
	}
	return 0
}

type GetMemberListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId  int64  `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	NextSeq int64  `protobuf:"varint,2,opt,name=next_seq,json=nextSeq,proto3" json:"next_seq,omitempty"`
	Cnt     int64  `protobuf:"varint,3,opt,name=cnt,proto3" json:"cnt,omitempty"`
	Token   string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetMemberListReq) Reset() {
	*x = GetMemberListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMemberListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemberListReq) ProtoMessage() {}

func (x *GetMemberListReq) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemberListReq.ProtoReflect.Descriptor instead.
func (*GetMemberListReq) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{10}
}

func (x *GetMemberListReq) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *GetMemberListReq) GetNextSeq() int64 {
	if x != nil {
		return x.NextSeq
	}
	return 0
}

func (x *GetMemberListReq) GetCnt() int64 {
	if x != nil {
		return x.Cnt
	}
	return 0
}

func (x *GetMemberListReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetMemberListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NextSeq    int64     `protobuf:"varint,2,opt,name=next_seq,json=nextSeq,proto3" json:"next_seq,omitempty"`
	Cnt        int64     `protobuf:"varint,3,opt,name=cnt,proto3" json:"cnt,omitempty"`
	MemberList []*Member `protobuf:"bytes,4,rep,name=member_list,json=memberList,proto3" json:"member_list,omitempty"`
}

func (x *GetMemberListResp) Reset() {
	*x = GetMemberListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMemberListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemberListResp) ProtoMessage() {}

func (x *GetMemberListResp) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemberListResp.ProtoReflect.Descriptor instead.
func (*GetMemberListResp) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{11}
}

func (x *GetMemberListResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetMemberListResp) GetNextSeq() int64 {
	if x != nil {
		return x.NextSeq
	}
	return 0
}

func (x *GetMemberListResp) GetCnt() int64 {
	if x != nil {
		return x.Cnt
	}
	return 0
}

func (x *GetMemberListResp) GetMemberList() []*Member {
	if x != nil {
		return x.MemberList
	}
	return nil
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Ctime int64 `protobuf:"varint,2,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Role  int32 `protobuf:"varint,3,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{12}
}

func (x *Member) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Member) GetCtime() int64 {
	if x != nil {
		return x.Ctime
	}
	return 0
}

func (x *Member) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x0b, 0x0a, 0x09, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x61, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x04, 0x75,
	0x69, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x20, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x03,
	0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x63,
	0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68,
	0x61, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x69, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x68,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x73,
	0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x53, 0x65,
	0x71, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x63, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0xa6, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4d, 0x73, 0x67, 0x52,
	0x07, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x22, 0x52, 0x0a, 0x0d, 0x4a, 0x6f, 0x69, 0x6e,
	0x54, 0x6f, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x04, 0x75, 0x69, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5a, 0x0a, 0x17,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a,
	0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x73,
	0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x53, 0x65,
	0x71, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x63, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x67, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x43, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x24, 0x0a, 0x04, 0x6d, 0x73, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4d, 0x73, 0x67, 0x52,
	0x04, 0x6d, 0x73, 0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x63, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x73, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x53,
	0x65, 0x71, 0x22, 0x6e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x53, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x63, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x86, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6e, 0x65, 0x78, 0x74,
	0x53, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x63, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x06, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x32,
	0xd2, 0x03, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x12, 0x45, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x48, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x74, 0x73, 0x12, 0x1b,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x63,
	0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74,
	0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x40, 0x0a, 0x0a, 0x6a, 0x6f, 0x69,
	0x6e, 0x54, 0x6f, 0x43, 0x68, 0x61, 0x74, 0x12, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x54, 0x6f, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x54, 0x0a, 0x14, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x43,
	0x68, 0x61, 0x74, 0x12, 0x24, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x46, 0x72,
	0x6f, 0x6d, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x51, 0x0a, 0x0e, 0x67, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x4e, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_chat_proto_goTypes = []interface{}{
	(*EmptyResp)(nil),               // 0: chat_server.EmptyResp
	(*CreateChatReq)(nil),           // 1: chat_server.CreateChatReq
	(*CreateChatResp)(nil),          // 2: chat_server.CreateChatResp
	(*Msg)(nil),                     // 3: chat_server.Msg
	(*RecentChatsReq)(nil),          // 4: chat_server.RecentChatsReq
	(*RecentChatsResp)(nil),         // 5: chat_server.RecentChatsResp
	(*JoinToChatReq)(nil),           // 6: chat_server.JoinToChatReq
	(*RemoveMemberFromChatReq)(nil), // 7: chat_server.RemoveMemberFromChatReq
	(*GetChatMsgListReq)(nil),       // 8: chat_server.GetChatMsgListReq
	(*GetChatMsgListResp)(nil),      // 9: chat_server.GetChatMsgListResp
	(*GetMemberListReq)(nil),        // 10: chat_server.GetMemberListReq
	(*GetMemberListResp)(nil),       // 11: chat_server.GetMemberListResp
	(*Member)(nil),                  // 12: chat_server.Member
}
var file_chat_proto_depIdxs = []int32{
	3,  // 0: chat_server.RecentChatsResp.last_msg:type_name -> chat_server.Msg
	3,  // 1: chat_server.GetChatMsgListResp.msgs:type_name -> chat_server.Msg
	12, // 2: chat_server.GetMemberListResp.member_list:type_name -> chat_server.Member
	1,  // 3: chat_server.chat.createChat:input_type -> chat_server.CreateChatReq
	4,  // 4: chat_server.chat.recentChats:input_type -> chat_server.RecentChatsReq
	6,  // 5: chat_server.chat.joinToChat:input_type -> chat_server.JoinToChatReq
	7,  // 6: chat_server.chat.removeMemberFromChat:input_type -> chat_server.RemoveMemberFromChatReq
	8,  // 7: chat_server.chat.getChatMsgList:input_type -> chat_server.GetChatMsgListReq
	10, // 8: chat_server.chat.getMemberList:input_type -> chat_server.GetMemberListReq
	2,  // 9: chat_server.chat.createChat:output_type -> chat_server.CreateChatResp
	5,  // 10: chat_server.chat.recentChats:output_type -> chat_server.RecentChatsResp
	0,  // 11: chat_server.chat.joinToChat:output_type -> chat_server.EmptyResp
	0,  // 12: chat_server.chat.removeMemberFromChat:output_type -> chat_server.EmptyResp
	9,  // 13: chat_server.chat.getChatMsgList:output_type -> chat_server.GetChatMsgListResp
	11, // 14: chat_server.chat.getMemberList:output_type -> chat_server.GetMemberListResp
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResp); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChatReq); i {
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
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChatResp); i {
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
		file_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
		file_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecentChatsReq); i {
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
		file_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecentChatsResp); i {
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
		file_chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinToChatReq); i {
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
		file_chat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveMemberFromChatReq); i {
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
		file_chat_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatMsgListReq); i {
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
		file_chat_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatMsgListResp); i {
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
		file_chat_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMemberListReq); i {
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
		file_chat_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMemberListResp); i {
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
		file_chat_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}
