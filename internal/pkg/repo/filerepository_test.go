package repo

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileRepository_SaveGetDelete(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	repo := NewFileRepository(tempDir)

	data := []byte("test data")
	fileName := "test.txt"

	uniqueFileName, err := repo.Save(data, fileName)
	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	retrievedData, err := repo.Get(uniqueFileName)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}

	if string(retrievedData) != string(data) {
		t.Errorf("Retrieved data does not match saved data")
	}

	err = repo.Delete(uniqueFileName)
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
}
