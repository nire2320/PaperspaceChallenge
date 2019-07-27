package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//MyFile data structure for JSON encoding
type MyFile struct {
	Filename string
	Packages []string
}

func main() {

	//Currently reading the working project folder, this can be changed to any specific folder the user requires
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.Contains(f.Name(), ".go") {
			//content, err := ioutil.ReadFile(f.Name())
			content := importFileReader(f.Name())

			//Iterate high to low so that when we append the new content array, we keep positions correct and don't go out of bounds
			for i := len(content) - 1; i >= 0; i-- {
				if content[i] == "import" || content[i] == "(" || content[i] == ")" {
					content = append(content[:i], content[i+1:]...)
				}
			}

			newFile := MyFile{
				Filename: f.Name(),
				Packages: content,
			}

			b, err := json.Marshal(newFile)
			if err != nil {
				log.Fatal(err)
			}
			os.Stdout.Write(b)
			fmt.Println()
		}
	}
}

// ImportFileReader will read in a given filename (string) and output a content string array containing packages imported
func importFileReader(f string) []string {

	doubleNewline := []byte{13, 10, 13, 10}              //"\r\n\r\n" byte characters (carriage return + newline)
	importString := []byte{105, 109, 112, 111, 114, 116} //"import" byte characters
	commentStart := []byte{47, 42}                       // the "/*" beginning characters of a comment
	commentEnd := []byte{42, 47}                         // the "*/" end characters of a comment

	file, err := os.Open(f)
	byteData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(string(byteData), "import") {

		//We're buffering/checking the first 10 bytes in order to check if comments are just at the beginning of the file
		//If our files contain the word "import" in a comment at the beginning of the file, it will break, but trimming out all initial comments removes this possibility
		if bytes.Contains(byteData[:10], commentStart) {
			byteData = bytes.TrimPrefix(byteData, byteData[:bytes.Index(byteData, commentEnd)])
		}

		//Modify the byte data by chopping off the prefix of anything before the string "import"
		byteData = bytes.TrimPrefix(byteData, byteData[:bytes.Index(byteData, importString)])

		//Modify the byte data by chopping off the suffix of anything after a double newline character (13 10 == "\n")
		byteData = bytes.TrimSuffix(byteData, byteData[bytes.Index(byteData, doubleNewline):])

		strArr := strings.Fields(string(byteData))

		return strArr
	}
	return nil
}
