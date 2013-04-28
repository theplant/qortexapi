package qortexapi

type EntryInput struct {
	Id               string
	EType            string
	Title            string
	Slug             string
	Content          string
	GroupId          string
	IsToGroup        string
	ToUserIds        string //seperate with "," for example: "1234,4567" means []string{"1234", "5678"}
	MentionedUserIds string
	IsPublished      string

	//task
	IsAcknowledgement string
	TaskRequireType   string
	TaskDue           string

	//comment
	RootId                   string
	IsCommentAcknowledgement string

	// convert to wiki
	BaseOnEntryId string

	//update
	NewVersion string
	OldGroupId string

	KnowledgeBase bool
	AnyoneCanEdit bool
	Presentation  bool

	//incoming mail
	IsFromEmail bool

	//External Comment
	Email string
	Name  string

	IsInnerMessage bool
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

type OrgSettingsInput struct {
	AllowUsersCreateGroups bool
	AllowUsersInvitePeople bool
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
	AuthorId                 string
	MemberIds                []string
	GroupIds                 []string
	QortexURL                string
	LogoURL                  string
	ChatToken                string
	RegistrationMode         int
}

// Like or Unlike an entry action input
type LikeInput struct {
	EntryId string
	GroupId string
	Like    string // "0" for Unlike, "1" for Like
}

type PreferencesInput struct {
	Timezone                 string
	TimezoneOffset           string
	PreferFullName           string
	EnterForNewLine          string
	AsideGroupsCollapse      string
	AsideOtherGroupsCollapse string
	ShowMarkUnreadThreshold  string
	AdminModeOn              string
	PreferMarkdown           string
	AutoFollowPublicGroup    string
}

type MemberAccountInput struct {
	FirstName string
	LastName  string
	AvatarURL string
}

type NewsletterInput struct {
	Email string
}

type ShareChatInput struct {
	Title         string
	Content       string
	BasedConvId   string
	BaseOnEntryId string
	GroupId       string
}

type ContactInput struct {
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

type UserProfileInput struct {
	Summary       string
	Title         string
	Department    string
	Location      string
	Expertise     string
	Interests     string
	BirthMonth    string
	BirthDay      string
	WorkPhone     string
	Mobile        string
	Twitter       string
	Skype         string
	Facebook      string
	OtherWebsites []string
}

type MailUpdatesInput struct {
	IndividualIsOn    bool
	SendLag           int
	AckRequest        bool
	AckConfirmation   bool
	Todo              bool
	TodoConfirmation  bool
	SystemMessage     bool
	EntryNotification bool
	Like              bool
	SendTimeIsOn      bool
	Mon               bool
	Tue               bool
	Wed               bool
	Thu               bool
	Fri               bool
	Sat               bool
	Sun               bool
	SendHoursIsOn     bool
	StartAt           int
	EndAt             int
	DailyIsOn         bool
}
