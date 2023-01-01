package math

import "log"

func Add(a, b int) int {
	log.Printf("%d + %d = %d", a, b, a+b)
	return a + b
}
