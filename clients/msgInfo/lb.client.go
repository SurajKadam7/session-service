package msginfo

import (
	"log"
	"net/url"

	msgSrv "github.com/SurajKadam7/msg-info-service/msginfo_srv"
	msgInfoHttp "github.com/SurajKadam7/msg-info-service/msginfo_srv/transport/http"
)

func GetClient(rawURL string) (client msgSrv.Service) {
	url, err := url.Parse(rawURL)
	if err != nil {
		log.Fatal("Error come while parsing msgInfo url", err)
	}
	client = msgInfoHttp.NewHTTPClient(url)
	return
}
