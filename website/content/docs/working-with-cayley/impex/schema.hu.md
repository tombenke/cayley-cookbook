---
title: "Adat import / export struct annotációval"
weight: 2
bookCollapseSection: true
---

# Adat import / export struct annotációval

## Struktúra annotációk

Ez a fejezet arra mutat példát, hogyan lehet egyszerű vagy összetett objektumokat, azok tulajdonságait quad-okká alakítani, és feltölteni a store-ba, az adatstruktúra deklarációjában elhelyezett annotációk segítségével. Továbbá hogyan lehet a store-ból, az eredmény listákból egyszerűen összeszedni az objektumok property-jeit, és objektumként visszakapni az eredményeket.

Ezzel a módszerrel jelentősen egyszerűbb az összetett adat-szerkezetek store-ba való feltöltése és lekérdezése, nem kell saját implementációt készíteni arra például, hogy egy adott Blank Node-dal, vagy IRI-vel azonosított objektum minden egyes tulajdonság-értékét egyenként szedegessük össze.

## Store létrehozása

A példákban egyszerű, in-memory store-t fogunk használni, aminek a megnyitását a `store.go` file-ban implementált `InitStore()` funkció végzi el.

{{< details title="A store.go megjelenítése" open=false >}}
{{% code file="/static/src/kbase/store.go" language="go" %}}
{{< /details >}}

## Teszt adatok

A `bookmarks.yml` file YAML formátumban reprezentált teszt adatokat tartalmaz. Könyvjelzők listája található benne.

Ezeket a könyvjelzőket fogjuk feltölteni, és lekérdezni a séma annotációk segítségével.

{{< details title="A bookmarks.yml megjelenítése" open=false >}}
{{% code file="/static/src/schema/bookmarks.yml" language="yaml" %}}
{{< /details >}}

Az adatok file-ból történő beolvasására a `yaml.go`-ban található `ReadFromYaml()` funkciót használjuk.

A `yaml.go` két függvény tartalmaz:
- A `ReadFromYaml()` az adatok YAML file-ból történő beolvasását végzi el,
- a `SaveToYaml()` a YAML formátumba történő kiírását teszi lehetővé.

{{< details title="A yaml.go megjelenítése" open=false >}}
{{% code file="/static/src/kbase/impex/yaml.go" language="go" %}}
{{< /details >}}

## Struktúra annotációk definiálása

A könyvjelző objektumok struktúráját, valamint a séma definícióját, az importáló és exportáló műveleteket mind a `bookmarks.go` file tartalmazza.

{{< details title="A bookmarks.go megjelenítése" open=false >}}
{{% code file="/static/src/schema/bookmarks.go" language="go" %}}
{{< /details >}}

A forráskódban látható, hogy kétféle struktúra van definiálva:
- a `Bookmark` a teljes változat,
- a `BookmarkShort` pedig egy rövidített változat.

A `struct`-ban vannak elhelyezve a `yaml` és a `quad` cimkével azonosított annotációk, tehát egyszerre minkét típusú átalakításhoz megadhatjuk az instrukciókat.

## Adat importálás struktúra annotációk alapján

Az adatokat először beolvassuk egy `[]Bookmarks` típusú változóba, majd abból feltöltjük a store-ba. A feltöltést a `"github.com/cayleygraph/cayley/schema"` package-ben definiált, `schema.WriteAsQuads()` funkcióval végezzük el, ahogyan az `import_bookmarks_with_schema.go` file-ban látható.

{{< details title="Az import_bookmarks_with_schema.go megjelenítése" open=false >}}
{{% code file="/static/src/schema/import_bookmarks_with_schema.go" language="go" %}}
{{< /details >}}

Futtassuk a programot:
[![Run this code on Repl.it](https://repl.it/badge/github/tombenke/cayley-cookbook-src)](https://repl.it/@tombenke/cayley-cookbook-src#schema/import_bookmarks_with_schema.go)

```bash
    cd schema
    go run impex_bookmarks_with_schema.go yamlImpex.go store.go bookmarks.go
```

A program kimenete:

{{% code file="/static/src/schema/import_bookmarks_with_schema.out" language="txt" %}}

## Adat exportálás struktúra annotációk alapján

Az exportálás ugyanúgy kezdődik, mint az importálás. Először beolvassuk az adatokat a YAML file-ból, majd feltöltjük a store-ba.

Ezt követően lekérdezzük az össze Bookmarkot a store-ból, kétféle módon:
- A `GetAllBookmarks()` funció a teljes részletességgel fogja lekérdezni az adatokat,
- a `GetAllShortBookmarks()` pedig az egyszerűsített változatban.

Mindkét funkció a `bookmarks.go` file-ban van implementálva, és a `schemaConfig.LoadTo()` metódust használja a tulajdonságok összegyűjtésére. Ehhez előzőleg létre kell hoznunk egy `schemaConfig` objektumot, a `schema.NewConfig()` funció segítségével.

{{< details title="Az impex_bookmarks_with_schema.go megjelenítése" open=false >}}
{{% code file="/static/src/schema/impex_bookmarks_with_schema.go" language="go" %}}
{{< /details >}}

Futtassuk a programot:
[![Run this code on Repl.it](https://repl.it/badge/github/tombenke/cayley-cookbook-src)](https://repl.it/@tombenke/cayley-cookbook-src#schema/impex_bookmarks_with_schema.go)

```bash
    cd schema
    go run impex_bookmarks_with_schema.go yamlImpex.go store.go bookmarks.go
```

A lekérdezésekeket követően a program kiírja az eredményeket.

A program kimenete:

{{< details title="Az eredmények megjelenítése" open=false >}}
{{% code file="/static/src/schema/impex_bookmarks_with_schema.out" language="txt" %}}
{{< /details >}}

{{< seealso >}}
