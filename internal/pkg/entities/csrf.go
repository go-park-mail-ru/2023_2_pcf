package entities

type Csrf struct {
	Token  string `json:"token"`
	UserId int    `json:"-"`
}

type CsrfUseCaseInterface interface {
	CsrfCreate(userId int) (*Csrf, error)
	//CsrfRead(sr *Csrf) (*Csrf, error)
	CsrfRemove(sr *Csrf) error
	//CsrfContains(sr *Csrf) (bool, error)
	//Auth(*User) (*Session, error)
	GetByUserId(id int) (*Csrf, error)
}

type CsrfRepoInterface interface {
	Create(s *Csrf) (*Csrf, error)
	Remove(s *Csrf) error
	Read(userId int) (*Csrf, error)
	//Contains(s *Csrf) (bool, error)
}
