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

## 标签
官方提供了多种类型的标签，比如字段、格式、字符串、比较，下面罗列部分常用的标签

完整标签见文档：https://github.com/go-playground/validator#baked-in-validations


### 格式
|  标签        |  解释       |
|  :---       |  :---       |
|  base64     |  base64     |
|  btc_addr   |  比特币地址   |
|  datetime   |  日期        |
|  email      |  邮箱        |
|  jwt        |  jwt        |
|  uuid       |  uuid       |
|  md5        |  md5 哈希    |

### 字符串
|  标签        |  解释        |
|  :---        |  :---       |
|  alpha       |  只含希腊字母 |
|  ascii       |  只含 ascii  |
|  contains    |  包含某字符串  |
|  startswith  |  以某字符串开头 |
|  endswith    |  以某字符串结尾 |
|  number      |  数字 （正整型）|
|  numeric     |  数字          |
|  lowercase   |  只包含小写字母  |
|  uppercase   |  只包含大写字母  |

### 比较
|  标签            |  解释            |
|  :---           |  :---            |
|  eq             |  相等             |
|  eq_ignore_case |  相等（忽略大小写） |
|  gt             |  大于             |
|  gte            |  大于等于          |
|  lt             |  小于             |
|  lte            |  小于等于          |
|  ne             |  不等于            |
|  ne_ignore_case |  不等于（忽略大小写）|

