package main

import (
	"fmt"
	"os"
)

// Subject описывает структура каждого предмета и набор параметров
type Subject struct {
	Name          string
	Hours         uint
	HoursVisited  uint
	ValueOfLesson float32
	IsExam        bool
	IsProfile     bool
	IsPersonalFav bool
}

// currentScore озвращает текущее число баллов
func (s *Subject) currentScore() float32 {
	return float32(s.HoursVisited) * s.ValueOfLesson
}

// maxScore возвращает максимальное число баллов
func (s *Subject) maxScore() float32 {
	return float32(s.Hours) * s.ValueOfLesson
}

func (s *Subject) MakeDecision() bool {
	p := s.currentScore() / s.maxScore()

	// Если посещено более 80 пар, то можно и пропустить (теоретически)
	if p > 0.8 {
		return false
	}

	// Если посещено более 70 пар и предмент не профильный, то можно и пропустить (теоретически)
	if p > 0.7 && !s.IsProfile {
		return false
	}

	// Если посещено более 60 пар и предмент не интересен, то можно и пропустить (теоретически)
	if p > 0.6 && !s.IsPersonalFav {
		return false
	}

	// Если посещено более 50 пар и будет зачет, то можно и пропустить (теоретически)
	if p > 0.5 && !s.IsExam {
		return false
	}

	// В противном случае нужно идти
	return true
}

// Слайс с предметами
var subjects = []Subject{
	{
		Name:          "Физическая культура",
		Hours:         53,
		HoursVisited:  11,
		ValueOfLesson: 0.37736,
		IsExam:        false,
		IsProfile:     false,
		IsPersonalFav: false,
	},
	{
		Name:          "Иностранный язык профессионального общения",
		Hours:         20,
		HoursVisited:  12,
		ValueOfLesson: 1,
		IsExam:        false,
		IsProfile:     false,
		IsPersonalFav: true,
	},
	{
		Name:          "Системное программирование",
		Hours:         21,
		HoursVisited:  2,
		ValueOfLesson: 0.95238,
		IsExam:        false,
		IsProfile:     true,
		IsPersonalFav: true,
	},
	{
		Name:          "Компьютерное моделирование",
		Hours:         29,
		HoursVisited:  9,
		ValueOfLesson: 1.42857,
		IsExam:        true,
		IsProfile:     true,
		IsPersonalFav: false,
	},
	{
		Name:          "Исследование операций и методы оптимизации",
		Hours:         14,
		HoursVisited:  9,
		ValueOfLesson: 0.95238,
		IsExam:        true,
		IsProfile:     true,
		IsPersonalFav: false,
	},
}

func main() {
	fmt.Println("Выбери интересующий предмет:")

	for i := range subjects {
		fmt.Printf("%d. %s\n", i, subjects[i].Name)
	}

	var inp int
	fmt.Print("> ")
	fmt.Scanf("%d", &inp)

	if inp >= len(subjects) || inp < 0 {
		fmt.Fprintln(os.Stderr, "Введено недопустимое число")
		return
	}

	if !subjects[inp].MakeDecision() {
		fmt.Println("Не пойду")
		return
	}

	fmt.Println("Нужно идти")
}
