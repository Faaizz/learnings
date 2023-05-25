package noblock

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

// reference: https://stackoverflow.com/a/22892986
func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq() string {
	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return strings.ToUpper(string(b))
}

// Make request
func Request(baseUrl, jobPath string) error {
	c := &http.Client{
		// timeout must be set to prevent the client from waiting forever
		Timeout: 25 * time.Second,
	}

	initR, err := c.Get(baseUrl + jobPath)
	if err != nil {
		return err
	}

	defer initR.Body.Close()
	rIDBytes, err := io.ReadAll(initR.Body)
	if err != nil {
		return err
	}
	rID := string(rIDBytes)

	go func() {
		r, err := c.PostForm(baseUrl+jobPath, map[string][]string{"id": {rID}})
		if err != nil {
			panic(err)
		}
		if r.StatusCode != http.StatusOK {
			panic(r.Status)
		}

		defer r.Body.Close()
		sB, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		s := string(sB)

		log.Printf("response: %s", s)
		if s == "done" {
			log.Printf("done")
			os.Exit(0)
		}

		panic("request failed")
	}()

	for {
		sD := time.Duration(rand.Intn(5)) * time.Second
		time.Sleep(sD)
		log.Printf("doing some random stuff... " + randSeq())
	}
}
