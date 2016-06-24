package commenttags

import (
	"encoding/json"
	"fmt"
)

// Source store the tags of one text source
type Source struct {
	TotalLines int `json:"total_lines"`
	// TotalTODOs  int   `json:"total_todos"`
	// TotalFIXMEs int   `json:"total_fixmes"`
	Tags []Tag `json:"tags"`
}

func (c *Source) TagsFound() bool {
	if len(c.Tags) != 0 {
		return true
	}
	return false
}

// Pretty return the Source data as string
func (c *Source) PrettySource() string {
	out := ""
	for _, v := range c.Tags {
		out += v.Pretty() + "  \n"
	}
	return out
}

// PrettyPrint print the Source data to stdout
func (c *Source) PrettyPrint() {
	fmt.Println(c.PrettySource())
}

// JSON return the Store data as formatted json byte array
func (d *Source) JSON() ([]byte, error) {
	data, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	return data, nil
}
