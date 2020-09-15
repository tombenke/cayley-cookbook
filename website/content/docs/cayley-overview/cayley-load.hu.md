---
title: "The `cayley load` parancs"
weight: 3
bookCollapseSection: true
---

# The `cayley load` parancs

A `cayley load` paranccsal betölthetünk resource leíró file-ok tartalmát egy korábban létrehozott store-ba.

{{< figure src="/cayley-cookbook/cayley-load.png" title="Figure 1.: cayley load" >}}

Az alábbi parancs feltölti a `data/testdata.nq` file tartalmát egy adatbázisba:

```bash

    $ cayley load --db bolt --dbpath /home/tombenke/tmp/cayley/ --load data/testdata.nq 

    I0915 18:01:07.140303   25803 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0915 18:01:07.140500   25803 database.go:187] using backend "bolt" (/home/tombenke/tmp/cayley/)

```

