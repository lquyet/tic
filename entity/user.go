package entity

import (
	"fmt"
	"math/rand"
)

type User struct {
	ID               string `json:"id" faker:"uuid_digit"`
	Name             string `json:"name" faker:"name"`
	Email            string `json:"email" faker:"email"`
	Phone            string `json:"phone" faker:"phone_number"`
	CreditCardNumber string `json:"creditCardNumber" faker:"cc_number"`
	Avatar           string `json:"avatar"`
	JoinedDate       int64  `json:"joinedDate" faker:"unix_time"`
	Age              int    `json:"age" faker:"oneof:18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33"`
}

type GetUsersResponse struct {
	Data       []User     `json:"data"`
	Pagination Pagination `json:"pagination"`
}

func (u *User) GenImageURL() {
	u.Avatar = fmt.Sprintf("https://picsum.photos/id/%d/300/200", rand.Intn(99)+1)
}
