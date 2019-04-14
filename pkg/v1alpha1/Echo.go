package v1alpha1

import (
	"time"
	"net/http"
	"encoding/json"
)

type LivyBody struct {
	File         string   `json:"file"`
	ClassName    string   `json:"className"`
	NumExecutors int      `json:"numExecutors"`
	Name         string   `json:"name"`
	Args         []string `json:"args"`
	CreatedAt    time.Time
}

// EchoHandler echoes the request
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	livyBody := LivyBody{} //initialize empty user

	//Parse json request body and use it to set fields on user
	//Note that user is passed as a pointer variable so that it's fields can be modified
	err := json.NewDecoder(r.Body).Decode(&livyBody)
	if err != nil {
		panic(err)
	}

	//Set CreatedAt field on user to current local time
	livyBody.CreatedAt = time.Now().Local()

	//Marshal or convert user object back to json and write to response
	userJson, err := json.Marshal(livyBody)
	if err != nil {
		panic(err)
	}

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response
	w.Write(userJson)
}