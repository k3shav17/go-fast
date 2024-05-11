package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"github.com/schollz/progressbar/v2"
)

const (
	urlCount = 3                                  // url count defines the no of urls to be called to read a file to buffer
	token    = "YXNkZmFzZGxmbnNkYWZoYXNkZmhrYWxm" // this is a static token and one can find at https://fast.com/app-ab2f99.js
)

type TempFileLinks struct {
	Url string `json:"url"`
}

func main() {
	url := fmt.Sprintf("https://api.fast.com/netflix/speedtest?https=true&token=%s&urlCount=%d", token, urlCount)
	resp, er := http.Get(url)
	if er != nil {
		log.Fatalf("unable to read the url")
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("")
	}

	var tempFiles []TempFileLinks
	error := json.Unmarshal(data, &tempFiles)
	if len(tempFiles) == 0 {
		log.Fatalf("deentlo em ledhu ra gootley %s", error)
	}

	var multipleSpeeds []float64
	var sumOfSpeeds float64
	progressBar := progressbar.New(100)

	for _, file := range tempFiles {
		multipleSpeeds = append(multipleSpeeds, callThis(file.Url))
		progressBar.Add((progressBar.GetMax() / urlCount))
	}
	progressBar.Finish()

	for _, speed := range multipleSpeeds {
		sumOfSpeeds += speed
	}
	avgSpeed := sumOfSpeeds / float64(len(multipleSpeeds))
	fmt.Printf("\nspeed -> %.2f Mbps\n", avgSpeed)
}

func callThis(url string) float64 {
	chunkSize := 26214400
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("unable to read the url")
	}

	buffer := make([]byte, chunkSize) // creating a buffer of 25MB

	startTime := time.Now()
	_, err = io.ReadFull(resp.Body, buffer) // whatever is read from the response body is read into the buffer of size 25MB
	endTime := time.Now()
	if err != nil {
		log.Fatalf("can't write to buffer")
	}

	downloadRate := float64(chunkSize) / float64(endTime.Sub(startTime).Seconds())
	return (downloadRate / (1024 * 1024)) * 8 // returning the speed in Mege bits / second

}
