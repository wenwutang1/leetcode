package main
import (
	"bytes"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)
const (
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)
var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: true,
}

type Client struct {
	//webSocket连接器
	hub *Hub
	conn *websocket.Conn
	//广播的消息
	send chan []byte
}

//read message
func (this *Client) ReadMsg(){
	//失败收关闭连接删除通道
	defer func() {
		//连接错误的通道
		this.hub.unregister <- this
		this.conn.Close()
	}()
	this.conn.SetReadLimit(maxMessageSize)
	//设置连接超时时间
	this.conn.SetReadDeadline(time.Now().Add(pongWait))
	this.conn.SetPongHandler(func(appData string) error {
		this.conn.SetReadDeadline(time.Now().Add(pongWait));return nil
	})
	for {
		_,message,err:=this.conn.ReadMessage()
		if err!=nil{
			log.Println(err.Error())
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		//把消息传给接收器
		this.hub.broadcast <- message
	}
}

//send message
func (this *Client)WriteMsg(){
	ticker:=time.NewTicker(pingPeriod)//定时器
	defer func() {
		ticker.Stop()
		this.conn.Close()
	}()
	for {
		select {
		case message,ok:=<-this.send:
			this.conn.SetReadDeadline(time.Now().Add(writeWait))
			if !ok{
				// The hub closed the channel.
				this.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			//UTF-8的格式广播数据
			writer,err:=this.conn.NextWriter(websocket.TextMessage)
			if err!=nil{
				return
			}
			//第一个数据写入连接中
			writer.Write(message)
			//写入后续的数据 中间加空格页面打印清晰一点
			n:=len(this.send)
			for i:=0;i<n;i++{
				writer.Write(newline)
				writer.Write(<-this.send)
			}
			//写完关闭
			if err := writer.Close();err!=nil{
				return
			}
		//超时
		case <-ticker.C:
			this.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := this.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	//web中兼容http连接从捂手状态升级协议
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WriteMsg()
	go client.ReadMsg()
}