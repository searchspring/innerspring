package google

import (
	"encoding/json"

	"searchspring.com/innerspring/common"
)

var companyDirectory = "1GEZId7qG4uMk76j0bEMukLI5tGP00zhY4u5Tz663EEY"

type DAO interface {
	CheckUserLoggedIn(token string) (string, error)
}
type DAOImpl struct {
	Client *common.Client
}

func NewDAO(client *common.Client) DAO {
	return &DAOImpl{
		Client: client,
	}
}

// CheckUserLoggedIn is this user signed in with google.
func (d *DAOImpl) CheckUserLoggedIn(token string) (string, error) {
	body, err := d.Client.AuthorizedGet(token, "https://www.googleapis.com/oauth2/v2/userinfo?access_token="+token)
	if err != nil {
		return "", err
	}
	type emailHolder struct {
		Email string `json:"email"`
	}
	user := &emailHolder{}
	err = json.Unmarshal(body, user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}
