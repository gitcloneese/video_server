package defs


//requests

type UserCredential struct{
	UserName string `json:"user_name"`
	Pwd		string	`json:"pwd`
}

//Data Model
type VideoInfo struct{
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}

type Comment struct{
	Id string
	VideoId string
	Author string
	Content string
}