package client

import (
	"encoding/json"
	"log"
	"net/http"
	"prueba/models"
	"sync"
	"time"
)

func ClientChuck() (*models.Chucks, error) {

	wg := sync.WaitGroup{}
	chucks := &models.Chucks{}
	lock := sync.Mutex{}

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	for i := 0; i < 25; i++ {

		wg.Add(1)

		go func(chucks *models.Chucks) {
			defer wg.Done()
			resp, err := client.Get("https://api.chucknorris.io/jokes/random")
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
		}(chucks)
	}

	wg.Wait()
	return chucks, nil
}
