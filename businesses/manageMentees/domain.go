package manage_mentees

type domain struct{}

type Repository interface{}

type Usecase interface {
	// usecase delete access course mentee
	DeleteAccess(menteeId string, courseId string) error
}
