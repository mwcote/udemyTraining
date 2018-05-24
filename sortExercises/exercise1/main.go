package main

import (
	"fmt"
	"sort"
)

type people []string

func (ppl people) Len() int {
	return len(ppl)
}

func (ppl people) Swap(i int, j int) {
	temp := ppl[i]
	ppl[i] = ppl[j]
	ppl[j] = temp
}

func (ppl people) Less(i int, j int) bool {
	return ppl[i] < ppl[j]
}

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

	studyGroup2 := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup2)
	sort.Strings(studyGroup2)
	fmt.Println(studyGroup2)

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println(n)
	sort.Ints(n)
	fmt.Println(n)

	fmt.Println("Reverse")

	sort.Slice(studyGroup, func(i int, j int) bool {
		return studyGroup[i] >= studyGroup[j]
	})

	fmt.Println(studyGroup)

	sort.Slice(studyGroup2, func(i int, j int) bool {
		return studyGroup2[i] >= studyGroup2[j]
	})

	fmt.Println(studyGroup2)

	sort.Slice(n, func(i int, j int) bool {
		return n[i] >= n[j]
	})

	fmt.Println(n)
}
