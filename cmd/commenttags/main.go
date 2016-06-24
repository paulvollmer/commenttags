package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/paulvollmer/commenttags"
)

const (
	Version          = "0.1.3"
	SourceTypeFile   = 0
	SourceTypeFolder = 1
)

var (
	sourcePaths     []string
	sourceType      int
	flagVersion     = flag.Bool("v", false, "print out the `version`")
	flagFormat      = flag.String("f", "pretty", "set the output `format`. (pretty, json or json-pretty)")
	flagWrite       = flag.String("w", "", "`write` to file")
	flagMaxFilesize = flag.Int64("m", 5000000, "the `maximum filesize` to process")
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: commenttags <flags> [sources...]\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if *flagVersion {
		fmt.Println(Version)
		return
	}
	sourcePaths = make([]string, 0)
	totalArgs := len(os.Args)

	// check args and set source paths array
	if totalArgs == 1 {
		sourcePaths = append(sourcePaths, ".")
	}
	if totalArgs > 1 {
		sourcePaths = append(sourcePaths, os.Args[1:]...)
	}

	// process data
	// fmt.Println("Total sources to process:", len(sourcePaths), sourcePaths)
	for _, v := range sourcePaths {
		processSourcePath(v)
	}
}

func processSourcePath(src string) {
	folderResult := &commenttags.DirectoryData{}
	fileResult, err := commenttags.ProcessFile(src)
	if err != nil {
		// check if directory
		if err.Error() == "read "+src+": is a directory" {
			sourceType = SourceTypeFolder
			// Try to read as directory...
			folderResult, err = commenttags.ProcessDirectory(src, *flagMaxFilesize)
			if err != nil {
				fmt.Println("read directory failed!", err)
				os.Exit(1)
			}
		}
	}
	// else {
	// 	fileResult.PrettyPrint()
	// }

	//
	// format the result result
	//
	var formatted string
	switch *flagFormat {
	case "pretty":
		if sourceType == SourceTypeFile {
			if fileResult.TagsFound() {
				formatted = fileResult.Pretty()
			}
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
		fmt.Print(formatted)
	}
}
