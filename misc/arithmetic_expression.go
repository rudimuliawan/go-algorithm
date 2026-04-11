package misc

import "fmt"

func calculate_expression(expression string) int {
	for _, c := range expression {
		fmt.Println(c)
	}
}
