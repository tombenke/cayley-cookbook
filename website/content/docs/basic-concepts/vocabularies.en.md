---
weight: 4
keywords: ["vocabulary", "prefix", "namespace", "foaf"]
bookCollapseSection: true
title: "Vocabularies"
---

# Vocabularies

## Vocabulary

One of the big advantages of using knowledge-graphs is that the domain knowledge of the subject area is represented and manipulated in the symbolic space, so it is relatively easy to understand for human beings. What is even more important, it is possible to obviously identify and transmit concepts among the information processing agents, whether they are humans or computers. This way the mapping of tangible physical things, and abstract concepts are possible without misunderstanding.

When we give a name an entity, or we use statements by predicates, the machine and the programmer must have to mean the same meaning behind that name. The human languages are rather redundant, and they are full with homonyms and synonyms. So we need to use well defined vocabularies that clearly define these one-to-one mappings among these entities and their names. In most of the cases these vocabularies contain IRIs to name the entities and predicates, accompanied with some details, that mention their meaning. The vocabularies may also contain application rules about the predicates and resources.

The IRIs of the vocabularies mostly are English words or acronyms. The same word -whether we use it to identify an entity or a predicate- may mean different thing in different contexts. For example the word `robot` can mean an industrial assembly robot or welding robot, in case our context is a car-assembly factory, but it can also mean an intelligent humanoid droid in case of a sci-fi novel. So the predicate IRIs such as `<knows>`, `<Person>` that we've been using so far were not precise enough.

In order to make the IRIs more specific, and unambiguous, it is typical to create dedicated vocabularies to the specific business or scientific domains. These vocabularies may contain generic terms, e. g. `person`, `friend`, `knows`, `parent`, `typeOf`, but they can be very particular, which belong to a narrow field, such as antibiotics research. It is also possible that the same word exists in several vocabularies, with similar or even with different meanings.

The [FOAF](http://xmlns.com/foaf/spec/) is a very good example to a generic vocabulary, that holds IRIs to describe personal and/or organizational properties and profiles.

When we use the IRIs, we usually indicate which vocabulary the IRI belongs to so that ensure unambiguity of its meaning. We can do this via an identifier string, that we put to the front of the IRI. This string uniquely and globally identifies the vocabulary, the given IRI belongs to.

Let's extend the previously defined predicates with an identifier, which indicates that these predicates belong to the FOAF vocabulary:
```txt
<http://xmlns.com/foaf/0.1/#age>
<http://xmlns.com/foaf/0.1/#familyName>
<http://xmlns.com/foaf/0.1/#givenName>
<http://xmlns.com/foaf/0.1/#Person>
```

## Namespaces

In the previous section we introduced the vocabulary identifier string (e.g. `http://xmlns.com/foaf/0.1/#`), which makes unambiguous the meaning of the IRIs that we use as subjects, objects or predicates. This identifier string defines a so called __namespace__, in which the words are interpreted. These IRIs are unambiguous, but not too readable for humans. In order to make it more convenient to read and write the predicates, we apply a so called __prefix__, which is a short version of the ID, and that we append to the front of the IRI, separated by an `:` from the word.

In case of the FOAF vocabulary we use the `foaf:` prefix:

- namespace: `http://xmlns.com/foaf/0.1/`
- prefix: `foaf`

Now, let's see our predicates in this prefixed form:

```txt
<foaf:age>
<foaf:familyName>
<foaf:givenName>
<foaf:Person>
```

The full IRIs and the prefixed IRIs are equivalent. The only difference between the two formats is that the full IRIs can be used everywhere, but the prefixed version may used only we have introduced the prefix and namespace, and tied them up with each other. It means, we need to define which prefix identifies which namespace. Each standard knowledge representational formats (RDF, Turtle, etc.) makes this possible via special syntax elements.

These are the most relevant vocabularies, that are registered by default in {{< cayley >}}:

| prefix | namespace |
|-|-|
| `owl` | `http://www.w3.org/2002/07/owl#` |
| `rdf` | `http://www.w3.org/1999/02/22-rdf-syntax-ns#` |
| `rdfs` | `http://www.w3.org/2000/01/rdf-schema#` |
| `schema` | `http://schema.org/` |
| `xsd` | `http://www.w3.org/2001/XMLSchema#` |

We can create vocabularies too. For example we can introduce the following namespace/prefix pair for the Star-Wars Universe:

- namespace: `http://starwars.universe/#`
- prefix: `starwars`

then we can use it, in the followings:

```txt
<starwars:Luke-Skywalker>
<starwars:Leia-Organa>
```

Within a specific knowledge representation file, we can combine the IRIs defined by several vocabularies. Let's see how do our statements look like using the prefixed IRIs:

```txt
<starwars:Luke-Skywalker> <rdf:type> <Person> .
<starwars:Luke-Skywalker> <foaf:givenName> "Luke" .
<starwars:Luke-Skywalker> <foaf:familyName> "Skywalker" .
<starwars:Luke-Skywalker> <foaf:age> 23 .
<starwars:Leia-Organa> <rdf:type> <Person> .
<starwars:Leia-Organa> <foaf:givenName> "Leia" .
<starwars:Leia-Organa> <foaf:familyName> "Organa" .
<starwars:Leia-Organa> <knows> <starwars:Luke-Skywalker> .
```

## Ontologies

For the sake of completeness, it is important to note that many of the vocabularies offers much more information, than simply just introduce the list of words with prefixes.

The semantic web technology enables us to define resources and predicates that hold additional, meta-information on other resources, and predicates. For example, we can describe the properties of predicates, whether they are reflexive, or transitive, etc.). This meta-information can be used by the inference engines to deduce implicit information from the graphs, that are not written explicitly, and do reasoning, by applying the rules of logic. These extended knowledge representations, that hold such meta-informations are called Ontologies.

For example, if we would know that the `<foaf:knows>` relation is reflexive, then it were enough to say that

```txt
<starwars:Luke_Skywalker> <foaf:knows> <starwars:Leia_Organa> .
```

moreover, we could apply the rule of logic to the statement, and infer another, implicit statement that is also true:

```txt
<starwars:Leia_Organa> <foaf:knows> <starwars:Luke_Skywalker> .
```

{{< seealso >}}

