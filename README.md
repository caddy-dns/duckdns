Duck DNS module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records for Duck DNS.

## Caddy module name

```
dns.providers.duckdns
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "duckdns",
				"api_token": "YOUR_DUCKDNS_API_TOKEN"
			}
		}
	}
}
```

or with the Caddyfile:

```
tls {
	dns cloudflare {env.DUCKDNS_API_TOKEN}
}
```

You can replace `{env.DUCKDNS_API_TOKEN}` with the actual auth token if you prefer to put it directly in your config instead of an environment variable.


## Authenticating

See [the associated README in the libdns package](https://github.com/libdns/duckdns) for important information about credentials. Your token can be found at the top of the page when logged in at https://www.duckdns.org.
