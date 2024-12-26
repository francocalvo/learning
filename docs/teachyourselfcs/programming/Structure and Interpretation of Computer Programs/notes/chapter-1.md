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

```scheme
(define (cc amount kinds-of-coins)
  (cond ((= amount 0) 1)
        ((or (< amount 0) (= kinds-of-coins 0)) 0)
        (else (+ (cc amount
                     (- kinds-of-coins 1))
                 (cc (- amount
                       (first-denomination kinds-of-coins))
                     kinds-of-coins)))))

(define (first-denomination kinds-of-coins)
  (cond
    ((= kinds-of-coins 1 ) 1)
    ((= kinds-of-coins 2) 5)
    ((= kinds-of-coins 3) 10)
    ((= kinds-of-coins 4) 25)
    ((= kinds-of-coins 5) 50)))

(define (count-change amount)
  (cc amount  5))

(count-change 100) => 292
```

In this case, the important observation is that creating a recursive process to
solve this is more or less simple once we arrive at the algorithm, but creating
a iterative process for this is not very obvious.

## Section 1.3: Formulating Abstractions with Higher-Order Procedures

In Lisp-like languages: procedures themselves can be treated as data. This
allows us to create higher-order procedures — procedures that manipulate other
procedures as if they were simple data values.

This is important, as for example, with cubing, we could get away by always
using the computation `(* b b b)` to express the cube of `b`, but even tho we
could compute it, our language would not be able to express the concept of
cubing.

### Section 1.3.1

A key idea is that we can pass procedures as arguments to other procedures. For
example, consider the problem of summing the values of a function at integer
points over a given range. Instead of writing separate summation procedures for
different functions (e.g., summing squares, summing cubes, etc.), we can write a
single, general sum procedure that takes as an argument the function to be
applied:

```scheme
(define (sum term a next b)
  (if (> a b)
      0
      (+ (term a) (sum term (next a) next b))))
```

This allows as to create the sum of many other things, like integers,
pi-approximation, etc.

Also, it allows us to use it as a black box. If we build things on the `sum`
function, which is implemented as a recursive process, we wouldn't have to
change anything on top of it if we later change it for a iterative process.

### Section 1.3.2

- **Anonymous Procedures**: lambda allows the creation of procedures without
  naming them, which is essential for defining short-lived functions used in
  specific contexts without polluting the namespace.
- **Higher-Order Functions**: By enabling functions to be passed as arguments
  and returned as results, lambda facilitates the creation of more abstract and
  reusable code structures.

```scheme
(define (make-adder n)
  (lambda (x) (+ x n)))

(define add-five (make-adder 5))
(add-five 10) ; Returns 15

(define add-ten (make-adder 10))
(add-ten 10) ; Returns 20
```
