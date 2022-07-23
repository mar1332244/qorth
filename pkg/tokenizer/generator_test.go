package tokenizer

import (
	"io"
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
		if entry.IsDir() {
			continue
		}
		fname := dirname + "/" + entry.Name()
		tokens, err := GetTokensFromFile(fname)
		if err != nil {
			t.Error(err)
		}
		t.Log(tokens)
		fp, err := os.Open(fname)
		if err != nil {
			continue
		}
		defer fp.Close()
		buf, err := io.ReadAll(fp)
		if err != nil {
			continue
		}
		t.Logf("\n%s", buf)
	}
}
