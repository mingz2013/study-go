package main

import "fmt"

type Aer interface{
    Name()string
    PrintName()
}

type A struct {
}

func (a *A) Name() string {
    return "a"
}

func (a *A) PrintName() {
    fmt.Println(a.Name())
}

type B struct {
    A
}

func (b *B) Name() string {
    return "b"
}

func getAer() Aer {
    return &B{}
}

func main() {
    a := getAer()
	a.PrintName() // a
}
