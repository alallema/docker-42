package Users

import (
	"fmt"
	"log"
	"github.com/go-redis/redis"
	"Redis"
)

func AllUsers() *[]User {
	var userList []User
	var user *User
	var err error

	resultRequest, err := Redis.Client.Keys("email*").Result()
	if err == redis.Nil {
		fmt.Println("No Users in Database")
	} else if err != nil {
		log.Fatal(err)
	}

	for index, email := range resultRequest {
		_ = index
		user = FindUserByEmail(email)
		userList = append(userList, *user)
	}
	return &userList
}

func FindUserByEmail(email string) *User {
	var user User
	var err error

	result, err := Redis.Client.HGetAll(email).Result()
	if err == redis.Nil {
		fmt.Println("This User doesn't exist in Database")
	} else if err != nil {
		log.Fatal(err)
	}
	if len(result) == 0 {
		return nil
	}
	user.Email = result["Email"]
	user.Lastname = result["Name"]
	user.Firstname = result["Firstname"]
	user.Phone = result["Phone"]
	return &user
}

func DelUserByEmail(Key string) {
	resultRequest, err := Redis.Client.Del("email:" + Key).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		log.Fatal(err)
	}
	_ = resultRequest
}