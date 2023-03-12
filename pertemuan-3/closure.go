package main

import (
	"fmt"
	"sort"
	"strings"
)

type SearchData func(string) bool

var sorting = func(values []int) (ascend, descend []int) {
	ascend = sort.IntSlice(values)

	for idxI, i := range values {
		var temp int
		var listIdx []int

		temp = i

		for _, idx := range listIdx {
			if idxI == idx {
				continue
			} else {
				for idxJ, j := range values {
					if idxJ == idx {
						continue
					} else {
						if temp > j {
							temp = i
						} else {
							temp = j
						}
					}
				}
			}

			listIdx = append(listIdx, idxI)
		}

		descend = append(descend, temp)
	}

	return ascend, descend
}

func clasifier(items []string, search string) (int, func() []string) {
	var class []string
	for _, item := range items {
		if strings.Contains(item, search) {
			class = append(class, item)
		}
	}

	return len(items), func() []string {
		return class
	}
}

func selection(library []string, filter SearchData) map[string]string {
	var result = make(map[string]string)
	var bookIdx []int

	for idx, book := range library {
		if test := filter(book); test {
			bookIdx = append(bookIdx, idx)
			result["name"] = book
			result["count"] = fmt.Sprint(len(bookIdx))
		}
	}

	return result
}
