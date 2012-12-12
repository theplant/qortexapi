package qortexapi

type InputEntry struct {
	Id               string
	EType            string
	Title            string
	Content          string
	GroupId          string
	IsToGroup        string
	ToUserIds        string
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

type GroupInput struct {
	Id            string
	Name          string
	Description   string
	Type          string
	LogoURL       string
	IconName      string
	IsPrivate     bool
	IsShared      bool
	GroupOwners   []string
	InvitedOrgIds []string
}
