---
title: "Find nodes by pattern matching"
weight: 4
apifuncs: ["graph.V()", "graph.Vertex()", "path.Filter()", "regex()", "like()", "path.All()"]
usecases: ["filtering"]
bookCollapseSection: true
---

# Find nodes by pattern matching

## Problem

How can I find nodes by pattern matching?
For example, how can I find all movies with title that has the `".*Lord of.*"` pattern?

## Solution

### Gizmo Query with `regex()`
{{% code file="/static/src/query/filter/find_movies_with_regex.gizmo" language="javascript" %}}

Results:
{{% code file="/static/src/query/filter/find_movies_with_regex.out" language="txt" %}}

### Gizmo Query with `like()`
{{% code file="/static/src/query/filter/find_movies_with_like.gizmo" language="javascript" %}}

Results:
{{% code file="/static/src/query/filter/find_movies_with_like.out" language="txt" %}}

