---
weight: 5
bookCollapseSection: true
keywords: ["n-quads","n-triples","triplet","rdf"]
title: "Tudás Gráf reprezentációk"
---

# Tudás Gráf reprezentációk

Az előző fejezetekben bemutattuk, hogyan lehet tényállításokat kifejezni resource-ok, predikátumok, és literál értékek alkalmazásával. Ezeket a tényállítás hármasokat tripleteknek is nevezzük.

Bevezettük az IRI fogalmát, ami alkalmas arra, hogy a tényállításokat formalizáltan, általános érvénnyel írjuk le. Rendelkezésünkre áll tehát egy egyszerű módszer arra, hogy a tudásgráf csomópontjait és éleit, valamint azok címkéit szövegszerkesztővel szerkeszthető, egyszerű text formátumba alakítsuk, file-okban tároljuk, és ezen file-ok tartalmát betöltsük a tudásgráfot megvalósító adatbázisba, a számítógép memóriájába, továbbá, hogy onnan kinyerjük, és visszatároljuk file-ba.

Létezik tehát a tudásgráfnak egy olyan reprezentációja, ami a számítógép memóriájában helyezkedik el, és olyan is, ami valamilyen szabványosított formátumú file-ban kerül tárolásra. Ezekből a file formátumokból számos létezik. Az alábbiakban azokat a formátumokat soroljuk fel, amelyekkel a {{< cayley >}} használata során találkozni fogunk.

Az erőforrások (resource-ok) formális ábrázolását definiáló keretrendszert RDF-nek nevezzük. Az RDF a [Resource Description Framework](https://en.wikipedia.org/wiki/Resource_Description_Framework) kifejezés rövidítése.

Az RDF többféle file-szerializációs formátumot is definiál. A leggyakoribb formátumok általában szöveg alapúak (text, XML), amiknek szabványos mime azonosítójuk is van.

Az alábbi táblázat összefoglalja a leggyakrabban használatos formátumokat.

{{< hint info >}}
__Megjegyzés:__

A táblázat csak a teljesség kedvéért sorolja fel a formátumokat. Számunkra első közelítésben, szinte kizárólag az `n-quads`, `n-triples` és a `Turtle` formátumok fontosak. A többit nem szükséges mélyebben tanulmányozni, ahhoz, hogy a szakácskönyv példáit megértsük, és hogy a {{< cayley >}}-t használhassuk.
{{< /hint >}}

| formátum | kiterjesztés | mime-type | mikor használjuk? |
| --- | --- | --- | --- |
| [n-triples](https://www.w3.org/TR/n-triples/) | `.nt` | `application/n-triples` | Ha elfogadható performanciát, és magas szintű kompatibilitást akarunk biztosítani. |
| [n-quads](https://www.w3.org/TR/n-quads/) | `.nq` | `application/n-quads` | ua. mint n-triples |
| [Turtle](https://www.w3.org/TR/turtle/) | `.ttl` | `application/x-turtle` | Ha manuálisan kell szerkeszteni, olvasni. |
| [N3](https://www.w3.org/TeamSubmission/n3/) | `.n3` | `text/n3` | Ha szükségünk van RDF szabályok alkalmazására. |
| [RDF/XML](https://www.w3.org/TR/rdf-syntax-grammar/) | `.xml` | `application/rdf+xml` | Ha XML-t kell használnunk. |
| [JSON-LD](https://www.w3.org/2018/jsonld-cg-reports/json-ld/) | `.jsonld` | `application/ld+json` | Ha JSON API-t akarunk biztosítani, vagy ahhoz csatlakozunk, és nincs nagy performancia-igény. |

## n-triples, n-quads

Ez a két formátum szinte azonos. Az `n-triples` lényegében nem más, mint amit korábban a kijelentések c. fejezetben bemutattunk. Minden kijelentés `subject predicate object` formában, külön sorban szerepel, amit egy `.` karakter zár le. Például:

```txt
<Luke-Skywalker-from-StarWars> <is-a> <Person> .
<Luke-Skywalker-from-StarWars> <has-given-name> "Luke" .
<Luke-Skywalker-from-StarWars> <has-family-name> "Skywalker" .
<Luke-Skywalker-from-StarWars> <has-age> 23 .
```

az `n-quads` formátum ettől annyiban tér el, hogy az `object` és a záró `.` közé még elhelyezünk egy opcionális `label` stringet. Ez a `label` (címke), arra jó, hogy az egy csoportba (gráfba) tartozó kijelentéseket megjelölhessük vele. Ezzel az adatbázis-kezelő performanciáját javíthatjuk, ha a keresést csak meghatározott gráfokra szűkítjük, az összes feltöltött gráf helyett.

Például az alábbi kódrészlet `n-quads` kijelentéseket tartalmaz, a `"a-new-hope"` címkével kiegészítve:

```txt
<Luke-Skywalker-from-StarWars> <is-a> <Person> "a-new-hope" .
<Luke-Skywalker-from-StarWars> <has-given-name> "Luke" "a-new-hope" .
<Luke-Skywalker-from-StarWars> <has-family-name> "Skywalker" "a-new-hope" .
<Luke-Skywalker-from-StarWars> <has-age> 23 "a-new-hope" .
```

Ez a címke az epizódra utal, amiben a szereplők először feltűntek.

A következő kódrészlet a Dantooine bolygót írja le `n-triples` formátumban. Ezt a többi formátummal való összehasonlítás kedvéért hoztuk létre. Ugyanezen bolygó-leírást láthatjuk a további fejezetekben, többféle formátumban megjelenítve.

```txt
<https://swapi.co/resource/planet/25> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <https://swapi.co/vocabulary/Planet> .
<https://swapi.co/resource/planet/25> <http://www.w3.org/2000/01/rdf-schema#label> "Dantooine"^^<http://www.w3.org/2001/XMLSchema#string> .
<https://swapi.co/resource/planet/25> <http://www.ontotext.com/business-object/type> "Planet" .
<https://swapi.co/resource/planet/25> <https://swapi.co/vocabulary/climate> "temperate"^^<http://www.w3.org/2001/XMLSchema#string> .
<https://swapi.co/resource/planet/25> <https://swapi.co/vocabulary/desc> "None"^^<http://www.w3.org/2001/XMLSchema#string> .
<https://swapi.co/resource/planet/25> <https://swapi.co/vocabulary/diameter> "9830"^^<http://www.w3.org/2001/XMLSchema#integer> .
<https://swapi.co/resource/planet/25> <https://swapi.co/vocabulary/gravity> "1 standard"^^<http://www.w3.org/2001/XMLSchema#string> .
<https://swapi.co/resource/planet/25> <https://swapi.co/vocabulary/orbitalPeriod> "378"^^<http://www.w3.org/2001/XMLSchema#integer> .
<https://swapi.co/resource/planet/25> <https://swapi.co/vocabulary/population> "1000"^^<http://www.w3.org/2001/XMLSchema#integer> .
<https://swapi.co/resource/planet/25> <https://swapi.co/vocabulary/rotationPeriod> "25"^^<http://www.w3.org/2001/XMLSchema#integer> .
<https://swapi.co/resource/planet/25> <https://swapi.co/vocabulary/terrain> "oceans, savannas, mountains, grasslands"^^<http://www.w3.org/2001/XMLSchema#string> .
```

## Turtle

A Turle egy nagyon népszerű formátum. Nagyon hasonlít az `n-quads`-ra, de sokkal tömörebb. A file elején definiálhatjuk a prefixeket és namespace-eket, és a kijelentéseinkben a subject-et elegendő csak egyszer megadni, majd ezt követhetik a predicate-object párosok, amelyeket a `;` karakter zárja le, majd a teljes kijelentéshalmazt a `.` zárja le.

A Turtle formátum további könnyítéseket is tartalmaz, amiből adódóan kiválóan alkalmazható olyan esetben, ahol olvasnunk, vagy kézzel kell írnunk a tartalmat.

A Dantooine bolygó leírása Turtle formátumban:

```txt
@prefix bo: <http://www.ontotext.com/business-object/> .
@prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
@prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .
@prefix voc: <https://swapi.co/vocabulary/> .
@prefix xml: <http://www.w3.org/XML/1998/namespace> .
@prefix xsd: <http://www.w3.org/2001/XMLSchema#> .

<https://swapi.co/resource/planet/25> a voc:Planet ;
    rdfs:label "Dantooine"^^xsd:string ;
    bo:type "Planet" ;
    voc:climate "temperate"^^xsd:string ;
    voc:desc "None"^^xsd:string ;
    voc:diameter 9830 ;
    voc:gravity "1 standard"^^xsd:string ;
    voc:orbitalPeriod 378 ;
    voc:population 1000 ;
    voc:rotationPeriod 25 ;
    voc:terrain "oceans, savannas, mountains, grasslands"^^xsd:string .
```

## N3

Az N3-at, más néven Notation3 formátumot Tim-Berners Lee dolgozta ki, az eredeti RDF/XML formátum helyettesítésére. Az N3 a szemantikus információ teljes körű leírására alkalmas. Egyenértékű az RDF/XML formátummal, de annál sokkal jobban olvasható, átlátható emberi szem számára. Az N3 nagyon hasonlít a Turtle-re, a Turtle valójában az N3 egy részhalmaza.

A Dantooine bolygó leírása N3 formátumban:

```txt
@prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .
@prefix xsd: <http://www.w3.org/2001/XMLSchema#> .
@prefix ns0: <http://www.ontotext.com/business-object/> .
@prefix ns1: <https://swapi.co/vocabulary/> .

<https://swapi.co/resource/planet/25>
  a <https://swapi.co/vocabulary/Planet> ;
  rdfs:label "Dantooine"^^xsd:string ;
  ns0:type "Planet" ;
  ns1:climate "temperate"^^xsd:string ;
  ns1:desc "None"^^xsd:string ;
  ns1:diameter 9830 ;
  ns1:gravity "1 standard"^^xsd:string ;
  ns1:orbitalPeriod 378 ;
  ns1:population 1000 ;
  ns1:rotationPeriod 25 ;
  ns1:terrain "oceans, savannas, mountains, grasslands"^^xsd:string .
```

## RDF/XML

Az RDF/XML az a szerializációs formátum, amit legelőször hoztak létre, és ami teljes körűen lefedi a szemantikus web tudás-reprezentációval szemben támasztott igényeit.

Sajnos az RDF/XML meglehetősen bőbeszédű, ember számára nem túl jól olvasható. Leginkább akkor használjuk, ha az XML formátum alkalmazása szükséges.

A Dantooine bolygó leírása XML/RDF formátumban:

```xml
<?xml version="1.0" encoding="utf-8" ?>
<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
         xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#"
         xmlns:ns0="http://www.ontotext.com/business-object/"
         xmlns:ns1="https://swapi.co/vocabulary/">

  <rdf:Description rdf:about="https://swapi.co/resource/planet/25">
    <rdf:type rdf:resource="https://swapi.co/vocabulary/Planet"/>
    <rdfs:label rdf:datatype="http://www.w3.org/2001/XMLSchema#string">Dantooine</rdfs:label>
    <ns0:type>Planet</ns0:type>
    <ns1:climate rdf:datatype="http://www.w3.org/2001/XMLSchema#string">temperate</ns1:climate>
    <ns1:desc rdf:datatype="http://www.w3.org/2001/XMLSchema#string">None</ns1:desc>
    <ns1:diameter rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">9830</ns1:diameter>
    <ns1:gravity rdf:datatype="http://www.w3.org/2001/XMLSchema#string">1 standard</ns1:gravity>
    <ns1:orbitalPeriod rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">378</ns1:orbitalPeriod>
    <ns1:population rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">1000</ns1:population>
    <ns1:rotationPeriod rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">25</ns1:rotationPeriod>
    <ns1:terrain rdf:datatype="http://www.w3.org/2001/XMLSchema#string">oceans, savannas, mountains, grasslands</ns1:terrain>
  </rdf:Description>

</rdf:RDF>
```

## JSON-LD

A JSON-LD formátumú reprezentáció egy szabályos JSON dokumentum. A JSON-t elterjedten alkalmazzák REST API-ok request/response formátumaként. Jól olvasható ember számára is. Mivel mérete meglehetősen nagy, a feldolgozása sok erőforrást igényelhet, ezért nem a legoptimálisabb formátum, ha a feldolgozási sebesség fontos szempont.

A Dantooine bolygó leírása JSON-LD formátumban:

```json
[
    {
        "@id": "https://swapi.co/resource/planet/25",
        "@type": [
            "https://swapi.co/vocabulary/Planet"
        ],
        "http://www.ontotext.com/business-object/type": [
            {
                "@value": "Planet"
            }
        ],
        "http://www.w3.org/2000/01/rdf-schema#label": [
            {
                "@value": "Dantooine"
            }
        ],
        "https://swapi.co/vocabulary/climate": [
            {
                "@value": "temperate"
            }
        ],
        "https://swapi.co/vocabulary/desc": [
            {
                "@value": "None"
            }
        ],
        "https://swapi.co/vocabulary/diameter": [
            {
                "@value": 9830
            }
        ],
        "https://swapi.co/vocabulary/gravity": [
            {
                "@value": "1 standard"
            }
        ],
        "https://swapi.co/vocabulary/orbitalPeriod": [
            {
                "@value": 378
            }
        ],
        "https://swapi.co/vocabulary/population": [
            {
                "@value": 1000
            }
        ],
        "https://swapi.co/vocabulary/rotationPeriod": [
            {
                "@value": 25
            }
        ],
        "https://swapi.co/vocabulary/terrain": [
            {
                "@value": "oceans, savannas, mountains, grasslands"
            }
        ]
    },
    {
        "@id": "https://swapi.co/vocabulary/Planet"
    }
]
```
