package main

import (
	"fmt"
	"log"

	"io/ioutil"

	"github.com/jbowtie/gokogiri/xml"
	"github.com/jbowtie/ratago/xslt"
)

const file = "test.xml"

var globalStyle, globalDoc *xml.XmlDocument
var globalStylesheet *xslt.Stylesheet
var err error

func main() {
	globalStyle, err = xml.ReadFile("test.xsl", xml.StrictParseOption)
	if err != nil {
		log.Fatalf("error reading xsl file: %v", err)
	}
	globalStylesheet, err = xslt.ParseStylesheet(globalStyle, "test.xsl")
	if err != nil {
		log.Fatalf("error parsing style sheet: %v", err)
	}
	globalDoc, err = xml.ReadFile(file, xml.StrictParseOption)
	if err != nil {
		log.Fatalf("error reading xml file: %v", err)
	}
	globalStylesheet.OutputMethod = "html"
	output, errr := globalStylesheet.Process(globalDoc, xslt.StylesheetOptions{true, nil})
	if errr != nil {
		log.Fatalf("error parsing style sheet: %v", err)
	}
	fmt.Println("---")
	fmt.Println(output)
	fmt.Println("---")
	ioutil.WriteFile("output.html", []byte(output), 0755)
}
