package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getToken() {
	reqBody, err := json.Marshal(map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     dropify_client_id,
		"client_secret": dropify_client_secret,
	})

	if err != nil {
		log.Printf("[error] %s", err)
	}
	url := fmt.Sprintf("%s/%s", dropify_host, "oauth")
	// log.Printf("host: %s", url)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Printf("[error] %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[error] %s", err)
	}
	log.Printf("[debug] %s", string(body))
	// {"access_token":"3343009811c86944f8e811b2c6790655d212d087","expires_in":86400,"token_type":"Bearer","scope":null}

	// // Request product add.
	// client := &http.Client{}
	// req, err := http.NewRequest("POST", dropify_host +"/oauth", bytes.NewBuffer(reqBody))
	// req.Header.Set("Content-Type", "application/json")
	// HandleError(w, err)
	// req.SetBasicAuth(zunkaSiteUser(), zunkaSitePass())
	// res, err := client.Do(req)
	// HandleError(w, err)

	// // res, err := http.Post("http://localhost:3080/setup/product/add", "application/json", bytes.NewBuffer(reqBody))
	// defer res.Body.Close()
	// HandleError(w, err)

	// // Result.
	// resBody, err := ioutil.ReadAll(res.Body)
	// HandleError(w, err)
	// // No 200 status.
	// if res.StatusCode != 200 {
	// HandleError(w, errors.New(fmt.Sprintf("Error ao solicitar a criação do produto allnations no servidor zunka.\n\nstatus: %v\n\nbody: %v", res.StatusCode, string(resBody))))
	// return
	// }

}