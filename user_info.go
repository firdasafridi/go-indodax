package indodax

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type respGetInfo struct {
	Success int
	Return  *UserInfo
	Message string
}

//
// User Info containt balance info, wallet address, user id, profile picture, username, and email of user.
//
type UserInfo struct {
	Balance        map[string]float64
	BalanceHold    map[string]float64
	WalletAddress  map[string]string
	UserId         int
	ProfilePicture string
	UserName       string
	ServerTime     time.Time
	Email          string
}

func (UserInfo *UserInfo) UnmarshalJSON(b []byte) (err error) {
	var kv map[string]interface{}

	err = json.Unmarshal(b, &kv)
	if err != nil {
		return err
	}

	for k, v := range kv {
		k = strings.ToLower(k)

		switch k {
		case fieldNameBalance:
			UserInfo.Balance, err = jsonToMapStringFloat64(v.(map[string]interface{}))
		case fieldNameBalanceHold:
			UserInfo.BalanceHold, err = jsonToMapStringFloat64(v.(map[string]interface{}))
		case fieldNameWalletAddress:
			UserInfo.WalletAddress, err = jsonToMapStringString(v.(map[string]interface{}))
		case fieldNameUserId:
			val64, err := strconv.Atoi(v.(string))
			if err != nil {
				return err
			}
			UserInfo.UserId = int(val64)
		case fieldNameProfilePicture:
			UserInfo.ProfilePicture = fmt.Sprintf("%v", v)
		case fieldNameEmail:
			UserInfo.Email = fmt.Sprintf("%v", v)
		case fieldNameUserName:
			UserInfo.UserName = fmt.Sprintf("%v", v)
		case fieldNameUserServerTime:
			ts, err := strconv.ParseInt(fmt.Sprintf("%.0f", v), 10, 64)
			if err != nil {
				return err
			}
			UserInfo.ServerTime = time.Unix(ts, 0)
		}
		if err != nil {
			return err
		}
	}

	return nil
}
