[![GitHub release](https://img.shields.io/github/release/cybozu-go/myso.svg?maxAge=60)][releases]
[![CircleCI](https://circleci.com/gh/cybozu-go/myso.svg?style=svg)](https://circleci.com/gh/cybozu-go/myso)
[![GoDoc](https://godoc.org/github.com/cybozu-go/myso?status.svg)][godoc]
[![Go Report Card](https://goreportcard.com/badge/github.com/cybozu-go/myso)](https://goreportcard.com/report/github.com/cybozu-go/myso)
[![Docker Repository on Quay](https://quay.io/repository/cybozu/myso/status "Docker Repository on Quay")](https://quay.io/repository/cybozu/myso)

MySO
====

MySO is a MySQL operator to construct and manage lossless semi-sync replicated MySQL instances using binary logs.

MySO is designed to the following properties:

- Integrity
    - Do not lose any data under a given degree of faults (double or triple).
- Availability
    - Keep the MySQL cluster available under a given degree of faults (single or double).
- Serviceability
    - Even if a large-scale failure (e.g. data center blackout) occurs, MySO performs a quick recovery by combining full backup and binary logs.

Features
--------

TBD

Documentation
--------------

[docs](docs/) directory contains documents about designs and specifications.

License
-------
SQL Operator is licensed under MIT license.

[releases]: https://github.com/cybozu-go/myso/releases
[godoc]: https://godoc.org/github.com/cybozu-go/myso
