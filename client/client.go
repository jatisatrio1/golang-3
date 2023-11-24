package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

func hitAPI() {
	water := rand.Int()
	wind := rand.Int()

	requestBody := map[string]int{
		"water": water,
		"wind":  wind,
	}

	marshalReq, _ := json.Marshal(requestBody)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:1323/weather/", bytes.NewBuffer(marshalReq))

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
}
