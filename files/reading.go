package files

import (
	"bufio"
	"log"
	"os"
)

func ReadLineByLine(str *string) []string {

	answer := []string{}

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
		line := s.Text()
		// fmt.Println(line)
		answer = append(answer, line)
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	return answer

}
