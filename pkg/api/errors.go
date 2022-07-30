package api

type ErrNonExistentUser struct {
}

func (e *ErrNonExistentUser) Error() string {
	return "Non existent user"
}

func (e *ErrNonExistentUser) Is(target error) bool {
	_, isType := target.(*ErrNonExistentUser)

	return isType
}
