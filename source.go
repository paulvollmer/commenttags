package commenttags

import (
	"encoding/json"
)

type Source struct {
	TotalLines int `json:"total_lines"`
	// TotalTODOs  int   `json:"total_todos"`
	// TotalFIXMEs int   `json:"total_fixmes"`
	Tags []Tag `json:"tags"`
}

func (c *Source) Pretty() string {
	out := ""
	for _, v := range c.Tags {
		out += v.Pretty()
	}
	return out
}

func (c *Source) PrettyPrint() {
	for _, v := range c.Tags {
		v.PrettyPrint()
	}
}

func (d *Source) JSON() ([]byte, error) {
	data, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	return data, nil
}
