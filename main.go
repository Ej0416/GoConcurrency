package main

import (
	"fmt"
)

// func sendEmail(message string) {
// 	go func() {
// 		time.Sleep(time.Millisecond * 250)
// 		fmt.Printf("Email received: '%s'\n", message)
// 	}()
// 	fmt.Printf("Email sent: '%s'\n", message)
// }

// // Don't touch below this line

// func test(message string) {
// 	sendEmail(message)
// 	time.Sleep(time.Millisecond * 500)
// 	fmt.Println("========================")
// }

// ------------------------------------ lesson 2

// type email struct {
// 	body string
// 	date time.Time
// }

// func checkEmailAge(emails [3]email) [3]bool {
// 	isOldChan := make(chan bool)

// 	go sendIsOld(isOldChan, emails)

// 	isOld := [3]bool{}
// 	isOld[0] = <-isOldChan
// 	isOld[1] = <-isOldChan
// 	isOld[2] = <-isOldChan
// 	return isOld
// }

// // don't touch below this line

// func sendIsOld(isOldChan chan<- bool, emails [3]email) {
// 	for _, e := range emails {
// 		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
// 			isOldChan <- true
// 			continue
// 		}
// 		isOldChan <- false
// 	}
// }

// ------------------------------------------------------ lesson 3

// func waitForDBs(numDBs int, dbChan chan struct{}) {
// 	for range numDBs{
// 		<-dbChan
// 	}
// }

// // don't touch below this line

// func getDBsChannel(numDBs int) (chan struct{}, *int) {
// 	count := 0
// 	ch := make(chan struct{})

// 	go func() {
// 		for i := range numDBs {
// 			ch <- struct{}{}
// 			fmt.Printf("Database %v is online\n", i+1)
// 			count++
// 		}
// 	}()

// 	return ch, &count
// }

// ------------------------------------------------- lesson 4

// func addEmailsToQueue(emails []string) chan string {
// 	emailMsg := make(chan string, len(emails))

// 	for _, em := range emails{
// 		emailMsg <- em
// 	}

// 	close(emailMsg)

// 	return emailMsg
// }

// ------------------------------------------------- lesson 5

// func countReports(numSentCh chan int) int {
// 	count := 0;

// 	for {
// 		ch, ok := <- numSentCh;
// 		if  !ok {
// 			break
// 		}
// 		count += ch
// 	}

// 	return count
// }

// // don't touch below this line

// func sendReports(numBatches int, ch chan int) {
// 	for i := range numBatches {
// 		numReports := i*23 + 32%17
// 		ch <- numReports
// 	}
// 	close(ch)
// }

// ------------------------------------------------- lesson 6

func concurrentFib(n int) []int {
	s := make(chan int)
	intSlice := make([]int,0)
	go fibonacci(n,s)

	for val := range s{
		intSlice = append(intSlice, val)
	}
	
	return intSlice
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	fmt.Println("app started")
}
