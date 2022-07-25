package tokenizer

import (
	"path/filepath"
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
		name := entry.Name()
		t.Logf("%s\n", name)
		tokens, err := GetTokensFromFile(filepath.Join(dirname, name))
		if err != nil {
			t.Error(err)
		}
		for _, token := range tokens {
			t.Log(token.Repr)
		}
	}
}
