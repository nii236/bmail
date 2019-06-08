package service

// S are long running services for bmail
type S interface {
	Start()
	Stop()
	Name() string
	Description() string
}
