package controllers

import (
	"chat-im/global"
	"encoding/json"
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
	userId := c.GetString("user_id")
	auth := true
	if userId == "" {
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
	}
	global.ClientMap[userId] = node

	//当我们想要同时发送和接收消息时--协程
	// 同步等待组
	global.WG.Add(2)
	go ReadMessage(node)  // 读-我向服务端发送消息，服务端读取
	go WriteMessage(node) // 读-我向服务端发送消息，服务端读取
	global.WG.Wait()

}

// 消息内容结构体
type MessageType struct {
	UserId  string `json:"user_id"`  // 发布方id
	DistId  string `json:"dist_id"`  // 接收方id
	Content string `json:"content"`  // 消息内容
	MsgType string `json:"msg_type"` // 消息类型 （系统/个人）
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
		// 对方链接信息是否存在
		if _, ok := global.ClientMap[msg.DistId]; ok {
			// 信息存在, 将消息后放入对方接信息的对应数据通道中
			global.ClientMap[msg.DistId].Data <- message
		} else {
			msg = MessageType{
				UserId:  msg.UserId,
				DistId:  msg.DistId,
				Content: "对方不在线",
				MsgType: "admin",
			}
			str, _ := json.Marshal(&msg)
			// 不存在--给我发送系统级消息--对方不在线
			global.ClientMap[msg.UserId].Data <- str
		}
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
					}
					str, _ := json.Marshal(&msg)
					//对发消息的用户做出提示
					global.ClientMap[msg.UserId].Data <- str
					// 将发送失败或未发送（对方不在线）的消息也存入数据库
				}
			} else {
				// 无数据情况处理
			}
		}
	}
}
