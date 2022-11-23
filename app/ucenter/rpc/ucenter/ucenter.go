// Code generated by goctl. DO NOT EDIT!
// Source: ucenter.proto

package ucenter

import (
	"context"

	"gomicro/app/ucenter/rpc/types/ucenter"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	UserInfoReply = ucenter.UserInfoReply
	UserInfoReq   = ucenter.UserInfoReq

	Ucenter interface {
		UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoReply, error)
	}

	defaultUcenter struct {
		cli zrpc.Client
	}
)

func NewUcenter(cli zrpc.Client) Ucenter {
	return &defaultUcenter{
		cli: cli,
	}
}

func (m *defaultUcenter) UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoReply, error) {
	client := ucenter.NewUcenterClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}
