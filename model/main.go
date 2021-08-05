package main

import (
	"fmt"
	"strconv"
)

func main() {

	var users []User
	for i := 1; i <= 10; i++ {
		users = append(users, User{
			name: strconv.Itoa(i) + "user",
			city: strconv.Itoa(i) + "city",
			age:  i,
		})
	}

	fmt.Println(users)

	// 获取到的size 和page
	size := 4
	page := 3

	totalCount := len(users)

	pageCount := totalCount / size
	if totalCount%size > 0 {
		pageCount += 1
	}

	// 分页数据
	var res []User
	if page*size > totalCount {
		res = users[(page-1)*size : totalCount]
	} else {
		res = users[(page-1)*size : page*size]
	}

	fmt.Println(res)

	fmt.Printf(" %v\n", res)

}
