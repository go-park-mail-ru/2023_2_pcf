package entities

type ULink struct {
	Token string `json:"token"`
	AdId  int    `json:"-"`
}

//go:generate /Users/bincom/go/bin/mockgen -source=uniquelink.go -destination=mock_entities/uniquelink_mock.go
type ULinkUseCaseInterface interface {
	ULinkCreate(sr *ULink) (*ULink, error)
	ULinkRead(sr *ULink) (*ULink, error)
	ULinkRemove(sr *ULink) error
	ULinkContains(sr *ULink) (bool, error)
	GetAdId(token string) (int, error)
}

type ULinkRepoInterface interface {
	Create(s *ULink) (*ULink, error)
	Remove(s *ULink) error
	Read(s *ULink) (*ULink, error)
	Contains(s *ULink) (bool, error)
}
