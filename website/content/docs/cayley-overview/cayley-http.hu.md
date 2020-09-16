---
title: "A `cayley http` parancs"
weight: 7
bookCollapseSection: true
---

# A `cayley http` parancs

Talán a `cayley http` a legismertebb, és leggyakrabban használt a cayley parancsok közül.
Ez egy olyan alkalmazás, amely tartalmaz quad-store-t, query engine-t és egy HTTP server-t is, ami a {{< cayley >}} lényegében összes funkcionalitását elérhetővé teszi a REST API interfészen keresztül külső alkalmazások számára, ezen felül egy web-es grafikus felületet is biztosít a felhasználóinak.

Az 1. ábra szemlélteti az összes olyan komponenst, amit használni lehet a `cayley http` üzemmódban futó {{< cayley >}}-vel.

{{< figure src="/cayley-cookbook/cayley-http.png" title="1. ábra: cayley http" >}}

A `cayley http` által biztosított REST API végpontok leírását a [hivatalos dokumentációban](https://cayley.gitbook.io/cayley/usage/http) olvashatjuk. A [3rd party APIs](https://cayley.gitbook.io/cayley/usage/3rd-party-apis) dokumentációs oldalak leírják, hogy milyen programozási nyelveken, és milyen könyvtárakkal férhetünk hozzá a `cayley http` server-hez

## A Cayley HTTP server elindítása

### Indítás alapértelmezett beállításokkal

A `cayley http` parancs végrehajtása elindítja a server-t az alapértelmezett beállításokkal:

```bash

    $ cayley http

    I0916 08:00:12.107989    6702 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0916 08:00:12.108196    6702 database.go:187] using backend "memstore"
    I0916 08:00:12.108322    6702 http.go:197] using assets from "/snap/cayley/2"
    I0916 08:00:12.108805    6702 http.go:42] listening on 127.0.0.1:64210, web interface at http://127.0.0.1:64210

```

A server ekkor egy üres, in-memory adatbázist hoz létre. A server alapértelmezés szerint a `64210` porton figyel. Egy böngészővel megnyithatjuk a `http://127.0.0.1:64210` URL-t, és ekkor megjelenik a web UI, amelyen keresztül query-ket hajthatunk végre az adatbázison.

### Indítás in-memory store-ral, automatikus adatbetöltéssel

Jelezhetjük a `cayley http` parancs számára, hogy az indulást követően azonnal töltse fel egy adatfile tartalmát az adatbázisba. Erre a célra a `--load <data-file-path>` argumentumot alkalmazzuk.

```bash

    $ cayley http --load data/testdata.nq 
    I0916 08:10:40.473364    7087 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0916 08:10:40.473602    7087 database.go:187] using backend "memstore"
    I0916 08:10:40.473935    7087 database.go:250] loaded "data/testdata.nq" in 266.525µs
    I0916 08:10:40.474084    7087 http.go:197] using assets from "/snap/cayley/2"
    I0916 08:10:40.474498    7087 http.go:42] listening on 127.0.0.1:64210, web interface at http://127.0.0.1:64210

```

A feltöltött adatok azonnal rendelkezésre fognak állni, és mind a web-es UI-on, mind a REST API-n keresztül lekérdezéseket és módosításokat hajthatunk végre azokon. Az adatok ebben az esetben is még csak in-memory tárolódnak.

### Indítás perzisztens store-ral, és adatfeltöltéssel

A `cayley http` parancsot úgy is indíthatjuk, hogy az perzisztens adatbázist használjon. Ebben az esetben két lehetőség közül választhatunk:
1. Inicializálunk egy üres, perzisztens store-t, és esetleg adatokat is töltünk fel rögtön, az indulást követően.
2. Egy már korábban létrehozott üres, vagy adatokkal feltöltött adatbázist nyitunk meg. Egy ilyen adatbázis létrehozásához a `cayley init` és `cayley load` parancsokat használhatjuk.

Indításkor hozzunk létre egy perzisztens adatbázist, és rögtön töltsünk is bele adatokat:
```bash

    $ cayley http --init --db bolt --dbpath /home/tombenke/tmp/cayley/ --load data/testdata.nq 
    I0916 08:20:05.039047    7367 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0916 08:20:05.039234    7367 database.go:187] using backend "bolt" (/home/tombenke/tmp/cayley/)
    I0916 08:20:05.070109    7367 database.go:250] loaded "data/testdata.nq" in 9.150519ms
    I0916 08:20:05.070347    7367 http.go:197] using assets from "/snap/cayley/2"
    I0916 08:20:05.071829    7367 http.go:42] listening on 127.0.0.1:64210, web interface at http://127.0.0.1:64210
    I0916 08:20:11.371920    7367 http.go:100] started POST /api/v1/query/gizmo for 127.0.0.1:39426
    I0916 08:20:11.372573    7367 http.go:102] completed 200 OK /api/v1/query/gizmo in 666.087µs
    I0916 08:20:17.928216    7367 http.go:100] started POST /api/v1/query/gizmo for 127.0.0.1:39426
    I0916 08:20:17.932551    7367 http.go:102] completed 200 OK /api/v1/query/gizmo in 4.335239ms

```

készítsünk elő egy perzisztens store-t az `init` és `load` parancsokkal, majd indítsuk el a `http` server-t:
```bash

    $ cayley init --db bolt --dbpath /home/tombenke/tmp/cayley/
    I0916 08:21:00.701903    7460 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0916 08:21:00.702107    7460 database.go:187] using backend "bolt" (/home/tombenke/tmp/cayley/)

    $ cayley load --db bolt --dbpath /home/tombenke/tmp/cayley/ --load data/testdata.nq
    I0916 08:21:42.843042    7509 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0916 08:21:42.843269    7509 database.go:187] using backend "bolt" (/home/tombenke/tmp/cayley/)

    $ cayley http --db bolt --dbpath /home/tombenke/tmp/cayley/
    I0916 08:22:08.376538    7573 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0916 08:22:08.376726    7573 database.go:187] using backend "bolt" (/home/tombenke/tmp/cayley/)
    I0916 08:22:08.379455    7573 http.go:197] using assets from "/snap/cayley/2"
    I0916 08:22:08.380016    7573 http.go:42] listening on 127.0.0.1:64210, web interface at http://127.0.0.1:64210

```

Mindkét esetben elérhetjük az adatokat a REST API-n keresztül, rögtön az indítást követően.

{{< hint info >}}
__Note:__

Mind a `cayley load` és a `cayley save` parancsok használhatóak adatfeltöltésre, ill. kiírásra olyan adat adatbázisból, amit a `cayley http` server is el tud érni, de csak akkor, ha a server le van állítva, mivel az adatbázist egyidejűleg csak egy alkalmazás használhatja közvetlenül. Emiatt az 1. ábrán ezt a két segédprogramot szaggatott vonallal kötöttük össze a store-ral.

{{< /hint >}}

## A Web-es felhasználói felület

Indítsuk el a `cayley http` server-t akár in-memory, akár perzisztens store-ral, úgy, hogy az tartalmazza a feltöltött teszt adatokat, ahogyan ezt az imént leírtuk a megelőző alfejezetekben. Ezt követően nyissuk meg a `http://127.0.0.1:64210` URL-t egy böngészővel. A böngésző ablakban meg fog jelenni a {{< cayley >}} web-es felhasználói felülete. A jobb felső területen elhelyezkedő szövegbeviteli területen gépeljük be a `g.V().GetLimit(2)` query kifejezést, majd nyomjuk meg a jobb felső részel látható "Run Query" gombot Ekkor a query végrehajtódik, és az ablak jobb alsó területén megjelenik a keresés eredménye, ahogyan azt a 2. ábrán láthatjuk.

{{< figure src="/cayley-cookbook/cayley-http-ui.png" title="Figure 2.: cayley http UI" >}}

A web-es felület használatának leírása a megtalálható a [hivatalos dokumentációban](https://cayley.gitbook.io/cayley/usage/ui-overview).


## Adat importálás és exportálás

Ha olyankor is szeretnénk adatokat fel és letölteni az adatbázisból, amikor a `cayley http` server éppen fut, akkor azt a REST interfészen keresztül tehetjük meg. létezik két segédprogram, amit a {{< cayley >}} forráskódjában találhatunk: a [`cayleyimport`](https://github.com/cayleygraph/cayley/tree/master/cmd/cayleyimport) és a [`cayleyexport`](https://github.com/cayleygraph/cayley/tree/master/cmd/cayleyexport).

Klónozzuk le a [cayleygraph/cayley](https://github.com/cayleygraph/cayley) repository-t, és akkor tudjuk futtatni ezeket a programokat `go run` paranccsal, ahogyan azt a következő példában lentebb láthatjuk.

Adatok importálása az adatbázisba a REST API-n keresztül:

```bash

    $ go run <path-to-the-cayley-source>/cmd/cayleyimport/cayleyimport.go data/testdata.nq 
    Successfully wrote 15 quads.


```

Adatok exportálása file-ba az adatbázisból a REST API-n keresztül:

```bash

    go run <path-to-the-cayley-source>/cmd/cayleyexport/cayleyexport.go -o out.nq

```

Adatok exportálása az adatbázisból a standard output-ra, a REST API-n keresztül. Mivel ebben az esetben a kimeneti file neve nincs megadva, ezért a program nem tudja meghatározni a kimeneti reprezentációs formátumot. Emiatt azt külön meg kell határoznunk a `--format` argumentummal:


```bash

    $ go run ~/sandbox/cayley/cayley/cmd/cayleyexport/cayleyexport.go --format jsonld
    [{"@id":"alice","follows":[{"@id":"bob"}]},{"@id":"bob","follows":[{"@id":"fred"}],"status":[{"@value":"cool_person"}]},{"@id":"charlie","follows":[{"@id":"bob"},{"@id":"dani"}]},{"@id":"dani","follows":[{"@id":"bob"},{"@id":"greg"}],"status":[{"@value":"cool_person"}]},{"@id":"emily","follows":[{"@id":"fred"}]},{"@id":"fred","follows":[{"@id":"greg"}]},{"@id":"greg","status":[{"@value":"cool_person"}]},{"@id":"predicates","are":[{"@id":"follows"},{"@id":"status"}]},{"@graph":[{"@id":"emily","status":[{"@value":"smart_person"}]},{"@id":"greg","status":[{"@value":"smart_person"}]}],"@id":"smart_graph"}]

```


