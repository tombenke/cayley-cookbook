---
title: "Store-ok"
weight: 3
# bookFlatSection: false
# bookToc: true
# bookHidden: false
# bookCollapseSection: true
# bookComments: true
---

# Store-ok

## Teszt adatok előállítása
{{% code file="/static/src/store/data.go" language="go" %}}

## Node-ok a Store-ban

{{% code file="/static/src/store/print_all_nodes.go" language="go" %}}

A program kimenete:

{{% code file="/static/src/store/all_nodes.out" language="txt" %}}

## Quad-ok a Store-ban
{{% code file="/static/src/store/print_all_quads.go" language="go" %}}

A program kimenete:

{{% code file="/static/src/store/all_quads.out" language="txt" %}}

## A node-ok és quad-ok belső reprezentációja

{{< figure src="/cayley-cookbook/internal-representation.png" title="Belső reprezentáció" >}}


