package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/CarlosEduardoNop/apphook/internal/models/calls"
)

func RetryRequest(url string, body []byte, attempt int, id int64) {
	timeout, _ := strconv.Atoi(os.Getenv("ATTEMPT_TIMEOUT"))

	time.Sleep(time.Duration(timeout) * time.Second)

	r, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		fmt.Println(err)
		return
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	resBody := make([]byte, 1024)
	res.Body.Read(resBody)

	if res.StatusCode != http.StatusOK {
		fmt.Println(res.Status)
		attempts := attempt + 1
		currentTime := time.Now()
		calls.Update(map[string]interface{}{"response": string(resBody), "attempts": attempt + 1, "updated_at": currentTime.Format("2006-01-02 15:04:05")}, map[string]interface{}{"id": id})
		attemptsLimit, _ := strconv.Atoi(os.Getenv("CALL_ATTEMPTS"))

		if attempts >= attemptsLimit {
			return
		}

		go RetryRequest(url, body, attempts, id)

		return
	}
}
