---
title: "Csomópontok keresése minta-illesztéssel"
weight: 4
apifuncs: ["graph.V()", "graph.Vertex()", "path.Filter()", "regex()", "like()", "path.All()"]
usecases: ["filtering"]
bookCollapseSection: true
---

# Csomópontok keresése minta-illesztéssel

## Probléma

Hogyan tudok keresni, vagy szűrni szövegminta-illesztéssel?
Például, hogyan tudom megtalálni azokat a filmeket, amelyek címében megtalálható a `".*Lord of.*"` szövegtöredék?

## Megoldás

### Gizmo Query `regex()`-el
{{% code file="/static/src/query/filter/find_movies_with_regex.gizmo" language="javascript" %}}

Eredmények:
{{% code file="/static/src/query/filter/find_movies_with_regex.out" language="txt" %}}

### Gizmo Query with `like()`-kal
{{% code file="/static/src/query/filter/find_movies_with_like.gizmo" language="javascript" %}}

Eredmények:
{{% code file="/static/src/query/filter/find_movies_with_like.out" language="txt" %}}

