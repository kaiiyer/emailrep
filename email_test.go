package emailrep

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type people struct {
	Number int `json:"number"`
}

type GetWebRequest interface {
	FetchBytes(url string) []byte
}

type LiveGetWebRequest struct {
}

func (LiveGetWebRequest) FetchBytes(url string) []byte {
	emailrepClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "emailrep")

	res, getErr := emailrepClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	if body != nil {
		defer body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return body
}
