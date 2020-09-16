---
title: "A `cayley query` parancs"
weight: 6
bookCollapseSection: true
---

# A `cayley query` parancs


A `cayley query` paranccsal végrehajthatunk egy query-t egy létező adatbázison.


A query kifejezést, amit vigre kívánunk hajtani, a standard input-on keresztül továbbíthatjuk a parancshoz. Az egyszerűbb kifejezések esetében elegendő lehet, ha azokat az `echo` utility-vel továbbítjuk, és közvetlenül beírjuk a parancssorba. A hosszabb kifejezéseket célszerűbb lehet query file-okban elhelyezni, és a `cat` utility-vel továbbítani a parancs számára..

A `--lang <query-language-type>` parameterrel kijelölhetjük, hogy a query kifejezést melyik nyelven írtuk. A lehetséges válastható nyelvek: "gizmo", "graphql", "mql", "sexp". Az elepértelmezett query nyelv a "gizmo".

{{< figure src="/cayley-cookbook/cayley-query.png" title="Figure 4.: cayley query" >}}

Futtasunk egy query-t egy korábban létrehozott store tartalmán:
```bash

    echo "g.V().All();" | cayley query --logtostderr false --db bolt --dbpath /home/tombenke/tmp/cayley --logs ~/tmp
    {"id":"bob"}
    {"id":"status"}
    {"id":"cool_person"}
    {"id":"alice"}
    {"id":"greg"}
    {"id":"emily"}
    {"id":"smart_graph"}
    {"id":"predicates"}
    {"id":"dani"}
    {"id":"fred"}
    {"id":"smart_person"}
    {"id":"charlie"}
    {"id":"are"}
    {"id":"follows"}

```

Amennyiben a query eredményét közvetlenül egy másik programba szeretnénk betáplálni, akkor szükségünk lehet arra, hogy a logokat eltüntessük.
Az alábbi példa bemutatja, hogyan tehetjük ezt meg.
A logok a `/home/tombenke/tmp` directory-ban lesznek elhelyezve:

```bash

    $ cat ~/tmp/cayley.tombenke-Latitude-E5470.tombenke.log.INFO.20200915-182850.27351 
    Log file created at: 2020/09/15 18:28:50
    Running on machine: tombenke-Latitude-E5470
    Binary: Built with gc go1.11.2 for linux/amd64
    Log line format: [IWEF]mmdd hh:mm:ss.uuuuuu threadid file:line] msg

```

Ezt a trükköt, szükség esetén, a többi parancs esetében is alkalmazhatjuk.

Gyakran előfordul, hogy a query-k fejlesztése, tesztelése során valójában nincs szükségünk perzisztens store-ra. Ilyenkor sokkal egyszerűbb lehet egy átmeneti, in-memory adatbázis használata, amibe a futás idejére betöltjük a tesztelésghez szükséges adatokat a `--load <data-file>` parancssori parameterel.
Az alábbi parancs ugyanazt a műveletet hajtja végre, mint a fentebb bemutatott példa, de a query-t egy in-memory adatbázison hajtja végre:

```bash

    $ echo "g.V().All();" | cayley query --logtostderr false --load data/testdata.nq --logs ~/tmp

    {"id":"alice"}
    {"id":"follows"}
    {"id":"bob"}
    {"id":"fred"}
    {"id":"status"}
    {"id":"cool_person"}
    {"id":"dani"}
    {"id":"charlie"}
    {"id":"greg"}
    {"id":"emily"}
    {"id":"predicates"}
    {"id":"are"}
    {"id":"smart_person"}
    {"id":"smart_graph"}

```


