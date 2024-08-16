package data

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
)

type Store struct {
	Date map[string]int
	Mux  sync.Mutex
}

var values = Store{make(map[string]int), sync.Mutex{}}
var sep = regexp.MustCompile(`[ \n,!?;.]`)
var saveFile = "src/info/data.json"
var credentialeFile = "src/info/credentials.json"
var credentiale = make(map[string]string)
var counterPosts = 0

func StoreData(text string) {
	counterPosts++
	values.Mux.Lock()
	cuvinte := sep.Split(text, -1)
	for _, cuvant := range cuvinte {
		if len(cuvant) > 0 {
			_, exists := values.Date[cuvant]
			if exists {
				values.Date[cuvant]++
			} else {
				values.Date[cuvant] = 1
			}
		}
	}
	fmt.Println(values.Date)
	values.Mux.Unlock()
	if counterPosts == 20 {
		counterPosts = 0
		DumpData()
	}

}

func GetData(text string) map[string]int {
	cuvinte := strings.Split(text, "_")
	res := make(map[string]int)
	for _, cuvant := range cuvinte {
		if len(cuvant) > 0 {
			val, exists := values.Date[cuvant]
			if exists {
				res[cuvant] = val
			} else {
				res[cuvant] = 0
			}
		}
	}
	return res
}

func InitData() {
	data, err := os.ReadFile(saveFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	err = json.Unmarshal(data, &values.Date)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	jsoncredentials, err := os.ReadFile(credentialeFile)
	if err != nil {
		fmt.Println("Error opening credentials file:", err)
		return
	}

	err = json.Unmarshal(jsoncredentials, &credentiale)
	if err != nil {
		fmt.Println("Error unmarshalling credentails JSON:", err)
		return
	}

}

func DumpData() {
	jsondump, _ := json.Marshal(values.Date)
	file, err := os.OpenFile(saveFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	_, err = file.Write(jsondump)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("JSON data written to file successfully!")
	defer file.Close()
}

func CheckCredentials(username, password string) bool {
	val, exists := credentiale[username]
	if exists && val == password {
		return true
	}

	return false
}
