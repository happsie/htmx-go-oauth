package router

import (
	"github.com/happise/pixelwars/container"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"slices"
)

func GetJwtConfig(container container.Container) echojwt.Config {
	jwtWhitelist := []string{"/api/auth/login", "/api/auth/callback", "", "/", "/css*", "/images*"}
	return echojwt.Config{
		TokenLookup: "cookie:token",
		SigningKey:  []byte(container.GetConfig().JWT.Secret),
		Skipper: func(c echo.Context) bool {
			return slices.Contains(jwtWhitelist, c.Path())
		},
	}
}

/*

SuccessHandler: func(c echo.Context) {
cookie, err := c.Request().Cookie("token")
if err != nil {
c.NoContent(http.StatusUnauthorized)
return
}
tokenString := strings.Split(cookie.Value, " ")[1]
tokenString = strings.ReplaceAll(tokenString, `"`, "")
token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
container.GetLogger().Warn("unexpected signing method", "signing method", token.Header["alg"])
return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
}
return []byte(container.GetConfig().JWT.Secret), nil
})
if err != nil {
container.GetLogger().Warn("could not parse jwt", "error", err.Error())
c.NoContent(http.StatusUnauthorized)
return
}
if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
/*exp, convOk := strconv.ParseInt(fmt.Sprintf("%v", claims["exp"]), 10, 64)
if convOk != nil {
container.GetLogger().Warn("could not parse expire time", "exp", claims["exp"])
c.NoContent(http.StatusUnauthorized)
return
}
// TODO: exp seems broken.
fmt.Printf("%v", claims)
c.Set("auth", model.JwtInfo{
	UserId:   fmt.Sprintf("%v", claims["userId"]),
	Username: fmt.Sprintf("%v", claims["username"]),
	//Exp:      exp,
})
} else {
	c.NoContent(http.StatusUnauthorized)
	return
}
},
*/
