package main

import (
    "net/url"
    "fmt"
    "net"
)

func main() {

    stdoutLn := fmt.Println

    sqlUrl := "postgres://user:pass@host.com:5432/somePath?k=v#f"
    stdoutLn("Url:", sqlUrl)
    stdoutLn("----------------------------------------")

    urlParts, error := url.Parse(sqlUrl)
    if error != nil {
        panic(error)
    }

    stdoutLn("Url Schema:", urlParts.Scheme)

    stdoutLn("User:", urlParts.User)
    stdoutLn("Username:", urlParts.User.Username())

    password, _ := urlParts.User.Password()
    stdoutLn("Password:", password)

    stdoutLn("Full Host:", urlParts.Host)
    host, port, _ := net.SplitHostPort(urlParts.Host)
    stdoutLn("Host:", host)
    stdoutLn("Port:", port)

    stdoutLn("Path:", urlParts.Path)
    stdoutLn("Fragments", urlParts.Fragment)

    stdoutLn(urlParts.RawQuery)
    rawQueryMap, _ := url.ParseQuery(urlParts.RawQuery)
    stdoutLn("Raw Query String Map:", rawQueryMap)
    stdoutLn("Query Parameter 'k':", rawQueryMap["k"][0])
}
