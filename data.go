package qortexapi

import (
	"html/template"
	// "text/template"
	"time"
)

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
	FollowingGroups      []*Group
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

type Group struct {
	Id                string
	Name              string
	Description       string
	GType             string
	LogoURL           string
	IconName          string
	Link              string
	Author            EmbedUser
	IsAdmin           bool
	IsPrivate         bool
	Editable          bool
	Managable         bool
	FollowedByMe      bool
	AdministratedByMe bool
	IsShared          bool
	HostOrgName       string
	EntriesCount      int
	FollowersCount    int
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
	ContentType   string
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

	TaskBarHtml template.HTML
}

type Wiki struct {
	IsLastVersion          bool
	LocalUpdatedAt         string
	UpdatedAtUnixNano      string
	CurrentVersionEditor   EmbedUser
	LinkedEntries          []*LinkedEntry
	BaseOnEntryId          string
	BaseOnEntryTitle       string
	BaseOnEntryLink        template.HTMLAttr
	Versions               []*WikiVersion
	CurrentVersionComments []*Entry
	OtherVersionsComments  []*Entry
}

type WikiVersion struct {
	Id                string
	GroupId           string
	UpdatedAt         time.Time
	LocalUpdatedAt    string
	UpdatedAtUnixNano string
	CurVerEditor      EmbedUser
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

type Entry struct {
	Id         string
	EType      string
	Title      string
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

	HtmlTitle         template.HTML
	HtmlContent       template.HTML
	HtmlContentPart   template.HTML
	WatchlistHtml     template.HTML
	ToUsersHtml       template.HTML
	CommentsCountHtml template.HTML
	LikedByUsersHtml  template.HTML

	Link      template.HTMLAttr
	UploadURL template.HTMLAttr

	IsBroadcast        bool
	IsSystemMessage    bool
	IsRequest          bool
	IsWiki             bool
	IsPost             bool
	IsTask             bool
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

	AllAttachmentsCount int
	CommentsCount       int

	Author            EmbedUser
	Group             *Group
	Task              *Task
	Wiki              *Wiki
	ShareGroupRequest *ShareGroupRequest

	ToUsers        []EmbedUser
	MentionedUsers []EmbedUser
	LikedByUsers   []EmbedUser
	Attachments    []*Attachment
	Comments       []*Entry
	NewComment     *Entry
	NewEntry       *Entry
	GroupSlector   *GroupSelector
}

type WatchList struct {
	Items         []*WatchItem
	WhatWatchList bool
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
