package httpserver

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"ip-pool/pkg/global"
	"ip-pool/pkg/v2ray"
	"ip-pool/pkg/zipfile"
	"log"
	"os"
	"strings"
)

func StartServer() {
	g := gin.Default()

	api := g.Group("/api")
	api.Use(func(ctx *gin.Context) {
		if ctx.Query("token") != "ai" {
			ctx.AbortWithStatus(403)
			return
		}
		ctx.Next()
	})
	api.GET("/register", register)
	api.POST("/del", DelUser)
	g.GET("/link/:username", queryUUID)

	g.GET("/download/:username", download)

	g.Run(":9123")
}

func register(ctx *gin.Context) {
	username := ctx.Query("username")
	if username == "" {
		ctx.AbortWithStatus(403)
		return
	}
	id, err := v2ray.CreateUser(ctx.Request.Context(), global.UserTag, username)
	if err != nil {
		log.Println("[v2ray.CreateUser] err", err)
		ctx.Error(err)
		return
	}
	uri := fmt.Sprintf(data, id)
	vmess := fmt.Sprintf("vmess://%s", base64.StdEncoding.EncodeToString([]byte(uri)))
	ctx.String(200, vmess)
}

func download(ctx *gin.Context) {
	user := strings.TrimPrefix(ctx.Param("username"), "/")
	if user == "" {
		ctx.String(200, "参数错误")
		return
	}
	uuid, err := v2ray.QueryUUID(user)
	if err != nil || uuid == "" {
		ctx.String(200, "用户不存在")
		return
	}
	dst, err := zipfile.ZipFile("/tmp/data/v2ray", uuid)
	if err != nil {
		ctx.String(200, "服务器内部错误")
		return
	}
	defer os.Remove(dst)
	// 设置响应头
	ctx.Header("Content-Disposition", "attachment; filename=v2rayN.zip")
	ctx.Header("Content-Type", "application/zip")
	ctx.File(dst)
}

func DelUser(ctx *gin.Context) {
	username := ctx.Query("username")
	if username == "" {
		ctx.AbortWithStatus(403)
		return
	}
	err := v2ray.DelUser(ctx.Request.Context(), username)
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
	vmess := fmt.Sprintf("vmess://%s", base64.StdEncoding.EncodeToString([]byte(uri)))
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
