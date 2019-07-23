package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	//  Create a Buffer value and write a string to the buffer.
	// Using the Write method that implements io.Writer
	var b bytes.Buffer
	b.Write([]byte("Hello "))

	// Use Fprintf to concatenate a string to the Buffer
	// Passing the address of a bytes.Buffer value for the io.Writer
	// Fprintf formats according to a format specifier and writes to w. It
	// returns the number of bytes written and any write error encountered.
	// func Fprintf(w io.Writer, format string, a ...interface{})(n int, err error)
	fmt.Fprintf(&b, "World!")

	// Write the content of the Buffer to the stdout device
	//  Passing the address of a os.File value for io.Writer
	b.WriteTo(os.Stdout)
}
