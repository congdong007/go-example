package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	pl := fmt.Println
	s := "postgres://user:pass@host.com:5432/path?k=v#fl"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	pl(u.Scheme)
	pl(u.User)
	pl(u.User.Username())

	pass, _ := u.User.Password()
	pl(pass)

	pl(u.Host)

	host, port, _ := net.SplitHostPort(u.Host)
	pl(host)
	pl(port)

	pl(u.Path)

	pl(u.RawQuery)

	m, _ := url.ParseQuery(u.RawQuery)

	pl(m)
	pl(m["k"][0])

	pl(u.Fragment)
}
