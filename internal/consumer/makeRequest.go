package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/CarlosEduardoNop/apphook/internal/models/calls"
)

func MakeRequest(url string, body []byte, input map[string]interface{}) {
	if input["delay"] != 0 {
		time.Sleep(time.Duration(input["delay"].(float64)) * time.Second)
	}

	r, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		fmt.Println(err)
		return
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)

	currentTime := time.Now()

	call, _ := calls.Create(map[string]interface{}{
		"url":        input["url"],
		"event":      input["event"],
		"payload":    string(body),
		"attempts":   1,
		"delay":      input["delay"],
		"created_at": currentTime.Format("2006-01-02 15:04:05"),
		"updated_at": currentTime.Format("2006-01-02 15:04:05"),
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	resBody := make([]byte, 1024)
	res.Body.Read(resBody)

	id, _ := call.LastInsertId()

	if res.StatusCode != http.StatusOK {
		// fmt.Println(res.Status)
		calls.Update(map[string]interface{}{"response": string(resBody)}, map[string]interface{}{"id": id})

		go RetryRequest(url, body, 1, id)

		return
	}
}
