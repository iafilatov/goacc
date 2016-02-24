Tiny accout/balance manager
===========================

An exercise on HTTP/JSON/databases in Go. I'm new to Go so don't be harsh on this. Not meant for any other kind of use.

At all.


Installation
------------

1. PostgreSQL

2. `go get github.com/iafilatov/goacc`


Running
-------

    goacc "dsn string" createdb

to create the DB table.

Then

    goacc 

or

    goacc "dsn string"

> Default dsn is `password=goacc user=goacc dbname=goacc sslmode=disable`

It will start on `0.0.0.0:8080`


API
---

**GET** `/users`

Return a list of users like .

    [
      {
        "id":2,
        "name":"John",
        "balance":"1164"
      },
      {
        "id":1,
        "name":"Jane",
        "balance":"23.14"
      }
    ]


**POST** `/create`

    {
      "name":"John",
      "balance":"1164"
    }

Returns the newly created user.

    {
      "id":2,
      "name":"John",
      "balance":"1164"
    }
    

**GET** /balance?user=2

The balance of the user.

    {"balance": 1000}


**POST** /deposit

    {"user": 2, "amount": 100}
    
Deposits 100 to the user with id 2.


**POST** /withdraw

    {"user": 2, "amount": 50}
    
Withdraws 50 form the account of user with id 2.


**POST** /transfer

    {"from": 2, "to": 1, amount: 25}

Transfers 25 from user 2 to user 1.
