package main

import "fmt"

type Printer interface {
	Print(name string) string
}

func NewPrinter(lang string) Printer {
	switch lang {
	case "cn":
		return new(CnPrinter)
	case "en":
		return new(EnPrinter)
	default:
		return new(EnPrinter)
	}
}

type CnPrinter struct {
}

func (c *CnPrinter) Print(name string) string {
	return fmt.Sprintf("你好, %s\n", name)
}

type EnPrinter struct {
}

func (e *EnPrinter) Print(name string) string {
	return fmt.Sprintf("Hello, %s\n", name)
}

func main() {
	printer := NewPrinter("en")
	fmt.Println(printer.Print("wz"))
}
