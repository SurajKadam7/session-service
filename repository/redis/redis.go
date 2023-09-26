package redis

import (
	"context"
	"strconv"

	"github.com/SurajKadam7/session-service/model"
	"github.com/SurajKadam7/session-service/repository"
	"github.com/redis/go-redis/v9"
)

const (
	sessionKey string = "sessionKey"
)

type repo struct {
	client *redis.Client
}

func New(client *redis.Client) repository.Repository {
	return &repo{
		client: client,
	}
}

func (r *repo) Add(ctx context.Context, sessionInfo model.SessionInfo) (res model.SessionInfo, err error) {
	_, err = r.client.HSet(ctx, sessionKey, sessionInfo.UserId, sessionInfo.ServerIp).Result()
	if err != nil {
		return
	}
	res = sessionInfo
	return
}

func (r *repo) Remove(ctx context.Context, userId int) (res int, err error) {

	_, err = r.client.HDel(ctx, sessionKey, strconv.Itoa(userId)).Result()
	if err != nil {
		return
	}
	res = userId
	return
}

func (r *repo) GetIp(ctx context.Context, userId int) (serverIp string, err error) {
	serverIp, err = r.client.HGet(ctx, sessionKey, strconv.Itoa(userId)).Result()
	return
}
