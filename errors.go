package qortexapi
import(
	"errors"
)

type CanNotBlankError struct {
	Field string
}

func (this *CanNotBlankError) Error() string {
	return this.Field + " can't be blank"
}

var GroupNotFoundError = errors.New("group not found")
var UserNotFoundError = errors.New("user not found")
var PermissionDenied = errors.New("permiession denied")
