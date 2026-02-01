# Ip Locate 


# Télecharger le package

```shell
go get github.com/ZecolikFR/IpLocate
```

# Code exemple

```golang
info, err := iplocate.Request("8.8.8.8")
if err != nil {
	log.Fatal(err)
}

fmt.Println(info.Country, info.City, info.Isp)´´´
}
