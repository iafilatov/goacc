package main

import (
	"database/sql"
	"fmt"

	"github.com/shopspring/decimal"
)

func PerformBalanceUpdate(db *sql.DB, amount decimal.Decimal, userId int) (err error) {
	var tx *sql.Tx
	if tx, err = db.Begin(); err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	return BalanceUpdateTx(tx, amount, userId)
}

func BalanceUpdateTx(tx *sql.Tx, amount decimal.Decimal, userId int) (err error) {
	var (
		zero          decimal.Decimal
		newBalanceRaw []uint8
		newBalance    decimal.Decimal
	)

	zero, _ = decimal.NewFromString("0")
	if amount.Cmp(zero) == 0 {
		return fmt.Errorf("Zero amount")
	}

	if err = tx.QueryRow(
		"UPDATE users SET balance = balance + $2 WHERE id = $1 RETURNING balance",
		userId,
		amount,
	).Scan(&newBalanceRaw); err != nil {
		return
	}
	newBalance, _ = decimal.NewFromString(string(newBalanceRaw))
	if newBalance.Cmp(zero) == -1 {
		err = fmt.Errorf("Insufficient balance")
		return
	}
	return
}

func PerformBalanceTransfer(db *sql.DB, amount decimal.Decimal, idFrom, idTo int) (err error) {
	var tx *sql.Tx
	if tx, err = db.Begin(); err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	err = BalanceUpdateTx(tx, amount, idTo)
	neg, _ := decimal.NewFromString("-1")
	err = BalanceUpdateTx(tx, amount.Mul(neg), idFrom)
	return
}

func CreateUser(db *sql.DB, user *User) (err error) {
	_, err = db.Exec(
		"INSERT INTO users (name, balance) VALUES ($1, $2)",
		user.Name,
		user.Balance,
	)
	return
}

func GetBalance(db *sql.DB, userId int) (balance decimal.Decimal, err error) {
	var balanceStr string
	err = db.QueryRow(
		"SELECT balance FROM users WHERE id = $1",
		userId,
	).Scan(&balanceStr)
	switch {
	case err == sql.ErrNoRows:
		err = fmt.Errorf("No user with id '%d'", userId)
		return
	case err != nil:
		return
	}
	return decimal.NewFromString(balanceStr)
}
