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

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.Contains(f.Name(), ".go") {
			//content, err := ioutil.ReadFile(f.Name())
			content := importFileReader(f.Name())

			//Iterate high to low so that when we append the new content array, we keep positions and don't go out of bounds
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

func importFileReader(f string) []string {

	doubleNewline := []byte{13, 10, 13, 10}              //"\n\n" byte characters
	importString := []byte{105, 109, 112, 111, 114, 116} //"import" byte characters

	file, err := os.Open(f)
	byteData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(string(byteData), "import") {
		//Modify the byte data by chopping off the prefix of anything before the string "import" (this will be buggy if files contain comments with "import" in them...)
		byteData = bytes.TrimPrefix(byteData, byteData[:bytes.Index(byteData, importString)])

		//Modify the byte data by chopping off the suffix of anything after a double newline character (13 10 == "\n")
		byteData = bytes.TrimSuffix(byteData, byteData[bytes.Index(byteData, doubleNewline):])

		strArr := strings.Fields(string(byteData))

		return strArr
	}
	return nil
}

//for each file in dir ---> DONE
//read file ---> DONE
//list all packages and filename ---> DONE
//output JSON object with Filename and all imported packages
/* EX :
{
	"Filename":"Filename1",
	"Packages": ["fmt", "math", "string"]
}
*/
