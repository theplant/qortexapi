package qortexapi

import (
	"errors"
)

type CanNotBlankError struct {
	Field string
}

func (this *CanNotBlankError) Error() string {
	return this.Field + " can't be blank"
}

var OrganizationNotFoundError = errors.New("organization not found")
var GroupNotFoundError = errors.New("group not found")
var UserNotFoundError = errors.New("user not found")
var PermissionDeniedError = errors.New("permiession denied")
var UserSavedError = errors.New("can not save user")
var SaveGroupError = errors.New("can not save group")
var DeleteGroupError = errors.New("can not delete group")
var BackupEntryError = errors.New("can not backup entry")
