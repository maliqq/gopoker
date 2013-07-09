package server

import (
	"fmt"
	"log"
)

import (
	"gopoker/model"
)

type Room struct {
	Id string

	Game  *model.Game
	Table *model.Table

	recv chan string
	send chan string
}

func (room Room) String() string {
	return fmt.Sprintf("id=%s", room.Id)
}

func (room Room) Pause() {

}

func (room Room) Close() {

}

func (room Room) Destroy() {

}

func (room Room) Start() {
	log.Printf("starting room %s", room)
	for {
		select {
		case m := <-room.recv:
			reply := m
			log.Printf("got: %s", m)
			room.send <- reply
		}
	}
}
