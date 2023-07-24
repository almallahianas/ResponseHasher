package hashing

import (
	"log"
	"sync"
)

func HandleCliConcurrently() {
	// input
	cliReader := CliReader{}
	iv, _ := cliReader.Read()

	cliWriter := CliWriter{}

	// process
	var wg sync.WaitGroup
	ch := make(chan *OutputVars, len(iv.Urls))
	go func() {
		// output
		for y := range ch {
			if y == nil {
				return
			}
			_ = cliWriter.Write(*y)
		}
	}()
	for batchStart := 0; batchStart < len(iv.Urls); batchStart += iv.MaxConcurrent {
		batchEnd := batchStart + iv.MaxConcurrent
		// handling index out of range for the inner loop
		if batchEnd > len(iv.Urls) {
			batchEnd = len(iv.Urls)
		}
		// process a batch
		for _, reqURL := range iv.Urls[batchStart:batchEnd] {
			wg.Add(1)
			go func(ch chan<- *OutputVars, reqURL string) {
				defer wg.Done()
				httpResponse, err := FetchUrl(reqURL)
				if err != nil {
					log.Printf("error:%v", err)
				}
				hashedResponse, _ := MD5HashFromResponse(httpResponse)
				ch <- &OutputVars{Url: reqURL, ResponseHash: hashedResponse}
			}(ch, reqURL)
		}
		wg.Wait()
	}
	ch <- nil
	close(ch)

}
