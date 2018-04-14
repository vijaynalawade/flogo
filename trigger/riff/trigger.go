package riff

import (
	"context"
	syslog "log"

	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"

	// Import the aws-lambda-go. Required for dep to pull on app create
	_ "github.com/aws/aws-lambda-go/lambda"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
)

// log is the default package logger
var log = logger.GetLogger("trigger-flogo-riff")
var singleton *RiffTrigger

// RiffTrigger AWS Riff trigger struct
type RiffTrigger struct {
	metadata *trigger.Metadata
	config   *trigger.Config
	handlers []*trigger.Handler
}

//NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &RiffFactory{metadata: md}
}

// RiffFactory AWS Riff Trigger factory
type RiffFactory struct {
	metadata *trigger.Metadata
}

//New Creates a new trigger instance for a given id
func (t *RiffFactory) New(config *trigger.Config) trigger.Trigger {
	singleton = &RiffTrigger{metadata: t.metadata, config: config}
	return singleton
}

// Metadata implements trigger.Trigger.Metadata
func (t *RiffTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

func (t *RiffTrigger) Initialize(ctx trigger.InitContext) error {
	t.handlers = ctx.GetHandlers()
	return nil
}

func Invoke(input interface{}) (interface{}, error) {

	log.Info("Starting Project Riff Trigger")
	syslog.Println("Starting Project Riff Trigger")

	log.Debugf("Received Input: '%+v'\n", input)
	syslog.Printf("Received Input: '%+v'\n", input)

	//select handler, use 0th for now
	handler := singleton.handlers[0]

	idata := map[string]interface{}{
		"input": input,
	}

	results, err := handler.Handle(context.Background(), idata)

	var replyData interface{}

	if len(results) != 0 {
		dataAttr, ok := results["output"]
		if ok {
			replyData = dataAttr.Value()
		}
	}

	if err != nil {
		log.Debugf("Riff Trigger Error: %s", err.Error())
		syslog.Printf("Riff Trigger Error: %s", err.Error())
		return nil, err
	}

	log.Debugf("Riff Trigger Reply: '%+v'\n", replyData)
	syslog.Printf("Riff Trigger Reply: '%+v'\n", replyData)

	//Workaround
	strigifyData, _ := data.CoerceToString(replyData)

	return  strigifyData, err
}

func (t *RiffTrigger) Start() error {
	return nil
}

// Stop implements util.Managed.Stop
func (t *RiffTrigger) Stop() error {
	return nil
}
