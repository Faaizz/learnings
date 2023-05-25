package block

import (
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

// Make request
func Request(baseUrl, jobPath string, pInterval, pCtr int) error {
	c := &http.Client{}

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

	for i := 0; i < pCtr; i++ {
		r, err := c.PostForm(baseUrl+jobPath, map[string][]string{"id": {rID}})
		if err != nil {
			return err
		}
		if r.StatusCode != http.StatusOK {
			return errors.New(r.Status)
		}

		defer r.Body.Close()
		sB, err := io.ReadAll(r.Body)
		if err != nil {
			return err
		}
		s := string(sB)

		log.Printf("response: %s", s)
		if s == "done" {
			log.Printf("done")
			return nil
		}

		pD := time.Duration(pInterval) * time.Second
		time.Sleep(pD)
	}

	return errors.New("request took too long")
}
