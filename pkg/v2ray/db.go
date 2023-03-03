package v2ray

import (
	"github.com/boltdb/bolt"
	"log"
	"v2ray.com/core/common/errors"
)

var (
	db *bolt.DB
)

func InitDB(path string) {
	_db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	db = _db
}

func createUserDB(username string, uuid string) error {
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	bkt, err := tx.CreateBucketIfNotExists([]byte("users"))
	if err != nil {
		return err
	}
	err = bkt.Put([]byte(username), []byte(uuid))
	if err != nil {
		return err
	}
	return tx.Commit()
}

func QueryUUID(username string) (string, error) {
	var (
		uid string
	)
	err := db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte("users"))
		if bkt == nil {
			return nil
		}
		name := bkt.Get([]byte(username))
		uid = string(name)
		return nil
	})
	if err != nil {
		return "", err
	}
	if uid == "" {
		return "", errors.New("not found")
	}
	return uid, nil
}
