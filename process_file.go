package commenttags

import (
	"fmt"
	"io/ioutil"
)

type FileData struct {
	Filename string
	CommentTags
}

func ProcessFile(src string) (*FileData, error) {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return nil, err
	}
	tags := ProcessData(data)
	return &FileData{src, *tags}, nil
}

func (f *FileData) PrettyPrint() {
	fmt.Printf("### %s\n%s\n", f.Filename, f.Pretty())
}
