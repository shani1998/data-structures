package main

import "fmt"

type DisplayStrategy interface {
	Display(msg string)
}

func NewDisplayStrategy() DisplayStrategy {
	return &DisplayFmt{}
}

type DisplayFmt struct{}

func (d *DisplayFmt) Display(msg string) {
	fmt.Println(msg)
}
