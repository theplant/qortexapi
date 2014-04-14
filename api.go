package qortexapi

type PublicService interface {
	GetSession(email string, password string, locale string) (session string, err error)
	GetAuthUserService(session string, orgId string) (authUserService AuthUserService, err error)
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

	// Demo related
	CreateSandboxOrg(idOrQortexURL string) (r *Organization, err error)
	CreateSandboxMember(firstName string, lastName string, avatarURL string) (r *Member, err error)
}

// User registered and confirmed email and logged in but haven't join or create any organization.
type AuthMemberService interface {
	// For creating new Organization
	GetNewOrganization(memberId string) (org *Organization, err error)
	GetMyOrganizations() (orgs []*Organization, err error)
	CreateOrganization(input *OrganizationInput) (apiOrg *Organization, err error)
	JoinOrganization(orgId string) (err error)
	LeaveOrganization(orgId string) (err error)

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
	CloseTask(entryId string, groupId string, taskId string) (entry *Task, err error)
	UpdateTask(taskInput *TaskInput) (task *Task, err error)
	CreateComment(input *EntryInput) (entry *Entry, err error)
	GetComment(entryId string, groupId string, languageCode string) (entry *Entry, err error)
	EditComment(entryId string, groupId string, languageCode string) (entry *Entry, err error)
	UpdateComment(input *EntryInput) (entry *Entry, err error)
	UpdateEntry(input *EntryInput) (entry *Entry, err error)
	GetLatestUpdatedEntryIdByTitle(title string, groupId string) (entryId string, err error)
	GetTitle(groupId string, entryId string) (title string, err error)                                                                                              //When languageCode is empty, use default
	GetEntry(entryId string, groupId string, updateAtUnixNanoForVersion string, hightlightKeywords string, languageCode string) (entry *Entry, err error)           //When languageCode is empty, use default
	SwitchEntryVersion(entryId string, groupId string, updateAtUnixNanoForVersion string, hightlightKeywords string, languageCode string) (entry *Entry, err error) //When languageCode is empty, use default
	EditEntry(entryId string, groupId string, updateAtUnixNanoForVersion string, hightlightKeywords string, languageCode string) (entry *Entry, err error)          //When languageCode is empty, use default
	SwitchEntryLanguage(entryId string, groupId string, languageCode string) (entry *Entry, err error)
	GetKnowledgeOverview(groupId string, languageCode string) (r *KnowledgeOverview, err error) //When languageCode is empty, use default
	SwitchKnowledgeOverviewVersion(groupId string, entryId string, languageCode string) (r *KnowledgeOverview, err error)
	UpdateKnowledgeOverview(input *KnowledgeOverviewInput) (r *KnowledgeOverview, err error)
	GetEntryToTranslate(entryId string, groupId string) (entry *Entry, err error)
	GetWikiSectionToTranslate(entryId string, groupId string) (entry *KnowledgeOverview, err error)

	// dType "all": delete all versions of the entry, "version": delete current version of the entry
	DeleteEntry(entryId string, groupId string, dType string) (delType string, err error)
	MuteEntry(entryId string, groupId string) (err error)
	UndoMuteEntry(entryId string, groupId string) (err error)
	GetMachineTranslatableLangauges() (options *LanguageSelector, err error)
	MachineTranslate(entryId string, groupId string, currentLang string, targetlang string) (translatedThread *TranslatedThread, err error)
	MachineTranslateWikiSection(entryId string, groupId string, targetlang string) (translatedThread *TranslatedThread, err error)
	OriginalThread(entryId string, groupId string) (translatedThread *TranslatedThread, err error)

	GetEntryAttachments(entryId string, groupId string) (attachments []*Attachment, err error)
	GetDocViewSession(doi string, groupId string, attachmentId string) (sessionId string, err error)
	DeleteEntryAttachment(doi string, groupId string, attachmentId string, ownerId string) (attachments []*Attachment, err error)
	GetOtherVersionsComments(entryId string, groupId string, updateAtUnixNanoForVersion string) (comments []*Entry, err error)
	GetOtherVersionsTaskLogs(entryId string, groupId string, updateAtUnixNanoForVersion string) (taskLogs []*TaskLog, err error)

	GetGroupEntries(groupId string, entryType string, before string, limit int, withComments bool) (entries []*Entry, err error)
	GetMyFeedEntries(entryType string, before string, limit int, withComments bool) (entries []*Entry, err error)
	GetGroupAside() (ga *GroupAside, err error)
	GetNewGroupAside() (nga *NewGroupAside, err error)

	GetMyFeedEntriesLite(before string, limit int) (entries []*Entry, err error)

	// Get Unread entries for type entryType since the fromTimeUnixNano unix nanoseconds, and return max limit entries.
	// entryType could be:	"post", "knowledge", "task"
	GetNewFeedEntries(entryType string, fromTimeUnixNano string, limit int) (entries []*Entry, err error)
	GetUserEntries(userId string, entryType string, before string, limit int) (entries []*Entry, err error)

	GetMyNotifications(before string, limit int) (mynotis *MyNotifications, err error)
	GetMyNotificationItems(before string, limit int) (notificationItems []*NotificationItem, err error)
	MarkAllAsRead(groupId string) (mycount *MyCount, err error)

	// watchlist related
	GetWatchList(before string, limit int) (watchlist *WatchList, err error)
	AddToWatchList(entryId string, groupId string, remindMode string) (err error)
	StopWatching(entryId string, groupId string) (err error)
	RemindMe() (reminded bool, err error)
	StartSmartReminding(groupId string, watchItemId string) (stopped bool, err error)
	StopReminding(groupId string, watchItemId string) (stopped bool, err error)

	// Like action
	UpdateLike(input *LikeInput) (entry *Entry, err error)

	// draft related
	GetDraftList(before string, limit int) (draftlist *DraftList, err error)
	GetDraft(entryId string, groupId string) (entry *Entry, err error)
	DeleteDraft(entryId string, groupId string) (err error)
	ChooseMarkdownEditor() (err error)
	ChooseStyledEditor() (err error)

	// Group related
	GetNewGroup() (group *Group, err error)
	GetGroup(groupId string) (group *Group, err error)
	CreateGroup(input *GroupInput) (group *Group, err error)
	UpdateGroup(input *GroupInput) (group *Group, err error)
	UpdateGroupLogo(groupId string, logoURL string) (err error)
	DeleteGroup(groupId string) (err error)
	GetGroupBySlug(slug string) (group *Group, err error)
	GetGroups(keyword string) (groups []*Group, err error)
	GetPublicGroups(keyword string) (groups []*Group, err error)
	AddUserToGroup(groupId string, userId string) (err error)
	RemoveUserFromGroup(groupId string, userId string) (err error)
	GetClassifiedGroups() (anouncementGroup *Group, smGroup *Group, followedNormalGroups []*Group, followedSharedGroups []*Group, unFollowedNormalGroups []*Group, unFollowedSharedGroups []*Group, err error)
	GetAllGroupCollections() (gcs []*GroupCollection, err error)
	ToggleGroupArchiving(gids string, signal bool) (err error)
	BulkUpdateTasksInGroup(groupId string, taskPwMap []*TaskPwMap, taskInputs []*TaskInput, markerInputs []*ToDoMarkerInput) (err error)
	UpdateCollection(gId, colId, colName string) (group *Group, err error)

	// User related
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
	SetPreferredLanguages(languageCodes []string) (err error)
	ToggleGroupCol(gtype int, colIdStr string) (err error)

	// Count related
	GetMyCount() (myCount *MyCount, err error)
	ReadEntry(entryId, groupId string) (myCount *MyCount, err error)
	ReadNotificationItem(itemId, groupId string) (myCount *MyCount, err error)

	// Organization Related
	GetJoinOrgInvitations() (invitations []*Invitation, err error)
	GetOrganization(orgId string) (org *Organization, err error)
	GetOrganizations(orgIds []string) (orgs []*Organization, err error)
	GetMyOrgsUnreadInfo() (unreadInfo []*OrgUnreadInfo, err error)
	GetMyJoinedOrganizations() (orgs []*Organization, err error)
	GetCurrentOrganization() (org *Organization, err error)
	SearchOrganizations(keyword string) (orgs []*SearchOrganization, err error)
	UpdateOrganization(input *OrganizationInput) (org *Organization, err error)
	SwitchOrganization(orgId string) (err error)
	MarkAsSampleOrg() (err error)
	MarkAsStandardOrg() (err error)
	GetSampleOrgs() (orgs []*Organization, err error)
	GetSandboxOrgs() (orgs []*Organization, err error)
	DeleteSandboxOrg(orgId string) (err error)
	DeleteAllSandboxOrg() (err error)

	AcceptShareRequestByAdmin(requestId string) (err error)
	RejectShareRequestByAdmin(requestId string) (err error)

	StartTrial() (err error)

	// Settings related
	GetOrgSettings() (orgSetting *OrgSettings, err error)
	UpdateOrgSettings(orgSettingInput *OrgSettingsInput) (err error)
	CanCreateGroup() (ok bool, err error)
	CanInvitePeople() (ok bool, err error)
	InvitePeople(emails []string, allowEmpty bool, skipInvalidEmail bool, customMessage string) (sendedEmails []string, err error)
	CancelInvitation(email string) (err error)
	ResendInvitation(email string) (err error)
	UpdateGroupAdvancedToDoSettings(gId, settings string) (err error)

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

	// preferences
	DismissPresentationTip() (err error)
	DismissTutorialsTip() (err error)

	// chat
	GetMyChatEntries(before string, limit int) (entries []*Entry, err error)
	GetPrivateChat(conversationId string, searchKeyWords string) (chatEntry *Entry, err error)
	OpenConversation(cid, fromJid, toJid string) (err error)

	// Qortex Support
	CreateQortexSupport(input *QortexSupportInput) (entry *Entry, err error)
	CreateQortexSupportComment(input *QortexSupportInput) (entry *Entry, err error)
	GetQortexSupport(entryId string, languageCode string) (entry *Entry, err error) //When languageCode is empty, use default
	GetQortexSupportComment(entryId string) (entry *Entry, err error)
	UpdateQortexSupport(input *QortexSupportInput) (entry *Entry, err error)
	UpdateQortexSupportComment(input *QortexSupportInput) (entry *Entry, err error)
	GetQortexSupportHelpLink(title string) (link string, err error)

	// Advand Task Related
	NewTask(groupId string) (task *Task, err error)
	EditTask(groupId string, taskId string) (task *Task, err error)
	GetAdvancedTask(taskId string) (at *AdvancedTask, err error)
	ClaimTask(taskId string, groupId string) (task *Task, err error)
	UpdateSimpleTask(input *TaskInput) (task *Task, err error)
	GetTasksForMe() (needActionTasks []*TaskOutline, groupTasks []*GroupTasksOutline, err error)
	GetOpenTasksIMade() (groupTasks []*GroupTasksOutline, err error)
	GetClosedTasksIMade(before string, limit int) (tasks []*TaskOutline, err error)
	GetOpenTasksIWorkedOn() (groupTasks []*GroupTasksOutline, err error)
	GetClosedTasksIWorkedOn(before string, limit int) (tasks []*TaskOutline, err error)
	GetUserTasks(userId string, groupId string) (needActionTasks []*TaskOutline, groupTasks []*GroupTasksOutline, err error)

	GetGroupGeneralSettingPage(gId string) (page *GroupGeneralSettingPage, err error)
	GetGroupUsersPage(gId string) (page *GroupUsersPage, err error)
	GetGroupSharingExternallyPage(gId string) (page *GroupSharingExternallyPage, err error)

	// [Group] Advanced To-Dos Related
	GetGroupAdvancedToDoSetting(gId string) (page *GroupAdvancedSettingPage, err error)
	AllOpenAdvancedToDosInGroup(groupId string) (bucket *OpenAdvancedToDosBucket, err error)
	AllOpenAdvancedToDosGroupingByUserInGroup(groupId string) (atos []*OpenAdvancedToDosPage, err error)
	AllOpenAdvancedToDosGroupingByStatusInGroup(groupId string) (page []*OpenAdvancedToDosBucket, apiGroup *Group, err error)
	AllOpenAdvancedToDosGroupingByLabelInGroup(groupId string) (page []*OpenAdvancedToDosBucket, apiGroup *Group, err error)
	AllOpenBasicToDosInGroup(groupId string) (taskOutlines []*TaskOutline, err error)
	AllOpenBasicToDosGroupingByUserInGroup(groupId string) (atos []*BasicOpenToDoOutlines, err error)

	AllClosedBasicToDosInGroup(groupId string, afterTimeS string) (taskOutlines []*TaskOutline, err error)
	AllClosedAdvancedToDosInGroup(groupId string) (closedOutlines []*ClosedAdvancedToDoOutline, err error)
	MoreClosedAdvancedToDosWithStatusInGroup(groupId string, status int, afterTime string) (taskOutlines []*TaskOutline, apiGroup *Group, hasMore bool, err error)
	CountOfClosedToDosInGroup(ttype int, groupId string) (count int, err error)
	CountOfActionNeededToDosInGroup(gid string) (count int, err error)
	ToDoCSV(groupId string) (todos []*ToDoCSVItem, err error)

	// Apple device service
	RegisterAppleDeviceForUser(userId string, token string) (err error)
	UnregisterAppleDeviceForUser(userId string, token string) (err error)

	//payment
	GetPaymentSession() (session string, err error)
	CanSeeBilling() (yes bool, err error)
	GetBillingInfo() (billing *BillingInfo, err error)
	GetReceiptInfo(id string) (receipt *ReceiptInfo, err error)
	SyncBilling() (err error)
	SyncBillingDetails() (err error)
	// SyncPastPayments() (err error)
	ValidatePayment() (err error)
	CancelSubscription() (err error)
	DismissPaymentTips() (err error)

	GetContactUsInfo() (info *ContactUsInfo, err error)
	// For Mobile Specifically
	GetInitInfo() (info *InitInfo, err error)
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

	GetTotalOnlineUsers() (embedUsers []*EmbedUser, err error)

	MarkOrgFreeOrPay(orgId string) (free bool, err error)

	GetOrgPayment() (orgPaymentInfos []*OrgPaymentInfo, err error)

	GetPaymentHistory(orgId string) (history []OrgPaymentHistory, err error)

	SetTrial(orgId string, deadLine string) (err error)

	SetExpiredAt(orgId string, deadLine string) (err error)
	//for test
	SendPaymentWarnEmail(orgId string) (err error)
}
