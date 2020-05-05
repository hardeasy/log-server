package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"log-server/config"
	v1 "log-server/internal/server/http/controler/api/v1"
	pushv1 "log-server/internal/server/http/controler/push/v1"
	"log-server/internal/server/http/middleware"
	"log-server/internal/service"
	"net/http"
	"os"
	"time"
)

func NewHttpServer(cfg *config.Config, serv *service.Service) *HttpServer {
	hs := &HttpServer{cfg, serv, nil}
	hs.Run()
	return hs
}

type HttpServer struct {
	Conf *config.Config
	Service *service.Service
	InnerServer *http.Server
}

func (this *HttpServer) Run() {
	apiController := v1.NewController(this.Service)
	pushController := pushv1.NewController(this.Service)
	mdw := middleware.NewMiddleWare(this.Service)

	var r *gin.Engine
	if this.Conf.Application.Env != config.Application_DEV {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
		f, _ := os.Create(this.Conf.Application.LogPath)
		gin.DefaultWriter = io.MultiWriter(f)
	}

	r = gin.Default()
	r.Use(mdw.CorsFilter)
	api := r.Group("/api/v1")
	api.Use(mdw.AuthUserFilter)
	{
		api.POST("/login", apiController.Login)

		api.POST("/logout", apiController.Logout)

		api.GET("/logs", apiController.GetLogList)

		api.GET("/logs/:id", apiController.GetLogDetail)

		api.GET("/indices", apiController.Getindices)
		api.DELETE("/indices/:index", apiController.DeleteIndex)

		api.GET("/users/me", apiController.GetUserById)

		api.GET("/users", apiController.GetUserList)
		api.POST("/users", apiController.AddUser)

		api.PUT("/users/:id", apiController.EditUser)

		api.GET("/apps", apiController.GetAppAll)
		api.POST("/apps", apiController.AddApp)
		api.PUT("/apps", apiController.EditApp)

		api.GET("/apps/:appid/members", apiController.GetAppMembers)
		api.POST("/apps/:appid/members", apiController.AddAppMember)
		api.DELETE("/apps/:appid/members/:userid", apiController.DeleteAppMember)
	}

	push := r.Group("/push/v1")
	push.Use(mdw.PushAuthFilter)
	{
		push.POST("/logs", pushController.Push)
	}

	this.InnerServer = &http.Server{Addr:this.Conf.HTTPServer.Addr, Handler: r}
	go func() {
		if err := this.InnerServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen error\n", err)
		}
	}()

}

func (this *HttpServer) Shutdown() {
	log.Println("http server shutdown....")
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := this.InnerServer.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown error: ", err)
	}
	log.Println("http server exiting...")
}
