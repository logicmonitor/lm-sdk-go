# lm-sdk-go
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/logicmonitor/lm-sdk-go)

Getting Started
---------------

```go
package main

import (
    "fmt"
    "github.com/logicmonitor/lm-sdk-go/client"
    "github.com/logicmonitor/lm-sdk-go/client/lm"
)
func main() {
    // Configure API key authorization: LMv1
    domain := "YOUR_COMPANY.logicmonitor.com"
    accessID := "YOUR_ACCESS_ID"
    accessKey := "YOUR_ACCESS_KEY"

    config := client.NewConfig()
    config.SetAccountDomain(&domain)
    config.SetAccessID(&accessID)
    config.SetAccessKey(&accessKey)
    
    // create an instance of the API class
    lmSdk := client.New(config)
    params := lm.NewGetDeviceListParams()

    // ack alert by id
    resp, err := lmSdk.LM.GetDeviceList(params)
    if err != nil {
        fmt.Printf("Exception when calling client.LM.AckAlertByID: %v",err.Error())
    }
    fmt.Print(resp)
}
```

### License
[![license](https://img.shields.io/github/license/logicmonitor/lm-sdk-go.svg?style=flat-square)](https://github.com/logicmonitor/lm-sdk-go/blob/master/LICENSE)
