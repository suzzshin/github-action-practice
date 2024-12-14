package main

import "fmt"

// variable for version when build with ldflags
var version string

func main() {
	fmt.Println("Version:", version)
}
