package main

type Message struct {
	Value    float64
	List     int
	Position int
}

func FindAllGreaterThen(data [][]float64, targetNum float64) []Message {

	var greaterNumbers []Message

	for listCount, list := range data {
		for idx, num := range list {
			if num > targetNum {
				greaterNumbers = append(greaterNumbers, Message{
					Value:    num,
					List:     listCount + 1,
					Position: idx + 1,
				})
			}

		}
	}

	return greaterNumbers
}
