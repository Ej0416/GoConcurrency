package main

import (
	"fmt"
	"time"
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

// func concurrentFib(n int) []int {
// 	s := make(chan int)
// 	intSlice := make([]int,0)
// 	go fibonacci(n,s)

// 	for val := range s{
// 		intSlice = append(intSlice, val)
// 	}

// 	return intSlice
// }

// // don't touch below this line

// func fibonacci(n int, ch chan int) {
// 	x, y := 0, 1
// 	for range n {
// 		ch <- x
// 		x, y = y, x+y
// 	}
// 	close(ch)
// }

// -------------------------------------------- lesson 7

// func logMessages(chEmails, chSms chan string) {
// 	for{
// 		select{
// 		case sms, ok := <-chSms:
// 			if !ok{
// 				return
// 			}
// 			logSms(sms)

// 		case email, ok := <- chEmails:
// 			if !ok{
// 				return
// 			}
// 			logEmail(email)
// 		}
// 	}
// }

// // don't touch below this line

// func logSms(sms string) {
// 	fmt.Println("SMS:", sms)
// }

// func logEmail(email string) {
// 	fmt.Println("Email:", email)
// }

// func test(sms []string, emails []string) {
// 	fmt.Println("Starting...")

// 	chSms, chEmails := sendToLogger(sms, emails)

// 	logMessages(chEmails, chSms)
// 	fmt.Println("===============================")
// }

// func main() {
// 	fmt.Println("app started")
// 		test(
// 		[]string{
// 			"hi friend",
// 			"What's going on?",
// 			"Welcome to the business",
// 			"I'll pay you to be my friend",
// 		},
// 		[]string{
// 			"Will you make your appointment?",
// 			"Let's be friends",
// 			"What are you doing?",
// 			"I can't believe you've done this.",
// 		},
// 	)
// 	test(
// 		[]string{
// 			"this song slaps hard",
// 			"yooo hoooo",
// 			"i'm a big fan",
// 		},
// 		[]string{
// 			"What do you think of this song?",
// 			"I hate this band",
// 			"Can you believe this song?",
// 		},
// 	)
// }

// func sendToLogger(sms, emails []string) (chSms, chEmails chan string) {
// 	chSms = make(chan string)
// 	chEmails = make(chan string)
// 	randReader := rand.New(rand.NewSource(0))
// 	go func() {
// 		for i := 0; i < len(sms) && i < len(emails); i++ {
// 			done := make(chan struct{})
// 			s := sms[i]
// 			e := emails[i]
// 			t1 := time.Millisecond * time.Duration(randReader.Intn(1000))
// 			t2 := time.Millisecond * time.Duration(randReader.Intn(1000))
// 			go func() {
// 				time.Sleep(t1)
// 				chSms <- s
// 				done <- struct{}{}
// 			}()
// 			go func() {
// 				time.Sleep(t2)
// 				chEmails <- e
// 				done <- struct{}{}
// 			}()
// 			<-done
// 			<-done
// 			time.Sleep(10 * time.Millisecond)
// 		}
// 		close(chSms)
// 		close(chEmails)
// 	}()
// 	return chSms, chEmails
// }

// ------------------------------------------------- lesson 8

// func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
// 	for{
// 		select {
// 			case _, ok := <-snapshotTicker:
// 				if ok {
// 					takeSnapshot(logChan)
// 				}

// 			case _,ok := <- saveAfter:
// 				if ok{
// 					saveSnapshot(logChan)
// 					return
// 				}

// 			default:
// 				waitForData(logChan)
// 				time.Sleep(500 * time.Millisecond)
// 		}
// 	}
// }

// don't touch below this line

// func takeSnapshot(logChan chan string) {
// 	logChan <- "Taking a backup snapshot..."
// }

// func saveSnapshot(logChan chan string) {
// 	logChan <- "All backups saved!"
// 	close(logChan)
// }

// func waitForData(logChan chan string) {
// 	logChan <- "Nothing to do, waiting..."
// }

// --------------------------------------------- lesson 11

func pingPong(numPings int) {
	pings := make(chan struct{})
	pongs := make(chan struct{})

	go ponger(pings, pongs)
	go pinger(pings, numPings)

	func() {
		i := 0
		for range pongs {
			fmt.Println("got pong", i)
			i++
		}
		fmt.Println("pongs done")
	}()
}

// don't touch below this line
func pinger(pings chan struct{}, numPings int) {
	sleepTime := 50 * time.Millisecond
	for i := 0; i < numPings; i++ {
		fmt.Printf("sending ping %v\n", i)
		pings <- struct{}{}
		time.Sleep(sleepTime)
		sleepTime *= 2
	}
	close(pings)
}

func ponger(pings, pongs chan struct{}) {
	i := 0
	for range pings {
		fmt.Printf("got ping %v, sending pong %v\n", i, i)
		pongs <- struct{}{}
		i++
	}
	fmt.Println("pings done")
	close(pongs)
}

func test(numPings int) {
	fmt.Println("Starting game...")
	pingPong(numPings)
	fmt.Println("===== Game over =====")
}

func main() {
	fmt.Println("app started")
	test(4)
	test(3)
	test(2)
}
