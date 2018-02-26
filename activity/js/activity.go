package js

import (
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/robertkrimen/otto"
)

var activityLog = logger.GetLogger("activity-vijay-js")

const (
	ivInputVars = "inputVars"
	ivJs        = "javascript"
	ivOutputVars = "outputVars"
	ovOutput    = "output"
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

	inputVars, _ := context.GetInput(ivInputVars).(map[string]interface{})
	jsCode, _ := context.GetInput(ivJs).(string)

	activityLog.Debugf("Input Variables: %v", inputVars)
	activityLog.Debugf("JavaScript Code: %s", jsCode)

	vm := otto.New()

	//Set Input Variables
	for k, v := range inputVars {
		vm.Set(k, v)
	}

	v, err := vm.Run(jsCode)
	if err != nil {
		return false, activity.NewError(fmt.Sprintf("Failed to execute JavaScript code due to error: %s", err.Error()), "", nil)
	}

	outputVars, ok := context.GetInput(ivOutputVars).(map[string]interface{})
	if ok && len(outputVars) > 0 {
		//Specific variables
		result := make(map[string]interface{}, len(outputVars))
		for k := range outputVars {
			value, err := vm.Get(k)
			if err != nil || value.IsUndefined() {
				return false, activity.NewError(fmt.Sprintf("Variable:%s is not set in the java script", k), "", nil)
			}
			goVal, _ := value.Export()
			result[k] = goVal
		}
		context.SetOutput(ovOutput, result)
	} else {
		output, _ := v.Export()
		context.SetOutput(ovOutput, output)
	}

	activityLog.Debugf("Output: %v", context.GetOutput(ovOutput))
	return true, nil
}
