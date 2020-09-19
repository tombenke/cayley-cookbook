---
title: "Create a New Vocabulary"
weight: 3
keywords: ["foaf", "vocabulary"]
---

# Create a New Vocabulary

## Problem

I want to create my own vocabulary to the domain I am working with.
I want to use the terms defined by this vocabulary in my Go code, to make it simpler, and more readable.
How can I create and register such a new vocabulary?

## Solution

The following code demonstrates how can a new vocabulary be defined.
The `foaf.go` package defines the most frequently used terms of the [FOAF Vocabulary Specification](http://xmlns.com/foaf/spec/).
This package is used by several other examples of the cookbook as a vocabulary to use some terms as predicates.

{{% code file="/static/src/kbase/voc/foaf/foaf.go" language="go" %}}

{{< seealso >}}

