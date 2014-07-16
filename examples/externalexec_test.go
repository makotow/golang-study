package examples

import (
	"testing"
	"os"
	"path/filepath"
)

// TODO testingパッケージの使い方を学び書き換える。
func TestCall(t *testing.T) {
	err := ExecuteCommand(filepath.Join(os.Getenv("HOME"), "tmp", "db", "HackGeo"), "load.sh")
	if err != nil {
		t.Log("OK")
	}
	err = ExecuteCommand(filepath.Join(os.Getenv("HOME"), "tmp", "db", "HackGeo"), "./test.sh")
	if err != nil {
		t.Log(err)
		t.Log("Expected: Permission denied. not executalbe.")
	}
	err = ExecuteCommand(filepath.Join(os.Getenv("HOME"), "tmp", "db", "HackGeo"), "./test2.sh")
	if err != nil {
		t.Log(err)
		t.Error()
	}
}
