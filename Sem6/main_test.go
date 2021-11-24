package main

import (
	"reflect"
	"testing"
)

func TestSliceToMap(t *testing.T) {
	cases := []struct {
		cards         []Card
		expectedValue map[string]Card
		expectedError error
	}{
		{
			cards:         nil,
			expectedValue: nil,
			expectedError: errConvertingSliceToMap,
		},
		{
			cards:         []Card{},
			expectedValue: nil,
			expectedError: errConvertingSliceToMap,
		},
		{
			cards: []Card{
				{From: "Москва", To: "Мельбурн"},
			},
			expectedValue: map[string]Card{
				"Москва": {From: "Москва", To: "Мельбурн"},
			},
			expectedError: nil,
		},
		{
			cards: []Card{
				{From: "Москва", To: "Мельбурн"},
				{From: "Мельбурн", To: "Париж"},
			},
			expectedValue: map[string]Card{
				"Москва":   {From: "Москва", To: "Мельбурн"},
				"Мельбурн": {From: "Мельбурн", To: "Париж"},
			},
			expectedError: nil,
		},
		{
			cards: []Card{
				{From: "Москва", To: "Мельбурн"},
				{From: "Мельбурн", To: "Париж"},
				{From: "Париж", To: "Воронеж"},
			},
			expectedValue: map[string]Card{
				"Москва":   {From: "Москва", To: "Мельбурн"},
				"Мельбурн": {From: "Мельбурн", To: "Париж"},
				"Париж":    {From: "Париж", To: "Воронеж"},
			},
			expectedError: nil,
		},
	}

	for _, tc := range cases {
		m, err := sliceToMap(tc.cards)

		if !reflect.DeepEqual(m, tc.expectedValue) || err != tc.expectedError {
			t.Errorf("sliceToMap(%v): Expected %v, %v; Got %v,%v\n", tc.cards, tc.expectedValue, tc.expectedError, m, err)
		}
	}
}

func TestFindStartCity(t *testing.T) {
	cases := []struct {
		m             map[string]Card
		expectedValue string
		expectedError error
	}{
		{
			m:             nil,
			expectedValue: "",
			expectedError: errEmptyMap,
		},
		{
			m:             make(map[string]Card),
			expectedValue: "",
			expectedError: errEmptyMap,
		},
		{
			m: map[string]Card{
				"Москва": {From: "Москва", To: "Мельбурн"},
			},
			expectedValue: "Москва",
			expectedError: nil,
		},
		{
			m: map[string]Card{
				"Москва":   {From: "Москва", To: "Мельбурн"},
				"Мельбурн": {From: "Мельбурн", To: "Москва"},
			},
			expectedValue: "",
			expectedError: errNoStartCity,
		},
		{
			m: map[string]Card{
				"Москва":   {From: "Москва", To: "Мельбурн"},
				"Париж":    {From: "Париж", To: "Воронеж"},
				"Мельбурн": {From: "Мельбурн", To: "Париж"},
			},
			expectedValue: "Москва",
			expectedError: nil,
		},
	}

	for _, tc := range cases {
		s, err := findStartCity(tc.m)

		if s != tc.expectedValue || err != tc.expectedError {
			t.Errorf("findStartCity(%v): Expected %v, %v; Got %v,%v\n", tc.m, tc.expectedValue, tc.expectedError, s, err)
		}
	}
}

func TestGetSortedCards(t *testing.T) {
	cases := []struct {
		cards         []Card
		expectedValue []Card
		expectedError error
	}{
		{
			cards:         nil,
			expectedValue: nil,
			expectedError: errNoCards,
		},
		{
			cards:         []Card{},
			expectedValue: nil,
			expectedError: errNoCards,
		},
		{
			cards: []Card{
				{From: "Москва", To: "Воронеж"},
			},
			expectedValue: []Card{
				{From: "Москва", To: "Воронеж"},
			},
			expectedError: nil,
		},
		{
			cards: []Card{
				{From: "Москва", To: "Воронеж"},
				{From: "Воронеж", To: "Москва"},
			},
			expectedValue: nil,
			expectedError: errNoStartCity,
		},
		{
			cards: []Card{
				{From: "Мельбурн", To: "Кельн"},
				{From: "Москва", To: "Париж"},
				{From: "Кельн", To: "Москва"},
			},
			expectedValue: []Card{
				{From: "Мельбурн", To: "Кельн"},
				{From: "Кельн", To: "Москва"},
				{From: "Москва", To: "Париж"},
			},
			expectedError: nil,
		},
		{
			cards: []Card{
				{From: "Мельбурн", To: "Кельн"},
				{From: "Кельн", To: "Москва"},
				{From: "Москва", To: "Париж"},
				{From: "Париж", To: "Амстердам"},
				{From: "Амстердам", To: "Лондон"},
			},
			expectedValue: []Card{
				{From: "Мельбурн", To: "Кельн"},
				{From: "Кельн", To: "Москва"},
				{From: "Москва", To: "Париж"},
				{From: "Париж", To: "Амстердам"},
				{From: "Амстердам", To: "Лондон"},
			},
			expectedError: nil,
		},
	}

	for _, tc := range cases {
		sorted, err := getSortedCards(tc.cards)

		if !reflect.DeepEqual(sorted, tc.expectedValue) || err != tc.expectedError {
			t.Errorf("getSortedCards(%v): Expected %v, %v; Got %v,%v\n", tc.cards, tc.expectedValue, tc.expectedError, sorted, err)
		}
	}
}
