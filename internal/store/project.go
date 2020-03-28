package store

import (
	"encoding/json"
	"fmt"
	"net"

	"go.etcd.io/bbolt"
)

type Project struct {
	ID   int
	Name string
	Path string
}

type ProjectStore interface {
	CreateProject(p *Project) error
	ProjectByPath(path string) (*Project, error)
}

func (p *Project) IP() net.IP {
	ip := fmt.Sprintf("192.168.220.%d", p.ID) // TODO(andremedeiros): Get a better IP schema
	return net.ParseIP(ip)
}

func (s *Store) CreateProject(p *Project) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(projectsBucket)
		id, _ := b.NextSequence()
		p.ID = int(id)
		buf, err := json.Marshal(p)
		if err != nil {
			return err
		}

		return b.Put([]byte(p.Path), buf)
	})
}

func (s *Store) ProjectByPath(path string) (*Project, error) {
	var p Project

	err := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(projectsBucket)
		c := b.Cursor()

		_, v := c.Seek([]byte(path))
		return json.Unmarshal(v, &p)
	})

	if err != nil {
		return nil, err
	}

	return &p, nil
}
