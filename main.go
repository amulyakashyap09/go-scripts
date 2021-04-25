package main

import (
	"fmt"
	"go-scripts/dir"
)

func main() {
	fmt.Printf("Current Working Directory %s\n", dir.GetWorkingDirectory())
	fmt.Printf("Current Executing File %s\n", dir.GetCurrentExecutable())
	fmt.Printf("Parent Directory %s\n", dir.GetParentDirectory())
}
