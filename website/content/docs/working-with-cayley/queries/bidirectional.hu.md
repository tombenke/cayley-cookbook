---
title: "Kétirányú relációk"
weight: 3
apifuncs: ["graph.V()", "graph.Vertex()", "path.In()", "path.Out()","path.All()"]
usecases: ["subgraph-matching"]
# bookFlatSection: false
# bookToc: true
# bookHidden: false
# bookCollapseSection: true
# bookComments: true
---

# Kétirányú relációk

## Probléma

Ha van egy `A <-1-> B <-2-> C` gráfunk ahol `A`, `B`, `C` subject-ek, és object-ek az `1`, `2` predikátumokkal vannak összekapcsolva mindkét irányban (bidirectional),
akkor hogyan tudjuk `C`-t elérni `A`-ból egy path segítségével?

## Teszt adatok

{{< figure src="/cayley-cookbook/src/query/bidirectional/data.png" title="Fig 1.: A tesz adat gráfja" >}}

{{< details title="A teszt adatok megjelenítése in nquads formátumban" open=false >}}
{{% code file="/static/src/query/bidirectional/data.nq" language="txt" %}}
{{< /details >}}

## Gizmo Query

### Egyszerű Query
{{% code file="/static/src/query/bidirectional/query.gizmo" language="javascript" %}}

Kimenet:
{{% code file="/static/src/query/bidirectional/query.out" language="txt" %}}

### Query kétirányú predicate alkalmazásával
{{% code file="/static/src/query/bidirectional/query_with_both.gizmo" language="javascript" %}}

Kimenet:
{{% code file="/static/src/query/bidirectional/query_with_both.out" language="txt" %}}

### Query morphism alkalmazásával
{{% code file="/static/src/query/bidirectional/query_with_morphism.gizmo" language="javascript" %}}

Kimenet:
{{% code file="/static/src/query/bidirectional/query_with_morphism.out" language="txt" %}}

## Golang implementáció

{{< details title="A forráskód megjelenítése" open=false >}}
{{% code file="/static/src/query/bidirectional/main.go" language="go" %}}
{{< /details >}}

Results:
{{% code file="/static/src/query/bidirectional/main.out" language="txt" %}}


