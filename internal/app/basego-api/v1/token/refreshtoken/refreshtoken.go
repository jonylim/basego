package refreshtoken

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/crypto/hash"
	"github.com/jonylim/basego/internal/pkg/common/helper"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWTClaims defines information saved in a JWT.
type JWTClaims struct {
	TokenString string `json:"tkn"`
	TokenID     int64  `json:"tid"`
	SessionID   int64  `json:"sid"`
	AccountID   int64  `json:"uid"`
	jwt.StandardClaims
}

// TokenTTL defines how long a refresh token is valid for before it expires, in seconds.
const TokenTTL = 86400 * 30

// Parse error code.
const (
	ErrParseFailed   = 1
	ErrTokenInvalid  = 2
	ErrDeviceInvalid = 3
	ErrNotOwner      = 4
	ErrTokenExpired  = 5
)

var seed = 999999

// GenerateRefreshToken returns a new refresh token for a user's session.
func GenerateRefreshToken(sessionID int64) string {
	now := time.Now()
	hash := hash.SHA512inHex("reftkn:" + helper.Int64ToString(sessionID) + "@" + helper.Int64ToString(now.UnixNano()) + "x" + getSeed())
	return hash
}

// GenerateJWT converts a refresh token string into JWT.
func GenerateJWT(tokenID int64, tokenString string, issuedAt, expiresAt int64, sessionID int64, accountID int64) (string, error) {
	strTokenID := helper.Int64ToString(tokenID)
	jwtID := helper.Int64ToString(issuedAt) + "r" + strTokenID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		TokenString: tokenString,
		TokenID:     tokenID,
		SessionID:   sessionID,
		AccountID:   accountID,
		StandardClaims: jwt.StandardClaims{
			Id:        jwtID,
			Issuer:    "cstd/customer",
			IssuedAt:  issuedAt,
			ExpiresAt: expiresAt,
		},
	})
	str, err := token.SignedString(createJWTSecretKey(jwtID))
	if err != nil {
		logger.Error("refreshtoken", err.Error())
	}
	return str, err
}

// ParseJWT returns a parsed token containing the refresh token's details.
func ParseJWT(jwtToken string) (*JWTClaims, int, error) {
	claims, err := parseJWT(jwtToken)
	if err != nil {
		var code = ErrParseFailed
		errText := err.Error()
		if errText == jwt.ErrSignatureInvalid.Error() {
			err = errors.New("Token signature is invalid, please get a new token")
		} else {
			if claims != nil && claims.VerifyExpiresAt(time.Now().Unix(), false) == false {
				code = ErrTokenExpired
			}
			err = errors.New(helper.UcFirst(errText))
		}
		return claims, code, err
	} else if claims == nil {
		return claims, ErrParseFailed, errors.New("Failed retrieving token detail")
	}
	return claims, 0, nil
}

// ValidateState compares the parsed token with saved account session's details and refresh token.
func (claims *JWTClaims) ValidateState(accountSession model.CstAccountSession, sessionToken model.CstAccountSessionToken, apiKey model.XAPIKey, deviceID string) (int, error) {
	if claims.TokenString != sessionToken.RefreshToken || claims.TokenID != sessionToken.ID {
		return ErrTokenInvalid, errors.New("Refresh token is invalid")
	} else if apiKey.AppPlatform != accountSession.Platform ||
		deviceID != accountSession.DeviceID {
		logger.Debug("refreshtoken", "Refresh token does not belong to the device"+
			"\n    Platform-1: "+accountSession.Platform+
			"\n    Platform-2: "+apiKey.AppPlatform+
			"\n    DeviceID-1: "+accountSession.DeviceID+
			"\n    DeviceID-2: "+deviceID)
		return ErrDeviceInvalid, errors.New("Refresh token does not belong to the device")
	} else if claims.SessionID != accountSession.ID || claims.AccountID != accountSession.AccountID {
		return ErrNotOwner, errors.New("Refresh token does not belong to the user")
	} else if now := helper.UnixMillisecond(time.Now()); now > sessionToken.RefreshTokenExpiry {
		logger.Debug("refreshtoken", fmt.Sprintf("ErrTokenExpired: (%v) %v > %v", accountSession.ID, now, sessionToken.RefreshTokenExpiry))
		return ErrTokenExpired, errors.New("Refresh token is expired")
	}
	return 0, nil
}

func parseJWT(jwtToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate expected alg.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key.
		claims, ok := token.Claims.(*JWTClaims)
		if claims != nil && ok {
			return createJWTSecretKey(claims.Id), nil
		}
		return nil, errors.New("Failed to parse token")
	})
	if token != nil {
		claims, ok := token.Claims.(*JWTClaims)
		if claims != nil && ok {
			return claims, err
		}
	}
	return nil, err
}

func createJWTSecretKey(jwtID string) []byte {
	var key = "refs:" + jwtID
	logger.Trace("refreshtoken", fmt.Sprintf(`createJWTSecretKey: { jwtID: "%s", key: "%s" }`, jwtID, key))
	b := sha256.Sum256([]byte(key))
	return b[:]
}

func getSeed() string {
	seed = seed - 1
	if seed < 100000 {
		seed = 999999
	}
	return helper.IntToString(seed)
}
