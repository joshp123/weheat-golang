# weheat-golang

Batteries-included Go client for the Weheat cloud API.

## Status
- Read-only endpoints (no control APIs discovered yet).
- Mirrors the public OpenAPI shapes from the upstream `weheat` Python client.

## Base URLs
- Default API: `https://api.weheat.nl`
- Legacy/alt: `https://api.weheat.nl/third_party` (if your account requires it)

## Auth
Weheat uses OAuth2 (Keycloak). This client accepts a token source so you can plug in
Gohome’s OAuth manager or a static access token.

## Quick start
```go
package main

import (
  "context"
  "fmt"

  weheat "github.com/joshp123/weheat-golang"
)

func main() {
  ctx := context.Background()
  client, _ := weheat.NewClient(
    weheat.WithTokenSource(weheat.StaticToken("ACCESS_TOKEN")),
  )

  pumps, _ := client.ListHeatPumps(ctx, weheat.ListHeatPumpsParams{
    State: weheat.DeviceStateActive,
  })

  fmt.Println(len(pumps.Data))
}
```

## Helpers
```go
ctx := context.Background()
client, _ := weheat.NewClient(
  weheat.WithTokenSource(weheat.StaticToken("ACCESS_TOKEN")),
)

pumps, _ := client.DiscoverActiveHeatPumps(ctx)
if len(pumps) > 0 {
  hp := weheat.NewHeatPump(client, pumps[0].ID)
  _ = hp.RefreshStatus(ctx, weheat.RequestOptions{})
  fmt.Println(hp.COP())
}
```

## License
MIT

## Supported endpoints
- `GET /api/v1/users/me`
- `GET /api/v1/heat-pumps`
- `GET /api/v1/heat-pumps/{heatPumpId}`
- `GET /api/v1/heat-pumps/{heatPumpId}/logs/latest`
- `GET /api/v1/heat-pumps/{heatPumpId}/logs/raw`
- `GET /api/v1/heat-pumps/{heatPumpId}/logs`
- `GET /api/v1/energy-logs/{heatPumpId}`
- `GET /api/v1/energy-logs/{heatPumpId}/total`
