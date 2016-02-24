package main

import (
	"log"
	"net/http"
	"os"

	"github.com/shopspring/decimal"
)

func balanceHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserId(r)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}

	db := GetDb()
	defer db.Close()

	var balance decimal.Decimal
	balance, err = GetBalance(db, userId)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}

	resp := map[string]decimal.Decimal{"balance": balance}
	JsonResponse(w, resp)
}

func depositHandler(w http.ResponseWriter, r *http.Request) {
	update, err := UnmarshalBalanceUpdate(r)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}

	db := GetDb()
	defer db.Close()

	if err := PerformBalanceUpdate(db, update.Amount, update.User); err != nil {
		http.Error(w, err.Error(), 422)
	}
}

func withdrawHandler(w http.ResponseWriter, r *http.Request) {
	update, err := UnmarshalBalanceUpdate(r)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}

	db := GetDb()
	defer db.Close()

	neg, _ := decimal.NewFromString("-1")
	if err := PerformBalanceUpdate(db, update.Amount.Mul(neg), update.User); err != nil {
		http.Error(w, err.Error(), 422)
	}
}

func tranferHandler(w http.ResponseWriter, r *http.Request) {
	trans, err := UnmarshalBalanceTransfer(r)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}

	db := GetDb()
	defer db.Close()

	if err := PerformBalanceTransfer(
		db,
		trans.Amount,
		trans.From,
		trans.To,
	); err != nil {
		http.Error(w, err.Error(), 422)
	}
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	user, err := UnmarshalUser(r)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}

	db := GetDb()
	defer db.Close()
	if err := CreateUser(db, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JsonResponse(w, user)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	db := GetDb()
	defer db.Close()

	users, err := GetUsers(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	JsonResponse(w, users)
}

func main() {
	var dsn string
	if len(os.Args) >= 2 {
		dsn = os.Args[1]
	} else {
		dsn = "password=goacc user=goacc dbname=goacc sslmode=disable"
	}
	if len(os.Args) >= 3 {
		if os.Args[2] == "createdb" {
			CreateDb()
			log.Println("Database created successfully.")
			os.Exit(0)
		} else {
			log.Fatal("Usage: goacc [dsn [createdb]]")
		}
	}

	SetDsn(dsn)
	log.Fatal(http.ListenAndServe(":8080", GetRouter()))
}
