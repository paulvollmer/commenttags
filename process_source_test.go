package commenttags

import (
	"testing"
)

func Test_ProcessSource(t *testing.T) {
	tags := ProcessSource([]byte("TODO: hello-test"))
	if len(tags.Tags) != 1 {
		t.Error("ProcessSource Tags length not equal.")
	}
	if tags.Tags[0].Type != TagTODO {
		t.Error("ProcessSource Tag not equal.")
	}
	if tags.Tags[0].Message != "hello-test" {
		t.Error("ProcessSource Message not equal.")
	}
	if tags.Tags[0].Line != 1 {
		t.Error("ProcessSource Line not equal.")
	}
	tags.PrettyPrint()
	tagsJSON, err := tags.JSON()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(tagsJSON))
}
