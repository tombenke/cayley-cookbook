---
title: "Regisztrált Namespace-ek listázása"
weight: 1
keywords: ["prefix", "namespace"]
---

# Regisztrált Namespace-ek listázása

## Probléma

Nem tudom, hogy jelenleg milyen nammespace-ek vannak érvényben az adatbázisban.
Hogyan tudom kideríteni, hogy melyek azok amiket használhatok, vagy szükséges használnom?

## Megoldás

Az alábbi kód kilistázza az összes, jelenleg regisztrált namespace-t:

{{% code file="/static/src/namespaces/list_namespaces.go" language="go" %}}

[![Run this code on Repl.it](https://repl.it/badge/github/tombenke/cayley-cookbook-src)](https://repl.it/@tombenke/cayley-cookbook-src#namespaces/list_namespaces.go)

```bash
    cd namespaces
    go run list_namespaces.go`
```

A program kimenete:

{{% code file="/static/src/namespaces/list_namespaces.out" language="go" %}}

