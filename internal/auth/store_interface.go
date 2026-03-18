package auth

type Store interface {
	Save(session *Session) error
	Get(id string) (*Session, error)
	Delete(id string) error
}
