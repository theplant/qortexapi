package qortexapi

import (
	"html/template"
	"time"
)

type OrgSettings struct {
	AllowUsersCreateGroups bool
	AllowUsersInvitePeople bool
}

type Organization struct {
	Id                       string
	Name                     string
	QortexURL                string
	Summary                  string
	LogoURL                  string
	Address                  string
	Phone                    string
	Website                  string
	Domains                  []string
	RestrictSubscriptionMail bool
}

type SharingInvitationItem struct {
	FromOrgId       string
	FromUserId      string
	GroupId         string
	IsNewAccount    bool
	Email           string
	Token           string
	JoinedOrgs      []EmbedOrg
	IsAccepted      bool
	IsRejected      bool
	IsPending       bool
	IsForwarded     bool
	PendingDuration string
	ToOrgName       string
}

type User struct {
	Id                   string
	Email                string
	Firstame             string
	LastName             string
	FullName             string
	Title                string
	Avatar               string
	JID                  string
	Timezone             string
	IsSuperUser          bool
	IsSharedUser         bool
	OrgId                string
	OriginalOrgId        string
	PrefixURL            string
	ProfileURL           string
	IsLoggedInUser       bool
	IsAvailable          bool
	IsDisabled           bool
	IsDeleted            bool
	FromSharedGroup      bool
	FromOrganizationName string
	Editable             bool
	Followable           bool
	FollowedByMe         bool
	AdminingAGroup       bool
	AdminingTheGroup     bool
	FollowingTheGroup    bool
	Department           string
	Location             string
	FollowingGroups      []*Group
	Preferences          *Preferences
}

type GroupUsers struct {
	GroupId    string
	EmbedUsers []EmbedUser
}

type Preferences struct {
	Timezone                 string
	TimezoneOffset           string
	PreferFullName           bool
	EnterForNewLine          bool
	AsideGroupsCollapse      bool
	AsideOtherGroupsCollapse bool
	ShowMarkUnreadThreshold  int
	AdminModeOn              bool
	PreferMarkdown           bool
	AutoFollowPublicGroup    bool
}

type EmbedOrg struct {
	Id            string
	Name          string
	LogoURL       string
	NoNeedToShare bool
}

type EmbedUser struct {
	Id             string
	Email          string
	Name           string
	Title          string
	Avatar16       string
	Avatar32       string
	JID            string
	Timezone       string
	IsSuperUser    bool
	IsShare        bool
	OrganizationId string
	OriginalOrgId  string
	ProfileURL     template.HTMLAttr
}

type PanelStatus struct {
	AsideGroupsCollapse      bool
	AsideOtherGroupsCollapse bool
	HasToDo                  bool
	HasDraft                 bool
	HasWatchList             bool
	HasChat                  bool
	ShowMarkUnreadThreshold  int
}

type Group struct {
	Id                  string
	Name                string
	Description         string
	GType               string
	LogoURL             string
	IconName            string
	Link                string
	Slug                string
	Author              EmbedUser
	IsAdmin             bool
	IsPrivate           bool
	Editable            bool
	Managable           bool
	FollowedByMe        bool
	AdministratedByMe   bool
	IsShared            bool
	HostOrgName         string
	IsDispayHostOrgName bool
	EntriesCount        int
	FollowersCount      int
	IsAnnoucement       bool
	GroupOwners         []EmbedUser
}

type GroupSelectorItem struct {
	Id         string
	Name       string
	IsSelected bool
}

type GroupSelector struct {
	Header            template.HTML
	SelectedGroupId   string
	SysMessage        *GroupSelectorItem
	FollowingGroups   []*GroupSelectorItem
	UnFollowingGroups []*GroupSelectorItem
}

type Attachment struct {
	Id            string
	OwnerId       []string
	Category      string
	Filename      string
	ShortFilename string
	ContentType   string
	ContentId     string
	MD5           string
	ContentLength int64
	Error         string
	GroupId       []string
	UploadTime    time.Time
	Width         int
	Height        int

	URL          template.HTMLAttr
	ImageIconURL template.HTMLAttr
	FileIconURL  template.HTMLAttr
	HumanSize    string
	IsImage      bool
	FileKind     string
}

type Task struct {
	IsTaskOwner       bool
	IsTaskAssignee    bool
	IsOthers          bool
	IsCurrentUserDone bool

	IsAcknowledgement bool
	IsTodoForOne      bool
	IsTodoForAll      bool

	IsCompleted bool
	IsClosed    bool

	IsDueToday bool
	IsOverDue  bool

	CreatedAt         time.Time
	Due               time.Time
	CompletedAt       time.Time
	LocalCreatedDate  string
	LocalDue          string
	LocalDueShortDate string

	TotalUsersCount     int
	CompletedUsersCount int
	PendingUsersCount   int

	Owner          EmbedUser
	ToUsers        []EmbedUser
	PendingUsers   []EmbedUser
	CompletedUsers []EmbedUser

	ColorCssClass string
	TaskBarHtml   template.HTML
}

// type Wiki struct {
// 	IsLastVersion          bool
// 	LocalUpdatedAt         string
// 	UpdatedAtUnixNano      string
// 	CurrentVersionEditor   EmbedUser
// 	LinkedEntries          []*LinkedEntry
// 	BaseOnEntryId          string
// 	BaseOnEntryTitle       string
// 	BaseOnEntryLink        template.HTMLAttr
// 	Versions               []*WikiVersion
// 	CurrentVersionComments []*Entry
// 	OtherVersionsComments  []*Entry
// 	FirstPicture           *Attachment
// }

type EntryVersion struct {
	Id                   string
	GroupId              string
	UpdatedAt            time.Time
	LocalUpdatedAt       string
	UpdatedAtUnixNano    string
	CurrentVersionEditor EmbedUser
	IsNewVersion         bool
}

type ShareGroupRequest struct {
	IsCurrentOrgSent     bool
	IsCurrentOrgAccepted bool
	IsCurrentOrgRejected bool
}

type LinkedEntry struct {
	Id             string
	EType          string
	Title          string
	GroupId        string
	AuthorId       string
	IsRoot         bool
	RootId         string
	RootEntryTitle string
	Link           template.HTMLAttr
}

type Request struct {
	Info           template.HTML
	ActionButton   template.HTML
	FromOrg        EmbedOrg
	ToOrgs         []EmbedOrg
	AcceptedOrgs   []EmbedOrg
	RejectedOrgs   []EmbedOrg
	PendingOrgs    []EmbedOrg
	AcceptedComma  string
	RejectedComma  string
	SharedGroup    *Group
	SharedOrgIdHex string
}

type Conversation struct {
	Id                  string
	Title               string
	UserIds             []string
	Participants        []EmbedUser
	CreatedAt           time.Time
	EndedAt             time.Time
	LocalHumanCreatedAt string
	Topic               string
	Private             bool
	IsClose             bool
	IsShared            bool
	SharedMessageIds    []string
	MessagesCount       int
	Messages            []*Message
}

type Message struct {
	Id                 string
	ConversationId     string
	UserId             string
	Content            string
	HtmlContent        template.HTML
	CreatedAt          time.Time
	EmbedUser          EmbedUser
	ShowUser           bool
	HighlightedContent template.HTML
}

type Entry struct {
	Id         string
	EType      string
	Title      string
	Slug       string
	Content    string
	TypeTitle  string
	RootId     string
	GroupId    string
	AuthorId   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	BumpedUpAt time.Time

	AllAttachmentsURL    string
	Permalink            string
	IconName             string
	LocalHumanCreatedAt  string
	LocalHumanUpdatedAt  string
	WholeLastUpdateAtAgo string
	LastUpdateAtAgo      string
	MentionedUserIds     string
	DomainURL            string
	UpdatedAtUnixNano    string

	HtmlTitle           template.HTML
	HtmlContent         template.HTML
	HtmlContentPart     template.HTML
	TaskHtmlContentPart template.HTML
	WatchlistHtml       template.HTML
	ToUsersHtml         template.HTML
	LikedByUsersHtml    template.HTML
	NotifyOptionsHtml   template.HTML

	Link      template.HTMLAttr
	UploadURL template.HTMLAttr

	IsShared      bool
	IsPublished   bool
	IsCanPublish  bool
	Uncommentable bool

	IsSystemMessage               bool
	IsBroadcast                   bool
	IsBroadcastTypeToAllAdmins    bool
	IsBroadcastTypeToAllUsers     bool
	IsBroadcastTypeToSomeOrgs     bool
	IsFromSuperOrg                bool
	FromOrg                       EmbedOrg
	ToOrgs                        []EmbedOrg
	ToOrgsHtml                    template.HTML
	IsRequest                     bool
	Request                       *Request
	VisibleForSuperUserInSuperOrg bool
	VisibleForSuperOrg            bool

	IsKnowledgeBase    bool
	IsPost             bool
	IsComment          bool
	IsTask             bool
	IsChat             bool
	IsTaskToDo         bool
	IsInWatchList      bool
	IsToGroup          bool
	CurrentUserCanEdit bool
	CanEdit            bool
	ManagerCanEdit     bool
	LikedByMe          bool
	HasInlineTask      bool
	TaskIsCompleted    bool
	IsRoot             bool
	IsUnread           bool
	IsUpdated          bool
	IsLastVersion      bool

	AllAttachmentsCount int
	CommentsCount       int
	AllLikesCount       int

	Author               EmbedUser
	CurrentVersionEditor EmbedUser
	Group                *Group
	Task                 *Task
	// Wiki                 *Wiki
	ShareGroupRequest *ShareGroupRequest
	Conversation      *Conversation

	LinkedEntries []*LinkedEntry
	Versions      []*EntryVersion

	ToUsers                []EmbedUser
	MentionedUsers         []EmbedUser
	LikedByUsers           []EmbedUser
	Attachments            []*Attachment
	FirstPicture           *Attachment
	Comments               []*Entry
	ExternalComments       []*Entry
	CurrentVersionComments []*Entry
	OtherVersionsComments  []*Entry
	NewComment             *Entry
	NewEntry               *Entry
	GroupSlector           *GroupSelector
}

type EmbedEntry struct {
	Id        string
	GroupId   string
	Title     string
	HtmlTitle template.HTML
	EType     string
	Author    EmbedUser
	ToUsers   []EmbedUser
	Link      template.HTMLAttr
}

type WatchList struct {
	Items         []*WatchItem
	WhatWatchList bool
}

type MyTask struct {
	TasksForMe     []*Entry
	MyCreatedTasks []*Entry
	AboutTodos     bool
}

type MyChats struct {
	ChatEntries      []*Entry
	HasMore          bool
	LatestCreateTime int64
	WhatChats        bool
}

type MyNotifications struct {
	NotificationItems []*NotificationItem
	HasMore           bool
	LatestNotifyTime  int64
}

type NotificationItem struct {
	Id            string
	GroupId       string
	ToUser        EmbedUser
	ForEntry      EmbedEntry
	FromUser      EmbedUser
	FromOrg       EmbedOrg
	CausedByEntry EmbedEntry
	NotifiedAt    time.Time
	ReadAt        time.Time
	Readed        bool
	Type          string
	Link          template.HTMLAttr
}

type WatchItem struct {
	AttachCnt  int
	CommentCnt int
	LikeCnt    int

	AttachCntStr  template.HTML
	CommentCntStr template.HTML
	LikeCntStr    template.HTML

	WatchTime time.Time

	WatchEntry *Entry
}

type DraftList struct {
	// UserId     string
	DraftItems []*Entry
}

type MyCount struct {
	UserId                  string
	FollowedUnreadCount     int
	NotificationUnreadCount int
	ActiveTasksCount        int
	GroupCounts             []*GroupCount
}

type GroupCount struct {
	GroupId     string
	UnreadCount int
}

type GroupHeaderItem struct {
	HasToFollow     bool
	IsFollowing     bool
	IsManaging      bool
	HasFileTab      bool
	HasToDoTab      bool
	IsSystemMessage bool
	SelectedGroup   bool
}

type InlineHelp struct {
	WhatFeed        bool
	WhatGroup       bool
	WhatNext        bool
	WhatChats       bool
	WhatWatchList   bool
	CreateGroups    bool
	InviteOthers    bool
	AboutTodos      bool
	GettingOut      bool
	InviteOthersURL string
	WhatNextURL     string
	WhatChatsURL    string
}

type EmailChanger struct {
	Token string
	Email string
}

type Newsletter struct {
	Email string
}
