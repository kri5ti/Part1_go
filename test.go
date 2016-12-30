package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"encoding/json"
)

func main() {

type js struct {
Name string `json:"title"`
Count int `json:"lines"`
}
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	s := string(robots[0:len(robots)])
	cn := strings.Count(s, "\n")

	r := js{Name:"robots", Count: cn }
	b, err := json.Marshal(r)

	fmt.Println(string(b))
	
}