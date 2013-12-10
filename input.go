package qortexapi

type EntryInput struct {
	Id string

	// "post": normal entry
	// "task": acknowledgement or To-Do (IsAcknowledgement == true ,IsToGroup == "2")
	// "comment": entry's comment
	EType string

	Title   string
	Content string
	GroupId string

	IsToGroup        string // “0”:Notify People ,“1”:Notify Group
	ToUserIds        string // notify users  seperate with "," for example: "1234,4567" means []string{"1234", "5678"}
	TodoUserIds      string // Todo users seperate with "," for example: "1234,4567" means []string{"1234", "5678"}
	MentionedUserIds string // @users seperate with "," for example: "1234,4567" means []string{"1234", "5678"}

	IsAcknowledgement bool   // if IsAcknowledgement == true, get acknowledgement from notified people(ToUserIds).
	IsToDo            bool   // IsToDo == true, will create todo for entry.
	TaskDue           string // if AddToDo == true and want to set a deadline. format:20130507
	TodoStatus        int    // set it in group setting
	Priority          int    // Now :0 ,Soon :1, Someday:2
	Label             int    // set it in group setting
	EstimateTime      string // task's time Estimate

	RootId string // required if etype == "comment"

	NewVersion   bool   // if NewVersion == true will create new version.
	OldGroupId   string // when update entry required
	LastUpdateAt string // when update entry required

	KnowledgeBase bool // if KnowledgeBase == true, this entry is KnowledgeBase.
	AnyoneCanEdit bool // if AnyoneCanEdit == true, anyone can edit this entry.
	Presentation  bool // if Presentation == true, this entry is a Presentation

	IsFromEmail bool // IsFromEmail == true, entry create through email.

	IsPublished bool   // if IsPublished == true, this entry is a public blog.
	Slug        string // if IsPublished == true required
	Email       string // Blog Comment required
	Name        string // Blog Comment required

	InlineHelp       bool
	LinkTitle        string //for qortex support knowledge base
	BaseOnEntryId    string // when share chat,BaseOnEntryId = chat entry id
	PublishedToUsers bool

	LocaleName       string
	BingLanguageCode string //Bing translation code

	// For Creating To-Dos From Comment
	BasedPostId        string
	GroupIdOfBasedPost string
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

	LocaleName string
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
	LinkTitle        string //for qortex support knowledge base

	LocaleName string
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
	Name                     string
	Summary                  string
	Address                  string
	Phone                    string
	Website                  string
	Country                  string
	Size                     string
	Domains                  []string
	QortexURL                string
	LogoURL                  string
	SharingToken             string
	ContactWay               string
	NeedDemo                 bool
	RestrictSubscriptionMail bool
	AnyoneCanJoin            bool
	LanguageCodes            []string
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
	UserLocationCityName     string
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
	IsAgreed    bool
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

type TaskInput struct {
	TaskId                    string
	GroupId                   string
	AssigneeId                string
	TodoStatus                int
	Label                     int
	Priority                  int
	EstimateTime              string
	SpentTime                 string
	IsClaiming                bool
	ToUserIds                 string
	TimetrackingHistoryUpdate string
}

type TaskPwMap struct {
	Id             string `bson:"-"`
	PriorityWeight float64
}

type KnowledgeOverviewInput struct {
	GroupId  string
	Title    string
	Content  string
	IsHidden bool
}
