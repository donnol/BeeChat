package archive

import (
	"container/list"
	"fmt"
	"strconv"
)

type EventType int

const (
	EVENT_JOIN = iota
	EVENT_LEAVE
	EVENT_MESSAGE
)

type Event struct {
	Type      EventType `json:"type"` // JOIN, LEAVE, MESSAGE
	User      string    `json:"user"`
	RUser     string    `json:"ruser"`
	Timestamp int       `json:"timestamp"` // Unix timestamp (secs)
	Content   string    `json:"content"`
}

const archiveSize = 20

// Event archives.
var archive = list.New()

// NewArchive saves new event to archive list.
func NewArchive(event Event) {
	if archive.Len() >= archiveSize {
		archive.Remove(archive.Front())
	}
	archive.PushBack(event)
}

// GetEvents returns all events after lastReceived.
func GetEvents(lastReceived int, ruser string) []Event {
	events := make([]Event, 0, archive.Len())
	for event := archive.Front(); event != nil; event = event.Next() {
		e := event.Value.(Event)
		fmt.Println(e)
		fmt.Println(e.RUser, ruser)
		if e.User == ruser {
			continue
		}
		if e.RUser != "all" && e.RUser != ruser {
			continue
		}
		if e.RUser == ruser {
			e.Content = "Only for u -> " + e.Content
		}
		if e.Timestamp > int(lastReceived) {
			events = append(events, e)
		}
	}
	fmt.Println("len of event:" + strconv.Itoa(len(events)))
	return events
}
