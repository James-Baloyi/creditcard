package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/validate/{creditCard}", ValidateCreditCard)
	http.ListenAndServe(":8080", r)
}

type ErrorStruct struct {
	Error string
}

type CalculatedStruct struct {
	IsValid bool
}

func ValidateCreditCard(w http.ResponseWriter, r *http.Request) {
	var toProcess bool
	vars := mux.Vars(r)
	id, notNull := vars["creditCard"]

	if notNull {
		toProcess = true
		if !checkNumeric(id) {
			toProcess = false
		}
	} else {
		toProcess = false
	}

	var isValid bool

	if toProcess {
		if creditCardValid(id) {
			isValid = true
		} else {
			isValid = false
		}

		res := &CalculatedStruct{
			IsValid: isValid,
		}
		content, _ := json.Marshal(res)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(content)

	} else {
		res := &ErrorStruct{
			Error: "Credit card number must be numeric",
		}

		content, _ := json.Marshal(res)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(content)

	}
}

func checkNumeric(i string) bool {
	_, err := strconv.ParseInt(i, 10, 0)
	return err == nil
}

func creditCardValid(card string) bool {
	n := len(card)
	if n < 13 || n > 19 {
		return false
	}

	var sum int
	for i := 0; i < n; i++ {
		digit, err := strconv.Atoi(string(card[i]))
		if err != nil {
			return false
		}

		if i%2 == n%2 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}
