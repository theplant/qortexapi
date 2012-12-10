package qortexapi

type CanNotBlankError struct {
	Field string
}

func (this *CanNotBlankError) Error() string {
	return this.Field + " can't be blank."
}
