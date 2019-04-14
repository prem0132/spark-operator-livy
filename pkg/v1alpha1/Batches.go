package v1alpha1

import (
	"net/http"
	"encoding/json"

)

type LivyAppBody struct {
	File         string   `json:"file"`
	ClassName    string   `json:"className"`
	NumExecutors int      `json:"numExecutors"`
	Name         string   `json:"name"`
	Args         []string `json:"args"`
}

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}


// CreateBatch handles the /batches api
func Batches(w http.ResponseWriter, r *http.Request) {
	//user := User{} //initialize empty user
	//
	////Parse json request body and use it to set fields on user
	////Note that user is passed as a pointer variable so that it's fields can be modified
	//err := json.NewDecoder(r.Body).Decode(&user)
	//if err != nil {
	//    panic(err)
	//}
	//
	////Set CreatedAt field on user to current local time
	//user.CreatedAt = time.Now().Local()
	//
	////Marshal or convert user object back to json and write to response
	//userJson, err := json.Marshal(user)
	//if err != nil {
	//    panic(err)
	//}
	//
	////Set Content-Type header so that clients will know how to read response
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	////Write json response back to response
	//w.Write(userJson)
	if r.URL.Path != "/batches" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":

		//t := T{}

		livybody := LivyAppBody{}
		err := json.NewDecoder(r.Body).Decode(&livybody)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("postBodylB:\n",lb.className)

		//livybody.CreatedAt = time.Now().Local()

		respdata, err := json.Marshal(livybody)
		if err != nil {
			panic(err)
		}

		//fmt.Printf("postBodylB:\n",respdata)

		//Set Content-Type header so that clients will know how to read response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		//Write json response back to response
		w.Write(respdata)

		//err1 := yaml.Unmarshal([]byte(yamldata), &t)
		//if err1 != nil {
		//    log.Fatalf("error: %v", err1)
		//}
		//fmt.Printf("--- t:\n%v\n\n", t)
		//
		//d, err := yaml.Marshal(&t)
		//if err != nil {
		//    log.Fatalf("error: %v", err)
		//}
		//fmt.Printf("--- t dump:\n%s\n\n", string(d))
		//
		//m := make(map[interface{}]interface{})
		//
		//err = yaml.Unmarshal([]byte(yamldata), &m)
		//if err != nil {
		//    log.Fatalf("error: %v", err)
		//}
		//fmt.Printf("--- m:\n%v\n\n", m)
		//
		//d, err = yaml.Marshal(&m)
		//if err != nil {
		//    log.Fatalf("error: %v", err)
		//}
		//fmt.Printf("--- m dump:\n%s\n\n", string(d))
	}
}