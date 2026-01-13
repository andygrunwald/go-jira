# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

## [2.0]() (UNRELEASED)

Version 2.0 is a bigger change with the main goal to make this library more reliable and future safe.
See https://github.com/andygrunwald/go-jira/issues/489 for details.

For detailed migration instructions from v1.x to v2.0, see [MIGRATE.md](MIGRATE.md).

### Breaking changes

* Jira On-Premise and Jira Cloud have now different clients, because the API differs
* `client.NewRawRequestWithContext()` has been removed in favor of `client.NewRawRequest()`, which requires now a context as first argument
* `client.NewRequestWithContext()` has been removed in favor of `client.NewRequest()`, which requires now a context as first argument
* `client.NewMultiPartRequestWithContext()` has been removed in favor of `client.NewMultiPartRequest()`, which requires now a context as first argument
* `context` is now a first class citizen in all API calls. Functions that had a suffix like `...WithContext` have been removed entirely. The API methods support the context now as first argument.
* `BoardService.GetAllSprints` has been removed and `BoardService.GetAllSprintsWithOptions` has been renamed to `BoardService.GetAllSprints`
* `GroupService.Get` has been removed and `GroupService.GetWithOptions` has been renamed to `GroupService.Get`
* `Issue.Update` has been removed and `Issue.UpdateWithOptions` has been renamed to `Issue.Update`
* `Issue.GetCreateMeta` has been removed and `Issue.GetCreateMetaWithOptions` has been renamed to `Issue.GetCreateMeta`
* `Project.GetList` has been removed and `Project.ListWithOptions` has been renamed to `Project.GetAll`
* Cloud/Authentication: Removed `BearerAuthTransport`, because it was a (kind of) duplicate of `BasicAuthTransport`
* Cloud/Authentication: Removed `PATAuthTransport`, because it was a (kind of) duplicate of `BasicAuthTransport`
* Cloud/Authentication: `BasicAuthTransport.Password` was renamed to `BasicAuthTransport.APIToken`
* Cloud/Authentication: Removes `CookieAuthTransport` and `AuthenticationService`, because this type of auth is not supported by the Jira cloud offering
* Cloud/Component: The type `CreateComponentOptions` was renamed to `ComponentCreateOptions`
* Cloud/User: Renamed `User.GetSelf` to `User.GetCurrentUser`
* Cloud/Group: Renamed `Group.Add` to `Group.AddUserByGroupName`
* Cloud/Group: Renamed `Group.Remove` to `Group.RemoveUserByGroupName`

### Features

* UserAgent: Client HTTP calls are now identifable via a User Agent. This user agent can be configured (default: `go-jira/2.0.0`)
* The underlying used HTTP client for API calls can be retrieved via `client.Client()`
* API-Version: Official support for Jira Cloud API in [version 3](https://developer.atlassian.com/cloud/jira/platform/rest/v3/intro/)

### Bug Fixes

* README: Fixed all (broken) links

### API-Endpoints

* Workflow status categories: Revisited and fully implemented for Cloud and On Premise (incl. examples)

### Other

* Replace all "GET", "POST", ... with http.MethodGet (and related) constants
* Development: Added `make` commands to collect (unit) test coverage
* Internal: Replaced `io.ReadAll` and `json.Unmarshal` with `json.NewDecoder`

## [1.17.0](https://github.com/andygrunwald/go-jira/compare/v1.16.1...v1.17.0) (2025-09-16)

All changes in [#735](https://github.com/andygrunwald/go-jira/pull/735).

### Breaking Changes

* Minimum Go version increased from 1.15 to 1.21 (due to google/go-cmp dependency update)

### Security

* Updated github.com/golang-jwt/jwt/v4 v4.4.2 to v4.5.2 (CVE-2025-30204)

### Maintenance

* Replaced deprecated `ioutil.ReadFile` with `os.ReadFile`
* Replaced deprecated `ioutil.ReadAll` with `io.ReadAll`
* Replaced deprecated `ioutil.Discard` with `io.Discard`
* Upgraded staticcheck v2022.1 to v2023.1
* Code formatting updates (issue.go, metaissue.go, sprint.go)

### Dependencies

* github.com/google/go-cmp v0.5.8 to v0.7.0
* github.com/golang-jwt/jwt/v4 v4.4.2 to v4.5.2

### CI/CD

* Upgraded actions/checkout from v3 to v5
* Upgraded actions/setup-go from v3 to v6
* Upgraded dominikh/staticcheck-action from v1.2 to v1.4

## [1.16.1](https://github.com/andygrunwald/go-jira/compare/v1.16.0...v1.16.1) (2025-09-13)

### Features

* Added `IssueService.SearchV2JQL()` and `IssueService.SearchV2JQLWithContext()` to handle Atlassian's deprecation of `GET /rest/api/2/search` endpoint (effective October 31, 2024) ([#725](https://github.com/andygrunwald/go-jira/pull/725))

*New contributor: @conor-naranjo*

## [1.16.0](https://github.com/andygrunwald/go-jira/compare/v1.15.1...v1.16.0) (2022-07-08)

### Features

* Added example demonstrating how to add labels to issues ([#442](https://github.com/andygrunwald/go-jira/pull/442))

### Maintenance

* Test matrix expanded to include macOS ([#449](https://github.com/andygrunwald/go-jira/pull/449))
* Deprecated Go v1.15 and Go v1.16; added Go v1.18 ([#449](https://github.com/andygrunwald/go-jira/pull/449))
* Implemented nightly CI runs ([#449](https://github.com/andygrunwald/go-jira/pull/449))
* Removed greetings workflow ([#450](https://github.com/andygrunwald/go-jira/pull/450))
* Upgraded staticcheck action to v1.2.0 ([#451](https://github.com/andygrunwald/go-jira/pull/451))

### Dependencies

* github.com/golang-jwt/jwt/v4 v4.3.0 to v4.4.2 ([#443](https://github.com/andygrunwald/go-jira/pull/443), [#445](https://github.com/andygrunwald/go-jira/pull/445), [#464](https://github.com/andygrunwald/go-jira/pull/464))
* github.com/google/go-cmp v0.5.7 to v0.5.8 ([#453](https://github.com/andygrunwald/go-jira/pull/453))
* actions/checkout v2 to v3 ([#448](https://github.com/andygrunwald/go-jira/pull/448))
* actions/cache v2 to v3.0.1 ([#447](https://github.com/andygrunwald/go-jira/pull/447))
* actions/setup-go v2 to v3 ([#452](https://github.com/andygrunwald/go-jira/pull/452))

*New contributor: @lucribas*

## [1.15.1](https://github.com/andygrunwald/go-jira/compare/v1.15.0...v1.15.1) (2022-02-26)

### Features

* Added customer and request endpoints for Service Desk API ([#391](https://github.com/andygrunwald/go-jira/pull/391))
* Added "environment" field to Issue struct ([#440](https://github.com/andygrunwald/go-jira/pull/440))
* User module improvements ([#439](https://github.com/andygrunwald/go-jira/pull/439))

### Bug Fixes

* Added support for update section in DoTransitionWithPayload struct ([#423](https://github.com/andygrunwald/go-jira/pull/423))
* Fixed resource leaks ([#399](https://github.com/andygrunwald/go-jira/pull/399))

*New contributors: @CondensedTea, @nansuri, @metalbreeze, @iamjem, @dustin-decker*

## [1.15.0](https://github.com/andygrunwald/go-jira/compare/v1.14.0...v1.15.0) (2022-02-18)

### Features

* Upgraded to JWT Library V4 ([#403](https://github.com/andygrunwald/go-jira/pull/403))
* Added Bearer token support for OAuth 2.0 ([#397](https://github.com/andygrunwald/go-jira/pull/397))
* Added Personal Access Token (PAT) authentication support ([#422](https://github.com/andygrunwald/go-jira/pull/422))
* Added method for updating remote issue links by ID ([#411](https://github.com/andygrunwald/go-jira/pull/411))
* Added support for ServiceDeskID as strings ([#419](https://github.com/andygrunwald/go-jira/pull/419))
* Added missing min/max properties to BoardConfigurationColumn ([#430](https://github.com/andygrunwald/go-jira/pull/430))

### Bug Fixes

* Fixed missing struct omitempty tags in Parent struct ([#435](https://github.com/andygrunwald/go-jira/pull/435))

### Documentation

* Fixed import alias example in README ([#428](https://github.com/andygrunwald/go-jira/pull/428))

### Dependencies

* github.com/golang-jwt/jwt/v4 v4.1.0 to v4.3.0 ([#424](https://github.com/andygrunwald/go-jira/pull/424), [#438](https://github.com/andygrunwald/go-jira/pull/438))
* github.com/google/go-cmp v0.5.6 to v0.5.7 ([#432](https://github.com/andygrunwald/go-jira/pull/432))

## [1.14.0](https://github.com/andygrunwald/go-jira/compare/v1.13.0...v1.14.0) (2021-07-28)

### Features

* Added description field for component ([#365](https://github.com/andygrunwald/go-jira/pull/365))
* Added Jira ServiceDesk support with organizations ([#342](https://github.com/andygrunwald/go-jira/pull/342))
* Implemented delete issue link API ([#341](https://github.com/andygrunwald/go-jira/pull/341))
* Added Items and Custom properties to FieldSchema ([#332](https://github.com/andygrunwald/go-jira/pull/332))
* Added StartedAfter property to GetWorklogsQueryOptions ([#344](https://github.com/andygrunwald/go-jira/pull/344))

### Bug Fixes

* Fixed potential nil reference issues ([#342](https://github.com/andygrunwald/go-jira/pull/342))
* Fixed context parameter propagation ([#363](https://github.com/andygrunwald/go-jira/pull/363))
* Fixed variable total calculation in pagination example ([#354](https://github.com/andygrunwald/go-jira/pull/354))
* Fixed boolean type handling for release versions ([#336](https://github.com/andygrunwald/go-jira/pull/336))
* Fixed RemoteLinkStatus ([#303](https://github.com/andygrunwald/go-jira/pull/303))

### Documentation

* Updated examples to use accountID instead of deprecated username ([#374](https://github.com/andygrunwald/go-jira/pull/374))
* Added example for creating issues with custom fields ([#358](https://github.com/andygrunwald/go-jira/pull/358))
* Fixed examples/create response output ([#357](https://github.com/andygrunwald/go-jira/pull/357))
* Added pagination example ([#354](https://github.com/andygrunwald/go-jira/pull/354))
* Switched examples from deprecated terminal.ReadPassword to term.ReadPassword ([#359](https://github.com/andygrunwald/go-jira/pull/359))

### Maintenance

* CI/CD workflow cleanup ([#384](https://github.com/andygrunwald/go-jira/pull/384))
* Added funding information ([#383](https://github.com/andygrunwald/go-jira/pull/383))
* Upgraded to GitHub-native Dependabot ([#375](https://github.com/andygrunwald/go-jira/pull/375))
* Resolved nogo tautological linting error ([#364](https://github.com/andygrunwald/go-jira/pull/364))
* Updated Go versions in testing workflow ([#349](https://github.com/andygrunwald/go-jira/pull/349))

### Dependencies

* JWT library upgrade ([#387](https://github.com/andygrunwald/go-jira/pull/387))
* github.com/google/go-cmp v0.3.0 to v0.5.6 ([#347](https://github.com/andygrunwald/go-jira/pull/347), [#360](https://github.com/andygrunwald/go-jira/pull/360), [#379](https://github.com/andygrunwald/go-jira/pull/379))
* github.com/pkg/errors v0.8.0 to v0.9.1 ([#345](https://github.com/andygrunwald/go-jira/pull/345))
* github.com/trivago/tgo v1.0.1 to v1.0.7 ([#346](https://github.com/andygrunwald/go-jira/pull/346))
* github.com/fatih/structs v1.0.0 to v1.1.0 ([#348](https://github.com/andygrunwald/go-jira/pull/348))

## [1.13.0](https://github.com/andygrunwald/go-jira/compare/v1.11.1...v1.13.0) (2020-10-25)

### Features

* add AddRemoteLink method ([f200e15](https://github.com/andygrunwald/go-jira/commit/f200e158b997a303db081cbbc5a9d8ad5d89566d)), closes [/developer.atlassian.com/cloud/jira/platform/rest/v2/#api-rest-api-2](https://github.com/andygrunwald//developer.atlassian.com/cloud/jira/platform/rest/v2//issues/api-rest-api-2)
* Add Names support on Issue struct ([#278](https://github.com/andygrunwald/go-jira/issues/278)) ([1fc10e0](https://github.com/andygrunwald/go-jira/commit/1fc10e0606784f745673ccc4d8d706c36f385a7a))
* Extend Makefile for more source code quality targets ([5e52236](https://github.com/andygrunwald/go-jira/commit/5e5223631a29d10a13e598318a6abe47384e2982))
* **context:** Add support for context package ([e1f4265](https://github.com/andygrunwald/go-jira/commit/e1f4265e2b467b938fe0c095caf6d36f3136d2ff))
* **issues:** Add GetEditMeta on issue ([a783764](https://github.com/andygrunwald/go-jira/commit/a783764b52dc890773658ddd0483a9d0393e385d)), closes [/docs.atlassian.com/DAC/rest/jira/6.1.html#d2e1364](https://github.com/andygrunwald//docs.atlassian.com/DAC/rest/jira/6.1.html/issues/d2e1364)
* **IssueService:** allow empty JQL ([#268](https://github.com/andygrunwald/go-jira/issues/268)) ([4b91cf2](https://github.com/andygrunwald/go-jira/commit/4b91cf2b135355de7ecee41727c3e65f4e7067bc))
* **project:** Add cronjob to check for stale issues ([#287](https://github.com/andygrunwald/go-jira/issues/287)) ([2096b04](https://github.com/andygrunwald/go-jira/commit/2096b04e52b434c1fb1c841bab487a94674a271e))
* **project:** Add GitHub Actions testing workflow ([#289](https://github.com/andygrunwald/go-jira/issues/289)) ([80c0282](https://github.com/andygrunwald/go-jira/commit/80c02828ca9e4eb0e4a1877275baae14d330a2d9)), closes [#290](https://github.com/andygrunwald/go-jira/issues/290)
* **project:** Add workflow to greet new contributors ([#288](https://github.com/andygrunwald/go-jira/issues/288)) ([c357b61](https://github.com/andygrunwald/go-jira/commit/c357b61a40f62a919ebd94a555390958f99c8db7))

### Bug Fixes

* change millisecond time format ([8c77107](https://github.com/andygrunwald/go-jira/commit/8c77107df3757c4ec5eae6e9d7c018618e708bfa))
* paging with load balancer going to endless loop ([19d3fc0](https://github.com/andygrunwald/go-jira/commit/19d3fc0aecde547ffe1ab547c5ffb6c7972d387c)), closes [#260](https://github.com/andygrunwald/go-jira/issues/260)
* **issue:** IssueService.Search() with a not empty JQL triggers 400 bad request ([#292](https://github.com/andygrunwald/go-jira/issues/292)) ([8b64c7f](https://github.com/andygrunwald/go-jira/commit/8b64c7f005fbceb11fa43a7aff3de61eb3166fca)), closes [#291](https://github.com/andygrunwald/go-jira/issues/291)
* **IssueService.GetWatchers:** UserService.GetByAccountID support accountId params ([436469b](https://github.com/andygrunwald/go-jira/commit/436469b62d4d62037f380b38c918a13f4a5f0ab2))
* **product:** Make product naming consistent, rename JIRA to Jira ([#286](https://github.com/andygrunwald/go-jira/issues/286)) ([146229d](https://github.com/andygrunwald/go-jira/commit/146229d2ab58a3fb128ddc8dcbe03aff72e20857)), closes [#284](https://github.com/andygrunwald/go-jira/issues/284)
* **tests:** Fix TestIssueService_PostAttachment unit test ([f6b1dca](https://github.com/andygrunwald/go-jira/commit/f6b1dcafcfdd8fe69f842b1053c4030da6c97c7f))
* removing the use of username field in searching for users ([#297](https://github.com/andygrunwald/go-jira/issues/297)) ([f50cb07](https://github.com/andygrunwald/go-jira/commit/f50cb07b297d79138b13e5ab49ea33965d32f5c1))

## [1.12.0](https://github.com/andygrunwald/go-jira/compare/v1.11.1...v1.12.0) (2019-12-14)

### Features

* Add IssueLinkTypeService with GetList and test ([261889a](https://github.com/andygrunwald/go-jira/commit/261889adc63623fcea0fa8cab0d5da26eec37e68))
* add worklog update method ([9ff562a](https://github.com/andygrunwald/go-jira/commit/9ff562ae3ea037961f277be10412ad0a42ff8a6f))
* Implement get remote links method ([1946cac](https://github.com/andygrunwald/go-jira/commit/1946cac0fe6ee91f784e3dda3c12f3f30f7115b8))
* Implement issue link type DELETE ([e37cc6c](https://github.com/andygrunwald/go-jira/commit/e37cc6c6897830492c070667ab8b68bd85683fc3))
* Implement issue link type GET ([57538b9](https://github.com/andygrunwald/go-jira/commit/57538b926c558e97940760a30bdc16cdd37ef4f1))
* Implement issue link type POST ([75b9df8](https://github.com/andygrunwald/go-jira/commit/75b9df8b01557f01dc318d33c0bc2841a9c084eb))
* Implement issue link type PUT ([48a15c1](https://github.com/andygrunwald/go-jira/commit/48a15c10443a3cff78f0fb2c8034dd772320e238))
* provide access to issue transitions loaded from JIRA API ([7530b7c](https://github.com/andygrunwald/go-jira/commit/7530b7cd8266d82cdb4afe831518986772e742ba))

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
