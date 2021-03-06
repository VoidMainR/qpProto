// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: friendRoom/friendRoom.proto

/*
Package friendRoom is a generated protocol buffer package.

It is generated from these files:
	friendRoom/friendRoom.proto

It has these top-level messages:
	RoomMessage
	CreateRoomRequest
	CreateRoomResponse
	EnterRoomRequest
	RoomData
	GetPlayingRequest
	GetPlayingResponse
	RoomOption
*/
package friendRoom

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/VoidMainR/qpProto/room"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for FriendRoom service

type FriendRoomService interface {
	CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...client.CallOption) (*CreateRoomResponse, error)
	EnterRoom(ctx context.Context, in *EnterRoomRequest, opts ...client.CallOption) (*RoomData, error)
	GetPlaying(ctx context.Context, in *GetPlayingRequest, opts ...client.CallOption) (*GetPlayingResponse, error)
	Stream(ctx context.Context, opts ...client.CallOption) (FriendRoom_StreamService, error)
}

type friendRoomService struct {
	c    client.Client
	name string
}

func NewFriendRoomService(name string, c client.Client) FriendRoomService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "friendRoom"
	}
	return &friendRoomService{
		c:    c,
		name: name,
	}
}

func (c *friendRoomService) CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...client.CallOption) (*CreateRoomResponse, error) {
	req := c.c.NewRequest(c.name, "FriendRoom.CreateRoom", in)
	out := new(CreateRoomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendRoomService) EnterRoom(ctx context.Context, in *EnterRoomRequest, opts ...client.CallOption) (*RoomData, error) {
	req := c.c.NewRequest(c.name, "FriendRoom.EnterRoom", in)
	out := new(RoomData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendRoomService) GetPlaying(ctx context.Context, in *GetPlayingRequest, opts ...client.CallOption) (*GetPlayingResponse, error) {
	req := c.c.NewRequest(c.name, "FriendRoom.GetPlaying", in)
	out := new(GetPlayingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendRoomService) Stream(ctx context.Context, opts ...client.CallOption) (FriendRoom_StreamService, error) {
	req := c.c.NewRequest(c.name, "FriendRoom.Stream", &RoomMessage{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &friendRoomServiceStream{stream}, nil
}

type FriendRoom_StreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*RoomMessage) error
	Recv() (*RoomMessage, error)
}

type friendRoomServiceStream struct {
	stream client.Stream
}

func (x *friendRoomServiceStream) Close() error {
	return x.stream.Close()
}

func (x *friendRoomServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *friendRoomServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *friendRoomServiceStream) Send(m *RoomMessage) error {
	return x.stream.Send(m)
}

func (x *friendRoomServiceStream) Recv() (*RoomMessage, error) {
	m := new(RoomMessage)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for FriendRoom service

type FriendRoomHandler interface {
	CreateRoom(context.Context, *CreateRoomRequest, *CreateRoomResponse) error
	EnterRoom(context.Context, *EnterRoomRequest, *RoomData) error
	GetPlaying(context.Context, *GetPlayingRequest, *GetPlayingResponse) error
	Stream(context.Context, FriendRoom_StreamStream) error
}

func RegisterFriendRoomHandler(s server.Server, hdlr FriendRoomHandler, opts ...server.HandlerOption) error {
	type friendRoom interface {
		CreateRoom(ctx context.Context, in *CreateRoomRequest, out *CreateRoomResponse) error
		EnterRoom(ctx context.Context, in *EnterRoomRequest, out *RoomData) error
		GetPlaying(ctx context.Context, in *GetPlayingRequest, out *GetPlayingResponse) error
		Stream(ctx context.Context, stream server.Stream) error
	}
	type FriendRoom struct {
		friendRoom
	}
	h := &friendRoomHandler{hdlr}
	return s.Handle(s.NewHandler(&FriendRoom{h}, opts...))
}

type friendRoomHandler struct {
	FriendRoomHandler
}

func (h *friendRoomHandler) CreateRoom(ctx context.Context, in *CreateRoomRequest, out *CreateRoomResponse) error {
	return h.FriendRoomHandler.CreateRoom(ctx, in, out)
}

func (h *friendRoomHandler) EnterRoom(ctx context.Context, in *EnterRoomRequest, out *RoomData) error {
	return h.FriendRoomHandler.EnterRoom(ctx, in, out)
}

func (h *friendRoomHandler) GetPlaying(ctx context.Context, in *GetPlayingRequest, out *GetPlayingResponse) error {
	return h.FriendRoomHandler.GetPlaying(ctx, in, out)
}

func (h *friendRoomHandler) Stream(ctx context.Context, stream server.Stream) error {
	return h.FriendRoomHandler.Stream(ctx, &friendRoomStreamStream{stream})
}

type FriendRoom_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*RoomMessage) error
	Recv() (*RoomMessage, error)
}

type friendRoomStreamStream struct {
	stream server.Stream
}

func (x *friendRoomStreamStream) Close() error {
	return x.stream.Close()
}

func (x *friendRoomStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *friendRoomStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *friendRoomStreamStream) Send(m *RoomMessage) error {
	return x.stream.Send(m)
}

func (x *friendRoomStreamStream) Recv() (*RoomMessage, error) {
	m := new(RoomMessage)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
