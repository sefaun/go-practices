package httprequest

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"practice/types"
)

func Login(response chan types.LoginResponse) {

	values := map[string]string{"email": "email", "password": "password"}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://example.com/api/v1/login", "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var res types.LoginResponse

	json.NewDecoder(resp.Body).Decode(&res)

	response <- res
	close(response)
}
