package storage

import (
	"errors"
	"fmt"
	"go.etcd.io/bbolt"
	"log"
	"os"
)

const DBNAME = "urlshort.db"

type Connection struct {
	db *bbolt.DB
}

var Conn Connection

type BucketCreationError struct {
	bucketName string
}

type RouteData struct {
	Path string `yaml:"path" json:"path"`
	Url  string `yaml:"url" json:"url"`
}

func (e *BucketCreationError) Error() string {
	return fmt.Sprintf("failed to create bucket %s", e.bucketName)
}

func InitDB() {
	if Conn.db != nil {
		return
	}

	var err error
	Conn.db, err = bbolt.Open(DBNAME, 0600, nil)

	if !errors.Is(err, nil) {
		log.Panic(err)
	}

	err = Conn.db.Update(func(tx *bbolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte("DB"))

		if !errors.Is(err, nil) {
			return &BucketCreationError{bucketName: "DB"}
		}

		_, err = root.CreateBucketIfNotExists([]byte("ROUTES"))

		if !errors.Is(err, nil) {
			return &BucketCreationError{bucketName: "ROUTES"}
		}

		return nil
	})

	if !errors.Is(err, nil) {
		log.Fatal("failed to create DB")
	}

	log.Println("DB created")
}

func (Conn *Connection) AddRoute(route RouteData) error {
	err := Conn.db.Update(func(tx *bbolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte("ROUTES")).Put([]byte(route.Path), []byte(route.Url))

		if !errors.Is(err, nil) {
			return fmt.Errorf("failed to insert new route: %v", route)
		}

		return nil
	})

	return err
}

func (Conn *Connection) GetRoute(path string) (string, error) {
	var url []byte
	err := Conn.db.View(func(tx *bbolt.Tx) error {
		url = tx.Bucket([]byte("DB")).Bucket([]byte("ROUTES")).Get([]byte(path))
		return nil
	})

	return string(url), err
}

func (Conn *Connection) CloseConnection() {
	Conn.db.Close()
	os.Remove(Conn.db.Path())
}
