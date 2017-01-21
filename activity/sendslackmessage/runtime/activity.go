package sendWSMessage

import (
	"bytes"
	"encoding/json"
	"github.com/TIBCOSoftware/flogo-lib/flow/activity"
	"github.com/op/go-logging"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// log is the default package logger
var log = logging.MustGetLogger("activity-sendSlackMessage")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

type Payload struct {
	Channel   string `json:"channel"`
	Message   string `json:"text"`
	Username  string `json:"username"`
	Iconemoji string `json:"icon_emoji"`
}

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(&MyActivity{metadata: md})
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval - Sends a message to a Slack channel
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the activity data from the context
	webHookUrl := context.GetInput("WebHook").(string)
	channelName := context.GetInput("Channel").(string)
	message := context.GetInput("Message").(string)
	user := context.GetInput("Username").(string)
	iconEmoji := context.GetInput("Iconemoji").(string)

	if len(webHookUrl) == 0 {
		panic("WebHook URL must be configured.")
	}

	payload := &Payload{}
	if len(channelName) > 0 {
		payload.Channel = channelName
	}

	if len(user) > 0 {
		payload.Username = user
	}

	if len(iconEmoji) > 0 {
		payload.Iconemoji = iconEmoji
	}

	if len(message) == 0 {
		panic("Message must be configured.")
	}
	
	payload.Message = message
	b, _ := json.Marshal(payload)
	data := url.Values{}
	data.Set("payload", string(b))
	req, _ := http.NewRequest("POST", webHookUrl, bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	client := &http.Client{}
	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)
	context.SetOutput("result", string(body))

	return true, nil
}
