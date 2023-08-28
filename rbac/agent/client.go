package agent

type User struct{}

type Permission string

type agent struct{}

func NewAgent() (*agent, error)

// api
func (a *agent) HasPermission(u User, p Permission) bool
