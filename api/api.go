package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var url = "https://stream.wikimedia.org/v2/stream/recentchange"

type WikipediaChange struct {
	Title    string `json:"title"`
	TitleURL string `json:"title_url"`
	User     string `json:"user"`
	Type     string `json:"type"`
	Wiki     string `json:"wiki"`
}

func ListenToWikipediaChanges() {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "data: ") {
			jsonData := strings.TrimPrefix(line, "data: ")
			var change WikipediaChange
			if err := json.Unmarshal([]byte(jsonData), &change); err != nil {
				fmt.Println(err.Error())
				continue
			}

			fmt.Println("Title:", change.Title)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())
	}
}
