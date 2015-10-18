package main

import (
	"fmt"
	"github.com/paulvollmer/commenttags"
	"io/ioutil"
)

type FileData struct {
	Filename string
	commenttags.CommentTags
}

func ProcessFile(src string) (*FileData, error) {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return &FileData{}, err
	}
	tags := commenttags.ProcessData(data)
	return &FileData{src, *tags}, nil
}

func (f *FileData) PrettyPrint() {
	fmt.Printf("### %s\n%s\n", f.Filename, f.Pretty())
}
