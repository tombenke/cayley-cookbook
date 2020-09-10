---
title: "Összes predikátum kiírása"
weight: 2
api: ["graph.V()", "graph.Vertex()", "path.OutPredicates()", "path.Union", "path.InPredicates()", "path.All()"]
# bookFlatSection: false
# bookToc: true
# bookHidden: false
# bookCollapseSection: true
# bookComments: true
---

# Összes predikátum kiírása

## Probléma

Hogyan tudom kilistázni az összes predikátumot, ami szerepel a gráfban?

## Teszt adatok
{{< figure src="/cayley-cookbook/src/data/testdata.png" title="1. ábra: Teszt adatok Gráf diagramja" >}}

{{< details title="Kiinduló adatok megtekintése nquads formában" open=false >}}
{{% code file="/static/src/data/testdata.nq" language="txt" %}}
{{< /details >}}

## Gizmo query
{{% code file="/static/src/query/find_all_predicates/query.gizmo" language="javascript" %}}

Eredmények:
{{% code file="/static/src/query/find_all_predicates/query.out" language="txt" %}}


