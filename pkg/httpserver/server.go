package httpserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ip-pool/pkg/global"
	"ip-pool/pkg/v2ray"
	"log"
	"strings"
)

func StartServer() {
	g := gin.Default()

	api := g.Group("/api")
	api.Use(func(ctx *gin.Context) {
		if ctx.GetHeader("Authorization") != "hiai" {
			ctx.AbortWithStatus(403)
			return
		}
		ctx.Next()
	})
	api.POST("/register", register)
	api.POST("/del", DelUser)
	g.GET("/link/:username", queryUUID)

	g.Run(":9123")
}

type RegisterRequest struct {
	Username string `json:"username"`
}

func register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.BindJSON(&req); err != nil {
		log.Println("[bind] err 400", err)
		ctx.Error(err)
		return
	}
	id, err := v2ray.CreateUser(ctx.Request.Context(), global.UserTag, req.Username)
	if err != nil {
		log.Println("[v2ray.CreateUser] err", err)
		ctx.Error(err)
		return
	}
	ctx.JSON(200, gin.H{"status": 0, "id": id})
}

func DelUser(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.BindJSON(&req); err != nil {
		log.Println("[bind] err 400", err)
		ctx.Error(err)
		return
	}
	err := v2ray.DelUser(ctx.Request.Context(), req.Username)
	if err != nil {
		log.Println("[v2ray.delUser] err", err)
		ctx.Error(err)
		return
	}
	ctx.JSON(200, gin.H{"status": 0})
}

func queryUUID(ctx *gin.Context) {
	user := strings.TrimPrefix(ctx.Param("username"), "/")
	uuid, err := v2ray.QueryUUID(user)
	if err != nil || uuid == "" {
		ctx.String(200, "用户不存在")
		return
	}
	uri := fmt.Sprintf(data, uuid)
	vmess := fmt.Sprintf("vmess://%s", uri)
	ctx.String(200, vmess)
}

var data = `{
  "v": "2",
  "ps": "hi-ai.top",
  "add": "test.hi-ai.top",
  "port": "443",
  "id": "%s",
  "aid": "0",
  "net": "ws",
  "type": "none",
  "host": "test.hi-ai.top",
  "path": "/puppet",
  "tls": "tls"
}
`
