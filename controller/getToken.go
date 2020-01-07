package controller

import ("github.com/gin-gonic/gin"
			"github.com/dgrijalva/jwt-go"
			// jt "github.com/gbrlsnchs/jwt"
			"fmt"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("qwerty")

func GetToken(c *gin.Context) {

// expirationTime := time.Now().Add(5 * time.Minute)

claims := &Claims{
	Username: "123456",
	StandardClaims: jwt.StandardClaims{
	},
}

token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
tokenString, err := token.SignedString(jwtKey)

	c.JSON(200, gin.H{
		"token": tokenString,
	})

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		fmt.Println("err")
	}
		
}

func VerifyToken(c *gin.Context) {
	tokenString := c.Query("token")
	fmt.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    // Don't forget to validate the alg is what you expect:
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
    }
    return jwtKey, nil
	})
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Validated !!")
		// TODO: Check for valid authtoken
		// custom validation function checking for valid authtoken from database
		// if not valid throw validation error
	} else {
    fmt.Println(err)
	}
}