package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	storage "users/pkg/db"

	"github.com/boltdb/bolt"
)

type Bucket string

const (
	// RoleBucket db
	RoleBucket Bucket = "ROLE"
	// UserBucket db
	UserBucket Bucket = "USER"
)

// Bolt is the database driver.
type Bolt struct {
	// client is the Bolt client.
	client *bolt.DB
	bucket Bucket
}

// NewBoltDB creates a Bolt DB database driver given an underlying client.
func NewBoltDB(client *bolt.DB, bucket Bucket) (*Bolt, error) {
	err := client.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return fmt.Errorf("could not %v user bucket: %v", bucket, err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &Bolt{client, bucket}, nil
}

// NewID returns a unique ID for the given bucket.
func (b *Bolt) NewID() (string, error) {
	var sid string
	err := b.client.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(b.bucket))
		id, err := bkt.NextSequence()
		if err != nil {
			return err
		}
		sid = strconv.FormatUint(id, 10)
		return nil
	})
	return sid, err
}

// Save writes the record to the DB and returns the corresponding new ID.
// data must contain a value that can be marshaled by the encoding/json package.
func (b *Bolt) Save(id string, data interface{}) error {
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return b.client.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(b.bucket))
		if err := bkt.Put([]byte(id), buf); err != nil {
			return err
		}
		return nil
	})
}

// Delete deletes a record by ID.
func (b *Bolt) Delete(id string) error {
	return b.client.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(b.bucket)).Delete([]byte(id))
	})
}

// Load reads a record by ID. data is unmarshaled into and should hold a pointer.
func (b *Bolt) Load(id string, data interface{}) error {
	err := b.client.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(b.bucket))
		v := bkt.Get([]byte(id))
		if v == nil {
			return storage.ErrNotFound
		}

		err := json.Unmarshal(v, data)
		return err
	})

	return err
}

// LoadAll returns all the records in the given bucket. data should be a pointer
// to a slice. Don't do this in a real service :-)
func (b *Bolt) LoadAll(data interface{}) error {

	buf := &bytes.Buffer{}
	err := b.client.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(b.bucket))
		buf.WriteByte('[')
		if bkt != nil {
			first := true
			if err := bkt.ForEach(func(_, v []byte) error {
				if len(v) > 0 {
					if first {
						first = false
					} else {
						buf.WriteByte(',')
					}

					buf.Write(v)
				}
				return nil
			}); err != nil {
				return err
			}

		}
		buf.WriteByte(']')
		return nil
	})
	if err != nil {
		return err
	}
	fmt.Printf("Buffer: %s\n", buf.Bytes())
	if err = json.Unmarshal(buf.Bytes(), data); err != nil {
		return err
	}

	return nil
}
