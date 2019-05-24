package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	create := flag.Bool("create", false, "If specified, will create the tag if it doesn't exist")
	tagList := flag.String("tag-list", "", "The list of tags to search (normally obtained with 'git tag'")
	tagFormat := flag.String("tag-format", "", "The format that the tag must respect")

	flag.Parse()

	fmt.Println(*create)
	fmt.Println(*tagList)
	fmt.Println(*tagFormat)

	fmt.Println(strings.Fields(*tagList))

	//format := tagging.Format{"", nil}
}
