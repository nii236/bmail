package services

// S are long running services for bmail
type S interface {
	Run() error
	Name() string
	Description() string
}
