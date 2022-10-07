package errors 

type DuplicateUserError struct {

}

func (e DuplicateUserError) Error() string {
	return "duplicate user"
}