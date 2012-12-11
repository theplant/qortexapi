package qortexapi

import (
	"github.com/sunfmin/govalidations"
	"time"
)

type AuthUserService interface {
	CreatePost(input *InputEntry) (r *Entry, validated *govalidations.Validated, err error)
	// UpdatePost() (err error)
	// CreateWiki() (err error)
	// UpdateWiki() (err error)
	// CreatePostWithTask() (err error)
	// UpdatePostWithTask() (err error)
	// MyFeedEntries(entryType string, before time.Time, limit int) (r []*Entry, err error)
	GroupEntries(groupId string, entryType string, before time.Time, limit int) (g *Group, r []*Entry, err error)
	// UserEntries(userId string, entryType string, before time.Time, limit int) (u *User, r []*Entry, err error)
	// LoadEntry(groupId string, entryId string) (g *Group, r *Entry, err error)
	GetWatchList(before time.Time, limit int) (r *WatchList, err error)
	AddToWatchList(entryId string, groupId string) (added bool, err error)
	StopWatching(entryId string, groupId string) (stopped bool, err error)
	ReadWatching(entryId string, groupId string) (err error)

	GetDraftList(before time.Time, limit int) (r *DraftList, err error)
	DeleteDraft(entryId string, groupId string) (err error)

	GetAllGroups() (r []*Group, err error)
	AddUserToGroup(groupId string, userId string) (err error)
	RemoveUserFromGroup(groupId string, userId string) (err error)
}
