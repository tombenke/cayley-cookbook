---
title: "Cayley Overview"
weight: 3
bookCollapseSection: true
---

This section gives an overview about the elements of the Cayley ecosystem.
It shortly summarizes how can we use the individual commands and elements.

The detailed guidelines and help to these components can be found on the [official documentation pages](https://cayley.gitbook.io/cayley/).

## Installation

This cookbook has been written on a Linux machine. The simplest way to install {{< cayley >}} on Linux is using snap:

```bash

    snap install --edge --devmode cayley

```

## The Usage of Cayley

The {{< cayley >}} can be used mainly in two ways:
- as a stand-alone server,
- as an embedded database, built-into another application, written in Go.

### Using cayley as an embedded database

The [Working with Knowledge Graph]({{< relref "/docs/working-with-cayley" >}}) chapter brings several examples about how to use {{< cayley >}} as an embedded database.

### Using cayley as stand-alone server

The `cayley` binary application, provides several commands.

You can list all the available commands, by requesting the help:

```bash

    $ cayley -h
    Cayley is a graph store and graph query layer.

    Usage:
      cayley [command]

    Available Commands:
      convert     Convert quad files between supported formats.
      dedup       Deduplicate bnode values
      dump        Bulk-dump the database into a quad file.
      help        Help about any command
      http        Serve an HTTP endpoint on the given host and port.
      init        Create an empty database.
      load        Bulk-load a quad file into the database.
      query       Run a query in a specified database and print results.
      repl        Drop into a REPL of the given query language.
      upgrade     Upgrade Cayley database to current supported format.
      version     Prints the version of Cayley.
      
```

The following commands can be used to start the triple-store and the built-in query engine in different ways:
- as a HTTP server: `cayley http`.
- as a REPL server: `cayley repl`.
- as a Query server: `cayley query`.

This binary is also good for data management, such as:
- Creating a persistent database: `cayley init`,
- Load data into a (persistent) database from a file: `cayley load`,
- Save data from a (persistent) database into a file: `cayley dump`.
- convert resource descriptor files among different representational formats: `cayley conv`.

In the following subsections we will briefly summarize, how to use these commands.

### In-memory vs. persistent stores

The {{< cayley >}} is a triple-store, which can load n-quads and/or n-triples into its database, and allows to run queries on the loaded data. There are several query languages supported. This cookbook is focusing only on the so called Gizmo language.

The database can be either in-memory or persistent, and we mostly use the _store_ terminology instead of the _database_ term. Actually {{< cayley >}} offers several backing databases to store the quads persistently. The default persistent database is the [Bolt key-value database](https://github.com/boltdb/bolt), but we can change this via configuration. We can choose from several types of key-value, NoSQL and SQL databases. Read the [documentation pages on configuration settings](https://cayley.gitbook.io/cayley/configuration) for more details. 

The in-memory store can be immediately used by both the stand-alone server application and by the embedded database. The in-memory store is the default one. The persistent store is different. It needs to be initialized before we can use it.

When we use the stand-alone {{< cayley >}} binary, we can initialize a new database by either using the `--init` switch or explicitly initialize the database with the `cayley init` command, before we start using this with the other commands.

If we want to use a persistent store we have to tell the `cayley` command which type of database we aim to use, that we can do either via the command line arguments or via a configuration file.

The config file can be written both in JSON or YAML format. This is an example config file, that defines a persistent key-value store, using Bolt, and the database file will be placed under the `/home/tombenke/cayley/` folder:

{{% code file="/static/src/config/bolt.json" language="json" %}}

{{< hint info >}}
__Note:__

When we run a `cayley` command, it locks the database, so it is not possible for example to load data into a persistent database that is currently used by another command, such as `http`, etc. So there are two ways to manipulate data in a persistent store:
- Use the `cayley` commands individually, one-by-one, for example: `init`, `load`, `http`, `dump`,
- Use the REST API of the `cayley http` mode, which makes possible to alter the content in parallel using the query functionalities at the same time.

{{< /hint >}}

