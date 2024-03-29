// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: chat.proto

package chat_server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Chat_CreateChat_FullMethodName           = "/chat_server.chat/createChat"
	Chat_RecentChats_FullMethodName          = "/chat_server.chat/recentChats"
	Chat_JoinToChat_FullMethodName           = "/chat_server.chat/joinToChat"
	Chat_RemoveMemberFromChat_FullMethodName = "/chat_server.chat/removeMemberFromChat"
	Chat_GetChatMsgList_FullMethodName       = "/chat_server.chat/getChatMsgList"
	Chat_GetMemberList_FullMethodName        = "/chat_server.chat/getMemberList"
)

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	CreateChat(ctx context.Context, in *CreateChatReq, opts ...grpc.CallOption) (*CreateChatResp, error)
	RecentChats(ctx context.Context, in *RecentChatsReq, opts ...grpc.CallOption) (*RecentChatsResp, error)
	JoinToChat(ctx context.Context, in *JoinToChatReq, opts ...grpc.CallOption) (*EmptyResp, error)
	RemoveMemberFromChat(ctx context.Context, in *RemoveMemberFromChatReq, opts ...grpc.CallOption) (*EmptyResp, error)
	GetChatMsgList(ctx context.Context, in *GetChatMsgListReq, opts ...grpc.CallOption) (*GetChatMsgListResp, error)
	GetMemberList(ctx context.Context, in *GetMemberListReq, opts ...grpc.CallOption) (*GetMemberListResp, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) CreateChat(ctx context.Context, in *CreateChatReq, opts ...grpc.CallOption) (*CreateChatResp, error) {
	out := new(CreateChatResp)
	err := c.cc.Invoke(ctx, Chat_CreateChat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) RecentChats(ctx context.Context, in *RecentChatsReq, opts ...grpc.CallOption) (*RecentChatsResp, error) {
	out := new(RecentChatsResp)
	err := c.cc.Invoke(ctx, Chat_RecentChats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) JoinToChat(ctx context.Context, in *JoinToChatReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	out := new(EmptyResp)
	err := c.cc.Invoke(ctx, Chat_JoinToChat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) RemoveMemberFromChat(ctx context.Context, in *RemoveMemberFromChatReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	out := new(EmptyResp)
	err := c.cc.Invoke(ctx, Chat_RemoveMemberFromChat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetChatMsgList(ctx context.Context, in *GetChatMsgListReq, opts ...grpc.CallOption) (*GetChatMsgListResp, error) {
	out := new(GetChatMsgListResp)
	err := c.cc.Invoke(ctx, Chat_GetChatMsgList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetMemberList(ctx context.Context, in *GetMemberListReq, opts ...grpc.CallOption) (*GetMemberListResp, error) {
	out := new(GetMemberListResp)
	err := c.cc.Invoke(ctx, Chat_GetMemberList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	CreateChat(context.Context, *CreateChatReq) (*CreateChatResp, error)
	RecentChats(context.Context, *RecentChatsReq) (*RecentChatsResp, error)
	JoinToChat(context.Context, *JoinToChatReq) (*EmptyResp, error)
	RemoveMemberFromChat(context.Context, *RemoveMemberFromChatReq) (*EmptyResp, error)
	GetChatMsgList(context.Context, *GetChatMsgListReq) (*GetChatMsgListResp, error)
	GetMemberList(context.Context, *GetMemberListReq) (*GetMemberListResp, error)
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) CreateChat(context.Context, *CreateChatReq) (*CreateChatResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedChatServer) RecentChats(context.Context, *RecentChatsReq) (*RecentChatsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecentChats not implemented")
}
func (UnimplementedChatServer) JoinToChat(context.Context, *JoinToChatReq) (*EmptyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinToChat not implemented")
}
func (UnimplementedChatServer) RemoveMemberFromChat(context.Context, *RemoveMemberFromChatReq) (*EmptyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMemberFromChat not implemented")
}
func (UnimplementedChatServer) GetChatMsgList(context.Context, *GetChatMsgListReq) (*GetChatMsgListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatMsgList not implemented")
}
func (UnimplementedChatServer) GetMemberList(context.Context, *GetMemberListReq) (*GetMemberListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMemberList not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_CreateChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).CreateChat(ctx, req.(*CreateChatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_RecentChats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecentChatsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).RecentChats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_RecentChats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).RecentChats(ctx, req.(*RecentChatsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_JoinToChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinToChatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).JoinToChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_JoinToChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).JoinToChat(ctx, req.(*JoinToChatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_RemoveMemberFromChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveMemberFromChatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).RemoveMemberFromChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_RemoveMemberFromChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).RemoveMemberFromChat(ctx, req.(*RemoveMemberFromChatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetChatMsgList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatMsgListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetChatMsgList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_GetChatMsgList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetChatMsgList(ctx, req.(*GetChatMsgListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetMemberList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemberListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetMemberList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_GetMemberList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetMemberList(ctx, req.(*GetMemberListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat_server.chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createChat",
			Handler:    _Chat_CreateChat_Handler,
		},
		{
			MethodName: "recentChats",
			Handler:    _Chat_RecentChats_Handler,
		},
		{
			MethodName: "joinToChat",
			Handler:    _Chat_JoinToChat_Handler,
		},
		{
			MethodName: "removeMemberFromChat",
			Handler:    _Chat_RemoveMemberFromChat_Handler,
		},
		{
			MethodName: "getChatMsgList",
			Handler:    _Chat_GetChatMsgList_Handler,
		},
		{
			MethodName: "getMemberList",
			Handler:    _Chat_GetMemberList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
