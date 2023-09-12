## SQL 语句到结构体的转换

```shell
go run root.go sql struct --username root --password 123456 --db=feature --table user
```

```go
type User struct {
	// id
	Id int32 `json:"id"`
	// user_name
	UserName string `json:"user_name"`
	// avatar
	Avatar string `json:"avatar"`
	// create_time
	CreateTime int8 `json:"create_time"`
	// user_id
	UserId int32 `json:"user_id"`
	// password
	Password string `json:"password"`
	// phone
	Phone int8 `json:"phone"`
}

func (model User) TableName() string {
	return "user"
}
```
