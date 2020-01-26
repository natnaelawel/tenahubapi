package entity

import (
	"time"
)

// User represents user data
var MockUser = User{
	ID      :  1,
	FirstName :  "Nathaniel",
	LastName:"Awel",
	Email:"natnael.awel@gmail.com",
	Password :"secretpassword",
	ProfilePic:"",
	PhoneNumber: "0912443434",
}
var MockAgent = Agent{
	ID      :  1,
	FirstName :  "Nathaniel",
	LastName:"Awel",
	UserName:"natnaelawel",
	Email:"natnael.awel@gmail.com",
	Password :"secretpassword",
	ProfilePic:"",
	PhoneNumber: "0912443434",
}
var MockAdmin = Admin{
	ID      :  1,
	FirstName :  "Nathaniel",
	LastName:"Awel",
	UserName:"natnaelawel",
	Email:"natnael.awel@gmail.com",
	Password :"secretpassword",
	ProfilePic:"",
	PhoneNumber: "0912443434",
}
var MockHealthCenter = HealthCenter{
	ID      :  1,
	Name :  "Nathaniel",
	City:"Awel",
	AgentID:1,
	Email:"natnael.awel@gmail.com",
	Password :"secretpassword",
	ProfilePic:"",
	PhoneNumber: "0912443434",
}

var MockRating = Rating{
	ID:1,
	UserID:3,
	HealthCenterID:2,
	PlacedAt:time.Now(),
}
var MockComment = Comment{
	ID:2,
	Comment:"some comment",
	UserID:3,
	HealthCenterID:2,
	PlacedAt:time.Now(),
}
var MockService = Service{
	ID:2,
	Name:"service name",
	Description:"service description",
	HealthCenterID:3,
	Status:"pending",
}

