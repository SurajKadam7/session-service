package sessions_svc

import (
	"context"
	"errors"
	"fmt"

	msgModel "github.com/SurajKadam7/msg-info-service/model"
	msgSrv "github.com/SurajKadam7/msg-info-service/msginfo_srv"
	"github.com/SurajKadam7/session-service/model"
	"github.com/SurajKadam7/session-service/repository"
)

// @microgen http, middleware, logging
type Service interface {
	// @http-method POST
	// @http-path add/
	Add(ctx context.Context, sessionInfo model.SessionInfo) (res model.SessionInfo, err error)
	// @http-method POST
	// @http-path sendmsg/
	SendMsg(ctx context.Context, userId int, message model.Message) (status bool, err error)
	// @http-method POST
	// @http-path sendgroup/
	SendGroupMsg(ctx context.Context, userId int, groupId int, message model.Message) (status bool, err error)
	// @http-method POST
	// @http-path remove/
	Remove(ctx context.Context, userId int) (res model.SessionInfo, err error)
}

type service struct {
	repo          repository.Repository
	msgInfoClient msgSrv.Service
}

func New(repo repository.Repository, msgInfoClient msgSrv.Service) Service {
	return &service{
		repo:          repo,
		msgInfoClient: msgInfoClient,
	}
}

// @http-method POST
// @http-path add/
func (s *service) Add(ctx context.Context, sessionInfo model.SessionInfo) (res model.SessionInfo, err error) {
	if sessionInfo.ServerIp != "" || sessionInfo.UserId != 0 {
		return res, errors.New("invalid request")
	}

	res, err = s.repo.Add(ctx, sessionInfo)
	return
}

// @http-method POST
// @http-path sendmsg/
func (s *service) SendMsg(ctx context.Context, userId int, message model.Message) (status bool, err error) {
	if userId == 0 {
		return status, errors.New("invalid request")
	}
	if message.SentTo == 0 || message.SentBy == 0 {
		return status, errors.New("invalid request")
	}

	msgInfo := msgModel.MsgInfo{
		From:   message.SentBy,
		To:     message.SentBy,
		Msg:    message.Data,
		Status: msgModel.Sent,
	}
	
	msgId, err := s.msgInfoClient.Add(ctx, msgInfo)
	if err != nil {
		// TODO
		return
	}

	serverIp, err := s.repo.GetIp(ctx, userId)
	if err != nil {
		return status, errors.New("serverIp not found")
	}
	fmt.Println(serverIp)
	// TODO sending request to the msg-storage & api-gateway is remaining

	err = s.msgInfoClient.Update(ctx, message.SentBy, msgId, msgModel.Delivered)
	if err != nil {
		// TODO
		return
	}

	return
}

// @http-method POST
// @http-path sendgroup/
func (s *service) SendGroupMsg(ctx context.Context, userId int, groupId int, message model.Message) (status bool, err error) {
	panic("not implemented") // TODO: Implement
}

// @http-method POST
// @http-path remove/
func (s *service) Remove(ctx context.Context, userId int) (res model.SessionInfo, err error) {
	if userId != 0 {
		return res, errors.New("invalid request")
	}
	res, err = s.Remove(ctx, userId)
	return
}
