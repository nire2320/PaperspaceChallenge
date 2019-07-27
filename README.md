# PaperspaceChallenge
Go program that reads the files in a Go project folder and lists all the imported packages in the directory it's run, along with the name of the file in which the import occurs. 

# Building
Just use "go build main.go" to build the file and run main

# Testing
Currently using 3 test cases 
  - testBasic
    -- Just a basic "Hello World" test with fmt package
    
  - testComment
    -- Testing out if we would potentially have "import" in the comments at the beginning of a file
    
  - testDoubleImport
    -- Testing out somewhat less used golang syntax with multiple "import" package listings
