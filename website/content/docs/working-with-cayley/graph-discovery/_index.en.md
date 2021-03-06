---
title: "Discover and unknown graph"
weight: 1
bookCollapseSection: true
---

## Problem:

How can I execute queries in an unknown graph (e.g. the sample film database), if I do not know its internal structure?

## Solution

In the following we will go through a possible sequence of steps of analysis. At the end of the process, we will be able to query or modify the unknown graph. We do not need to know anything about this graph. At the same time, we suppose that this unknown graph is not made of randomly created and connected nodes, but it has it own "system". At least at the theoretical level, there is a virtual schema of it, and the predicates used in the graph follow some naming conventions, and they were created among some logic, that may refer to their meaning and roles.

We will visualize the schema on a diagram, step-by-step as we move forward on the process of analysis. We will draw with yellow color those IRIs that we know their actual IDs. The literal nodes will be drawn with green color including the name of the type. The unknown nodes will be drawn with blue color, and we will assign a variable name to them.

### Step 0.: Start the REPL with the database

Let's start the REPL console of {{< cayley >}}-t using an in-memory database, immediately loading the data to be investigated:

```bash

    $ cayley repl --load data/30kmoviedata.nq.gz 
    I0917 19:12:42.483992   18114 cayley.go:63] Cayley version: 0.7.5 (cf576babb7db)
    I0917 19:12:42.484240   18114 database.go:187] using backend "memstore"
    I0917 19:12:46.897340   18114 database.go:250] loaded "data/30kmoviedata.nq.gz" in 4.413006871s
    cayley>

```

In the followings we will use this utility to execute the steps of the analysis.

### Step 1.: Find all the predicates used in the graph

The following query finds all outgoing predicates, starting from the set of all nodes in the graph. Then it also collects all incoming predicates as well, finally creates the union of the two result sets. In order to have each predicate, only once, it also executes the `Unique()` operation on the union set:

```bash

    cayley> g.V().OutPredicates().Union(g.V().InPredicates()).Unique().All();

    ****
    id : </film/performance/actor>
    ****
    id : <name>
    ****
    id : <type>
    ****
    id : </film/performance/character>
    ****
    id : </film/film/directed_by>
    ****
    id : </film/film/starring>
    -----------
    6 Results
    Elapsed time: 720.486208 ms

```

In the results we can see that these are the predicates that used in the graph:

- `</film/performance/actor>`
- `<name>`
- `<type>`
- `</film/performance/character>`
- `</film/film/directed_by>`
- `</film/film/starring>`

Try to visualize an initial schema that holds the collected info, that we will gradually extend as we discover the graph step-by-step. The Figure 1. shows the diagram.


{{< figure src="/cayley-cookbook/discover_movie_graph_1.png" title="Figure 1.: The predicates of the database" >}}

### Step 2.: Form predicate groups according to their co-occurrence

Let's identify those predicates based on their co-occurrence on nodes, and form groups out of these predicates.


#### Step 2.1.: Aggregate the outgoing predicates

First collect those outgoing predicates, that have their inbound leg on the same node.

```bash

    cayley> g.V().In("</film/performance/actor>").OutPredicates().All()

    ****
    id : </film/performance/actor>
    ****
    id : </film/performance/character>
    -----------
    2 Results
    Elapsed time: 174.893162 ms

```
The result list shows that the `</film/performance/actor>` and the `</film/performance/character>` predicates belong to the same group, so it is very probable that usually the connect to the same node types as outbound predicates.

We have to gain the same result, whichever predicate from this group we use for the same type of query:

```bash

    cayley> g.V().In("</film/performance/character>").OutPredicates().All()

    ****
    id : </film/performance/actor>
    ****
    id : </film/performance/character>
    -----------
    2 Results
    Elapsed time: 45.116628 ms

```

Continue with the next predicate, that does not belong to this group, that we have just identified:

```bash

    cayley> g.V().In("<name>").OutPredicates().All()

    ****
    id : <name>
    ****
    id : <type>
    ****
    id : </film/film/directed_by>
    ****
    id : </film/film/starring>
    -----------
    4 Results
    Elapsed time: 178.340908 ms

```

We could identify another group that contains all the remaining predicates. This way we have grouped all outgoing the predicates contained by the graph.

Now assign these predicates to node variables, according to their grouping, and put them to the graph diagram, as we can see on Figure 2.

{{< figure src="/cayley-cookbook/discover_movie_graph_2_1.png" title="figure 2.: Outgoing predicates grouped by their co-occurrence" >}}

#### Step 2.2: Aggregate the incoming predicates

Let's crawl on the predicates that arrive as incoming edge to the previously discovered nodes, and assign them to these node variables.

In order to identify these incoming predicates, let's iterate through each items of the grouped outgoing predicates, and check their co-occurrences. We can do this by executing the following series of queries:



```bash

    cayley> g.V().In("</film/performance/actor>").InPredicates().All()

    ****
    id : </film/film/starring>
    -----------
    1 Result
    Elapsed time: 190.712052 ms

```

```bash

    cayley> g.V().In("</film/performance/character>").InPredicates().All()

    ****
    id : </film/film/starring>
    -----------
    1 Result
    Elapsed time: 42.079836 ms

```
 
```bash

    cayley> g.V().In("<type>").InPredicates().All()

    ****
    id : </film/performance/actor>
    ****
    id : </film/film/directed_by>
    -----------
    2 Results
    Elapsed time: 149.433856 ms

```

```bash

    cayley> g.V().In("<film/film/directed_by>").InPredicates().All()

```

In this latter case, there is no co-occurrence found.


```bash

    cayley> g.V().In("<film/film/starring>").InPredicates().All()

```

and neither in this case.

Put the results of these queries onto the schema diagram as inbound edges, connecting to the corresponding node variables, as we can see on Figure 3.

{{< figure src="/cayley-cookbook/discover_movie_graph_2_2.png" title="Figure 3.: Incoming and outgoing predicates grouped by their co-occurrence" >}}

#### Merge the incoming and outgoing predicate groups

Merge the results sets as well as the schema diagram, according to the Figure 4.

{{< figure src="/cayley-cookbook/discover_movie_graph_2_3.png" title="Figure 4.: The co-occurring predicates." >}}

{{< hint info >}}
__Important note:__

The `x` and `y` node variables are not necessarily represent only one individual node type, because we put every inbound and outbound predicates into single groups.

{{< /hint >}}

### Step 3.: Find the terminal nodes

The terminal nodes are the ones that has no more outgoing predicate to other nodes.

Let's study for each predicate that what kind of node it goes to.

Start with the `</film/performance/character>` predicate:

```bash

    cayley> g.V().Out("</film/performance/character>").All()

    ****
    id : Luther Krank
    ****
    id : Roland
    ****
    id : Tomás de Torquemada
    ****
    id : Ferdinand VII of Spain
    ****
    id : Christopher Columbus
    ...

```

Only the first 100 items of the result list will be printed by default.
We can see that the results are string literals, the character names of the films.

Count their number:

```bash

    cayley> g.V().Out("</film/performance/character>").Count()

    => 15058
    -----------
    1 Result
    Elapsed time: 15.394561 ms

```

Now let's check, if there is any predicate that goes further from any of these nodes:

```bash

    cayley> g.V().Out("</film/performance/character>").Out().Count()

     => 0
    -----------
    1 Result
    Elapsed time: 65.014376 ms

```

The number of such nodes are 0, so these are terminal nodes, since they have no outgoing connections.

Now do the same operation with the `<name>` predicate:

```bash

    cayley> g.V().Out("<name>").All()

    ****
    id : David Fisher
    ****
    id : 002 Operazione Luna
    ****
    id : 008: Operation Exterminate
    ****
    id : 02:37:00 AM
    ****
    id : 06/05
    ****
    id : 10,000 BC
    ****
    ...

```

These are also string literals.

```bash

    cayley> g.V().Out("<name>").Count()

    => 74950
    -----------
    1 Result
    Elapsed time: 44.480211 ms

```

There are 74950.

```bash

    cayley> g.V().Out("<name>").Out().Count()

    => 0
    -----------
    1 Result
    Elapsed time: 52.163808 ms

```

And these are also terminal nodes.

Let's continue with the `<type>` predicate:

```bash

    cayley> g.V().Out("<type>").All()

    ****
    id : </people/person>
    ****
    id : </film/film>
    ****
    id : </film/film>
    ...

```

These are IRIs, and we have many different ones.

Let's collect these IRIs individually:


```bash

    cayley> g.V().Out("<type>").Unique().All()
     
    ****
    id : </people/person>
    ****
    id : </film/film>
    -----------
    2 Results
    Elapsed time: 43.759725 ms

```

In total, there are only two such IRIs, the `</people/person>` and the `</film/film>`.

Let's check if there were terminal nodes:

```bash

    cayley> g.V().Out("<type>").Out().Count()

    => 0
    -----------
    1 Result
    Elapsed time: 48.857924 ms

```

Yes, they are terminal nodes.

Put these newly discovered IRI and literal nodes to the schema diagram, accompanied with the predicates that lead to them, as we can see on the Figure 5.

{{< figure src="/cayley-cookbook/discover_movie_graph_3_1.png" title="5. ábra: The schema diagram extended with the terminal node types." >}}

### 4. Separate the different node types from each other

As much as possible, try to differentiate among the node types, and separate the different types from each other in case they may receive or send the same predicates, but refer to different kind of entities.

#### Analysis of the `y` node variable

We can observe on the diagram that the `y` node has the `<type>` outgoing predicate with two different values, so this `y` node might be split to two different node types, that identify different entities. Let's separate the `y` node variable to two types, according to our hypotheses, as we can see on the Figure 6.

{{< figure src="/cayley-cookbook/discover_movie_graph_4_1.png" title="6. ábra: Separation of `person` and `film` node-types." >}}

Use multi-step (2, 3, ...n) predicate chains for the further analysis.

Let's see which predicates points out from, and arrives to the node types derived from `y`, when we restrict the starting nodes to have their `<type>` relation with either the one type then the other type:

```bash

    cayley> y_film = g.V().Has("<type>", "</film/film>")

    => [internal Iterator]
    -----------
    1 Result
    Elapsed time: 0.469018 ms

    cayley> y_film.OutPredicates().All()

    ****
    id : </film/film/directed_by>
    ****
    id : </film/film/starring>
    ****
    id : <name>
    ****
    id : <type>
    -----------
    4 Results
    Elapsed time: 113.54646 ms

```

Looking at the results, in case of films these seems to be normal.

Now do the same check with the case of `person` node-type candidates:


```bash

    cayley> y_person = g.V().Has("<type>", "</people/person>")

    => [internal Iterator]
    -----------
    1 Result
    Elapsed time: 0.552285 ms

    cayley> y_person.OutPredicates().All()

    ****
    id : <name>
    ****
    id : <type>
    ****
    id : </film/film/directed_by>
    ****
    id : </film/film/starring>
    -----------
    4 Results
    Elapsed time: 95.020065 ms

```

We got the same results, that is a bit surprising, because it is hard to imagine that it makes sense to use the `</film/film/directed_by>` nor the `</film/film/starring>` predicates as outgoing connections.

We caught suspicion, so take a look at this a bit deeper:

A film probably has a director, so it worth to check this kind of connection:

```bash

    cayley> y_film.Out("</film/film/directed_by>").Count()

    => 33310
    -----------
    1 Result
    Elapsed time: 108.365305 ms

```

We found 33310 hits of such kind of relation, and it seems to be obvious.

Let's check, if the target nodes this predicate points to, are really persons!

```bash

    cayley> y_film.Out("</film/film/directed_by>").Has("<type>", "</people/person>").Count()

    => 33310
    -----------
    1 Result
    Elapsed time: 168.508677 ms

```

Taking into consideration the numbers, it looks true. Let's see some of them by name:

```bash

    cayley> y_film.Out("</film/film/directed_by>").Has("<type>", "</people/person>").Out("<name>").GetLimit(5)

    ****
    id : Lucio Fulci
    ****
    id : Umberto Lenzi
    ****
    id : Murali K. Thalluri
    ****
    id : Theo van Gogh
    ****
    id : Roland Emmerich
    -----------
    5 Results
    Elapsed time: 1.071004 ms

```

it seems that this type of connection, and node type are absolutely O.K.

Now let's see what about the same relation with the person node types.
Has any node with `</people/person>` type any "directed_by" relation?

```bash

    cayley> y_person.Out("</film/film/directed_by>").Count()

    => 6
    -----------
    1 Result
    Elapsed time: 96.554532 ms

```

Surprisingly we found 6 of them.

Let's see these 6 persons closer. We use tags to visualize the details, that makes easier to understand what is the situation.

```bash

    cayley> y_person.Tag("person").Out("</film/film/directed_by>").Tag("directed_by").All()

    ****
    directed_by : </en/quentin_tarantino>
    id : </en/quentin_tarantino>
    person : </en/death_proof>
    ****
    directed_by : </en/mysterio>
    id : </en/mysterio>
    person : </en/jazmin>
    ****
    directed_by : </en/robert_rodriguez>
    id : </en/robert_rodriguez>
    person : </en/planet_terror>
    ****
    directed_by : </en/keenen_ivory_wayans>
    id : </en/keenen_ivory_wayans>
    person : </en/scary_movie_2>
    ****
    directed_by : </en/david_zucker>
    id : </en/david_zucker>
    person : </en/scary_movie_3>
    ****
    directed_by : </en/ralph_bakshi>
    id : </en/ralph_bakshi>
    person : </en/the_lord_of_the_rings_1978>
    -----------
    6 Results
    Elapsed time: 95.293387 ms

```

As we can see on the results list, in case of these 6 nodes, the `<directed_by>` predicate is really point to persons, however we also can realize that the IRI of the origin node shows that in fact these origins are not persons, but films. We can say, it is very likely that the `<type>` predicate were defined badly in these six cases. These `<type>` predicates point to the `</people/person>` IRI, instead of the correct `</film/film>` IRI. So we consider this connection invalid, in spite the database contains such wrong links.

The figure 7. shows the schema drawing that contains our latest findings too, that are: the `y_film` node-types have outgoing `</film/film/directed_by>` predicate that points to `y_person` node types, an the `y_person` node types has no such predicates.

{{< figure src="/cayley-cookbook/discover_movie_graph_4_2.png" title="Figure 7.: The outgoing predicates of `y_film` node types." >}}

Now let's continue with the analysis of the `</film/film/starring>` predicate.

```bash

    cayley> y_film.Out("</film/film/starring>").Count()

    => 136737
    -----------
    1 Result
    Elapsed time: 140.230179 ms

    cayley> y_film.Out("</film/film/starring>").GetLimit(5)

    ****
    id : _:117646
    ****
    id : _:117647
    ****
    id : _:117648
    ****
    id : _:117649
    ****
    id : _:117240
    -----------
    5 Results
    Elapsed time: 1.052197 ms

```

The `y_film` node type has 136737 connections with this predicate, that point to Blank Nodes.

Let's see what kind of incoming and outgoing predicates these "character" nodes have:

```bash

    cayley> y_film.Out("</film/film/starring>").OutPredicates().All()

    ****
    id : </film/performance/actor>
    ****
    id : </film/performance/character>
    -----------
    2 Results
    Elapsed time: 255.873653 ms

    cayley> y_film.Out("</film/film/starring>").InPredicates().All()

    ****
    id : </film/film/starring>
    -----------
    1 Result
    Elapsed time: 242.836008 ms

```

The "character" Blank Nodes can be considered as the node type that we named to `x` on the schema diagram.

As a check, let's see if we can find actors and roles connected to films, using the predicates:

```bash

    cayley> y_film = g.V().Has("<type>", "</film/film>")

    => [internal Iterator]
    -----------
    1 Result
    Elapsed time: 0.897193 ms

    cayley> filmWithTitle = y_film.Tag("film").Out("<name>").Tag("filmTitle").Back("film")

    => [internal Iterator]
    -----------
    1 Result
    Elapsed time: 0.680148 ms

    cayley> filmStar = filmWithTitle.Out("</film/film/starring>").Tag("star")

    => [internal Iterator]
    -----------
    1 Result
    Elapsed time: 0.340275 ms

    cayley> {t("</film/performance/character>").Tag("character").Back("star").Out("</film/performance/actor>").Out("<name>").Tag("actorName").GetLimit("10")

    ****
    actorName : Katherine Heigl
    character : Arlene
    film : </en/100_girls>
    filmTitle : 100 Girls
    id : Katherine Heigl
    star : _:239
    ****
    actorName : Joely Richardson
    character : Anita
    film : </en/101_dalmatians>
    filmTitle : 101 Dalmatians
    id : Joely Richardson
    star : _:457
    ****
    actorName : Susanne Blakeslee
    character : Cruella de Vil
    film : </en/101_dalmatians_ii_patchs_london_adventure>
    filmTitle : 101 Dalmatians II: Patch's London Adventure
    id : Susanne Blakeslee
    star : _:71373
    ****
    actorName : Brian Markinson
    character : Daniel
    film : </en/10_5_apocalypse>
    filmTitle : 10.5: Apocalypse
    id : Brian Markinson
    star : _:218
    ****
    actorName : Heath Ledger
    character : Patrick Verona
    film : </en/10_things_i_hate_about_you>
    filmTitle : 10 Things I Hate about You
    id : Heath Ledger
    star : _:503
    ****
    actorName : James Marsden
    character : Tommy
    film : </en/10th_wolf>
    filmTitle : 10th & Wolf
    id : James Marsden
    star : _:284
    ****
    actorName : Jeana Tomasino
    character : Karen
    film : </en/10_to_midnight>
    filmTitle : 10 to Midnight
    id : Jeana Tomasino
    star : _:100658
    ****
    actorName : Jason Segel
    character : Leon (paramedic 1)
    film : </en/11_14>
    filmTitle : 11:14
    id : Jason Segel
    star : _:108561
    ****
    actorName : Henry Fonda
    character : Juror #8
    film : </en/12_angry_men>
    filmTitle : 12 Angry Men
    id : Henry Fonda
    star : _:462
    ****
    actorName : Lee J. Cobb
    character : Juror #3
    film : </en/12_angry_men>
    filmTitle : 12 Angry Men
    id : Lee J. Cobb
    star : _:463
    -----------
    10 Results
    Elapsed time: 7.233902 ms

```

It definitely works. So we can say that the `x` node variable represents the film stars. 

We also can say that both the `</film/performance/actor>` predicate, and the `</film/performance/character>` predicate that origins from `x` and targets `y_person` are valid predicates, and represent real connections.

Now emerges a new question, whether there is an outgoing `</film/performance/actor>` predicate from `x` that targets a `y_film` node?

```bash

    filmStar.Out("</film/performance/actor>").Has("<type>", "</film/film>").Count()

    => 1
    -----------
    1 Result
    Elapsed time: 1315.917655 ms

```

It is strange, but we have found one.

Which is this starring character, and which actor does it point to?

```bash

    cayley> filmStar.Out("</film/performance/actor>").Has("<type>", "</film/film>").All()

    ****
    film : </en/jazmin>
    filmTitle : Jazmin
    id : </en/jazmin>
    star : _:23742
    -----------
    1 Result
    Elapsed time: 1340.16577 ms

```

We can see that the ID of this starring character is the `_:23742` Blank Node, which points to a film instead of a person as the actor. Let's check if this Blank Node also has a similar connection to a real person too:

```bash

    cayley> g.V("_:23742").Out("</film/performance/actor>").Has("<type>", "</people/person").Count()

    => 0
    -----------
    1 Result
    Elapsed time: 0.68618 ms

```

It has no such connection, so we have identified another bug in the database. We also make changes on the schema chart according to the new learnings, as it is shown on the Figure 8.

{{< figure src="/cayley-cookbook/discover_movie_graph_4_3.png" title="figure 8.: The connections between the `y_person` and `x_character` node types." >}}

Finally there is only one predicate left, that we haven't analyzed so far. It the the `</film/film/starring>` that origins from the `y_person` node types.

```bash

    cayley> y_person = g.V().Has("<type>", "</people/person>").Out("</film/film/starring>").Count()

    => 96
    -----------
    1 Result
    Elapsed time: 85.149164 ms

```

It is again very surprising, but we have found 96 connections. Let's find out how is this possible, and where do these connections point to.

```bash

    cayley> y_person = g.V().Has("<type>", "</people/person>").Out("</film/film/starring>").GetLimit(5)

    ****
    id : _:95530
    ****
    id : _:95531
    ****
    id : _:95532
    ****
    id : _:95533
    ****
    id : _:95534
    -----------
    5 Results
    Elapsed time: 12.859301 ms

```

it is clearly visible that the Blank Nodes very likely points to `x_character` type nodes.

Execute a query, where we start from person type nodes ("originPeson") through a `</film/film/starring>` predicate, towards an actor ("actor"), and retrieve the name and type of the node that is at the end of the chain.

```bash

    cayley> g.V().Tag("originPerson").Has("<type>", "</people/person>").Out("<type>").Tag("originType").Back("originPerson").Out("</film/film/starring>").Out("</film/performance/actor>").Tag("actor").Out("<type>").Tag("actorType").Back("person").All()

    ****
    actor : </en/jazmin>
    actorType : </film/film>
    id : </en/jazmin>
    originPerson : </en/jazmin>
    originType : </film/film>
    ****
    actor : </en/jazmin>
    actorType : </people/person>
    id : </en/jazmin>
    originPerson : </en/jazmin>
    originType : </film/film>
    ****
    actor : </en/jazmin>
    actorType : </people/person>
    id : </en/jazmin>
    originPerson : </en/jazmin>
    originType : </people/person>
    ****

    ...


    ****
    actor : </en/anthony_daniels>
    actorType : </people/person>
    id : </en/the_lord_of_the_rings_1978>
    originPerson : </en/the_lord_of_the_rings_1978>
    originType : </people/person>
    ****
    actor : </guid/9202a8c04000641f800000000112e07f>
    actorType : </people/person>
    id : </en/the_lord_of_the_rings_1978>
    originPerson : </en/the_lord_of_the_rings_1978>
    originType : </film/film>
    -----------
    100 Results
    Elapsed time: 965.035817 ms
```

If we look through the result list, again we can observe contradictions among the relations in the database. We can detect several errors. Based on the results list, we can assume that some of the `y` nodes may have `<type>` predicates that points both to the `</film/film>` and the `</people/person>` IRI nodes.

Prove that this assumption is true:

```bash

    cayley> g.V().Has("<type>", "</people/person>").And(g.V().Has("<type>", "</film/film>")).All()

    ****
    id : </en/death_proof>
    ****
    id : </en/jazmin>
    ****
    id : </en/planet_terror>
    ****
    id : </en/scary_movie_2>
    ****
    id : </en/scary_movie_3>
    ****
    id : </en/the_lord_of_the_rings_1978>
    -----------
    6 Results
    Elapsed time: 72.999153 ms

```

Unfortunately, our assumptions were correct, and these are real inconsistencies in the database.
We can state that from the `person` node type, there is no `</film/film/starring>` outgoing predicate towards the `character` nodes in normal cases.

Now we can finish the schema diagram, that you can see on the figure 9.

{{< figure src="/cayley-cookbook/discover_movie_graph_4_4.png" title="Figure 9.: The schema diagram of the film database." >}}

