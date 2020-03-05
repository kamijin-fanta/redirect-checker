package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func writeRow(url string, status int, redirectTo string, message string) {
	fmt.Printf("%s\t%d\t%s\t%s\n", url, status, redirectTo, message)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		url := scanner.Text()
		client := &http.Client{}
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
		req, _ := http.NewRequest("GET", url, nil)
		resp, err := client.Do(req)
		if err != nil {
			writeRow(url, 0, "", err.Error())
		} else {
			writeRow(url, resp.StatusCode, resp.Header.Get("Location"), "")
		}
	}
}
