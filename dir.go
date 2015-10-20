package commenttags

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type DirectoryData struct {
	Dirname string
	Files   []FileData
}

func ProcessDirectory(name string, maxSize int64) (*DirectoryData, error) {
	fmt.Println("ProcessDirectory", name)

	data := DirectoryData{}
	data.Dirname = name

	dir, err := ioutil.ReadDir(name)
	if err != nil {
		return &data, err
	}
	if len(dir) != 0 {
		for _, v := range dir {
			// fmt.Printf("Check %s %d %# v\n", name, k, v.Name())
			// check file or folder
			if v.IsDir() {
				// ignore som filders like .git
				if v.Name() != ".git" {
					tmp, _ := ProcessDirectory(name+"/"+v.Name(), maxSize)
					// HACK: exception handler...
					for _, v := range tmp.Files {
						data.Files = append(data.Files, v)
					}
				}
			} else {
				// fmt.Println("Is File...")
				// TODO: check filesize  v.Size() < maxSize
				tmpFiles, errFiles := ProcessFile(name + "/" + v.Name())
				if errFiles != nil {
					fmt.Errorf("ERROR...", errFiles)
				}
				// if tags exist, add to the files array
				if len(tmpFiles.Tags) > 0 {
					data.Files = append(data.Files, *tmpFiles)
				}
			}
		}
	}
	return &data, nil
}

func (d *DirectoryData) PrettyPrint() {
	for _, v := range d.Files {
		v.PrettyPrint()
	}
}

func (d *DirectoryData) SaveJSON(filename string, perm os.FileMode) error {
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, perm)
	return err
}
