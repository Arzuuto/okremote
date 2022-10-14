package server

import (
	"Backend/ent"
	"Backend/ent/users"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

func (s Server) ApplyRoutes() {
	s.Router.Use(cors.Default())
	s.Router.GET("/health", health)
	s.Router.POST("/register", s.register)
	s.Router.POST("/login", s.CheckLogin)
	s.Router.POST("/CreateAnnonce", s.CreateAnnonce)
	s.Router.GET("/home", s.Home)
	// s.Router.GET("/homeCon", s.HomeCon)
}

func (s Server) Home(c *gin.Context) {
	body := s.Db.AskAdvice(c)
	c.JSON(http.StatusOK, body)
}
// func (s Server) HomeCon(c *gin.Context) {
// 	var mySigningKey = []byte("remoteOk")
// 	cookie, err := c.Request.Cookie("Token")

// 	if (err != nil) {
// 		c.String(http.StatusForbidden, "Token not found")
// 		return
// 	}
// 	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("There was an error in parsing")
// 		}
// 		return mySigningKey, nil
// 	})
// 	if err != nil {
// 		c.String(http.StatusForbidden, "Token not found")
// 	}
// 	s.
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		if claims["email"] == v.Email && claims["password"] == v.Password {
// 			c.JSON(http.StatusOK,claims["email"])
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusUnauthorized, "Not authorized")
// }
	


func health(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

type User struct {
	Name     string
	Prenom   string
	Email    string
	Password string
}

type Login struct {
	Email    string
	Password string
}

type Announce struct {
	Contrat      string
	Date         string
	Entreprise   string
	Image        string
	Location     string
	Post_name    string
	Remuneration string
	Sector       string
}

func (s Server) CreateAnnonce(c *gin.Context) {
	body := Announce{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	u, err := s.Db.CreateAdv(c, body.Contrat, body.Date, body.Entreprise, body.Image, body.Location, body.Post_name, body.Remuneration, body.Sector)
	if err != nil {
		c.String(http.StatusBadRequest, "%v", err)
		return
	}
	c.JSON(http.StatusCreated, u)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s Server) CheckLogin(c *gin.Context) {
	body := Login{}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	_, err := s.Db.Client.Users.Query().Select("email").Where(users.Email(body.Email)).Only(c.Request.Context())
	if ent.IsNotFound(err) {
		c.String(http.StatusBadRequest, "%s", "Email not found")
		return
	}
	pass, err := s.Db.Client.Users.Query().Select("password").Strings(c.Request.Context())
	if ent.IsNotFound(err) {
		fmt.Println(pass)
		c.String(http.StatusBadRequest, "%v", err)
		log.Printf("%v\n", err)
		return
	}
	for _, hash := range pass {
		if CheckPasswordHash(body.Password, hash) {
			claims["authorized"] = true
			claims["email"] = body.Email
			claims["password"] = body.Password
			claims["exp"] = time.Now().Add(time.Minute * 120).Unix()
			tokenString, err := token.SignedString([]byte("remoteok"))

			if err != nil {
				c.String(http.StatusBadRequest, "%v", err)
				log.Printf("%v\n", err)
				return
			}
			c.SetCookie("Token", tokenString, 10, "/", "localhost", false, true)
			c.String(http.StatusOK, "next")
			return
		}
	}
	c.String(http.StatusBadRequest, "%s", "Password not found")

}

func (s Server) register(c *gin.Context) {
	user := User{}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	_, err := s.Db.Client.Users.Query().Select("email").Where(users.Email(user.Email)).Only(c.Request.Context())
	if ent.IsNotFound(err) {
		user.Password, err = HashPassword(user.Password)
		if err != nil {
			c.String(http.StatusBadRequest, "%v", err)
			log.Printf("%v\n", err)
			return
		}
		_, err := s.Db.CreateUser(c.Request.Context(), user.Name, user.Prenom, user.Email, user.Password)
		if err != nil {
			c.String(http.StatusBadRequest, "%v", err)
			log.Printf("%v\n", err)
			return
		}
		claims["authorized"] = true
		claims["email"] = user.Email
		claims["password"] = user.Password
		claims["exp"] = time.Now().Add(time.Minute * 120).Unix()
		tokenString, err := token.SignedString([]byte("remoteok"))
		if err != nil {
			c.String(http.StatusBadRequest, "%v", err)
			log.Printf("%v\n", err)
			return
		}
		c.SetCookie("Token", tokenString, 10, "/", "localhost", false, true)
		c.String(http.StatusOK, "Next")
		fmt.Println(user.Name + " " + user.Prenom + " " + user.Email + " " + user.Password)
	} else {
		c.String(http.StatusBadRequest, "%v", err)
		log.Printf("%v\n", err)
	}
}

func (s Server) LogIn(c *gin.Context) {
	user := User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		log.Printf("%v\n", err)
		return
	}
	sp := s.Db.AskPass(c)
	for _, vari := range sp {
		if vari.Email == user.Email && vari.Password == user.Password {
			log.Println("ca passe")
			break
		}
	}
}
