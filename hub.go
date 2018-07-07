// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"sync"
	"encoding/json"
	"log"
)



// hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	//clients map[*Client]bool
	clients sync.Map

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	//count of client
	ClientNum uint16

}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// send json object to a client with peerId
func (this *Hub) sendJsonToClient(peerId string, value interface{})  {
	b, err := json.Marshal(value)
	if err != nil {
		//logrus.Errorf("[Client.jsonResponse] Marshal err: %s", err.Error())
		return
	}
	client, ok := this.clients.Load(peerId)
	if !ok {
		log.Printf("sendJsonToClient error")
		return
	}
	if err := client.(*Client).sendMessage(b); err != nil {
		//logrus.Errorf("[Client.jsonResponse] sendMessage err: %s", err.Error())
	}
	//if err := client.(*Client).conn.WriteJSON(value); err != nil {
	//	//logrus.Errorf("[Client.jsonResponse] sendMessage err: %s", err.Error())
	//}
}

func (this *Hub) run() {
	for {
		select {
		case client := <-this.register:
			this.doRegister(client)
		case client := <-this.unregister:
			this.doUnregister(client)
		}
	}
}

func (this *Hub) doRegister(client *Client) {
	//	logrus.Debugf("[Hub.doRegister] %s", client.id)
	if client.PeerId != "" {
		this.clients.Store(client.PeerId, client)
		this.ClientNum ++
	}
}

func (this *Hub) doUnregister(client *Client) {
	//	logrus.Debugf("[Hub.doUnregister] %s", client.id)

	if client.PeerId == "" {
		return
	}

	_, ok := this.clients.Load(client.PeerId)

	if ok {
		//delRecordCh <- client.id
		this.clients.Delete(client.PeerId)
		close(client.send)
		this.ClientNum --
	}

}
