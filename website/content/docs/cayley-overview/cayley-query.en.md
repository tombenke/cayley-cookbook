---
title: "The `cayley query` command"
weight: 6
bookCollapseSection: true
---

# The `cayley query` command

With the `cayley query` we can run a query expression on an existing database.

The query that we want to execute can be forwarded to the command through the standard input. The simpler expressions we can directly put into the command line via the `echo` utility, the longer ones we'd better put into query files, then forward the via the `cat` utility.

We can select the query language type by the `--lang <query-language-type>` parameter, where the possible values for the query language are: "gizmo", "graphql", "mql", "sexp". The default is "gizmo".

{{< figure src="/cayley-cookbook/cayley-query.png" title="Figure 4.: cayley query" >}}

Run a query on a previously created database that has data:
```bash

    echo "g.V().All();" | cayley query --logtostderr false --db bolt --dbpath /home/tombenke/tmp/cayley --logs ~/tmp
    {"id":"bob"}
    {"id":"status"}
    {"id":"cool_person"}
    {"id":"alice"}
    {"id":"greg"}
    {"id":"emily"}
    {"id":"smart_graph"}
    {"id":"predicates"}
    {"id":"dani"}
    {"id":"fred"}
    {"id":"smart_person"}
    {"id":"charlie"}
    {"id":"are"}
    {"id":"follows"}

```

In case we want to feed the results into another program, then we may want to hide the logs printed out by the command. The next example shows how can we do this. The logs are written into the `/home/tombenke/tmp` directory:

```bash

    $ cat ~/tmp/cayley.tombenke-Latitude-E5470.tombenke.log.INFO.20200915-182850.27351 
    Log file created at: 2020/09/15 18:28:50
    Running on machine: tombenke-Latitude-E5470
    Binary: Built with gc go1.11.2 for linux/amd64
    Log line format: [IWEF]mmdd hh:mm:ss.uuuuuu threadid file:line] msg

```

This trick we can apply to the other commands as well.

Often, when we need to test our query expressions, we do not need a persistent database. It can be much easier and shorter to load the test data into an in-memory store by the `--load <data-file>` CLI parameter, and run the query on it. The following command executes the same query on the same data like in the previous example, but it uses a temporary, in-memory database.

```bash

    $ echo "g.V().All();" | cayley query --logtostderr false --load data/testdata.nq --logs ~/tmp

    {"id":"alice"}
    {"id":"follows"}
    {"id":"bob"}
    {"id":"fred"}
    {"id":"status"}
    {"id":"cool_person"}
    {"id":"dani"}
    {"id":"charlie"}
    {"id":"greg"}
    {"id":"emily"}
    {"id":"predicates"}
    {"id":"are"}
    {"id":"smart_person"}
    {"id":"smart_graph"}

```

