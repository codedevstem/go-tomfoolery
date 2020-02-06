package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var jsonStr = []byte(`{"name": "Kristian Hamre-Os", "email": "codedevstem@gmail.com", "source_code": "github.com/codedevstem"}`)
	req, err := http.NewRequest("POST", "http://10.200.233.89:8080/contest", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	
	for true {		resp, err := client.Do(req)

		if err != nil {
			fmt.Println("Heloo")
			fmt.Printf("%v", err)
		}
		defer resp.Body.Close()
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
