package main

func main() {
	// example #1
	// basic sending and receiving operations
	//ch := make(chan int)
	//go func() {
	//	ch <- 5 // sending
	//}()
	//a := <-ch
	//fmt.Println(a)

	// =====

	// example #2
	// receiving from closed channel
	//ch := make(chan int)
	//go func() {
	//	ch <- 6
	//}()
	//go func() {
	//	ch <- 5
	//}()
	//
	//a, ok := <-ch
	//
	//fmt.Println(a, ok)
	//fmt.Println(<-ch)
	//
	//a, ok = <-ch
	//if !ok {
	//	log.Fatal("Channel is closed")
	//}
	//close(ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)

	// =====

	// example #3
	// receiving more than sending
	//ch := make(chan int)
	//go func() {
	//	ch <- 5
	//}()
	//go func() {
	//	ch <- 6
	//}()
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)

	// =====

	// example #4
	// sending goroutine is blocked (stops) until the value is received from channel
	//ch := make(chan int)
	//go func() {
	//	ch <- 5
	//	fmt.Println("Anonymous func: - I sent the value")
	//}()
	//time.Sleep(time.Second * 3) // try to comment out
	//fmt.Println(<-ch)
	//fmt.Println("main func: - I received the value")
	//time.Sleep(time.Second) // try to comment out

	// =====

	// example #5
	// receiving goroutine is blocked (stops) until the value is sent to the channel
	//ch := make(chan int)
	//go func() {
	//	fmt.Println("Goroutine 1 started")
	//	fmt.Println("Goroutine 1 is working")
	//	time.Sleep(time.Second * 7)
	//	ch <- 5
	//	fmt.Println("Anonymous func 1: - I sent the value")
	//}()
	//
	////go func() {
	////	fmt.Println("Goroutine 2 started")
	////	fmt.Println(<-ch)
	////	fmt.Println("Goroutine 2 is working")
	////	fmt.Println("Anonymous func 2: - I received the value")
	////}()
	//fmt.Println("main func is working")
	//fmt.Println(<-ch)
	//fmt.Println("main func: - I received the value")
	//
	//time.Sleep(time.Second * 10)

	// =====

	// example #6
	// synchronizing goroutines (waiting for complete execution of goroutine)
	//done := make(chan bool)
	//
	//go func() {
	//	fmt.Println("Hello World!")
	//	time.Sleep(time.Second)
	//	done <- true
	//}()
	//
	//fmt.Println(<-done)
	//fmt.Println("exiting program")

	// =====

	// example #7
	// synchronizing goroutines (multiple)
	//done := make(chan int)
	//
	//go func() {
	//	fmt.Println("Hello World!")
	//	time.Sleep(time.Second)
	//	done <- 1
	//}()
	//
	//go func() {
	//	fmt.Println("Привет Мир!")
	//	time.Sleep(time.Second * 5)
	//	done <- 2
	//}()
	//
	//totalGoroutines := 2
	//for totalGoroutines > 0 {
	//	x := <-done
	//	if x == 1 {
	//		fmt.Println("Finished first goroutine")
	//		totalGoroutines--
	//	} else if x == 2 {
	//		fmt.Println("Finished second goroutine")
	//		totalGoroutines--
	//	}
	//}
	//
	//fmt.Println("exiting program")
}
