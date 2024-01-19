package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
    Subject string `json:"Sub"`    // 重命名
    Score float64 `json:"score,float32"`    // 改变类型，注意：还需要提供字段名用来和重命名区分(可以改名)；名字和类型中间只有逗号没有空格
    Name string `json:"-"`    // json编解码忽略
    Age int `json:"Age,omitempty"`    // omitempty忽略空，如果字段是空的就忽略此字段不参与编码
    address string // 小写字段名(非公开)不参与编解码
}


func main() {
    Andy := Student{
	Subject: "math",
	Score: 100.0,
        Name: "Andy",
	Age: 0,
	address: "CN",
    }

    // 编码，返回字节slice和err
    // func Marshal(v interface{}) (byte[], error)
    info, err := json.Marshal(&Andy)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("encoding to: ", string(info))

    // 解码，参数为字节slice和要将json解析到的变量(比如struct，需要提前定义)
    var stu Student
    // func UnMarshal([]byte, v interface{}) error
    if err = json.Unmarshal(info, &stu); err != nil {
        fmt.Println(err)
    }
    fmt.Println("decoding to: ", stu)    // 不能对外开放的字段不会被编解码，所以不会出现score
}
