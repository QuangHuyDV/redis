package handler

import (
	"context"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Name struct {
	Name string
}

var ctx = context.Background()

func conn() (*redis.Client) {
	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", 
        DB:       0,  
    })
	return rdb
}
//ex8

// viết api get /names để lấy dữ liệu từ trong redis ra.(nhớ mở cái expire 5s của list).
func GetAll(c *gin.Context) {
	rdb := conn()
	s, err := rdb.LLen(ctx, "names").Result()
	CheckErr(c,err)

	for i := 0; i < int(s); i++ {
		names, err := rdb.LIndex(ctx, "names",int64(i)).Result()
		CheckErr(c,err)

		a := map[int32]string{
			int32(i) : names,
		}
		c.JSON(200, gin.H{
			"name" : a,
		})
	}
}

// viết api post /name để tạo 1 name theo struct Name {Name string} Lpush vào redis.
func InsertName(c *gin.Context) {
	rdb := conn()
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	CheckErr(c,err)

	err = rdb.LPush(ctx, "names",value).Err()
	CheckErr(c,err)

	c.JSON(200, gin.H{
		"them thanh cong": string(value),
	})
}

// viết api post /name/{index} để update tên theo vị trí trong redis.
func UpdateName(c *gin.Context) {
	rdb := conn()
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	CheckErr(c,err)

	index, err := strconv.ParseInt(c.Param("index"), 16, 64)
	CheckErr(c,err)

	name := Name{
		Name: string(value),
	}
	err = rdb.LSet(ctx, "names", index, name.Name).Err()
	CheckErr(c,err)

	c.JSON(200, gin.H{
		"Sua thanh cong" : name.Name,
	})
}

// viết api get /name/{index} để lấy name theo vị trí trong redis.
func ReadName(c *gin.Context, ) {
	rdb := conn()
	index, err := strconv.ParseInt(c.Param("index"),16,64)
	CheckErr(c,err)

	name, err := rdb.LIndex(ctx, "names", index).Result()
	CheckErr(c,err)

	c.JSON(200, gin.H{
		"Name": name,
	})
}

// viết api delete /name/{index} để xóa theo 1 index trong redis.
func DeleteName(c *gin.Context) {
	rdb := conn()
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	CheckErr(c,err)

	index, err := strconv.ParseInt(c.Param("index"), 16, 64)
	CheckErr(c,err)

	err = rdb.LRem(ctx, "names", index, string(value)).Err()
	CheckErr(c,err)

	c.JSON(200, gin.H{
		"Xoa thanh cong": string(value),
		"So luong": index,
	})
}