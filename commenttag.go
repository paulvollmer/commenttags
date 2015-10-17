package commenttags

import (
	"fmt"
	"strings"
)

// CommentTag store data for a code comment tag
type CommentTag struct {
	Tag     string
	Line    int
	Message string
}

// Use the ParseComment func to parse a source and return a new CommentTag.
func ParseComment(source string) (*CommentTag, bool) {
	var tag *CommentTag
	var found bool
	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			tag, found = ParseCommentTag(source, TagFIXME)
			break
		case 1:
			tag, found = ParseCommentTag(source, TagTODO)
			break
		case 2:
			tag, found = ParseCommentTag(source, TagHACK)
			break
		case 3:
			tag, found = ParseCommentTag(source, TagUNDONE)
			break
		case 4:
			tag, found = ParseCommentTag(source, TagXXX)
			break
		}
		if found {
			return tag, true
		}
	}
	return &CommentTag{}, false
}

func ParseCommentTag(source, tag string) (*CommentTag, bool) {
	r := CommentTag{}
	// check if tag exist
	search := tag + ":"
	// fmt.Println("Search:", search, "Tag:", tag)
	if strings.Contains(source, search) {
		// get message
		splitted := strings.SplitAfter(source, search)
		if len(splitted) > 1 {
			r.Message = strings.TrimSpace(splitted[1])
			r.Tag = tag
			return &r, true
		}
	}
	return &r, false
}

// Print the data to stdout
func (t *CommentTag) Pretty() string {
	return fmt.Sprintf("Tag: %s \tLine: %d \tMessage: '%s'\n", t.Tag, t.Line, t.Message)
}

func (t *CommentTag) PrettyPrint() {
	fmt.Printf(t.Pretty())
}
