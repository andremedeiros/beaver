package store

import (
	"os"
	"path"

	"go.etcd.io/bbolt"
)

var (
	projectsBucket = []byte("projects")

	buckets = [][]byte{
		projectsBucket,
	}
)

type Store struct {
	db *bbolt.DB

	ProjectStore
}

func NewStore(root string) (*Store, error) {
	if err := os.MkdirAll(path.Join(root, ".beaver"), 0777); err != nil {
		return nil, err // TODO(andremedeiros): Return a more useful error here
	}

	db, err := bbolt.Open(path.Join(root, ".beaver", "beaver.db"), 0660, nil)
	if err != nil {
		return nil, err // TODO(andremedeiros): Return a more useful error here
	}

	// Set up buckets
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, err := tx.CreateBucketIfNotExists(b); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err // TODO(andremedeiros): Return a better error here
	}

	return &Store{db: db}, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}
