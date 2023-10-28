package entities

type Session struct {
	Token  string `json:"token"`
	UserId int    `json:"-"`
}

type SessionUseCaseInterface interface {
	SessionCreate(sr *Session) (*Session, error)
	SessionRead(sr *Session) (*Session, error)
	SessionRemove(sr *Session) error
	SessionContains(sr *Session) (bool, error)
}

type SessionRepoInterface interface {
	Create(s *Session) (*Session, error)
	Remove(s *Session) error
	Read(s *Session) (*Session, error)
	Contains(s *Session) (bool, error)
}
