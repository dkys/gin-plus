package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

type User struct {
}

type Menu struct {
}

func (u *Menu) GetName(ctx *gin.Context) {
	ctx.String(http.StatusOK, "menu Name")
}

func (u *User) PostAddr(ctx *gin.Context) {
	ctx.String(http.StatusOK, "user addr")
}

func (u *User) GetIndex(ctx *gin.Context) {
	ctx.String(http.StatusOK, "user index")
}

func (u *User) test(ctx *gin.Context) {
	ctx.String(http.StatusOK, "my name is luxi")
}

func TestGin(t *testing.T) {
	gin.SetMode(gin.DebugMode)
	r := Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello")
	})
	//routers := r.Routes()
	//for _, router := range routers {
	//	elog.Debug(router.Method, router.Path, router.HandlerFunc, router.Handler)
	//}
	r.Handles("/user", new(User), new(Menu))

	r.Run()
}
