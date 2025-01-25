### GO WEB APPLICATION

> Example of go based web server to demonstrate all the 
> supported functionalities of web server in GO


### Structure

The is structured based on the guidline provided [here](https://go.dev/doc/modules/layout)

```text
+go-web
    |
    | +cmd
        |
        | +go-web
           | +main.go
    | +internal
        | +handlers
        | +server
    | +pkg     
```

**Note**: `internal` is special kind of package which prevent accessing code outside the `internal` package
    - https://pkg.go.dev/cmd/go#hdr-Internal_Directories

The idea is the package can be imported by nearby code not any other package


### HTMX

Configured web-server to serve HTML templates on root endpoint i.e / using html/template package.
`index.html` file is created with HTMX script tag.


### Middleware
Added middleware to log message of endpoint, HTTP method and time taken to serve response from the endpoint


### Websocket
Added support for websocket with gorilla websocket package. Follow below steps to test sending and receiving message using websocket

1. Open console in browser
2. Define websocket variable `var ws = new Websocket("ws://localhost:8080/ws")`
3. Send message `ws.send("test")`