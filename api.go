package qortexapi

import (
	"github.com/sunfmin/govalidations"
	"time"
)

type PublicService interface {
	GetSession(email string, password string) (session string, validated *govalidations.Validated, err error)
	GetAuthUserService(session string) (authUserService AuthUserService, err error)

	// Change Email
	PrepareChangingEmail(memberId string, newEmail string) (changer *EmailChanger, validated *govalidations.Validated, err error)
	ConfirmChangingEmail(token string) (activationToken string, err error)
	CancelChangingEmail(token string) (err error)

	// Sharing Flow
	ChangeEmailToAcceptSharing(token string, newEmail string) (validated *govalidations.Validated, err error)
	GetSharingInviation(sharingInviationToken string, memberId string) (invitation *SharingInvitation, err error)

	ContactUs(input *ContactInput) (contact *ContactInfo, validated *govalidations.Validated, err error)

	// Blog
	GetBlogEntries(doi string, pageNum int, limit int) (blog *Blog, blogEntries []*BlogEntry, totalPageNum int, err error)
	GetBlogEntryBySlug(doi string, slug string) (blog *Blog, blogEntry *BlogEntry, err error)
	CreateExternalComment(doi string, input *EntryInput) (blogEntry *BlogEntry, validated *govalidations.Validated, err error)
	GenerateBlogEntrySlug(doi string, slug string) (validSlug string, err error)
	CreateNewsletter(input *NewsletterInput) (newsletter *Newsletter, validated *govalidations.Validated, err error)
}

// User registered and confirmed email and logged in but haven't join or create any organization.
type AuthMemberService interface {
	SwitchOrganization(orgId string) (err error)
	GetAbandonInfo(abandonOrgId string, memberId string) (info *AbandonInfo, err error)
	GetSharingInviationByToken(sharingInviationToken string) (invitation *SharingInvitation, err error)
	RejectSharingBeforeForwarding(groupId string, email string) (err error)
	RespondSharingRequest(token string, fromOrgId string, fromUserId string, forSharingOrgId string, groupId string) (prefixURL string, validated *govalidations.Validated, err error)
}

// Normal user and joined organization.
type AuthUserService interface {
	GetNewEntry(groupId string) (entry *Entry, err error)
	GetQortexMessages(messsageType string, before string, limit int) (entries []*Entry, err error) // when messageType is empty or equals "all", return all kinds of messages
	CreateBroadcast(input *BroadcastInput) (entry *Entry, validated *govalidations.Validated, err error)
	CreateBroadcastComment(input *BroadcastInput) (entry *Entry, validated *govalidations.Validated, err error)
	GetSharingRequestEntry(entryId string) (entry *Entry, err error)
	GetBroadcast(entryId string) (entry *Entry, err error)
	GetBroadcastComment(entryId string) (entry *Entry, err error)
	UpdateBroadcast(input *BroadcastInput) (entry *Entry, validated *govalidations.Validated, err error)
	UpdateBroadcastComment(input *BroadcastInput) (entry *Entry, validated *govalidations.Validated, err error)
	CreateEntry(input *EntryInput) (entry *Entry, validated *govalidations.Validated, err error)
	CreateTask(input *EntryInput) (entry *Entry, validated *govalidations.Validated, err error)
	CloseTask(entryId string, groupId string) (entry *Task, err error)
	CreateComment(input *EntryInput) (entry *Entry, validated *govalidations.Validated, err error)
	GetComment(entryId string, groupId string) (entry *Entry, err error)
	UpdateComment(input *EntryInput) (entry *Entry, validated *govalidations.Validated, err error)
	UpdateEntry(input *EntryInput) (entry *Entry, validated *govalidations.Validated, err error)
	GetLatestUpdatedEntryIdByTitle(title string, groupId string) (entryId string, err error)
	GetEntry(entryId string, groupId string, updateAtUnixNanoForVersion string, hightlightKeywords string) (entry *Entry, err error)
	DeleteEntry(entryId string, groupId string, dType string) (delType string, err error)
	MuteEntry(entryId string, groupId string) (err error)
	UndoMuteEntry(entryId string, groupId string) (err error)

	GetEntryAttachments(entryId string, groupId string) (attachments []*Attachment, err error)
	GetOtherVersionsComments(entryId string, groupId string, updateAtUnixNanoForVersion string, hightlightKeywords string) (comments []*Entry, err error)

	GetGroupEntries(groupId string, entryType string, before string, limit int) (entries []*Entry, err error)
	GetMyFeedEntries(entryType string, before string, limit int) (entries []*Entry, err error)
	GetNewFeedEntries(entryType string, From string, limit int) (entries []*Entry, err error)
	GetMyTaskEntries(active bool, before string, limit int) (TasksForMe []*Entry, MyCreatedTasks []*Entry, err error)
	GetUserEntries(userId string, entryType string, before string, limit int) (entries []*Entry, err error)

	GetMyChatEntries(before string, limit int) (entries []*Entry, err error)

	GetMyNotificationItems(before string, limit int) (notificationItems []*NotificationItem, err error)
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
	GetDraft(entryId string, groupId string) (entry *Entry, err error)
	DeleteDraft(entryId string, groupId string) (err error)

	//Group related
	GetNewGroup() (group *Group, err error)
	GetGroup(groupId string) (group *Group, err error)
	CreateGroup(input *GroupInput) (group *Group, validated *govalidations.Validated, err error)
	UpdateGroup(input *GroupInput) (validated *govalidations.Validated, err error)
	UpdateGroupLogo(groupId string, logoURL string) (err error)
	// UpdateGroupSlug(id string, slug string) (validated *govalidations.Validated, err error)
	DeleteGroup(groupId string) (err error)
	GetGroupBySlug(slug string) (group *Group, err error)
	GetGroups(keyword string) (groups []*Group, err error)
	GetPublicGroups(keyword string) (groups []*Group, err error)
	AddUserToGroup(groupId string, userId string) (err error)
	RemoveUserFromGroup(groupId string, userId string) (err error)
	GetGroupHeader(groupId string) (header *GroupHeader, err error)
	GetClassifiedGroups() (anouncementGroup *Group, followedGroups []*Group, unFollowedGroups []*Group, err error)

	//User related
	GetOrgUsers(keyword string, startFullName string, limit int) (users []*User, nextFullName string, err error)
	GetGroupUsers(groupId string, keyword string, onlyFollowers bool, startFullName string, limit int) (users []*User, nextFullName string, err error)
	GetUser(userId string) (user *User, err error)
	EnableUser(userId string) (err error)
	DisableUser(userId string) (err error)
	DeleteUser(userId string) (err error)
	PromoteToSuperUser(userId string) (err error)
	DemoteFromSuperUser(userId string) (err error)
	FollowUser(userId string) (err error)
	UnfollowUser(userId string) (err error)
	GetMyFollowingUsers() (followingUsers []*User, err error)
	GetPanelStatus() (panelStatus *PanelStatus, err error)
	GetUserPreferences() (preferences *Preferences, err error)
	UpdateUserPreferences(input *PreferencesInput) (preferences *Preferences, validated *govalidations.Validated, err error)
	GetOrgEmbedUsers() (users []*EmbedUser, err error)
	GetNonStandardGroupEmbedUsers() (groupUsers []*GroupUsers, err error)
	UpdateUserProfile(input *UserProfileInput) (validated *govalidations.Validated, err error)

	// Count related
	MyCount() (myCount *MyCount, err error)
	ReadEntry(entryId, groupId string) (myCount *MyCount, err error)

	//Organization Related
	GetInvitationsInfo() (invitaions []*Invitation, err error)
	OrganizationsInfo(orgIds []string) (orgs []*Organization, err error)
	OrganizationInfo(orgId string) (org *Organization, err error)
	SearchOrganizations(query string) (org []*Organization, err error)
	UpdateOrganization(input *OrganizationInput) (org *Organization, validated *govalidations.Validated, err error)
	SwitchOrganization(orgId string) (err error)
	AcceptSharedGroupRequest(fromOrgId string, sharedOrgId string, sharedGroupId string, fromUserId string) (req *Request, err error)
	RejectSharedGroupRequest(fromOrgId string, sharedOrgIdHex string, sharedGroupIdHex string, fromUserId string) (req *Request, err error)
	//GetSharedGroupRequest(sharedOrgId string, sharedGroupId string) (entry *Entry, err error)

	//Settings related
	GetOrgSettings() (orgSetting *OrgSettings, err error)
	UpdateOrgSettings(orgSettingInput *OrgSettingsInput) (err error)
	CanCreateGroup() (ok bool, err error)
	CanInvitePeople() (ok bool, err error)
	InvitePeople(emails []string) (validated *govalidations.Validated, err error)
	CancelInvitation(email string) (err error)
	ResendInvitation(email string) (err error)
	UpdateMailUpdates(input *MailUpdatesInput) (err error)

	PrepareChangingEmail(newEmail string) (changer *EmailChanger, validated *govalidations.Validated, err error)
	ConfirmChangingEmail(token string) (err error)
	UpdateAccount(input *MemberAccountInput) (validated *govalidations.Validated, err error)

	SendSharingInvitation(groupId string, email string, isResend bool) (si *SharingInvitation, validated *govalidations.Validated, err error)
	GetSharingInvitationItems(groupId string) (sis []*SharingInvitation, err error)
	CancelSharing(groupId string, email string) (err error)
	StopSharingGroup(GroupId string, toStopOrgId string) (err error)
	LeaveSharingGroup(GroupId string) (err error)

	//chat
	ShareChat(input *ShareChatInput) (chatEntry *Entry, validated *govalidations.Validated, err error)
}
