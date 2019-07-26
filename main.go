package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.Contains(f.Name(), ".go") {
			content, err := ioutil.ReadFile(f.Name())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("File " + f.Name() + " Go File Opened!")
			fmt.Println(f.Name() + " Contents: " + string(content))
		}
	}
}

//for each file in dir ---> DONE
//read file ---> DONE
//list all packages
//output JSON object with Filename and all imported packages
/* EX :
{
	"Filename":"Filename1",
	"Filename1Packs": ["Package1":"fmt", "Package2":"math", "Package3":"string"]
	  "File2": {
		"Package1":"fmt",
		"Package2":"string"
	  }
}
*/
