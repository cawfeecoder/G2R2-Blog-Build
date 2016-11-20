package servicesAuth

import (
	"fmt"
	"time"

	r "github.com/dancannon/gorethink"
	"github.com/dchest/uniuri"
	"github.com/dgrijalva/jwt-go"
	"github.com/nfrush/G2R2-Blog-Build/database/rethink"
	"github.com/nfrush/G2R2-Blog-Build/models/user"
	"github.com/nfrush/G2R2-Blog-Build/services/token"
	"golang.org/x/crypto/bcrypt"
)

var session = rethink.GetSession()

var signingKey = InitSigningKey()

//InitSigningKey - Initalize Our Key To Sign With
func InitSigningKey() string {
	return uniuri.NewLen(32)
}

//CompareHash conmpare hash to stored hash
func CompareHash(u *modelUser.User) error {
	result, errA := r.Table("users").Filter(map[string]interface{}{"Name": u.Username}).Pluck("Password").Run(session)
	if errA != nil {
		return errA
	}
	var user modelUser.User
	result.One(&user)
	result.Close()

	fmt.Println([]byte(user.Password))

	errB := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if errB != nil {
		return errB
	}
	return nil
}

//Login - Authenicate and reutrn JWT
func Login(u *modelUser.User) (string, error) {
	if err := CompareHash(u); err != nil {
		return "", err
	}
	exists, errA := servicesToken.TokenExistsUser(u)
	if errA != nil {
		return "", errA
	}
	if exists == false {
		token, errB := servicesToken.IssueToken(u)
		if errB != nil {
			return "", errB
		}
		return token, nil
	}
	if errC := servicesToken.RevokeToken(u); errC != nil {
		return "", errC
	}
	token, errD := servicesToken.IssueToken(u)
	if errD != nil {
		return "", errD
	}
	return token, nil
}

//Logout - Revoke JWT Token
func Logout(u *modelUser.User) error {
	if err := servicesToken.RevokeToken(u); err != nil {
		return err
	}
	return nil
}

//Refresh Token
func Refresh(u *modelUser.User) (string, error) {
	token, err := servicesToken.RefreshToken(u)
	if err != nil {
		return "errors", err
	}
	return token, nil
}

//TestKey - Generate Test Key
func TestKey() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "Aurelia Development LTD",
		"aud": "Tester",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"jti": "http://example.com",
	})

	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	fmt.Println(tokenString)

	return tokenString, nil
}
