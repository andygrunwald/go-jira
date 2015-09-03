# go-jira

[![GoDoc](https://godoc.org/github.com/andygrunwald/go-jira?status.svg)](https://godoc.org/github.com/andygrunwald/go-jira)
[![Build Status](https://travis-ci.org/andygrunwald/go-jira.svg?branch=master)](https://travis-ci.org/andygrunwald/go-jira)
[![Coverage Status](https://coveralls.io/repos/andygrunwald/go-jira/badge.svg?branch=master&service=github)](https://coveralls.io/github/andygrunwald/go-jira?branch=master)

[Go](https://golang.org/) client library for [Atlassian JIRA](https://www.atlassian.com/software/jira).

![Go client library for Atlassian JIRA](./img/go-jira-compressed.png "Go client library for Atlassian JIRA.")

The code structure of this package was inspired by [google/go-github](https://github.com/google/go-github).

## Features

* Authentication (HTTP Basic, OAuth, Session Cookie)
* Create and receive issues
* Call every (not implemented) API endpoint of the JIRA

> Attention: This package is not JIRA API complete (yet), but you can call every API endpoint you want. See ["Call a not implemented API endpoint"](#call-a-not-implemented-api-endpoint) how to do this. For all possible API endpoints have a look at [latest JIRA REST API documentation](https://docs.atlassian.com/jira/REST/latest/).

## Installation

It is go gettable

    $ go get github.com/andygrunwald/go-jira

(optional) to run unit / example tests:

    $ cd $GOPATH/src/github.com/andygrunwald/go-jira
    $ go test -v

## API

Please have a look at the [GoDoc documentation](https://godoc.org/github.com/andygrunwald/go-jira) for a detailed API description.

The [latest JIRA REST API documentation](https://docs.atlassian.com/jira/REST/latest/) was the base document for this package.

## Examples

Further a few examples how the API can be used.
A few more examples are available in the [GoDoc examples section](https://godoc.org/github.com/andygrunwald/go-jira#pkg-examples).

### Get a single issue

Lets retrieve [MESOS-3325](https://issues.apache.org/jira/browse/MESOS-3325) from the [Apache Mesos](http://mesos.apache.org/) project.

```go
package main

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

func main() {
	jiraClient, _ := jira.NewClient(nil, "https://issues.apache.org/jira/")
	issue, _, _ := jiraClient.Issue.Get("MESOS-3325")

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
	fmt.Printf("Type: %s\n", issue.Fields.Type.Name)
	fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)

	// MESOS-3325: Running mesos-slave@0.23 in a container causes slave to be lost after a restart
	// Type: Bug
	// Priority: Critical
}
```

### Authenticate with session cookie

```go
package main

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

func main() {
	jiraClient, err := jira.NewClient(nil, "https://your.jira-instance.com/")
	if err != nil {
		panic(err)
	}

	res, err := jiraClient.Authentication.AcquireSessionCookie("username", "password")
	if err != nil || res == false {
		fmt.Printf("Result: %v\n", res)
		panic(err)
	}

	issue, _, err := jiraClient.Issue.Get("SYS-5156")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
}
```

### Call a not implemented API endpoint

Lets get all public projects of [Atlassian`s JIRA instance](https://jira.atlassian.com/).

```go
package main

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

func main() {
	jiraClient, _ := jira.NewClient(nil, "https://jira.atlassian.com/")
	req, _ := jiraClient.NewRequest("GET", "/rest/api/2/project", nil)

	projects := new([]jira.Project)
	_, err := jiraClient.Do(req, projects)
	if err != nil {
		panic(err)
	}

	for _, project := range *projects {
		fmt.Printf("%s: %s\n", project.Key, project.Name)
	}

	// ...
	// BAM: Bamboo
	// BAMJ: Bamboo JIRA Plugin
	// CLOV: Clover
	// CONF: Confluence
	// ...
}
```

## Implementations

TODO

## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).