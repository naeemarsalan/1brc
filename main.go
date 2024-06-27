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

	if len(os.Args) < 2 {
		fmt.Println("Please provide a file to read")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	buff := bufio.NewReader(file)

	stations := make(map[string]Measurement)
	// lineNum := 0
	// chunk := 1000000
	for {
		line, err := buff.ReadString('\n')
		// lineNum++

		// if lineNum%chunk == 0 {
		// 	fmt.Println("Processing line: ", lineNum)
		// }

		if err != nil {
			if err == io.EOF {
				fmt.Println("End of file")
				break
			}
			panic(err)
		}
		station := strings.Split(line, ";")
		_, ok := stations[station[0]]
		temp, err := strconv.ParseFloat(strings.TrimSpace(station[1]), 64)
		if err != nil {
			panic(err)
		}
		if !ok {
			stations[station[0]] = Measurement{
				Min:   temp,
				Max:   temp,
				Count: 1,
				Total: 1,
				Mean:  0,
			}
		} else {
			selectedStation := stations[station[0]]
			selectedStation.Count++
			selectedStation.Total += temp
			if temp < selectedStation.Min {
				selectedStation.Min = temp
			}

			if temp > selectedStation.Max {
				selectedStation.Max = temp
			}

			stations[station[0]] = selectedStation
		}
		// fmt.Println(stations[station[0]])
	}

	writeFile, err := os.OpenFile("./data/results.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer writeFile.Close()

	w := bufio.NewWriter(writeFile)
	for key, _ := range stations {
		selectedStation := stations[key]
		selectedStation.Mean = (selectedStation.Total / float64(selectedStation.Count))
		stations[key] = selectedStation
		if err != nil {
			panic(err)
		}
		_, err := w.Write([]byte(fmt.Sprintf("%s;%.1f;%.1f;%.1f\n", key, selectedStation.Min, selectedStation.Max, selectedStation.Mean)))
		if err != nil {
			panic(err)
		}
	}
	w.Flush()
}
