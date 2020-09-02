---
title: "Erőforrások"
weight: 2
# bookFlatSection: false
# bookToc: true
# bookHidden: false
# bookCollapseSection: true
# bookComments: true
---

# Erőforrások

## Erőforrás

Az erőforrás egy általános fogalom, amelyet a web-en, és a szemantikus web-en gyakran használunk. Lényegében a világban létező konkrét és absztrakt dolgokat értjük alatta. A resource-nak van két fontos tulajdonsága:
- van egyedi azonossága, identitása (identity),
- ábrázolható valamilyen módon (representation).

Tekintetbe véve a fenti két jellemzőt, a resource megfeleltethető a Domain Driven Design entity (entitás) és aggregate fogalmainak. A mi vizsgálatunk szempontjából a két dolgot azonosnak tekinthetjük, a lényeg, hogy egy adott dologot egyértelműen tudjunk azonosítani, akág glóbálisan, akár egy szükebb értelmezési tartományban, vagyis rendelkezzen legalább egy egyedi azonosítóval, valamint le tudjuk írni valamilyen formában.

A leírásnak sokféle formája lehet: szöveges, képi, számítógépes program kód, memóriabeli állapot, stb. Mi a szabványos reprezentációs formátumokra fogunk fókuszálni. Ezek közül is azokra, amelyek a tudás szimbolikus leírására lettek kidolgozva, mint például a szemantikus web számára.

Példák resource-okra:
- konkrét, vagy elképzelt személyek
- tárgyak
- absztrakt fogalmak

Az alábbi kód részlet egy képzeletbeli személyt reprezentál JSON formátumban:

```json
{
    "id": "<star-wars-universe:luke_skywalker>",
    "firstName": "Luke",
    "lastName": "Skywalker",
    "age": 23
}
```

Ugyanez a személy YAML reprezentációs formátumban:

```yaml
---
id: <star-wars-universe:luke_skywalker>
firstName: Luke
lastName: Skywalker
age: 23
```

Két különböző reprezentáció, de ugyanaz a személy.

Tipikusan a resource-ok alkotják a tudásgráf csomópontjait, de gyakran előfordul olyan eset, hogy az entitások között fennáló kapcsolat, gráf él szerepel mint resource, amennyiben rá vonatkozóa akarunk további megállapításokat, szabályokat leírni.


## IRI-k

Az erőforrásokat, másként nevezve az entitásokat azonosítani kell. 
Erre szolgálnak az IRI-k. Az IRI az Internationalized Resource Identifier kifejezés rövidítése.

Az IRI egy Unicode string, amelynek a formátumát az [RFC 3987](https://www.ietf.org/rfc/rfc3987.txt) definiálja, és ami egyedileg azonosít egy entitást. Az IRI stringeket a `<` és `>` karakterek zárják közre, és a karaktereken és számokon kívül numerikus escape szekvenciákat is tartalmazhatnak. formájukban nagyon hasonlítanak az URI-kre, a Uniform Resource Identifier-ekre, mint amilyenek a webről jól ismert URL-ek.

Az IRI-k az URI-k általánosított változatának tekinthetők, Csak abszolút formában léteznek. Minden abszolút URI és URL IRI-ként tekintendő, viszont nem minden IRI számít URI-nek.

Az IRI-k alapértelmezés szerint globális hatáskörrel rendelkeznek. Ennek megfelelően egy IRI két különbőző környezetben történő megjelenése ugyanazt az entitást jelenti.

Minden IRI egy jól meghatározott dolgot, erőforrrást azonosít a tárgyalás körét jelentő világban.
Az erőforrást, amire az IRI hivatkozik, és amit azonosít az azonosítás tárgyának (referent) nevezzük.

Példák IRI-kre:

```nquads
<star-wars-universe:luke_skywalker>
<star-wars-universe:han_solo>
<35345322-543555-534534532>
<https://en.wikipedia.org/wiki/Millennium_Falcon>
```

## Literál értékek

Az erőforrásnak általában vannak tulajdonságai, amelyeket az ábrázolásban leírunk. Ezek a tulajdonságok valamilyen értéket vesznek fel.

Ilyen tulajdonság értékek például egy konkrét személy esetében a személy neve, születési időpontja, stb. a tulajdonságok értékei pedig a konkrét adatok (név, időpont) amiket az említett tulajdonságok felvesznek.

Ezeket a tulajdonság értékeket literál értékeknek, vagy csak egyszerűen literáloknak hívjuk.

A literál értékeknek megadhatjuk a típusát, pl.: string, egész szám, dátum, stb..

A típusokat postfix formában adjuk meg, az értékhez hozzáragasztva a `^^` jelőléssel.

Példák típusokra:
```rdf
"Dantooine"^^xsd:string
"23"^^<xsd:integer>
```

Azt is megadhatjuk, hogy milyen nyelven kell értelmezni az adott literál értéket. Ezt a nyelvi tag-gel, mint postfix-szel adhatjuk meg, a `@` karaktert használva elválasztóként.

Példák nyelvi azonosításra:

```rdf
"people"@en
"emberek"@hu
```

A nyelv azonosítása mellett az írás irányát is lehet, jelezni, és még sok finom részletet tárgyalnak a specifikációk, de ezek ismerete nem szükséges az itt leírtak megértéséhez és a {{< cayley >}} használatához.

{{< button relref="graph-based-data-model" >}}&#9669; Gráf alapú adat modell{{< /button >}}
{{< button relref="statements" >}}Kifejezések &#9659;{{< /button >}}

