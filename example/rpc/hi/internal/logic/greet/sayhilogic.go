package greetlogic

import (
	"context"

	"github.com/allen-ping/goctl/v2/example/rpc/hi/internal/svc"
	"github.com/allen-ping/goctl/v2/example/rpc/hi/pb/hi"
	"github.com/zeromicro/go-zero/core/logx"
)

type SayHiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayHiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHiLogic {
	return &SayHiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SayHiLogic) SayHi(in *hi.HiReq) (*hi.HiResp, error) {
	// todo: add your logic here and delete this line

	return &hi.HiResp{}, nil
}
