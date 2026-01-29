# weheat-golang

Batteries-included Go client for the Weheat cloud API.

## Status
- Read-only endpoints (no control APIs discovered yet).
- Mirrors the public OpenAPI shapes from the upstream `weheat` Python client.
- Includes helper abstractions (discovery + `HeatPump` computed metrics).

You can do things like this, when paired with a home-assistant-like knockoff like my [GoHome](https://github.com/joshp123/gohome) tool:

<img width="2308" height="1831" alt="image" src="https://github.com/user-attachments/assets/02665151-8f7a-4a78-97b9-d8bec9b7f61c" />


## Installation
```sh
go get github.com/joshp123/weheat-golang
```

## Base URLs
- Default API: `https://api.weheat.nl`
- Legacy/alt: `https://api.weheat.nl/third_party` (if your account requires it)

## Wiring it up (OAuth)
Weheat uses Keycloak OAuth2. You need a **client_id** and **client_secret** from Weheat.

### 1) Exchange credentials for a refresh token (example)
```go
import (
  "context"
  "golang.org/x/oauth2"
  weheat "github.com/joshp123/weheat-golang"
)

ctx := context.Background()
conf := &oauth2.Config{
  ClientID:     clientID,
  ClientSecret: clientSecret,
  Endpoint: oauth2.Endpoint{
    TokenURL: weheat.DefaultTokenURL,
  },
  Scopes: weheat.DefaultScopes,
}

token, err := conf.PasswordCredentialsToken(ctx, username, password)
refreshToken := token.RefreshToken
```

Store the refresh token in your secret manager (e.g., agenix). If direct grant is disabled
for your account, you’ll need Weheat to provide an alternate flow.

### 2) Create a client with refresh-token auto-rotation
```go
source, _ := weheat.OAuthTokenSource(weheat.OAuthConfig{
  ClientID:     clientID,
  ClientSecret: clientSecret,
  RefreshToken: refreshToken,
})

client, _ := weheat.NewClient(
  weheat.WithTokenSource(source),
  // use WithBaseURL("https://api.weheat.nl/third_party") if needed
)
```

### 3) Or use a static access token
```go
client, _ := weheat.NewClient(
  weheat.WithTokenSource(weheat.StaticToken(accessToken)),
)
```

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
