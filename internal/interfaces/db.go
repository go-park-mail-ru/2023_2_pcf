package interfaces

type Db interface {
	New() *Db
	Open() (*Db, error)
	Close()
}
