# gin-plus

#### 介绍
**gin-plus** 对gin框架的二次包装，使其支持路由，解析结构体方法。
#### 使用方法 
```go 
package main

import (
	"github.com/dkys/elog"
	"github.com/dkys/gin-plus"
	"github.com/gin-gonic/gin"
	"net/http"
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

func main() {
	gin.SetMode(gin.DebugMode)
	r := gplus.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello")
	})
	elog.Error()
	r.Handles("/user", new(User), new(Menu))

	r.Run()
}
```