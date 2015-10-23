package commenttags

import (
	"strings"
)

func ProcessSource(data []byte) *Source {
	c := &Source{}
	lines := strings.Split(string(data), "\n")
	c.TotalLines = len(lines)
	for k, line := range lines {
		// find code tags...
		tag := ProcessLine(line)
		if tag != nil {
			// add a new todo message to the Todos array
			tag.Line = k + 1
			c.Tags = append(c.Tags, *tag)
		}
	}
	return c
}
