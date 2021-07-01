package main

import (
	"flag"
	"log"
	"net/http"
)

type Hub struct {
	client map[*Client]bool
	broadcast chan []byte
	register chan *Client
	unregister chan *Client
}
//初始化
func HubInit()*Hub{
	return &Hub{
		client: make(map[*Client]bool,0),
		broadcast:make(chan []byte,0),
		register:make(chan *Client,0),
		unregister:make(chan *Client,0),
	}
}
//从
func (this *Hub)Run(){
	for {
		select {
		case message:=<-this.broadcast:
			for c:=range this.client{
				select {
				case c.send<-message:
					log.Println("广播消息")
				default:
					close(c.send)
					delete(this.client,c)
				}
			}
		}
	}
}

//////////////////////////////////////////////////////////////////////////////////

var addr =flag.String("addr","127.0.0.1:8080","服务器地址和端口号")
//页面跳转
func serverHome(w http.ResponseWriter,r *http.Request){
	log.Println(r.URL)
	if r.URL.Path!="/"{
		http.Error(w,"路径不存在",404)
		return
	}
	if r.Method!="GET"{
		http.Error(w,"request method error",http.StatusNotAcceptable)
		return
	}
	http.ServeFile(w,r,"home.html")
}

func main(){
	hub := HubInit()
	go hub.Run()
	http.HandleFunc("/",serverHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub,w,r)
	})
	err:=http.ListenAndServe(*addr,nil)
	if err!=nil{
		log.Println(err.Error())
		return
	}
}
