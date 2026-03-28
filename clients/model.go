package clients

import "fmt"

type Client struct {
	FIO      string
	Login    string
	Birthday string
	Email    string
}

func (c Client) String() string{
	fmt.Println("=======================")
	return fmt.Sprintf(
		"FIO: %s\nLogin: %s\nBirthday: %s\nEmail: %s\n",
		c.FIO,
		c.Login,
		c.Birthday,
		c.Email,
	)
}