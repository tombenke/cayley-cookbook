---
title: "The `cayley http` command"
weight: 7
bookCollapseSection: true
---

# The `cayley http` command

Maybe the `cayley http` is the most well known, and most frequently used command.
It is an application, which contains the quad-store, the query engines and a HTTP server, that provides all of the functionalities through a REST API for utilities and 3rd-party applications, as well as provides a web-based graphical UI for the end users.

The Figure 1. shows all the components that can be used together with the {{< cayley >}} running in the  `cayley http` mode.

{{< figure src="/cayley-cookbook/cayley-http.png" title="Figure 1.: cayley http" >}}

The `cayley http` exposes REST API endpoints which are described in the [official documentation](https://cayley.gitbook.io/cayley/usage/http). See also the [documentation pages of the 3rd party APIs](https://cayley.gitbook.io/cayley/usage/3rd-party-apis) avaiable for accessing to the `cayley http` server.

## Starting the Cayley HTTP Server

### Start with default settings

The execution of the `cayley http` command starts the server with the default settings:

```bash

    $ cayley http

    I0916 08:00:12.107989    6702 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0916 08:00:12.108196    6702 database.go:187] using backend "memstore"
    I0916 08:00:12.108322    6702 http.go:197] using assets from "/snap/cayley/2"
    I0916 08:00:12.108805    6702 http.go:42] listening on 127.0.0.1:64210, web interface at http://127.0.0.1:64210

```

The server uses an in-memory database, and the store will be empty. The server will listen on the `64210` port by default. We can open the `http://127.0.0.1:64210` URL with a browser, then the web UI will appear, where we can run queries on the database.

### Start with in-memory store with automatic data loading

We can tell to the `cayley http` to load data into the store immediately after starting it. We can use the `--load <data-file-path>` argument for this.

```bash

    $ cayley http --load data/testdata.nq 
    I0916 08:10:40.473364    7087 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0916 08:10:40.473602    7087 database.go:187] using backend "memstore"
    I0916 08:10:40.473935    7087 database.go:250] loaded "data/testdata.nq" in 266.525µs
    I0916 08:10:40.474084    7087 http.go:197] using assets from "/snap/cayley/2"
    I0916 08:10:40.474498    7087 http.go:42] listening on 127.0.0.1:64210, web interface at http://127.0.0.1:64210

```

The data will be available immediately for querying through both the UI and through the REST API for the client applications. Yet the store is in-memory.

### Start with persistent store, including data

The `cayley http` can be started with a persistent database as well. Here we have two options:
1. We initialize an empty persistent store, and may also load data into it immediately after the database has been created.
2. Open a previously created persistent store, that may be empty or filled with some data. To prepare such a database we can use the `cayley init` and `cayley load` commands.

Start an empty persistent store, and load data into it:
```bash

    $ cayley http --init --db bolt --dbpath /home/tombenke/tmp/cayley/ --load data/testdata.nq 
    I0916 08:20:05.039047    7367 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0916 08:20:05.039234    7367 database.go:187] using backend "bolt" (/home/tombenke/tmp/cayley/)
    I0916 08:20:05.070109    7367 database.go:250] loaded "data/testdata.nq" in 9.150519ms
    I0916 08:20:05.070347    7367 http.go:197] using assets from "/snap/cayley/2"
    I0916 08:20:05.071829    7367 http.go:42] listening on 127.0.0.1:64210, web interface at http://127.0.0.1:64210
    I0916 08:20:11.371920    7367 http.go:100] started POST /api/v1/query/gizmo for 127.0.0.1:39426
    I0916 08:20:11.372573    7367 http.go:102] completed 200 OK /api/v1/query/gizmo in 666.087µs
    I0916 08:20:17.928216    7367 http.go:100] started POST /api/v1/query/gizmo for 127.0.0.1:39426
    I0916 08:20:17.932551    7367 http.go:102] completed 200 OK /api/v1/query/gizmo in 4.335239ms

```

Prepare a persistent store with `init` and `load`, start the `http` server, using it:
```bash

    $ cayley init --db bolt --dbpath /home/tombenke/tmp/cayley/
    I0916 08:21:00.701903    7460 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0916 08:21:00.702107    7460 database.go:187] using backend "bolt" (/home/tombenke/tmp/cayley/)

    $ cayley load --db bolt --dbpath /home/tombenke/tmp/cayley/ --load data/testdata.nq
    I0916 08:21:42.843042    7509 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0916 08:21:42.843269    7509 database.go:187] using backend "bolt" (/home/tombenke/tmp/cayley/)

    $ cayley http --db bolt --dbpath /home/tombenke/tmp/cayley/
    I0916 08:22:08.376538    7573 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0916 08:22:08.376726    7573 database.go:187] using backend "bolt" (/home/tombenke/tmp/cayley/)
    I0916 08:22:08.379455    7573 http.go:197] using assets from "/snap/cayley/2"
    I0916 08:22:08.380016    7573 http.go:42] listening on 127.0.0.1:64210, web interface at http://127.0.0.1:64210

```

In both cases, we will be able to reach the same database content through the HTTP server.

{{< hint info >}}
__Note:__

Both the `cayley load` and `cayley save` can be used to store data into, and dump from the database that the `cayley http` uses, however it can happen only in case the server is stopped. That is why these utilities are connected via dotted line on Figure 1.

{{< /hint >}}

## The web UI

Start the `cayley http` server either with an in-memory or persistent database, that has data loaded into it (according to examples of the previous subsections), then open the `http://127.0.0.1:64210` URL with a browser. In the text area on the right top, enter the `g.V().GetLimit(2)` query expression, then press the "Run Query" button on the left. The results will appear on the right bottom part of the window, as you can see on Figure 2.

{{< figure src="/cayley-cookbook/cayley-http-ui.png" title="Figure 2.: cayley http UI" >}}

See the [official documentation on the UI](https://cayley.gitbook.io/cayley/usage/ui-overview) for further details:


## Data import and export

If you want to upload data into the database, or want to dumt the database, meanwhile the `cayley http` server is working, you have to use the REST interface. There are two utilities, that you can find in the source code of {{< cayley >}}: [`cayleyimport`](https://github.com/cayleygraph/cayley/tree/master/cmd/cayleyimport) and [`cayleyexport`](https://github.com/cayleygraph/cayley/tree/master/cmd/cayleyexport).

Clone the [cayleygraph/cayley](https://github.com/cayleygraph/cayley), and you can run these utilities with the `go run` command, as the examples show below.

Import data to the database, through the REST API:

```bash

    $ go run <path-to-the-cayley-source>/cmd/cayleyimport/cayleyimport.go data/testdata.nq 
    Successfully wrote 15 quads.


```

Export data from the database, into a file, through the REST API:

```bash

    go run <path-to-the-cayley-source>/cmd/cayleyexport/cayleyexport.go -o out.nq

```

Export data from the database, to the standard output, through the REST API. Since there is no output file name is defined, the utility can not determine the format, so we need to set it by the `--format` argument:


```bash

    $ go run ~/sandbox/cayley/cayley/cmd/cayleyexport/cayleyexport.go --format jsonld
    [{"@id":"alice","follows":[{"@id":"bob"}]},{"@id":"bob","follows":[{"@id":"fred"}],"status":[{"@value":"cool_person"}]},{"@id":"charlie","follows":[{"@id":"bob"},{"@id":"dani"}]},{"@id":"dani","follows":[{"@id":"bob"},{"@id":"greg"}],"status":[{"@value":"cool_person"}]},{"@id":"emily","follows":[{"@id":"fred"}]},{"@id":"fred","follows":[{"@id":"greg"}]},{"@id":"greg","status":[{"@value":"cool_person"}]},{"@id":"predicates","are":[{"@id":"follows"},{"@id":"status"}]},{"@graph":[{"@id":"emily","status":[{"@value":"smart_person"}]},{"@id":"greg","status":[{"@value":"smart_person"}]}],"@id":"smart_graph"}]

```


