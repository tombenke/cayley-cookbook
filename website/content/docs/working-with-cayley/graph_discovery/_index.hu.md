---
title: "Ismeretlen gráf felderítése"
weight: 1
bookCollapseSection: true
---

## Probléma:

Hogyan tudok egy ismeretlen gráfban (pl. film adatbázis) lekérdezéseket végrehajtani, ha nem ismerem a szerkezetét?

## Megoldás

Az alábbiakban végrehajtunk egy lehetséges lépéssorozatot, aminek a végén rendelkezni fogunk az ahhoz szükséges ismeretekkel, hogy az ismeretlen gráf tartalmát lekérdezzük, vagy módosítsuk. Mindehhez nincs szükségünk arra, hogy ismerjük a gráfban tárol modell sémáját. Ugyanakkor feltételezzük, hogy a gráf nem véletlenszerűen létrehozott csomópontokból, és kapcsolatokból áll, hanem van benne "rendszer", létezik hozzá legalább egy képzeletbeli séma, és a benne található predikátum elnevezése is követ bizonyos konvenciókat, azokat valamilyen logika mentén hozták létre, és amelyek utalnak azok szerepére.

A felderített sémát egy diagramon fogjuk ábrázolni. Azokat az IRI típusú csomópontokat amiknek ismerjük a konkrét azonosítóját, a diagramon sárgával fogjuk jelölni. A literál típusú csomópontokat zöld színnel, és a típus megnevezésével. Az ismeretlen csomópontokat pedig kék színnel, és egy változó névvel jelöljük.

### 0. lépés: Indítsuk el a REPL-t az adatbázissal

Indítsunk el a REPL paranccsal a {{< cayley >}}-t egy in-memory adatbázissal, amibe rögtön betöltjük az adatokat:

```bash

    $ cayley repl --load data/30kmoviedata.nq.gz 
    I0917 19:12:42.483992   18114 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0917 19:12:42.484240   18114 database.go:187] using backend "memstore"
    I0917 19:12:46.897340   18114 database.go:250] loaded "data/30kmoviedata.nq.gz" in 4.413006871s
    cayley>

```

A továbbiakban ezzel a segédprogrammal fogjuk a feltérképezést végezni.

### 2. lépés: Csoportosítsuk a predikátumokat együttes előfordulás szerint

Azonosítsuk azokat a predikátum csoportokat, amelyek együttesen fordulnak elő.

#### 2.1. lépés: Kimenő élek csoportosítása

Először gyűjtsük össze, csoportonként a kimenő élként együtt előforduló predikátumokat.

```bash

    cayley> g.V().In("</film/performance/actor>").OutPredicates().All()

    ****
    id : </film/performance/actor>
    ****
    id : </film/performance/character>
    -----------
    2 Results
    Elapsed time: 174.893162 ms

```

Az eredmény alapján a `</film/performance/actor>` és a `</film/performance/character>` predikátumok azonos csoportban vannak, tehát többnyire mindannyian ugyanabból a csomópontból indulnak ki.

Ugyanazt az eredményt kell kapnunk az együtt szereplő predikátumok esetében, akármelyikkel is indítjuk a lekérdezést:

```bash

    cayley> g.V().In("</film/performance/character>").OutPredicates().All()

    ****
    id : </film/performance/actor>
    ****
    id : </film/performance/character>
    -----------
    2 Results
    Elapsed time: 45.116628 ms

```

Nézzük a következő predikátumot, ami nem szerepel az imént beazonosított csoportban:

```bash

    cayley> g.V().In("<name>").OutPredicates().All()

    ****
    id : <name>
    ****
    id : <type>
    ****
    id : </film/film/directed_by>
    ****
    id : </film/film/starring>
    -----------
    4 Results
    Elapsed time: 178.340908 ms

```

Sikerült egy másik csoportot azonosítanunk, ami tartalmazza az összes eddig nem besorolt predikátumot.
Ezzel lefedtük az összes predikátumot, tehát két csoportba lehet őket rendezni első közelítésben.

Az együttesen előforduló kimenő éleket rendeljük csomópontokhoz, a csomópontoknak adjunk változó nevet, majd, és rajzoljuk fel az eddig kapott eredményt, ahogy az a 2. ábrán látható:

{{< figure src="/cayley-cookbook/discover_movie_graph_2_1.png" title="2. ábra: A kimenő élként együtt előforduló predikátumok" >}}

#### 2.2. lépés: Bemenő élek csoportosítása

Térképezzük fel a bemenő élként előforduló predikátumokat, és rendeljük hozzá őket azokhoz a csoportokhoz (csomóponti változókhoz) amiket a kimenő élek esetében létrehoztunk. Ehhez nézzük végig minden egyes kimenő predikátumra vonatkozóan, hogy az milyen bejövő predikátumokkal társul. Ehhez hajtsuk végre a következő lekérdezés sorozatot:


```bash

    cayley> g.V().In("</film/performance/actor>").InPredicates().All()

    ****
    id : </film/film/starring>
    -----------
    1 Result
    Elapsed time: 190.712052 ms

```

```bash

    cayley> g.V().In("</film/performance/character>").InPredicates().All()

    ****
    id : </film/film/starring>
    -----------
    1 Result
    Elapsed time: 42.079836 ms

```
 
```bash

    cayley> g.V().In("<type>").InPredicates().All()

    ****
    id : </film/performance/actor>
    ****
    id : </film/film/directed_by>
    -----------
    2 Results
    Elapsed time: 149.433856 ms

```

```bash

    cayley> g.V().In("<film/film/directed_by>").InPredicates().All()

```

ebben az esetben nincs találat.


```bash

    cayley> g.V().In("<film/film/starring>").InPredicates().All()

```

és ebben az esetben sincs.

Helyezzük fel fel a lekérdezések eredményeképpen kapott, bemenő éleket is a gráfra, hozzákapcsolva a csomóponti változókhoz. Az eredményt a 3. ábrán láthatjuk.

{{< figure src="/cayley-cookbook/discover_movie_graph_2_2.png" title="3. ábra: A bemenő és kimenő élekként együtt előforduló predikátumok" >}}

#### Bemenő és kimenő predikátum csoportok egyesítése

Fésüljük össze az eredményhalmazokat, és a gráfot. Az eredményt a 4. ábrán láthatjuk.

{{< figure src="/cayley-cookbook/discover_movie_graph_2_3.png" title="4. ábra: Az együtt előforduló predikátumok egyesített ábrája." >}}

{{< hint info >}}
__Fontos megjegyzés:__

Az eredményül kapott `x` és `y` node változók nem feltétlenül egyetlen node-típust reprezentálnak, mert a keresés során a bejövő és kimenő predikátumokat közösítettük.

{{< /hint >}}

### 3. lépés: Keressük meg a terminális csomópontokat

Ezek azok a csomópontok, akikből már nem indul ki további predikátum.

Egyenként vizsgáljuk meg a kimenő predikátumokat, hogy azok milyen csomópontokhoz érkeznek.

Kezdjük a `</film/performance/character>` predikátummal:

```bash

    cayley> g.V().Out("</film/performance/character>").All()

    ****
    id : Luther Krank
    ****
    id : Roland
    ****
    id : Tomás de Torquemada
    ****
    id : Ferdinand VII of Spain
    ****
    id : Christopher Columbus
    ...

```

Az eredménylista kiírása az első 100 elem után befejeződik, alapértelmezés szerint.
Látható, hogy az eredmények string literál értékek, a filmben szereplő karakterek nevei.

Nézzük meg, hogy hányan vannak:

```bash

    cayley> g.V().Out("</film/performance/character>").Count()

    => 15058
    -----------
    1 Result
    Elapsed time: 15.394561 ms

```

Most ellenőrizzük, hogy indul-e belőlük további predikátum:

```bash

    cayley> g.V().Out("</film/performance/character>").Out().Count()

     => 0
    -----------
    1 Result
    Elapsed time: 65.014376 ms

```

Az eredmény 0, tehát ezek terminális csomópontok, mivel további kimenő kapcsolatok már nincsenek.

Most vizsgáljuk meg a `<name>` predikátumot:

```bash

    cayley> g.V().Out("<name>").All()

    ****
    id : David Fisher
    ****
    id : 002 Operazione Luna
    ****
    id : 008: Operation Exterminate
    ****
    id : 02:37:00 AM
    ****
    id : 06/05
    ****
    id : 10,000 BC
    ****
    ...

```

Ezek string literálok.

```bash

    cayley> g.V().Out("<name>").Count()

    => 74950
    -----------
    1 Result
    Elapsed time: 44.480211 ms

```

A számuk 74950.

```bash

    cayley> g.V().Out("<name>").Out().Count()

    => 0
    -----------
    1 Result
    Elapsed time: 52.163808 ms

```

Ezek szintén terminális csomópontok.

Folytassuk a sort a `<type>` predikátummal:

```bash

    cayley> g.V().Out("<type>").All()

    ****
    id : </people/person>
    ****
    id : </film/film>
    ****
    id : </film/film>
    ...

```

Ebből több is van, és nem literálok, hanem IRI-k.

Állapítsuk meg egyedileg, hogy pontosan melyek ezek az IRI-k:


```bash

    cayley> g.V().Out("<type>").Unique().All()
     
    ****
    id : </people/person>
    ****
    id : </film/film>
    -----------
    2 Results
    Elapsed time: 43.759725 ms

```

Összesen két ilyen IRI létezik, a `</people/person>` és a `</film/film>`.
Ellenőrizzük ezek esetében is, hogy terminális csomópontok-e, vagy sem:

```bash

    cayley> g.V().Out("<type>").Out().Count()

    => 0
    -----------
    1 Result
    Elapsed time: 48.857924 ms

```

Terminálisok.

Vezessük fel az újonnan felfedezett literál és IRI csomópontokat a gráfra, a hozzájuk vezető predikátumokkal együtt, ahogyan az 5. ábrán látható.


{{< figure src="/cayley-cookbook/discover_movie_graph_3_1.png" title="5. ábra: Terminális csomópontokkal kiegészített séma." >}}

### 4. válasszuk szét a különálló csomópont típusokat

Lehetőség szerint válasszuk szét azokat a csomópont változókat, amelyek hasonló beérkező és kimenő predikátumokat fogadnak, de eltérő típusú entitásokat azonosítanak.

#### Az `y` csomóponti változó elemzése

Az 5. ábrán megfigyelhetjük, hogy az `y` csomópont `<type>` kimenő irányú predikátuma, a következő értékeket veheti fel: `</people/person>` és `</film/film>`. Ebből következően feltételezhetjük, hogy `y` csomópont több különböző entitás típust fedhet. Válasszuk szét a feltételezésnek megfelelően az `y` csomópontot kétfelé, a 6. ábra szerint:

{{< figure src="/cayley-cookbook/discover_movie_graph_4_1.png" title="6. ábra: A person és a film csomópont változók szétválasztása." >}}

Alkalmazzunk több (2, 3, ...n) lépésből álló kapcsolati láncokat, a további elemzéshez.


Vizsgáljuk meg, hogy milyen predikátumok indul és érkeznek `y`-ból származtatott, szétválasztott csomópontokhoz, ha azok értelmezését úgy korlátozzuk, hogy a `<type>` egyszer az egyik, másszor a másik értéket veszi fel:

```bash

    cayley> y_film = g.V().Has("<type>", "</film/film>")

    => [internal Iterator]
    -----------
    1 Result
    Elapsed time: 0.469018 ms

    cayley> y_film.OutPredicates().All()

    ****
    id : </film/film/directed_by>
    ****
    id : </film/film/starring>
    ****
    id : <name>
    ****
    id : <type>
    -----------
    4 Results
    Elapsed time: 113.54646 ms

```

A nevükből ítélve ezek a predikátumok egy film esetében, normálisnak tűnnek.

Most vizsgáljuk meg a person típusú csomópont jelöltek esetében, hogy mi a helyzet:


```bash

    cayley> y_person = g.V().Has("<type>", "</people/person>")

    => [internal Iterator]
    -----------
    1 Result
    Elapsed time: 0.552285 ms

    cayley> y_person.OutPredicates().All()

    ****
    id : <name>
    ****
    id : <type>
    ****
    id : </film/film/directed_by>
    ****
    id : </film/film/starring>
    -----------
    4 Results
    Elapsed time: 95.020065 ms

```

Ugyanazt az eredményt kaptuk, ami meglepő, mert egy személy esetében nehéz értelmezni mind a `</film/film/directed_by>`, mind pedig a `</film/film/starring>` kimenő predikátumokat.

Az eredmény gyanús, ezért vizsgáljuk meg kicsit tüzetesebben azt.

A filmeknek feltehetően van rendezője, tehát értelmes lehet erre vonatkozó kapcsolatokat vizsgálni:

```bash

    cayley> y_film.Out("</film/film/directed_by>").Count()

    => 33310
    -----------
    1 Result
    Elapsed time: 108.365305 ms

```

Ilyenből találunk 33310-et, és ez logikusnak látszik.

Ellenőrizzük, hogy ezek a predikátumok valóban személyekre mutatnak-e!

```bash

    cayley> y_film.Out("</film/film/directed_by>").Has("<type>", "</people/person>").Count()

    => 33310
    -----------
    1 Result
    Elapsed time: 168.508677 ms

```

A számosság alapján ítélve igen. Nézzünk meg néhányat név szerint is:

```bash

    cayley> y_film.Out("</film/film/directed_by>").Has("<type>", "</people/person>").Out("<name>").GetLimit(5)

    ****
    id : Lucio Fulci
    ****
    id : Umberto Lenzi
    ****
    id : Murali K. Thalluri
    ****
    id : Theo van Gogh
    ****
    id : Roland Emmerich
    -----------
    5 Results
    Elapsed time: 1.071004 ms

```

Úgy tűnik, hogy ez a kapcsolat és csomópont típus rendben van.

Most nézzük meg, hogy a `</people/person>` típusú csomópontoknak van-e "rendezte" kapcsolata:

```bash

    cayley> y_person.Out("</film/film/directed_by>").Count()

    => 6
    -----------
    1 Result
    Elapsed time: 96.554532 ms

```

Meglepő, de van 6 ilyen.
Nézzük meg ezt a 6 személyt részletesebben, tag-ekkel kiegészítve, hogy pontosan lássuk mi is a helyzet:

```bash

    cayley> y_person.Tag("person").Out("</film/film/directed_by>").Tag("directed_by").All()

    ****
    directed_by : </en/quentin_tarantino>
    id : </en/quentin_tarantino>
    person : </en/death_proof>
    ****
    directed_by : </en/mysterio>
    id : </en/mysterio>
    person : </en/jazmin>
    ****
    directed_by : </en/robert_rodriguez>
    id : </en/robert_rodriguez>
    person : </en/planet_terror>
    ****
    directed_by : </en/keenen_ivory_wayans>
    id : </en/keenen_ivory_wayans>
    person : </en/scary_movie_2>
    ****
    directed_by : </en/david_zucker>
    id : </en/david_zucker>
    person : </en/scary_movie_3>
    ****
    directed_by : </en/ralph_bakshi>
    id : </en/ralph_bakshi>
    person : </en/the_lord_of_the_rings_1978>
    -----------
    6 Results
    Elapsed time: 95.293387 ms

```
Az eredmény listából látható, hogy ennél a 6 csomópontnál a `<directed_by>` predikátum valóban személyekre mutat, de a kiinduló csomópontok IRI-jéből látszik, hogy azok valójában nem személyek, hanem filmek. Tehát igen nagy valószínűséggel állíthatjuk, hogy ebben a hat esetben a filmre vonatkozó IRI kimenő `<type>` predikátumai hibásan a `</people/person>` csomópontra mutat, a `</film/film>` csomópont helyett. Ezért ezt a kapcsolatot nem tekintjük érvényesnek, bár az adatbázis kétség kívül tartalmaz ilyen éleket.

A 7. ábra a korrigált séma rajzát szemlélteti, amire felvezettük a legutóbbi megállapításainkat, azaz a `y_film` típusú csomópont `</film/film/directed_by>` predikátuma `y_person` típusú csomópontra mutat, a `y_person`-nak pedig nincs ilyen kimenő predikátuma.

{{< figure src="/cayley-cookbook/discover_movie_graph_4_2.png" title="7. ábra: A `y_film` csomópont kimenő predikátumai." >}}


Most nézzük meg mi a helyzet a `</film/film/starring>` predikátummal.

```bash

    cayley> y_film.Out("</film/film/starring>").Count()

    => 136737
    -----------
    1 Result
    Elapsed time: 140.230179 ms

    cayley> y_film.Out("</film/film/starring>").GetLimit(5)

    ****
    id : _:117646
    ****
    id : _:117647
    ****
    id : _:117648
    ****
    id : _:117649
    ****
    id : _:117240
    -----------
    5 Results
    Elapsed time: 1.052197 ms

```

A `y_film` típusú csomópontnak van 136737 ilyen kapcsolata, amelyek Blank Node-ok.

Nézzük meg, hogy ezek a "starring" node-ok milyen bemenő és kimenő predikátumokkal rendelkeznek:

```bash

    cayley> y_film.Out("</film/film/starring>").OutPredicates().All()

    ****
    id : </film/performance/actor>
    ****
    id : </film/performance/character>
    -----------
    2 Results
    Elapsed time: 255.873653 ms

    cayley> y_film.Out("</film/film/starring>").InPredicates().All()

    ****
    id : </film/film/starring>
    -----------
    1 Result
    Elapsed time: 242.836008 ms

```

A "starring"-ra utaló Blank node-ok bejövő és kimenő predikátumai megegyeznek azzal a csomópont változóval, amit a séma diagramon `x`-szel jelöltünk.

Ellenőrzésképpen nézzük meg, hogy tudunk-e filmekhez szerepeket és konkrét színészeket találni, a predikátumok segítségével:

```bash

    cayley> y_film = g.V().Has("<type>", "</film/film>")

    => [internal Iterator]
    -----------
    1 Result
    Elapsed time: 0.897193 ms

    cayley> filmWithTitle = y_film.Tag("film").Out("<name>").Tag("filmTitle").Back("film")

    => [internal Iterator]
    -----------
    1 Result
    Elapsed time: 0.680148 ms

    cayley> filmStar = filmWithTitle.Out("</film/film/starring>").Tag("star")

    => [internal Iterator]
    -----------
    1 Result
    Elapsed time: 0.340275 ms

    cayley> {t("</film/performance/character>").Tag("character").Back("star").Out("</film/performance/actor>").Out("<name>").Tag("actorName").GetLimit("10")

    ****
    actorName : Katherine Heigl
    character : Arlene
    film : </en/100_girls>
    filmTitle : 100 Girls
    id : Katherine Heigl
    star : _:239
    ****
    actorName : Joely Richardson
    character : Anita
    film : </en/101_dalmatians>
    filmTitle : 101 Dalmatians
    id : Joely Richardson
    star : _:457
    ****
    actorName : Susanne Blakeslee
    character : Cruella de Vil
    film : </en/101_dalmatians_ii_patchs_london_adventure>
    filmTitle : 101 Dalmatians II: Patch's London Adventure
    id : Susanne Blakeslee
    star : _:71373
    ****
    actorName : Brian Markinson
    character : Daniel
    film : </en/10_5_apocalypse>
    filmTitle : 10.5: Apocalypse
    id : Brian Markinson
    star : _:218
    ****
    actorName : Heath Ledger
    character : Patrick Verona
    film : </en/10_things_i_hate_about_you>
    filmTitle : 10 Things I Hate about You
    id : Heath Ledger
    star : _:503
    ****
    actorName : James Marsden
    character : Tommy
    film : </en/10th_wolf>
    filmTitle : 10th & Wolf
    id : James Marsden
    star : _:284
    ****
    actorName : Jeana Tomasino
    character : Karen
    film : </en/10_to_midnight>
    filmTitle : 10 to Midnight
    id : Jeana Tomasino
    star : _:100658
    ****
    actorName : Jason Segel
    character : Leon (paramedic 1)
    film : </en/11_14>
    filmTitle : 11:14
    id : Jason Segel
    star : _:108561
    ****
    actorName : Henry Fonda
    character : Juror #8
    film : </en/12_angry_men>
    filmTitle : 12 Angry Men
    id : Henry Fonda
    star : _:462
    ****
    actorName : Lee J. Cobb
    character : Juror #3
    film : </en/12_angry_men>
    filmTitle : 12 Angry Men
    id : Lee J. Cobb
    star : _:463
    -----------
    10 Results
    Elapsed time: 7.233902 ms

```

Ez határozottan működőképesnek tűnik.

Tehát kimondhatjuk, hogy az `x` csomópont változó képviseli a film-sztárokat.

Azt is kijelenthetjük, hogy az `x` kimenő predikátumai közül mind a `y_person`-ra mutató `</film/performance/actor>` predikátum, és a `</film/performance/character>` literál értékre mutató predikátum is valós kapcsolatokat jelölnek.

Felmerül a kérdés, hogy az `x`-ből mutat-e `</film/performance/actor>` predikátum a `y_film` csomópont típus irányába?

```bash

    filmStar.Out("</film/performance/actor>").Has("<type>", "</film/film>").Count()

    => 1
    -----------
    1 Result
    Elapsed time: 1315.917655 ms

```

Fura módon ebből is találunk egyet.

Ki is ez a film-sztár, és kire mutat, mint színész a kapcsolat?

```bash

    cayley> filmStar.Out("</film/performance/actor>").Has("<type>", "</film/film>").All()

    ****
    film : </en/jazmin>
    filmTitle : Jazmin
    id : </en/jazmin>
    star : _:23742
    -----------
    1 Result
    Elapsed time: 1340.16577 ms

```

Jól látható, hogy ennek a filmsztárnak az azonosítója a `_:23742` Blank Node, ami színészként személy helyett egy filmre mutat. Nézzük meg, hogy személyre vonatkozóan is létezik-e ilyen kapcsolata:

```bash

    cayley> g.V("_:23742").Out("</film/performance/actor>").Has("<type>", "</people/person").Count()

    => 0
    -----------
    1 Result
    Elapsed time: 0.68618 ms

```

Nincs ilyen kapcsolata. Kijelenthetjük, hogy egy újabb hibát fedeztünk fel az adatbázisban, és átvezethetjük a tapasztalatainkat a séma rajzra, a 8. ábrának megfelelően.

{{< figure src="/cayley-cookbook/discover_movie_graph_4_3.png" title="8. ábra: Az `y_person` és `x_character` csomópont típusok kapcsolatai." >}}


Végül elemezzük az utolsó kapcsolatot, amit még nem vizsgáltunk meg eddig. Az `y_person`-ból kiinduló `</film/film/starring>` predikátumot.

```bash

    cayley> y_person = g.V().Has("<type>", "</people/person>").Out("</film/film/starring>").Count()

    => 96
    -----------
    1 Result
    Elapsed time: 85.149164 ms

```

Meglepő, de ebből is találunk 96 darabot. Járjunk végére, hogy hová mutatnak ezek a kapcsolatok!

```bash

cayley> y_person = g.V().Has("<type>", "</people/person>").Out("</film/film/starring>").GetLimit(5)

****
id : _:95530
****
id : _:95531
****
id : _:95532
****
id : _:95533
****
id : _:95534
-----------
5 Results
Elapsed time: 12.859301 ms

```

Jól látható, hogy Blank node-ok, tehát valószínűleg `x_character` típusú csomópontokra mutatnak.

Hajtsunk végre egy olyan lekérdezést, ahol személy típusú csomópontból ("originPerson") kiindulva, a `</film/film/starring>` kapcsolaton keresztül továbbmegy a színész ("actor") irányába, és nézzük meg a lánc végén álló csomópont nevét és típusát. Minden kiinduló csomópontot csak egyszer jelenítsen meg.

```bash

    cayley> g.V().Tag("originPerson").Has("<type>", "</people/person>").Out("<type>").Tag("originType").Back("originPerson").Out("</film/film/starring>").Out("</film/performance/actor>").Tag("actor").Out("<type>").Tag("actorType").Back("person").All()

    ****
    actor : </en/jazmin>
    actorType : </film/film>
    id : </en/jazmin>
    originPerson : </en/jazmin>
    originType : </film/film>
    ****
    actor : </en/jazmin>
    actorType : </people/person>
    id : </en/jazmin>
    originPerson : </en/jazmin>
    originType : </film/film>
    ****
    actor : </en/jazmin>
    actorType : </people/person>
    id : </en/jazmin>
    originPerson : </en/jazmin>
    originType : </people/person>
    ****

    ...


    ****
    actor : </en/anthony_daniels>
    actorType : </people/person>
    id : </en/the_lord_of_the_rings_1978>
    originPerson : </en/the_lord_of_the_rings_1978>
    originType : </people/person>
    ****
    actor : </guid/9202a8c04000641f800000000112e07f>
    actorType : </people/person>
    id : </en/the_lord_of_the_rings_1978>
    originPerson : </en/the_lord_of_the_rings_1978>
    originType : </film/film>
    -----------
    100 Results
    Elapsed time: 965.035817 ms
```

Ha végignézzük az eredménylistát, láthatjuk, hogy ismét sikerült felfedeznünk ellentmondásos kapcsolatokat az adatbázisban. Többféle hibát is felfedezhetünk. Gyanítható, hogy bizonyos `y` egyszerre rendelkeznek `</film/film>` és `</people/person>` csomópontokra mutató `<type>` predikátumokkal.

Ellenőrizzük, hogy feltevésünk megfelel-e a valóságnak:

```bash

    cayley> g.V().Has("<type>", "</people/person>").And(g.V().Has("<type>", "</film/film>")).All()

    ****
    id : </en/death_proof>
    ****
    id : </en/jazmin>
    ****
    id : </en/planet_terror>
    ****
    id : </en/scary_movie_2>
    ****
    id : </en/scary_movie_3>
    ****
    id : </en/the_lord_of_the_rings_1978>
    -----------
    6 Results
    Elapsed time: 72.999153 ms

```

Sajnos feltevésünk helyesnek bizonyult. Ezek azonban egyértelműen hibák. Kijelenthetjük tehát, hogy a `person` típusú csomópontból, normális esetben nem indul ki `</film/film/starring>` predikátum sem a `character`, sem más csomópont irányába.

A végleges séma tehát, ami csak az értelmes kapcsolatokat szemlélteti, a 9. ábrán látható.

{{< figure src="/cayley-cookbook/discover_movie_graph_4_4.png" title="9. ábra: A film adatbázis sémája." >}}

