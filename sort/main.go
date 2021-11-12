package main

import (
	"log"
	"sort"
)

type sortable []int

var _ sort.Interface = sortable{}

func (s sortable) Len() int {
	return len(s)
}

func (s sortable) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortable) Swap(i, j int) {
	s[i] += s[j]
	s[j] = s[i] - s[j]
	s[i] = s[i] - s[j]
}

func main() {
	s := sortable{5, 4, 2, 3, 2, 1}
	sort.Sort(s)
	log.Println(s)
}
