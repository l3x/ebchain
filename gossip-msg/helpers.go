package main

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"time"
	"log"
	"math/rand"
	"strconv"
	"os"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}


func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}

func readLines(path string) ([]string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to read %", path))
	}
	return strings.Split(string(content), "\n")
}


func maybeDisableTimer() bool {
	disableTimerLines := readLines(myTimerFilePath)
	return stringInSlice("true", disableTimerLines)
}


func getIntevalSecs() (intSecs int) {
	intervalSecs := readLines(myIntervalFilePath)
	if len(intervalSecs) == 0 {
		log.Fatal(fmt.Sprintf("myIntervalFilePath (%s) does not have valid data.", myIntervalFilePath))
	} else {
		var err error
		intSecs, err = strconv.Atoi(intervalSecs[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	}
	return intSecs
}


func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max - min) + min
}