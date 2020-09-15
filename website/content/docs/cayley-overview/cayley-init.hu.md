---
title: "A `cayley init` parancs"
weight: 2
bookCollapseSection: true
---

# A `cayley init` parancs

A `cayley init` paranccsal létrehozhatunk egy új, üres adatbázist.

{{< figure src="/cayley-cookbook/cayley-init.png" title="Figure 1.: cayley init" >}}

Az alábbi parancs létrehoz egy perzisztens Bolt típusú key-value store-t, a `/home/tombenke/tmp/cayleydb` folder-ben:

```bash

    $ cayley init --db bolt --dbpath /home/tombenke/tmp/cayleydb 
    
    I0915 17:52:56.529687   24933 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0915 17:52:56.529904   24933 database.go:187] using backend "bolt" (/home/tombenke/tmp/cayleydb)

```

Az eredményképpen létrejött adatbázist használhatjuk a többi programmal, úgymint: `load`, `dump`, `repl`, `query` and `http`.

