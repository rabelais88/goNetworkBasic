package statusCheck

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
)

type urlState struct {
	url   string
	state bool
}

func checkLink(link string, c chan<- urlState, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, `is not working`)
		c <- urlState{
			url:   link,
			state: false,
		}
	} else {
		fmt.Println(link, `is working`)
		c <- urlState{
			url:   link,
			state: true,
		}
	}
}

func Check(links []string) bool {
	result := true
	stats := make(map[string]bool)
	var wg sync.WaitGroup
	fmt.Println(`total sites: `, len(links))
	wg.Add(len(links)) // add workgroups of exactly same amount as links array length

	c := make(chan urlState, len(links))
	for _, link := range links {
		stats[link] = false
		fmt.Println(`requesting...`, link)
		go checkLink(link, c, &wg)
		fmt.Println(`goroutines: `, runtime.NumGoroutine())
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(`recv from channel, assigning result for `, v.url)
		stats[v.url] = v.state
		fmt.Println(`goroutines: `, runtime.NumGoroutine())
	}

	fmt.Println(`work finished!`)
	for key, value := range stats {
		fmt.Println(key, `---`, value)
		if value == false {
			result = false
		}
	}
	return result
}
