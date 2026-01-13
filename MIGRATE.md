# Migration Guide: v1.x to v2.0.0

> **Guide version:** Written on 2026-01-12, based on commit [`351048e`](https://github.com/andygrunwald/go-jira/commit/351048edf7197c79f975a01ca4446ef3c7fabda0) (2026-01-01).

This guide helps you migrate your code from `go-jira` v1.x (including v1.17.0) to v2.0.0.

Version 2.0.0 introduces significant breaking changes to make the library more reliable, future-proof, and aligned with the diverging Jira Cloud and On-Premise APIs.
For the full context on the v2.0.0 goals, see [Issue #489](https://github.com/andygrunwald/go-jira/issues/489).

## Table of Contents

- [Quick Reference](#quick-reference)
- [Prerequisites](#prerequisites)
- [Which Package Should I Use?](#which-package-should-i-use)
- [Breaking Changes](#breaking-changes)
  - [1. Module and Import Path Changes](#1-module-and-import-path-changes)
  - [2. Client Initialization](#2-client-initialization)
  - [3. Context is Now Required](#3-context-is-now-required)
  - [4. Authentication Changes](#4-authentication-changes)
  - [5. Service Method Changes](#5-service-method-changes)
- [Complete Migration Examples](#complete-migration-examples)
- [New Features in v2.0.0](#new-features-in-v200)
- [Troubleshooting](#troubleshooting)
- [Resources](#resources)

---

## Quick Reference

Here's a checklist of the most common changes you'll need to make:

- [ ] Update `go.mod` to use `github.com/andygrunwald/go-jira/v2`
- [ ] Change import to `github.com/andygrunwald/go-jira/v2/cloud` or `github.com/andygrunwald/go-jira/v2/onpremise`
- [ ] Swap `NewClient` arguments: `NewClient(baseURL, httpClient)` instead of `NewClient(httpClient, baseURL)`
- [ ] Add `context.Context` as the first argument to all service method calls
- [ ] Remove `WithContext` suffix from method names (e.g., `GetWithContext` becomes `Get`)
- [ ] Update authentication transport field names (Cloud: `Password` -> `APIToken`)
- [ ] Update renamed methods (see [Service Method Changes](#5-service-method-changes))

---

## Prerequisites

### Go Version

v2.0.0 requires **Go 1.21 or later**.

### Determine Your Jira Deployment Type

Before migrating, you need to know whether you're connecting to:

- **Jira Cloud**: Hosted by Atlassian at `*.atlassian.net` domains
- **Jira On-Premise**: Self-hosted Jira Server or Jira Data Center

This determines which package you'll import.

---

## Which Package Should I Use?

| Your Jira Instance | Import Path |
|-------------------|-------------|
| Jira Cloud (`*.atlassian.net`) | `github.com/andygrunwald/go-jira/v2/cloud` |
| Jira Server (self-hosted) | `github.com/andygrunwald/go-jira/v2/onpremise` |
| Jira Data Center (self-hosted) | `github.com/andygrunwald/go-jira/v2/onpremise` |

**Why the split?** The Jira Cloud and On-Premise APIs have diverged significantly over time. Atlassian develops them independently with different endpoints, authentication methods, and data structures. Maintaining a single client led to confusion and bugs.
See [PR #503](https://github.com/andygrunwald/go-jira/pull/503) for the full rationale.

---

## Breaking Changes

### 1. Module and Import Path Changes

The module path now includes `/v2` following Go module versioning conventions.

**Update your `go.mod`:**

```diff
- require github.com/andygrunwald/go-jira v1.17.0
+ require github.com/andygrunwald/go-jira/v2 v2.0.0
```

**Update your imports:**

```go
// v1.x
import "github.com/andygrunwald/go-jira"

// v2.0.0 - Jira Cloud
import jira "github.com/andygrunwald/go-jira/v2/cloud"

// v2.0.0 - Jira On-Premise (Server/Data Center)
import jira "github.com/andygrunwald/go-jira/v2/onpremise"
```

> **Tip:** Using the alias `jira` allows you to keep most of your existing code unchanged (e.g., `jira.NewClient`, `jira.Issue`).

**Rationale:** [PR #503](https://github.com/andygrunwald/go-jira/pull/503) - The Cloud and On-Premise APIs differ significantly. Separate packages make API differences explicit and prevent using Cloud-only features with On-Premise instances (or vice versa).

---

### 2. Client Initialization

The argument order for `NewClient` has been reversed.

**Before (v1.x):**

```go
// httpClient first, then baseURL
client, err := jira.NewClient(httpClient, "https://jira.example.com/")

// With nil httpClient
client, err := jira.NewClient(nil, "https://jira.example.com/")
```

**After (v2.0.0):**

```go
// baseURL first, then httpClient
client, err := jira.NewClient("https://jira.example.com/", httpClient)

// With nil httpClient (uses default http.Client)
client, err := jira.NewClient("https://jira.example.com/", nil)
```

**Rationale:** [PR #509](https://github.com/andygrunwald/go-jira/pull/509) - The base URL is always required while the HTTP client is optional. The new order is more intuitive and follows the pattern of other popular Go API clients.

---

### 3. Context is Now Required

All API methods now require a `context.Context` as their first argument. The `*WithContext` suffix methods have been removed.

#### Request Methods

**Before (v1.x):**

```go
// Without context
req, err := client.NewRequest("GET", "rest/api/2/issue/KEY-123", nil)

// With context
req, err := client.NewRequestWithContext(ctx, "GET", "rest/api/2/issue/KEY-123", nil)
```

**After (v2.0.0):**

```go
// Context is always required
req, err := client.NewRequest(ctx, "GET", "rest/api/2/issue/KEY-123", nil)
```

The same applies to:
- `NewRawRequest` (was `NewRawRequestWithContext`)
- `NewMultiPartRequest` (was `NewMultiPartRequestWithContext`)

#### Service Methods

**Before (v1.x):**

```go
// Without context
issue, resp, err := client.Issue.Get("KEY-123", nil)

// With context
issue, resp, err := client.Issue.GetWithContext(ctx, "KEY-123", nil)
```

**After (v2.0.0):**

```go
// Context is always required
issue, resp, err := client.Issue.Get(ctx, "KEY-123", nil)
```

**Migration pattern:**

If you weren't using contexts before, add `context.Background()` or `context.TODO()`:

```go
// Quick migration
issue, resp, err := client.Issue.Get(context.Background(), "KEY-123", nil)

// Better: use a proper context with timeout
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
issue, resp, err := client.Issue.Get(ctx, "KEY-123", nil)
```

**Rationale:** [PR #513](https://github.com/andygrunwald/go-jira/pull/513) - Go 1.13+ natively supports context in `http.NewRequest`. Since the library no longer supports older Go versions, having both `Method()` and `MethodWithContext()` was redundant and cluttered the API.

---

### 4. Authentication Changes

Authentication has been simplified and now differs between Cloud and On-Premise.

#### Jira Cloud Authentication

Cloud authentication has been streamlined. Several transports were removed because they were duplicates or not applicable to Cloud.

**Removed transports:**
- `BearerAuthTransport` - Use `BasicAuthTransport` instead
- `PATAuthTransport` - Use `BasicAuthTransport` instead
- `CookieAuthTransport` - Not supported by Jira Cloud
- `AuthenticationService` - Removed from client

**Field rename:**
- `BasicAuthTransport.Password` is now `BasicAuthTransport.APIToken`

**Before (v1.x) - Using BearerAuthTransport:**

```go
import "github.com/andygrunwald/go-jira"

tp := jira.BearerAuthTransport{
    Token: "your-api-token",
}
client, err := jira.NewClient(tp.Client(), "https://example.atlassian.net/")
```

**Before (v1.x) - Using BasicAuthTransport with Password:**

```go
import "github.com/andygrunwald/go-jira"

tp := jira.BasicAuthTransport{
    Username: "user@example.com",
    Password: "your-api-token",
}
client, err := jira.NewClient(tp.Client(), "https://example.atlassian.net/")
```

**After (v2.0.0 Cloud):**

```go
import jira "github.com/andygrunwald/go-jira/v2/cloud"

tp := jira.BasicAuthTransport{
    Username: "user@example.com",
    APIToken: "your-api-token",  // Note: Field renamed from Password to APIToken
}
client, err := jira.NewClient("https://example.atlassian.net/", tp.Client())
```

> **Note:** For Jira Cloud, you must use an [API token](https://id.atlassian.com/manage-profile/security/api-tokens), not your account password.

**Rationale:** [PR #578](https://github.com/andygrunwald/go-jira/pull/578), [PR #579](https://github.com/andygrunwald/go-jira/pull/579) - Jira Cloud only supports API token authentication via Basic Auth. The other transports were either duplicates or not applicable to Cloud.

#### Jira On-Premise Authentication

On-Premise retains more authentication options to support various deployment configurations.

**Available transports:**
- `BasicAuthTransport` - Username/password authentication
- `BearerAuthTransport` - OAuth 2.0 / Personal Access Tokens (Jira 8.14+)
- `CookieAuthTransport` - Session-based authentication
- `PersonalAccessTokenAuthTransport` - PAT authentication
- `JWTAuthTransport` - JWT authentication for add-ons

**Before (v1.x):**

```go
import "github.com/andygrunwald/go-jira"

tp := jira.BasicAuthTransport{
    Username: "admin",
    Password: "secret",
}
client, err := jira.NewClient(tp.Client(), "https://jira.example.com/")
```

**After (v2.0.0 On-Premise):**

```go
import jira "github.com/andygrunwald/go-jira/v2/onpremise"

// Basic Auth (unchanged field names for On-Premise)
tp := jira.BasicAuthTransport{
    Username: "admin",
    Password: "secret",
}
client, err := jira.NewClient("https://jira.example.com/", tp.Client())
```

**Using Personal Access Tokens (Jira 8.14+):**

```go
import jira "github.com/andygrunwald/go-jira/v2/onpremise"

tp := jira.BearerAuthTransport{
    Token: "your-personal-access-token",
}
client, err := jira.NewClient("https://jira.example.com/", tp.Client())
```

---

### 5. Service Method Changes

Several methods have been renamed or consolidated. Methods with `*WithOptions` variants have been merged into a single method that accepts optional parameters.

#### BoardService

**`GetAllSprints` / `GetAllSprintsWithOptions` merged:**

```go
// v1.x
sprints, resp, err := client.Board.GetAllSprints("123")
sprints, resp, err := client.Board.GetAllSprintsWithOptions("123", opts)

// v2.0.0 - Pass nil for default options
sprints, resp, err := client.Board.GetAllSprints(ctx, 123, nil)
sprints, resp, err := client.Board.GetAllSprints(ctx, 123, opts)
```

#### GroupService

**`Get` / `GetWithOptions` merged:**

```go
// v1.x
members, resp, err := client.Group.Get("developers")
members, resp, err := client.Group.GetWithOptions("developers", opts)

// v2.0.0
members, resp, err := client.Group.Get(ctx, "developers", nil)
members, resp, err := client.Group.Get(ctx, "developers", opts)
```

**Cloud-only: `Add` renamed to `AddUserByGroupName`:**

```go
// v1.x (Cloud)
group, resp, err := client.Group.Add("groupname", "accountId")

// v2.0.0 (Cloud)
group, resp, err := client.Group.AddUserByGroupName(ctx, "groupname", "accountId")
```

**Cloud-only: `Remove` renamed to `RemoveUserByGroupName`:**

```go
// v1.x (Cloud)
resp, err := client.Group.Remove("groupname", "accountId")

// v2.0.0 (Cloud)
resp, err := client.Group.RemoveUserByGroupName(ctx, "groupname", "accountId")
```

#### IssueService

**`Update` / `UpdateWithOptions` merged:**

```go
// v1.x
issue, resp, err := client.Issue.Update(issue)
issue, resp, err := client.Issue.UpdateWithOptions(issue, opts)

// v2.0.0
issue, resp, err := client.Issue.Update(ctx, issue, nil)
issue, resp, err := client.Issue.Update(ctx, issue, opts)
```

**`GetCreateMeta` / `GetCreateMetaWithOptions` merged:**

```go
// v1.x
meta, resp, err := client.Issue.GetCreateMeta("PROJ")
meta, resp, err := client.Issue.GetCreateMetaWithOptions(opts)

// v2.0.0
meta, resp, err := client.Issue.GetCreateMeta(ctx, &jira.GetQueryOptions{
    ProjectKeys: "PROJ",
    Expand:      "projects.issuetypes.fields",
})
```

#### ProjectService

**`GetList` / `ListWithOptions` renamed to `GetAll`:**

```go
// v1.x
projects, resp, err := client.Project.GetList()
projects, resp, err := client.Project.ListWithOptions(opts)

// v2.0.0
projects, resp, err := client.Project.GetAll(ctx, nil)
projects, resp, err := client.Project.GetAll(ctx, opts)
```

#### UserService (Cloud-only)

**`GetSelf` renamed to `GetCurrentUser`:**

```go
// v1.x (Cloud)
user, resp, err := client.User.GetSelf()

// v2.0.0 (Cloud)
user, resp, err := client.User.GetCurrentUser(ctx)
```

#### ComponentService (Cloud-only)

**Type renamed: `CreateComponentOptions` -> `ComponentCreateOptions`:**

```go
// v1.x
opts := &jira.CreateComponentOptions{
    Name:    "Component Name",
    Project: "PROJ",
}

// v2.0.0
opts := &jira.ComponentCreateOptions{
    Name:    "Component Name",
    Project: "PROJ",
}
```

---

## Complete Migration Examples

### Example 1: Jira Cloud - Basic Setup

**Before (v1.x):**

```go
package main

import (
    "fmt"

    "github.com/andygrunwald/go-jira"
)

func main() {
    tp := jira.BasicAuthTransport{
        Username: "user@example.com",
        Password: "api-token",
    }

    client, err := jira.NewClient(tp.Client(), "https://example.atlassian.net/")
    if err != nil {
        panic(err)
    }

    user, _, err := client.User.GetSelf()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Logged in as: %s\n", user.DisplayName)
}
```

**After (v2.0.0):**

```go
package main

import (
    "context"
    "fmt"

    jira "github.com/andygrunwald/go-jira/v2/cloud"
)

func main() {
    tp := jira.BasicAuthTransport{
        Username: "user@example.com",
        APIToken: "api-token",  // Changed from Password
    }

    client, err := jira.NewClient("https://example.atlassian.net/", tp.Client())  // Arguments swapped
    if err != nil {
        panic(err)
    }

    user, _, err := client.User.GetCurrentUser(context.Background())  // Renamed + context added
    if err != nil {
        panic(err)
    }

    fmt.Printf("Logged in as: %s\n", user.DisplayName)
}
```

### Example 2: Jira On-Premise - Basic Setup

**Before (v1.x):**

```go
package main

import (
    "fmt"

    "github.com/andygrunwald/go-jira"
)

func main() {
    tp := jira.BasicAuthTransport{
        Username: "admin",
        Password: "secret",
    }

    client, err := jira.NewClient(tp.Client(), "https://jira.example.com/")
    if err != nil {
        panic(err)
    }

    user, _, err := client.User.GetSelf()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Logged in as: %s\n", user.DisplayName)
}
```

**After (v2.0.0):**

```go
package main

import (
    "context"
    "fmt"

    jira "github.com/andygrunwald/go-jira/v2/onpremise"
)

func main() {
    tp := jira.BasicAuthTransport{
        Username: "admin",
        Password: "secret",  // On-Premise keeps Password field
    }

    client, err := jira.NewClient("https://jira.example.com/", tp.Client())  // Arguments swapped
    if err != nil {
        panic(err)
    }

    user, _, err := client.User.GetSelf(context.Background())  // Context added
    if err != nil {
        panic(err)
    }

    fmt.Printf("Logged in as: %s\n", user.DisplayName)
}
```

### Example 3: Fetching an Issue

**Before (v1.x):**

```go
issue, resp, err := client.Issue.Get("PROJ-123", nil)
if err != nil {
    panic(err)
}
fmt.Printf("Issue: %s - %s\n", issue.Key, issue.Fields.Summary)
```

**After (v2.0.0):**

```go
issue, resp, err := client.Issue.Get(context.Background(), "PROJ-123", nil)
if err != nil {
    panic(err)
}
fmt.Printf("Issue: %s - %s\n", issue.Key, issue.Fields.Summary)
```

### Example 4: Creating an Issue

**Before (v1.x):**

```go
i := jira.Issue{
    Fields: &jira.IssueFields{
        Project: jira.Project{
            Key: "PROJ",
        },
        Type: jira.IssueType{
            Name: "Bug",
        },
        Summary:     "Something is broken",
        Description: "A detailed description of the bug.",
    },
}

issue, resp, err := client.Issue.Create(&i)
if err != nil {
    panic(err)
}
fmt.Printf("Created issue: %s\n", issue.Key)
```

**After (v2.0.0):**

```go
i := jira.Issue{
    Fields: &jira.IssueFields{
        Project: jira.Project{
            Key: "PROJ",
        },
        Type: jira.IssueType{
            Name: "Bug",
        },
        Summary:     "Something is broken",
        Description: "A detailed description of the bug.",
    },
}

issue, resp, err := client.Issue.Create(context.Background(), &i)  // Context added
if err != nil {
    panic(err)
}
fmt.Printf("Created issue: %s\n", issue.Key)
```

### Example 5: Searching with JQL

**Before (v1.x):**

```go
jql := "project = PROJ AND status = Open"
issues, resp, err := client.Issue.Search(jql, nil)
if err != nil {
    panic(err)
}

for _, issue := range issues {
    fmt.Printf("%s: %s\n", issue.Key, issue.Fields.Summary)
}
```

**After (v2.0.0):**

```go
jql := "project = PROJ AND status = Open"
issues, resp, err := client.Issue.Search(context.Background(), jql, nil)  // Context added
if err != nil {
    panic(err)
}

for _, issue := range issues {
    fmt.Printf("%s: %s\n", issue.Key, issue.Fields.Summary)
}
```

---

## New Features in v2.0.0

### User-Agent Header

The client now identifies itself with a User-Agent header: `go-jira/2.0.0`. This helps Jira administrators identify API traffic from this library.

You can customize the User-Agent if needed:

```go
client, err := jira.NewClient("https://example.atlassian.net/", nil)
client.UserAgent = "my-app/1.0 go-jira/2.0.0"
```

### Retrieve Underlying HTTP Client

You can now retrieve a copy of the underlying HTTP client:

```go
httpClient := client.Client()
```

This is useful for debugging or when you need to inspect the client configuration.

### Jira Cloud API v3 Support

The Cloud package now officially targets Jira Cloud REST API v3. This provides better compatibility with Atlassian's latest API features.

---

## Troubleshooting

### Common Compilation Errors

**Error: `too many arguments in call to jira.NewClient`**

You're using the v1.x argument order. Swap the arguments:

```go
// Wrong (v1.x order)
client, err := jira.NewClient(httpClient, baseURL)

// Correct (v2.0.0 order)
client, err := jira.NewClient(baseURL, httpClient)
```

**Error: `undefined: jira.BearerAuthTransport` (Cloud)**

`BearerAuthTransport` was removed from the Cloud package. Use `BasicAuthTransport` instead:

```go
tp := jira.BasicAuthTransport{
    Username: "user@example.com",
    APIToken: "your-api-token",
}
```

**Error: `tp.Password undefined (type BasicAuthTransport has no field or method Password)` (Cloud)**

The `Password` field was renamed to `APIToken` in the Cloud package:

```go
// Wrong
tp := jira.BasicAuthTransport{
    Username: "user@example.com",
    Password: "token",  // Field doesn't exist in Cloud
}

// Correct
tp := jira.BasicAuthTransport{
    Username: "user@example.com",
    APIToken: "token",
}
```

**Error: `not enough arguments in call to client.Issue.Get`**

All service methods now require a context as the first argument:

```go
// Wrong
issue, _, err := client.Issue.Get("KEY-123", nil)

// Correct
issue, _, err := client.Issue.Get(context.Background(), "KEY-123", nil)
```

**Error: `client.User.GetSelf undefined`**

On Cloud, `GetSelf` was renamed to `GetCurrentUser`:

```go
// Wrong (Cloud v2.0.0)
user, _, err := client.User.GetSelf()

// Correct (Cloud v2.0.0)
user, _, err := client.User.GetCurrentUser(ctx)
```

### FAQ

**Q: Can I use the same code for both Cloud and On-Premise?**

No. The Cloud and On-Premise packages are separate and have different APIs. You need to choose the appropriate package based on your Jira deployment. If you need to support both, you'll need conditional compilation or separate code paths.

**Q: What if I'm not sure whether my Jira is Cloud or On-Premise?**

- If your Jira URL ends with `.atlassian.net`, it's Jira Cloud
- If your Jira is hosted on your own domain or servers, it's On-Premise

**Q: Why was context made mandatory?**

Context allows you to set timeouts, cancellation, and pass request-scoped values. It's a Go best practice for any operation that may block or take significant time. Making it mandatory ensures all API calls can be properly controlled and cancelled.

**Q: Do I need to update all my code at once?**

Yes, v2.0.0 is a breaking change. You cannot incrementally migrate. However, using find-and-replace and the compiler errors as a guide makes the migration straightforward.

---

## Resources

### Examples

The repository includes working examples for both Cloud and On-Premise:

- **Cloud examples:** [`cloud/examples/`](https://github.com/andygrunwald/go-jira/tree/main/cloud/examples)
- **On-Premise examples:** [`onpremise/examples/`](https://github.com/andygrunwald/go-jira/tree/main/onpremise/examples)

### Related Pull Requests

- [PR #503](https://github.com/andygrunwald/go-jira/pull/503) - Split On-Premise and Cloud Clients
- [PR #509](https://github.com/andygrunwald/go-jira/pull/509) - Client changes for future-proofing
- [PR #513](https://github.com/andygrunwald/go-jira/pull/513) - Make context a first-class citizen
- [PR #578](https://github.com/andygrunwald/go-jira/pull/578) - Cloud authentication cleanup
- [PR #579](https://github.com/andygrunwald/go-jira/pull/579) - Cloud authentication cleanup (related)

### Documentation

- [CHANGELOG.md](https://github.com/andygrunwald/go-jira/blob/main/CHANGELOG.md) - Full changelog
- [Jira Cloud REST API v3](https://developer.atlassian.com/cloud/jira/platform/rest/v3/intro/) - Atlassian's API documentation
- [Jira Server REST API](https://docs.atlassian.com/software/jira/docs/api/REST/latest/) - On-Premise API documentation

---

If you encounter any issues during migration, please [open an issue](https://github.com/andygrunwald/go-jira/issues) on GitHub.
