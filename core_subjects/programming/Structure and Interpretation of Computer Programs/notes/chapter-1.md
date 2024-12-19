# Chapter 1

## Section 1.1: The elements of programming

### Subsection 1.1.1

- Every powerful language has these three mechanisms:
  - **Primitive expressions**, which represent the simplest entities the
    language is concerned with.
  - **Means of combinations**, by which compound elements are built from simpler
    ones.
  - **Means of abstraction** , by which compound elements can be named and
    manipulated as units.

In Scheme we can define a variable as:

### Subsection 1.1.4

```scheme
(define a 1)
```

Or a function like this:

```scheme
(define (square x) (* x x))
```

It's important to note how the functions are called:

```scheme
(square 2) ;; 4
```

### Subsection 1.1.5

The way we can think about how the interpreter executes ours commands can be
simplified with the **substitution model** . This means:

> To apply a compound procedure to arguments, evaluate the body of the procedure
> with each formal parameter replaced by the corresponding argument.

Also, it's interesting do quote this:

> [The] alternative "fully expand and then reduce" evaluation method is known as
> _normal-order evaluation_.

This is contrary to what Lisp uses, which is applicative-order evaluation. This
allows to reduce redundant calculations.

### Subsection 1.1.8

Using a decomposition strategy, we can take apart any problem in more atomic
subproblems. For these, we then can create procedures that solved them, and use
these as 'black boxes' that we use in other places, but don't care about the
internal implementation.

One of the things we do care about in a procedure are it's parameters, which
specify the expected input for the procedure:

> A formal parameter of a procedure has a very special role in the procedure
> definition, in that it doesn't matter what name the formal parameter has. Such
> a name is called a _bound variable_, and we say that the procedure definition
> _binds_ its formal parameters. The set of expression for which a binding
> defines a name is called the _scope_ of that name.

## Section 1.2: Procedures and the processes they generate

### Subsection 1.2.1:

So far, we've seen recursion. But in Lisp, not all recursion is born the same.
We can make the distinction in a procedure that's recursive, but the process
itself is not.

We can see two types of processes in this space: recursive processes and
iterative processes. Both created by recursive procedures.

Consider this implementation of a factorial using a recursive process:

```scheme
(define (factorial n)
  (if (= n 0)
      1
      (* n (factorial (- n 1)))))
```

On the other hand, we can also have a similar process that's going to be
iterative. This is done by passing along partial results to the next iteration.

```scheme
(define (factorial n)
  (define (iter product counter)
    (if (> counter n)
        product
        (iter (* counter product)
              (+ counter 1))))
  (iter 1 1))
```

Normally, recursive processes are easier to come up with, as they follow a more
"mathematical" definition, as seen with the factorial definition. In most cases,
a iterative process is going to have a smaller space complexity.

### Subsection 1.2.2

I'll write down the solution for the example of counting change, as it's a neat
way to see how we can create a iterative process while also traversing a tree of
options.

We can ask: how many different way can we make change of $1.00, given
half-dollars, quarters, dimes, nickels and pennies?

To solve this, we can see that the full range of possibilities can be thought of
as the sum of all possibilities where _N_ is in the solution, (for example, a
half-dollars), and all the possibilities that don't have _N_.




