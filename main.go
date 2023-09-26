package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	msginfo "github.com/SurajKadam7/session-service/clients/msgInfo"
	redisRepo "github.com/SurajKadam7/session-service/repository/redis"
	"github.com/SurajKadam7/session-service/sessions_svc"
	"github.com/SurajKadam7/session-service/sessions_svc/service"
	"github.com/SurajKadam7/session-service/sessions_svc/transport"
	transporthttp "github.com/SurajKadam7/session-service/sessions_svc/transport/http"
	"github.com/go-kit/log"
	"github.com/redis/go-redis/v9"
)

const msgInfoUrl string = "http://localhost:8081/msginfoservice"

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		PoolSize: 100,
	})

	logger := log.With(log.NewJSONLogger(os.Stdout), "service", "sessions_srv")

	repo := redisRepo.New(client)
	msgClient := msginfo.GetClient(msgInfoUrl)

	srv := sessions_svc.New(repo, msgClient)
	srv = service.LoggingMiddleware(logger)(srv)
	endpoints := transport.Endpoints(srv)
	handler := transporthttp.NewHTTPHandler(&endpoints)

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	// grasefull shutdown

	quite := make(chan os.Signal, 1)
	signal.Notify(quite, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := httpServer.ListenAndServe()
		// error will never be nil here ...
		logger.Log("server error ", err.Error())
		close(quite)
	}()

	sig := <-quite

	if sig == syscall.SIGINT || sig == syscall.SIGTERM {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()

		logger.Log("server status", "shutdown called")
		time.Sleep(time.Second * 5)

		err := httpServer.Shutdown(ctx)
		if err != nil {
			logger.Log("error while shutdown", err)
		}
	}

	logger.Log("server status", "closed")

}

// g := group.Group{}

// g.Add(httpServer.ListenAndServe, func(err error) {
// 	logger.Log("server error", err)
// 	err1 := httpServer.Shutdown(context.Background())
// 	logger.Log("server shut down", err1)
// })

// quite := make(chan os.Signal, 1)

// g.Add(func() error {
// 	signal.Notify(quite, syscall.SIGINT, syscall.SIGTERM)
// 	sig := <-quite
// 	fmt.Println("channel closed .. ")
// 	return fmt.Errorf("recived close signal %s", sig)
// }, func(err error) {
// 	close(quite)
// })

// intrupt := make(chan struct{})
// g.Add(func() error {
// 	// time.AfterFunc(time.Second*10, func() {
// 	// 	close(intrupt)
// 	// })
// 	<-intrupt
// 	fmt.Println("third function .. ")
// 	return fmt.Errorf("function3 error ")
// }, func(err error) {
// 	fmt.Println("one more time ...............")
// 	select {
// 	case <-intrupt:
// 		return
// 	default:
// 		close(intrupt)
// 	}
// })

// g.Run()
