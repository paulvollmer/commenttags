package commenttags

import (
	"testing"
)

func Test_ProcessDir(t *testing.T) {
	data, err := ProcessDirectory("./fixture", 500000)
	if err != nil {
		t.Error(err)
	}
	// if len(dir.Files) != 1 {
	// 	t.Error("ProcessFile Tags length not equal.", len(dir.Files))
	// }

	err = data.SaveJSON("./tmp.json", 0777)
	if err != nil {
		t.Error(err)
	}
}

func Test_ReadJSON(t *testing.T) {
	data, err := ReadJSON("./tmp.json")
	if err != nil {
		t.Error(err)
	}
	if len(data.Files) != 3 {
		t.Error("ReadJSON length of Files array not equal")
	}
}
