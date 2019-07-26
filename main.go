package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		os.Open(f.Name())
		fmt.Println("File " + f.Name() + " Opened!")
	}

	//for each file in dir
	//read file
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

}
