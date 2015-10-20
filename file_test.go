package main

import (
	"testing"
)

func Test_ProcessFile(t *testing.T) {
	tags, err := ProcessFile("../../fixture/sample.js")
	if err != nil {
		t.Error(err)
	}
	if tags.Filename != "../../fixture/sample.js" {
		t.Error("ProcessFile Filename not equal.")
	}

	if len(tags.Tags) != 8 {
		t.Error("ProcessFile Tags length not equal.")
	}
	if tags.Tags[0].Tag != "TODO" {
		t.Error("ProcessFile Tag not equal.")
	}
	if tags.Tags[0].Message != "test-1" {
		t.Error("ProcessFile Message not equal.")
	}
	// if tags.Tags[0].Line != 2 {
	// 	t.Error("ProcessFile Line not equal.")
	// }
}
