package main

import (
	"errors"
	"fmt"
	"log"
)

const dbReady = true
const balance = 200
const numberToSuccess = 2

func connectDb(nTy int) error {
	if nTy == numberToSuccess {
		return nil
	}
	return errors.New("busy")
}

func waitForDatabase() error {
	connectDb(1)
	connectDb(2)
	return nil
}

func getBalance() (int, error) {
	if !dbReady {
		if err := waitForDatabase(); err != nil {
			return 0, fmt.Errorf("getBalance: %v", err)
		}

	}
	return balance, nil
}

func withdraw(amount int) (int, error) {

	balance, err := getBalance()
	if err != nil {
		return 0, fmt.Errorf("withdraw: %v", err)
	}

	if !dbReady {
		return 0, errors.New("withdraw: database is down")
	}
	if amount > balance {
		return 0, errors.New("withdraw: insufficient fund")
	}
	return amount, nil
}

func main() {
	log.SetFlags(0)

	amount, err := withdraw(200)
	if err != nil {

		// fmt.Println("main: ", err)
		// os.Exit(1)
		log.Fatalf("main: %v", err)
	}
	fmt.Println("Please collect your money :", amount)
}
