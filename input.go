package qortexapi

type InputEntry struct {
	Id               string
	Title            string
	Content          string
	GroupId          string
	RootId           string
	BaseOnEntryId    string
	ToUserIds        []string
	MentionedUserIds []string
}
