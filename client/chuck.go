package client

import (
	"encoding/json"
	"log"
	"net/http"
	"prueba/models"
	"sync"
)

func ClientChuck() (*models.Chucks, error) {

	// set the limit of concurrent requests
	c := make(chan int, 4)
	wg := sync.WaitGroup{}
	chucks := &models.Chucks{}
	lock := sync.Mutex{}

	for i := 0; i < 25; i++ {

		wg.Add(1)

		c <- 1
		go func(chucks *models.Chucks) {
			defer wg.Done()
			resp, err := http.Get("https://api.chucknorris.io/jokes/random")
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			chuck := &models.Chuck{}

			if err := json.NewDecoder(resp.Body).Decode(chuck); err != nil {
				log.Fatal(err)
			}
			lock.Lock()
			chucks.Chucks = append(chucks.Chucks, *chuck)
			lock.Unlock()
			<-c
		}(chucks)
	}

	wg.Wait()
	return chucks, nil
}
