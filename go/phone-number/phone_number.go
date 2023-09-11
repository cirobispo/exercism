package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

type NANP struct {
	CountryId       string
	AreaCode        string
	ExchageCode     string
	SubscribeNumber string
}

type ParsePhoneNumberError int

const (
	NoError                            ParsePhoneNumberError = 0
	WrongCountryError                  ParsePhoneNumberError = 1
	WrongAreaError                     ParsePhoneNumberError = 2
	WrongCountryAndAreaError           ParsePhoneNumberError = 3
	WrongExchangeError                 ParsePhoneNumberError = 4
	WrongCountryAndExchangeError       ParsePhoneNumberError = 5
	WrongAreaAndExchangeError          ParsePhoneNumberError = 6
	WrongCountryAndAreaAndExchageError ParsePhoneNumberError = 7
)

func parsePhoneNumber(phone string) (NANP, error) {
	re := regexp.MustCompile(`((\d)|(\d\D+))|((\d+)(\d+\D+))`)
	numbers := re.FindAllString(phone, -1)

	var data string
	for _, item := range numbers {
		data += item
	}

	var err error = nil
	result := NANP{}
	if (data == "") || (len(data) < 10 || len(data) > 11) {
		err = errors.New("phone number incorrect")
	} else {
		//2234567890
		result.CountryId = "1"

		index := 0
		if len(data) == 11 {
			index = 1
			result.CountryId = data[0:1]
		}
		result.AreaCode = data[index : 3+index]
		result.ExchageCode = data[3+index : 6+index]
		result.SubscribeNumber = data[6+index : 10+index]
	}
	return result, err
}

func isNANPOK(data NANP) ParsePhoneNumberError {
	wrongCountry := data.CountryId != "1"
	wrongArea := len(data.AreaCode) > 0 && (data.AreaCode[0] == '0' || data.AreaCode[0] == '1')
	wrongExchange := len(data.ExchageCode) > 0 && (data.ExchageCode[0] == '0' || data.ExchageCode[0] == '1')

	var result ParsePhoneNumberError

	if wrongCountry {
		result += WrongCountryError
	}

	if wrongArea {
		result += WrongAreaError
	}

	if wrongExchange {
		result += WrongExchangeError
	}

	return result
}

func verifyError(en ParsePhoneNumberError) error {
	switch en {
	case WrongCountryError:
		return errors.New("country code incorrect")
	case WrongAreaError:
		return errors.New("area code incorrect")
	case WrongExchangeError:
		return errors.New("exchange code incorrect")
	case WrongCountryAndAreaError:
		return errors.New("country code and area code are incorrect")
	case WrongCountryAndExchangeError:
		return errors.New("country code and exchange code are incorrect")
	case WrongAreaAndExchangeError:
		return errors.New("area code and exchange code are incorrect")
	case WrongCountryAndAreaAndExchageError:
		return errors.New("country code, area code and exchange code are incorrect")
	default:
		return nil
	}
}

func Number(phoneNumber string) (string, error) {
	if pn, err := parsePhoneNumber(phoneNumber); err == nil {
		err = verifyError(isNANPOK(pn))
		return pn.AreaCode + pn.ExchageCode + pn.SubscribeNumber, err
	} else {
		return "", err
	}
}

func AreaCode(phoneNumber string) (string, error) {
	if pn, err := parsePhoneNumber(phoneNumber); err == nil {
		err = verifyError(isNANPOK(pn))
		return pn.AreaCode, err
	} else {
		return "", err
	}
}

func Format(phoneNumber string) (string, error) {
	if pn, err := parsePhoneNumber(phoneNumber); err == nil {
		err = verifyError(isNANPOK(pn))
		return fmt.Sprintf("(%s) %s-%s", pn.AreaCode, pn.ExchageCode, pn.SubscribeNumber), err
	} else {
		return "", err
	}
}
