package api

type ErrNonExistentResource struct {
}

func (e *ErrNonExistentResource) Error() string {
	return "Non existent resource"
}

func (e *ErrNonExistentResource) Is(target error) bool {
	_, isType := target.(*ErrNonExistentResource)

	return isType
}
