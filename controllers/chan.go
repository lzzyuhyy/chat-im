package controllers

import (
	"chat-im/global"
	"chat-im/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type ChanController struct {
	beego.Controller
}

func (c *ChanController) Chan() {
	// 模拟中间件鉴权
	userId, _ := c.GetInt("user_id")
	log.Println(userId)
	auth := true
	if userId == 0 {
		auth = false
	}

	//将http请求升级成websocket请求
	// 链接配置信息
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024, // 读取数据缓冲区大小
		WriteBufferSize: 1024, // 写入数据缓冲区大小
		CheckOrigin: func(r *http.Request) bool { // 校验源---是否允许请求链接
			return auth
		},
	}

	// 建立链接  conn---链接信息/标识
	conn, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// 初始化映射关系
	node := global.Node{
		Conn: conn,
		Data: make(chan []byte, 50), // 初始化数据通道
		//SetGroup: set.New(set.ThreadSafe),
	}
	global.ClientMap[userId] = node

	//// 获取用户所在的所有群
	//ug := models.GroupChat{
	//	OwnerId: uint(userId),
	//}
	//list, res := ug.GetGroupList()
	//if res.Error != nil {
	//	logs.Error("用户所在群获取失败", res.Error)
	//}
	//for _, v := range list {
	//	node.SetGroup.Add(v.ID) // 将用户所在群id放入集合
	//}

	//当我们想要同时发送和接收消息时--协程
	// 同步等待组
	global.WG.Add(2)
	go ReadMessage(node)  // 读-我向服务端发送消息，服务端读取
	go WriteMessage(node) // 读-我向服务端发送消息，服务端读取
	global.WG.Wait()

}

// 消息内容结构体
type MessageType struct {
	UserId  uint   `json:"user_id"`  // 发布方id
	DistId  uint   `json:"dist_id"`  // 接收方id
	Content string `json:"content"`  // 消息内容
	MsgType int    `json:"msg_type"` // 消息类型 （系统/个人）
	Cmd     int    `json:"cmd"`      // 消息类型 （群聊/私聊）
}

// 发消息---我向服务端发送数据
func ReadMessage(node global.Node) {
	defer global.WG.Done()
	var msg MessageType
	for {
		// 我发的消息，服务器来读
		_, message, _ := node.Conn.ReadMessage()
		// 读到消息后转码
		json.Unmarshal(message, &msg)

		//是否群聊
		//if msg.Cmd == 1 { // 是群聊
		//	// 找到当前用户聊天的群
		//
		//} else if msg.Cmd == 0 { // 不是群聊
		// 对方链接信息是否存在
		if _, ok := global.ClientMap[int(msg.DistId)]; ok {
			// 信息存在, 将消息后放入对方接信息的对应数据通道中
			global.ClientMap[int(msg.DistId)].Data <- message
		} else {
			msg = MessageType{
				UserId:  msg.UserId,
				DistId:  msg.DistId,
				Content: "对方不在线",
				MsgType: 2,
				Cmd:     msg.Cmd,
			}
			str, _ := json.Marshal(&msg)
			// 不存在--给我发送系统级消息--对方不在线
			global.ClientMap[int(msg.UserId)].Data <- str
		}
		//}
	}
}

// 收消息---服务端向我的客户端写入
func WriteMessage(node global.Node) {
	defer global.WG.Done()
	var msg MessageType
	for {
		// 多路复用
		select {
		case message, ok := <-node.Data: // 从节点获取数据---如果没有数据会阻塞
			if ok {
				// 转码
				json.Unmarshal(message, &msg)
				err := node.Conn.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					// 发送失败
					msg = MessageType{
						UserId:  msg.UserId,
						DistId:  msg.DistId,
						Content: msg.Content,
						MsgType: msg.MsgType,
						Cmd:     msg.Cmd,
					}
					str, _ := json.Marshal(&msg)
					//对发消息的用户做出提示
					global.ClientMap[int(msg.UserId)].Data <- str
					// 将发送失败或未发送（对方不在线）的消息也存入数据库
					msgHis := models.MsgHistories{
						OwnerId: msg.UserId,
						DistId:  msg.DistId,
						Content: msg.Content,
						IsRead:  0,
						IsSend:  2,
						Cmd:     uint8(msg.Cmd),
						Status:  0,
					}
					err = msgHis.AddMessage()
					if err != nil {
						fmt.Println("数据添加成功")
						return
					}
				}
				// 发送成功
				msgHis := models.MsgHistories{
					OwnerId: msg.UserId,
					DistId:  msg.DistId,
					Content: msg.Content,
					IsRead:  0,
					IsSend:  2,
					Cmd:     uint8(msg.Cmd),
					Status:  0,
				}
				err = msgHis.AddMessage()
				if err != nil {
					fmt.Println("数据添加成功")
					return
				}
			} else {
				// 无数据情况处理
			}
		}
	}
}
