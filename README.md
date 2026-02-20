# Goverland Core SDK

<a href="https://github.com/goverland-labs/goverland-core-sdk-go?tab=License-1-ov-file" rel="nofollow"><img src="https://img.shields.io/github/license/goverland-labs/goverland-core-sdk-go" alt="GPL 3.0" style="max-width:100%;"></a>
![unit-tests](https://github.com/goverland-labs/goverland-core-sdk-go/workflows/unit-tests/badge.svg)
![golangci-lint](https://github.com/goverland-labs/goverland-core-sdk-go/workflows/golangci-lint/badge.svg)

Go SDK client for the [Goverland Core Web API](https://github.com/goverland-labs/goverland-core-web-api). Provides typed methods for querying DAOs, proposals, votes, delegates, feed, ENS resolution, and more.

## Installation

```bash
go get github.com/goverland-labs/goverland-core-sdk-go
```

## Usage

```go
package main

import (
    "context"
    "fmt"

    coresdk "github.com/goverland-labs/goverland-core-sdk-go"
)

func main() {
    client := coresdk.NewClient("https://core.goverland.xyz/v1")

    // Get a DAO by ID or alias
    dao, err := client.GetDao(context.Background(), "aavedao.eth")
    if err != nil {
        panic(err)
    }
    fmt.Println(dao.Name)

    // List proposals with filters
    proposals, err := client.GetProposalList(context.Background(), coresdk.GetProposalListRequest{
        Dao:   dao.ID.String(),
        Limit: 10,
    })
    if err != nil {
        panic(err)
    }
    fmt.Printf("Found %d proposals\n", proposals.TotalCnt)
}
```

## API Reference

### Client

```go
// Create a new client with the Core Web API base URL.
// Optionally pass a custom *http.Client.
client := coresdk.NewClient("https://core.goverland.xyz/v1")
```

### DAOs

| Method | Description |
|---|---|
| `GetDao(ctx, id)` | Get a single DAO by UUID or Snapshot space ID (alias) |
| `GetDaoList(ctx, params)` | List DAOs with filtering by query, category, IDs, or fungible IDs |
| `GetDaoTop(ctx, params)` | Get top DAOs grouped by category |
| `GetDaoRecommendations(ctx)` | Get DAO recommendations |
| `GetDaoTokenInfo(ctx, daoID)` | Get token info for a DAO |
| `GetDaoTokenChart(ctx, daoID, period)` | Get token price chart for a DAO |

### Proposals

| Method | Description |
|---|---|
| `GetProposal(ctx, id)` | Get a single proposal by ID |
| `GetProposalList(ctx, params)` | List proposals with filtering by DAO, category, title, active status, and sorting |
| `GetProposalTop(ctx, params)` | Get top proposals |
| `GetProposalVotes(ctx, proposalID, params)` | Get votes for a proposal with search and ordering |
| `GetProposalVpList(ctx, proposalID)` | Get voting power distribution for a proposal |
| `ValidateVote(ctx, proposalID, params)` | Validate a vote before submission |
| `PrepareVote(ctx, proposalID, params)` | Prepare a vote for signing |
| `Vote(ctx, params)` | Submit a signed vote |

### Delegates

| Method | Description |
|---|---|
| `GetDelegates(ctx, daoID, params)` | List delegates for a DAO |
| `GetDelegatesV2(ctx, daoID, params)` | List delegates (V2 format with detailed delegation info) |
| `GetDelegateProfile(ctx, daoID, params)` | Get delegate profile for a specific address |
| `GetDelegators(ctx, daoID, params)` | List delegators for a delegate in a DAO |
| `GetUserDelegatesTopV2(ctx, address)` | Get top delegates for a user across all DAOs |
| `GetUserDelegatesListV2(ctx, daoID, address, params)` | List delegates for a user in a specific DAO |
| `GetUserDelegatorsTopV2(ctx, address)` | Get top delegators for a user across all DAOs |
| `GetUserDelegatorsV2(ctx, daoID, address, params)` | List delegators for a user in a specific DAO |
| `GetUserTopDelegatorsByDaoV2(ctx, address, daoID)` | Get top delegators for a user in a specific DAO |

### Votes & Users

| Method | Description |
|---|---|
| `GetUserVotes(ctx, address, params)` | Get votes by a user with optional DAO filter |
| `GetUserParticipatedDaos(ctx, address)` | Get DAOs where a user has voted |

### ENS

| Method | Description |
|---|---|
| `GetEnsNames(ctx, params)` | Resolve Ethereum addresses to ENS names |
| `GetAddressesByEnsNames(ctx, params)` | Resolve ENS names to Ethereum addresses |

### Feed

| Method | Description |
|---|---|
| `GetDaoFeed(ctx, daoID, params)` | Get activity feed for a DAO |
| `GetFeedByFilters(ctx, params)` | Get feed items with filters (DAOs, types, actions, active status) |

### Subscriptions

| Method | Description |
|---|---|
| `CreateSubscriber(ctx, webhookURL)` | Create a webhook subscriber |
| `UpdateSubscriber(ctx, subscriberID, webhookURL)` | Update subscriber webhook URL |
| `SubscribeOnDao(ctx, subscriberID, daoID)` | Subscribe to DAO events |
| `UnsubscribeFromDao(ctx, subscriberID, daoID)` | Unsubscribe from DAO events |

### Stats

| Method | Description |
|---|---|
| `GetStatsTotals(ctx)` | Get platform-wide statistics |

## Error Handling

The SDK returns typed errors for common HTTP status codes:

```go
import coresdk "github.com/goverland-labs/goverland-core-sdk-go"

dao, err := client.GetDao(ctx, "unknown")
if errors.Is(err, coresdk.ErrNotFound) {
    // 404 - resource not found
}
if errors.Is(err, coresdk.ErrUnauthorized) {
    // 401 - unauthorized
}
if errors.Is(err, coresdk.ErrForbidden) {
    // 403 - forbidden
}

// Validation errors (400)
var validationErr coresdk.ValidationError
if errors.As(err, &validationErr) {
    fmt.Println(validationErr.Error())
}

// Rate limiting (429) includes retry-after duration
var rateLimitErr coresdk.TooManyRequestsError
if errors.As(err, &rateLimitErr) {
    time.Sleep(rateLimitErr.RetryAfter)
}
```

## Examples

See the [examples/](examples/) directory for runnable code samples.

## Contribution Rules

[CONTRIBUTING.md](CONTRIBUTING.md)

## Changelog

[CHANGELOG.md](CHANGELOG.md)
