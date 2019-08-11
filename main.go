package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	correct := 0
	wrong := 0
	dat, _ := ioutil.ReadFile("./quiz1.csv")
	str := string(dat)
	r := csv.NewReader(strings.NewReader(str))
	reader := bufio.NewReader(os.Stdin)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(record[0], " -> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare(record[1], text) == 0 {
			correct++
		} else {
			wrong++
		}
	}
	fmt.Printf("Correct: %d, Wrong: %d\nPercent Correct: %d%%\n", correct, wrong, int(float64(correct)/float64(correct+wrong)*100))
}
