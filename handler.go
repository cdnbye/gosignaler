package main

import (
	"encoding/json"
	"fmt"
	"log"
)


type Handler interface {
	Handle()
}

type SignalMsg struct {
	To_peer_id string          `json:"to_peer_id"`
	Data  interface{}          `json:"data"`
}

type SignalResp struct {
	Action string              `json:"action"`
	FromPeerId string          `json:"from_peer_id"`
	Data interface{}           `json:"data,omitempty"`
}


func (this *Client) handle(message []byte) {
	//	logrus.Debugf("[Client.handle] %s", string(message))
	action := struct {
		Action     string `json:"action"`
	}{}
	if err := json.Unmarshal(message, &action); err != nil {
		//log.Println(err)
		return
	}

	this.CreateHandler(action.Action, message).Handle()
}

func (this *Client) CreateHandler(action string, payload []byte) Handler {

	switch action {
	case "signal":
		msg := SignalMsg{}
		if err := json.Unmarshal(payload, &msg); err != nil {
			log.Println(err)
			return  &ExceptionHandler{err.Error()}
		}
		return &SignalHandler{msg, this}
	}

	return &ExceptionHandler{message: fmt.Sprintf("unregnized action %s", action)}
}

type ExceptionHandler struct {
	message string
}

func (this *ExceptionHandler) Handle() {
	//log.Printf("[ExceptionHandler] err %s", this.message)
}

type SignalHandler struct {
	message SignalMsg
	client *Client
}

func (this *SignalHandler) Handle()  {
	_, ok := this.client.hub.clients.Load(this.message.To_peer_id)        //判断节点是否还在线
	if ok {
		resp := SignalResp{
			Action: "signal",
			FromPeerId: this.client.PeerId,
			Data: this.message.Data,
		}
		this.client.hub.sendJsonToClient(this.message.To_peer_id, resp)
	} else {
		//log.Println("Peer not found")
		resp := SignalResp{
			Action: "signal",
			FromPeerId: this.message.To_peer_id,
		}
		// 发送一次后，同一peerId下次不再发送，节省带宽
		if !this.client.InvalidPeers[this.message.To_peer_id] {
			this.client.InvalidPeers[this.message.To_peer_id] = true
			this.client.hub.sendJsonToClient(this.client.PeerId, resp)
		}
	}
}





