# Learning Spark

I've read this book before, but I wanted a quick refresher, so I'm skimming it
again. The main target for me is to understand better how to setup an
environment from scratch in different places, like a Docker container, or
concepts like Dataset or Dataframes.

For the notes on this book, as it's mostly theory and I've read it before, I'm
only going to transcribe the highlights I've done on my Kindle. If there are no
notes on a given chapter or section, I'll simply skip it.

## Chapter 1: Introduction to Apache Spark: A Unified Analytics Engine

### Apache Spark's Distributed Execution

> At a high level in the Spark architecture, a Spark application consists of a
> driver program that is responsible for orchestrating parallel operations on
> the Spark cluster. The drive accesses the distributed components in the
> cluster (the Spark executors and cluster manager) through a `SparkSession`.

_-- Page 10_

**Spark driver**

> As the part of the Spark application responsible for instantiating a
> `SparkSession`, the Spark driver has multiple roles: it communicates with the
> cluster manager; it requests resources (CPU, memory, etc.) from the cluster
> manager for Spark's executors; and it transforms all the Spark operations into
> DAG computations

_-- Page 10_

**SparkSession**

> In Spark 2.0, the `SparkSession` became the unified conduit to all Spark
> operations and data. Not only did it subsume previous entry points to Spark
> like the `SparkContext`, `SQLContext`, `HiveContext`, `SparkConf`, and
> `StreamingContext`, but it also made working with Spark simpler and easier.

_-- Page 11_

> In a standalone Spark application, you can create a `SparkSession` using one
> of the APIs in the programming language of your choice. In the Spark shell the
> `SparkSession` is created for you, and you can access it via a global variable
> called `spark` or `sc`.

_-- Page 11_

**Cluster manager**

> The cluster manager is responsible for managing and allocating resources for
> the cluster of nodes on which your Spark application runs. Currently, Spark
> supports four cluster managers: the built-in standalone cluster manager,
> Apache Hadoop YARN, Apache Mesos, and Kubernetes.

_-- Page 12_

**Spark executor**

> A Spark executor runs on each worker node in the cluster. The executors
> communicate with the driver program and are responsible for executing tasks on
> the workers. In most deployments modes, only a single executor runs per node.

_-- Page 12_

### The Developer's Experience

**Data engineering tasks**

> The engineers use Spark because it provides a simple way to parallelize
> computations and hides all complexity of distribution and fault tolerance.
> This leaves them free to focus on using high-level DataFrame based APIs and
> domain-specific language queries to do ETL, reading and combining data from
> multiple sources.

## Chapter 2: Downloading Apache Spark and Getting Started

### Using the Local Machine

> Every computation expressed in high-level Structured APIs is decomposed into
> low-level optimized and generated RDD operations and then converted into Scala
> bytecode for the executors' JVMs. This generated RDD operation code is not
> accessible to users, nor is it the same as the user-facing RDD APIs.

_-- Page 25_

### Understanding Spark Application Concepts

> - Application: a user program built on Spark using its APIs. It consists of a
>   driver program and executors on the cluster.
> - SparkSession: an object that provides a point of entry to interact with
>   underlying Spark functionality and allows programming Spark with its APIs.
>   In an interactive shell, the Spark driver instantiates it for you, while in
>   a Spark application, you create it yourself.
> - Job: a parallel computation consisting of multiple tasks that get spawned in
>   response to a Spark action.
> - Stage: each job gets divided into smaller sets of tasks called stages that
>   depend on each other.
> - Task: a single unit of work or execution that will be sent to a Spark
>   executor.

_-- Page 26_

**Spark Jobs**

> During interactive sessions with Spark shells, the driver converts your Spark
> application into one or more Spark jobs. It then transforms each job into a
> DAG.

_-- Page 27_

**Spark Stages**

> As part of the DAG nodes, stages are created based on what operations can be
> performed serially or in parallel. Not all Spark operations can happen in a
> single stage, so they may be divided into multiple stages. Often stages are
> delineated on the operator's computation boundaries, where they dictate data
> transfer among Spark executors.

_-- Page 28_

**Spark Tasks**

> Each stage is comprised of Spark tasks (a unit of execution), which are then
> federated across each Spark executor; each task maps to a single core and
> works on a single partition of data.

_-- Page 28_

**Transformations, Actions, and Lazy Evaluation**

> Spark operations on distributed data can be classified into two types:
> _transformations_ and actions. Transformations, as the name suggests,
> transform a Spark DataFrame into a new DataFrame without altering the original
> data, giving it the property of immutability.

_-- Page 28_

> The actions and transformations contribute to a Spark query plan, [...].
> Nothing in a query plan is executed until an action is invoked.

_-- Page 29_

#### Narrow and Wide Transformations

> As noted, transformations are operations that Spark evaluates lazily. A huge
> advantage of the lazy evaluation scheme is that Spark can inspect your
> computational query and ascertain how it can optimize it. This optimization
> can be done by either joining or pipelining some operations and assigning them
> to a stage, or breaking them into stages by determining which operations
> require a shuffle or exchange of data across clusters.
>
> Transformations can be classified as having either narrow dependencies or wide
> dependencies. Any transformation where a single output partition can be
> computed from a single input partition is a _narrow_ transformation. For
> example [...] `filter()` and `container()` represent narrow transformations
> because they can operate on a single partition and produce the resulting
> output partition without any exchange of data.
>
> However, `groupBy()` or `orderBy()` instruct Spark to perform _wide_
> transformations, where data from other partition is read on, combined and
> written to disk.

_-- Page 30_

## Chapter 3: Apache Spar's Structured APIs

### Spark: What's Underneath an RDD?

> The RDD is the most basic abstraction in Spark. There are three vital
> characteristics associated with an RDD:
>
> - Dependencies
> - Partitions (with some locality information)
> - Compute function: Partition => `Iterator[T]`
>
> All three are integral to the simple RDD programming API model upon which all
> higher-level functionality is constructed. First, a list of _dependencies_
> that instructs Spark how an RDD is constructed with its inputs is required.
> When necessary to reproduce results, Spark can recreate an RDD from these
> dependencies and replicate operations on it. This characteristic gives RDDs
> resiliency.
>
> Second, _partitions_ provide Spark the ability to split the work parallelize
> computation on partitions across executors.
>
> And finally, an RDD has a _compute function_ that produces an `Iterator[T]`
> for the data that will be stored in the RDD.

_-- Page 44_

> [...] because it's unable to inspect the computation or expression in the
> function, Spark has no way to optimize the expression--it has no comprehension
> of its intention. And finally, Spark has no knowledge of the specific data
> type in `T`. TO Spark it's an opaque object; it has no idea if you are
> accessing a column of a certain type within an object. Therefore, all Spark
> can do is serialize the opaque object as a series of bytes, without using any
> data compression techniques.

_-- Page 44_

### Structuring Spark

#### Key merits and benefits

> Structure yields a number of benefits, including better performance and space
> efficiency across Spark components. [...] advantages: expressivity,
> simplicity, composability, and uniformity.

#### The DataFrame API

> Inspired by pandas DataFrames in structure, format and a few specific
> operations, Spark DataFrames are like distributed in-memory tables with named
> columns and schemas, where each column has a specific data type.

_-- Page 48_

> A _schema_ in Spark defines the column names and associated data types for a
> DataFrame. Defining a schema up front as opposed to taking a schema-on-read
> approach offers benefits.

_-- Page 50_

#### The Dataset API

> Conceptually, you can think of a DataFrame in Scala as an alias for a
> collection of generic objects, `Dataset[Row]`, where a `Row` is a generic
> untyped JVM object that may hold different types or fields. A Dataset, by
> contrast, is a collection of strongly typed JVM objects in Scala or a class in
> Java.

_-- Page 69_

#### DataFrames versus Datasets

> - If you want to tell Spark _what to do_, not _how to do it_, use DataFrames
>   or Datasets.
> - If you want rich semantics, high-level abstractions, and DLS operators, use
>   DataFrames or Datasets. If you want strict compile-time type safety and
>   don't mind creating multiple case classes for a specific `Dataset[T]`, use
>   Datasets.
> - If your processing demands high-level expression, filters, maps,
>   aggregations, computing averages or sums, SQL queries, columnar access, or
>   use of relational operators on semi-structured data, use DataFrames or
>   Datasets.
> - If your processing dictates relational transformations similar to SQL-like
>   queries, use DataFrames.
> - If you want to take advantage of and benefit from Tungsten¿s efficient
>   serialization with Encoders, use Datasets.
> - If you want unification, code optimization, and simplification of APIs
>   across Spark components, use DataFrames.
> - If you are a Python user, use DataFrames and drop down to RDDs if you need
>   more control.
> - If you want space and speed efficiency, use DataFrames.

_-- Pages 74-75_

#### Spark SQL and the Underlying Engine

> At a programmatic level, Spark SQL allows developers to issue ANSI SQL:2003
> compatible queries on structured data with a schema.

_-- Page 76_

> At the core of the Spark SQL engine are the Catalyst optimizer and Project
> Tungsten. Together, these support the high-level DataFrame and Dataset APIs
> and SQL queries.

_Page 77_

##### The Catalyst Optimizer

> The Catalyst optimizer takes a computational query and converts it into an
> execution plan. It goes through four transformational phases.

_Page 77_

> **Phase 1: analysis -** the Spark SQL engine beings by generating an abstract
> syntax tree for the SQL or DataFrame query. In this initial phase, any columns
> or table names will be resolved by consult an internal `Catalog`, a
> programmatic interface to Spark SQL that hold a list of names of columns, data
> types, functions, tables, databases, etc. Once they've all been successfully
> resolved, the query proceeds to the next phase.
>
> **Phase 2: Logical optimization -** this phase compromises two internal
> stages. Applying a standard rule based optimization approach, the Catalyst
> optimizer will first construct a set of multiple plans and then, using its
> cost-based optimizer, assign costs to each plan. This logical plan is the
> input into the physical plan.
>
> **Phase 3: Physical planning -** in this phase, Spark SQL generated an optimal
> physical plan for the selected logical plan.
>
> **Phase 4: Code Generation -** the final phase of query optimization involves
> generating efficient Java bytecode to run on each machine. Because Spark SQL
> can operate on data sets loaded in memory, Spark can use state-of-the-art
> compiler technology for code generation to speed up execution. In other words,
> it acts as a compiler. Project Tungsten, which facilites whole-stage code
> generation, plays a role here.
