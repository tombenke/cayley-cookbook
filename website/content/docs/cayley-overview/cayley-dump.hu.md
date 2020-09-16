---
title: "The `cayley dump` parancs"
weight: 4
bookCollapseSection: true
---

# The `cayley dump` parancs

A `cayley dump` parancs kiírja a resource leírókat az aktuális adatbázisból egy szabványos formátumú file-ba. A formátumot a kiterjesztés alapján határozza meg, de megadható neki a `--dump_format <format-name>` argumentummal is, ahol a `<format-name>` lehetséges értékei: "gml", "graphml", "graphviz", "json", "json-stream", "jsonld", "nquads", "pquads".

{{< figure src="/cayley-cookbook/cayley-dump.png" title="1. ábra: cayley dump" >}}

Az alábbi parancs kiírja az adatokat a `dump.nq` file-ba:

```bash

    $ cayley dump --db bolt --dbpath /home/tombenke/tmp/cayley/ -o dump.out
    I0916 07:15:25.453665    5286 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0916 07:15:25.453848    5286 database.go:187] using backend "bolt" (/home/tombenke/tmp/cayley/)
    writing quads to file "dump.out"
    15 entries were written

```


Az aktuális adatok kiíratása az `stdout`-ra:

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

