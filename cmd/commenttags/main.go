package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/paulvollmer/commenttags"
	"io/ioutil"
	"os"
)

const (
	Version     = "0.1.0"
	MaxSizeFile = 5000000
)

func main() {
	flagVersion := flag.Bool("v", false, "Print out the version")
	flagFormat := flag.String("f", "pretty", "format as pretty formatted text or json")
	flagWrite := flag.String("write", "", "write to file")
	flag.Parse()
	if *flagVersion {
		fmt.Println(Version)
		return
	}

	if len(os.Args) > 1 {
		// TODO: check if file or directory
		sourcePath := os.Args[len(os.Args)-1]
		fileResult, err := commenttags.ProcessFile(sourcePath)
		if err != nil {
			// fmt.Println("File Processing failed!", err)
			if err.Error() == "read "+sourcePath+": is a directory" {
				// fmt.Println("Try to read as Directory...")
				dirResult, errDir := commenttags.ProcessDirectory(sourcePath, MaxSizeFile)
				if errDir != nil {
					fmt.Println("Read Directory failed!", err)
					return
				}
				// dirResult PrettyPrint()
				// fmt.Println(dirResult)
				for _, v := range dirResult.Files {
					v.PrettyPrint()
				}
			}
			return
		}

		// format the result result
		formatted := ""
		switch *flagFormat {
		case "pretty":
			formatted = fileResult.Pretty()
			break
		case "json":
			out, _ := json.Marshal(fileResult)
			formatted = string(out)
			break
		case "json-pretty":
			out, _ := json.MarshalIndent(fileResult, "", "  ")
			formatted = string(out)
			break
		default:
			fmt.Println("Format not supported")
			break
		}

		// write or print out the result
		if *flagWrite != "" {
			ioutil.WriteFile(*flagWrite, []byte(formatted), 0777)
		} else {
			fmt.Println(formatted)
		}

	} else {
		fmt.Println("Missing Filepath")
	}
}
