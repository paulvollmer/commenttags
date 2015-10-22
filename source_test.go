package commenttags

import (
	"testing"
)

func Test_ProcessData(t *testing.T) {
	tags := ProcessData([]byte("TODO: hello-test"))
	if len(tags.Tags) != 1 {
		t.Error("ProcessData Tags length not equal.")
	}
	if tags.Tags[0].Tag != TagTODO {
		t.Error("ProcessData Tag not equal.")
	}
	if tags.Tags[0].Message != "hello-test" {
		t.Error("ProcessData Message not equal.")
	}
	if tags.Tags[0].Line != 1 {
		t.Error("ProcessData Line not equal.")
	}
}
