package httpreq

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func CallHTTP() {

	// I. THE SIMPLE WAY
	resp, err := http.Get("https://reqres.in/api/users?page=2")
	if err != nil {
		log.Println(errors.New("failed to call endpoint"))
	}

	defer resp.Body.Close()
	body1, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(errors.New("failed to read response body"))
	}

	// Convert Byte to string: https://stackoverflow.com/questions/40632802/how-to-convert-byte-array-to-string-in-go
	// We can also json.unmarshal the body but we need to state the struct
	log.Println(string(body1))

	//========================================================================================================================
	log.Println("===========================================================================================================")

	// II. THE MORE IN-DEPTH OPTION
	// 1. Declare http client
	client := &http.Client{}

	// 2. Convert Request Body string to io.reader format
	reqBody, err := json.Marshal(map[string]string{
		"name": "kokondao",
		"job":  "pecundang",
	})
	if err != nil {
		log.Println(errors.New("failed to json marshal request body"))
	}

	req, err := http.NewRequest("POST", "https://reqres.in/api/users", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println(errors.New("failed to wrap request body"))
	}

	// 3. Add headers
	req.Header.Add("Content-Type", `application/json"`)
	req.Header.Add("Authorization", `Bearer bwhb3hb3h34b53hj5435b34h5b35ij"`)

	// 4. Hit request
	res, err := client.Do(req)
	if err != nil {
		log.Println(errors.New("failed to read response body"))
	}

	// 5. Read Response Body
	defer res.Body.Close()
	body2, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(errors.New("failed to read response body"))
	}

	log.Println(string(body2))
}
