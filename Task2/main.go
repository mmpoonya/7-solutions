package main

import "fmt"

func main() {
	var input string
	fmt.Printf("Please type the encoded string: ")
	fmt.Scan(&input)
	result := decodeText(input)
	fmt.Printf("result=%s\n", result)
}

func decodeText(text string) string {
	min := 0
	buff := []int{0}
	for i := 0; i < len(text); i++ {
		switch text[i] {
		case 'L':
			buff = append(buff, buff[i]-1)
		case 'R':
			buff = append(buff, buff[i]+1)
		case '=':
			buff = append(buff, buff[i])
		default:
			return ""
		}
		if min > buff[i+1] {
			min = buff[i+1]
		}
	}

	result := ""
	for _, v := range buff {
		c := '0' + v - min
		result += fmt.Sprintf("%c", c)
	}
	return result
}
