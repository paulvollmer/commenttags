# commenttags [![Build Status](https://travis-ci.org/paulvollmer/commenttags.svg?branch=master)](https://travis-ci.org/paulvollmer/commenttags)

Find TODO, FIXME, HACK, UNDONE and XXX [comment tags](https://en.wikipedia.org/wiki/Comment_(computer_programming)#Tags) at your files.

## Goal
Fast and easy to distribute (binary) tool. No language like node.js or ruby required for running the tool.

## Installation

Install the commandline interface by running
```
go get github.com/paulvollmer/commenttags/cmd/commenttags
```

## Usage

To find all tags at one file, simple run...
```
commenttags yourfile.ext
```

If you want to print the result as pretty formatted json, run...
```
commenttags -f json-pretty yourFile.ext
```

To write a json file...
```
commenttags -f json -w yourReport.json yourFile.ext
```
