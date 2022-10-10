package main

import (
	"fmt"
	//"log"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"

	"github.com/hedenface/mootex/pkg"
)

type jsonBody struct {
	Key string `json:"key"`
}

func lock(key string, w http.ResponseWriter) {
	fmt.Printf("lock(%s)\n", key)
	mootex.Lock(key)
	fmt.Fprintf(w, "%s", "locked")
}

func unlock(key string, w http.ResponseWriter) {
	fmt.Printf("unlock(%s)\n", key)
	mootex.Unlock(key)
	fmt.Fprintf(w, "%s", "unlocked")
}

func getKeyParameter(u *url.URL) (string, error) {
	fmt.Println("getKeyParameter()")
	key := u.Query().Get("key")

	if key != "" {
		fmt.Printf("got key (%s)\n", key)
		return key, nil
	}

	return "", fmt.Errorf("%s", "no parameter named 'key' in url")
}

func lockParameter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("lockParameter()")
	key, err := getKeyParameter(r.URL)

	if err == nil {
		fmt.Printf("got key (%s)\n", key)
		lock(key, w)
	} else {
		fmt.Printf("error: %v\n", err)
	}
}

func unlockParameter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("unlockParameter()")
	key, err := getKeyParameter(r.URL)

	if err == nil {
		fmt.Printf("got key (%s)\n", key)
		unlock(key, w)
	} else {
		fmt.Printf("error: %v\n", err)
	}
}

func getPlaintextBody(r *http.Request) (string, error) {
	fmt.Println("getPlaintextBody()")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", fmt.Errorf("%s", "Unable to read request body")
	}

	fmt.Printf("body = [%s]\n", body)

	return string(body), nil
}

func lockPlaintextBody(w http.ResponseWriter, r *http.Request) {
	fmt.Println("lockPlaintextBody()")
	key, err := getPlaintextBody(r)

	if err == nil {
		fmt.Printf("got key (%s)\n", key)
		lock(key, w)
	} else {
		fmt.Printf("error: %v\n", err)
	}
}

func unlockPlaintextBody(w http.ResponseWriter, r *http.Request) {
	fmt.Println("unlockPlaintextBody()")
	key, err := getPlaintextBody(r)

	if err == nil {
		fmt.Printf("got key (%s)\n", key)
		unlock(key, w)
	} else {
		fmt.Printf("error: %v\n", err)
	}
}

func getJSONKey(r *http.Request) (string, error) {
	fmt.Println("getJSONKey()")
	body, err := getPlaintextBody(r)
	if err != nil {
		return "", err
	}

	fmt.Printf("body = [%s]\n", body)

	var JSONBody jsonBody

	err = json.Unmarshal([]byte(body), &JSONBody)
	if err != nil {
		return "", fmt.Errorf("%s", "json.Unmarshal error")
	}

	fmt.Printf("%v\n", JSONBody)

	if JSONBody.Key == "" {
		return "", fmt.Errorf("%s", "no key named 'key' in map")
	}

	return JSONBody.Key, nil
}

func lockJSONBody(w http.ResponseWriter, r *http.Request) {
	fmt.Println("lockJSONBody()")
	key, err := getJSONKey(r)

	if err == nil {
		fmt.Printf("got key (%s)\n", key)
		lock(key, w)
	} else {
		fmt.Printf("error: %v\n", err)
	}
}

func unlockJSONBody(w http.ResponseWriter, r *http.Request) {
	fmt.Println("unlockJSONBody()")
	key, err := getJSONKey(r)

	if err == nil {
		fmt.Printf("got key (%s)\n", key)
		unlock(key, w)
	} else {
		fmt.Printf("error: %v\n", err)
	}
}

func main() {
	http.HandleFunc("/lock", lockParameter)
	http.HandleFunc("/unlock", unlockParameter)
	http.HandleFunc("/lockParameter", lockParameter)
	http.HandleFunc("/unlockParameter", unlockParameter)
	http.HandleFunc("/lockPlaintextBody", lockPlaintextBody)
	http.HandleFunc("/unlockPlaintextBody", unlockPlaintextBody)
	http.HandleFunc("/lockJSONBody", lockJSONBody)
	http.HandleFunc("/unlockJSONBody", unlockJSONBody)

	http.ListenAndServe(":3000", nil)
}
