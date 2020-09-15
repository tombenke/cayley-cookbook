---
title: "The `cayley dump` command"
weight: 3
bookCollapseSection: true
---

# The `cayley dump` command

Dump the actual content of a database to the `stdout`:

```bash

    $ cayley dump --db bolt --dbpath /home/tombenke/tmp/cayley/ 
    I0915 18:03:30.188245   25985 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0915 18:03:30.188453   25985 database.go:187] using backend "bolt" (/home/tombenke/tmp/cayley/)
    I0915 18:03:30.191235   25985 dump.go:20] writing quads to stdout
    <alice> <follows> <bob> .
    <bob> <follows> <fred> .
    <bob> <status> "cool_person" .
    <dani> <follows> <bob> .
    <charlie> <follows> <bob> .
    <charlie> <follows> <dani> .
    <dani> <follows> <greg> .
    <dani> <status> "cool_person" .
    <emily> <follows> <fred> .
    <fred> <follows> <greg> .
    <greg> <status> "cool_person" .
    <predicates> <are> <follows> .
    <predicates> <are> <status> .
    <emily> <status> "smart_person" <smart_graph> .
    <greg> <status> "smart_person" <smart_graph> .

```

