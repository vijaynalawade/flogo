package sendWSMessage

import (
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/flow/activity"
)

func TestRegistered(t *testing.T) {
	act := activity.Get("sendSlackMessage")

	if act == nil {
		t.Error("Activity Not Registered")
		t.Fail()
		return
	}
}
