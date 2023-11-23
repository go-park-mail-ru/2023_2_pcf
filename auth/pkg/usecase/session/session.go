package session

import (
	"AdHub/auth/pkg/entities"
	"AdHub/pkg/cryptoUtils"
)

type SessionUseCase struct {
	repo entities.SessionRepoInterface
}

func New(r entities.SessionRepoInterface) *SessionUseCase {
	return &SessionUseCase{
		repo: r,
	}
}

func (su SessionUseCase) SessionCreate(sr *entities.Session) (*entities.Session, error) {
	return su.repo.Create(sr)
}

func (su SessionUseCase) SessionRead(sr *entities.Session) (*entities.Session, error) {
	return su.repo.Read(sr)
}

func (su SessionUseCase) SessionRemove(sr *entities.Session) error {
	return su.repo.Remove(sr)
}

func (su SessionUseCase) SessionContains(sr *entities.Session) (bool, error) {
	return su.repo.Contains(sr)
}

func (su SessionUseCase) Auth(userFromDB *entities.User) (*entities.Session, error) {
	var tokenLen = 32
	newSession := &entities.Session{UserId: userFromDB.Id}

	var err error
	newSession.Token, err = cryptoUtils.GenToken(tokenLen)
	if err != nil {
		return nil, err
	}

	//Проверка уникальности токена, регенерация если он уже занят
	for contains, err := su.SessionContains(newSession); contains; su.SessionContains(newSession) {
		newSession.Token, err = cryptoUtils.GenToken(tokenLen)
		if err != nil {
			return nil, err
		}
	}
	newSession, err = su.SessionCreate(newSession)
	if err != nil {
		return nil, err
	}

	return newSession, nil
}

func (su SessionUseCase) GetUserId(token string) (int, error) {
	s, err := su.repo.Read(&entities.Session{
		Token:  token,
		UserId: 0,
	})
	return s.UserId, err
}
