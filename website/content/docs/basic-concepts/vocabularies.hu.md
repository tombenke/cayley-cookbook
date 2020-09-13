---
weight: 4
keywords: ["vocabulary", "prefix", "namespace", "foaf"]
bookCollapseSection: true
title: "Szótárak"
---

# Szótárak

## Szótár
A tudásgráfok alkalmazásának egyik fő előnye, hogy a segítségükkel leírt tudás-területek ismeretanyagát a szimbolikus térben tárgyalja, vagyis olyan kódolási, és reprezentációs formákat alkalmaz, ami az emberek számára is érthetővé teszi a fogalmak, dolgok egyértelmű azonosítását.

Amikor egy entitást megnevezünk, vagy állításokat teszünk, predikátumok alkalmazásával, a gépnek és az embernek ugyanazt a jelentéstartalmat kell értenie a kifejezések mögött. Ahhoz, hogy a kifejezések jelentése is egyértelmű legyen, szótárakat hozunk létre. A szótárak definiálják mindazokat a azonosítókat és predikátumokat, amelyek a szótár hatáskörébe tartoznak. A szótárak a legtöbb esetben IRI-ket sorolnak fel, azok jelentésével és alkalmazási lehetőségeivel, szabályaival.

A szótárakban szereplő IRI-k általában angol szavak, vagy rövidítések. Ugyanaz a szó akár egy entitás azonosítására használjuk, akár predikátumként egy viszony vagy tulajdonság kifejezésére, különböző szövegkörnyezetekben mást és mást jelenthet. Pl. az a szó, hogy robot egy autó-összeszerelő ipari üzemben szerelő robotot fog jelenteni, de egy sci-fi novellában már egy intelligens droidot.

Az eddig bemutatott predikátumaink, mint `<knows>`, `<Person>`, tehát nem eléggé precízek, és egyértelműek.

Azért, hogy az IRI-k globális érvényessége ne sérüljön, és, hogy a jelentés egyértelmű legyen, az egyes szakterületekre külön-külön szótárakat hoznak létre. Ezek a szótárak lehetnek általános fogalmakat (`ember`, `barát`, `ismer`, `szülő`, `gyerek`, stb.) leíró szótárak, és fókuszálhatnak egy szűk területre, pl. gyógyszerkutatás. Az is lehetséges, hogy ugyanaz a szó több szótárban is szerepel, hasonló, vagy némileg eltérő jelentéssel.

Jó példa egy szótárra a [FOAF](http://xmlns.com/foaf/spec/), ami jól használható emberek és szervezetek profiljának leírására.

Amikor használjuk az IRI-ket, akkor azt is jelezzük, hogy melyik szótár értelmezésében értjük azt. Ezt úgy tehetjük meg, hogy megadjuk a szótár azonosítóját, ami egy egyedi string, többnyire egy URL. Ezt a szótár azonosítót az IRI elejére illesztjük.

Ha most átalakítjuk a predicate-ként használt IRI-jeinket, úgy, hogy azok megfeleljenek a FOAF szótárnak, akkor az alábbi formát fogják ölteni:

```txt
<http://xmlns.com/foaf/0.1/#age>
<http://xmlns.com/foaf/0.1/#familyName>
<http://xmlns.com/foaf/0.1/#givenName>
<http://xmlns.com/foaf/0.1/#Person>
```

## Névtartományok

Az előző fejezetben bevezetett szótár azonosító string (pl.: `http://xmlns.com/foaf/0.1/#`) egyértelműen meghatároz egy név-tartományt, angolul __namespace__-t, amiben pontosan definiáltuk a szó azonosságát, és jelentését. Az IRI-k ebben a hosszú formában nem túl kényelmesen olvashatóak ember számára, ezért azokban a reprezentációkban, amiket szerkeszteni, vagy olvasni akarunk, és nem csak a gépek olvasnak, lehet egy rövidebb változatot is alkalmazni, amit __prefix__-nek nevezünk. A FOAF szótár esetében jellemzően a `foaf:` prefixet használjuk, tehát:
- namespace: `http://xmlns.com/foaf/0.1/`
- prefix: `foaf`

A prefix-es változatban a predikátumaink a következőképpen festenek:

```txt
<foaf:age>
<foaf:familyName>
<foaf:givenName>
<foaf:Person>
```

A prefixes és a teljes namespace-szel leírt formák azonosak egymással. A különbség az, hogy a teljes változatot bárhol használhatjuk, mert az a teljes információt és azonosságot magában hordozza. A prefix-es változatnál minden esetben jeleznünk kell az adott reprezentációban (forrás file), a prefix definícióját. Erre minden reprezentációs formátumnak (RDF, Turtle, stb.) ami megengedi a prefix-ek alkalmazását meghatározott nyelvi elemei vannak.

Fontos szótárak, amelyek alapértelmezésben a {{< cayley >}}-ben is definiálva vannak:

| prefix | namespace |
|-|-|
| `owl` | `http://www.w3.org/2002/07/owl#` |
| `rdf` | `http://www.w3.org/1999/02/22-rdf-syntax-ns#` |
| `rdfs` | `http://www.w3.org/2000/01/rdf-schema#` |
| `schema` | `http://schema.org/` |
| `xsd` | `http://www.w3.org/2001/XMLSchema#` |

Mi magunk is létre tudunk hozni namespace-eket és prefix-eket. Például a StarWars univerzum számára bevezethetjük a következő párost:

- namespace: `http://starwars.universe/#`
- prefix: `starwars`

és ezen túl, ebben tudjuk értelmezni a resource-ainkat:

```txt
<starwars:Luke-Skywalker>
<starwars:Leia-Organa>
```
A szótárakat egy adott tudásreprezentációs file-on, ill. adatbázison belül kombináltan használhatjuk egymással, a kijelentésekben. A prefixekkel kiegészített kijelentéseink az alábbiak lesznek:

```txt
<starwars:Luke-Skywalker> <rdf:type> <Person> .
<starwars:Luke-Skywalker> <foaf:givenName> "Luke" .
<starwars:Luke-Skywalker> <foaf:familyName> "Skywalker" .
<starwars:Luke-Skywalker> <foaf:age> 23 .
<starwars:Leia-Organa> <rdf:type> <Person> .
<starwars:Leia-Organa> <foaf:givenName> "Leia" .
<starwars:Leia-Organa> <foaf:familyName> "Organa" .
<starwars:Leia-Organa> <knows> <starwars:Luke-Skywalker> .
```

## Ontológiák

A teljesség kedvéért fontos megemlíteni, hogy a szótárak az esetek jelentős részében jóval több mindent tartalmaznak, mint a nevek listáját, és szöveges értelmezését.

A szemantikus web technológiájával olyan erőforrásokat és predikátumokat is használunk, amelyek meta információt tartalmaznak más erőforrásokról, és predikátumokról. Pl. a predikátumoknak, mint relációknak leírhatóak a tulajdonságai (reflexív, tranzitív, stb.). Ezek olyan tudás-leírások, amelyek a következtetés végrehajtásához szükséges meta információkat is tartalmazzák, és megfelelő eszköz, pl. következtető gép segítségével olyan adatok is kinyerhetőek ezekből a gráfokból, amelyek explicit módon nincsenek leírva, de implicit módon ki lehet következtetni őket. Az ilyen leírásokat ontológiáknak nevezzük.

Például ha a `<foaf:knows>` kapcsolatról tudjuk, hogy az reflexív, akkor elegendő azt kijelenteni, hogy

```txt
<starwars:Luke_Skywalker> <foaf:knows> <starwars:Leia_Organa> .
```

Megfelelő séma alkalmazása esetén, ebből a kijelentésből, a logika szabályai alapján, kikövetkeztethető, hogy az alábbi állítás is igaz:

```txt
<starwars:Leia_Organa> <foaf:knows> <starwars:Luke_Skywalker> .
```

{{< seealso >}}

