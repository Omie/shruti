package pusher

import (
	"github.com/pusher/pusher-http-go"
)

var (
	client         pusher.Client
	usePusher      bool = false
	channel, event string
)

func InitPusher(appId, key, secret, pchannel, pevent string, use bool) {
	client = pusher.Client{
		AppId:  appId,
		Key:    key,
		Secret: secret,
	}
	channel = pchannel
	event = pevent
	usePusher = use
}

func Push(data interface{}) (err error) {
	if usePusher {
		_, err = client.Trigger(channel, event, data)
		return
	}
	return nil
}
