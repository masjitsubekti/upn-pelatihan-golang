package producer

import "gitlab.com/upn-belajar-go/event/model"

// Producer represents an event producer interface.
type Producer interface {
	Publish(request model.PublishRequest)
}
