package repo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type FileRepository struct {
	storagePath string
}

func NewFileRepository(storagePath string) *FileRepository {
	return &FileRepository{storagePath}
}

func (repo *FileRepository) GenerateUniqueFileName(originalName string) string {
	uniqueID := uuid.New()

	fileExtension := filepath.Ext(originalName)

	uniqueFileName := fmt.Sprintf("%s%s", uniqueID, fileExtension)

	return uniqueFileName
}

func (repo *FileRepository) Save(fileData []byte, originalName string) (string, error) {
	uniqueFileName := repo.GenerateUniqueFileName(originalName)
	filePath := repo.getFilePath(uniqueFileName)
	//wd, _ := os.Getwd()
	err := os.WriteFile(filePath, fileData, 0644)
	if err != nil {
		return "", fmt.Errorf("%s", err.Error())
	}

	return uniqueFileName, nil
}

func (repo *FileRepository) Get(fileName string) ([]byte, error) {
	filePath := repo.getFilePath(fileName)

	data, err := os.ReadFile(filePath)
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
	return filepath.Join(repo.storagePath, fileName)
}
