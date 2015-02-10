package qortexapi

import (
	"html/template"
	"time"

	"labix.org/v2/mgo/bson"

	paymentapi "github.com/theplant/theplant_payment/api"
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
	Country                  string
	Phone                    string
	Website                  string
	Size                     string
	SizeText                 string
	SharingToken             string
	Domains                  []string
	RestrictSubscriptionMail bool
	IsActive                 bool
	AnyoneCanJoin            bool
	NeedDemo                 bool
	ContactWay               string
	EnableMultilingual       bool
	LanguageSelectors        *LanguageSelectors
	SizeOptions              map[string]string

	IsSample      bool
	IsSandBox     bool
	PublicDemoURL string
	TutorialsURL  string
	CreatedAt     string

	// for current loggind user, added for AuthUserService.GetInitInfo
	UnreadCount             int `json:",omitempty"`
	NotificationUnreadCount int `json:",omitempty"`
	OfflineMessageCount     int `json:",omitempty"`
	ActiveTasksCount        int `json:",omitempty"`
	ActionNeededTasksCount  int `json:",omitempty"`
}

type SearchOrganization struct {
	Id        string
	Name      string
	QortexURL string
	LogoURL   string
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
	LocalCreatedAt   string
	Permalink        string
	CreateCommentURL string
	HtmlContent      template.HTML
	HtmlContentPart  template.HTML
	Author           EmbedUser
	PrevBlogUrl      string
	NextBlogUrl      string
	TweetUrl         string
}

type User struct {
	Id                    string
	Email                 string
	Firstame              string
	LastName              string
	Name                  string
	LocaleName            map[string]string `json:"-"` // Keeping names for different locale, which has different name order for first and last name
	KatakanaName          string
	Avatar                string
	JID                   string
	Timezone              string
	IsSuperUser           bool
	IsSuperUserInSuperOrg bool
	IsSharedUser          bool
	IsOfficial            bool
	OrgId                 string `json:",omitempty"`
	OriginalOrgId         string `json:",omitempty"`
	OrgName               string `json:",omitempty"`
	PrefixURL             string `json:",omitempty"`
	ProfileURL            template.HTMLAttr
	IsLoggedInUser        bool
	IsAvailable           bool
	IsDisabled            bool
	IsDeleted             bool
	Followable            bool
	Editable              bool
	FollowingTheGroup     bool
	FollowingGroups       []*Group     `json:",omitempty"`
	Preferences           *Preferences `json:",omitempty"`
	Profile               Profile      `json:",omitempty"`
	NoDetail              bool         `json:",omitempty"`
	EnabledMobilePush     bool         `json:",omitempty"`
	PreferredLanguageCode string
}

type GroupUsers struct {
	GroupId    string
	IsShared   bool
	IsPrivate  bool
	IsFollowed bool
	EmbedUsers []EmbedUser
}

type Preferences struct {
	Timezone                   string             `json:",omitempty"`
	FirstDayOfWeek             string             `json:",omitempty"`
	TimezoneOffset             string             `json:",omitempty"`
	PreferFullName             bool               `json:",omitempty"`
	EnterForNewLine            bool               `json:",omitempty"`
	AsideGroupsCollapse        bool               `json:",omitempty"`
	AsideOtherGroupsCollapse   bool               `json:",omitempty"`
	ShowMarkUnreadThreshold    int                `json:",omitempty"`
	AdminModeOn                bool               `json:",omitempty"`
	PreferMarkdown             bool               `json:",omitempty"`
	AutoFollowPublicGroup      bool               `json:",omitempty"`
	EnableHTML5Notification    bool               `json:",omitempty"`
	PreferredLanguageSelectors *LanguageSelectors `json:",omitempty"`
}

type Profile struct {
	Summary    string   `json:",omitempty"`
	Title      string   `json:",omitempty"`
	Department string   `json:",omitempty"`
	Location   string   `json:",omitempty"`
	Expertise  string   `json:",omitempty"`
	Interests  string   `json:",omitempty"`
	BirthMonth string   `json:",omitempty"`
	BirthDay   string   `json:",omitempty"`
	WorkPhone  string   `json:",omitempty"`
	Mobile     string   `json:",omitempty"`
	Twitter    string   `json:",omitempty"`
	Skype      string   `json:",omitempty"`
	Facebook   string   `json:",omitempty"`
	Others     []string `json:",omitempty"`
}

type EmbedOrg struct {
	Id            string
	Name          string
	LogoURL       string
	NoNeedToShare bool
}

type EmbedUser struct {
	Id                    string
	Email                 string
	Name                  string
	LocaleName            map[string]string `json:"-"`
	KatakanaName          string
	Title                 string
	Avatar                string
	JID                   string
	Timezone              string
	OrgName               string
	IsSuperUser           bool
	IsDisabled            bool `json:"-"`
	IsDeleted             bool `json:"-"`
	IsAvailable           bool `json:"-"`
	IsShare               bool
	OrganizationId        string
	OriginalOrgId         string
	ProfileURL            template.HTMLAttr `json:"-"`
	NoDetail              bool
	UnfollowSharedGroup   bool
	PreferredLanguageCode string
}

type PanelStatus struct {
	// HasToDo                 bool
	HasDraft                bool
	HasWatchList            bool
	HasChat                 bool
	ShowMarkUnreadThreshold int
}

type Group struct {
	Id                     string `json:",omitempty"`
	Name                   string `json:",omitempty"`
	Description            string `json:",omitempty"`
	GType                  string `json:",omitempty"`
	LogoURL                string `json:",omitempty"`
	IconName               string `json:",omitempty"`
	Link                   string `json:",omitempty"`
	TaskLink               string `json:",omitempty"`
	Slug                   string `json:",omitempty"`
	Author                 EmbedUser
	IsArchived             bool
	IsAdmin                bool              `json:",omitempty"`
	IsPrivate              bool              `json:",omitempty"`
	Editable               bool              `json:",omitempty"`
	Managable              bool              `json:",omitempty"`
	CanShareGroup          bool              `json:",omitempty"`
	Accessible             bool              `json:",omitempty"`
	FollowedByMe           bool              `json:",omitempty"`
	AdministratedByMe      bool              `json:",omitempty"`
	IsPreShared            bool              `json:",omitempty"`
	IsShared               bool              `json:",omitempty"`
	IsDefaultLogoURL       bool              `json:",omitempty"`
	HostOrgName            string            `json:",omitempty"`
	EntriesCount           int               `json:",omitempty"`
	FollowersCount         int               `json:",omitempty"`
	IsAnnouncement         bool              `json:",omitempty"`
	IsQortexSupport        bool              `json:",omitempty"`
	GroupOwners            []EmbedUser       `json:",omitempty"`
	SharingInfo            *GroupSharingInfo `json:",omitempty"`
	GroupEmailAddress      string
	ToDoSettings           *AdvancedToDoSettings
	TodoGroupingRoute      string `json:",omitempty"`
	Collection             *GroupCollection
	IsCollectionFirstGroup bool

	// For Group Selector
	FormattedName string `json:",omitempty"`
	IsSelected    bool

	UnreadCount          int `json:",omitempty"` // for current loggind user
	IsSandboxGroup       bool
	IsInSandboxOrg       bool
	IsSharedInSandboxOrg bool

	// is the current logged-in user the project manager of this group
	AmIPM bool `json:",omitempty"`

	SharedWithOrgs []EmbedOrg `json:",omitempty"`
}

type GroupCollection struct {
	Id   string
	Name string
}

type AdvancedToDoSettings struct {
	Enabled                 bool
	EnableTimeEstimate      bool
	EnableTimeTracking      bool
	TimeUnit                int
	ProjectManager          EmbedUser
	Labels                  []*TagIndex
	NotYetOpenStatuses      []*TagIndex
	OpenStatuses            []*TagIndex
	ClosedStatuses          []*TagIndex
	DefaultNotYetOpenStatus *TagIndex
	DefaultOpenStatus       *TagIndex
	DefaultClosedStatus     *TagIndex
	LabelCounter            int
	NotYetOpenStatusCounter int
	OpenStatusCounter       int
	ClosedStatusCounter     int
	EnabledSharingOrgs      []string

	ShowedThrowawayStatusSuggestion bool
}

type TagIndex struct {
	Tag   string
	Index int
}

type GroupGeneralSettingPage struct {
	Group       *Group
	CurrentOrg  *Organization
	Collections []*GroupCollection `json:",omitempty"`

	IsNewGroup bool   `json:"-"`
	ActionURL  string `json:"-"`
}

type GroupUsersPage struct {
	Group      *Group
	CurrentOrg *Organization

	IsNewGroup bool `json:"-"`
}

type GroupSharingExternallyPage struct {
	Group         *Group
	CurrentOrg    *Organization
	ShareRequests []*ShareRequest
	SharingOrgs   []*EmbedOrg

	IsNewGroup        bool `json:"-"`
	DisableProFeatrue bool
}

type GroupAdvancedSettingPage struct {
	Group       *Group
	Followers   []EmbedUser
	CurrentOrg  *Organization
	SharingInfo *GroupSharingInfo

	IsNewGroup        bool `json:"-"`
	DisableProFeatrue bool

	// Shit...
	ThrowawayStatusSuggestions    map[string]string
	ThrowawayNotYetOpenTagIndexes map[string][]string
	ThrowawayOpenTagIndexes       map[string][]string
	ThrowawayClosedTagIndexes     map[string][]string
}

type EmbedGroup struct {
	Id       string
	Name     string
	IconName string `json:",omitempty"`
	Link     string `json:",omitempty"`
	TaskLink string `json:",omitempty"`
}

type Attachment struct {
	Id                  string
	OwnerId             []string  `json:",omitempty"`
	Category            string    `json:",omitempty"`
	Filename            string    `json:",omitempty"`
	ShortFilename       string    `json:",omitempty"`
	ContentType         string    `json:",omitempty"`
	ContentId           string    `json:",omitempty"`
	MD5                 string    `json:",omitempty"`
	ContentLength       int64     `json:",omitempty"`
	Error               string    `json:",omitempty"`
	GroupId             []string  `json:",omitempty"`
	UploadTime          time.Time `json:",omitempty"`
	LocalHumanCreatedAt string    `json:",omitempty"`
	Width               int
	Height              int
	Resolution          string `json:",omitempty"`

	URL          template.HTMLAttr `json:",omitempty"` // origin
	S1ThumbURL   template.HTMLAttr `json:",omitempty"` // width 240
	MThumbURL    template.HTMLAttr `json:",omitempty"` // width 640
	LThumbURL    template.HTMLAttr `json:",omitempty"` // width 720
	ImageIconURL template.HTMLAttr `json:",omitempty"`
	FileIconURL  template.HTMLAttr `json:",omitempty"`
	DownloadUrl  template.HTMLAttr `json:",omitempty"`
	HumanSize    string            `json:",omitempty"`
	IsImage      bool
	FileKind     string

	// this is 0 if the attachment has no crocodoc version
	// (should not - because it's not an office file)
	// for other statuses see qortex/entry/attachment_crocodoc
	CrocodocStatus int `json:",omitempty"`
	// Crocodoc session Id for iframe viewer
	SessionId string `json:",omitempty"`

	// For Search
	GroupName        string `json:",omitempty"`
	GroupLink        string `json:",omitempty"`
	LinkWithKeywords template.HTMLAttr
	RootId           string `json:",omitempty"`
	EntryTitle       string `json:",omitempty"`
}

type File struct {
	Attachment *Attachment
	Entry      *FileEntry
}

type FileEntry struct {
	Id        string
	GroupId   string
	GroupName string
	Title     string
	EType     string
	Author    EmbedUser
	Link      template.HTML
}

type AdvancedTask struct {
	CurrentAssignee         EmbedUser
	IsTimeEstimationEnabled bool
	IsTimeTrackingEnabled   bool
	IsPendingEstimation     bool

	// PriorityCode int    // The value of priority: 0, 10 or 20.
	// Priority     string // The text of priority: "Please Set", "Soon" or "Someday".
	LabelCode  int
	Label      string
	StatusCode int    // The value of status: 0, 1 or 2.
	Status     string // The text of status: new, open or closed.

	TimeUnit           string
	EstimatedTimeValue float64
	SpentTimeTracking  []*TimeTrackingItem
	TotalSpentTime     string

	AssignableUsers []*AssignableUser

	TaskFlowNewStatuses    []*TaskFlowStatus // All available statuses
	TaskFlowOpenStatuses   []*TaskFlowStatus
	TaskFlowClosedStatuses []*TaskFlowStatus

	TaskLabels []*TaskLabel // All available labels
	// TaskPriorities []*TaskPriority // All available priorities
	TaskLogs []*TaskLog // All the actions have been taken
}

type TaskLog struct {
	Id                string
	IsClaimed         bool // {Author} will do this
	IsAssigneeChanged bool // {Author} reassigned the To-Do from {OldAssignee} to {Assignee}.
	// IsTimingChanged   bool // {Author} set Start Timing to {Priority}.
	IsStatusChanged       bool // {Author} changed the status to {Status}.
	IsEstimationChanged   bool // {Author} set an estimate of {EstimatedTimeValue} {TimeUnit}.
	IsTimeTrackingAdded   bool // {Author} worked {NewSpentTime} {TimeUnit} on this.
	IsTimeTrackingUpdated bool
	IsReopened            bool // {Author} reopened this To-Do.
	IsLabelChanged        bool // {Author} set the label to Bug.
	IsDueChanged          bool // {Author} set the due date from 2013/09/23 to 2014/09/23.
	HtmlContent           template.HTML

	CreatedAt           time.Time
	VersionAt           time.Time
	LocalHumanCreatedAt string
	Author              EmbedUser // User who takes the action
	Assignee            EmbedUser
	OldAssignee         EmbedUser
	EstimatedTimeValue  float64
	TimeUnit            string

	NewSpentTime float64

	Status string
	// Priority string
	Label string

	LocalOldDue string
	LocalNewDue string

	TimeTrackingUpdateLogs []*TimeTrackingUpdateLog
}

type TimeTrackingUpdateLog struct {
	IsUpdateLog  bool
	IsDeleteLog  bool
	IsSelfUpdate bool
	OldSpentTime float64
	NewSpentTime float64
	Owner        EmbedUser
	Updater      EmbedUser
}

type AssignableUser struct {
	UserId     string
	Name       string
	Avatar     string
	IsAssigned bool
}

type TaskSelectorItem struct {
	StoreKey    int
	DisplayText string
	IsCurrent   bool
}

type TaskPriority TaskSelectorItem

type TaskLabel TaskSelectorItem

type TaskFlowStatus TaskSelectorItem

type TimeTrackingItem struct {
	Id         string
	UserName   string
	FinishDate string
	SpentTime  float64
	TimeUnit   string
}

type Task struct {
	Id         string
	GroupId    string
	EntryId    string `bson:",omitempty"`
	EntryTitle string `bson:",omitempty"`
	CommentId  string `bson:",omitempty"` // only for ack in comment

	IsTaskOwner       bool `json:",omitempty"` // Is this task created by current user
	IsTaskAssignee    bool `json:",omitempty"` // Is this task assigned to current user
	IsOthers          bool `json:",omitempty"` //
	IsCurrentUserDone bool `json:",omitempty"` // use in ack to multi-user, if current user click ack

	IsAcknowledgement bool `json:",omitempty"` // is task a ack
	IsTodoForOne      bool `json:",omitempty"` // is task a todo
	IsTodoForAll      bool `json:",omitempty"` // not use in current system

	IsCompleted bool `json:",omitempty"` // task assignee finish or task creator close the task , IsCompleted =true
	IsClosed    bool `json:",omitempty"` // task creator close the task , IsCompleted =true
	IsToGroup   bool `json:",omitempty"`

	IsDueToday bool `json:",omitempty"`
	IsOverDue  bool `json:",omitempty"`

	CreatedAt         time.Time `json:",omitempty"`
	Due               time.Time `json:",omitempty"`
	CompletedAt       time.Time `json:",omitempty"`
	LocalCreatedDate  string    `json:",omitempty"`
	LocalDue          string    `json:",omitempty"`
	LocalDueWithYear  string    `json:",omitempty"`
	LocalDueShortDate string    `json:",omitempty"`
	DueInputValue     string    `json:",omitempty"`

	TotalUsersCount     int `json:",omitempty"`
	CompletedUsersCount int `json:",omitempty"`
	PendingUsersCount   int `json:",omitempty"`

	Owner          EmbedUser   `json:",omitempty"` // task creator
	ToUsers        []EmbedUser `json:",omitempty"` // used in ack and todo, to track all assignees
	ToUsersText    string      `json:",omitempty"` // format   user1Id:user1Name,user2Id:user2Name
	PendingUsers   []EmbedUser `json:",omitempty"` // used in muti-usr ack, to track incompleted users
	CompletedUsers []EmbedUser `json:",omitempty"` // used in muti-usr ack,to track completed users
	Assignee       EmbedUser   `json:",omitempty"` // used in todo, current assignee

	ColorCssClass       string        `json:",omitempty"`
	ReopenColorCssClass string        `json:",omitempty"`
	TaskBarHtml         template.HTML `json:",omitempty"`

	TaskFlow int //For Advanced todo ,NEW:0 ,OPEN:1,CLOSED:2 ,For BasicToDo, NEW:0, CLOSED:2

	IsClaimed      bool // used in muti-user todo, if someone has clicked "i will do it",IsClaimed = true
	IsAdvancedTask bool // True when the PM feature is enabled and the AdvancedTask will be not nil.
	AdvancedTask   *AdvancedTask

	NeedShowAppliedText bool
	NeedToBeEditMode    bool
	IsEditing           bool
	CanEditDueDate      bool
	FarAwayCorner       bool

	PriorityPos        int
	PriorityWeight     float64             `json:",omitempty"`
	SimpleTaskOutlines []SimpleTaskOutline `json:",omitempty"`
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

type InnerMessage struct {
	IsGroupCreated bool
	IsGroupDeleted bool
	IsOrgCreated   bool
	UserName       string
	GroupName      string
	GroupLink      string
	OrgName        string
}

type GroupSharingInfo struct {
	IsSharing       bool
	FromOrg         EmbedOrg
	AccpetedOrgs    []EmbedOrg
	ForwardedOrgs   []EmbedOrg
	PendingToEmails []string
}

type Entry struct {
	Id            string
	EType         string    `json:",omitempty"`
	Title         string    `json:",omitempty"`
	Slug          string    `json:",omitempty"`
	Content       string    `json:",omitempty"`
	TypeTitle     string    `json:",omitempty"`
	RootId        string    `json:",omitempty"`
	GroupId       string    `json:",omitempty"`
	AuthorId      string    `json:",omitempty"`
	CreatedAt     time.Time `json:",omitempty"`
	UpdatedAt     time.Time `json:",omitempty"`
	BumpedUpAt    time.Time `json:",omitempty"`
	VersionAt     time.Time `json:",omitempty"`
	BaseOnEntryId string    `json:",omitempty"`

	IconName             string    `json:",omitempty"`
	LocalHumanCreatedAt  string    `json:",omitempty"`
	LocalHumanUpdatedAt  string    `json:",omitempty"`
	WholeLastUpdateAtAgo string    `json:",omitempty"`
	WholeLastUpdateAt    time.Time `json:",omitempty"`
	LastUpdateAtAgo      string    `json:",omitempty"`
	WatchedAtAgo         string    `json:",omitempty"`
	NextRemindAtLater    string    `json:",omitempty"`
	// MentionedUserIds     string    `json:",omitempty"`
	DomainURL         string `json:",omitempty"`
	UpdatedAtUnixNano string `json:",omitempty"`
	// last version's update time
	LastUpdateAt string `json:",omitempty"`

	HtmlContent         template.HTML `json:",omitempty"`
	EditingHtmlContent  template.HTML `json:",omitempty"`
	HtmlContentPart     template.HTML `json:",omitempty"`
	TaskHtmlContentPart template.HTML `json:",omitempty"`
	WatchlistHtml       template.HTML `json:",omitempty"`
	ToUsersHtml         template.HTML `json:",omitempty"`
	LikedByUsersHtml    template.HTML `json:",omitempty"`
	HistoryPanelHtml    template.HTML `json:",omitempty"`

	Link             template.HTMLAttr `json:",omitempty"`
	BaseOnLink       template.HTMLAttr `json:",omitempty"`
	BaseOnLinkTitle  string            `json:",omitempty"`
	PresentationLink template.HTMLAttr `json:",omitempty"`
	UploadURL        template.HTMLAttr `json:",omitempty"`

	IsShared bool `json:",omitempty"`
	// blog
	IsPublished bool `json:",omitempty"`
	// qortex support knowledge base
	PublishedToUsers bool `json:",omitempty"`
	IsCanPublish     bool `json:",omitempty"`
	IsPreferMarkdown bool `json:",omitempty"`
	IsMuted          bool `json:",omitempty"`
	IsReminding      bool `json:",omitempty"`
	IsSmartReminding bool `json:",omitempty"`
	IsNoReminding    bool `json:",omitempty"`
	IsNotified       bool `json:",omitempty"`

	IsBroadcast     bool          `json:",omitempty"` // Deprecated, should be removed
	IsFromSuperOrg  bool          `json:",omitempty"`
	IsFromSuperUser bool          `json:",omitempty"`
	IsFeedback      bool          `json:",omitempty"` // Deprecated, should be removed
	FromOrg         EmbedOrg      `json:",omitempty"`
	ToOrgs          []EmbedOrg    `json:",omitempty"`
	ToOrgsHtml      template.HTML `json:",omitempty"`

	IsHidePresentationTip bool `json:",omitempty"`

	IsKnowledgeBase bool `json:",omitempty"`
	IsPost          bool `json:",omitempty"`
	IsComment       bool `json:",omitempty"`
	IsTask          bool `json:",omitempty"` // when entry is ack or todo , IsTask = true
	IsNotification  bool `json:",omitempty"`
	Notification    *Notification
	IsChat          bool   `json:",omitempty"`
	IsTaskToDo      bool   `json:",omitempty"` // when entry is todo , IsTaskToDo = true
	IsTaskAck       bool   `json:",omitempty"` // when entry is ack , IsTaskAck = true
	IsTaskLog       bool   `json:",omitempty"` // use IsTaskLog to distinguish between task log and normal comment
	IsInWatchList   bool   `json:",omitempty"`
	IsToGroup       string `json:",omitempty"`
	CanEdit         bool
	CanReply        bool `json:",omitempty"`
	LikedByMe       bool `json:",omitempty"`
	HasInlineTask   bool `json:",omitempty"` // when comment has a ack , HasInlineTask = true
	TaskIsCompleted bool `json:",omitempty"` // obsolete ?  use Todo.IsCompleted or Ack.IsCompleted
	IsRoot          bool `json:",omitempty"`
	IsUnread        bool `json:",omitempty"`
	HasUnread       bool `json:",omitempty"` // if entry or entry's comments have unread, HasUnread = true
	IsUpdated       bool `json:",omitempty"`
	IsLastVersion   bool `json:",omitempty"`
	Presentation    bool `json:",omitempty"`
	AnyoneCanEdit   bool `json:",omitempty"`
	IsInGroup       bool `json:",omitempty"`
	IsFromEmail     bool `json:",omitempty"`
	InlineHelp      bool `json:",omitempty"`
	IsEvent         bool
	HasResource     bool

	VisibleForSuperUserInSuperOrg bool `json:",omitempty"`
	VisibleForSuperOrg            bool `json:",omitempty"`

	AllAttachmentsCount         int
	CommentsCount               int
	CurrentVersionCommentsCount int
	AllLikesCount               int
	VersionCount                int
	UnreadCommentCount          int
	TranslationsCount           int

	Author               EmbedUser
	CurrentVersionEditor EmbedUser
	LastNewVersionEditor EmbedUser
	Group                *Group `json:",omitempty"`
	// Task                 *Task         `json:",omitempty"`
	Todo         *Task         `json:",omitempty"` // when entry is a todo(IsTaskToDo=true), this exsits
	Ack          *Task         `json:",omitempty"` // when entry is a ack(IsTaskAck=true), this exsits
	Conversation *Conversation `json:",omitempty"`
	Event        *Event        // when entry is a event (IsEvent == true), this exists

	Versions []*EntryVersion `json:",omitempty"`

	ToUsers                        []EmbedUser   `json:",omitempty"`
	MentionedUsers                 []EmbedUser   `json:",omitempty"`
	LikedByUsers                   []EmbedUser   `json:",omitempty"`
	IsLikedByUsersCountMoreThanOne bool          `json:",omitempty"`
	Attachments                    []*Attachment `json:",omitempty"`
	CommentsAttachments            []*Attachment `json:",omitempty"`
	FirstPicture                   *Attachment   `json:",omitempty"`
	ExternalComments               []*Entry      `json:",omitempty"`
	CurrentVersionComments         []*Entry
	OtherVersionsComments          []*Entry
	NewComment                     *Entry `json:",omitempty"`

	GroupsLists []*GroupsList `json:",omitempty"`

	// Qortex Support Type
	IsQortexSupport              bool              `json:",omitempty"`
	QortexSupport                *QortexSupport    `json:",omitempty"`
	IsQortexSupportKnowledgeBase bool              `json:",omitempty"`
	LinkTitle                    string            `json:",omitempty"`
	QortexSupportNotifyOptions   map[string]string `json:",omitempty"`

	// System Message Type
	IsSystemMessage bool `json:",omitempty"`

	// Is Share Request type of System Message
	IsRequest    bool          `json:",omitempty"`
	ShareRequest *ShareRequest `json:",omitempty"`

	//Multi locales related
	CurrentLocaleName      string                   `json:",omitempty"`
	LocaleTitleMap         map[string]string        `json:",omitempty"`
	LocaleContentMap       map[string]template.HTML `json:",omitempty"`
	LocaleHtmlContentMap   map[string]template.HTML `json:",omitempty"`
	LanguageCode           string                   `json:",omitempty"`
	ComparedDefaultTitle   string                   `json:",omitempty"`
	ComparedDefaultContent template.HTML            `json:",omitempty"`

	HasMoreThanOneLanguages bool
	IsAllTranslated         bool
	EntryLanguages          []*EntryLanguage
	ToLanguages             []*SupportedLanguage
	IsEditingTranslation    bool

	// Is Inner Message type of System Message
	IsInnerMessage bool          `json:",omitempty"`
	InnerMessage   *InnerMessage `json:",omitempty"`

	// For Advanced To-Dos
	DerivedToDoEntries []*RelatedEntry // For Comment, All embeded items
	RelatedToDoEntries []*RelatedEntry // For Entry
	BasedOnPost        *BasedOnPost

	DisableProFeatrue bool

	LinkedEntries []*LinkedEntry `json:",omitempty"`
}

type Notification struct {
	Type       string
	EntryId    string
	CommentId  string
	OrgId      string //JoinSharedGroupNotification M_SETUP_ORGANIZATION
	User       EmbedUser
	GroupId    string      // innermessage
	EventUsers *EventUsers `json:",omitempty"`
}
type EventUser struct {
	EmbedUser
	ReplyText string
	ReplyVal  string
}

type EventUsers struct {
	JustNowReplied []*EventUser
	OthersReplied  []*EventUser
	NotYetReplied  []*EventUser
}

type Event struct {
	Id            string
	GroupId       string
	AuthorId      string
	EntryId       string
	EntryTitle    string
	IsInviteGroup bool

	InvitedUIDs  []string
	GoingUIDs    []string
	MaybeUIDs    []string
	NotGoingUIDs []string

	Resource []*Resource

	WholeDay       bool
	TimezoneOffset int64
	StartAt        time.Time
	EndAt          time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Resource struct {
	Id        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EntryLanguage struct {
	Code         string
	LanguageName string
	IsCurrent    bool
}

type RelatedEntry struct {
	Title               template.HTML
	Link                template.HTMLAttr
	LocalHumanCreatedAt string
	Author              EmbedUser
	IsComment           bool
	IsEmbedded          bool
}

type BasedOnPost struct {
	RootTitle template.HTML
	Link      template.HTMLAttr
	IsComment bool
	Content   template.HTML
}

type QortexSupport struct {
	Audience          string        `json:",omitempty"`
	IsToOffical       bool          `json:",omitempty"` // Is customer's feedback
	IsToAllUsers      bool          `json:",omitempty"` // Is Qortex Support Message to all users
	IsToAllAdmins     bool          `json:",omitempty"` // Is Qortex Support Message to all admins
	IsToOrganizations bool          `json:",omitempty"` // Is Qortex Support Message to some organizations
	FromOrg           EmbedOrg      `json:",omitempty"`
	ToOrgs            []EmbedOrg    `json:",omitempty"`
	ToOrgsHtml        template.HTML `json:",omitempty"`
	IsPublished       bool          `json:",omitempty"`
}

type EmbedEntry struct {
	Id        string
	GroupId   string
	GroupName string
	Title     string
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
	PrefixURL       string
	UserName        string
	NeedActionTasks []*TaskOutline
	GroupTasks      []*GroupTasksOutline
	ClosedTasks     []*TaskOutline
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
	WatchTime        time.Time
	IsSmartReminding bool
	IsNoReminding    bool

	WatchEntry *Entry
}

type DraftList struct {
	DraftItems []*Entry
}

type MyCount struct {
	UserId                    string
	FollowedUnreadCount       int
	NotificationUnreadCount   int // Deprecated
	NotificationUnreadCountV2 int
	ActiveTasksCount          int
	ActionNeededTasksCount    int // Deprecated
	OfflineMessageCount       int
	GroupCounts               []*GroupCount
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
	IsQortexSupport bool
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
	Id                 string
	Name               string
	Email              string
	ComfirmationSentAt string
	SignupConfirmedAt  string
	SignupStatus       string
	JoinedOrgs         []*Organization
	IsSandbox          bool
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

type OrgPaymentInfo struct {
	OrgId         string
	OrgName       string
	FreeUserLimit int
	IsFreeOrg     bool
	IsSharingOrg  bool
	HasPaid       bool
	TrialDeadline time.Time
	ExpiredAt     time.Time
}

type MinOrgInfo struct {
	OrgId   string
	OrgName string
}

type OrgUserCache struct {
	UserMap               map[string]*User
	EmbedUsersMap         map[string]*EmbedUser
	UserNameMap           map[string]bson.ObjectId
	GroupToSharingOrgsMap map[string][]string
}

type OrgPaymentHistory struct {
	PaymentDescription string
	PaymentDate        time.Time
	CurrencyCode       string
	PaymentAmount      int64
	Status             string
	Term               int64
	UserCount          int64
	NextPaymentDate    time.Time
}

type AccessReq struct {
	Email      string
	AccessCode string
	Status     string
	ApprovedBy string
	CreatedAt  string
	UpdatedAt  string
}

// Deprecated. Now is NewGroupAside
type GroupAside struct {
	IsMyGroupsCollapse     bool
	IsOtherGroupsCollapse  bool
	ShowNewGroupButton     bool
	HaveOtherGroup         bool
	AnnounGroup            *Group
	SMGroup                *Group
	FollowingNormalGroups  []*Group
	FollowingSharedGroups  []*Group
	UnfollowedNormalGroups []*Group
	UnfollowedSharedGroups []*Group
}

type OrgUnreadInfo struct {
	OrgId                   string
	FeedUnreadCount         int
	NotificationUnreadCount int
}

const (
	MS_RECEIVE_JOIN_ORG_INVITATION        = "InvitedNotSignedUpNudgeToUser"
	MS_RECEIVE_SHARED_GROUP_INVITATION    = "InvitedToSharedGroupNotSignedUpNudge"
	MS_APPROVED                           = "ReceivedBetaInvitationNotSignedUpNudge"
	MS_SETUP_ACCOUNT                      = "JustSignedUpDirectlyNudgeToAdmin"
	MS_ACTIVATE_BY_JOIN_ORG_INVITATION    = "JustSignedUpViaInvitedNudgeToAdmin"
	MS_ACTIVATE_BY_SHARE_GROUP_INVITATION = "JustSignedUpViaSharedGroupNudgeToAdmin"
	MS_CREATE_ORG                         = "JustCreatedOrgNudgeToAdmin"
	MS_JOIN_ORG                           = "JustJoinedOrgNudgeToAdmin"
	MS_CREATE_ORG_VIA_SHARE_GROUP         = "JustCreatedOrgViaSharedGroupNudgeToAdmin"
	MS_INACTIVE_DURING_TRIAL              = "InactiveDuringFreeTrialNudgeToAdmin"
	MS_INACTIVE_AT_TRIAL                  = "InactiveAtTrialEndNudgeToAdmin"
	MS_ACTIVE_AT_TRIAL                    = "ActiveAtTrialEndNudgeToAdmin"
	MS_ACTIVE_NEAR_TRIAL                  = "ActiveNearTrialEndNudgeToAdmin"
)

type MailChimpUserListItem struct {
	Email     string
	FirstName string
	LastName  string

	PreferredLanguageCode string
}

type SimpleTaskOutline struct {
	Id                  string
	EntryTitle          string
	AbovePriorityWeight float64
}

// for My Tasks  and Group tasks
type TaskOutline struct {
	Id                  string
	RootId              string
	CommentId           string // only for ack in comment
	EntryTitle          string
	EntryLink           template.HTMLAttr
	IsComment           bool
	Assignee            EmbedUser
	AuthorName          string
	Group               *EmbedGroup
	Age                 string
	LastUpdate          string
	CreatedAt           time.Time
	Status              string
	StatusCode          int
	Due                 string
	Label               string
	LabelCode           int
	EstimateTime        string
	EstimateTimeFloat64 float64
	SpentTime           string
	PriorityWeight      float64
	Priority            int
	CompleteAtStr       string
	CompleteAtUnixNano  int64
	IsTitleCreatedBy    bool
	ActionNeeded        bool
	IsUnprioritized     bool
	EnableTimeEstimate  bool
	EnableTimeTracking  bool
}

type GroupTasksOutline struct {
	Group                   *EmbedGroup
	AcksAndPendingToDos     []*TaskOutline // Action Needed
	SimpleToDos             []*TaskOutline
	OpenToDos               []*TaskOutline
	OpenEstimateTotal       float64
	OpenSpentTotal          float64
	NotStartedToDos         []*TaskOutline
	NotStartedEstimateTotal float64
	NotStartedSpentTotal    float64
	EstimateUnit            string
	EnableTimeEstimate      bool
	EnableTimeTracking      bool
	AllToDos                []*TaskOutline
	AllEstimateTotal        float64
	AllSpentTotal           float64
}

type GroupTasks struct {
	Group                   *EmbedGroup
	AcksAndPendingToDos     []*Task // Action Needed
	SimpleToDos             []*Task
	OpenToDos               []*Task
	NotStartedToDos         []*Task
	OpenEstimateTotal       float64
	NotStartedEstimateTotal float64
	EstimateUnit            string
}

// Bucket is a historical name after many runs of feature iteration,
// basically, it represents a bunch of to-dos and its relevant information.
type OpenAdvancedToDosBucket struct {
	Title        string
	ToDoSettings *AdvancedToDoSettings
	ToDos        []*TaskOutline

	NoStat                       bool
	EstimateTotal, TrackingTotal float64
	EstimateUnit                 string

	Followers   []EmbedUser
	ToDoMarkers []*ToDoMarker

	HasUnprioritizedToDos bool `json:"-"`

	// TODO: to remove
	LenOfToDos int
	Editable   bool
}

type OpenAdvancedToDosPage struct {
	Assignee     EmbedUser `json:",omitempty"`
	ToDosBuckets []*OpenAdvancedToDosBucket
}

type ToDoMarker struct {
	Id                           string
	Label                        string
	Date                         string
	BeforeIndex                  int
	EstimateTotal, TrackingTotal float64
	EstimateUnit                 string
	PriorityWeight               float64
}

type BasicOpenToDoOutlines struct {
	Assignee     EmbedUser
	TaskOutlines []*TaskOutline
	Editable     bool
}

type ClosedAdvancedToDoOutline struct {
	Status           *TagIndex
	Count            int
	WithLoadMoreLink bool
	Tasks            []*TaskOutline
	Group            *Group
}

type ToDoCSVItem struct {
	Creator         string `Title:"i18n.csv.Creator"`
	Title           string `Title:"i18n.csv.Title"`
	Content         string `Title:"i18n.csv.Description"`
	Status          string `Title:"i18n.csv.Status"`
	EstimateTime    string `Title:"i18n.csv.Estimate Time"`
	TotalSpentTime  string `Title:"i18n.csv.Total Spent Time"`
	TimeUnit        string `Title:"i18n.csv.Time Unit"`
	SpentTimeDetail string `Title:"i18n.csv.Spent Time Detail"`
	CreateTime      string `Title:"i18n.csv.Create Time"`
	UpdateTime      string `Title:"i18n.csv.Update Time"`
	Due             string `Title:"i18n.csv.Due"`
}

type TranslatedThread struct {
	Title         string
	Content       string
	Comments      []*TranslatedComment
	IsCommentOnly bool
	IsWikiSection bool
}

type TranslatedComment struct {
	Id      string
	Content string
}

type LanguageSelectors struct {
	Selectors []*LanguageSelector
	LabelText string
}

type LanguageSelector struct {
	IsFirst            bool
	Index              string
	SupportedLanguages []*SupportedLanguage
	PreferredLanguages []*SupportedLanguage
	RestLanguages      []*SupportedLanguage
	UILanguages        []*SupportedLanguage
}

type SupportedLanguage struct {
	StoreKey    string
	DisplayText string
	IsCurrent   bool
}

type BillingInfo struct {
	IsSharedGroupAccount bool
	IsFreeAccount        bool
	IsProAccount         bool
	FreeTrialLeftDays    int
	ExpiredLeftDays      int
	ActiveUserCount      int
	FreeUserLimit        int
	Country              string
	Phone                string
	Billing              *paymentapi.Billing
	BillingDetails       *paymentapi.BillingDetails
	PastPayments         []paymentapi.Payment
	Package              *paymentapi.Package
	CurrencySymbol       string
	MonthlyPrice         int
	HasWaitingBilling    bool
	HasSubscribedBilling bool
	HasPaidBilling       bool
	IsOverdue            bool
	DismissPaymentTips   bool
	PrefixURL            string
	FreeTrialDays        int

	//Date display related
	PackageExpiredAt                 string
	PackageSubscribedNextPaymentDate string
}

type ReceiptInfo struct {
	Id          string
	OrgName     string
	OrgAdress   string
	OrgCountry  string
	PaymentTerm string
	UserCount   int
	PaymentDate string
	PriceUnit   string
	Cost        string
	Tax         string

	// for jp
	CostWithoutTax string
	YearOrMonth    string
	// for countries ,not jp ,de
	CurrencySymbol string
}

type KnowledgeOverview struct {
	Author                  EmbedUser
	PrefixURL               string
	GroupId                 string
	EntryId                 string
	Content                 string
	Title                   string
	HtmlContent             template.HTML
	LocaleTitleMap          map[string]string        `json:",omitempty"`
	LocaleHtmlContentMap    map[string]template.HTML `json:",omitempty"`
	IsPreferMarkdown        bool
	IsHidden                bool
	IsAtQortexSupport       bool
	Editable                bool
	CanSeeHiddenBanner      bool
	HasVersions             bool
	Versions                []*EntryVersion `json:",omitempty"`
	LanguageCode            string
	EntryLanguages          []*EntryLanguage
	ToLanguages             []*SupportedLanguage
	HasMoreThanOneLanguages bool
	IsAllTranslated         bool
	UploadURL               template.HTMLAttr `json:",omitempty"` //just for reuse the mannual translation form
	IsHidePresentationTip   bool              `json:",omitempty"` //just for reuse the mannual translation form
	Id                      string            `json:",omitempty"` //just for reuse the mannual translation form
	DisableProFeatrue       bool
	IsEditingTranslation    bool
}

type ContactUsInfo struct {
	FirstName   string
	LastName    string
	Email       string
	Address     string
	Phone       string
	Country     string
	CompanyName string
	CompanySize string
}

type InitInfo struct {
	CurrentUser User

	CurrentOrg *Organization
	JoinedOrgs []*Organization

	// Return by GetInitInfo
	AnnouncementGroup      *Group
	SmGroup                *Group
	FollowedNormalGroups   []*Group
	FollowedSharedGroups   []*Group
	UnFollowedNormalGroups []*Group
	UnFollowedSharedGroups []*Group

	// Return by GetNewInitInfo
	Lists []*GroupsList
}

type NewGroupAside struct {
	ShowNewGroupButton bool

	// Lists contains at least one element, and up to three elements.
	// The first one is My Groups, then Other Groups, the last one Archived Groups.
	Lists []*GroupsList
}

type GroupsList struct {
	IsCollapsed   bool
	Announcement  *Group // only exists in My Groups
	QortexSupport *Group // only exists in My Groups
	Metas         []*MetaGroup

	// Deprecated in V2
	// Now standard and shared groups are mixed together.
	SharedMetas []*MetaGroup
}

type MetaGroup struct {
	Collection *GroupedGroups
	Group      *Group
}

type GroupedGroups struct {
	ColId            string
	ColName          string
	FormattedColName string
	IsCollapsed      bool
	Groups           []*Group

	// Deprecated in V2
	// Now standard and shared groups are mixed together.
	SharedGroups []*Group
}

type GroupColCollapseState struct {
	ColId       string
	IsCollapsed bool
}

type (
	SearchResult struct {
		Entities    []SearchEntity
		Attachments []*Attachment // TODO: should be removed
		GroupsLists []*GroupsList

		HighlightedTerms []string

		SearchAllGroupsURL string
		Keywords           string

		Page      int
		TotalPage int
		PrePage   int
		PrePages  []int
		NextPage  int
		NextPages []int

		TotalItemsCount   int
		EntriesCount      int
		FilesCount        int
		FileIndexableIds  []string
		PrivateChatsCount int
	}

	// A SearchEntity could be only be a Chat or a Entry
	SearchEntity struct {
		Chat  *Chat
		Entry *SearchEntry
		File  *Attachment
	}

	SearchEntry struct {
		Id     string
		Title  template.HTML
		Author EmbedUser

		GroupId       string
		GroupName     string
		GroupIconName string

		LocalHumanCreatedAt string
		LinkWithKeywords    template.HTMLAttr
		HighlightedChunks   []template.HTML // array is overdesign, or just unnecessary?

		IsChat          bool
		IsKnowledgeBase bool
		AnyoneCanEdit   bool
		IsPublished     bool
		IsShared        bool
		IsTask          bool
		IsFromEmail     bool
	}

	SearchLink struct {
		IndexableId string

		Title                   string
		Link                    string
		LinkToEntryWithKeywords template.HTMLAttr
		GroupName               string
		GroupLink               string
	}
)

type PushInfo struct {
	OrgId     string
	UserId    string
	GroupId   string
	EntryId   string
	CommentId string
}

type Token struct {
	Id            string
	MemberId      string
	OrgId         string
	IsForAllGroup bool
	GroupIds      []string
	Key           string
	Label         string
	AccessLevel   int
}

// ----- Chat Related -----
type (
	Chat struct {
		Id          string
		UpdatedAt   string
		CurrentUser EmbedUser
		WithUser    EmbedUser
		Convs       []*Conversation `json:",omitempty"`
		HasMore     bool
		Messages    []*Message `json:",omitempty"` // For holding the chat search results
	}

	Conversation struct {
		Id                  string
		Title               string
		UserIds             []string
		Participants        []EmbedUser
		CreatedAtUnixNano   int64
		LocalHumanCreatedAt string
		Topic               string
		Private             bool
		IsClose             bool
		IsShared            bool
		HasOfflineMessage   bool
		OfflineLocalTime    string
		SharedMessageIds    []string
		MessagesCount       int
		Messages            []*Message

		LinkWithKeywords template.HTMLAttr `json:",omitempty"` // for Search
	}

	Message struct {
		Id                  string
		ConversationId      string
		UserId              string
		Content             string
		HtmlContent         template.HTML
		CreatedAt           time.Time
		LocalHumanCreatedAt string
		EmbedUser           EmbedUser
		ShowUser            bool
		IsOffline           bool
	}
)

// Calendar

type (
	CalendarItemGroup struct {
		Name   string
		WeekNo int
		Items  []CalendarItem
	}

	CalendarItem struct {
		EntryId        string
		EventId        string
		TaskId         string
		EntryLink      string
		Title          string
		IsAllDay       bool
		ItemType       string      `json:",omitempty"`
		Attending      string      `json:",omitempty"` // YES, NO, MAYBE
		Owner          EmbedUser   `json:",omitempty"`
		Assignee       EmbedUser   `json:",omitempty"`
		ToUsers        []EmbedUser `json:",omitempty"`
		Group          EmbedGroup  `json:",omitempty"`
		Date           string      `json:",omitempty"` // 2014-01-02
		StartTime      time.Time   `json:",omitempty"`
		EndTime        time.Time   `json:",omitempty"`
		ImportanceType string      `json:",omitempty"`
		SortValue      string      `json:",omitempty"`
	}
)

// The data shared between qortex and qortexpush
type (
	UniqushPushData struct {
		UserId      string `form:"UserId" binding:"required"`
		Message     string `form:"Message" binding:"required"`
		AlertNumber string `form:"AlertNumber" binding:"required"`
		ItemId      string `form:"ItemId" binding:"required"`
		Sound       bool   `form:"Sound"`
	}

	UniqushSubData struct {
		UserId       string `form:"UserId" binding:"required"`
		ServiceType  string `form:"ServiceType" binding:"required"`
		IsSubscribe  bool   `form:"IsSubscribe" binding:"required"`
		TokenOrRegId string `form:"TokenOrRegId" binding:"required"`
	}
)
