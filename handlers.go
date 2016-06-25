package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func DriverIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(drivers); err != nil {
		panic(err)
	}
}

func DriverShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// driverId := vars["driverId"]
	driverId, err := strconv.Atoi(vars["driverId"])
	if err != nil {
		panic(err)
	}

	driver := RepoFindDriver(driverId)
	// fmt.Fprintln(w, "Driver show:", driverId)
	if err := json.NewEncoder(w).Encode(driver); err != nil {
		panic(err)
	}
}

func DriverCreate(w http.ResponseWriter, r *http.Request) {
	var driver Driver
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &driver); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateDriver(driver)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
