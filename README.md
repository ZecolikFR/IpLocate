# Ip Locate 


# Télecharger le package

```shell
go get github.com/ZecolikFR/IpLocate
```

# Code exemple

```golang
package main

import(
"fmt"
"github.com/ZecolikFR/IpLocate"
)

func main(){
Locate.Request("8.8.8.8")
fmt.Println(Locate.GetCity())
