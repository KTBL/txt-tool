package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/KTBL/txt-tool/files"
)

// go run ./main.go
// cls;go run ./main.go
func main() {

	fmt.Printf("voll das crazy zeug hier ... hab ab!! U+2191U+2191\n")

	file := "output/filtered-makost.log"
	lines := files.ReadLineByLine(&file)
	all := []string{}

	for _, l := range lines {
		f := strings.Split(l, " ")
		all = append(all, f[0])
	}

	filtered := unique(all)
	sort.Strings(filtered)

	fmt.Printf("length: %d\n", len(filtered))
	for _, l := range filtered {
		fmt.Printf("%s\n", l)
	}

	fmt.Printf("... ths is the end, my only friend the end!!")

}

func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func SimpleSplits() {

	s := "146.0.106.198 - - [27/Jun/2022:11:32:37 +0200] \"POST /makost/HEARTBEAT/?v-uiId=1 HTTP/1.1\" 200 5475 \"https://daten.ktbl.de/makost/\" \"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Safari/537.36 Edg/103.0.1264.37\""
	v := strings.Split(s, " ")

	// fmt.Println(v) // [This is a delimited string]
	fmt.Printf("%s\n", v[0])

}

func filter() {

	word := "makost"
	file := "data/access-20220627.log"
	lines := files.ReadLineByLine(&file)

	filtered := []string{}

	for _, l := range lines {
		if strings.Contains(l, word) {
			filtered = append(filtered, l)
		}
	}

	files.WriteLineByLine(filtered, "output/filtered-makost.log")

}

func countOccurance() {

	// count occurance word

	word := "makost"
	file := "data/access-20220627.log"
	lines := files.ReadLineByLine(&file)

	counter := 0
	for _, l := range lines {
		if strings.Contains(l, word) {

			counter++
		}
	}

	fmt.Printf("%d\n", counter)

}

func findDoublettes() {

	file := "data/txt-2.txt"

	lines := files.ReadLineByLine(&file)

	sort.Strings(lines)

	for _, l := range lines {
		fmt.Printf("%s\n", l)
	}

}
