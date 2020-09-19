---
weight: 2
bookCollapseSection: true
keywords: ["entity","resource","literal","property","IRI"]
title: "Resources"
---

# Resources

## Resource

The resource is a generic term, that we use both on the normal web and in the domain of semantic web.
We usually mean a tangible or an abstract thing. 

A resource has two important properties:
- it has unique identity,
- it can be represented (described in a way that holds each relevant properties of the resource).

Taken into consideration these two properties, we can say that the resource is compatible with the entity term of the Domain Driven Design methodology.
From the perspective of our study, we can take these two terms to be equal with each other.
It is very important to note that, we have to be able to uniquely identify the resource of discussion both globally or in a local scope, so it needs to have at least one unique ID, and we also have to describe it somehow.

There are several possible ways of describing a thing. it can be textual, image, binary program code, in-memory state, etc. In the followings we are going to focus standard representational formats, especially to the ones that are created for symbolic representations, such as the standards worked out for the semantic web.

Examples of resources:
- specific, or imagined persons,
- tangible, real-world objects,
- abstract concepts.

The following code fragment represents an imagined person in JSON format:

```json
{
    "id": "<star-wars-universe:luke_skywalker>",
    "firstName": "Luke",
    "lastName": "Skywalker",
    "age": 23
}
```

The same person in YAML representation:

```yaml
---
id: <star-wars-universe:luke_skywalker>
firstName: Luke
lastName: Skywalker
age: 23
```

The same person is represented in two different formats.

Typically the resources are the nodes of the knowledge-graphs, but often, the edges among the nodes are also resources, as they represent abstract concepts. In many cases, we can use these edges, as connections between resources, then they also appear in the form of nodes, when we want to make assertions about their properties, or want to define rules about their validity and usage.

## IRIs

The resources (or entities) have to be identified.
For this purpose we use the so called IRIs.
IRI is the acronym of Internationalized Resource Identifier.

An IRI is a Unicode character string. Its standard format is defined by [RFC 3987](https://www.ietf.org/rfc/rfc3987.txt), and it uniquely identifies an entity.

The IRI strings can contain alpha-numeric characters, escape sequences, and these are enclosed between the `<` and `>` characters. An IRI is very similar to the URIs, that stands for Uniform Resource Identifier, which are well known from the World-Wide-Web.

The IRIs can be taken to the generalized version of URIs. The exist only in absolute form. every absolute URI can be used as an IRI, but not every IRI can be used as URL.

By definition, the IRIs are defined within global scope. Accordingly, Appearances of a given IRI in different situations, and environments mean the very same entity.

Each IRI identifies a well defined thing as resource in the matter of discussion. The resource, that the IRI refers to is called _referent_.

Examples for IRIs:

```txt
<star-wars-universe:luke_skywalker>
<star-wars-universe:han_solo>
<35345322-543555-534534532>
<https://en.wikipedia.org/wiki/Millennium_Falcon>
```

The IRIs can identify abstract things and concepts as well, for example the attributes of entities as well as the relations among entities. 
Here are some examples for this:

```txt
<has>
<is>
<nameOf>
<foaf:knows>
```

## Literal values

The resources we want to represent, usually have internal details, that we also want to describe. These internal properties have specific values.

For example, in case of a specific person, such properties can be the name of the person, the date of birth, etc. The values of these properties then the actual name and date, that the properties can hold.

These property values are named _literal values_, or simply _literals_.

We can define the type of a literal value, e.g.: string, integer number, date, etc.

We usually add the type information to the literal in postfix form, using the `^^` separator.

Examples for typed values:
```txt
"Dantooine"^^xsd:string
"23"^^<xsd:integer>
```

In case of a string, beside the type, we also can tell what language is used for definition. It is defined via a language tag, as a postfix value, using the `@` character as separator.

Examples for language identification:
```txt
"people"@en
"emberek"@hu
```

Among others, it is also possible to define the direction of writing in addition to the language definition, however these fine details are not necessary to know in order to start using the {{< cayley >}} system.
