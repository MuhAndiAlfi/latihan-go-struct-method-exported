package main

import (
	"latihan-go-struct/management"
)

func main() {
	// user := User{}
	// user.ID = 1
	// user.FristName = "Andi"
	// user.LastName = "Alfiansyah"
	// user.email = "andialfiansyah@gmail.com"
	// user.isActive = true

	// user2 := User{
	// 	ID:        2,
	// 	FristName: "Andi",
	// 	LastName:  "ALf",
	// 	email:     "a@gmail",
	// 	isActive:  true,
	// }

	user3 := management.User{3, "LE", "LA", "LELA@gmail.com", true}
	// result := user3.display()

	// fmt.Println(result)

	user4 := management.User{4, "LA", "LA", "LALA@gmail.com", true}

	// fmt.Println(user)
	// fmt.Println(user2)

	// displayUser1 := displayUser(user3)
	// displayUse2 := displayUser(user4)

	// fmt.Println(displayUser1)
	// fmt.Println(displayUse2)

	users := []management.User{user3, user4}
	group := management.Group{"gamer", user3, users, true}

	group.DisplayGroup()

}

// func displayGroup(group Group) {
// 	fmt.Printf("Name: %s", group.Name)
// 	fmt.Println("")
// 	fmt.Printf("Member count: %d", len(group.Users))
// 	fmt.Println("")

// 	fmt.Println("user name: ")
// 	for _, user := range group.Users {
// 		fmt.Println(user.FristName)
// 	}
// }

// func displayUser(user User) string {
// 	result := fmt.Sprintf("Name: %s %s, Email: %s", user.FristName, user.LastName, user.email)

// 	return result
// }
