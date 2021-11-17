package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type ProcessResult struct {
	ServersNames [][]string
	ReplaceInfo  []string
}

func (pr ProcessResult) String() string {
	str := "ServersNames: [\n"
	for i := range pr.ServersNames {
		str += fmt.Sprintf("\t[%s]\n", strings.Join(pr.ServersNames[i], ","))
	}
	str += fmt.Sprintf("],\nReplaceInfo: [\n\t%s\n]", strings.Join(pr.ReplaceInfo, "\n\t"))

	return str
}

func main() {
	servers := []string{
		"US_MLRSP401", "EU_SZTPR1004", "US_MLRSP403", "AS_WPPR01",
		"RUS_MLP303", "US_MLRSP405", "US_MLRSP406", "AS_WPPR02",
		"US_MLRSP402", "AS_WPPR03", "EU_SZTPR1001", "EU_SZTPR1002",
		"US_MLRSP404", "EU_SZTPR1003", "RUS_MLP301",
	}

	res := solver(servers)
	fmt.Println(res)
}

// Solver just solves assignment
func solver(list []string) ProcessResult {
	var (
		res = ProcessResult{
			ServersNames: make([][]string, 0),
			ReplaceInfo:  make([]string, 0),
		}
		m             = make(map[string][]string)
		currentSymbol = '#' // Starts from ASCII's 35th symbol
	)

	// Create map from list
	// where key – region
	// 	 value – remaining info
	for i := range list {
		// splited[0] -> region, splited[1] -> group + number
		splited := strings.SplitN(list[i], "_", 2)

		if _, exist := m[splited[0]]; exist {
			m[splited[0]] = append(m[splited[0]], splited[1])
			continue
		}

		m[splited[0]] = []string{splited[1]}
	}

	// Process all map entries
	for k, v := range m {
		// If less than 3 servers then just sort it
		if len(v) < 3 {
			sort.Strings(v)

			// Concatenate region with group and number
			restoreNames(k, v)

			// And append it to res struct
			res.ServersNames = append(res.ServersNames, v)

			// Remove this from map
			delete(m, k)

			continue
		}

		idx := 0
		running := true
		for running {
			for i, j := 0, 1; j < len(v); i, j = i+1, j+1 {
				// Convert strings to rune slices
				runes1, runes2 := []rune(v[i]), []rune(v[j])

				// If out of bounds then stop processing
				if idx >= len(runes1) || idx >= len(runes2) {
					running = false
					break
				}

				// If reach any digit then stop processing
				if unicode.IsDigit(runes1[idx]) || unicode.IsDigit(runes2[idx]) {
					running = false
					break
				}

				if v[i][:idx] != v[j][:idx] {
					running = false
					break
				}
			}

			// Increment only if we still running
			if running {
				idx++
			}
		}

		// Save replaced sequence
		res.ReplaceInfo = append(res.ReplaceInfo, fmt.Sprintf("%s_%s = %c", k, v[0][:idx], currentSymbol))

		// Replace sequence in the lines
		for i := range v {
			v[i] = fmt.Sprintf("%c%s", currentSymbol, v[i][idx:])
		}

		// Save lines in res
		res.ServersNames = append(res.ServersNames, v)

		// Set next ASCII char
		currentSymbol += 1
	}

	return res
}

// Restore default server name format
func restoreNames(region string, list []string) {
	for i := range list {
		list[i] = fmt.Sprintf("%s_%s", region, list[i])
	}
}
