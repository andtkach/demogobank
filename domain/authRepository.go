package domain

import (
	"encoding/json"
	"fmt"
	"github.com/andtkach/demogobank/logger"
	"net/http"
	"net/url"
	"os"
)

type AuthRepository interface {
	IsAuthorized(token string, routeName string, vars map[string]string) bool
}

type RemoteAuthRepository struct {
	AuthServer string
}

func (r RemoteAuthRepository) IsAuthorized(token string, routeName string, vars map[string]string) bool {

	u := buildVerifyURL(token, routeName, vars, r.AuthServer)

	if response, err := http.Get(u); err != nil {
		fmt.Println("Error while sending..." + err.Error())
		return false
	} else {
		m := map[string]bool{}
		if err = json.NewDecoder(response.Body).Decode(&m); err != nil {
			logger.Error("Error while decoding response from auth server:" + err.Error())
			return false
		}
		return m["isAuthorized"]
	}
}

/*
  This will generate a url for token verification in the below format

  /auth/verify?token={token string}
              &routeName={current route name}
              &customer_id={customer id from the current route}

  Sample: /auth/verify?token=aaaa.bbbb.cccc&routeName=MakeTransaction&customer_id=2000
*/
func buildVerifyURL(token string, routeName string, vars map[string]string, authServer string) string {
	u := url.URL{Host: authServer, Path: "/auth/verify", Scheme: "http"}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func NewAuthRepository() RemoteAuthRepository {
	return RemoteAuthRepository{
		AuthServer: os.Getenv("AUTH_SERVER"),
	}
}
