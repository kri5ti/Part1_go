package main

import (
	
	"log"
	"net/http"
	"os"
	"fmt"
	"encoding/json"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s $URL", os.Args[0])
	}
	result := make(map[string]string)
	resp, err := http.Get(os.Args[1])
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	err2 := dec.Decode(&result)
 		if err2 != nil {
			log.Fatal(err2)
			}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result["origin"])
	

	
}
