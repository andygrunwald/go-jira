# go-jira

[![GoDoc](https://pkg.go.dev/badge/github.com/andygrunwald/go-jira?utm_source=godoc)](https://pkg.go.dev/github.com/andygrunwald/go-jira)
[![Build Status](https://github.com/andygrunwald/go-jira/actions/workflows/testing.yml/badge.svg)](https://github.com/andygrunwald/go-jira/actions/workflows/testing.yml)
[![Go Report Card](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=flat)](https://goreportcard.com/report/github.com/andygrunwald/go-jira)

[Go](https://go.dev/) client library for [Atlassian Jira](https://www.atlassian.com/software/jira).

![Go client library for Atlassian Jira](./img/logo_small.png "Go client library for Atlassian Jira.")

## :warning: State of this library :warning:

**v2 of this library is in development.**
**v2 will contain breaking changes :warning:**
**The current main branch contains the development version of v2.**

The goals of v2 are:

* idiomatic go usage
* proper documentation
* being compliant with different kinds of Atlassian Jira products (on-premise vs. cloud)
* remove flaws introduced during the early times of this library

See our milestone [Road to v2](https://github.com/andygrunwald/go-jira/milestone/1) and provide feedback in [Development is kicking: Road to v2 🚀 #489](https://github.com/andygrunwald/go-jira/issues/489).
Attention: The current `main` branch represents the v2 development version - we treat this version as unstable and breaking changes are expected.

**If you want to stay more stable, please use v1.\*** - See our [releases](https://github.com/andygrunwald/go-jira/releases).
Latest stable release: [v1.16.0](https://github.com/andygrunwald/go-jira/releases/tag/v1.16.0)

## Features

* Authentication (HTTP Basic, OAuth, Session Cookie, Bearer (for PATs))
* Create and retrieve issues
* Create and retrieve issue transitions (status updates)
* Call every API endpoint of the Jira, even if it is not directly implemented in this library

This package is not Jira API complete (yet), but you can call every API endpoint you want. See [Call a not implemented API endpoint](#call-a-not-implemented-api-endpoint) how to do this. For all possible API endpoints of Jira have a look at [latest Jira REST API documentation](https://developer.atlassian.com/cloud/jira/platform/rest/v3/intro/).

## Requirements

* Go >= 1.14
* Jira v6.3.4 & v7.1.2.

Note that we also run our tests against 1.13, though only the last two versions
of Go are officially supported.

## Installation

It is go gettable

```sh
go get github.com/andygrunwald/go-jira
```

## API

Please have a look at the [GoDoc documentation](https://pkg.go.dev/github.com/andygrunwald/go-jira) for a detailed API description.

The [latest Jira REST API documentation](https://developer.atlassian.com/cloud/jira/platform/rest/v3/intro/) was the base document for this package.

## Examples

Further a few examples how the API can be used.
A few more examples are available in the [GoDoc examples section](https://pkg.go.dev/github.com/andygrunwald/go-jira#section-directories).

### Get a single issue

Lets retrieve [MESOS-3325](https://issues.apache.org/jira/browse/MESOS-3325) from the [Apache Mesos](http://mesos.apache.org/) project.

```go
package main

import (
	"fmt"
	jira "github.com/andygrunwald/go-jira"
)

func main() {
	jiraClient, _ := jira.NewClient(nil, "https://issues.apache.org/jira/")
	issue, _, _ := jiraClient.Issue.Get("MESOS-3325", nil)

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
	fmt.Printf("Type: %s\n", issue.Fields.Type.Name)
	fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)

	// MESOS-3325: Running mesos-slave@0.23 in a container causes slave to be lost after a restart
	// Type: Bug
	// Priority: Critical
}
```

### Authentication

The `go-jira` library does not handle most authentication directly.  Instead, authentication should be handled within
an `http.Client`. That client can then be passed into the `NewClient` function when creating a jira client.

For convenience, capability for basic and cookie-based authentication is included in the main library.

#### Token (Jira on Atlassian Cloud)

Token-based authentication uses the basic authentication scheme, with a user-generated API token in place of a user's password. You can generate a token for your user [here](https://id.atlassian.com/manage-profile/security/api-tokens). Additional information about Atlassian Cloud API tokens can be found [here](https://support.atlassian.com/atlassian-account/docs/manage-api-tokens-for-your-atlassian-account/).

A more thorough, [runnable example](cloud/examples/basic_auth/main.go) is provided in the examples directory.

```go
func main() {
	tp := jira.BasicAuthTransport{
		Username: "<username>",
		Password: "<api-token>",
	}

	client, err := jira.NewClient(tp.Client(), "https://my.jira.com")

	u, _, err = client.User.GetCurrentUser(context.Background())

	fmt.Printf("Email: %v\n", u.EmailAddress)
	fmt.Println("Success!")
}
```

#### Bearer - Personal Access Tokens (self-hosted Jira)

For **self-hosted Jira** (v8.14 and later), Personal Access Tokens (PATs) were introduced.
Similar to the API tokens, PATs are a safe alternative to using username and password for authentication with scripts and integrations.
PATs use the Bearer authentication scheme.
Read more about Jira PATs [here](https://confluence.atlassian.com/enterprise/using-personal-access-tokens-1026032365.html).

See [examples/bearerauth](onpremise/examples/bearerauth/main.go) for how to use the Bearer authentication scheme with Jira in Go.

#### Basic (self-hosted Jira)

Password-based API authentication works for self-hosted Jira **only**, and has been [deprecated for users of Atlassian Cloud](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-basic-auth-and-cookie-based-auth/).

Depending on your version of Jira, either of the above token authentication examples may be used, substituting a user's password for a generated token.

#### Authenticate with OAuth

If you want to connect via OAuth to your Jira Cloud instance checkout the [example of using OAuth authentication with Jira in Go](https://gist.github.com/Lupus/edafe9a7c5c6b13407293d795442fe67) by [@Lupus](https://github.com/Lupus).

For more details have a look at the [issue #56](https://github.com/andygrunwald/go-jira/issues/56).

### Create an issue

Example how to create an issue.

```go
package main

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

func main() {
	base := "https://my.jira.com"
	tp := jira.BasicAuthTransport{
		Username: "username",
		Password: "token",
	}

	jiraClient, err := jira.NewClient(tp.Client(), base)
	if err != nil {
		panic(err)
	}

	i := jira.Issue{
		Fields: &jira.IssueFields{
			Assignee: &jira.User{
				Name: "myuser",
			},
			Reporter: &jira.User{
				Name: "youruser",
			},
			Description: "Test Issue",
			Type: jira.IssueType{
				Name: "Bug",
			},
			Project: jira.Project{
				Key: "PROJ1",
			},
			Summary: "Just a demo issue",
		},
	}
	issue, _, err := jiraClient.Issue.Create(&i)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
}
```

### Change an issue status

This is how one can change an issue status. In this example, we change the issue from "To Do" to "In Progress."

```go
package main

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

func main() {
        testIssueID := "FART-1"
	base := "https://my.jira.com"
	tp := jira.BasicAuthTransport{
		Username: "username",
		Password: "token",
	}

	jiraClient, err := jira.NewClient(tp.Client(), base)
	if err != nil {
		panic(err)
	}

	issue, _, _ := jiraClient.Issue.Get(testIssueID, nil)
	currentStatus := issue.Fields.Status.Name
	fmt.Printf("Current status: %s\n", currentStatus)

	var transitionID string
	possibleTransitions, _, _ := jiraClient.Issue.GetTransitions(testIssueID)
	for _, v := range possibleTransitions {
		if v.Name == "In Progress" {
			transitionID = v.ID
			break
		}
	}

	jiraClient.Issue.DoTransition(testIssueID, transitionID)
	issue, _, _ = jiraClient.Issue.Get(testIssueID, nil)
	fmt.Printf("Status after transition: %+v\n", issue.Fields.Status.Name)
}
```

### Get all the issues for JQL with Pagination

Jira API has limit on maxResults it can return. You may have a usecase where you need to get all issues for given JQL.
This example shows reference implementation of GetAllIssues function which does pagination on Jira API to get all the issues for given JQL.

Please look at [Pagination Example](https://github.com/andygrunwald/go-jira/blob/main/cloud/examples/pagination/main.go)

### Call a not implemented API endpoint

Not all API endpoints of the Jira API are implemented into *go-jira*.
But you can call them anyway:
Lets get all public projects of [Atlassian`s Jira instance](https://jira.atlassian.com/).

```go
package main

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

func main() {
	base := "https://my.jira.com"
	tp := jira.BasicAuthTransport{
		Username: "username",
		Password: "token",
	}

	jiraClient, err := jira.NewClient(tp.Client(), base)
	req, _ := jiraClient.NewRequest("GET", "rest/api/2/project", nil)

	projects := new([]jira.Project)
	_, err = jiraClient.Do(req, projects)
	if err != nil {
		panic(err)
	}

	for _, project := range *projects {
		fmt.Printf("%s: %s\n", project.Key, project.Name)
	}

	// ...
	// BAM: Bamboo
	// BAMJ: Bamboo Jira Plugin
	// CLOV: Clover
	// CONF: Confluence
	// ...
}
```

## Implementations

* [andygrunwald/jitic](https://github.com/andygrunwald/jitic) - The Jira Ticket Checker

## Development

### Code structure

The code structure of this package was inspired by [google/go-github](https://github.com/google/go-github).

There is one main part (the client).
Based on this main client the other endpoints, like Issues or Authentication are extracted in services. E.g. `IssueService` or `AuthenticationService`.
These services own a responsibility of the single endpoints / usecases of Jira.

### Unit testing

To run the local unit tests, execute

```sh
$ make test
```

To run the local unit tests and view the unit test code coverage in your local web browser, execute

```sh
$ make test-coverage-html
```

## Contribution

We ❤️ PR's

Contribution, in any kind of way, is highly welcome!
It doesn't matter if you are not able to write code.
Creating issues or holding talks and help other people to use [go-jira](https://github.com/andygrunwald/go-jira) is contribution, too!
A few examples:

* Correct typos in the README / documentation
* Reporting bugs
* Implement a new feature or endpoint
* Sharing the love of [go-jira](https://github.com/andygrunwald/go-jira) and help people to get use to it

If you are new to pull requests, checkout [Collaborating on projects using issues and pull requests / Creating a pull request](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request).

### Supported Go versions

We follow the [Go Release Policy](https://go.dev/doc/devel/release#policy):

> Each major Go release is supported until there are two newer major releases. For example, Go 1.5 was supported until the Go 1.7 release, and Go 1.6 was supported until the Go 1.8 release. We fix critical problems, including [critical security problems](https://go.dev/security/), in supported releases as needed by issuing minor revisions (for example, Go 1.6.1, Go 1.6.2, and so on).

### Supported Jira versions

#### Jira Server (On-Premise solution)

We follow the [Atlassian Support End of Life Policy](https://confluence.atlassian.com/support/atlassian-support-end-of-life-policy-201851003.html):

> Atlassian supports feature versions for two years after the first major iteration of that version was released (for example, we support Jira Core 7.2.x for 2 years after Jira 7.2.0 was released).

#### Jira Cloud

We support Jira Cloud API in [version 3](https://developer.atlassian.com/cloud/jira/platform/rest/v3/intro/).
Even if this API version is _currently_ in beta (by Atlassian):

[Version 2](https://developer.atlassian.com/cloud/jira/platform/rest/v2/) and version 3 of the API offer the same collection of operations.
However, version 3 provides support for the [Atlassian Document Format (ADF)](https://developer.atlassian.com/cloud/jira/platform/apis/document/structure/) in a subset of the API.

### Official Jira API documentation

* [Jira Server (On-Premise solution)](https://developer.atlassian.com/server/jira/platform/rest-apis/)
* Jira Cloud API in [version 2](https://developer.atlassian.com/cloud/jira/platform/rest/v2/intro/)
* Jira Cloud API in [version 3](https://developer.atlassian.com/cloud/jira/platform/rest/v3/intro/)

### Sandbox environment for testing

Jira offers sandbox test environments at http://go.atlassian.com/cloud-dev.

You can read more about them at https://blog.developer.atlassian.com/cloud-ecosystem-dev-env/.

## Releasing

Install [standard-version](https://github.com/conventional-changelog/standard-version)

```bash
npm i -g standard-version
```

```bash
standard-version
git push --tags
```

Manually copy/paste text from changelog (for this new version) into the release on Github.com. E.g.

[https://github.com/andygrunwald/go-jira/releases/edit/v1.11.0](https://github.com/andygrunwald/go-jira/releases/edit/v1.11.0)

## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).
