package duckdns

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/duckdns"
)

// Provider wraps the provider implementation as a Caddy module.
type Provider struct{ *duckdns.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.duckdns",
		New: func() caddy.Module { return &Provider{new(duckdns.Provider)} },
	}
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// duckdns [<api_token>] {
//     api_token <api_token>
//     override_domain <duckdns_domain>
// }
//
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	repl := caddy.NewReplacer()
	for d.Next() {
		if d.NextArg() {
			p.Provider.APIToken = repl.ReplaceAll(d.Val(), "")
		}
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_token":
				if !d.NextArg() {
					return d.ArgErr()
				}
				if p.Provider.APIToken != "" {
					return d.Err("API token already set")
				}
				p.Provider.APIToken = repl.ReplaceAll(d.Val(), "")
				if d.NextArg() {
					return d.ArgErr()
				}
			case "override_domain":
				if !d.NextArg() {
					return d.ArgErr()
				}
				if p.Provider.OverrideDomain != "" {
					return d.Err("Override domain already set")
				}
				p.Provider.OverrideDomain = repl.ReplaceAll(d.Val(), "")
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.APIToken == "" {
		return d.Err("missing API token")
	}
	return nil
}

// Interface guard
var _ caddyfile.Unmarshaler = (*Provider)(nil)
