package main

import (
	"github.com/husobee/vestigo"
)

func GetRouter() (router *vestigo.Router) {
	router = vestigo.NewRouter()
	router.Get("/balance", balanceHandler)
	router.Post("/deposit", depositHandler)
	router.Post("/withdraw", withdrawHandler)
	router.Post("/transfer", tranferHandler)
	router.Post("/create", createHandler)
	router.Get("/users", usersHandler)
	return
}
