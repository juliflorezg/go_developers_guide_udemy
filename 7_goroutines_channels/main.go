package main

import (
	"fmt"
	"net/http"
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
	for {
		go checkLink(<-c, c)
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
