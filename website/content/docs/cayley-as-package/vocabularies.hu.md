---
title: "Namespace-ek és szótárak"
weight: 1
keywords: ["vocabulary", "prefix", "namespace", "foaf"]
# bookFlatSection: false
# bookToc: true
# bookHidden: false
# bookCollapseSection: true
# bookComments: true
---

# Namespace-ek és szótárak

## Namespace-ek

Az alábbi kód kilistázza az összes, jelenleg regisztrált namespace-t:

{{% code file="/static/src/namespaces/list_namespaces.go" language="go" %}}

[![Run this code on Repl.it](https://repl.it/badge/github/tombenke/cayley-cookbook-src)](https://repl.it/@tombenke/cayley-cookbook-src#namespaces/list_namespaces.go)

```bash
    cd namespaces
    go run list_namespaces.go`
```

a program kimenete:

{{% code file="/static/src/namespaces/list_namespaces.out" language="go" %}}


Az alábbi kód regisztrál egy új namespace-t:

{{% code file="/static/src/namespaces/register_namespace.go" language="go" %}}

[![Run this code on Repl.it](https://repl.it/badge/github/tombenke/cayley-cookbook-src)](https://repl.it/@tombenke/cayley-cookbook-src#namespaces/register_namespace.go)

```bash
    cd namespaces
    go run register_namespace.go
```

a program kimenete:

{{% code file="/static/src/namespaces/register_namespace.out" language="go" %}}


## Szótárak

A következő kódrészlet bemutatja, hogyan lehet egy új szótárat definiálni.

A `foaf.go` package [FOAF Vocabulary Specification](http://xmlns.com/foaf/spec/) leggyakrabban használt kifejezéseit definiálja.
Ezt a package-et több példában is használjuk. A szótárban definiált szavakat predikátumként alkalmazzuk.

{{% code file="/static/src/kbase/voc/foaf/foaf.go" language="go" %}}

{{< seealso >}}

