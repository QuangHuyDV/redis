package handler

import (
	"bufio"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckErr(c *gin.Context, err error) {
	if err != nil {
		c.JSON(200, gin.H{
			"err" : err.Error(),
		})
		return
	}
}



// ex7 
//Lưu đúng số dòng trên vào file name.txt đọc file ra theo từng dòng và Lpush vào list names. Nếu dòng ko có dữ liệu hoặc dữ liệu là "\s+" thì ko lưu vào redis.

// Lpop và inra gía trị lấy được.

// Set thời gian expire cho list là 5s. sau đó thêm 1 phần tử vào bên phải list là Dota2vn. Sau đó in ra danh sách.

func Ex7(c *gin.Context) {
	rdb := conn()
	f, err := os.Open("name.txt")
    CheckErr(c,err)
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
		if scanner.Text() == ""  {
			continue
		} else {
			err = rdb.LPush(ctx, "names", scanner.Text()).Err()
			CheckErr(c,err)
		}
    }
	err = rdb.RPush(ctx, "names", "Dota2vn").Err()
	CheckErr(c,err)

    err = scanner.Err(); 
	CheckErr(c,err)

	err = rdb.Expire(ctx,"names",5*time.Second).Err()
	CheckErr(c,err)

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

//ex6

//Sử dụng thư viện go-redis. Tạo 1 chương trình thêm một key demo_key có giá trị là time.Now().Unix() với thời gian expire 10s

//Sleep chương trình 12s đọc key demo_key ra nếu ko có key_demo thì lưu lại với giá trị thời gian hiện tại tính theo giây.

func Ex6(c *gin.Context) {
	rdb := conn()
	err := rdb.Set(ctx,"demo_key", time.Now().Unix(), time.Second*10).Err()
	CheckErr(c,err)

	time.Sleep(time.Second*12)

	val, err := rdb.Get(ctx,"demo_key").Result()
	if err != nil {
		a := time.Now().Second()
		err = rdb.Set(ctx, "demo_key", a, 10*time.Second).Err()
		CheckErr(c,err)

		c.JSON(200, gin.H{
			"Thoi gian theo giay hien tai:": a,
		})
		return 
	}
	c.JSON(200, gin.H{
		"Thoi gian da luu:": val,
	})
}
