package domain

import "time"

type User struct {
	Id         int
	Auth       string
	Email      string
	Status     string
	Mobile     string
	Name       string
	Gender     string
	Age        int
	Profession string
	CreatedOn  time.Time
	UpdatedOn  time.Time
}

type UserOtp struct {
	Id        int
	Uid       int
	SentOn    time.Time
	Validity  time.Time
	Otp       string
	Token     string
	SentTo    string
	Tries     int
	Validated int
	UpdatedOn time.Time
	CreatedOn time.Time
}

type UserAddress struct {
	Id        int
	Uid       int
	Line      string
	City      string
	District  string
	State     string
	Pincode   string
	UpdatedOn time.Time
	CreatedOn time.Time
}

type PersonalDetails struct {
	Name       string `json:"name"`
	Mobile     string `json:"mobile"`
	Gender     string `json:"gender"`
	Age        int    `json:"age"`
	Profession string `json:"profession"`
}

type AddressDetails struct {
	Line     string `json:"line"`
	City     string `json:"city"`
	District string `json:"district"`
	State    string `json:"state"`
	Pincode  string `json:"pincode"`
}

type UserDetails struct {
	Status string `json:"status"`
}

type NewToken struct {
	Token string `json:"token"`
}
