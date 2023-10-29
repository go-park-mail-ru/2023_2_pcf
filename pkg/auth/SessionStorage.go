package auth

import (
	"fmt"
	"sync"
)

//go:generate /Users/bincom/go/bin/mockgen -source=SessionStorage.go -destination=mocks/mock.go
type SessionStorageInterface interface {
	AddSession(session Session)
	RemoveSession(token string)
	RemoveUser(userId int)
	Contains(token string) bool
	GetUserId(token string) (userId int, err error)
}

type SessionStorage struct {
	Sessions map[string]int //key - session token, value - user_id
	mutex    sync.Mutex
}

var MySessionStorage = &SessionStorage{
	Sessions: make(map[string]int),
}

func (ss *SessionStorage) AddSession(session Session) {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	ss.Sessions[session.Token] = session.UserId
}

// RemoveSession removes 1 session (by its token)
func (ss *SessionStorage) RemoveSession(token string) {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	delete(ss.Sessions, token)
}

// RemoveUser removes all Sessions assigned to this user (by userId)
func (ss *SessionStorage) RemoveUser(userId int) {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()

	toDeleteBuf := []string{} //ищем все токены сессий пользователя
	for token, storedUserID := range ss.Sessions {
		if storedUserID == userId {
			toDeleteBuf = append(toDeleteBuf, token)
		}
	}

	//не удалять этот комментарий!
	//Удаление с помощью буфера позволяет гарантировать,
	//что будут удалены все сессии пользователя, так как в
	//таком случае порядок обхода мапы не будет изменен
	//во время самого обхода (операция удаления элемента
	//может привести к изменению структуры мапы)
	for _, key := range toDeleteBuf {
		delete(ss.Sessions, key)
	}
}

func (ss *SessionStorage) Contains(token string) bool {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	_, got := ss.Sessions[token]
	return got
}

func (ss *SessionStorage) GetUserId(token string) (userId int, err error) {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	userId, got := ss.Sessions[token]
	if !got {
		err = fmt.Errorf("Session not found")
		return -1, err
	}
	return userId, nil
}
