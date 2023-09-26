package repository

import (
	"context"

	"github.com/SurajKadam7/session-service/model"
)

type Repository interface {
	Add(ctx context.Context, sessionInfo model.SessionInfo) (res model.SessionInfo, err error)
	Remove(ctx context.Context, userId int) (res int, err error)
	GetIp(ctx context.Context, userId int) (serverIp string, err error)
}
