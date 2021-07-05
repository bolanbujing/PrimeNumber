package main

import "fmt"

func Processor(seq chan int, wait chan struct{}) {
	go func() {
		prime, ok := <-seq
		if !ok {
			close(seq)
			return
		}
		fmt.Println(prime)
		out := make(chan int)
		Processor(out, wait)
		for num := range seq {
			if num%prime != 0 {
				out <- num
			}
		}
		close(out)
	}()

}
func main() {
	orgin, wait := make(chan int), make(chan struct{})
	//hello github
	Processor(orgin, wait)
	for num := 2; num < 10000; num++ {
		orgin <- num
	}
	close(orgin)
	<-wait
}
