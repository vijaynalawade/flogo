package js

import (
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/robertkrimen/otto"
)

var activityLog = logger.GetLogger("activity-vijay-js")

const (
	ivInputVars = "jsInput"
	ivJs        = "javascript"
	ovOutput    = "jsOutput"
)

type JSActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &JSActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *JSActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Execute JS code
func (a *JSActivity) Eval(context activity.Context) (done bool, err error) {

	inputVars, _ := context.GetInput(ivInputVars).(interface{})
	jsCode, _ := context.GetInput(ivJs).(string)

	activityLog.Debugf("JavaScript Input: %v", inputVars)
	activityLog.Debugf("JavaScript Code: %s", jsCode)

	vm := otto.New()

	//Set Input Variable
	vm.Set(ivInputVars, inputVars)

	v, err := vm.Run(jsCode)
	if err != nil {
		return false, activity.NewError(fmt.Sprintf("Failed to execute JavaScript code due to error: %s", err.Error()), "", nil)
	}

	var jsOutput interface{}
	// Look for jsOutput variable value
	value, err := vm.Get(ovOutput)
	if value.IsNull() || value.IsUndefined() {
		// Set returned value
		jsOutput, _ = v.Export()
	} else {
		// Set jsOutput variable value
		jsOutput, _ = value.Export()
	}
	context.SetOutput(ovOutput, jsOutput)

	activityLog.Debugf("JavaScript Output: %v", jsOutput)
	return true, nil
}
