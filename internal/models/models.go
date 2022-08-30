package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Order struct {
	Id                      uint16
	SatoshiAmount           uint64
	USDAmount               uint64
	Address                 string
	BlockchainTransactionID string
	Status                  string
	CreatedAt               string
	UpdatedAt               string
}

func NewOrder(db *sql.DB, o *Order) (int64, error) {
	res, err := db.Exec("INSERT INTO `order` (satoshiAmount, usdAmount, address, blockchainTransactionId, status, createdAt, updatedAt) values (?, ?, ?, ?, 'NEW', ?, ?)",
		o.SatoshiAmount, o.USDAmount, o.Address, o.BlockchainTransactionID, time.Now(), time.Now())

	if err != nil {
		fmt.Println("error while inserting into database", err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		println("Error:", err.Error())
		return 0, err
	}

	return id, nil
}

func ListOrders(db *sql.DB) ([]Order, error) {
	var orders []Order
	res, err := db.Query("SELECT * FROM `order` ORDER BY createdAt DESC LIMIT 20")
	if err != nil {
		fmt.Println("error while selecting orders from database")
		return orders, err
	}

	for res.Next() {

		var o Order
		err := res.Scan(&o.Id, &o.SatoshiAmount, &o.USDAmount, &o.Address,
			&o.BlockchainTransactionID, &o.Status, &o.CreatedAt, &o.UpdatedAt)

		if err != nil {
			fmt.Println(err)
		}
		orders = append(orders, o)
	}
	return []Order{}, nil
}

func (o *Order) Save(db *sql.DB) error {
	if o.Id == 0 {
		return errors.New("unable to update, order does not exists")
	}
	fmt.Println("new status", o.Status, o.Id)
	res, err := db.Exec("UPDATE `order` set blockchainTransactionId=?, status=?, updatedAt=? where id=?", o.BlockchainTransactionID, o.Status, time.Now(), o.Id)

	fmt.Println(res)
	return err
}

func GetOrderById(db *sql.DB, id int64) Order {
	o := Order{}
	db.QueryRow(fmt.Sprintf("SELECT * FROM `order` WHERE `id`=%d", id)).Scan(&o.Id, &o.SatoshiAmount, &o.USDAmount, &o.Address,
		&o.BlockchainTransactionID, &o.Status, &o.CreatedAt, &o.UpdatedAt)

	return o
}
