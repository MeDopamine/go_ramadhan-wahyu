package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	total := 0
	for _, score := range s.score {
		total += score
	}
	return float64(total) / float64(len(s.score))
}

func (s Student) Min() (int, string) {
	minScore := s.score[0]
	minIndex := 0

	for i := 1; i < len(s.score); i++ {
		if s.score[i] < minScore {
			minScore = s.score[i]
			minIndex = i
		}
	}

	return minScore, s.name[minIndex]
}

func (s Student) Max() (int, string) {
	maxScore := s.score[0]
	maxIndex := 0

	for i := 1; i < len(s.score); i++ {
		if s.score[i] > maxScore {
			maxScore = s.score[i]
			maxIndex = i
		}
	}

	return maxScore, s.name[maxIndex]
}

func main() {
	var students Student

	for i := 0; i < 3; i++ {
		var name string
		fmt.Printf("Input Student's Name %d: ", i)
		fmt.Scan(&name)
		students.name = append(students.name, name)

		var score int
		fmt.Printf("Input Student's Score %d: ", i)
		fmt.Scan(&score)
		students.score = append(students.score, score)
	}

	fmt.Printf("\nAverage Score of Students is %.2f\n", students.Average())

	minScore, minName := students.Min()
	fmt.Printf("Min Score of Students is: %s (%d)\n", minName, minScore)

	maxScore, maxName := students.Max()
	fmt.Printf("Max Score of Students is: %s (%d)\n", maxName, maxScore)
}
