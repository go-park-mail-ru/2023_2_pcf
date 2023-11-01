package entities

type FileRepositoryInterface interface {
	Save(fileData []byte, fileName string) error
	Get(fileName string) ([]byte, error)
	Delete(fileName string) error
}

type FileUseCaseInterface interface {
	Save(fileData []byte, fileName string) error
	Get(fileName string) ([]byte, error)
	Delete(fileName string) error
}
