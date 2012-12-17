package qortexapi

import (
	"html/template"
	// "text/template"
	"time"
)

type Organization struct {
	Id        string
	Name      string
	QortexURL string
	LogoURL   string
}

type User struct {
	Id            string
	Email         string
	Name          string
	Title         string
	Avatar        string
	JID           string
	Timezone      string
	IsSuperUser   bool
	IsSharedUser  bool
	OrgId         string
	OriginalOrgId string
	URL           string
}

type EmbedUser struct {
	Id             string
	Email          string
	Name           string
	Title          string
	Avatar32       string
	JID            string
	Timezone       string
	IsSuperUser    bool
	IsShare        bool
	OrganizationId string
	OriginalOrgId  string
	UserPageUrl    template.HTMLAttr
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
	IsAcknowledgement  bool
	CurrentUserIsOwner bool
	CurrentUserIsDone  bool
	IsCompleted        bool
	IsClosed           bool
	IsOneCompleter     bool
	IsOnePendingUser   bool
	PendingUsers       []EmbedUser
	CompletedUsers     []EmbedUser
	OnlyPendingUser    EmbedUser
	OnlyCompleter      EmbedUser
	Due                time.Time
	CompletedAt        time.Time
	CreatedAt          time.Time
	LocalDueDate       string
	LocalDueShortDate  string
	LocalCompletedDate string

	ColorCssClass string
}

type Wiki struct {
	CurrentVersionEditor         EmbedUser
	EntryLinks                   []string
	LocalCurrentVersionUpdatedAt string
	LocalHistoryUpdatedAt        string
	// WikiVersion
}

type ShareGroupRequest struct {
	IsCurrentOrgSent     bool
	IsCurrentOrgAccepted bool
	IsCurrentOrgRejected bool
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
	WholeLastUpdateAtAgo string
	LastUpdateAtAgo      string
	MentionedUserIds     string

	HtmlTitle         template.HTML
	HtmlContent       template.HTML
	WatchlistHtml     template.HTML
	ToUsersHtml       template.HTML
	CommentsCountHtml template.HTML
	LikedByUsersHtml  template.HTML

	Link           template.HTMLAttr
	UploadURL      template.HTMLAttr
	CommentFormURL template.HTMLAttr
	UpdatePostURL  template.HTMLAttr

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
