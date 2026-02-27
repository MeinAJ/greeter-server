package logic

import (
	"context"

	"github.com/MeinAJ/greeter-server/greeter"
	"github.com/MeinAJ/greeter-server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SayHelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHelloLogic {
	return &SayHelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SayHelloLogic) SayHello(in *greeter.SayHelloReq) (*greeter.SayHelloResp, error) {
	// todo: add your logic here and delete this line
	return &greeter.SayHelloResp{Message: "Hello " + in.Name}, nil
}
