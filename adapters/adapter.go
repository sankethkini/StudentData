package adapters

//constans
const (
	FILE   = "file"
	MEMORY = "memory"
)

//interface for adapter
type IAdapter interface {
	Save(items interface{}) error
	Retrive(feild string, value string) bool
	RetriveAll(field string, order int) (interface{}, error)
	Delete(field string, value string) error
}

var fileAdapter IAdapter
var memoryAdapter IAdapter

//singleton implementation
//init function
func Init(typ string) (IAdapter, error) {
	switch typ {

	case FILE:
		if fileAdapter == nil {
			fileAdapter = &FileAdapter{}
		}
		return fileAdapter, nil

	case MEMORY:
		if memoryAdapter == nil {
			memoryAdapter = &MemoryAdapter{}
		}
		return memoryAdapter, nil

	default:
		return nil, NoAdapterFound

	}
}
