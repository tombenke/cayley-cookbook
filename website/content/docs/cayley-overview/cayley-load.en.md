---
title: "The `cayley load` command"
weight: 3
bookCollapseSection: true
---

# The `cayley load` command

The `cayley load` command loads resource descriptions into an existing database.

{{< figure src="/cayley-cookbook/cayley-load.png" title="Figure 1.: cayley load" >}}

The following command loads the content of the `data/testdata.nq` file into a previously created database:

```bash

    $ cayley load --db bolt --dbpath /home/tombenke/tmp/cayley/ --load data/testdata.nq 

    I0915 18:01:07.140303   25803 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0915 18:01:07.140500   25803 database.go:187] using backend "bolt" (/home/tombenke/tmp/cayley/)

```

