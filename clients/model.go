package clients

import "fmt"

type Client struct {
	Id       int
	FIO      string
	Login    string
	Birthday string
	Email    string
}

func (c Client) String() string {
	fmt.Println("=======================")
	return fmt.Sprintf(
		"ID: %d\nFIO: %s\nLogin: %s\nBirthday: %s\nEmail: %s\n",
		c.Id,
		c.FIO,
		c.Login,
		c.Birthday,
		c.Email,
	)
}
