package qortexapi

import (
	"errors"
)

type CanNotBlankError struct {
	Field string
}

func NewCanNotBlankError(field string) *CanNotBlankError {
	return &CanNotBlankError{
		Field: field,
	}
}

func (this *CanNotBlankError) Error() string {
	return this.Field + " can't be blank"
}

var ServerError = errors.New("Oops, something is wrong!")

var OrganizationNotFoundError = errors.New("organization not found")
var OrganizationsNotFoundError = errors.New("organizations not found")
var GroupNotFoundError = errors.New("group not found")
var GroupsNotFoundError = errors.New("groups not found")
var UserNotFoundError = errors.New("user not found")
var UsersNotFoundError = errors.New("users not found")
var PermissionDeniedError = errors.New("permiession denied")
var UserSavedError = errors.New("can not save user")
var SaveGroupError = errors.New("can not save group")
var DeleteGroupError = errors.New("can not delete group")
var BackupEntryError = errors.New("can not backup entry")
var EntryNotFoundError = errors.New("entry not found")
var RemoveIndexError = errors.New("can not remove index")
var EntrySaveError = errors.New("can not save entry")
