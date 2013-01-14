package qortexapi

import (
	"github.com/sunfmin/govalidations"
	"time"
)

type Global interface {
	GetSession(email string, password string) (session string, err error)
	GetAuthUserService(session string) (authUserService AuthUserService, err error)
}

type AuthUserService interface {
	NewEntry(groupId string) (entry *Entry, err error)
	QortexMessages(messsageType string, before string, limit int) (entries []*Entry, err error) // when messageType is empty or equals "all", return all kinds of messages
	CreateBroadcast(input *BroadcastInput) (entry *Entry, validated *govalidations.Validated, err error)
	CreateBroadcastComment(input *BroadcastInput) (entry *Entry, validated *govalidations.Validated, err error)
	EditBroadcast(entryId string) (entry *Entry, err error)
	EditBroadcastComment(entryId string) (entry *Entry, err error)
	UpdateBroadcast(input *BroadcastInput) (entry *Entry, validated *govalidations.Validated, err error)
	UpdateBroadcastComment(input *BroadcastInput) (entry *Entry, validated *govalidations.Validated, err error)
	CreatePost(input *EntryInput) (entry *Entry, validated *govalidations.Validated, err error)
	GetPost(entryId string, groupId string) (entry *Entry, err error)
	UpdatePost(input *EntryInput) (entry *Entry, validated *govalidations.Validated, err error)
	CreateTask(input *EntryInput) (entry *Entry, validated *govalidations.Validated, err error)
	GetTask(entryId string, groupId string) (entry *Entry, err error)
	CloseTask(entryId string, groupId string) (entry *Task, err error)
	CreateComment(input *EntryInput) (entry *Entry, validated *govalidations.Validated, err error)
	GetComment(entryId string, groupId string) (entry *Entry, err error)
	UpdateComment(input *EntryInput) (entry *Entry, validated *govalidations.Validated, err error)
	CreateWiki(input *EntryInput) (entry *Entry, validated *govalidations.Validated, err error)
	GetWiki(entryId string, groupId string, updateAtUnixNano string, searchKeyWords string) (entry *Entry, err error)
	GetWikiByTitle(title string, groupId string, updateAtUnixNano string, searchKeyWords string) (entry *Entry, err error)
	UpdateWiki(input *EntryInput) (entry *Entry, validated *govalidations.Validated, err error)
	GetEntry(entryId string, groupId string, searchKeyWords string) (entry *Entry, err error)

	GroupUnreadEntryIds(entryIds []string, groupId string) (unreadEntryIds []string, err error)
	UnreadEntryIds(entryIds []string, groupIds []string) (unreadEntryIds []string, err error)
	GroupEntries(groupId string, entryType string, before string, limit int) (entries []*Entry, err error)
	MyFeedEntries(entryType string, before string, limit int) (entries []*Entry, err error)
	// NewMyFeedEntries(entryType string, after time.Time, limit int) (entries []*Entry, err error)
	MyTaskEntries(active bool, before string, limit int) (TasksForMe []*Entry, MyCreatedTasks []*Entry, err error)

	MyChatEntries(before string, limit int) (entries []*Entry, err error)
	// UserEntries(userId string, entryType string, before time.Time, limit int) (u *User, entries []*Entry, err error)
	// LoadEntry(groupId string, entryId string) (g *Group, entry *Entry, err error)

	MyNotificationItems(before string, limit int) (notificationItems []*NotificationItem, err error)

	// watchlist related
	GetWatchList(before time.Time, limit int) (watchlist *WatchList, err error)
	AddToWatchList(entryId string, groupId string) (added bool, err error)
	StopWatching(entryId string, groupId string) (stopped bool, err error)
	ReadWatching(entryId string, groupId string) (err error)

	// draft related
	GetDraftList(before time.Time, limit int) (draftlist *DraftList, err error)
	DeleteDraft(entryId string, groupId string) (err error)

	//Group related
	CreateGroup(input *GroupInput) (group *Group, validated *govalidations.Validated, err error)
	UpdateGroup(input *GroupInput) (validated *govalidations.Validated, err error)
	DeleteGroup(groupId string) (err error)
	ConvertToSharedGroup(input *GroupInput) (validated *govalidations.Validated, err error)
	GetAllGroups(keyword string) (groups []*Group, err error)
	GetPublicGroups(keyword string) (groups []*Group, err error)
	AddUserToGroup(groupId string, userId string) (err error)
	RemoveUserFromGroup(groupId string, userId string) (err error)
	GetGroupHeaderItem(groupId string) (ghi *GroupHeaderItem)

	//User related
	OrganizationUsers(query string, pageNumber int, countPerPage int) (users []*User, pageCount int, err error)
	GroupUsers(groupId string, query string, OnlyFollowers bool, sortKey string, countPerPage int) (users []*User, newSortKey string, err error)
	GetUser(userId string) (user *User, err error)
	EnableUser(userId string) (err error)
	DisableUser(userId string) (err error)
	DeleteUser(userId string) (err error)
	PromoteToSuperUser(userId string) (err error)
	DemoteFromSuperUser(userId string) (err error)
	FollowUser(userId string) (err error)
	UnfollowUser(userId string) (err error)

	//Organization Related

	OrganizationsInfo(orgIds []string) (orgs []*Organization, err error)
	OrganizationInfo(orgId string) (org *Organization, err error)
	SearchOrganizations(query string) (org []*Organization, err error)
	UpdateOrganization(input *OrganizationInput) (org *Organization, validated *govalidations.Validated, err error)
	SwitchOrganization(orgId string) (err error)
	AcceptSharedGroupRequest(sharedOrgId string, sharedGroupId string) (req *Request, err error)
	RejectSharedGroupRequest(sharedOrgId string, sharedGroupId string) (req *Request, err error)
	//GetSharedGroupRequest(sharedOrgId string, sharedGroupId string) (entry *Entry, err error)
}
