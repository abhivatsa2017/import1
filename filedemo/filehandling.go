package filedemo

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readLineByLine() {
	file, err := os.Open("players.txt")
	if err != nil {
		//The system cannot find the file specified.
		log.Fatal(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func readAllLinesAtOnce() {
	file, err := os.Open("players.txt")
	if err != nil {
		//The system cannot find the file specified.
		log.Fatal(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	content, err := ioutil.ReadAll(reader)
	fmt.Printf(string(content))
}

func rewriteAllContent() {
	file, err := os.OpenFile("playersb.txt", os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.WriteString("Hej hopp\n")
	if err != nil {
		log.Fatal(err)
	}
}

func appendToFile() {
	file, err := os.OpenFile("playersb.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.WriteString("Hej hopp\n")
	if err != nil {
		log.Fatal(err)
	}
}

func Demo() {
}
