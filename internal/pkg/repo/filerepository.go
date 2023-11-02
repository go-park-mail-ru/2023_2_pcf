package repo

import (
	"fmt"
	"io/ioutil"
	"os"
)

type FileRepository struct {
	storagePath string
}

func NewFileRepository(storagePath string) *FileRepository {
	return &FileRepository{storagePath}
}

func (repo *FileRepository) Save(fileData []byte, fileName string) error {
	filePath := repo.getFilePath(fileName)

	err := ioutil.WriteFile(filePath, fileData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (repo *FileRepository) Get(fileName string) ([]byte, error) {
	filePath := repo.getFilePath(fileName)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *FileRepository) Delete(fileName string) error {
	filePath := repo.getFilePath(fileName)

	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}

func (repo *FileRepository) getFilePath(fileName string) string {
	return fmt.Sprintf("%s/%s", repo.storagePath, fileName)
}
