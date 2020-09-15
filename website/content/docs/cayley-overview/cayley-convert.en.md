---
title: "The `cayley convert` command"
weight: 1
bookCollapseSection: true
---

# The `cayley convert` command

We can convert knowledge-content files from one representation format to another with the `cayley convert` command.

{{< figure src="/cayley-cookbook/cayley-convert.png" title="2. ábra: cayley convert" >}}


The following command converts the test data from n-quads format to JSON-LD format:

```bash

    $ cayley convert --load data/testdata.nq --dump data/testdata.jsonld

```

The next command below does the same conversion, but the results will be written to the standard output:

```bash

    $ cayley convert --load data/testdata.nq --dump_format jsonld --dump -
    I0915 19:15:58.328789   30572 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0915 19:15:58.328974   30572 dump.go:20] writing quads to stdout
    I0915 19:15:58.328989   30572 convert.go:100] reading "data/testdata.nq"
    [{"@id":"alice","follows":[{"@id":"bob"}]},{"@id":"bob","follows":[{"@id":"fred"}],"status":[{"@value":"cool_person"}]},{"@id":"charlie","follows":[{"@id":"bob"},{"@id":"dani"}]},{"@id":"dani","follows":[{"@id":"bob"},{"@id":"greg"}],"status":[{"@value":"cool_person"}]},{"@id":"emily","follows":[{"@id":"fred"}]},{"@id":"fred","follows":[{"@id":"greg"}]},{"@id":"greg","status":[{"@value":"cool_person"}]},{"@id":"predicates","are":[{"@id":"follows"},{"@id":"status"}]},{"@graph":[{"@id":"emily","status":[{"@value":"smart_person"}]},{"@id":"greg","status":[{"@value":"smart_person"}]}],"@id":"smart_graph"}]
    [{"@id":"alice","follows":[{"@id":"bob"}]},{"@id":"bob","follows":[{"@id":"fred"}],"status":[{"@value":"cool_person"}]},{"@id":"charlie","follows":[{"@id":"bob"},{"@id":"dani"}]},{"@id":"dani","follows":[{"@id":"bob"},{"@id":"greg"}],"status":[{"@value":"cool_person"}]},{"@id":"emily","follows":[{"@id":"fred"}]},{"@id":"fred","follows":[{"@id":"greg"}]},{"@id":"greg","status":[{"@value":"cool_person"}]},{"@id":"predicates","are":[{"@id":"follows"},{"@id":"status"}]},{"@graph":[{"@id":"emily","status":[{"@value":"smart_person"}]},{"@id":"greg","status":[{"@value":"smart_person"}]}],"@id":"smart_graph"}]

```
