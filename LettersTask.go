package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Letters struct {
	Counter map[string]int
}

func (l Letters) String() string {
	letters := make([]string, 0, len(l.Counter))
	for t := range l.Counter {
		letters = append(letters, t)
	}
	sort.Strings(letters)
	var sb strings.Builder
	for _, t := range letters {
		sb.WriteString(t)
		sb.WriteString(strconv.Itoa(l.Counter[t]))
	}
	return sb.String()
}

func main() {
	//input := "aaaaУУУaaddddЯЯЯddvvwwwwaaaыыыыпSSSSппппп"
	var input string
	fmt.Scanln(&input)
	var r Letters
	r.Counter = make(map[string]int)
	for _, c := range input {
		r.Counter[string(c)]++
	}
	fmt.Println(r)
}
