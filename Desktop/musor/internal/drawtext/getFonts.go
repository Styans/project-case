package drawtext

import (
	"fmt"
	"os"
	"strings"
)

func GetFonts(dir string) map[string][]string {
	fonts := make(map[string][]string)

	a, _ := os.ReadDir(dir)
	for _, e := range a {
		// var arr []string
		// var text string
		data, _ := os.ReadFile(dir + e.Name())

		// tempFS := strings.ReplaceAll(string(data), "\r", string(' '))
		fs := strings.Split(string(data), string('\n'))
		fmt.Println(fs)
	}
	// res, _ := io.ReadAll()

	// os.ReadFile("./internal/drawtext/fonts/shadow.txt")
	// os.ReadFile("./internal/drawtext/fonts/shadow.txt")
	return fonts
}
