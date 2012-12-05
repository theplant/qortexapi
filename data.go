package qortexapi

import (
	"html/template"
	"labix.org/v2/mgo/bson"
	"time"
)

type User struct {
	Id            bson.ObjectId
	Email         string
	Name          string
	Title         string
	Avatar        string
	JID           string
	Timezone      string
	IsSuperUser   bool
	IsSharedUser  bool
	OrgId         bson.ObjectId
	OriginalOrgId bson.ObjectId
	Permalink     string
}

type Attachment struct {
	Permalink   string
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
	LocalCompletedDate string
}

type Wiki struct {
	CurrentVersionEditor         User
	EntryLinks                   []string
	LocalCurrentVersionUpdatedAt string
	LocalHistoryUpdatedAt        string
}

type ShareGroupRequest struct {
	IsCurrentOrgSent     bool
	IsCurrentOrgAccepted bool
	IsCurrentOrgRejected bool
}

type Entry struct {
	GroupId                 bson.ObjectId
	GroupName               string
	EType                   string
	HtmlTitle               template.HTML
	HtmlContent             template.HTML
	TypeTitle               string
	IconName                string
	LocalHumanCreatedAt     string
	Permalink               string
	IsBroadcast             bool
	IsSystemMessage         bool
	IsWiki                  bool
	IsTaskTodo              bool
	Author                  User
	ToUsers                 []User
	LikedByUsers            []User
	CurrentUserCanEdit      bool
	Attachments             []*Attachment
	AllAttachmentsCount     int
	AllAttachmentsPermalink string
	Comments                []*Entry
	Task                    *Task
	Wiki                    *Wiki
	CommentsCount           int
	ShareGroupRequest       *ShareGroupRequest
}
