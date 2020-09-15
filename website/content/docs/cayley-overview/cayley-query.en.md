---
title: "The `cayley query` command"
weight: 5
bookCollapseSection: true
---

# The `cayley query` command

With the `cayley query` we can run a query expression on an existing database.

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

In case we want to feed the results into another program, then we may want to hide the logs pronted out by the command.
The next example shows how can we do this. The logs are written into the `/home/tombenke/tmp` directory:

```bash

    $ cat ~/tmp/cayley.tombenke-Latitude-E5470.tombenke.log.INFO.20200915-182850.27351 
    Log file created at: 2020/09/15 18:28:50
    Running on machine: tombenke-Latitude-E5470
    Binary: Built with gc go1.11.2 for linux/amd64
    Log line format: [IWEF]mmdd hh:mm:ss.uuuuuu threadid file:line] msg

```

This trick we can apply to the other commands as well.

