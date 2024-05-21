package asciiart

import (
	"os"
)

func Printer(s string, b [][]string, outputfile *os.File) {
	// map each character of a string to its ascii character in our set
	for i := 0; i < 9; i++ {
		for _, char := range s {
			toPrint := char - 32
			outputfile.Write([]byte(b[toPrint][i])) // write the art onto the created file
		}
		outputfile.Write([]byte("\n"))
	}
}
