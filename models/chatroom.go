package models

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type ChatMessage struct {
	UserName  string
	Content   string
	TimeStamp int64
}

func NewChatMessage(username string, content string, timeStamp int64) *ChatMessage {
	chatMessage := ChatMessage{}
	chatMessage.UserName = username
	chatMessage.Content = content
	chatMessage.TimeStamp = timeStamp
	return &chatMessage
}

type ChatUser struct {
	UserName string
	Conn     *websocket.Conn
}

func NewChatUser(username string, conn *websocket.Conn) *ChatUser {
	chatUser := ChatUser{}
	chatUser.UserName = username
	chatUser.Conn = conn
	return &chatUser
}

type ChatRoom struct {
	ChatUserMap map[string]ChatUser
	Joiner      chan *ChatUser
	Leaver      chan string
	Messager    chan *ChatMessage
}

func NewChatRoom() *ChatRoom {
	chatRoom := ChatRoom{}
	chatRoom.ChatUserMap = make(map[string]ChatUser)
	chatRoom.Joiner = make(chan *ChatUser, 10)
	chatRoom.Leaver = make(chan string, 10)
	chatRoom.Messager = make(chan *ChatMessage, 10)
	go chatRoom.chatRoom()
	return &chatRoom
}

func (this *ChatRoom) SendMessage(username string, content string, timeStamp int64) {
	this.Messager <- NewChatMessage(username, content, timeStamp)
}

func (this *ChatRoom) Join(username string, ws *websocket.Conn) {
	fmt.Println(username + "join")
	this.Joiner <- NewChatUser(username, ws)
}

func (this *ChatRoom) Leave(username string) {
	this.Leaver <- username
}

func (this *ChatRoom) broadcast(chatMessage *ChatMessage) {
	for k, v := range this.ChatUserMap {
		if k == chatMessage.UserName {
			// var responseContent []interface{}
			// responseContent = append(responseContent, 1)
			// responseContent = append(responseContent, *chatMessage)
			// v.Conn.WriteJSON(responseContent)
			// fmt.Printf("%+v\n", responseContent)
			// fmt.Println(len(this.ChatUserMap))
			continue
		}
		var responseContent []interface{}
		responseContent = append(responseContent, 2)
		responseContent = append(responseContent, *chatMessage)
		fmt.Printf("%+v\n", responseContent)
		fmt.Println(len(this.ChatUserMap))
		v.Conn.WriteJSON(responseContent)

	}
}
func (this *ChatRoom) chatRoom() {
	for {
		select {
		case joiner := <-this.Joiner:
			this.ChatUserMap[joiner.UserName] = *joiner
		case leaver := <-this.Leaver:
			fmt.Println(leaver, "connect close")
			delete(this.ChatUserMap, leaver)
			if len(this.ChatUserMap) == 0 {
				fmt.Println("chatroom close")
				return
			}
		case message := <-this.Messager:
			this.broadcast(message)
		}
	}
}
