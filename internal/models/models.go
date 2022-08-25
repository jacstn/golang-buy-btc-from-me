package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Order struct {
	Id                      uint16
	BTCAmount               uint64
	USDAmount               uint64
	Address                 string
	BlockchainTransactionID string
	Status                  string
	CreatedAt               string
	UpdatedAt               string
}

func NewOrder(db *sql.DB, o *Order) error {
	_, err := db.Exec("INSERT into order (btcAmount, usdAmount, address, blockchainTransactionId, status, createdAt, updatedAt) values (?, ?, ?, ?, ?, 'NEW', ?, ?)",
		o.BTCAmount, o.USDAmount, o.Address, o.BlockchainTransactionID, time.Now(), time.Now())

	if err != nil {
		fmt.Println("error while inserting into database", err)
		return err
	}
	return nil
}

func ListOrders(db *sql.DB) ([]Order, error) {
	var orders []Order
	res, err := db.Query("SELEECT * FROM `order` ORDER BY createdAt DESC LIMIT 20")
	if err != nil {
		fmt.Println("error while selecting orders from database")
		return orders, err
	}

	for res.Next() {

		var o Order
		err := res.Scan(&o.Id, &o.BTCAmount, &o.USDAmount, &o.Address,
			&o.BlockchainTransactionID, &o.Status, &o.CreatedAt, &o.UpdatedAt)

		if err != nil {
			fmt.Println(err)

		}
		orders = append(orders, o)
	}
	return []Order{}, nil
}
