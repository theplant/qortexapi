package qortexapi

import (
	"time"
)

type EntryInput struct {
	Id               string
	EType            string
	Title            string
	Content          string
	GroupId          string
	IsToGroup        string
	ToUserIds        string //seperate with "," for example: "1234,4567" means []string{"1234", "5678"}
	MentionedUserIds string

	//task
	IsAcknowledgement string
	TaskRequireType   string
	TaskDue           string

	//comment
	RootId                   string
	IsCommentAcknowledgement string

	// convert to wiki
	BaseOnEntryId string
}

const (
	BT_TO_ALL_ADMINS         = "boradcast_type_to_all_admins"
	BT_TO_ALL_USERS          = "boradcast_type_to_all_users"
	BT_TO_SOME_ORGANIZATIONS = "boradcast_type_to_some_organizations"
)

type BroadcastInput struct {
	Id            string
	Title         string
	Content       string
	ToOrgIds      []string
	BroadcastType string
	RootId        string
}

//type RequestInput struct {
//Title          string
//Content        string
//SharedGroupId  string
//SharedOrgId    string
//HostOrgId      string
//ToOrgIds       []string
//AcceptedOrgIds []string
//RejectedOrgIds []string
//}

type GroupInput struct {
	Id            string
	Name          string
	Description   string
	Type          string
	LogoURL       string
	IconName      string
	Slug          string
	IsPrivate     bool
	IsShared      bool
	GroupOwners   []string
	InvitedOrgIds []string
	ActionOrgId   string
}

type OrganizationInput struct {
	Id                       string
	OType                    string
	Name                     string
	Summary                  string
	Address                  string
	Phone                    string
	Website                  string
	Domain                   string
	Domains                  []string
	RestrictSubscriptionMail bool
	CreatedAt                time.Time
	AuthorId                 string
	MemberIds                []string
	GroupIds                 []string
	QortexURL                string
	LogoURL                  string
	ChatToken                string
	RegistrationMode         int
}

type PreferencesInput struct {
	Timezone                 string
	PreferFullName           bool
	EnterForNewLine          bool
	AsideGroupsCollapse      bool
	AsideOtherGroupsCollapse bool
	ShowMarkUnreadThreshold  int
	AdminModeOn              bool
	PreferMarkdown           bool
	AutoFollowPublicGroup    bool
}
