package commenttags

import (
	"strings"
)

// ProcessLine parse the source and return a new Tag.
func ProcessLine(source string) *Tag {
	tag := &Tag{}
	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			tag = ProcessLineTag(source, TagFIXME)
			break
		case 1:
			tag = ProcessLineTag(source, TagTODO)
			break
		case 2:
			tag = ProcessLineTag(source, TagHACK)
			break
		case 3:
			tag = ProcessLineTag(source, TagUNDONE)
			break
		case 4:
			tag = ProcessLineTag(source, TagXXX)
			break
		}
		if tag != nil {
			return tag
		}
	}
	return nil
}

func ProcessLineTag(source, tag string) *Tag {
	r := Tag{}
	// check if tag exist
	search := tag + ":"
	// fmt.Println("Search:", search, "Tag:", tag)
	if strings.Contains(source, search) {
		// get message
		splitted := strings.SplitAfter(source, search)
		if len(splitted) > 1 {
			r.Message = strings.TrimSpace(splitted[1])
			r.Type = tag
			return &r
		}
	}
	return nil
}
