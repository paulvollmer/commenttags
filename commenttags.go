package commenttags

import (
	"strings"
)

type CommentTags struct {
	Tags []CommentTag
}

func ProcessData(data []byte) *CommentTags {
	c := &CommentTags{}
	lines := strings.Split(string(data), "\n")
	for k, line := range lines {
		// find code tags...
		tag, found := ParseComment(line)
		tag.Line = k + 1
		if found {
			// add a new todo message to the Todos array
			c.Tags = append(c.Tags, *tag)
		}
	}
	return c
}

func (c *CommentTags) Pretty() string {
	out := ""
	for _, v := range c.Tags {
		out += v.Pretty()
	}
	return out
}

func (c *CommentTags) PrettyPrint() {
	for _, v := range c.Tags {
		v.PrettyPrint()
	}
}
