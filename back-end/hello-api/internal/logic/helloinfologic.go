package logic

import (
	"context"

	"backend/hello-api/internal/svc"
	"backend/hello-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 你好，世界
func NewHelloInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloInfoLogic {
	return &HelloInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloInfoLogic) HelloInfo(req *types.HelloReq) (resp *types.HelloResp, err error) {
	return &types.HelloResp{
		HelloWorld: "hello " + req.World,
	}, nil
}
