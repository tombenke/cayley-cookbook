---
title: "Összes node kiírása"
weight: 1
api: ["graph.V()", "graph.Vertex()", "path.All()"]
# bookFlatSection: false
# bookToc: true
# bookHidden: false
# bookCollapseSection: true
# bookComments: true
---

# Összes node kiírása

## Probléma

Hogyan tudom gráf összes elemét (subject-eket, object-eket, predicate-eket) kiíratni? 

## Teszt adatok
{{< figure src="/cayley-cookbook/src/data/testdata.png" title="1. ábra: Teszt adatok Gráf diagramja" >}}

{{< details title="Kiinduló adatok megtekintése nquads formában" open=false >}}
{{% code file="/static/src/data/testdata.nq" language="txt" %}}
{{< /details >}}

## Gizmo query
{{% code file="/static/src/query/find_all_vertices/query.gizmo" language="javascript" %}}

Eredmények:
{{% code file="/static/src/query/find_all_vertices/query.out" language="txt" %}}


