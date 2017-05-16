package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func produce(ch chan int) {
	i := 0
	for true {
		time.Sleep(time.Second)
		ch <- i
		i++
	}
}

//
func get_first(A int, group ...[]int) int {
	ch := make(chan int)
	for _, a := range group {
		go func(a []int) {
			for _, j := range a {
				if j > A {
					ch <- j
					break
				}
			}
		}(a)
	}
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	return 0
	// return <-ch
}

// func test(a ...[]int) {
// 	for i, j := range a {
// 		fmt.Println(i)
// 		fmt.Println(j)
// 	}
//
// }

func cc(a int) (b int) {
	b = a
	defer func(b *int) {
		*b = *b + *b
	}(&b)
	return
}

type gao struct {
	a int
	b int
	c int
}

// var sum1, sum2 int64
//
// func add1() {
// 	for i := 0; i < 1000; i++ {
// 		for j := 0; j < 10000; j++ {
// 			sum1++
// 		}
// 	}
// }
// type b struct{}

type test struct {
}

func (*test) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello")
}

func test_http() {
	mux := http.NewServeMux()
	mux.Handle("/h", &test{})
	http.ListenAndServe(":8080", mux)
}

func main() {
	test_http()

	// a := &gao{1, 2, 3}
	// fmt.Println((*a).a)
	// var result interface{}
	// result = 5
	// switch 10 {
	// case 10:
	// 	fmt.Println(result.(int))
	// }
	// fmt.Println(cc(10))
	// var ch chan int
	// ch = make(chan int)
	// ch <- 1
	// fmt.Println(<-ch)
	// fmt.Println("jiji")
	// go add1()
	// for i := 0; i < 1000; i++ {
	// 	for j := 0; j < 1000; j++ {
	// 		sum2++
	// 	}
	// }
	// go add1()
	// fmt.Println(sum1, sum2)
	// var a [3]int = [3]int{1, 2, 3}
	// A := []int{70, 89, 92}
	// B := []int{85, 69, 98, 90}
	// fmt.Println(get_first(90, A, B))

	// ch := make(chan int, 10)
	// var a [3]int = [3]int{1, 2, 3}
	// c := 0
	// go consume(ch)
	// go produce(ch)
	// c, err := <-ch
	// fmt.Println(c)
	// fmt.Println(err)
	// select {
	// case c1 := <-ch:
	// 	fmt.Println(c1)
	// 	// if err1 {
	// 	// fmt.Println("yes" + fmt.Sprintf("%d", c1))
	// 	// }
	// }

	// for true {
	// 	time.Sleep(time.Second)
	// 	a := <-ch
	// 	fmt.Println("consume" + fmt.Sprintf("%d", a))
	// }

	// if err {
	// fmt.Println(c)
	// }
	// for true {
	// 	time.Sleep(time.Second)
	// 	c++
	// 	ch <- c
	// 	for i := range ch {
	// 		fmt.Println(i)
	// 	}
	// }

	// go consume(ch)
	// for true {
	// time.Sleep(time.Second)
	// ch <- i
	// i++
	// select {
	// case ch <- a[i]:
	// i++
	// i %= 3
	// fmt.Println(i)
	// case <-ch:
	// fmt.Println(c)
	// fmt.Println(err)
	// }
	// }
	// time.Sleep(time.Second * 10)
}
func consume(ch chan int) {
	for true {
		time.Sleep(time.Second)
		a := <-ch
		fmt.Println("consume" + fmt.Sprintf("%d", a))
	}
}
