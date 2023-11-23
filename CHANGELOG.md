# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres
to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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
