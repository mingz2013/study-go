package main

import (
	"log"
)

type Table struct {
	i int
}

// 这里要用指针绑定方法，如果不用指针，方法每次会复制一份参数，包括接收者table也会复制一份，所以修改不生效，如果要修改原table，就要用指针
func (t *Table) addI() {
	log.Println("before", t.i)
	t.i += 1
	log.Println("end", t.i)
}

func main() {

	//ch:= make(chan int)

	table := Table{}

	//var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	for{
	//		select {
	//		case <-ch:
	//			table.addI()
	//		case <-time.After(1*time.Second):
	//
	//		}
	//	}
	//
	//}()

	//for i:=0;i<4;i++{
	//	ch<-i
	//}

	//wg.Wait()

	for i := 0; i < 4; i++ {
		table.addI()
	}

}
