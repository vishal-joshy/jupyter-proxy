package jupyter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

const url = "http://localhost:8000/hub/api"

func AddHeaders(req *http.Request) {
	req.Header.Add("Authorization", "Token secret-token")
	req.Header.Add("Content-Type", "application/json")
}

type JupyterUser struct {
	Name string `json:"name"`
}
type Test map[string]any

var client = http.Client{}

func GetUsers(c echo.Context) error {
	req, err := http.NewRequest("GET", url+"/users", nil)
	AddHeaders(req)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	var userList []JupyterUser
	if err := json.NewDecoder(res.Body).Decode(&userList); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", userList)
	return c.JSON(http.StatusOK, userList)
}

func GetUser(c echo.Context) error {
	username := c.Param("name")
	req, err := http.NewRequest("GET", url+"/users/"+username, nil)
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

type CreateUserForm struct {
	Name string `json:"name"`
}

func CreateUser(c echo.Context) error {
	var u CreateUserForm
	if err := c.Bind(&u); err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url+"/users/"+u.Name, nil)
	AddHeaders(req)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	var resBody any
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		fmt.Println(err)
	}
	fmt.Println(resBody)
	return c.JSON(http.StatusOK, resBody)
}
