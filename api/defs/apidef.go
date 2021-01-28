package defs


//requests

type UserCredential struct{
	UserName string `json:"user_name"`
	Pwd		string	`json:"pwd`
}

//response
type SignedUp struct{
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
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
	Time string
}

type SimpleSession struct{
	Usemename string   //login name
	TTL int64   //time to live
}