---
title: "Data import / export via struct annotation"
weight: 2
bookCollapseSection: true
---

# Data import / export via struct annotation

## Structure annotations

This section demonstrates how can we convert simple or complex objects and their properties to quads.
it also shows how can we directly upload these data structures to the store, using the annotations placed into the `struct` declarations ob the objects. Moreover there are examples to the retrieval of these complex objects from the store.

Within this method it is much easier to store/retrieve complex data structures, than doing it directly with the quads that these data structures are made of, and collecting their properties one-by-one.

## Create store

In the examples we will use a simple in-memory store.
We can start using such a store via calling the `InitStore()` function, that we can find in the `store.go` file.

{{< details title="Show the source code of the `store.go`" open=false >}}
{{% code file="/static/src/kbase/store.go" language="go" %}}
{{< /details >}}

## Test data

The `bookmarks.yml` file holds the test data in YAML format. it contains a list of bookmarks.

We will upload these bookmarks, then we will query them. The results of the queries will be retrieved via the help of `struct` annotations.

{{< details title="Show the `bookmarks.yml` file" open=false >}}
{{% code file="/static/src/schema/bookmarks.yml" language="yaml" %}}
{{< /details >}}

In the `yaml.go` file, there is the `ReadFromYaml()` function, that we will use to read the data in.

The `yaml.go` contains two relevant functions:
- The `ReadFromYaml()` reads data in from YAML files,
- the `SaveToYaml()` saves data into a YAML format file.

{{< details title="Show the `yaml.go` file" open=false >}}
{{% code file="/static/src/kbase/impex/yaml.go" language="go" %}}
{{< /details >}}

## Define structures with annotations

The description of the bookmark objects, and their schema annotations are placed into the `bookmarks.go` file, as well as the functions to import and export these objects.

{{< details title="Show the `bookmarks.go` file" open=false >}}
{{% code file="/static/src/schema/bookmarks.go" language="go" %}}
{{< /details >}}

In the source code there are two structures defined:
- The `Bookmark` which is the full variant,
- ant the `BookmarkShort` which is a simplified variant.

In the `struct` declaration we placed the annotations to the `yaml` and `quad` serialization/deserialization. So these annotations can be used for both formats.

## Data import based on annotated structures

First we read the data into a `[]Bookmarks` variable, then upload it into the store.
For the uploading, we use the `schema.WriteAsQuads()` that is defined in the `"github.com/cayleygraph/cayley/schema"` package-ben, as we can see in the example in the `import_bookmarks_with_schema.go` file.

{{< details title="Show the `import_bookmarks_with_schema.go` file" open=false >}}
{{% code file="/static/src/schema/import_bookmarks_with_schema.go" language="go" %}}
{{< /details >}}

Run the program:
[![Run this code on Repl.it](https://repl.it/badge/github/tombenke/cayley-cookbook-src)](https://repl.it/@tombenke/cayley-cookbook-src#schema/import_bookmarks_with_schema.go)

```bash
    cd schema
    go run impex_bookmarks_with_schema.go yamlImpex.go store.go bookmarks.go
```

The results:

{{% code file="/static/src/schema/import_bookmarks_with_schema.out" language="txt" %}}

## Data export based on annotated structures

The export operation begins similarly to the import operation.
First we read the data in from the YAML file, then we upload it into the store.

The we retrieve all the Bookmarks from the store within the following two functions:
- The `GetAllBookmarks()` function will show the results with full details,
- the `GetAllShortBookmarks()` will show the simplified version.

Both functions are implemented in the `bookmarks.go` file, and they use the `schemaConfig.LoadTo()` method to collect the properties of the resulted objects. In order to call this method, first we need to create a `schemaConfig` object, using the `schema.NewConfig()` function.

{{< details title="Show the `impex_bookmarks_with_schema.go` file" open=false >}}
{{% code file="/static/src/schema/impex_bookmarks_with_schema.go" language="go" %}}
{{< /details >}}

Run the program:
[![Run this code on Repl.it](https://repl.it/badge/github/tombenke/cayley-cookbook-src)](https://repl.it/@tombenke/cayley-cookbook-src#schema/impex_bookmarks_with_schema.go)

```bash
    cd schema
    go run impex_bookmarks_with_schema.go yamlImpex.go store.go bookmarks.go
```

The results:

{{< details title="Show the results" open=false >}}
{{% code file="/static/src/schema/impex_bookmarks_with_schema.out" language="txt" %}}
{{< /details >}}

{{< seealso >}}
