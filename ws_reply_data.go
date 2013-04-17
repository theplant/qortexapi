package qortexapi

import (
	"html/template"
)

type CountNotification struct {
	Method           string
	GroupId          string
	NewEntry         bool
	EntryId          string
	DelType          string
	MyCount          *MyCount
	NewMessageNumber int
}

type DraftReply struct {
	Method   string
	Success  bool
	IsManual bool
	Message  string
}

type TaskState struct {
	Method  string
	TaskBar string
	EntryId string
	MyCount *MyCount
}

type PulseReply struct {
	Method string
}

//**************
// User related
const (
	REPLY_USER_DELETED = "Reply.DeleteUser"
)

type AvailableUser struct {
	Method  string
	TheUser EmbedUser
}

func NewUserDeleted(embedUser EmbedUser) *AvailableUser {
	return &AvailableUser{
		Method:  REPLY_USER_DELETED,
		TheUser: embedUser,
	}
}

//**************
// Group related
const (
	REPLY_GROUP_CREATED       = "Reply.GroupCreated"
	REPLY_GROUP_DELETED       = "Reply.GroupDeleted"
	REPLY_GROUP_STATE_CHANGED = "Reply.GroupStateChanged"
)

type GroupState struct {
	Method               string
	GroupId              string
	IsFollowed           bool
	GroupMenu            string
	GroupAside           template.HTML
	GroupHeaderUnderInfo template.HTML
	IsNormal             bool
}

func NewGroupCreated(isFollowed bool, groupAside string) *GroupState {
	return &GroupState{
		Method:     REPLY_GROUP_CREATED,
		IsFollowed: isFollowed,
		GroupMenu:  groupAside,
	}
}

func NewGroupDeleted(groupId string) *GroupState {
	return &GroupState{
		Method:  REPLY_GROUP_DELETED,
		GroupId: groupId,
	}
}

func NewGroupStateChanged(groupId string, groupAside string, groupHeaderUnderInfo string, isNormal bool) *GroupState {
	return &GroupState{
		Method:               REPLY_GROUP_STATE_CHANGED,
		GroupId:              groupId,
		GroupAside:           template.HTML(groupAside),
		GroupHeaderUnderInfo: template.HTML(groupHeaderUnderInfo),
		IsNormal:             isNormal,
	}
}
