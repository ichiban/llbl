# llbl

"Let Localhost be Localhost" DNS server which resolves `*.localhost` to `127.0.0.1`.

It relies on `/etc/resolver`.

## Installation

```bash
go get github.com/ichiban/llbl/cmd/llbl
```

## Usage

By default, `foo.localhost` doesn't resolve to anything.

```console
$ ping -c 1 foo.localhost
ping: cannot resolve foo.localhost: Unknown host
```

Run `llbl` with `sudo`. While it's running, any domains that end with `.localhost` will be resolved as `127.0.0.1`.

```console
$ sudo llbl
Password:
INFO[0000] listen                                        addr="0.0.0.0:62631"
```

```console
$ ping -c 1 foo.localhost
PING foo.localhost (127.0.0.1): 56 data bytes
64 bytes from 127.0.0.1: icmp_seq=0 ttl=64 time=0.048 ms

--- foo.localhost ping statistics ---
1 packets transmitted, 1 packets received, 0.0% packet loss
round-trip min/avg/max/stddev = 0.048/0.048/0.048/0.000 ms
```

## Back story

I wanted to test a web app which can be accessed as `*.example.com` where `*` can be any strings.
So I looked it up how to configure /etc/hosts to resolve any `*.example.com` to localhost.

I stumbled upon [a stackoverflow answer](https://stackoverflow.com/a/20446931) that says there's no way to configure /etc/hosts that way. So, instead [I needed to run a local DNS server](https://passingcuriosity.com/2013/dnsmasq-dev-osx/).
For that purpose, [the domain has to be `*.test` but there's a proposal to use `*.localhost`](https://ma.ttias.be/chrome-force-dev-domains-https-via-preloaded-hsts/).
I checked if `foo.localhost` works on my laptop. It didn't. So I made this.

## License

Distributed under the MIT license. See ``LICENSE`` for more information.

## Contributing

1. Fork it (<https://github.com/ichiban/llbl/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request