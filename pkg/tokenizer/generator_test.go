package tokenizer

import (
	"os"
	"testing"
)

func TestGetTokensFromFile(t *testing.T) {
	dirname := "./../../examples/"
	dir, err := os.ReadDir(dirname)
	if err != nil {
		t.Error(err)
	}
	for _, entry := range dir {
		t.Logf("%s\n", entry.Name())
	}
}
