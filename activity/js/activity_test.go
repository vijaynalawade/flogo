package js

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval_InputVars(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	inputVar := make(map[string]interface{}, 2)
	inputVar["n1"] = 2
	inputVar["n2"] = 3
	tc.SetInput(ivInputVars, inputVar)
	tc.SetInput(ivJs, `jsInput.n1 + jsInput.n2`)
	ok, err := act.Eval(tc)
	assert.NoError(t, err)

	if ok {
		sum, _ := data.CoerceToInteger(tc.GetOutput(ovOutput))
		assert.Equal(t, 5, sum)
	}

}

func TestEval_OutputVars(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//Set input variables
	inputVar := make(map[string]interface{}, 2)
	inputVar["n1"] = 2
	inputVar["n2"] = 3
	tc.SetInput(ivInputVars, inputVar)

	//Set JS code
	tc.SetInput(ivJs, `var jsOutput = {}; jsOutput["sum"] = jsInput.n1 + jsInput.n2; jsOutput["result"] = "Sum is " + jsOutput["sum"];`)

	success, err := act.Eval(tc)
	assert.NoError(t, err)

	if success {
		output, ok := tc.GetOutput(ovOutput).(map[string]interface{})
		assert.True(t, ok)

		sum, _ := data.CoerceToInteger(output["sum"])
		assert.Equal(t, 5, sum)

		result, _ := data.CoerceToString(output["result"])
		assert.Equal(t, "Sum is 5", result)
	}
}
