/* Program to convert arbitrary file size into Human-Readable format */

package main

import (
	"fmt"
	// "math"  // module for Mathematical functions
)

const (
	_ = iota  // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)  // left-shift operator
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main()  {
	fmt.Println("")
	
	fileSize := 4000000000.
	fmt.Printf("%.2fGB",fileSize/GB)

}