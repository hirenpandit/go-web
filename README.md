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