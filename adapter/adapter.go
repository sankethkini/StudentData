package adapter

type Adapter struct {
	FileAdapter   FileDriver
	MemoryAdapter MemoryDriver
}

type MemoryDriver interface {
	Save(items interface{}) error
	Retrive(feild string, value string) bool
	RetriveAll(field string, order int) (interface{}, error)
	Delete(field string, value string) error
}

type FileDriver interface {
	Save(items interface{}) error
	RetriveAll() (interface{}, error)
}

func NewAdapter(f FileDriver, m MemoryDriver) (*Adapter, error) {
	cur := Adapter{}
	cur.FileAdapter = f
	cur.MemoryAdapter = m
	return &cur, nil
}
