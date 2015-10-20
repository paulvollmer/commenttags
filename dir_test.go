package commenttags

import (
	"testing"
)

func Test_ProcessDir(t *testing.T) {
	_, err := ProcessDirectory("./fixture", 500000)
	if err != nil {
		t.Error(err)
	}
	// if len(dir.Files) != 1 {
	// 	t.Error("ProcessFile Tags length not equal.", len(dir.Files))
	// }
}
