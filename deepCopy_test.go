package deepCopy

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Transaction struct {
	AssociatedBanks []*Bank
	Name            *string

	Users map[string]User

	Dates []Date

	Me User
}

type Bank struct {
	Id    *string
	money int
}

type Date struct {
	Days []*string
}
type User struct {
	Name     string
	LastName *string
	Id       string
	Friends  []User
}

func copyTransaction(old Transaction) Transaction {
	return Transaction{
		AssociatedBanks: CopySlice(old.AssociatedBanks, func(b *Bank) *Bank {
			return CopyPointer(b)
		}),
		Name:  CopyPointer(old.Name),
		Users: CopyMap(old.Users, copyUser),
		Dates: CopySlice(old.Dates, copyDate),
		Me:    copyUser(old.Me),
	}
}

func copyUser(u User) User {
	return User{
		Name:     u.Name,
		LastName: CopyPointer(u.LastName),
		Id:       u.Id,
		Friends:  CopySlice(u.Friends, copyUser),
	}
}

func copyDate(d Date) Date {
	return Date{
		Days: CopySlice(d.Days, func(s *string) *string {
			return CopyPointer(s)
		}),
	}
}

func copyBank(b Bank) Bank {
	return Bank{
		Id:    CopyPointer(b.Id),
		money: b.money,
	}
}
func copyUserOld(u User) User {
	var lastName *string

	if u.LastName != nil {
		lastNameVal := *u.LastName
		lastName = &lastNameVal
	}

	var friends []User
	if u.Friends != nil {
		friends = make([]User, 0, len(u.Friends))
		for _, friend := range u.Friends {
			friends = append(friends, copyUser(friend))
		}
	}

	return User{
		Name:     u.Name,
		Id:       u.Id,
		LastName: lastName,
		Friends:  friends,
	}
}

func TestDeepCopy(t *testing.T) {
	myLast := "lev"
	user1 := User{
		Name:     "aryeh",
		LastName: &myLast,
		Id:       "0",
		Friends: []User{{
			Name:     "john",
			LastName: nil,
			Id:       "1",
			Friends: []User{{
				Name:     "",
				LastName: nil,
				Id:       "",
				Friends:  nil,
			}},
		}},
	}

	//var user2 = copyUser(user1)
	var user2 User
	b, _ := json.Marshal(user1)

	json.Unmarshal(b, &user2)

	if user2.LastName != nil {
		*user2.LastName = "notlev"
	}

	if len(user2.Friends) > 0 {
		user2.Friends[0] = User{
			Name:     "notjoe",
			LastName: nil,
			Id:       "",
			Friends:  nil,
		}
	}

	fmt.Println(*user1.LastName)
	fmt.Println(user1.Friends)

	var rando1 = map[string]string{"aryeh": "1"}

	rando2 := CopyMap(rando1)

	rando2["batel"] = "2"

	fmt.Println(rando1)
}
