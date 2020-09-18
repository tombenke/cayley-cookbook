---
title: Cayley Szakácskönyv
type: docs
keywords:
  - cayley
  - cayley-graph
  - knowledge-graph
  - graph-database
  - gremlin
  - tutorial
  - cookbook
  - cook-book
  - go
  - golang
  - gizmo
---

A [Cayley](https://github.com/cayleygraph/cayley) egy open-source Linked Data adatbáziskezelő.
Létrehozását a Google [Knowledge Graph](https://en.wikipedia.org/wiki/Knowledge_Graph) gráf adatbázisa inspirálta.

Ez a weboldal kódrészletek gyűjteménye, melyek azt mutatják be, működő példákon keresztül, hogyan lehet a
[Cayley Knowledge Graph](https://github.com/cayleygraph/cayley) [Golang](https://golang.org/) könyvtárként alkalmazni.

A könyv írásakor az aktuális {{< cayley >}} verzió: `0.7.5`. A példák ezzel a verzióval lettek kipróbálva.

A [github.com/tombenke/cayley-cookbook-src](https://github.com/tombenke/cayley-cookbook-src/tree/master)
repository tartalmazza a forráskódot. A repository-ban az egyes használati esetek directory-kba vannak szervezve.
A gyűjtemény létrehozásának célja az, hogy megkönnyítse a {{< cayley >}} tanulását, használatát.

A példákat kétféle módon használhatjuk:
1. Közvetlenül végrehajthatjuk őket a felhőben.
[![Run on Repl.it](https://repl.it/badge/github/tombenke/cayley-cookbook-src)](https://repl.it/github/tombenke/cayley-cookbook-src)
2. Klónozzuk a [repository](https://github.com/tombenke/cayley-cookbook-src/tree/master)-t a saját gépünkre, és ott futtatjuk.

Ha mélyebb betekintést szeretnénk abba, hogyan működik a {{< cayley >}} rendszer, akkor erősen ajánlott [installálni](https://cayley.gitbook.io/cayley/installation) azt a saját gépünkre, továbbá klónozni a forráskódját tartalmazó repository-kat ([cayleygraph/cayley](https://github.com/cayleygraph/cayley/) és [cayleygraph/quad](https://github.com/cayleygraph/quad/)) is.

<hr />
