package files

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadLineByLine(str *string) {

	// answer := []string{}

	f, err := os.Open(*str)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)

	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

}
