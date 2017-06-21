package main

import (
	"github.com/parnurzeal/gorequest"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
	"github.com/jhonata-menezes/kartolafc-backend/cmd"
)

func main() {
	var wg sync.WaitGroup
	inicio := time.Now()
	jobs, _ := strconv.Atoi(os.Args[2])
	loops, _ := strconv.Atoi(os.Args[3])
	count := 0
	for i := 0; i < jobs; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			request := gorequest.New()
			for j := 0; j < loops; j++ {
				resp, _, errs := request.Get(os.Args[1]).Set("Accept", "application/json, text/plain, */*").
					Set("Referer", "https://cartolafc.globo.com/").
					Set("Origin", "https://cartolafc.globo.com/").
					Set("Accept-Language", "pt-BR,pt;q=0.8,en-US;q=0.6,en;q=0.4,es;q=0.2").
					Set("X-GLB-Token", cmd.Config.TokenGlb).End()
				count++
				if errs == nil {
					log.Println(count, resp.StatusCode)
					resp.Body.Close()
				} else {
					log.Println(errs)
				}
			}
		}()
	}
	time.Sleep(1 * time.Second)
	wg.Wait()
	log.Println(time.Since(inicio))
}
