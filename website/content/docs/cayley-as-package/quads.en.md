---
title: "Quads"
weight: 2
# bookFlatSection: false
# bookToc: true
# bookHidden: false
# bookCollapseSection: true
# bookComments: true
---

# Quads

This code creates a set of quads, then exports them into several formats:

{{< details title="See the full list of source code" open=false >}}

{{% code file="/static/src/quad/make_quads.go" language="go" %}}
{{< /details >}}

[![Run this code on Repl.it](https://repl.it/badge/github/tombenke/cayley-cokbook)](https://repl.it/@tombenke/cayley-cookbook-1#quad/writer/export_nquads_to_stdout.go)

```bash
    cd quad/writer
    go run export_nquads_to_stdout.go data
```


The results:

{{< tabs "results-of-make-quads" >}}

{{< tab "Graphviz" >}}

{{% code file="/static/src/quad/writer/people.dot" language="graphviz" %}}

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
{{% code file="/static/src/quad/writer/people.nq" language="nquads" %}}
{{< /tab >}}

{{< /tabs >}}

{{< figure src="/cayley-cookbook/src/quad/writer/people_dot.png" title="Rendered Graphviz" >}}

{{< figure src="/cayley-cookbook/src/quad/writer/people_graphml.png" title="Rendered yEd GraphML" >}}


