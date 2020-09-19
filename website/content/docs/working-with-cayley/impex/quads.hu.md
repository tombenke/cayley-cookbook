---
title: "Quad-ok exportálása"
weight: 1
bookCollapseSection: true
---

# Quad-ok exportálása

## Probléma

Hogyan tudom az adatbázis tartalmát kiexportálni szabványos formátumú file-okba?

## Megoldás

Az alábbi példa létrehoz egy quad halmazt, majd kiexportálja többféle formátumban:

{{< details title="A teljes kódlista megtekintése" open=false >}}

{{% code file="/static/src/quad/make_quads.go" language="go" %}}
{{< /details >}}

Futtassuk a programot:
[![Run this code on Repl.it](https://repl.it/badge/github/tombenke/cayley-cookbook-src)](https://repl.it/@tombenke/cayley-cookbook-src#quad/writer/export_nquads_to_stdout.go)

```bash
    cd quad/writer
    go run export_nquads_to_stdout.go data
```

A program kimenete:

{{< tabs "results-of-make-quads" >}}

{{< tab "Graphviz" >}}

{{% code file="/static/src/quad/writer/people.dot" language="txt" %}}

{{< /tab >}}

{{< tab "GML" >}}
{{% code file="/static/src/quad/writer/people.gml" language="xml" %}}
{{< /tab >}}

{{< tab "GraphML" >}}
{{% code file="/static/src/quad/writer/people.graphml" language="xml" %}}
{{< /tab >}}

{{< tab "JSON" >}}
{{% code file="/static/src/quad/writer/people.json" language="json" %}}
{{< /tab >}}

{{< tab "n-quads" >}}
{{% code file="/static/src/quad/writer/people.nq" language="txt" %}}
{{< /tab >}}

{{< /tabs >}}

{{< figure src="/cayley-cookbook/src/quad/writer/people_dot.png" title="Graphviz-ből előállított diagram" >}}

{{< figure src="/cayley-cookbook/src/quad/writer/people_graphml.png" title="yEd GraphML-ből előállított diagram" >}}


