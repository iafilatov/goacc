package main

import (
	"encoding/json"
	"net/http"
)

func UnmarshalUser(r *http.Request) (user *User, err error) {
	user = &User{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(user)
	return
}

func UnmarshalBalanceUpdate(r *http.Request) (update *BalanceUpdate, err error) {
	update = &BalanceUpdate{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(update)
	return
}

func UnmarshalBalanceTransfer(r *http.Request) (trans *BalanceTransfer, err error) {
	trans = &BalanceTransfer{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(trans)
	return
}
