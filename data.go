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
	IsActive                 bool
}

type Blog struct {
	Title       string
	Description string
	SideContent template.HTML
}

type BlogEntry struct {
	Id               string
	Title            string
	HtmlTitle        template.HTML
	Slug             string
	CreatedAt        time.Time
	LocalCreatedAt   string
	Permalink        string
	CreateCommentURL string
	HtmlContent      template.HTML
	HtmlContentPart  template.HTML
	Author           EmbedUser
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
	NoDetail             bool `json:",omitempty"`
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
	EnableHTML5Notification  bool
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
	Avatar         string            `json:",omitempty"`
	JID            string            `json:",omitempty"`
	Timezone       string            `json:",omitempty"`
	IsSuperUser    bool              `json:",omitempty"`
	IsShare        bool              `json:",omitempty"`
	OrganizationId string            `json:",omitempty"`
	OriginalOrgId  string            `json:",omitempty"`
	ProfileURL     template.HTMLAttr `json:",omitempty"`
	NoDetail       bool              `json:",omitempty"`
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
	IsAdmin             bool              `json:",omitempty"`
	IsPrivate           bool              `json:",omitempty"`
	Editable            bool              `json:",omitempty"`
	Managable           bool              `json:",omitempty"`
	FollowedByMe        bool              `json:",omitempty"`
	AdministratedByMe   bool              `json:",omitempty"`
	IsPreShared         bool              `json:",omitempty"`
	IsShared            bool              `json:",omitempty"`
	IsDefaultLogoURL    bool              `json:",omitempty"`
	HostOrgName         string            `json:",omitempty"`
	IsDispayHostOrgName bool              `json:",omitempty"`
	EntriesCount        int               `json:",omitempty"`
	FollowersCount      int               `json:",omitempty"`
	IsAnnoucement       bool              `json:",omitempty"`
	GroupOwners         []EmbedUser       `json:",omitempty"`
	SharingInfo         *GroupSharingInfo `json:",omitempty"`
	GroupEmailAddress   string
}

type EmbedGroup struct {
	Id   string
	Name string
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
	Width         int
	Height        int

	URL          template.HTMLAttr `json:",omitempty"` // origin
	S1ThumbURL   template.HTMLAttr `json:",omitempty"` // width 240
	MThumbURL    template.HTMLAttr `json:",omitempty"` // width 640
	LThumbURL    template.HTMLAttr `json:",omitempty"` // width 720
	ImageIconURL template.HTMLAttr `json:",omitempty"`
	FileIconURL  template.HTMLAttr `json:",omitempty"`
	HumanSize    string            `json:",omitempty"`
	IsImage      bool
	FileKind     string
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

// TODO: should be replaced by ShareRequest
type Request struct {
	Id               string
	CurrentPrefixURL string
	Info             template.HTML
	ActionButton     template.HTML

	FromOrg         EmbedOrg
	ToOrg           EmbedOrg
	SharedGroup     EmbedGroup
	SharedOrgIdHex  string
	FromUserIdHex   string
	SharedInvitee   EmbedUser
	SharedInviter   EmbedUser
	SharedResponsor EmbedUser
	ToEmail         string
	State           string
}

// TODO: Deprecated! Remove me later. ShareRequest is the new one
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

type ShareRequest struct {
	Id              string
	FromUser        EmbedUser
	ToUser          EmbedUser
	Responser       EmbedUser
	FromOrg         EmbedOrg
	ToOrg           EmbedOrg
	JoinedOrgs      []EmbedOrg
	SharedGroup     EmbedGroup
	Token           string
	ToEmail         string
	PendingDuration string
	IsNewAccount    bool
	IsPending       bool
	IsAccepted      bool
	IsRejected      bool
	IsForwarded     bool
	IsCanceled      bool
	IsStopped       bool
	Info            template.HTML `json:",omitempty"`
	ActionButton    template.HTML `json:",omitempty"`
	RequestBarHtml  template.HTML `json:",omitempty"`
}

type GroupSharingInfo struct {
	IsSharing       bool
	FromOrg         EmbedOrg
	AccpetedOrg     []EmbedOrg
	ForwardedOrgs   []EmbedOrg
	PendingToEmails []string
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
	WatchedAtAgo         string `json:",omitempty"`
	NextRemindAtLater    string `json:",omitempty"`
	MentionedUserIds     string `json:",omitempty"`
	DomainURL            string `json:",omitempty"`
	UpdatedAtUnixNano    string `json:",omitempty"`
	// last version's update time
	LastUpdateAt string `json:",omitempty"`

	HtmlTitle           template.HTML `json:",omitempty"`
	HtmlContent         template.HTML `json:",omitempty"`
	HtmlContentPart     template.HTML `json:",omitempty"`
	TaskHtmlContentPart template.HTML `json:",omitempty"`
	WatchlistHtml       template.HTML `json:",omitempty"`
	ToUsersHtml         template.HTML `json:",omitempty"`
	LikedByUsersHtml    template.HTML `json:",omitempty"`

	Link             template.HTMLAttr `json:",omitempty"`
	PresentationLink template.HTMLAttr `json:",omitempty"`
	UploadURL        template.HTMLAttr `json:",omitempty"`

	IsShared         bool `json:",omitempty"`
	IsPublished      bool `json:",omitempty"`
	IsCanPublish     bool `json:",omitempty"`
	IsMuted          bool `json:",omitempty"`
	IsReminding      bool `json:",omitempty"`
	IsSmartReminding bool `json:",omitempty"`
	IsNoReminding    bool `json:",omitempty"`

	InnerMessage               *InnerMessage `json:",omitempty"`
	IsSystemMessage            bool          `json:",omitempty"`
	IsInnerMessage             bool          `json:",omitempty"`
	SystemMessageType          string        `json:",omitempty"`
	BroadcastType              string        `json:",omitempty"`
	IsBroadcast                bool          `json:",omitempty"`
	IsBroadcastTypeToAllAdmins bool          `json:",omitempty"`
	IsBroadcastTypeToAllUsers  bool          `json:",omitempty"`
	IsBroadcastTypeToSomeOrgs  bool          `json:",omitempty"`
	IsFromSuperOrg             bool          `json:",omitempty"`
	IsFeedback                 bool          `json:",omitempty"`
	FromOrg                    EmbedOrg      `json:",omitempty"`
	ToOrgs                     []EmbedOrg    `json:",omitempty"`
	ToOrgsHtml                 template.HTML `json:",omitempty"`

	IsHidePresentationTip bool `json:",omitempty"`

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
	InlineHelp         bool   `json:",omitempty"`

	AllAttachmentsCount         int
	CommentsCount               int
	CurrentVersionCommentsCount int
	AllLikesCount               int
	VersionCount                int
	UnreadCommentCount          int

	Author               EmbedUser
	CurrentVersionEditor EmbedUser
	Group                *Group        `json:",omitempty"`
	Task                 *Task         `json:",omitempty"`
	Conversation         *Conversation `json:",omitempty"`

	LinkedEntries []*LinkedEntry  `json:",omitempty"`
	Versions      []*EntryVersion `json:",omitempty"`

	ToUsers                []EmbedUser   `json:",omitempty"`
	MentionedUsers         []EmbedUser   `json:",omitempty"`
	LikedByUsers           []EmbedUser   `json:",omitempty"`
	Attachments            []*Attachment `json:",omitempty"`
	CommentsAttachments    []*Attachment `json:",omitempty"`
	FirstPicture           *Attachment   `json:",omitempty"`
	Comments               []*Entry
	ExternalComments       []*Entry `json:",omitempty"`
	CurrentVersionComments []*Entry
	OtherVersionsComments  []*Entry
	NewComment             *Entry         `json:",omitempty"`
	NewEntry               *Entry         `json:",omitempty"`
	GroupSlector           *GroupSelector `json:",omitempty"`

	// Aaron New Added
	QortexSupportNotifyOptions   map[string]string `json:",omitempty"`
	IsQortexSupport              bool              `json:",omitempty"`
	QortexSupport                *QortexSupport    `json:",omitempty"`
	IsQortexSupportKnowledgeBase bool              `json:",omitempty"`
	//only in single entry
	IsSuperOrg  bool `json:",omitempty"`
	IsSuperUser bool `json:",omitempty"`

	// Should be removed
	IsRequest                     bool          `json:",omitempty"`
	ShareRequest                  *ShareRequest `json:",omitempty"`
	VisibleForSuperUserInSuperOrg bool          `json:",omitempty"`
	VisibleForSuperOrg            bool          `json:",omitempty"`
}

type QortexSupport struct {
	Audience          string        `json:",omitempty"`
	IsToOffical       bool          `json:",omitempty"`
	IsToAllUsers      bool          `json:",omitempty"`
	IsToAllAdmins     bool          `json:",omitempty"`
	IsToOrganizations bool          `json:",omitempty"`
	FromOrg           EmbedOrg      `json:",omitempty"`
	ToOrgs            []EmbedOrg    `json:",omitempty"`
	ToOrgsHtml        template.HTML `json:",omitempty"`
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

	IsSmartReminding bool
	IsNoReminding    bool

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
	QortexOverview  bool
	WhatNext        bool
	WhatChats       bool
	InviteOthersURL string
	WhatChatsURL    string
}

type EmailChanger struct {
	Token        string
	Email        string
	SharingToken string
}

type Newsletter struct {
	Email string
}

type Invitation struct {
	Email             string
	Token             string
	SentAgo           string
	ByUser            EmbedUser
	HideInPendingList bool
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

type Member struct {
	Name               string
	Email              string
	ComfirmationSentAt string
	SignupConfirmedAt  string
	SignupStatus       string
	JoinedOrgs         []*Organization
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
	CreatedAt        string
	LastUpdate       string
	Author           EmbedUser
}

type AccessReq struct {
	Email        string
	PriorityCode string
	AccessCode   string
	Status       string
	ApprovedBy   string
	CreatedAt    string
	UpdatedAt    string
}

type InnerMessage struct {
	ByUser                string
	GroupName             string
	GroupLink             string
	OrgName               string
	IsDeletedGroupMessage bool
	IsCreatedGroupMessage bool
	IsSetupOrgMessage     bool
}

type GroupAside struct {
	IsMyGroupsCollapse      bool
	IsOtherGroupsCollapse   bool
	ShowNewGroupButton      bool
	HaveOtherGroup          bool
	ShowSharedExternallyBar bool
	AnnounGroup             *Group
	SMGroup                 *Group
	FollowingNormalGroups   []*Group
	FollowingSharedGroups   []*Group
	UnfollowedNormalGroups  []*Group
	UnfollowedSharedGroups  []*Group
}

type OrgUnreadInfo struct {
	OrgId           string
	FeedUnreadCount int
}

const (
	MS_RECEIVE_JOIN_ORG_INVITATION        = "receive join org inviation but not activate"
	MS_RECEIVE_SHARED_GROUP_INVITATION    = "receive shared group inviation but not activate"
	MS_APPROVED                           = "approve access but not setup account"
	MS_SETUP_ACCOUNT                      = "setup account but not create organization"
	MS_ACTIVATE_BY_JOIN_ORG_INVITATION    = "only activate account from join org inviation but not join the org"
	MS_ACTIVATE_BY_SHARE_GROUP_INVITATION = "only activate account from shared group inviation but not create any organization"
	MS_CREATE_ORG                         = "create first organization after setup account"
	MS_JOIN_ORG                           = "join first organization after activate account"
)

type MarketableMemberInfo struct {
	Email     string
	FirstName string
	LastName  string
	Status    string
}
