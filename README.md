# Go HTTP

[Net package](https://golang.org/pkg/net/http/)

```go
resp, err := http.Get("http://google.com")

if err != nil {
    fmt.Println("Error", err)
    os.Exit(1)
}

fmt.Println(resp)
```
![Reponse](https://github.com/alytvynov/go-http/blob/main/doc/response.png)

```go
//1st way
io.Copy(os.Stdout, resp.Body)
```

```go
//2nd way
bs := make([]byte, 999999)
len, readErr := resp.Body.Read(bs)

fmt.Println(len)
fmt.Println(string(bs))

if readErr != nil && readErr != io.EOF {
    fmt.Println("Read Error", readErr)
    os.Exit(1)
}
```