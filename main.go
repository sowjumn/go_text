package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {

	var scanner *bufio.Scanner
	if len(os.Args) < 2 {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		fileName := os.Args[1]
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()

		scanner = bufio.NewScanner(file)
	}

	space := regexp.MustCompile(`\s+`)
	countMap := make(map[string]int)
	r := strings.NewReplacer("!", " ", ",", " ", ";", " ", ":", " ", ".", " ", "\n", " ", "\t", " ", "\"", " ", "?", " ", "#", " ", "(", " ", ")", " ", "*", " ")

	for scanner.Scan() {
		var sequence string
		currLine := scanner.Text()
		currLine = r.Replace(currLine)
		currLine = space.ReplaceAllString(currLine, " ")
		currLine = strings.TrimSpace(currLine)
		if isEmptyLine(currLine) {
			continue
		}
		currLine = strings.ToLower(currLine)
		currWords := strings.Split(currLine, " ")
		for i := 0; i < len(currWords); i++ {
			if i+3 <= len(currWords) {
				sequence = strings.Join(currWords[i:i+3], " ")
				countMap[sequence] += 1
			}
		}
	}

	var valueCount []int
	for _, v := range countMap {
		valueCount = append(valueCount, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(valueCount)))

	var tophundred []int
	if len(valueCount) <= 100 {
		tophundred = valueCount[0:]
	} else {
		tophundred = valueCount[0:100]
	}

	tophundredMap := make(map[int][]string)
	for k, v := range countMap {
		if include(tophundred, v) {
			tophundredMap[v] = append(tophundredMap[v], k)

		}
	}

	finalCount := 0
	tophundred = unique(tophundred)
	sort.Sort(sort.Reverse(sort.IntSlice(tophundred)))
	for _, countVal := range tophundred {
		currStrs := tophundredMap[countVal]
		for _, currStr := range currStrs {
			if finalCount == 100 {
				break
			}
			fmt.Printf("%s - %d\n", currStr, countVal)
			finalCount += 1
		}
		if finalCount == 100 {
			break
		}
	}
}

func include(s []int, key int) bool {
	for _, val := range s {
		if val == key {
			return true
		}
	}

	return false
}

func unique(s []int) []int {
	uniqueMap := make(map[int]bool)
	for _, v := range s {
		uniqueMap[v] = true
	}

	var outSlice []int
	for k, _ := range uniqueMap {
		outSlice = append(outSlice, k)
	}
	return outSlice
}

func isEmptyLine(line string) bool {
	trimmedLine := strings.TrimSpace(line)
	return len(trimmedLine) == 0
}
