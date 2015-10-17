package commenttags

import (
	"testing"
)

var parseCommentTagTest = []struct {
	source string
	found  bool
	result CommentTag
}{
	{"FIXME: hello world", true, CommentTag{"FIXME", 1, "hello world"}},
	{"TODO: hello world", true, CommentTag{"TODO", 1, "hello world"}},
	{"HACK: hello world", true, CommentTag{"HACK", 1, "hello world"}},
	{"UNDONE: hello world", true, CommentTag{"UNDONE", 1, "hello world"}},
	{"XXX: hello world", true, CommentTag{"XXX", 1, "hello world"}},
}

func Test_ParseComment(t *testing.T) {
	for _, tt := range parseCommentTagTest {
		t.Logf("Testing '%s'", tt.source)
		result, found := ParseComment(tt.source)
		if found != tt.found || result.Message != tt.result.Message || result.Tag != tt.result.Tag {
			t.Errorf("ParseCommentTag '%s' failed!", tt.result.Tag)
		}
	}
}
