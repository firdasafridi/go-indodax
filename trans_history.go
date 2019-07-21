package indodax

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type respTransHistory struct {
	Success int
	Return  *TransHistory
	Message string
}

//
// Transaction History containt list of Deposit and withdraw.
//
type TransHistory struct {
	Withdraw map[string][]TransWithdraw
	Deposit  map[string][]TransDeposit
}

type TransWithdraw struct {
	Status      string
	Type        string
	Amount      float64
	Fee         float64
	SubmitTime  time.Time
	SuccessTime time.Time
	ID          int64
}

type TransDeposit struct {
	Status      string
	Type        string
	Amount      float64
	Fee         float64
	SubmitTime  time.Time
	SuccessTime time.Time
	ID          int64
}

func (transWithdraw *TransWithdraw) UnmarshalJSON(b []byte) (err error) {
	var kv map[string]interface{}

	err = json.Unmarshal(b, &kv)
	if err != nil {
		return err
	}

	for k, v := range kv {
		k = strings.ToLower(k)

		switch k {
		case fieldNameStatus:
			transWithdraw.Status = v.(string)
		case fieldNameType:
			transWithdraw.Type = v.(string)
		case fieldNameFee:
			fee := v.(string)
			transWithdraw.Fee, err = strconv.ParseFloat(fee, 64)
		case fieldNameSubmitTime:
			ts, err := strconv.ParseInt(v.(string), 10, 64)
			if err != nil {
				return err
			}
			transWithdraw.SubmitTime = time.Unix(ts, 0)
		case fieldNameSuccessTime:
			ts, err := strconv.ParseInt(v.(string), 10, 64)
			if err != nil {
				return err
			}
			transWithdraw.SuccessTime = time.Unix(ts, 0)
		case fieldNameWithdrawID:
			transWithdraw.ID, err = strconv.ParseInt(v.(string), 10, 64)
		default:
			amount := v.(string)
			transWithdraw.Amount, err = strconv.ParseFloat(amount, 64)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (transDeposit *TransDeposit) UnmarshalJSON(b []byte) (err error) {
	var kv map[string]interface{}

	err = json.Unmarshal(b, &kv)
	if err != nil {
		return err
	}

	for k, v := range kv {
		k = strings.ToLower(k)

		switch k {
		case fieldNameStatus:
			transDeposit.Status = v.(string)
		case fieldNameType:
			transDeposit.Type = v.(string)
		case fieldNameFee:
			fee := v.(string)
			transDeposit.Fee, err = strconv.ParseFloat(fee, 64)
		case fieldNameSubmitTime:
			ts, err := strconv.ParseInt(v.(string), 10, 64)
			if err != nil {
				return err
			}
			transDeposit.SubmitTime = time.Unix(ts, 0)
		case fieldNameSuccessTime:
			ts, err := strconv.ParseInt(v.(string), 10, 64)
			if err != nil {
				return err
			}
			transDeposit.SuccessTime = time.Unix(ts, 0)
		case fieldNameDepositID:
			transDeposit.ID, err = strconv.ParseInt(v.(string), 10, 64)
		case fieldNameAmount:
			amount := v.(string)
			transDeposit.Amount, err = strconv.ParseFloat(amount, 64)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
