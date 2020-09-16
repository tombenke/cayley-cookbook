---
title: "The `cayley dump` command"
weight: 4
bookCollapseSection: true
---

# The `cayley dump` command

The `cayley dump` command dumps out the resource descriptions from an existing database into a file or to the standard output. The format type is determined by the given file extension, but it is also can be selected by the `--dump_format <format-name>` argument, where the possible values of `<format-name>` are: "gml", "graphml", "graphviz", "json", "json-stream", "jsonld", "nquads", "pquads".

{{< figure src="/cayley-cookbook/cayley-dump.png" title="Figure 1.: cayley dump" >}}

The following command dumps the content to the `dump.nq` file:

```bash

    $ cayley dump --db bolt --dbpath /home/tombenke/tmp/cayley/ -o dump.out
    I0916 07:15:25.453665    5286 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0916 07:15:25.453848    5286 database.go:187] using backend "bolt" (/home/tombenke/tmp/cayley/)
    writing quads to file "dump.out"
    15 entries were written

```


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

