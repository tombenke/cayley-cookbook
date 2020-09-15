---
title: "The `cayley init` command"
weight: 2
bookCollapseSection: true
---

# The `cayley init` command

We can create an empty database with the `cayley init` command.

{{< figure src="/cayley-cookbook/cayley-init.png" title="Figure 1.: cayley init" >}}

The following command will create a persistent database, using the Bolt key-value store type, in the `/home/tombenke/tmp/cayleydb` folder:

```bash

    $ cayley init --db bolt --dbpath /home/tombenke/tmp/cayleydb 
    
    I0915 17:52:56.529687   24933 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0915 17:52:56.529904   24933 database.go:187] using backend "bolt" (/home/tombenke/tmp/cayleydb)

```

The resulted database can be used by the other commands, such as: `load`, `dump`, `repl`, `query` and `http`.

