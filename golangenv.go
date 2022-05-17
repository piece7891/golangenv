package main

import (
  "os"
  "fmt"  
  "strings"  
)

func main() {	
	// fetcha all env variables
	for _, element := range os.Environ() {
        variable := strings.Split(element, "=")
        fmt.Println(variable[0],"=",variable[1])		
    }
}
