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
	Version          = "0.1.1"
	SourceTypeFile   = 0
	SourceTypeFolder = 1
)

var (
	formatted  string
	sourceType int
)

func main() {
	flagVersion := flag.Bool("v", false, "print out the `version`")
	flagFormat := flag.String("f", "pretty", "set the output `format`. (pretty, json or json-pretty)")
	flagWrite := flag.String("w", "", "`write` to file")
	flagMaxFilesize := flag.Int64("m", 5000000, "the `maximum filesize` to process")
	flag.Parse()
	if *flagVersion {
		fmt.Println(Version)
		return
	}

	if len(os.Args) > 1 {
		//
		// process data
		//
		sourcePath := os.Args[len(os.Args)-1]
		folderResult := &commenttags.DirectoryData{}
		fileResult, err := commenttags.ProcessFile(sourcePath)
		if err != nil {
			// check if directory
			if err.Error() == "read "+sourcePath+": is a directory" {
				sourceType = SourceTypeFolder
				// Try to read as directory...
				folderResult, err = commenttags.ProcessDirectory(sourcePath, *flagMaxFilesize)
				if err != nil {
					fmt.Println("read directory failed!", err)
					os.Exit(1)
				}
			}
		}

		//
		// format the result result
		//
		switch *flagFormat {
		case "pretty":
			if sourceType == SourceTypeFile {
				formatted = fileResult.Pretty()
			} else if sourceType == SourceTypeFolder {
				formatted += folderResult.Pretty()
			}
			break
		case "json":
			if sourceType == SourceTypeFile {
				out, _ := fileResult.JSON()
				formatted = string(out)
			} else if sourceType == SourceTypeFolder {
				tmp, _ := folderResult.JSON()
				formatted = string(tmp)
			}
			break
		case "json-pretty":
			if sourceType == SourceTypeFile {
				out, _ := json.MarshalIndent(fileResult, "", "  ")
				formatted = string(out)
			} else if sourceType == SourceTypeFolder {
				out, _ := json.MarshalIndent(folderResult, "", "  ")
				formatted = string(out)
			}
			break
		default:
			fmt.Printf("Format '%s' not supported! Choose between the following formats:\n", *flagFormat)
			fmt.Println("- pretty (default)")
			fmt.Println("- json")
			fmt.Println("- json-pretty")
			break
		}

		//
		// write or print out the result
		//
		if *flagWrite != "" {
			ioutil.WriteFile(*flagWrite, []byte(formatted), 0777)
		} else {
			fmt.Println(formatted)
		}

	} else {
		fmt.Println("Missing Filepath, See the -h help out...")
	}
}
