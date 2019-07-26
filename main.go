package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	//"encoding/json"
)

//Our JSON data structure
/*type JSONFiles struct {
	filename string
	packages []string
}
*/

func main() {

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.Contains(f.Name(), ".go") {
			//content, err := ioutil.ReadFile(f.Name())
			content := importFileReader(f.Name())
			fmt.Println("File " + f.Name() + " Go File Opened!")
			for i := range content {
				if content[i] == ")" {
					content = content[:i]
					break
				}
			}
			fmt.Println(content[1:])
			fmt.Println()
		}
	}
}

func importFileReader(f string) []string {

	byteData, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	dataToString := string(byteData)

	if strings.Contains(dataToString, "import") {
		//Modify the data string by chopping off the prefix of anything before the string "import" (this will be buggy later if files contain comments with "import" in them...)
		dataToString = strings.TrimPrefix(dataToString, dataToString[:strings.Index(dataToString, "import")])
		strArr := strings.Fields(dataToString)
		if strArr[1] == "(" {
			return strArr[1:]
		}
		return strArr
	}
	return nil
}

//for each file in dir ---> DONE
//read file ---> DONE
//list all packages
//output JSON object with Filename and all imported packages
/* EX :
{
	"Filename":"Filename1",
	"Packages": ["fmt", "math", "string"]
}
*/
