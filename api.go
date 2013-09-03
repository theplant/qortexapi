package qortexapi

import (
	"time"
)

type PublicService interface {
	GetSession(email string, password string, locale string) (session string, err error)
	GetAuthUserService(session string) (authUserService AuthUserService, err error)
	GetAuthorizedAdmin(session string) (apiEmbedUser EmbedUser, err error)
	GetAuthAdminService(session string) (authAdminService AuthAdminService, err error)

	// Find Password
	FindPassword(email string) (err error)
	ResetPassword(token string, password string, confirmedPassword string) (memberId string, email string, err error)

	// Change Email
	PrepareChangingEmail(memberId string, newEmail string, sharingToken string, invitationToken string) (changer *EmailChanger, err error)
	ConfirmChangingEmail(token string) (activationToken string, sharingToken string, err error)
	CancelChangingEmail(token string) (err error)

	// Sharing Flow
	ChangeEmailToAcceptSharing(token string, newEmail string) (err error)

	GetShareRequest(token string, memberId string) (shareRequest *ShareRequest, err error)

	ContactUs(input *ContactInput) (contact *ContactInfo, err error)

	// Blog
	GetBlogEntries(doi string, pageNum int, limit int) (blog *Blog, blogEntries []*BlogEntry, totalPageNum int, err error)
	GetBlogEntryBySlug(doi string, slug string) (blog *Blog, blogEntry *BlogEntry, err error)
	GenerateBlogEntrySlug(doi string, slug string) (validSlug string, err error)
	CreateNewsletter(input *NewsletterInput) (newsletter *Newsletter, err error)
	RequestNewSignupToken(email string) (err error)
	RequestNewInvitationToken(orgId string, email string) (err error)
	RequestNewSharingToken(email string) (err error)

	InviteMe(organizationId string, email string) (err error)

	// Signup
	RequestSignup(email string) (err error)
}

// User registered and confirmed email and logged in but haven't join or create any organization.
type AuthMemberService interface {
	SwitchOrganization(orgId string) (err error)
	GetAbandonInfo(abandonOrgId string, memberId string) (info *AbandonInfo, err error)

	GetShareRequest(token string) (shareRequest *ShareRequest, err error)
	RejectShareRequestByInvitee(token string) (err error)
	AcceptShareRequestByInvitee(token string, toOrgId string) (err error)
}

// Normal user and joined organization.
type AuthUserService interface {
	GetNewEntry(groupId string) (entry *Entry, err error)
	GetNewChatEntry(chatId string) (entry *Entry, err error)
	GetQortexSupportEntries(before string, limit int, withComments bool) (entries []*Entry, err error)
	CreateEntry(input *EntryInput) (entry *Entry, err error)
	CreateTask(input *EntryInput) (entry *Entry, err error)
	CloseTask(entryId string, groupId string) (entry *Task, err error)
	CreateComment(input *EntryInput) (entry *Entry, err error)
	GetComment(entryId string, groupId string) (entry *Entry, err error)
	UpdateComment(input *EntryInput) (entry *Entry, err error)
	UpdateEntry(input *EntryInput) (entry *Entry, err error)
	GetLatestUpdatedEntryIdByTitle(title string, groupId string) (entryId string, err error)
	GetEntry(entryId string, groupId string, updateAtUnixNanoForVersion string, hightlightKeywords string) (entry *Entry, err error)

	//dType "all": delete all versions of the entry, "version": delete current version of the entry
	DeleteEntry(entryId string, groupId string, dType string) (delType string, err error)
	MuteEntry(entryId string, groupId string) (err error)
	UndoMuteEntry(entryId string, groupId string) (err error)

	GetEntryAttachments(entryId string, groupId string) (attachments []*Attachment, err error)
	GetOtherVersionsComments(entryId string, groupId string, updateAtUnixNanoForVersion string, hightlightKeywords string) (comments []*Entry, err error)

	GetGroupEntries(groupId string, entryType string, before string, limit int, withComments bool) (entries []*Entry, err error)
	GetMyFeedEntries(entryType string, before string, limit int, withComments bool) (entries []*Entry, err error)
	GetGroupAside() (ga *GroupAside, err error)

	// Get Unread entries for type entryType since the fromTimeUnixNano unix nanoseconds, and return max limit entries.
	// entryType could be:	"post", "knowledge", "task"
	GetNewFeedEntries(entryType string, fromTimeUnixNano string, limit int) (entries []*Entry, err error)
	GetMyTaskEntries(active bool, before string, limit int) (TasksForMe []*Entry, MyCreatedTasks []*Entry, err error)
	GetUserEntries(userId string, entryType string, before string, limit int) (entries []*Entry, err error)

	GetMyNotificationItems(before string, limit int) (notificationItems []*NotificationItem, err error)
	MarkAllAsRead(groupId string) (mycount *MyCount, err error)

	// watchlist related
	GetWatchList(before time.Time, limit int) (watchlist *WatchList, err error)
	AddToWatchList(entryId string, groupId string, remindMode string) (added bool, err error)
	StopWatching(entryId string, groupId string) (stopped bool, err error)
	ReadWatching(entryId string, groupId string) (err error)
	RemindMe() (reminded bool, err error)
	StartSmartReminding(groupId string, watchItemId string) (stopped bool, err error)
	StopReminding(groupId string, watchItemId string) (stopped bool, err error)

	// Like action
	UpdateLike(input *LikeInput) (entry *Entry, err error)

	// draft related
	GetDraftList(before time.Time, limit int) (draftlist *DraftList, err error)
	GetDraft(entryId string, groupId string) (entry *Entry, err error)
	DeleteDraft(entryId string, groupId string) (err error)

	//Group related
	GetNewGroup() (group *Group, err error)
	GetGroup(groupId string) (group *Group, err error)
	CreateGroup(input *GroupInput) (group *Group, err error)
	UpdateGroup(input *GroupInput) (err error)
	UpdateGroupLogo(groupId string, logoURL string) (err error)
	// UpdateGroupSlug(id string, slug string) (validated *govalidations.Validated, err error)
	DeleteGroup(groupId string) (err error)
	GetGroupBySlug(slug string) (group *Group, err error)
	GetGroups(keyword string) (groups []*Group, err error)
	GetPublicGroups(keyword string) (groups []*Group, err error)
	AddUserToGroup(groupId string, userId string) (err error)
	RemoveUserFromGroup(groupId string, userId string) (err error)
	GetGroupHeader(groupId string) (header *GroupHeader, err error)
	GetClassifiedGroups() (anouncementGroup *Group, smGroup *Group, followedNormalGroups []*Group, followedSharedGroups []*Group, unFollowedNormalGroups []*Group, unFollowedSharedGroups []*Group, err error)

	//User related
	GetAuthUser() (user *User, err error)
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
	UpdateUserPreferences(input *PreferencesInput) (preferences *Preferences, err error)
	GetOrgEmbedUsers() (users []*EmbedUser, err error)
	GetNonStandardGroupEmbedUsers() (groupUsers []*GroupUsers, err error)
	UpdateUserProfile(input *UserProfileInput) (err error)

	// Count related
	GetMyCount() (myCount *MyCount, err error)
	ReadEntry(entryId, groupId string) (myCount *MyCount, err error)

	//Organization Related
	GetJoinOrgInvitations() (invitations []*Invitation, err error)
	GetOrganization(orgId string) (org *Organization, err error)
	GetOrganizations(orgIds []string) (orgs []*Organization, err error)
	GetMyOrgsUnreadInfo() (unreadInfo []*OrgUnreadInfo, err error)
	GetMyJoinedOrganizations() (orgs []*Organization, err error)
	GetCurrentOrganization() (org *Organization, err error)
	SearchOrganizations(keyword string) (orgs []*Organization, err error)
	UpdateOrganization(input *OrganizationInput) (org *Organization, err error)
	SwitchOrganization(orgId string) (err error)

	AcceptShareRequestByAdmin(requestId string) (err error)
	RejectShareRequestByAdmin(requestId string) (err error)

	//Settings related
	GetOrgSettings() (orgSetting *OrgSettings, err error)
	UpdateOrgSettings(orgSettingInput *OrgSettingsInput) (err error)
	CanCreateGroup() (ok bool, err error)
	CanInvitePeople() (ok bool, err error)
	InvitePeople(emails []string, skipInvalidEmail bool, customMessage string) (sendedEmails []string, err error)
	CancelInvitation(email string) (err error)
	ResendInvitation(email string) (err error)
	ChangeLocale(localeName string) (err error)

	// TODO: mail-updates: remove it
	// UpdateMailUpdates(input *MailUpdatesInput) (err error)
	UpdateMailPreference(input *MailPreferenceInput) (err error)

	PrepareChangingEmail(newEmail string) (changer *EmailChanger, err error)
	ConfirmChangingEmail(token string) (err error)
	UpdateAccount(input *MemberAccountInput) (err error)

	SendShareRequest(groupId string, email string) (shareRequest *ShareRequest, err error)
	GetShareRequests(groupId string) (sis []*ShareRequest, err error)
	CancelShareRequest(requestId string) (err error)

	StopSharingGroup(requestId string) (err error)

	//preferences
	DismissPresentationTip() (err error)

	//chat
	GetMyChatEntries(before string, limit int) (entries []*Entry, err error)
	GetPrivateChat(conversationId string, searchKeyWords string) (chatEntry *Entry, err error)

	// Qortex Support
	CreateQortexSupport(input *QortexSupportInput) (entry *Entry, err error)
	CreateQortexSupportComment(input *QortexSupportInput) (entry *Entry, err error)
	GetQortexSupport(entryId string) (entry *Entry, err error)
	GetQortexSupportComment(entryId string) (entry *Entry, err error)
	UpdateQortexSupport(input *QortexSupportInput) (entry *Entry, err error)
	UpdateQortexSupportComment(input *QortexSupportInput) (entry *Entry, err error)
}

type AuthAdminService interface {
	// Get the overall statistics (excluding ThePlant)
	GetTotalStats() (totalStat *TotalStats, err error)
	// Get the weekly statistics (excluding ThePlant)
	GetWeeklyTotalStats() (totalStat *TotalStats, err error)
	// Get all the Organizations' statistics
	GetOrgStats() (orgStats []*OrgStats, err error)
	// Get all closed beta access requests
	GetAccessRequests() (accessReqs []*AccessReq, err error)
	// Approve user access request for closed beta
	ApproveAccess(email string) (err error)
	// Resend the approved mail
	ResendApprovedMail(email string) (err error)
	// Get all members
	GetAllMembers() (members []*Member, err error)
	// Ignore the access
	IgnoreAccess(email string) (err error)

	GetAutoApproveAccess() (enabled bool, err error)

	SetAutoApproveAccess(enable bool) (err error)

	GetMarketableUsers() (memberInfos []*MarketableMemberInfo, err error)
}
