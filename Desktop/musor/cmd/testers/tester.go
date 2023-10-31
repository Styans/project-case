package main

import (
	"ascii-art/internal/drawtext"
	"fmt"
)

func main() {
	test := drawtext.GetFonts("./internal/drawtext/fonts/")
	fmt.Println(test)
}
