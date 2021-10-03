package httpreq

import (
	"bytes"
	"encoding/json"
	"errors"
	"gobasics/entity"
	"io/ioutil"
	"log"
	"net/http"
)

// NOTE: THIS SECTION REQUIRES THREEDEE TO RUN

// I. THE SIMPLE SYNTAX
func CallHttpGet() {

	// 1. Call the request directly using the REST verb method of http package
	resp, err := http.Get("http://localhost:3000/print-requests/1")
	if err != nil {
		log.Println(errors.New("failed to call endpoint"))
	}

	// 2. Read the response body. Dont forget to close it afterwards
	defer resp.Body.Close()
	body1, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(errors.New("failed to read response body"))
	}

	// 3. Show response
	// a) By Converting Byte to string: https://stackoverflow.com/questions/40632802/how-to-convert-byte-array-to-string-in-go
	log.Println(string(body1))

	// b) We can also json.unmarshal the body but we need to state the struct (NEEDS NESTED STRUCT; See entity.HttpResponse)
	var receivedObj entity.HttpResponse
	err = json.Unmarshal(body1, &receivedObj)
	if err != nil {
		log.Println(err)
	}
	log.Println(receivedObj)
}

// II. THE IN-DEPTH SYNTAX
func CallHttpPost() {

	// 1. Declare http client
	client := &http.Client{}

	// 2. Convert Request Body string to io.reader format
	object := entity.NewPrintRequest("Tablet holder", 150, 20000, 32000, "http://drive.google.com/file/1789", "Karim", "received")

	reqBody, err := json.Marshal(object)
	if err != nil {
		log.Println(errors.New("failed to json marshal request body"))
	}

	req, err := http.NewRequest("POST", "http://localhost:3000/print-requests", bytes.NewBuffer(reqBody))
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
