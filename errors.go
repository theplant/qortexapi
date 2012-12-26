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

var (
	ServerError                    = errors.New("Oops, something is wrong!")
	OrganizationNotFoundError      = errors.New("organization not found")
	OrganizationNotJoinedError     = errors.New("not join organization yet")
	OrganizationsNotFoundError     = errors.New("organizations not found")
	NonStandardOrganizationError   = errors.New("this organization is not standard")
	GroupNotFoundError             = errors.New("group not found")
	GroupsNotFoundError            = errors.New("groups not found")
	UserNotFoundError              = errors.New("user not found")
	UsersNotFoundError             = errors.New("users not found")
	PermissionDeniedError          = errors.New("permiession denied")
	SaveMemberError                = errors.New("can not save member")
	SaveUserError                  = errors.New("can not save user")
	SaveGroupError                 = errors.New("can not save group")
	SaveOrganizationError          = errors.New("can not save organization")
	DeleteGroupError               = errors.New("can not delete group")
	BackupEntryError               = errors.New("can not backup entry")
	EntryNotFoundError             = errors.New("entry not found")
	RemoveIndexError               = errors.New("can not remove index")
	SaveEntryError                 = errors.New("can not save entry")
	UserIsDeletedError             = errors.New("user is already deleted")
	SelfActionError                = errors.New("can not do this for your self")
	FollowUserError                = errors.New("can not follow user")
	UnfollowUserError              = errors.New("can not unfollow user")
	AlreadyJoinedOrganizationError = errors.New("already joined the organization")
)
