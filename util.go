package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetUserId(r *http.Request) (id int, err error) {
	if idParams, ok := r.URL.Query()["user"]; ok {
		if id, err = strconv.Atoi(idParams[0]); err != nil {
			err = fmt.Errorf("'user' query parameter invalid")
			return
		}
		return
	}
	err = fmt.Errorf("'user' query parameter required")
	return
}
