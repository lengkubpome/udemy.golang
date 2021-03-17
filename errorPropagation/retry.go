package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

const dbReady = false
const balance = 200
const numberToSuccess = 10

func connectDb(nTy int) error {
	if nTy == numberToSuccess {
		return nil
	}
	return errors.New("busy")
}

func waitForDatabase() error {
	timeout := 5 * time.Second
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		err := connectDb(tries)
		if err == nil {
			return nil
		}
		log.Printf("database is not responding (%v), retrying...", err)
		time.Sleep(time.Second << tries) // 1 2 4 8 16s
	}
	return fmt.Errorf("waitForDatabase > timeout %v", timeout)
}

func getBalance() (int, error) {
	if !dbReady {
		if err := waitForDatabase(); err != nil {
			return 0, fmt.Errorf("getBalance > %v", err)
		}

	}
	return balance, nil
}

func withdraw(amount int) (int, error) {

	balance, err := getBalance()
	if err != nil {
		return 0, fmt.Errorf("withdraw > %v", err)
	}

	if !dbReady {
		return 0, errors.New("withdraw > database is down")
	}
	if amount > balance {
		return 0, errors.New("withdraw > insufficient fund")
	}
	return amount, nil
}

func main() {
	log.SetFlags(0)

	amount, err := withdraw(200)
	if err != nil {
		// fmt.Println("main: ", err)
		// os.Exit(1)
		log.Fatalf("main > %v", err)
	}
	fmt.Println("Please collect your money :", amount)
}
