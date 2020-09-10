---
title: "Bi-directional Relations"
weight: 3
api: ["graph.V()", "graph.Vertex()", "path.In()", "path.Out()","path.All()"]
# bookFlatSection: false
# bookToc: true
# bookHidden: false
# bookCollapseSection: true
# bookComments: true
---

# Bi-directional Relations

## Problem

If I have a graph: `A <-1-> B <-2-> C` where `A`, `B`, `C` are subjects and objects linked with predicates `1`, `2` (bidirectional).
How  can I get `C` using a path?

## Test data

{{< figure src="/cayley-cookbook/src/query/bidirectional/data.png" title="Fig 1.: Graph diagram of test data" >}}

{{< details title="Show the test data in nquads format" open=false >}}
{{% code file="/static/src/query/bidirectional/data.nq" language="txt" %}}
{{< /details >}}

## Gizmo Query

{{% code file="/static/src/query/bidirectional/query.gizmo" language="javascript" %}}

Results:
{{% code file="/static/src/query/bidirectional/query.out" language="txt" %}}

## Golang Implementation

{{< details title="Show the source code" open=false >}}
{{% code file="/static/src/query/bidirectional/main.go" language="go" %}}
{{< /details >}}

Results:
{{% code file="/static/src/query/bidirectional/main.out" language="txt" %}}


