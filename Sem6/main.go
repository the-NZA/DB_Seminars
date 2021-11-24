package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	errNoCards              = errors.New("Cards must contain at least 1 card")
	errNoStartCity          = errors.New("Can't find out start city")
	errEmptyMap             = errors.New("Map of cities must have at least 1 item")
	errConvertingSliceToMap = errors.New("Can't convert empty slice to map")
)

type Card struct {
	From string
	To   string
}

func (c Card) String() string {
	return fmt.Sprintf("%s -> %s", c.From, c.To)
}

var rawCards = []Card{
	{From: "Мельбурн", To: "Кельн"},
	{From: "Москва", To: "Париж"},
	{From: "Кельн", To: "Москва"},
	{From: "Милан", To: "Мельбурн"},
	{From: "Воронеж", To: "Милан"},
}

func main() {
	sorted, err := getSortedCards(rawCards)
	if err != nil {
		log.Fatalln(err)
	}

	// Display sorted cards
	for i := 0; i < len(sorted); i++ {
		fmt.Print(sorted[i])

		if (i + 1) < len(sorted) {
			fmt.Print(", ")
		} else {
			fmt.Println()
		}

	}
}

// get sorted cards | Complexity O(n)
func getSortedCards(cards []Card) ([]Card, error) {
	cLen := len(cards)

	if cards == nil || cLen < 1 {
		return nil, errNoCards
	}

	if cLen == 1 {
		return cards, nil
	}

	// convert slice to map
	m, err := sliceToMap(cards)
	if err != nil {
		return nil, err
	}

	// get first city
	start, err := findStartCity(m)
	if err != nil {
		return nil, err
	}

	// create slice for sorted cards with start city
	sortedCards := []Card{
		m[start],
	}

	for i := 1; i < cLen; i++ {
		sortedCards = append(sortedCards, m[sortedCards[i-1].To])
	}

	return sortedCards, nil
}

// convert slice cards to map | complexity O(n)
func sliceToMap(cards []Card) (map[string]Card, error) {
	cLen := len(cards)
	if cards == nil || cLen < 1 {
		return nil, errConvertingSliceToMap
	}

	m := make(map[string]Card)

	for i := range cards {
		m[cards[i].From] = cards[i]
	}

	return m, nil
}

// find start city | complexity O(n)
func findStartCity(m map[string]Card) (string, error) {
	var startCity string
	mLen := len(m)

	if m == nil || mLen < 1 {
		return "", errEmptyMap
	}

	// only one iteration
	if mLen == 1 {
		for k := range m {
			startCity = k
			return startCity, nil
		}
	}

	// create set (map of string as keys with empty struct as a value)
	set := make(map[string]struct{}, mLen)

	// fill set with map's keys
	for k := range m {
		set[k] = struct{}{}
	}

	// remove all
	for _, v := range m {
		delete(set, v.To)
	}

	// if there is less or more than one item left
	if len(set) != 1 {
		return "", errNoStartCity
	}

	// only one iteration
	for k := range set {
		startCity = k
	}

	return startCity, nil
}
