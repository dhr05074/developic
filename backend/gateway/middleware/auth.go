package middleware

import (
	"errors"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"time"
)

func ParseAndValidateJWT(plain string) (string, error) {
	token, err := jwt.Parse(
		[]byte(plain),
		jwt.WithKey(jwa.HS256, []byte("secret")),
		jwt.WithValidate(true),
		jwt.WithAcceptableSkew(10*time.Second))
	if err != nil {
		return "", err
	}

	// 파싱된 JWT 토큰에서 원하는 값을 가져올 수 있습니다.
	sub, ok := token.Get("sub")
	if !ok {
		return "", errors.New("sub not found")
	}

	return sub.(string), nil
}
