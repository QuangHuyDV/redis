package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserId 	int 	`json:"userId"`
	Id 		int		`json:"id"`
	Title 	string	`json:"title"`
	Body 	string	`json:"body"`
}


// get dữ liệu của api sau: https://jsonplaceholder.typicode.com/posts. Tạo 1 struct để hứng kết quả trả về.

func ReadUser(c *gin.Context) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		c.JSON(200, gin.H{
			"err": err,
		})
		return
	}
	defer resp.Body.Close()
	value, _ := ioutil.ReadAll(resp.Body)

	var result []User
	if err = json.Unmarshal(value, &result); err != nil {
		c.JSON(200, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"User": result,
	})
}

// post một đối tượng từ struct trên để tạo mới dữ liệu với api: https://jsonplaceholder.typicode.com/posts. Log ra kết quả ra màn hình.
func InsertUser(c *gin.Context) (){
	user := User{
		UserId: 11,
		Id: 111,
		Title: "Test",
		Body: "BodyTest",
	}

	body, _ := json.Marshal(user)

	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json",bytes.NewBuffer(body))

	if err != nil {
		c.JSON(200, gin.H{
			"err": err,
		})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		var result User
		if err = json.Unmarshal(body, &result); err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}
		c.JSON(200, gin.H{
			"User": result,
		})
	} else {
		if err != nil {
			c.JSON(200, gin.H{
				"Get failed with error: ": resp.Status,
			})
		}
	}
}
