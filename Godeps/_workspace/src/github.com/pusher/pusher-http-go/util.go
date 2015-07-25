package pusher

import (
	"errors"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

var channelValidationRegex = regexp.MustCompile("^[-a-zA-Z0-9_=@,.;]+$")
var socketIdValidationRegex = regexp.MustCompile(`\A\d+\.\d+\z`)

func authTimestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func parseAuthRequestParams(_params []byte) (string, string, error) {
	params, err := url.ParseQuery(string(_params))

	if err != nil {
		return "", "", err
	}

	return params["channel_name"][0], params["socket_id"][0], nil
}

func channelsAreValid(channels []string) bool {
	for _, channel := range channels {
		if len(channel) > 200 || !channelValidationRegex.MatchString(channel) {
			return false
		}
	}
	return true
}

func validateSocketId(socketId *string) (err error) {
	if (socketId == nil) || socketIdValidationRegex.MatchString(*socketId) {
		return
	} else {
		return errors.New("socket_id invalid")
	}
}
