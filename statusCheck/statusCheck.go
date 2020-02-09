package statusCheck

import (
	"fmt"
	"net/http"
	"sync"
)

func checkLink(link string) bool {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, `is not working`)
		return false
	}
	fmt.Println(link, `is working`)
	return true
}

func Check(links []string) bool {
	result := true
	stats := make(map[string]bool)
	var wg sync.WaitGroup
	wg.Add(len(links))
	for _, link := range links {
		stats[link] = checkLink(link)
		wg.Done()
	}
	wg.Wait()
	for key, value := range stats {
		fmt.Println(key, `---`, value)
		if value == false {
			result = false
		}
	}
	return result
}
