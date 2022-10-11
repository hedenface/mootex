package main

import (
	"fmt"
	"log"
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
	mootex.LogDebug("lock(%s)", key)
	mootex.Lock(key)
	mootex.LogDebug("%s was locked!", key)
	fmt.Fprintf(w, "%s", "locked")
}

func unlock(key string, w http.ResponseWriter) {
	mootex.LogDebug("unlock(%s)", key)
	mootex.Unlock(key)
	mootex.LogDebug("%s was unlocked!", key)
	fmt.Fprintf(w, "%s", "unlocked")
}

func getKeyParameter(u *url.URL) (string, error) {
	mootex.LogDebugln("getKeyParameter()")
	key := u.Query().Get("key")

	if key != "" {
		mootex.LogDebug("got key (%s)", key)
		return key, nil
	}

	return "", fmt.Errorf("%s", "no parameter named 'key' in url")
}

func lockParameter(w http.ResponseWriter, r *http.Request) {
	mootex.LogDebugln("lockParameter()")
	key, err := getKeyParameter(r.URL)

	if err == nil {
		lock(key, w)
	} else {
		mootex.LogError("%v", err)
	}
}

func unlockParameter(w http.ResponseWriter, r *http.Request) {
	mootex.LogDebugln("unlockParameter()")
	key, err := getKeyParameter(r.URL)

	if err == nil {
		unlock(key, w)
	} else {
		mootex.LogError("%v", err)
	}
}

func getPlaintextBody(r *http.Request) (string, error) {
	mootex.LogDebugln("getPlaintextBody()")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", fmt.Errorf("%s", "Unable to read request body")
	}

	mootex.LogDebug("body = [%s]", body)

	return string(body), nil
}

func lockPlaintextBody(w http.ResponseWriter, r *http.Request) {
	mootex.LogDebugln("lockPlaintextBody()")
	key, err := getPlaintextBody(r)

	if err == nil {
		mootex.LogDebug("got key (%s)", key)
		lock(key, w)
	} else {
		mootex.LogError("%v", err)
	}
}

func unlockPlaintextBody(w http.ResponseWriter, r *http.Request) {
	mootex.LogDebugln("unlockPlaintextBody()")
	key, err := getPlaintextBody(r)

	if err == nil {
		mootex.LogDebug("got key (%s)", key)
		unlock(key, w)
	} else {
		mootex.LogError("%v", err)
	}
}

func getJSONKey(r *http.Request) (string, error) {
	mootex.LogDebugln("getJSONKey()")
	body, err := getPlaintextBody(r)
	if err != nil {
		return "", err
	}

	mootex.LogDebug("body = [%s]", body)

	var JSONBody jsonBody

	err = json.Unmarshal([]byte(body), &JSONBody)
	if err != nil {
		return "", fmt.Errorf("%s", "json.Unmarshal error")
	}

	mootex.LogDebug("%v", JSONBody)

	if JSONBody.Key == "" {
		return "", fmt.Errorf("%s", "no key named 'key' in map")
	}

	return JSONBody.Key, nil
}

func lockJSONBody(w http.ResponseWriter, r *http.Request) {
	mootex.LogDebugln("lockJSONBody()")
	key, err := getJSONKey(r)

	if err == nil {
		mootex.LogDebug("got key (%s)", key)
		lock(key, w)
	} else {
		mootex.LogError("%v", err)
	}
}

func unlockJSONBody(w http.ResponseWriter, r *http.Request) {
	mootex.LogDebugln("unlockJSONBody()")
	key, err := getJSONKey(r)

	if err == nil {
		mootex.LogDebug("got key (%s)", key)
		unlock(key, w)
	} else {
		mootex.LogError("%v", err)
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

	log.Fatal(http.ListenAndServe(":3000", nil))
}
