package file

import (
	"AdHub/internal/pkg/entities"
)

type FileUseCase struct {
	repo entities.FileRepoInterface
}

func New(r entities.FileRepoInterface) *FileUseCase {
	return &FileUseCase{
		repo: r,
	}
}

func (uc *FileUseCase) Save(fileData []byte, originalName string) (string, error) {
	return uc.repo.Save(fileData, originalName)
}

func (uc *FileUseCase) Get(fileName string) ([]byte, error) {
	return uc.repo.Get(fileName)
}

func (uc *FileUseCase) Delete(fileName string) error {
	return uc.repo.Delete(fileName)
}
