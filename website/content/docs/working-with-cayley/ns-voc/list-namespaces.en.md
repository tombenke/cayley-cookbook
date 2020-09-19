---
title: "List Registered Namespaces"
weight: 1
keywords: ["prefix", "namespace"]
---

# List Registered Namespaces

## Problem

I do not know which namespaces are registered currently in the database.
How can I investigate what are these namespaces that I can/have to use?

## Solution

The following code fragment lists all the namespaces currently registered:

{{% code file="/static/src/namespaces/list_namespaces.go" language="go" %}}

[![Run this code on Repl.it](https://repl.it/badge/github/tombenke/cayley-cookbook-src)](https://repl.it/@tombenke/cayley-cookbook-src#namespaces/list_namespaces.go)

```bash
    cd namespaces
    go run list_namespaces.go`
```

The results:

{{% code file="/static/src/namespaces/list_namespaces.out" language="go" %}}

