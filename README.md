Duck DNS module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records for Duck DNS.

## Caddy module name

```
dns.providers.duckdns
```

## Caddyfile definition

```
duckdns [<api_token>] {
    api_token <api_token>
    override_domain <duckdns_domain>
}
```

- `api_token` may be specified as an argument to the `duckdns` directive, or in the body.
- `override_domain` is optional; see [Challenge delegation](#challenge-delegation) below.

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "duckdns",
				"api_token": "YOUR_DUCKDNS_API_TOKEN",
				"override_domain": "OPTIONAL_DUCKDNS_DOMAIN"
			}
		}
	}
}
```

or with the Caddyfile:

```
tls {
	dns duckdns {env.DUCKDNS_API_TOKEN}
}
```

You can replace `{env.DUCKDNS_API_TOKEN}` with the actual auth token if you prefer to put it directly in your config instead of an environment variable.


## Authenticating

See [the associated README in the libdns package](https://github.com/libdns/duckdns) for important information about credentials. Your token can be found at the top of the page when logged in at https://www.duckdns.org.

## Challenge delegation

To obtain a certificate using ACME DNS challenges, you'd use this module as described above. But, if you have a different domain (say, `my.example.com`) CNAME'd to your Duck DNS domain, you have two options:

1. Not use this module: Use a module matching the DNS provider for `my.example.com`.
2. [Delegate the challenge](https://letsencrypt.org/docs/challenge-types/#dns-01-challenge) to Duck DNS.

Delegating the challenge is actually quite easy, and may be useful if the DNS provider for `my.example.com` is difficult to configure or slow:

> Since Let’s Encrypt follows the DNS standards when looking up TXT records for DNS-01 validation, you can use CNAME records or NS records to delegate answering the challenge to other DNS zones. This can be used to delegate the _acme-challenge subdomain to a validation-specific server or zone. It can also be used if your DNS provider is slow to update, and you want to delegate to a quicker-updating server.

Let's say you have `my.example.com` as a CNAME to `example.duckdns.org`. Following the above instructions, you'd want to add a CNAME from `_acme-challenge.my.example.com` to `example.duckdns.org`. Then, in your Caddyfile, instruct this module to override the domain used when communicating with Duck DNS' API:

```
my.example.com {
	tls {
		dns duckdns <token> {
			override_domain example.duckdns.org
		}
	}
	...
}
```

If you do not set `override_domain`, this module will attempt to ask Duck DNS to update `my.example.com`. Duck DNS will return an error, because it does not know about `my.example.com`. Setting this value fixes this up and allows Caddy to ask for a certificate matching `my.example.com`, but by performing DNS challenges on `example.duckdns.org` instead!
