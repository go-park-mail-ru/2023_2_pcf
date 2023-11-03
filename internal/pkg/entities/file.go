package entities

//go:generate /Users/bincom/go/bin/mockgen -source=file.go -destination=mock_entities/file_mock.go
type FileRepoInterface interface {
	Save(fileData []byte, originalName string) (string, error)
	Get(fileName string) ([]byte, error)
	Delete(fileName string) error
}

type FileUseCaseInterface interface {
	Save(fileData []byte, originalName string) (string, error)
	Get(fileName string) ([]byte, error)
	Delete(fileName string) error
}
