---
title: "Namespaces and Vocabularies"
weight: 1
keywords: ["vocabulary", "prefix", "namespace", "foaf"]
# bookFlatSection: false
# bookToc: true
# bookHidden: false
# bookCollapseSection: true
# bookComments: true
---

# Namespaces and Vocabularies

## Namespaces

The following code fragment lists all the namespaces currently registered:

{{% code file="/static/src/namespaces/list_namespaces.go" language="go" %}}

[![Run this code on Repl.it](https://repl.it/badge/github/tombenke/cayley-cookbook-src)](https://repl.it/@tombenke/cayley-cookbook-src#namespaces/list_namespaces.go)

```bash
    cd namespaces
    go run list_namespaces.go`
```

The results:

{{% code file="/static/src/namespaces/list_namespaces.out" language="go" %}}


The next code registers a new namespace:

{{% code file="/static/src/namespaces/register_namespace.go" language="go" %}}

[![Run this code on Repl.it](https://repl.it/badge/github/tombenke/cayley-cookbook-src)](https://repl.it/@tombenke/cayley-cookbook-src#namespaces/register_namespace.go)

```bash
    cd namespaces
    go run register_namespace.go
```

The results:

{{% code file="/static/src/namespaces/register_namespace.out" language="go" %}}


## Vocabularies

The following code demonstrates how can a new vocabulary be defined.
The `foaf.go` package defines the most frequently used terms of the [FOAF Vocabulary Specification](http://xmlns.com/foaf/spec/).
This package is used by several other examples of the cookbook as a vocabulary to use some terms as predicates.

{{% code file="/static/src/kbase/voc/foaf/foaf.go" language="go" %}}

{{< seealso >}}

