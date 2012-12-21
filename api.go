package qortexapi

import (
	"github.com/sunfmin/govalidations"
	"time"
)

type AuthUserService interface {
	CreatePost(input *EntryInput) (r *Entry, validated *govalidations.Validated, err error)
	EditPost(entryId string, groupId string) (r *Entry, err error)
	UpdatePost(input *EntryInput) (r *Entry, validated *govalidations.Validated, err error)
	CreateComment(input *EntryInput) (r *Entry, validated *govalidations.Validated, err error)
	EditComment(entryId string, groupId string) (r *Entry, err error)
	CreateWiki(input *EntryInput) (r *Entry, validated *govalidations.Validated, err error)
	EditWiki(entryId string, groupId string) (r *Entry, err error)
	GetWiki(entryId string, groupId string, updateAtUnixNano string) (r *Entry, err error)
	GetWikiByTitle(title string, groupId string, updateAtUnixNano string) (r *Entry, err error)
	UpdateWiki(input *EntryInput) (r *Entry, validated *govalidations.Validated, err error)
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

	//Group related
	CreateGroup(input *GroupInput) (r *Group, validated *govalidations.Validated, err error)
	UpdateGroup(input *GroupInput) (validated *govalidations.Validated, err error)
	DeleteGroup(groupId string) (err error)
	ConvertToSharedGroup(input *GroupInput) (validated *govalidations.Validated, err error)
	GetAllGroups(keyword string) (r []*Group, err error)
	GetPublicGroups(keyword string) (r []*Group, err error)
	AddUserToGroup(groupId string, userId string) (err error)
	RemoveUserFromGroup(groupId string, userId string) (err error)
	OrganizationUsers(query string, pageNumber int, countPerPage int) (r []*User, pageCount int, err error)
	GroupUsers(groupId string, query string, OnlyFollowers bool, sortKey string, countPerPage int) (r []*User, newSortKey string, err error)

	//User related
	EnableUser(userId string) (err error)
	DisableUser(userId string) (err error)
	DeleteUser(userId string) (err error)
	PromoteToSuperUser(userId string) (err error)
	DemoteFromSuperUser(userId string) (err error)
	FollowUser(userId string) (err error)
	UnfollowUser(userId string) (err error)
}
