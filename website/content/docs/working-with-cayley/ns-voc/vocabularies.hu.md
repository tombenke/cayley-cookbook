---
title: "Új szótár létrehozása"
weight: 3
keywords: ["foaf", "vocabulary"]
---

# Új szótár létrehozása

## Probléma

Létre akarom hozni a saját szótáramat az adott szakterülethez, amin éppen dolgozom.
Az új szótárban definiált elemeket szeretném használni a Go forráskódban, hogy az tömörebb, és olvashatóbb legyen.

Hogyan tudok saját szótárat létrehozni, és betölteni?

## Megoldás


A következő kódrészlet bemutatja, hogyan lehet egy új szótárat definiálni.

A `foaf.go` package [FOAF Vocabulary Specification](http://xmlns.com/foaf/spec/) leggyakrabban használt kifejezéseit definiálja.
Ezt a package-et több példában is használjuk. A szótárban definiált szavakat predikátumként alkalmazzuk.

{{% code file="/static/src/kbase/voc/foaf/foaf.go" language="go" %}}

{{< seealso >}}

