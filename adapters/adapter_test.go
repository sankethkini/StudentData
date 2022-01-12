package adapters

import (
	"reflect"
	"testing"
)

func TestSingleton(t *testing.T) {
	//testing for file adapter
	file, err := Init(FILE)
	if err != nil {
		t.Error(err)
	}
	if (reflect.TypeOf(file) != reflect.TypeOf(&FileAdapter{})) {
		t.Error("not  a right type", file)
	}

	//testing for memory adapter
	memory, err := Init(MEMORY)
	if err != nil {
		t.Error(err)
	}
	if (reflect.TypeOf(memory) != reflect.TypeOf(&MemoryAdapter{})) {
		t.Error("not  a right type", file)
	}
}
