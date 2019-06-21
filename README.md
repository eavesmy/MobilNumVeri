# 各国手机号码验证

### Install
``` go get "github.com/zdy23216340/MobilNumVeri" ```

### Usage
```
package main

import(
    "fmt"
    nmv "github.com/zdy23216340/MobilNumVeri"
)

const Tel = "18000000000"

func main(){

    if is := mnv.Veri("cn",Tel); is {
        fmt.Println("Tel is China mobile number")
    }

    if is := mnv.Veri("cn",Tel); is {
        fmt.Println("Tel is China mobile number")
    }
    
    from := mnv.Country(Tel)
    fmt.Printf("Tel might be from these country: %v",from)
}

```
