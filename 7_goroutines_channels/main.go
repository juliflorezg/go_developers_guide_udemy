package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// receiving messages from a channel is blocking, the main routine waits for this messages and THEN continues its execution

	//* now, we want to check constantly for the website status, so as long as this program is running it will spawn goroutines to check all sites status

	// this syntax is equivalent to
	/*
		    for {
				go checkLink(<-, c)
			}
	*/
	// and is saying that is going to wait for a value to come out of the channel and assign it to the 'l' variable, and then pass it to the checkLink function
	for l := range c {
    //! this line of code right here will make the main routine pause for 5 seconds but it will be blocked, so it cannot receive any other messages from channels while in sleep, so it's not the appropriate place to put the sleep
		time.Sleep(5 * time.Second)
		go checkLink(l, c)
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
