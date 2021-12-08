package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var (
	withAsteriskIsCompleted    = regexp.MustCompile(`[А-Яа-я-,.]+`)
	withoutAsteriskIsCompleted = regexp.MustCompile(`[А-Яа-я]+-*[А-Яа-я]*`)
)

func Top10(text string, taskWithAsteriskIsCompleted bool) []string {
	wordsMap := make(map[string]int)
	words := strings.Fields(text)

	for _, word := range words {
		re := withAsteriskIsCompleted
		if taskWithAsteriskIsCompleted {
			re = withoutAsteriskIsCompleted
		}
		res := re.FindAllString(word, -1)
		if len(res) > 0 {
			resWord := res[0]
			if re == withoutAsteriskIsCompleted {
				resWord = strings.ToLower(resWord)
			}
			wordsMap[resWord]++
		}
	}

	type Dict struct {
		Word   string
		Weight int
	}
	d := make([]Dict, 0)

	for word, weight := range wordsMap {
		d = append(d, Dict{word, weight})
	}
	sort.Slice(d, func(i, j int) bool {
		return (d[i].Weight == d[j].Weight && d[i].Word < d[j].Word) || d[i].Weight > d[j].Weight
	})
	var strArray []string
	if len(d) > 10 {
		for i := 0; i < 10; i++ {
			strArray = append(strArray, d[i].Word)
		}
	}
	return strArray
}
