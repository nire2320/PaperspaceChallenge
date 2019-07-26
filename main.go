package main

import "fmt"

func main() {

	fmt.Println("Hello World!")

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
