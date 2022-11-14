package cmd

import (
	"fmt"
)

type LiveCommander struct {
	subscribers map[string]bool
	notifierId  string
}

func NewLiveCommander(notifierId string) *LiveCommander {
	return &LiveCommander{
		notifierId:  notifierId,
		subscribers: map[string]bool{},
	}
}
func (c *LiveCommander) Handle(s ApiNooter, m Message) {

	// We only allow this command on discord
	_, isDiscordNooter := s.(*DiscordNooter)
	if !isDiscordNooter {
		return
	}
	// Notifier refers to the discordId of the person allowed to notify subscribers
	if m.Author.Id == c.notifierId {
		s.NootMessage(formatNootMessage(c.subscribers))
	} else {
		//
		if _, ok := c.subscribers[m.Author.Id]; ok {
			delete(c.subscribers, m.Author.Id)
			s.NootMessage(fmt.Sprintf("<@%s> unsubscribed from livestream notifications, NooNoot!", m.Author.Id))
		} else {
			c.subscribers[m.Author.Id] = true
			s.NootMessage(fmt.Sprintf("<@%s> subscribed to livestream notifications, NootNoot!", m.Author.Id))
		}

	}
}
func formatNootMessage(subscribers map[string]bool) string {
	message := "Stream is Live! NootNoot"
	for name := range subscribers {
		message = fmt.Sprintf(message+" <@%s>", name)
	}
	return message
}
