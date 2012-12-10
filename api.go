package qortexapi

import (
	"time"
)

type UserService interface {
	// CreatePost() (err error)
	// UpdatePost() (err error)
	// CreateWiki() (err error)
	// UpdateWiki() (err error)
	// CreatePostWithTask() (err error)
	// UpdatePostWithTask() (err error)
	// MyFeedEntries(entryType string, before time.Time, limit int) (r []*Entry, err error)
	// GroupEntries(groupId string, entryType string, before time.Time, limit int) (g *Group, r []*Entry, err error)
	// UserEntries(userId string, entryType string, before time.Time, limit int) (u *User, r []*Entry, err error)
	// LoadEntry(groupId string, entryId string) (g *Group, r *Entry, err error)
	GetWatchList(before time.Time, limit int) (r *WatchList, err error)
	AddToWatchList(entryId string, groupId string) (IsAdded bool, err error)
	StopWatching(entryId string, groupId string) (IsAdded bool, err error)
	ReadWatching(entryId string, groupId string) (err error)
}
