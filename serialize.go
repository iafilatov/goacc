package main

import (
	"encoding/json"
	"net/http"
)

func UnmarshalUser(r *http.Request, user *User) (err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(user)
	return
}

func UnmarshalBalanceUpdate(r *http.Request, update *BalanceUpdate) (err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(update)
	return
}

func UnmarshalBalanceTransfer(r *http.Request, trans *BalanceTransfer) (err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(trans)
	return
}
