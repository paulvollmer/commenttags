package commenttags

import (
	"fmt"
	"strconv"
)

// Tag store data of one comment tag
type Tag struct {
	Type    string `json:"type"`
	Line    int    `json:"line"`
	Message string `json:"message"`
}

// Pretty return the Tag data as string
func (t *Tag) Pretty() string {
	return "Tag " + t.Type + " @line " + strconv.Itoa(t.Line) + " \t '" + t.Message + "'"
}

// PrettyPrint print the Tag data to stdout
func (t *Tag) PrettyPrint() {
	fmt.Println(t.Pretty())
}
