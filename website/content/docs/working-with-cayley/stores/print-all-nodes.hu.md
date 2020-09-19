---
title: "A store összes node-jának kilistázása"
weight: 3
keywords: ["store", "node"]
---

# A store összes node-jának kilistázása

## Probléma

Hogyan tudom az store-ban lévő összes node-ot lekérdezni, és kilistázni?

## Megoldás

### Teszt adatok előállítása:
{{% code file="/static/src/store/data.go" language="go" %}}

### A store-ban lévő összes node lekérdezése is listázása:
{{% code file="/static/src/store/print_all_nodes.go" language="go" %}}

A program kimenete:
{{% code file="/static/src/store/all_nodes.out" language="txt" %}}

