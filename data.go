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

type Group struct {
	Id                string
	Name              string
	Description       string
	GType             string
	LogoURL           string
	IconName          string
	Link              string
	Author            User
	IsAdmin           bool
	IsPrivate         bool
	Editable          bool
	Managable          bool
	FollowedByMe      bool
	AdministratedByMe bool
	IsShared          bool
}

type Attachment struct {
	URL         string
	Filename    string
	FileIconURL string
	HumanSize   string
	Size        int
	IsImage     bool
}

type Task struct {
	IsAcknowledgement  bool
	CurrentUserIsOwner bool
	CurrentUserIsDone  bool
	IsCompleted        bool
	IsOneCompleter     bool
	IsOnePendingUser   bool
	PendingUsers       []User
	CompletedUsers     []User
	OnlyPendingUser    User
	OnlyCompleter      User
	Due                time.Time `json:",omitempty"`
	CompletedAt        time.Time `json:",omitempty"`
	CreatedAt          time.Time
	LocalDueDate       string
	LocalDueShortDate  string
	LocalCompletedDate string
}

type Wiki struct {
	CurrentVersionEditor         User
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
	TypeTitle  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	BumpedUpAt time.Time

	AllAttachmentsURL    string
	Permalink            string
	IconName             string
	LocalHumanCreatedAt  string
	WholeLastUpdateAtAgo string
	LastUpdateAtAgo      string

	HtmlTitle     template.HTML
	HtmlContent   template.HTML
	Link          template.HTMLAttr
	WatchlistHtml template.HTML

	IsBroadcast        bool
	IsSystemMessage    bool
	IsWiki             bool
	IsTaskTodo         bool
	IsInWatchList      bool
	CurrentUserCanEdit bool

	AllAttachmentsCount int
	CommentsCount       int

	Author            *User
	Group             *Group
	Task              *Task
	Wiki              *Wiki
	ShareGroupRequest *ShareGroupRequest

	ToUsers      []*User
	LikedByUsers []*User
	Attachments  []*Attachment
	Comments     []*Entry
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
