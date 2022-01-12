package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Data struct {
	Source string `json:"source"`
	Title  string `json:"title"`
	Link   string `json:"link"`
	Id     int    `json:"id"`
}

func scrape(url string) ([]Data, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data []Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func separate(data []Data) []Data {
	var separated []Data
	for _, d := range data {
		separated = append(separated, d)
	}
	return separated
}

func drawBox(separated []Data) {
	fmt.Println("")
	for _, d := range separated {
		fmt.Println("")
		fmt.Println("Source:", d.Source)
		fmt.Println("Title:", d.Title)
		fmt.Println("Link:", d.Link)
		fmt.Println("Id:", d.Id)
		fmt.Println("")
	}
	fmt.Println("")
	for {
		fmt.Println("Enter an id or q to exit:")
		var input string
		fmt.Scanln(&input)
		if input == "q" {
			fmt.Println("")
			fmt.Println("Goodbye!")
			break
		} else {
			id, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("")
				fmt.Println("Invalid input, try again.")
				continue
			}
			for _, d := range separated {
				if d.Id == id {
					fmt.Println("")
					fmt.Println("Source:", d.Source)
					fmt.Println("Title:", d.Title)
					fmt.Println("Link:", d.Link)
					fmt.Println("Id:", d.Id)
					fmt.Println("")
				}
			}
		}
	}
}

// main function
func main() {
	url := "https://watchman-api.herokuapp.com/api/v1/news"
	data, err := scrape(url)
	if err != nil {
		fmt.Println(err)
	}
	separated := separate(data)
	drawBox(separated)
}
