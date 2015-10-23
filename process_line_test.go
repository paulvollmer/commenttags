package commenttags

import (
	"testing"
)

var ProcessLine_TestData = []struct {
	source string
	result Tag
}{
	{"FIXME: hello world", Tag{"FIXME", 1, "hello world"}},
	{"TODO: hello world", Tag{"TODO", 1, "hello world"}},
	{"HACK: hello world", Tag{"HACK", 1, "hello world"}},
	{"UNDONE: hello world", Tag{"UNDONE", 1, "hello world"}},
	{"XXX: hello world", Tag{"XXX", 1, "hello world"}},
}

func Test_ProcessLine(t *testing.T) {
	for _, tt := range ProcessLine_TestData {
		t.Logf("Testing '%s'", tt.source)

		result := ProcessLine(tt.source)
		if result == nil || result.Message != tt.result.Message || result.Type != tt.result.Type {
			t.Errorf("ProcessLine '%s' failed!", tt.result.Type)
		}

	}
}
