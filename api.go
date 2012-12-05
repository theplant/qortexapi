package qortexapi

type Service interface {
	CreatePost() (err error)
	UpdatePost() (err error)
	CreateWiki() (err error)
	UpdateWiki() (err error)
	CreatePostWithTask() (err error)
	UpdatePostWithTask() (err error)
}
