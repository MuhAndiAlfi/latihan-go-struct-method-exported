package management

import "fmt"

type User struct {
	ID        int
	FristName string
	LastName  string
	Email     string
	IsActive  bool
}

// METHOD

func (user User) Display() string {
	result := fmt.Sprintf("Name: %s %s, Email: %s", user.FristName, user.LastName, user.Email)

	return result
}
func displayUser(user User) string {
	result := fmt.Sprintf("Name: %s %s, Email: %s", user.FristName, user.LastName, user.Email)

	return result
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func (group Group) DisplayGroup() {
	fmt.Printf("Name: %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count: %d", len(group.Users))
	fmt.Println("")

	fmt.Println("user name: ")
	for _, user := range group.Users {
		fmt.Println(user.FristName)
	}
}
