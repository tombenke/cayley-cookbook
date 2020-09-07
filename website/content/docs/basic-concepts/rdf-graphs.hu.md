---
title: "Tudás Gráf reprezentációk"
weight: 5
# bookFlatSection: false
# bookToc: true
# bookHidden: false
# bookCollapseSection: true
# bookComments: true
---

# Tudás Gráf reprezentációk

Az előző fejezetekben bemutattuk, hogyan lehet tényállításokat kifejezni resource-ok, predikátumok, és literál értékek alkalmazásával. Ezeket a tényállítás hármasokat tripleteknek is nevezzük.

Bevezettük az IRI fogalmát, ami alkalmas arra, hogy a tényállításokat formalizáltan, általános érvénnyel írjuk le. Rendelkezésünkre áll tehát egy egyszerű módszer arra, hogy a tudásgráf csomópontjait és éleit, valamint azok címkéit szövegszerkesztővel szerkeszthető, egyszerű text formátumba alakítsuk, file-okban tároljuk, és ezen file-ok tartalmát betöltsük a tudásgráfot megvalósító adatbázisba, a számítógép memóriájába, továbbá, hogy onnan kinyerjük, és visszatároljuk file-ba.

Létezik tehát a tudásgráfnak egy olyan reprezentációja, ami a számítógép memóriájában helyezkedik el, és olyan is, ami valamilyen szabványosított formátumú file-ban. Ezekből a formátumokból számos létezik. Az alábbiakban azokat a formátumokat soroljuk fel, amelyekkel a {{< cayley >}} használata során találkozni fogunk.

A leggyakoribb formátumok általában valamilyen szöveg alapú (text, XML) formátumok, amiknek szabványos mime azonosítójuk is van.

| formátum | kiterjesztés | mime-type | használata |
| --- | --- | --- | --- |
| n-quads | `nq`, `nt` | `TBD` | tbd. |


## n-quads
