package main
import "encoding/json"
import "fmt"

type Student struct {
    Name string
    Age int
    score float64    // 不能对外开放的字段不会被编解码
}

func main() {
    Andy := Student{
        Name: "Andy",
        Age: 20,
        score: 100.0,
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
    // 注意：Unmarshal的m不会大写；第二个参数一般是引用&stu
    if err = json.Unmarshal(info, &stu); err != nil {
        fmt.Println(err)    
    }
    fmt.Println("decoding to: ", stu)    // 不能对外开放的字段不会被编解码，所以不会出现score
    fmt.Println(stu.Name)
}
