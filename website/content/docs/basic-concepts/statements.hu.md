---
weight: 3
keywords: ["triplet","subject","predicate","object","IRI","literal","statement"]
bookCollapseSection: true
title: "Állítások"
---

# Állítások

## Alany, Állítmány, Tárgy

A tudást tény-állításokként rögzítjük.
A tényeket kijelentő mondatokként megfogalmazott állítások formájában írtuk le.
Például:

```txt
Luke Skywalker életkora kora 23 év.
```

angolul:

```txt
Luke Skywalker has age 23.
```


A mondatokat át lehet alakítani olyan formába, hogy azok az `alany-állítmány-tárgy` szerkezetet tükrözzék.
Ezeket a szerkezeti elemeket angolul `subject-predicate-object` névvel azonosítjuk.

```txt
Luke-Skywalker has-age 23.
```
A mondat elemei:
- subject (alany): `Luke-Skywalker`,
- predicate (állítmány): `has-age`,
- object (tárgy): `23`.

Mivel sok Luke Skywalker létezhet, ezért az egyértelműség kedvéért a konkrét személyre a `Luke-Skywalker-from-StarWars` azonosítóval fogunk hivatkozni.
Mivel szabványosított formátumú, univerzális azonosítót akarunk használni, ezért Luke azonosítója a következő IRI lesz: `<Luke-Skywalker-from-StarWars>`.

Luke egyértelműen egy resource, vagy más néven entitás, akit egy IRI-vel egyértelműen azonosítunk.
Van egy tulajdonsága, az életkor, amit a `<has-age>` IRI-vel tudunk kifejezni, mivel ez egy általános fogalom.
Ezen felül szükségünk vagy egy literálra, aminek az értéke `23` és a típusa egész szám. Ezt az értéket veszi fel Luke esetében a `has-age` tulajdonság.
Végleges formában tehát a kijelentés az alábbi lesz:

```txt
    <Luke-Skywalker-from-StarWars> <has-age> 23^^integer .
```

Láthatjuk az 1. ábrán, hogy a `subject-predicate-object` hármas felrajzolható gráfként, ahol a subject és az object node-okként jelenik meg, és a predicate lesz a két node-ot összekötő, irányított él.

{{< figure src="/cayley-cookbook/subject-predicate-object-graph.png" title="1. ábra: Subject-Predicate-Object Gráf" >}}

Az ilyen formában felírt kijelentést __triplet__-nek is nevezik.

Példaképpen vizsgáljuk meg, hogyan tudnánk reprezentálni a számítógép számára az alábbi tényeket:

```txt
Luke Skywalker egy személy a StarWars világban. Családi neve Skywalker, keresztneve Luke.
Luke Skywalker jelenleg 23 éves.
Leia Organa egy személy a StarWars világban. Családi neve Organa, keresztneve Leia.
Leia ismeri Luke-ot.
```

Alakítsuk át a fenti mondatokat `subject-predicat-object` szerkezetűvé.

Nézzük hogyan néz ki átfogalmazva a fenti néhány mondat.
Az átrendezett mondatokat angolul fogjuk leírni, mert ezzel egyszerűbb lesz követni a gráffá alakítás során a jelölésmódot.
A subject-eket és predicate-eket rögtön IRI formátumba alakítottuk.

```txt
<Luke-Skywalker-from-StarWars> <is-a> <Person> .
<Luke-Skywalker-from-StarWars> <has-given-name> "Luke" .
<Luke-Skywalker-from-StarWars> <has-family-name> "Skywalker" .
<Luke-Skywalker-from-StarWars> <has-age> 23 .
<Leia-Organa-from-StarWars> <is-a> <Person> .
<Leia-Organa-from-StarWars> <has-given-name> "Leia" .
<Leia-Organa-from-StarWars> <has-family-name> "Organa" .
<Leia-Organa-from-StarWars> <knows> <Luke-Skywalker-from-StarWars> .
```

Gráfként fogjuk ábrázolni a tényeket, azokkal az elemekkel, amit az előző fejezetben bevezettünk, resource-okkal, IRI-kkel, és literál értékekkel.

A resource-ok (subject-ek és object-ek):
- Luke Skywalker a StarWars világban: `<Luke-Skywalker-from-StarWars>`,
- Leia Organa a StarWars világban: `<Leia-Organa-from-StarWars>`,
- Személy (mint fogalom): `<Person>`.

A predikátumok:
- `<is-a>`,
- `<has-given-name>`,
- `<has-family-name>`,
- `<has-age>`,
- `<knows>`.

A literálok:
- `"Luke"`,
- `"Skywalker"`,
- `"Leia"`,
- `"Organa"`,
- `23`.

A 2. ábra azt szemléltei, hogy a fenti állítások hogyan jelennek meg gráfként:

{{< figure src="/cayley-cookbook/statements-example.png" title="2. ábra: Predikátumok gráfként ábrázolva" >}}

Az entitások sárga színnel, a literálok pedig zöld színnel vannak megjelenítve.

## Üres node-ok

Előfordul olyan eset, amikor az állításokat úgy akarjuk hozzákapcsolni egy entitáshoz, hogy ahhoz nem akarunk globális azonosítót, IRI-t rendelni, mivel az entitásra csak az adott gráfon belül van szükségünk, és annak a kapcsolatai a lényegesek, nem pedig az, hogy a külvilág felé egyértelmű legyen az azonossága.
Ebben az esetben úgynevezett üres csomópontokat, angolul Blank Node-okat alkalmazunk.

A Blank Node-oknak is van saját azonosítója, ami egy adott gráfon belül egyedi ameddig a program fut.

A Blank Node-ok a `_:` prefix-szel kezdődnek, és többnyire valamilyen véletlenszerűen generált számmal, vagy stringgel folytatódnak.

A 3. ábra hasonló a 2. ábrán látható gráfhoz, azzal a különbséggel, hogy azon a két személyt nem azonosítjuk, globális hatáskörű IRI-kkel, hanem csak lokális hatáskörű Blank Node-okkal.

{{< figure src="/cayley-cookbook/blank-node-example.png" title="3. ábra: Predikátumok gráfként ábrázolva, Blank node-okkal" >}}

A Blank Node-okat szürke színnel jelöljük.

{{< seealso >}}

## A node-ok és quad-ok belső reprezentációja

Fontos Megjegyeznünk, hogy valójában, a predikátumok belső reprezentációja szintén gráf csomópont, ami jellemzően IRI, vagy esetenként literál érték lehet. Ezek a predikátumok megjelenhetnek a keresési eredménylistákban is.

A 4. ábra szemlélteti, hogyan néz ki a belső reprezentáció (bal oldal), ahhoz a jelölésmódhoz viszonyítva, amit eddig alkalmaztunk a quad-ok ábrázolására (jobb oldal).

A diagram a következő két quad-ot jeleníti meg:

```bash
 <starwars:leia_organa> <foaf:familyName> "Organa" "people" .
 <starwars:leia_organa> <foaf:knows> _:n353930893927990388 "people" .
```

{{< figure src="/cayley-cookbook/internal-representation.png" title="4. ábra: A quad-ok belső reprezentációja." >}}

A belső reprezentációval kapcsolatban az alábbi megállapításokat tehetjük, az ábra alapján:
- A subject-ek, object-ek ((IRI-k és Blank Node-ok), továbbá a literál értékek mindannyian gáf csomópontokként vannak reprezentálva,
- A predikátumok szintén gráf csomópontokkal vannak reprezentálva,
- A gráf címkék úgyszintén gráf csomópontokkal vannak reprezentálva,
- Minden egyes komplett n-quad kijelentés egy további, kiegészítő gráf csomóponttal van reprezentálva.
