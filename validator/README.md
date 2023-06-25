# validator
参考：
- https://github.com/go-playground/validator
- https://juejin.cn/post/6863765115456454664

## 安装
```
go get github.com/go-playground/validator/v10
```

## 使用

简单示例
```
import "github.com/go-playground/validator/v10"

type User struct {
	Name  string `validate:"required"`
	Age   uint8  `validate:"gte=0,lte=130"`
	Email string `validate:"required,email"`
}

validate := validator.New()


user := User{"username", 18, "xxx@gmail.com"}
if err := validate.Struct(user); err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println("validation pass")
}
```

## tag
TODO