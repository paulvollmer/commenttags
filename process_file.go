package commenttags

import (
	"fmt"
	"io/ioutil"
)

type FileData struct {
	Filename string `json:"filename"`
	Source
}

func ProcessFile(src string) (*FileData, error) {
	// fmt.Println("Call ProcessFile", src)
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return nil, err
	}
	tags := ProcessSource(data)
	return &FileData{src, *tags}, nil
}

func (f *FileData) Pretty() string {
	if f.TagsFound() {
		return fmt.Sprintf("### %s\n%s  \n", f.Filename, f.PrettySource())
	}
	return ""
}

func (f *FileData) PrettyPrint() {
	if f.TagsFound() {
		fmt.Printf("### %s  \n%s", f.Filename, f.Pretty())
	}
}
