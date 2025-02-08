package listener

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

var (
	url                = "https://stream.wikimedia.org/v2/stream/recentchange"
	mu                 sync.Mutex
	mostRecentChange   = make([]WikipediaChange, 0, 10)
	LanguageFilterChan = make(chan string)
)

func ListenToWikipediaChanges(wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	var currentFilter string

	go func() {
		for newFilter := range LanguageFilterChan {
			mu.Lock()
			currentFilter = newFilter
			mostRecentChange = mostRecentChange[:0]
			mu.Unlock()
			fmt.Println("Filter updated: ", newFilter)
		}
	}()

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "data: ") {
			jsonData := strings.TrimPrefix(line, "data: ")
			var change WikipediaChange
			if err := json.Unmarshal([]byte(jsonData), &change); err != nil {
				fmt.Println(err.Error())
				continue
			}

			mu.Lock()
			filter := currentFilter
			mu.Unlock()

			if change.Wiki != filter && filter != "" {
				continue
			}
			//fmt.Println(mostRecentChange)
			addChange(change)

		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())
	}
}
