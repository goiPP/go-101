package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://golang.org",
		"http://facebook.com",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// blocking line (call to channel and wait for the message from chan)
	// fmt.Println(<-c) // the main-Go-Routine will wait here until something happen then will exit. So we can see message only from first child-go-routine

	for l := range c { // wait for the channel to return a value and assign to 'l' then execute loop's body
		// time.Sleep(5 * time.Second) // pause routine for 5 sec, but this will pause main-Routine and we cannot take more request
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c) // l === '<- c'
		}(l) // we have to pass the link into literal function , so that child-Routine will consume copied of value in different address that main-routine is using
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
