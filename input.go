package qortexapi

type EntryInput struct {
	Id string
	// "post": normal entry
	// "task":  acknowledgement or To-Do (IsAcknowledgement == true ,IsToGroup == "2")
	// "comment":  entry's comment
	EType   string
	Title   string
	Content string
	GroupId string

	IsToGroup        string // “0”:Notify People ,“1”:Notify Group, “2”:To-Do
	ToUserIds        string // notify users  seperate with "," for example: "1234,4567" means []string{"1234", "5678"}
	MentionedUserIds string // @users        seperate with "," for example: "1234,4567" means []string{"1234", "5678"}

	IsAcknowledgement bool   // if IsAcknowledgement == true, get acknowledgement from notified people(ToUserIds).
	TaskDue           string // if IsToGroup == "2" and want to set a deadline. format:20130507

	RootId                   string // if etype == "comment"  required
	IsCommentAcknowledgement string // if etype == "comment"  required

	NewVersion   string // if NewVersion == "1"  will create new version.
	OldGroupId   string // when update entry  required
	LastUpdateAt string

	KnowledgeBase bool // if KnowledgeBase == true, this entry is KnowledgeBase.
	AnyoneCanEdit bool // if AnyoneCanEdit == true, anyone can edit this entry.
	Presentation  bool // if Presentation == true, this entry is a Presentation

	IsFromEmail bool // IsFromEmail == true, entry create through email.

	IsPublished bool   // if IsPublished == true, this entry is a public blog.
	Slug        string // if IsPublished == true required
	Email       string // Blog Comment required
	Name        string // Blog Comment required

	InlineHelp bool
	BaseOnEntryId string // when share chat,BaseOnEntryId = chat entry id
	PublishedToUsers bool
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

// TODO: Explaination needed.
type QortexSupportInput struct {
	Id               string
	Title            string
	Content          string
	ToOrgIds         []string
	RootId           string
	Audiance         string
	KnowledgeBase    bool
	InlineHelp       bool
	PublishedToUsers bool
}

type GroupInput struct {
	Id            string
	Name          string
	Description   string
	LogoURL       string
	IconName      string
	Slug          string
	IsPrivate     bool
	IsShared      bool
	GroupOwners   []string
	InvitedOrgIds []string
	AutoGenSlug   bool
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
	EnableHTML5Notification  string
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
	Fake        bool // always false
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

// TODO: mail-updates: remove it
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

type MailPreferenceInput struct {
	Expecting    bool
	SendInterval int
	SendLag      int
}
