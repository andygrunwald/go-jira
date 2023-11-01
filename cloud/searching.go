package cloud

import "fmt"

type (
	searchParam struct {
		name  string
		value string
	}

	search []searchParam

	searchF func(search) search
)

// WithMaxResults sets the max results to return
func WithMaxResults(maxResults int) searchF {
	return func(s search) search {
		s = append(s, searchParam{name: "maxResults", value: fmt.Sprintf("%d", maxResults)})
		return s
	}
}

// WithAccountId sets the account id to search
func WithAccountId(accountId string) searchF {
	return func(s search) search {
		s = append(s, searchParam{name: "accountId", value: accountId})
		return s
	}
}

// WithUsername sets the username to search
func WithUsername(username string) searchF {
	return func(s search) search {
		s = append(s, searchParam{name: "username", value: username})
		return s
	}
}
