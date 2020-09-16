---
title: "A `cayley repl` parancs"
weight: 5
bookCollapseSection: true
---

# A `cayley repl` parancs

A REPL aRead-Eval-Print-Loop kifejezés rövidítése.

A `cayley repl` paranccsal elindíthatunk egy konzolt ami egy már korábban létrehozott adatbázishoz kapcsolódik, és interaktívan végrehajthatunk az adatbázisban lekérdezéseket, és egyéb műveleket.

A konzol a normál Linux konzolhoz hasonlóan működik.

Beírhatjuk a kifejezéseket, amelyeket az Enter billenytű lenyomása után az interpreter értelmez, és végrehajt, az eredményt pedig kiírja a konzolra.
A fel/le nyilakat is használhatjuk, hogy navigáljunk a korábban végrehajtott parancsok között.
A kifejezések, amiket végrehajtunk, bekerülnek egy `.cayley_history` nevű file-ba, amit később megnyithatunk egy text editorral, és átemelhetjük belőle azokat a kifejezéseket, amelyek hasznosnak bizonyultak.
Ezen a módon a REPL-t hatékonyan alkalmazhatjuk a query-k kialakítására, és a kísérletek eredményét végül átemelhetjük a véglegesnek szánt query file-okba.

{{< figure src="/cayley-cookbook/cayley-repl.png" title="Figure 3.: cayley repl" >}}

A következő lépéssorozat elindít egy REPL session-t egy létező adatbázison, majd végrehajt egy egyszerű query-t, ami megkeresi és kilistázza az összes node-ot, amit az adatbázisban talál:

```bash

    $ cayley repl --db bolt --dbpath /home/tombenke/tmp/cayley/ 
    I0915 18:02:08.623789   25897 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0915 18:02:08.623971   25897 database.go:187] using backend "bolt" (/home/tombenke/tmp/cayley/)
    cayley> g.V().All()

    ****
    id : <bob>
    ****
    id : <status>
    ****
    id : cool_person
    ****
    id : <alice>
    ****
    id : <greg>
    ****
    id : <emily>
    ****
    id : <smart_graph>
    ****
    id : <predicates>
    ****
    id : <dani>
    ****
    id : <fred>
    ****
    id : smart_person
    ****
    id : <charlie>
    ****
    id : <are>
    ****
    id : <follows>
    -----------
    14 Results
    Elapsed time: 3.046727 ms

    cayley>

```

Lássuk, hogy a session-t követően mit találunk a `.cayley_history` file-ban:
```bash

    $ cat .cayley_history 
    
    g.V().All()

```
