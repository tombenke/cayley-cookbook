---
title: "Új namespace regisztrálása"
weight: 1
keywords: ["prefix", "namespace"]
---

# Új namespace regisztrálása

## Probléma

Szükségem van egy új namespace-re, hogyan tudok létrehozni egyet?

## Megoldás

Az alábbi kód regisztrál egy új namespace-t:

{{% code file="/static/src/namespaces/register_namespace.go" language="go" %}}

[![Run this code on Repl.it](https://repl.it/badge/github/tombenke/cayley-cookbook-src)](https://repl.it/@tombenke/cayley-cookbook-src#namespaces/register_namespace.go)

```bash
    cd namespaces
    go run register_namespace.go
```

a program kimenete:

{{% code file="/static/src/namespaces/register_namespace.out" language="go" %}}


