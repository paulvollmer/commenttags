package commenttags

import (
	"testing"
)

func Test_Tag(t *testing.T) {
	tag := Tag{TagTODO, 0, "hello tag"}
	tag.PrettyPrint()
}
