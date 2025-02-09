package wikipedia

import "strconv"

type WikipediaChange struct {
	Title     string `json:"title"`
	TitleURL  string `json:"title_url"`
	User      string `json:"user"`
	Timestamp int    `json:"timestamp"`
	Wiki      string `json:"wiki"`
}

func (wc WikipediaChange) String() string {
	return "Title: " + wc.Title + "\n" + "Url: " + wc.TitleURL + "\n" + "Author: " + wc.User + "\n" + "Timestamp: " + strconv.Itoa(wc.Timestamp) + "\n" +
		"Wiki: " + wc.Wiki + "\n"

}

func addChange(change WikipediaChange) {
	mu.Lock()
	defer mu.Unlock()

	if len(mostRecentChange) == capacity {
		mostRecentChange = mostRecentChange[1:]
	}

	mostRecentChange = append(mostRecentChange, change)
}

func GetRecentChanges() []WikipediaChange {
	mu.Lock()
	defer mu.Unlock()

	copiedChanges := make([]WikipediaChange, len(mostRecentChange))
	copy(copiedChanges, mostRecentChange)

	return copiedChanges

}
