# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

### [1.11.1](https://github.com/andygrunwald/go-jira/compare/v1.11.0...v1.11.1) (2019-10-17)

## [1.11.0](https://github.com/andygrunwald/go-jira/compare/v1.10.0...v1.11.0) (2019-10-17)


### Features

* Add AccountID and AccountType to GroupMember struct ([216e005](https://github.com/andygrunwald/go-jira/commit/216e0056d6385eba9d31cb37e6ff64314860d2cc))
* Add AccountType and Locale to User struct ([52ab347](https://github.com/andygrunwald/go-jira/commit/52ab34790307144087f0d9bf86c93a2b2209fe46))
* Add GetAllStatuses ([afc96b1](https://github.com/andygrunwald/go-jira/commit/afc96b18d17b77e32cec9e1ac7e4f5dec7e627f5))
* Add GetMyFilters to FilterService ([ebae19d](https://github.com/andygrunwald/go-jira/commit/ebae19dda6afd0e54578f30300bc36012381e99b))
* Add Search to FilterService ([38a755b](https://github.com/andygrunwald/go-jira/commit/38a755b407cd70d11fe2e2897d814552ca29ab51))
* add support for JWT auth with qsh needed by add-ons ([a8bdfed](https://github.com/andygrunwald/go-jira/commit/a8bdfed27ff42a9bb0468b8cf192871780919def))
* AddGetBoardConfiguration ([fd698c5](https://github.com/andygrunwald/go-jira/commit/fd698c57163f248f21285d5ebc6a3bb60d46694f))
* Replace http.Client with interface for extensibility ([b59a65c](https://github.com/andygrunwald/go-jira/commit/b59a65c365dcefd42e135579e9b7ce9c9c006489))


### Bug Fixes

* Fix fixversion description tag ([8383e2f](https://github.com/andygrunwald/go-jira/commit/8383e2f5f145d04f6bcdb47fb12a95b58bdcedfa))
* Fix typos in filter_test.go ([e9a261c](https://github.com/andygrunwald/go-jira/commit/e9a261c52249073345e5895b22e2cf4d7286497a))

# [1.10.0](https://github.com/andygrunwald/go-jira/compare/v1.9.0...v1.10.0) (2019-05-23)


### Bug Fixes

* empty SearchOptions causing malformed request ([b3bf8c2](https://github.com/andygrunwald/go-jira/commit/b3bf8c2))


### Features

* added DeleteAttachment ([e93c0e1](https://github.com/andygrunwald/go-jira/commit/e93c0e1))



# [1.9.0](https://github.com/andygrunwald/go-jira/compare/v1.8.0...v1.9.0) (2019-05-19)


### Features

* **issues:** Added support for AddWorklog and GetWorklogs ([1ebd7e7](https://github.com/andygrunwald/go-jira/commit/1ebd7e7))



# [1.8.0](https://github.com/andygrunwald/go-jira/compare/v1.7.0...v1.8.0) (2019-05-16)


### Bug Fixes

* Add PriorityService to the main ([8491cb0](https://github.com/andygrunwald/go-jira/commit/8491cb0))


### Features

* **filter:** Add GetFavouriteList to FilterService. ([645898e](https://github.com/andygrunwald/go-jira/commit/645898e))
* Add get all priorities ([1c63e25](https://github.com/andygrunwald/go-jira/commit/1c63e25))
* Add ResolutionService to retrieve resolutions ([fb1ce22](https://github.com/andygrunwald/go-jira/commit/fb1ce22))
* Add status category constants ([6223ddd](https://github.com/andygrunwald/go-jira/commit/6223ddd))
* Add StatusCategory GetList ([049a756](https://github.com/andygrunwald/go-jira/commit/049a756))



