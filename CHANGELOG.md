# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres
to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.5] - 2025-06-27

### Added
- Added fungible ids to dao list

## [0.2.4] - 2025-05-27

### Added
- Extended token info

## [0.2.3] - 2025-03-25

### Added
- Dao token symbol

## [0.2.2] - 2025-02-05

### Added
- Added token info and chart

## [0.2.1] - 2024-12-24

### Added
- Added vote status delegation

## [0.2.0] - 2024-12-04

### Added
- All delegations list by address
- All delegators list by address
- Delegation summary by address

## [0.1.14] - 2024-10-08

### Added
- Extend DAO object with active proposals ids

## [0.1.13] - 2024-10-02

### Added
- Support vote now

## [0.1.12] - 2024-09-22

### Added
- Added delegates total

## [0.1.11] - 2024-09-16

### Added
- Vp list for proposal

## [0.1.10] - 2024-09-09

### Added
- Added proxy api for delegates

## [0.1.9] - 2024-08-20

### Changed
- GetByID can search by internal and original DAO identifier

## [0.1.8] - 2024-08-12

### Changed
- The parameter's name for search

## [0.1.7] - 2024-07-22

### Added
- Search for votes 

## [0.1.6] - 2024-07-05

### Changed
- Extend vote response with proposal identifier

## [0.1.5] - 2024-04-19

### Added
- DAO recommendations endpoint

## [0.1.4] - 2024-04-10

### Added
- Daos where voter participates in

## [0.1.3] - 2024-03-22

### Added
- Stats endpoint

## [0.1.2] - 2024-03-15

### Added
- Total vp for votes endpoints

## [0.1.1] - 2024-03-13

### Added
- Ens names

## [0.1.0] - 2024-03-02

### Changed
- Changed the path name of the go module

### Added
- Added LICENSE information
- Added info for contributing
- Added github issues templates
- Added linter and unit-tests workflows for github actions
- Added badges with link to the license and passed workflows

## [0.0.22] - 2024-02-06

### Added
- Active votes, verified fields to dao

## [0.0.21] - 2024-02-05

### Added
- Order proposals votes by voter

## [0.0.20] - 2024-01-30

### Added
- User Votes

## [0.0.19] - 2023-12-14

### Added
- Author ens name field for votes

## [0.0.18] - 2023-12-06

### Added
- Author ens name field for proposals

## [0.0.17] - 2023-12-04

### Added
- Add voting methods

## [0.0.16] - 2023-10-09

### Added
- Voters count field for dao info

## [0.0.15] - 2023-10-06

### Added
- VotingEndsSoon const

## [0.0.14] - 2023-09-18

### Added
- Feed client

## [0.0.13] - 2023-09-07

### Changed
- Mark votes choice field as json.RawMessage due to multiple values

## [0.0.12] - 2023-08-23

### Added
- Proposal timeline field

## [0.0.11] - 2023-07-18

### Changed
- Extend vote model

## [0.0.10] - 2023-07-14

### Fixed
- Supported feed.timeline

## [0.0.9] - 2023-07-14

### Fixed
- Fixed parentID field (temporary removed)

## [0.0.8] - 2023-07-14

### Added
- Dao activity since field

## [0.0.7] - 2023-07-14

### Added
- Added golangci-lint config

### Changed
- Updated structure of the requests and responses
- Used uuid instead of strings if it's required

## [0.0.6] - 2023-07-11

### Added
- Proposal top endpoint

## [0.0.5] - 2023-07-11

### Fixed
- Fixed missed fields in DAO and Strategy

## [0.0.4] - 2023-07-07

### Added
- Filtering proposals by title

## [0.0.3] - 2023-07-06

### Added
- Prepare flat feed

## [0.0.2] - 2023-06-29

### Added
- Filter dao by ids

## [0.0.1] - 2023-06-16

### Added
- Basic client for core-web-api
