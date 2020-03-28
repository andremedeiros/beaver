package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectIPReturnsValidAddress(t *testing.T) {
	p := &Project{ID: 123, Name: "beaver"}
	assert.Equal(t, p.IP().String(), "192.168.220.123")
}

func TestCreateProject(t *testing.T) {
	s, teardown := setup(t)
	defer teardown()

	p := &Project{Name: "foobar", Path: "/home/beaver/src/foobar"}
	err := s.CreateProject(p)
	assert.NoError(t, err)

	p, err = s.ProjectByPath(p.Path)
	assert.NoError(t, err)
	assert.Equal(t, "foobar", p.Name)
	assert.Equal(t, 1, p.ID)
	assert.Equal(t, "/home/beaver/src/foobar", p.Path)
}
