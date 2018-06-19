# [simple-http-proxy](https://github.com/fnkr/simple-http-proxy)

Simple http proxy, written in Go.
Connections are limited to specified users/domains.

## Usage example

```
$ simple-http-proxy -bind '[::1]:8080' -user foo:bar -host api.github.com:443 -host-match '.*\.googleapis\.com:443'
$ curl -x 'http://foo:bar@[::1]:8080' https://api.github.com/emojis
```

Run `simple-http-proxy -help` for help.
