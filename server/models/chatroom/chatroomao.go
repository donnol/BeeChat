package chatroom

import (
	"container/list"
	"fmt"
	"time"

	"github.com/astaxie/beego"

	"beechat/models/archive"
)

type Subscription struct {
	Archive []archive.Event
	New     <-chan archive.Event
}

func NewEvent(ep archive.EventType, user, ruser, msg string) archive.Event {
	return archive.Event{ep, user, ruser, int(time.Now().Unix()), msg}
}

func Join(user string) {
	subscribe <- Subscriber{Name: user}
}

func Leave(user string) {
	unsubscribe <- user
}

type Subscriber struct {
	Name string
}

var (
	// Channel for new join users.
	subscribe = make(chan Subscriber, 10)
	// Channel for exit users.
	unsubscribe = make(chan string, 10)
	// Send events here to publish them.
	Publish = make(chan archive.Event, 10)
	// Long polling waiting list.
	WaitingList = list.New()
	subscribers = list.New()
)

// This function handles all incoming chan messages.
func chatroom() {
	fmt.Println("into chatroom.")
	for {
		select {
		case sub := <-subscribe:
			if !isUserExist(subscribers, sub.Name) {
				subscribers.PushBack(sub) // Add user to the end of list.
				// Publish a JOIN event.
				Publish <- NewEvent(archive.EVENT_JOIN, sub.Name, "all", "")
				beego.Info("New user:", sub.Name)
			} else {
				beego.Info("Old user:", sub.Name)
			}
		case event := <-Publish:
			// Notify waiting list.
			var n *list.Element
			for ch := WaitingList.Back(); ch != nil; ch = n {
				ch.Value.(chan bool) <- true
				n = ch.Prev()
				WaitingList.Remove(ch)
			}

			//broadcastWebSocket(event)
			archive.NewArchive(event)

			if event.Type == archive.EVENT_MESSAGE {
				beego.Info("Message from", event.User, ";Content:", event.Content)
			}
		case unsub := <-unsubscribe:
			for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(Subscriber).Name == unsub {
					subscribers.Remove(sub)
					Publish <- NewEvent(archive.EVENT_LEAVE, unsub, "all", "") // Publish a LEAVE event.
					break
				}
			}
		}
	}
}

func init() {
	go chatroom()
}

func isUserExist(subscribers *list.List, user string) bool {
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscriber).Name == user {
			return true
		}
	}
	return false
}
