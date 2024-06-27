package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Measurement struct {
	Min   float64
	Max   float64
	Count int
	Total float64
	Mean  float64
}

func main() {

	file, err := os.Open("./data/measurements.txt")

	if err != nil {
		panic(err)
	}

	buff := bufio.NewReader(file)

	stations := make(map[string]Measurement)

	for {
		line, err := buff.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		station := strings.Split(line, ";")
		_, ok := stations[station[0]]

		if !ok {
			temp, err := strconv.ParseFloat(strings.TrimSpace(station[1]), 64)
			if err != nil {
				panic(err)
			}
			stations[station[0]] = Measurement{
				Min:   temp,
				Max:   temp,
				Count: 1,
				Total: 1,
				Mean:  0,
			}
		}
		fmt.Println(stations[station[0]])
	}
}
