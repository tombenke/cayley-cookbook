---
title: "Find All Vertices"
weight: 1
api: ["graph.V()", "graph.Vertex()", "path.All()"]
# bookFlatSection: false
# bookToc: true
# bookHidden: false
# bookCollapseSection: true
# bookComments: true
---

# Find All Vertices

## Problem

How can I list all the nodes (subjects, predicates and objects) of the graph?

## Test Data
{{< figure src="/cayley-cookbook/src/data/testdata.png" title="Fig 1.: Graph diagram of test data" >}}

{{< details title="Show the test data in nquads format" open=false >}}
{{% code file="/static/src/data/testdata.nq" language="txt" %}}
{{< /details >}}

## Gizmo Query
{{% code file="/static/src/query/find_all_vertices/query.gizmo" language="javascript" %}}

Results:
{{% code file="/static/src/query/find_all_vertices/query.out" language="txt" %}}


