package qortexapi

import (
	"github.com/sunfmin/govalidations"
	"time"
)

type Global interface {
	GetSession(email string, password string) (session string, validated *govalidations.Validated, err error)
	GetAuthUserService(session string) (authUserService AuthUserService, err error)
}

type NoAuthUserService interface {
	CancelChangingEmail(token string) (err error)
	ChangeEmail(token string) (activationToken string, err error)
	PrepareChangeEmail(memberId string, newEmail string) (r *EmailChanger, validated *govalidations.Validated, err error)
	GetSharingInviation(sharingInviationToken string) (r *SharingInvitation, err error)

	ResponseSharingRequest(token string, fromOrgId string, forSharingOrgId string, groupId string) (prefixURL string, validated *govalidations.Validated, err error)

	/* Blog */
	BlogEntries(pageNum int, limit int) (totalPageNum int, r []*Entry, err error)
	BlogEntryBySlug(slug string) (r *Entry, err error)
	CreateExternalComment(input *EntryInput) (r *Entry, validated *govalidations.Validated, err error)
	CheckSlug(slug string) (validSlug string, err error)
	CreateNewsletter(input *NewsletterInput) (r *Newsletter, validated *govalidations.Validated, err error)

}

type AuthUserService interface {
	NewEntry(groupId string) (entry *Entry, err error)
	QortexMessages(messsageType string, before string, limit int) (entries []*Entry, err error) // when messageType is empty or equals "all", return all kinds of messages
	CreateBroadcast(input *BroadcastInput) (entry *Entry, validated *govalidations.Validated, err error)
	CreateBroadcastComment(input *BroadcastInput) (entry *Entry, validated *govalidations.Validated, err error)
	GetRequest(entryId string) (entry *Entry, err error)
	EditBroadcast(entryId string) (entry *Entry, err error)
	EditBroadcastComment(entryId string) (entry *Entry, err error)
	UpdateBroadcast(input *BroadcastInput) (entry *Entry, validated *govalidations.Validated, err error)
	UpdateBroadcastComment(input *BroadcastInput) (entry *Entry, validated *govalidations.Validated, err error)
	CreatePost(input *EntryInput) (entry *Entry, validated *govalidations.Validated, err error)
	// GetPost(entryId string, groupId string) (entry *Entry, err error)
	// UpdatePost(input *EntryInput) (entry *Entry, validated *govalidations.Validated, err error)
	CreateTask(input *EntryInput) (entry *Entry, validated *govalidations.Validated, err error)
	// GetTask(entryId string, groupId string) (entry *Entry, err error)
	CloseTask(entryId string, groupId string) (entry *Task, err error)
	CreateComment(input *EntryInput) (entry *Entry, validated *govalidations.Validated, err error)
	GetComment(entryId string, groupId string) (entry *Entry, err error)
	UpdateComment(input *EntryInput) (entry *Entry, validated *govalidations.Validated, err error)
	// CreateWiki(input *EntryInput) (entry *Entry, validated *govalidations.Validated, err error)
	// GetWiki(entryId string, groupId string, versionUpdateat string) (entry *Entry, err error)
	// GetWikiByTitle(title string, groupId string, updateAtUnixNano string, searchKeyWords string) (entry *Entry, err error)
	UpdateEntry(input *EntryInput) (entry *Entry, validated *govalidations.Validated, err error)
	GetLatestUpdatedEntryIdByTitle(title string, groupId string) (r string, err error)
	GetEntry(entryId string, groupId string, updateAtUnixNano string, searchKeyWords string) (entry *Entry, err error)
	DeleteEntry(entryId string, groupId string, dType string) (delType string, err error)

	EntryAttachments(entryId string, groupId string) (attachments []*Attachment, err error)
	OtherComments(entryId string, groupId string, versionUpdateat string) (comments []*Entry, err error)

	// GroupUnreadEntryIds(entryIds []string, groupId string) (unreadEntryIds []string, err error)
	// UnreadEntryIds(entryIds []string, groupIds []string) (unreadEntryIds []string, err error)
	GroupEntries(groupId string, entryType string, before string, limit int) (entries []*Entry, err error)
	MyFeedEntries(entryType string, before string, limit int) (entries []*Entry, err error)
	NewFeedEntries(entryType string, From string, limit int) (entries []*Entry, err error)
	// NewMyFeedEntries(entryType string, after time.Time, limit int) (entries []*Entry, err error)
	MyTaskEntries(active bool, before string, limit int) (TasksForMe []*Entry, MyCreatedTasks []*Entry, err error)
	UserEntries(userId string, entryType string, before string, limit int) (entries []*Entry, err error)

	MyChatEntries(before string, limit int) (entries []*Entry, err error)
	// LoadEntry(groupId string, entryId string) (g *Group, entry *Entry, err error)

	MyNotificationItems(before string, limit int) (notificationItems []*NotificationItem, err error)
	MarkAllAsRead(groupId string) (mycount *MyCount, err error)

	// watchlist related
	GetWatchList(before time.Time, limit int) (watchlist *WatchList, err error)
	AddToWatchList(entryId string, groupId string) (added bool, err error)
	StopWatching(entryId string, groupId string) (stopped bool, err error)
	ReadWatching(entryId string, groupId string) (err error)

	// Like action
	UpdateLike(input *LikeInput) (entry *Entry, err error)

	// draft related
	GetDraftList(before time.Time, limit int) (draftlist *DraftList, err error)
	DeleteDraft(entryId string, groupId string) (err error)

	//Group related
	NewGroup() (group *Group, err error)
	GetGroup(groupId string) (group *Group, err error)
	CreateGroup(input *GroupInput) (group *Group, validated *govalidations.Validated, err error)
	UpdateGroup(input *GroupInput) (validated *govalidations.Validated, err error)
	UpdateGroupLogo(groupId string, logoURL string) (err error)
	// UpdateGroupSlug(id string, slug string) (validated *govalidations.Validated, err error)
	DeleteGroup(groupId string) (err error)
	GroupBySlug(slug string) (group *Group, err error)
	ConvertToSharedGroup(input *GroupInput) (validated *govalidations.Validated, err error)
	GetAllGroups(keyword string) (groups []*Group, err error)
	GetPublicGroups(keyword string) (groups []*Group, err error)
	AddUserToGroup(groupId string, userId string) (err error)
	RemoveUserFromGroup(groupId string, userId string) (err error)
	GetGroupHeaderItem(groupId string) (ghi *GroupHeaderItem, err error)
	ClassifyMyGroups() (publicGroup *Group, followedGroups []*Group, unFollowedGroups []*Group, err error)

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
	PanelStatus() (panelStatus *PanelStatus, err error)
	Preferences() (preference *Preferences, err error)
	UpdatePreferences(input *PreferencesInput) (preference *Preferences, validated *govalidations.Validated, err error)
	AllEmbedUsers() (users []*EmbedUser, err error)
	GroupEmbedUsers() (groupUsers []*GroupUsers, err error)

	// Count related
	MyCount() (myCount *MyCount, err error)
	ReadEntry(entryId, groupId string) (myCount *MyCount, err error)

	//Organization Related
	OrganizationsInfo(orgIds []string) (orgs []*Organization, err error)
	OrganizationInfo(orgId string) (org *Organization, err error)
	SearchOrganizations(query string) (org []*Organization, err error)
	UpdateOrganization(input *OrganizationInput) (org *Organization, validated *govalidations.Validated, err error)
	SwitchOrganization(orgId string) (err error)
	AcceptSharedGroupRequest(sharedOrgId string, sharedGroupId string) (req *Request, err error)
	RejectSharedGroupRequest(sharedOrgId string, sharedGroupId string) (req *Request, err error)
	//GetSharedGroupRequest(sharedOrgId string, sharedGroupId string) (entry *Entry, err error)

	//Settings related
	GetOrgSettings() (orgSetting *OrgSettings, err error)
	UpdateOrgSettings(orgSettingInput *OrgSettingsInput) (err error)
	CanCreateGroup() (r bool)
	CanInvitePeople() (r bool)
	InvitePeople(emails []string) (validated *govalidations.Validated, err error)

	PrepareChangeEmail(newEmail string) (r *EmailChanger, validated *govalidations.Validated, err error)
	ChangeEmail(token string) (err error)
	UpdateAccount(input *MemberAccountInput) (validated *govalidations.Validated, err error)

	SendSharingInvitation(groupId string, email string) (validated *govalidations.Validated, err error)
}
