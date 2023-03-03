package v2ray

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"ip-pool/pkg/global"
	"log"
	"strings"
	"sync"
	"v2ray.com/core/app/proxyman/command"
	"v2ray.com/core/common/protocol"
	"v2ray.com/core/common/serial"
	"v2ray.com/core/proxy/vmess"
)

var (
	o      sync.Once
	client command.HandlerServiceClient
)

func InitClient(addr string) {
	o.Do(func() {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		client = command.NewHandlerServiceClient(conn)
	})
}

func CreateUser(ctx context.Context, tag string, username string) (string, error) {
	uid := uuid.New().String()
	err := createUserGRPC(username, uid, tag)
	if err != nil {
		log.Println("[CreateUser] 创建用户失败: ", username, uid)
		return "", err
	}
	if err := createUserDB(username, uid); err != nil {
		return "", err
	}
	log.Println("[CreateUser] 创建用户成功: ", username, uid)
	return uid, nil
}

func emailByUser(e string) string {
	if strings.Contains(e, "@") {
		return e
	}
	return fmt.Sprintf("%s@gmail.com", e)
}

func createUserGRPC(user string, uid string, tag string) error {
	_, err := client.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag: tag,
		Operation: serial.ToTypedMessage(&command.AddUserOperation{
			User: &protocol.User{
				Level: 0,
				Email: emailByUser(user),
				Account: serial.ToTypedMessage(&vmess.Account{
					Id:               uid,
					AlterId:          1,
					SecuritySettings: &protocol.SecurityConfig{Type: protocol.SecurityType_AUTO},
				}),
			},
		}),
	})
	return err
}

func DelUser(ctx context.Context, username string) error {
	uuid, err := QueryUUID(username)
	if err != nil {
		return err
	}
	resp, err := client.AlterInbound(ctx, &command.AlterInboundRequest{
		Tag: global.UserTag,
		Operation: serial.ToTypedMessage(&command.RemoveUserOperation{
			Email: emailByUser(username),
		}),
	})
	if err != nil {
		log.Println("[delUser] 删除用户失败: ", username, err, uuid)
		return err
	}
	_ = resp
	log.Println("[delUser] 删除用户成功: ", username)
	return nil
}

// InitUsers 将用户注册到服务器
func InitUsers() error {
	err := db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte("users"))
		if bkt == nil {
			return nil
		}
		var (
			key []byte
			val []byte
		)
		c := bkt.Cursor()
		key, val = c.First()
		for {
			if len(key) == 0 || len(val) == 0 {
				break
			}
			err := createUserGRPC(string(key), string(val), global.UserTag)
			if err != nil {
				log.Println("[InitUser] 初始化用户失败: ", string(key), string(val))
			} else {
				log.Println("[InitUser] 初始化用户成功: ", string(key), string(val))
			}
			key, val = c.Next()
		}
		return nil
	})
	return err
}
