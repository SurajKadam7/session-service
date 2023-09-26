package apigateway

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/SurajKadam7/msg-info-service/model"
	msgSrv "github.com/SurajKadam7/msg-info-service/msginfo_srv"
	msgInfoHttp "github.com/SurajKadam7/msg-info-service/msginfo_srv/transport/http"
)

// below code is for static load balancing No use for msginfoclient
type lbClient struct {
	port int
}

func New(port int) msgSrv.Service {
	return lbClient{
		port: port,
	}
}

type ip string

const Key ip = "ip"

func (l lbClient) getHost(ip string) string {
	return fmt.Sprintf("%s:%d", ip, l.port)
}

func getIp(ctx context.Context) (ip string, err error) {
	ip, ok := ctx.Value(Key).(string)
	if !ok {
		err = errors.New("ip not found")
	}
	return
}

func (l lbClient) getclient(host string) (client msgSrv.Service) {
	url := url.URL{
		Scheme: "http",
		Host:   host,
	}
	client = msgInfoHttp.NewHTTPClient(&url)
	return
}

func (lbm lbClient) Add(ctx context.Context, msg model.MsgInfo) (id int, err error) {
	ip, err := getIp(ctx)
	if err != nil {
		return
	}
	client := lbm.getclient(ip)
	return client.Add(ctx, msg)
}

func (lbm lbClient) Delete(ctx context.Context, userId int, msgId int) (err error) {
	ip, err := getIp(ctx)
	if err != nil {
		return
	}
	client := lbm.getclient(ip)
	return client.Delete(ctx, userId, msgId)
}

func (lbm lbClient) GetAll(ctx context.Context, userId int, status model.Status) (msgs []model.MsgInfo, err error) {
	ip, err := getIp(ctx)
	if err != nil {
		return
	}
	client := lbm.getclient(ip)
	return client.GetAll(ctx, userId, status)
}

func (lbm lbClient) Update(ctx context.Context, userId int, msgId int, status model.Status) (err error) {
	ip, err := getIp(ctx)
	if err != nil {
		return
	}
	client := lbm.getclient(ip)
	return client.Update(ctx, userId, msgId, status)
}
