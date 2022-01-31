package adapter

import "github.com/sankethkini/StudentData/domain/user"

type MemoryDriver interface {
	Save(item ...user.User) error
	Retrieve(feild string, value string) bool
	RetriveAll(field string, order int) ([]user.User, error)
	Delete(field string, value string) error
}

type FileDriver interface {
	Save([]user.User) error
	RetriveAll() ([]user.User, error)
}

type Adapter struct {
	FileAdapter   FileDriver
	MemoryAdapter MemoryDriver
}

func NewAdapter(f FileDriver, m MemoryDriver) (*Adapter, error) {
	cur := Adapter{}
	cur.FileAdapter = f
	cur.MemoryAdapter = m
	return &cur, nil
}
