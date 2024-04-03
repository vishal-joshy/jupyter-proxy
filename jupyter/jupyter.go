package jupyter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const hubUrl = "http://localhost:8000/hub/api/"
const userUrl = "http://localhost:8000/user/"

// Add admin auth headers to hub api requests.
// Token available in hub configuration
func AddHeaders(req *http.Request) {
	req.Header.Add("Authorization", "Token secret-token")
	req.Header.Add("Content-Type", "application/json")
}

type JupyterUser struct {
	Name string `json:"name"`
}

var client = http.Client{}

func GetUsers(c echo.Context) error {
	req, err := http.NewRequest("GET", hubUrl+"/users", nil)
	AddHeaders(req)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	var userList []JupyterUser
	if err := json.NewDecoder(res.Body).Decode(&userList); err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, userList)
}

func GetUser(c echo.Context) error {
	username := c.Param("name")
	req, err := http.NewRequest("GET", hubUrl+"/users/"+username, nil)
	AddHeaders(req)
	res, err := client.Do(req)
	var user JupyterUser
	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, user)
}

type CreateUserInput struct {
	Name string `json:"name"`
}

func CreateUserAndStartNotebook(c echo.Context) error {
	var u CreateUserInput
	if err := c.Bind(&u); err != nil {
		return err
	}
	req, err := http.NewRequest("POST", hubUrl+"/users/"+u.Name, nil)
	AddHeaders(req)
	fmt.Println("Creating User ", u.Name)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	var resBody any
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		fmt.Println(err)
	}
	fmt.Println("User Created", u.Name)

	userToken, err := GetToken(u.Name)
	isNotebookCreated := CreateNotebook(u.Name, userToken)
	if isNotebookCreated {
		notebookUrl := userUrl + u.Name + "?token=" + userToken
		return c.JSON(http.StatusOK, notebookUrl)
	}
	return c.JSON(http.StatusInternalServerError, nil)
}

type JupyterUserToken struct {
	Created string
	User    string
	Token   string
}

func GetToken(username string) (string, error) {
	req, err := http.NewRequest("POST", hubUrl+"/users/"+username+"/tokens", nil)
	AddHeaders(req)
	fmt.Println("Creating Token")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	var resBody JupyterUserToken
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		fmt.Println(err)
	}
	return resBody.Token, nil
}

// 201 server started
// 202 server requested not started
// 400 server running
func CreateNotebook(username string, token string) bool {
	fmt.Println("Starting Notebook for user" + username)
	req, _ := http.NewRequest("POST", hubUrl+"/users/"+username+"/server", nil)
	AddHeaders(req)
	for {
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("Notebook creation failed" + err.Error())
			return false
		}
		if res.StatusCode == 400 {
			return true
		}

		time.Sleep(4 * time.Second)
	}
}
