package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	var pingUrl string
	flag.StringVar(&pingUrl, "u", "", "url to GET")
	var timeout int
	flag.IntVar(&timeout, "t", 5, "timeout between pings (seconds)")
	var jwtPath string
	flag.StringVar(&jwtPath, "j", "bin/xposer.jwt", "path to signed jwt")
	var testing bool
	flag.BoolVar(&testing, "test", false, "don't run forever if testing")
	flag.Parse()
	jwt, err := ioutil.ReadFile(jwtPath)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{Timeout: time.Duration(5) * time.Second}
	req, err := http.NewRequest("GET", pingUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "Bearer "+string(jwt))
	req.Header.Add("Accept-Encoding", "application/json")

	for {
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Failed to connect: %s", err)
		} else {
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Printf("Couldn't read body: %s", err)
				} else {
					log.Printf("Failed to connect: %s", body)
				}
			}
		}
		if testing {
			break
		}
		time.Sleep(time.Duration(timeout) * time.Second)
	}
}
