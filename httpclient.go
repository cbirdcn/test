package main
import "net/http"
import "fmt"
import "io/ioutil"

func main() {
    client := http.Client{}
    resp, err := client.Get("https://www.baidu.com")
    if err != nil {
        fmt.Println(err)
    }
    
    // 获取响应协议内容
    // resp.Header.Content-Type
    // resp.Request.URL
    // resp.StatusCode
    // resp.Status
    
    // 获取响应协议体
    // func ReadAll(r io.Reader) ([]byte, error)    返回字节slice和err
    bodyStr, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("read body err")
    }
    fmt.Println(string(bodyStr))
}