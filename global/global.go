package global

import (
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"sync"
)

// 全局创建映射所有人的映射关系连接信息，方便使用--避免包的循环引用

// 全局链接信息
type Node struct {
	Conn       *websocket.Conn // 链接信息
	Data       chan []byte     // 收发的数据
	CloseRead  bool            // 是否关闭发送消息的链接
	CloseWrite bool            // 是否关闭接收消息的链接
}

// 所有链接映射关系
var ClientMap map[string]Node = make(map[string]Node)

// 等待组--等待读写协程
var WG sync.WaitGroup

// 数据库
var DB *gorm.DB

// 返回数据类型
type ResponseData struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
