package v2ray

/*package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"v2ray.com/core/app/proxyman/command"
	statsService "v2ray.com/core/app/stats/command"
	"v2ray.com/core/common/protocol"
	"v2ray.com/core/common/serial"
	"v2ray.com/core/proxy/vmess"
)

const (
	API_ADDRESS = "195.245.242.102"
	API_PORT    = 10080
	INBOUND_TAG = "proxy"
	LEVEL       = 0
	EMAIL       = "123@gmail.com"
	UUID        = "2601070b-ab53-4352-a290-1d44414581ae"
	ALTERID     = 32
)

func main() {
	cmdConn, err := grpc.Dial(fmt.Sprintf("%s:%d", API_ADDRESS, API_PORT), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	hsClient := command.NewHandlerServiceClient(cmdConn)
	addUser(hsClient)
}

func addUser(c command.HandlerServiceClient) {
	resp, err := c.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag: INBOUND_TAG,
		Operation: serial.ToTypedMessage(&command.AddUserOperation{
			User: &protocol.User{
				Level: LEVEL,
				Account: serial.ToTypedMessage(&vmess.Account{
					Id:               UUID,
					AlterId:          ALTERID,
					SecuritySettings: &protocol.SecurityConfig{Type: protocol.SecurityType_AUTO},
				}),
			},
		}),
	})
	if err != nil {
		log.Printf("failed to call grpc command: %v", err)
	} else {
		log.Printf("ok: %v", resp)
	}
}
func removeUser(c command.HandlerServiceClient) {
	resp, err := c.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag: INBOUND_TAG,
		Operation: serial.ToTypedMessage(&command.RemoveUserOperation{
			Email: EMAIL,
		}),
	})
	if err != nil {
		log.Printf("failed to call grpc command: %v", err)
	} else {
		log.Printf("ok: %v", resp)
	}
}
func queryUserTraffic(c statsService.StatsServiceClient) {
	resp, err := c.QueryStats(context.Background(), &statsService.QueryStatsRequest{
		Pattern: "*",   // 筛选用户表达式
		Reset_:  false, // 查询完成后是否重置流量
	})
	if err != nil {
		log.Printf("failed to call grpc command: %v", err)
	}
	// 获取返回值中的流量信息
	stat := resp.GetStat()
	// 返回的是一个数组，对其进行遍历输出
	for _, e := range stat {
		fmt.Println(e)
	}
}
*/
