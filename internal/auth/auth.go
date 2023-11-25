package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey - from Authorization header
//format
// Authorization: ApiKey {inserted key}
func GetApiKey(headers http.Header)(string,error) {
	val:= headers.Get("Authorization")
	if val==""{

		return "", errors.New("no authentication found")
	}
	vals := strings.Split(val, " ")
	if len(vals)!=2 || vals[0]!="ApiKey" {
		return "",errors.New("not an appropriate header ")
	}

	return vals[1],nil

}