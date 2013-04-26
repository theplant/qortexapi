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

type Blog struct {
	Title       string
	Description string
	SideContent template.HTML
}

type BlogEntry struct {
	Id               string
	Title            string
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Permalink        string
	CreateCommentURL string
	HtmlContent      template.HTML
	Author           EmbedUser
	Comments         []*BlogEntry
}

type SharingInvitation struct {
	FromOrg         EmbedOrg
	FromUserId      string
	SharedGroup     *Group
	IsNewAccount    bool
	Email           string
	Token           string
	JoinedOrgs      []EmbedOrg
	IsAccepted      bool
	IsRejected      bool
	IsPending       bool
	IsForwarded     bool
	IsCanceled      bool
	IsStopped       bool
	PendingDuration string
	ToOrgName       string
	ToOrgId         string
}

type User struct {
	Id                   string
	Email                string
	Firstame             string
	LastName             string
	Name                 string
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
	Title          string            `json:",omitempty"`
	Avatar16       string            `json:",omitempty"`
	Avatar32       string            `json:",omitempty"`
	JID            string            `json:",omitempty"`
	Timezone       string            `json:",omitempty"`
	IsSuperUser    bool              `json:",omitempty"`
	IsShare        bool              `json:",omitempty"`
	OrganizationId string            `json:",omitempty"`
	OriginalOrgId  string            `json:",omitempty"`
	ProfileURL     template.HTMLAttr `json:",omitempty"`
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
	Id                  string `json:",omitempty"`
	Name                string `json:",omitempty"`
	Description         string `json:",omitempty"`
	GType               string `json:",omitempty"`
	LogoURL             string `json:",omitempty"`
	IconName            string `json:",omitempty"`
	Link                string `json:",omitempty"`
	Slug                string `json:",omitempty"`
	Author              EmbedUser
	IsAdmin             bool        `json:",omitempty"`
	IsPrivate           bool        `json:",omitempty"`
	Editable            bool        `json:",omitempty"`
	Managable           bool        `json:",omitempty"`
	FollowedByMe        bool        `json:",omitempty"`
	AdministratedByMe   bool        `json:",omitempty"`
	IsPreShared         bool        `json:",omitempty"`
	IsShared            bool        `json:",omitempty"`
	IsDefaultLogoURL    bool        `json:",omitempty"`
	HostOrgName         string      `json:",omitempty"`
	IsDispayHostOrgName bool        `json:",omitempty"`
	EntriesCount        int         `json:",omitempty"`
	FollowersCount      int         `json:",omitempty"`
	IsAnnoucement       bool        `json:",omitempty"`
	GroupOwners         []EmbedUser `json:",omitempty"`
	SharedGroupFromOrg  EmbedOrg
	AcceptedEmbedOrgs   []EmbedOrg `json:",omitempty"`
	PreSharingEmails    []string   `json:",omitempty"`
	ForwardedOrgs       []EmbedOrg `json:",omitempty"`
	HasPendingItems     bool       `json:",omitempty"`
}

type GroupSelectorItem struct {
	Id         string
	Name       string
	IsSelected bool
}

type GroupSelector struct {
	Header                  template.HTML
	SelectedGroupId         string
	SysMessage              *GroupSelectorItem
	FollowingNormalGroups   []*GroupSelectorItem
	FollowingSharedGroups   []*GroupSelectorItem
	UnFollowingNormalGroups []*GroupSelectorItem
	UnFollowingSharedGroups []*GroupSelectorItem
}

type Attachment struct {
	Id            string
	OwnerId       []string  `json:",omitempty"`
	Category      string    `json:",omitempty"`
	Filename      string    `json:",omitempty"`
	ShortFilename string    `json:",omitempty"`
	ContentType   string    `json:",omitempty"`
	ContentId     string    `json:",omitempty"`
	MD5           string    `json:",omitempty"`
	ContentLength int64     `json:",omitempty"`
	Error         string    `json:",omitempty"`
	GroupId       []string  `json:",omitempty"`
	UploadTime    time.Time `json:",omitempty"`
	Width         int       `json:",omitempty"`
	Height        int       `json:",omitempty"`

	URL          template.HTMLAttr `json:",omitempty"`
	ImageIconURL template.HTMLAttr `json:",omitempty"`
	FileIconURL  template.HTMLAttr `json:",omitempty"`
	HumanSize    string            `json:",omitempty"`
	IsImage      bool              `json:",omitempty"`
	FileKind     string            `json:",omitempty"`
}

type Task struct {
	IsTaskOwner       bool `json:",omitempty"`
	IsTaskAssignee    bool `json:",omitempty"`
	IsOthers          bool `json:",omitempty"`
	IsCurrentUserDone bool `json:",omitempty"`

	IsAcknowledgement bool `json:",omitempty"`
	IsTodoForOne      bool `json:",omitempty"`
	IsTodoForAll      bool `json:",omitempty"`

	IsCompleted bool `json:",omitempty"`
	IsClosed    bool `json:",omitempty"`

	IsDueToday bool `json:",omitempty"`
	IsOverDue  bool `json:",omitempty"`

	CreatedAt         time.Time `json:",omitempty"`
	Due               time.Time `json:",omitempty"`
	CompletedAt       time.Time `json:",omitempty"`
	LocalCreatedDate  string    `json:",omitempty"`
	LocalDue          string    `json:",omitempty"`
	LocalDueShortDate string    `json:",omitempty"`
	DueInputValue     string    `json:",omitempty"`

	TotalUsersCount     int `json:",omitempty"`
	CompletedUsersCount int `json:",omitempty"`
	PendingUsersCount   int `json:",omitempty"`

	Owner          EmbedUser   `json:",omitempty"`
	ToUsers        []EmbedUser `json:"-"`
	PendingUsers   []EmbedUser `json:",omitempty"`
	CompletedUsers []EmbedUser `json:",omitempty"`

	ColorCssClass string        `json:",omitempty"`
	TaskBarHtml   template.HTML `json:",omitempty"`
}

type EntryVersion struct {
	Id                   string
	GroupId              string
	UpdatedAt            time.Time
	LocalUpdatedAt       string
	UpdatedAtUnixNano    string
	CurrentVersionEditor EmbedUser
	IsNewVersion         bool
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
	CurrentPrefixURL string
	Info             template.HTML
	ActionButton     template.HTML
	FromOrg          EmbedOrg
	ToOrg            EmbedOrg
	SharedGroup      *Group
	SharedOrgIdHex   string
	FromUserIdHex    string
	SharedInvitee    EmbedUser
	SharedInviter    EmbedUser
	SharedResponsor  EmbedUser
	ToEmail          string
	State            string
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
	EType      string    `json:",omitempty"`
	Title      string    `json:"-"`
	Slug       string    `json:",omitempty"`
	Content    string    `json:"-"`
	TypeTitle  string    `json:",omitempty"`
	RootId     string    `json:",omitempty"`
	GroupId    string    `json:",omitempty"`
	AuthorId   string    `json:",omitempty"`
	CreatedAt  time.Time `json:",omitempty"`
	UpdatedAt  time.Time `json:",omitempty"`
	BumpedUpAt time.Time `json:",omitempty"`

	AllAttachmentsURL    string `json:",omitempty"`
	Permalink            string `json:",omitempty"`
	IconName             string `json:",omitempty"`
	LocalHumanCreatedAt  string `json:",omitempty"`
	LocalHumanUpdatedAt  string `json:",omitempty"`
	WholeLastUpdateAtAgo string `json:",omitempty"`
	LastUpdateAtAgo      string `json:",omitempty"`
	MentionedUserIds     string `json:",omitempty"`
	DomainURL            string `json:",omitempty"`
	UpdatedAtUnixNano    string `json:",omitempty"`

	HtmlTitle           template.HTML `json:",omitempty"`
	HtmlContent         template.HTML `json:",omitempty"`
	HtmlContentPart     template.HTML `json:",omitempty"`
	TaskHtmlContentPart template.HTML `json:",omitempty"`
	WatchlistHtml       template.HTML `json:",omitempty"`
	ToUsersHtml         template.HTML `json:",omitempty"`
	LikedByUsersHtml    template.HTML `json:",omitempty"`
	NotifyOptionsHtml   template.HTML `json:",omitempty"`

	Link             template.HTMLAttr `json:",omitempty"`
	PresentationLink template.HTMLAttr `json:",omitempty"`
	UploadURL        template.HTMLAttr `json:",omitempty"`

	IsShared     bool `json:",omitempty"`
	IsPublished  bool `json:",omitempty"`
	IsCanPublish bool `json:",omitempty"`
	IsMuted      bool `json:",omitempty"`

	IsSystemMessage               bool          `json:",omitempty"`
	SystemMessageType             string        `json:",omitempty"`
	BroadcastType                 string        `json:",omitempty"`
	IsBroadcast                   bool          `json:",omitempty"`
	IsBroadcastTypeToAllAdmins    bool          `json:",omitempty"`
	IsBroadcastTypeToAllUsers     bool          `json:",omitempty"`
	IsBroadcastTypeToSomeOrgs     bool          `json:",omitempty"`
	IsFromSuperOrg                bool          `json:",omitempty"`
	IsFeedback                    bool          `json:",omitempty"`
	FromOrg                       EmbedOrg      `json:",omitempty"`
	ToOrgs                        []EmbedOrg    `json:",omitempty"`
	ToOrgsHtml                    template.HTML `json:",omitempty"`
	IsRequest                     bool          `json:",omitempty"`
	Request                       *Request      `json:",omitempty"`
	VisibleForSuperUserInSuperOrg bool          `json:",omitempty"`
	VisibleForSuperOrg            bool          `json:",omitempty"`

	IsKnowledgeBase    bool   `json:",omitempty"`
	IsPost             bool   `json:",omitempty"`
	IsComment          bool   `json:",omitempty"`
	IsTask             bool   `json:",omitempty"`
	IsChat             bool   `json:",omitempty"`
	IsTaskToDo         bool   `json:",omitempty"`
	IsTaskAck          bool   `json:",omitempty"`
	IsInWatchList      bool   `json:",omitempty"`
	IsToGroup          string `json:",omitempty"`
	CurrentUserCanEdit bool   `json:",omitempty"`
	CanEdit            bool   `json:",omitempty"`
	CanReply           bool   `json:",omitempty"`
	ManagerCanEdit     bool   `json:",omitempty"`
	LikedByMe          bool   `json:",omitempty"`
	HasInlineTask      bool   `json:",omitempty"`
	TaskIsCompleted    bool   `json:",omitempty"`
	IsRoot             bool   `json:",omitempty"`
	IsUnread           bool   `json:",omitempty"`
	IsUpdated          bool   `json:",omitempty"`
	IsLastVersion      bool   `json:",omitempty"`
	Presentation       bool   `json:",omitempty"`
	AnyoneCanEdit      bool   `json:",omitempty"`
	IsInGroup          bool   `json:",omitempty"`
	IsFromEmail        bool   `json:",omitempty"`

	AllAttachmentsCount int `json:",omitempty"`
	CommentsCount       int `json:",omitempty"`
	AllLikesCount       int `json:",omitempty"`
	VersionCount        int `json:",omitempty"`

	Author               EmbedUser
	CurrentVersionEditor EmbedUser
	Group                *Group        `json:",omitempty"`
	Task                 *Task         `json:",omitempty"`
	Conversation         *Conversation `json:",omitempty"`

	LinkedEntries []*LinkedEntry  `json:",omitempty"`
	Versions      []*EntryVersion `json:",omitempty"`

	ToUsers                []EmbedUser    `json:",omitempty"`
	MentionedUsers         []EmbedUser    `json:",omitempty"`
	LikedByUsers           []EmbedUser    `json:",omitempty"`
	Attachments            []*Attachment  `json:",omitempty"`
	FirstPicture           *Attachment    `json:",omitempty"`
	Comments               []*Entry       `json:",omitempty"`
	ExternalComments       []*Entry       `json:",omitempty"`
	CurrentVersionComments []*Entry       `json:",omitempty"`
	OtherVersionsComments  []*Entry       `json:",omitempty"`
	NewComment             *Entry         `json:",omitempty"`
	NewEntry               *Entry         `json:",omitempty"`
	GroupSlector           *GroupSelector `json:",omitempty"`
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
	PrefixURL        string
}

type MyNotifications struct {
	NotificationItems []*NotificationItem
	HasMore           bool
	LatestNotifyTime  int64
}

type NotificationItem struct {
	Id                    string
	GroupId               string
	ToUser                EmbedUser
	ForEntry              EmbedEntry
	FromUser              EmbedUser
	FromOrg               EmbedOrg
	CausedByEntry         EmbedEntry
	NotifiedAt            time.Time
	ReadAt                time.Time
	Readed                bool
	Type                  string
	Link                  template.HTMLAttr
	SharingRequestToEmail string
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

type GroupHeader struct {
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

type Invitation struct {
	Email   string
	Token   string
	SentAgo string
	ByUser  EmbedUser
}

type AbandonInfo struct {
	AbandonFromOrg EmbedOrg
	AvailableOrgs  []EmbedOrg
}

type ContactInfo struct {
	FirstName   string
	LastName    string
	CompanyName string
	CompanySize string
	Email       string
	Phone       string
	Country     string
	City        string
	HelpContent string
}

// Following are for Admin Service
type TotalStats struct {
	OrgCount     int
	MemberCount  int
	GroupCount   int
	EntryCount   int
	CommentCount int
	ChatCount    int
}

type OrgStats struct {
	Organization     *Organization
	UserCount        int
	GroupCount       int
	SharedGroupCount int
	EntryCount       int
	CommentCount     int
	ChatCount        int
}

type AccessReq struct {
	Email      string
	AccessCode string
	Status     string
	ApprovedBy string
	CreatedAt  string
	UpdatedAt  string
}
