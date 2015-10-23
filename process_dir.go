package commenttags

import (
	"encoding/json"
	"io/ioutil"
	"os"
	// "strconv"
)

type DirectoryData struct {
	Dirname             string     `json:"dirname"`
	TotalProcessedFiles int        `json:"total_processed_files"`
	TotalProcessedLines int        `json:"total_processed_lines"`
	Files               []FileData `json:"files"`
}

func ProcessDirectory(name string, maxSize int64) (*DirectoryData, error) {
	// fmt.Println("ProcessDirectory", name)
	data := &DirectoryData{}
	data.Dirname = name

	dir, err := ioutil.ReadDir(name)
	if err != nil {
		return nil, err
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
					data.TotalProcessedLines += tmp.TotalProcessedLines
					data.TotalProcessedFiles += tmp.TotalProcessedFiles
					for _, v := range tmp.Files {
						data.Files = append(data.Files, v)
					}
				}
			} else {
				// fmt.Println("Is File...")
				// TODO: check filesize  v.Size() < maxSize
				tmpFiles, errFiles := ProcessFile(name + "/" + v.Name())
				if errFiles != nil {
					return nil, errFiles
				}
				data.TotalProcessedLines += tmpFiles.TotalLines
				data.TotalProcessedFiles++
				// if tags exist, add to the files array
				if len(tmpFiles.Tags) > 0 {
					data.Files = append(data.Files, *tmpFiles)
				}
			}
		}
	}
	return data, nil
}

func (d *DirectoryData) Pretty() string {
	out := ""
	// additional infos...
	// out += "Dirname: " + d.Dirname + "\n"
	// out += "Total processed files: " + strconv.Itoa(d.TotalProcessedFiles) + "\n"
	// out += "Total processed lines: " + strconv.Itoa(d.TotalProcessedLines) + "\n"
	// out += "----------------------\n\n"
	for _, v := range d.Files {
		out += v.Pretty()
	}
	return out
}

func (d *DirectoryData) PrettyPrint() {
	for _, v := range d.Files {
		v.PrettyPrint()
	}
}

func (d *DirectoryData) JSON() ([]byte, error) {
	data, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *DirectoryData) SaveJSON(filename string, perm os.FileMode) error {
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, perm)
	return err
}

func ParseJSON(data []byte) (DirectoryData, error) {
	v := DirectoryData{}
	err := json.Unmarshal(data, &v)
	return v, err
}

func ReadJSON(filename string) (*DirectoryData, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return &DirectoryData{}, err
	}
	v, errParse := ParseJSON(data)
	return &v, errParse
}
