package commenttags

import (
	"fmt"
)

// Tag store data of one comment tag
type Tag struct {
	Type    string `json:"type"`
	Line    int    `json:"line"`
	Message string `json:"message"`
}

// Pretty return the Tag data as string
func (t *Tag) Pretty() string {
	return fmt.Sprintf("Tag %s @line %d \t '%s'\n", t.Type, t.Line, t.Message)
}

// PrettyPrint print the Tag data to stdout
func (t *Tag) PrettyPrint() {
	fmt.Printf(t.Pretty() + "\n")
}
