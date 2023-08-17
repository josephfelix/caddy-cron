package caddy_cron

import (
	"testing"

	"github.com/caddyserver/caddy/v2/caddytest"
)

func TestMinimal(t *testing.T) {
	tester := caddytest.NewTester(t)
	tester.InitServer(`
	{
		admin localhost:2999
		skip_install_trust
		http_port     9080
		https_port    9443
	}
	localhost:9080 {
		route /test {
			cronjob @every 1s echo "Hello, world!"
			respond "Hello, default!"
		}
	}`, "caddyfile")

	tester.AssertGetResponse(`http://localhost:9080/test`, 200, "Hello, default!")
}
