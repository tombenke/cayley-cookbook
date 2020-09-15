---
title: "A Cayley rendszer áttekintése"
weight: 3
bookCollapseSection: true
---

Ez a fejezet, abban segít eligazodni, hogy milyen elemekből áll a Cayley "ökoszisztéma", és hogyan lehet az egyes elemeket praktikusan felhasználni.

A {{< cayley >}} komponensek használatára vonatkozó információk a [hivatalos dokumentációs oldalakon](https://cayley.gitbook.io/cayley/) olvashatóak.

## Installation

Ez a szakácskönyv egy Linux-ot futtató gépen készült. A {{< cayley >}} installálásának a legegyszerűbb módja egy linuxos gépen, a `snap` használata:

```bash

    snap install --edge --devmode cayley

```

## A Cayley használata

A {{< cayley >}} kétféle módon használható:
- mint önálló bináris program (`cayley`),
- mint program könyvtár, beépítve egy másik alkalmazásba, amit Go nyelven írtak.

### A cayley beágyazása Go könyvtárként

[A Tudás-Gráf használata]({{< relref "/docs/working-with-cayley" >}}) című fejezet számos példát hoz arra vonatkozóan, hogyan lehet a {{< cayley >}}-t Go library-ként használni.

### A cayley mint önálló alkalmazás használata
Az önálló bináris alkalmazás több parancsot is támogat. A help lekérésével kilistázhatjuk ezeket a parancsokat:

```bash

    $ cayley -h
    Cayley is a graph store and graph query layer.

    Usage:
      cayley [command]

    Available Commands:
      convert     Convert quad files between supported formats.
      dedup       Deduplicate bnode values
      dump        Bulk-dump the database into a quad file.
      help        Help about any command
      http        Serve an HTTP endpoint on the given host and port.
      init        Create an empty database.
      load        Bulk-load a quad file into the database.
      query       Run a query in a specified database and print results.
      repl        Drop into a REPL of the given query language.
      upgrade     Upgrade Cayley database to current supported format.
      version     Prints the version of Cayley.

```

Az alábbi parancsok használhatóak arra, hogy elindítsuk a triple-store-t, és a lekérdező motort, különféle üzemmódokban, úgymint:
- HTTP server: `cayley http`.
- REPL server: `cayley repl`.
- Query server: `cayley query`.

A `cayley` alkalmazás, adatmanipulációs célokra is használható, a következő parancsokkal:
- Üres, perzisztens adatbázis létrehozása: `cayley init`,
- Adatok betöltése egy (perzisztens) adatbázisba, file-ból: `cayley load`,
- Adatok kiíratása file-ba, egy (perzisztens) adatbázisból: `cayley dump`.
- Erőforrás leíró file-ok konvertálása különféle formátumok között: `cayley conv`.

A soron következő alfejezetekben röviden ismertetjük, hogyan lehet az imént felsorolt parancsokat használni.

### In-memory vs. perzisztens store
A {{< cayley >}} egy ún. triple-store, ami lehetővé teszi n-quads-ok és/vagy n-triplet-ek betöltését egy adatbázisba, valamint lekérdezések végrehajtását az adatbázis tartalmán. Többféle lekérdező nyelvet támogat. Jelen szakácskönyv az ún. Gizmo query nyelvre koncentrál. Az adatbázis lehet in-memory, vagy perzisztens.

Az adatbázis, amiben a {{< cayley >}} a quad-okat tárolja, in-memory vagy perzisztens adatbázis lehet, amire leginkább a _store_ elnevezést használjuk. Valójában a {{< cayley >}} többféle adatbázist képes alkalmazni abból a célból, hogy a quad-okat tároljuk. Az alapértelmezett adatbázis a [Bolt key-value store](https://github.com/boltdb/bolt), de kicserélhetjük a konfigurációs beállítások segítségével más típusúra. Többféle, különböző típusú (key-value, NoSQL és SQL) adatbázis közül választhatunk. Az ide vonatkozó részletek a [konfigurációs beállítások](https://cayley.gitbook.io/cayley/configuration) oldalon találhatóak, a hivatalos dokumentációban.

Az in-memory store-t mind az önálló alkalmazásként futó `cayley`, mind pedig a beágyazott könyvtár azonnal képes használni. Az in-memory store az alapbeállítás szerinti adatbázis. A perzisztens store-okat előzőleg inicializálnunk kell, hogy használhassuk.

Amikor az önálló {{< cayley >}} bináris alkalmazást használjuk, inicializálhatunk egy új adatbázist az `--init` kapcsolóval, vagy megtehetjük ezt külön is, a `cayley init` parancs segítségével, mielőtt egy másik paranccsal használatba vennénk az új adatbázist.

Amikor egy perzisztens store-t akarunk használni, meg kell mondanunk a `cayley` parancsnak, hogy milyen típusút szeretnénk. Ez egyaránt megtehetjük parancssori argumentumokkal, vagy egy konfigurációs file alkalmazásával

A konfigurációs file-ok YAML vagy JSON formátumúak lehetnek. Az alábbi példa a Bolt key-value store használatát írja elő. A `cayley` a `/home/tombenke/cayley/` folder-be fogja tenni a file-okat:

{{% code file="/static/src/config/bolt.json" language="json" %}}

{{< hint info >}}
__Megjegyzés:__

Amikor egy `cayley` parancsot futtatunk, az kizárólagosan fogja használni a perzisztens store-t, tehát nem lehetséges, hogy például adatokat töltsünk fel egy adatbázisba a `load` paranccsal, miközben azt egy másik parancs, pl. a `http` használja. Ezért két lehetőségünk van az adatbázisok manipulálására és használatára többféle paranccsal:
- Meghatározott sorrendben, egymás után, külön-külön használjuk a `cayley` parancsokat, pl.: `init`, `store`, `http`, `dump`,
- Használjuk a REST API-t, amit a `cayley http` biztosít, és ami lehetővé teszi, hogy manipuláljuk az adatbázisban lévő adatokat, miközben lekérdezéseket is végrehajtunk.

{{< /hint >}}

