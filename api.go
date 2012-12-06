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
	MyFeedEntries(entryType string, before time.Time, limit int) (r []*Entry, err error)
	GroupEntries(groupId string, entryType string, before time.Time, limit int) (g *Group, r []*Entry, err error)
	UserEntries(userId string, entryType string, before time.Time, limit int) (u *User, r []*Entry, err error)
	LoadEntry(groupId string, entryId string) (g *Group, r *Entry, err error)
}
