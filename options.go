package httputils

type option func(*Client)

// AddAuth adds basic authentication to the client
func AddAuth(username string, password string) option {
	return func(c *Client) {
		c.Username = username
		c.Password = password
	}
}
