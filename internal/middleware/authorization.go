package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"log"
	"strings"
	"svc/proxy-service/internal/common"
)

func TokenParser() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken := c.Request.Header["Authorization"]
		if jwtToken == nil || len(jwtToken) == 0 {
			panic(common.ClientError{Code: 400, Message: "Empty JWT in header"})
		}
		parseJWT(jwtToken[0])
		// pass to other middleware
		c.Next()
	}
}

func parseJWT(rawToken string) {
	cleanToken := strings.Replace(rawToken, "Bearer ", "", -1)
	token, _, err := new(jwt.Parser).ParseUnverified(cleanToken, jwt.MapClaims{})
	if err != nil {
		log.Printf("Error when parse token %v \n", err)
		panic(common.ClientError{Code: 400, Message: "Unable to parse token"})
		return
	}

	// parse the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		jwtRoles := claims["cognito:groups"].([]interface{})
		jwtSiteIds := claims["custom:sites"].(string)
		jwtUsername := claims["cognito:username"].(string)
		jwtCustomerId := claims["custom:tenant"]

		siteIds := parseSiteId(jwtSiteIds)
		role := parseRole(jwtRoles)
		username := jwtUsername
		customerId := parseCustomerId(jwtCustomerId)

		reqContext := common.RequestContext{
			Username:   username,
			SiteIDs:    siteIds,
			Role:       role,
			CustomerId: customerId,
		}

		log.Printf("Request context: %v \n", reqContext)
	} else {
		log.Printf("Error when try to fill context %s \n", err)
		panic(common.ServerError{Code: 500, Message: "Error when try to fill context"})
	}
}

func parseSiteId(rawSiteIds string) []int {
	var numbers []int
	// Unmarshal (decode) the JSON data
	err := json.Unmarshal([]byte(rawSiteIds), &numbers)
	if err != nil {
		panic(common.ClientError{Code: 400, Message: "Error when parse token site ids"})
		return nil
	}
	return numbers
}

func parseRole(rawRoles []interface{}) []common.Role {
	roles := make([]common.Role, 0, len(rawRoles))
	for _, rawRole := range rawRoles {
		role := common.ToRole(rawRole.(string))
		roles = append(roles, role)
	}
	return roles
}
func parseCustomerId(customerId interface{}) *int {
	var result *int
	if customerId != nil {
		*result = customerId.(int)
	}
	return result
}
