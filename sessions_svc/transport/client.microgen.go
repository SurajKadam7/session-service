// Code generated by microgen 1.0.5. DO NOT EDIT.

package transport

import (
	"context"
	model "github.com/SurajKadam7/session-service/model"
)

func (set EndpointsSet) Add(arg0 context.Context, arg1 model.SessionInfo) (res0 model.SessionInfo, res1 error) {
	request := AddRequest{SessionInfo: arg1}
	response, res1 := set.AddEndpoint(arg0, &request)
	if res1 != nil {
		return
	}
	return response.(*AddResponse).Res, res1
}

func (set EndpointsSet) SendMsg(arg0 context.Context, arg1 int, arg2 model.Message) (res0 bool, res1 error) {
	request := SendMsgRequest{
		Message: arg2,
		UserId:  arg1,
	}
	response, res1 := set.SendMsgEndpoint(arg0, &request)
	if res1 != nil {
		return
	}
	return response.(*SendMsgResponse).Status, res1
}

func (set EndpointsSet) SendGroupMsg(arg0 context.Context, arg1 int, arg2 int, arg3 model.Message) (res0 bool, res1 error) {
	request := SendGroupMsgRequest{
		GroupId: arg2,
		Message: arg3,
		UserId:  arg1,
	}
	response, res1 := set.SendGroupMsgEndpoint(arg0, &request)
	if res1 != nil {
		return
	}
	return response.(*SendGroupMsgResponse).Status, res1
}

func (set EndpointsSet) Remove(arg0 context.Context, arg1 int) (res0 model.SessionInfo, res1 error) {
	request := RemoveRequest{UserId: arg1}
	response, res1 := set.RemoveEndpoint(arg0, &request)
	if res1 != nil {
		return
	}
	return response.(*RemoveResponse).Res, res1
}