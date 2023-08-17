package caddy_cron

import (
	"fmt"
	"net/http"
	"time"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(CronJob{})
	httpcaddyfile.RegisterHandlerDirective("cronjob", parseCaddyfile)
}

type CronJob struct {
	Schedule string `json:"schedule,omitempty"`
	Command  string `json:"command,omitempty"`
}

// CaddyModule returns the Caddy module information.
func (CronJob) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID: "http.handlers.cronjob",
		New: func() caddy.Module {
			return new(CronJob)
		},
	}
}

func (m *CronJob) Provision(ctx caddy.Context) error {
	fmt.Println("CronJob provisioned:", m.Schedule, m.Command)
	return nil
}

func (m CronJob) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	fmt.Println("Executing cron job:", m.Schedule, m.Command)

	// Simulate the cron job execution for demonstration purposes.
	ticker := time.NewTicker(parseDuration(m.Schedule))
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Running cron job:", m.Command)
		}
	}

	return next.ServeHTTP(w, r)
}

// UnmarshalCaddyfile implements caddyfile.Unmarshaler.
func (m *CronJob) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if !d.Args(&m.Command) {
			return d.ArgErr()
		}
	}
	return nil
}

// parseCaddyfile unmarshals tokens from h into a new Middleware.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m CronJob
	err := m.UnmarshalCaddyfile(h.Dispenser)
	return m, err
}

func parseDuration(durationStr string) time.Duration {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0
	}
	return duration
}

// Interface guards
var (
	_ caddy.Provisioner           = (*CronJob)(nil)
	_ caddyhttp.MiddlewareHandler = (*CronJob)(nil)
	_ caddyfile.Unmarshaler       = (*CronJob)(nil)
)
