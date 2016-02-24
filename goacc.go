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
	var update BalanceUpdate
	if err := UnmarshalBalanceUpdate(r, &update); err != nil {
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
	var update BalanceUpdate
	if err := UnmarshalBalanceUpdate(r, &update); err != nil {
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
	var trans BalanceTransfer
	if err := UnmarshalBalanceTransfer(r, &trans); err != nil {
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
	var user User
	if err := UnmarshalUser(r, &user); err != nil {
		http.Error(w, err.Error(), 422)
		return
	}

	db := GetDb()
	defer db.Close()
	if err := CreateUser(db, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JsonResponse(w, user)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	db := GetDb()
	defer db.Close()

	sqlGetAll := "SELECT id, name, balance FROM users"
	rows, err := db.Query(sqlGetAll)
	defer rows.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	users := []User{}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Name, &user.Balance); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		users = append(users, user)
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
