---
weight: 3
keywords: ["triplet","subject","predicate","object"]
bookCollapseSection: true
title: "Statements"
---

# Statements

## Subject, Predicate Object

We use statements of facts to define knowledge.
In general, we usually use declarative sentences to make statements of facts

For example:

```txt
Luke Skywalker is 23 years old.
```

or, in other words:

```txt
Luke Skywalker has age 23.
```

It is possible to convert these declarative sentences into a `subject-predicate-object` structure.

```txt
Luke-Skywalker has-age 23.
```
The parts of this sentence:
- subject: `Luke-Skywalker`,
- predicate: `has-age`,
- object: `23`.

As there may many Luke Skywalker exist in our World, therefor we will use the `Luke-Skywalker-from-StarWars` identifier to the specific person, that is the subject/object of our discussion.

We want to use this universal ID in standard format, so Luke's ID will be the following IRI: `<Luke-Skywalker-from-StarWars>`.

Luke is a resource, or entity in other words, that we can obviously identify by an IRI.
He has the _age_ property, that we can can express by the `<has-age>` IRI, since this is a generic term. In addition, we need a literal value, that is `23`, with the type of integer. This is the value that belongs to the `<has-age>` property, in case of Luke.

The final statement looks like this:

```txt
    <Luke-Skywalker-from-StarWars> <has-age> 23^^integer .
```

We can observe on the Figure 1. that the `subject-predicate-object` triple can be visualized as a simple graph, where the subject and object show up as nodes, and the predicate appears as a labeled edge betwee these two nodes.

{{< hint info >}}
__Important Note:__

Internally, the predicates are also represented in the form of graph nodes, such as IRIs or even as Literal values too. They can appear in the results of queries too. The Figure 1. shows the classical graph visualization of triples in the form of binary relations.

{{< /hint >}}

{{< figure src="/cayley-cookbook/subject-predicate-object-graph.png" title="Figure 1.: Subject-Predicate-Object Graph" >}}

The statements written in this form mentioned above are also called __triple__.

Let's see how could we define for the computer the following facts:

```txt
Luke Skywalker is a person in the StarWars Universe.
His family name is Skywalker, and his given name is Luke.
Luke is 23 years old.
Leia Organa is a person in the StarWars Universe.
Her family name is Organa, and her first name is Leia.
Leia knows Luke.
```

Let's transform the sentences listed above to the `subject-predicat-object` format, using IRIs for the predicates:

```txt
<Luke-Skywalker-from-StarWars> <is-a> <Person> .
<Luke-Skywalker-from-StarWars> <has-given-name> "Luke" .
<Luke-Skywalker-from-StarWars> <has-family-name> "Skywalker" .
<Luke-Skywalker-from-StarWars> <has-age> 23 .
<Leia-Organa-from-StarWars> <is-a> <Person> .
<Leia-Organa-from-StarWars> <has-given-name> "Leia" .
<Leia-Organa-from-StarWars> <has-family-name> "Organa" .
<Leia-Organa-from-StarWars> <knows> <Luke-Skywalker-from-StarWars> .
```

These statements will be visualized as graph, including the resources, predicates and literals we have just introduced.

The resources (subjects and objects):
- Luke Skywalker in the StarWars Universe: `<Luke-Skywalker-from-StarWars>`,
- Leia Organa in the StarWars Universe: `<Leia-Organa-from-StarWars>`,
- Person (as concept): `<Person>`.

The predicates:
- `<is-a>`,
- `<has-given-name>`,
- `<has-family-name>`,
- `<has-age>`,
- `<knows>`.

The literals:
- `"Luke"`,
- `"Skywalker"`,
- `"Leia"`,
- `"Organa"`,
- `23`.

The Figure 2. demonstrates how the statements listed above appear as graph:

{{< figure src="/cayley-cookbook/statements-example.png" title="Figure 2.: Predicates represented as graph" >}}

The entities appear in yellow color, and the literals are drawn in green color.

## Blank Nodes

Often, when we define the knowledge, we do not want to use IRIs for identification with global scope, only with scope internal to the given graph, or source file, that holds the statements. In these cases, we only need to be able to express somehow that which subject and object are connected to each other via a predicate, so it is enough to use an internal ID, that can be auto-generated, and even a hidden one. These nodes are called _Blank Nodes_.

A Blank Node also has a unique identifier, that has internal scope only for the actual graph, that is valid until the program is running.

The name of Blank nodes begin with the `_:`, that is followed by some (randomly-generated) alpha-numeric string.

The Figure 3. shows a graph, that is very similar to the one on Figure 2., but in this latter case the two persons are not identified by global IRIs, but by Blank nodes with local scope.

{{< figure src="/cayley-cookbook/blank-node-example.png" title="Figure 3.:  Predicates represented as graph, using Blank nodes" >}}

The Blank Nodes appear in gray color.

{{< seealso >}}

