package pwd

import "golang.org/x/crypto/bcrypt"

type Encoder interface {
	Encode(rawPwd string) (string, error)
	Matches(rawPwd, encodePwd string) bool
}

func Encode(rawPwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(rawPwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func Matches(rawPwd, encodePwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encodePwd), []byte(rawPwd)) == nil
}
