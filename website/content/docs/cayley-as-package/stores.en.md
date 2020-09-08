---
title: "Stores"
weight: 3
# bookFlatSection: false
# bookToc: true
# bookHidden: false
# bookCollapseSection: true
# bookComments: true
---

# Stores

## Test data
{{% code file="/static/src/store/data.go" language="go" %}}

## Nodes in the Store

{{% code file="/static/src/store/print_all_nodes.go" language="go" %}}

The results:

{{% code file="/static/src/store/all_nodes.out" language="txt" %}}

## Quads in the Store
{{% code file="/static/src/store/print_all_quads.go" language="go" %}}

The results:

{{% code file="/static/src/store/all_quads.out" language="txt" %}}

## The internal representation of nodes and quads

{{< figure src="/cayley-cookbook/internal-representation.png" title="Internal Representation" >}}


