package logic

import (
	"context"

	"one_for_all/services/account/rpc/internal/svc"
	"one_for_all/services/account/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByNameLogic {
	return &GetUserByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByNameLogic) GetUserByName(in *account.GetUserByNameReq) (*account.GetUserByNameRes, error) {
	// todo: add your logic here and delete this line

	return &account.GetUserByNameRes{}, nil
}
