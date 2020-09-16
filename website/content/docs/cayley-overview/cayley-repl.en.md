---
title: "The `cayley repl` command"
weight: 5
bookCollapseSection: true
---

# The `cayley repl` command

The REPL is the acronym of Read-Eval-Print-Loop.
We can use the `cayley repl` command to open a console to an existing database, and interactively execute queries and other operations on the database.

The console works similar to a normal Linux console. 
We can type in the expressions, then press the Enter, and the interpreter will execute, then prints the results.
It is also possible to use the up/down arrows to navigate among the previously entered commands.
The expressions we execute will be stored into a `.cayley_history` file, that we can open with a text editor, and extract the expressions we used previously.
This way the REPL is a very efficient tool to experiment with queries, then move the results of experiments to the final query files.

{{< figure src="/cayley-cookbook/cayley-repl.png" title="Figure 3.: cayley repl" >}}

The following session starts REPL on an existing database, then queries and prints all the nodes stored in the database:

```bash

    $ cayley repl --db bolt --dbpath /home/tombenke/tmp/cayley/ 
    I0915 18:02:08.623789   25897 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0915 18:02:08.623971   25897 database.go:187] using backend "bolt" (/home/tombenke/tmp/cayley/)
    cayley> g.V().All()

    ****
    id : <bob>
    ****
    id : <status>
    ****
    id : cool_person
    ****
    id : <alice>
    ****
    id : <greg>
    ****
    id : <emily>
    ****
    id : <smart_graph>
    ****
    id : <predicates>
    ****
    id : <dani>
    ****
    id : <fred>
    ****
    id : smart_person
    ****
    id : <charlie>
    ****
    id : <are>
    ****
    id : <follows>
    -----------
    14 Results
    Elapsed time: 3.046727 ms

    cayley>

```

Let's see the content of the `.cayley_history` file after the session:
```bash

    $ cat .cayley_history 
    
    g.V().All()

```
