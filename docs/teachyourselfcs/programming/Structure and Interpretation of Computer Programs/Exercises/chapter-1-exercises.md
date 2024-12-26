# Chapter 1

## [Section 1.1.6 Conditional Expressions and Predicates](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-10.html#%_sec_1.1.6)

### [Exercise 1.3](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-10.html#%_thm_1.3)

```scheme
(define (maxsum a b c) 
  (cond 
    ((> a b) (if (> b c) (+ a b) (+ a c)))
    ((> a c) (if (> b c) (+ a b) (+ a c)))
    (else (+ b c)))
    )

(display (maxsum 1 2 3))
```

## [Section 1.1.7 Example: Square Roots by Newton's Method](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-10.html#%_sec_1.1.7)

### [Exercise 1.8](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-10.html#%_thm_1.8)

```scheme
( define (cubert-iter guess x)
  ;; (display guess) (newline)
  (if (cube-good-enough? guess x)
    guess
    (cubert-iter (improve-cube-guess guess x) x)))

(define (cube-good-enough? guess x)
  (< 
    (abs (/ (- (* guess guess guess) x) x))
    0.1
  )
)

(define (improve-cube-guess guess x) 
  (/ 
    (+ (/ x (square guess)) (* 2 guess))
    3
  ))

(define (cuberoot x) (cubert-iter 1 x))

(exact->inexact (cuberoot 27))
(exact->inexact (cuberoot 125))
```

## [Section 1.2.2 Tree Recursion;](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-11.html#%_sec_1.2.2)

### [Exercise 1.11](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-11.html#%_thm_1.11)

```scheme
(define (f_r n)
  (cond ((< n 3) n)
        ((>= n 3)
         (+ 
           (f_r (- n 1))
           (* 2 (f_r (- n 2)))
           (* 3 (f_r (- n 3)))
           ))))

(define (f_i n)
  (define (f_i_iter a b c count)
      (cond 
        ((= count 0) (+ a (* 2 b) (* 3 c))) 
        ((> count 0) (f_i_iter (+ a (* 2 b) (* 3 c)) a b (- count 1)))))

  (cond ((< n 3) n)
        ((>= n 3)
          (f_i_iter 2 1 0 (- n 3)))))
                                 

(f_r 2)
(f_r 3)

(f_i 2)
(f_i 3)
```

## [Section 1.2.2 Tree Recursion](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-11.html#%_sec_1.2.2)

### [Exercise 1.12](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-11.html#%_thm_1.12)

```scheme
(define (pascal col row)
  (cond ((= row 1) 1)
        ((= col 1) 1)
        ((= col row) 1)
        (else (
               + 
               (pascal (- col 1) (- row 1))
               (pascal col (- row 1))
               ))))

(pascal 3 5)

(pascal 3 5)
```

## [Section 1.2.4 Exponentiation](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-11.html#%_sec_1.2.4)

### [Exercise 1.16](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-11.html#%_thm_1.16)

```scheme
(define (even a) (= (remainder a 2) 0))

(define (exp b n)
  (define (iter a b n) 
    (cond 
      ((= n 0) a)
      ((even n) (iter a (* b b) (/ n 2)))   
      (else (iter (* a b) b (- n 1)))))

  (iter 1 b n))


(exp 3 3)
```

### [Exercise 1.17](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-11.html#%_thm_1.17)

```scheme
(define (double a) (* a 2))
(define (halve a) (/ a 2))
(define (even? a) (= 0 (remainder a 2)))

(define (mult a b)
  (cond
    ((= b 1) a)
    ((even? b) (mult (double a) (halve b)))
    (else (+ a (mult a (- b 1))))))


(mult 8 7)
```

### [Exercise 1.18](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-11.html#%_thm_1.18)

```scheme
(define (double a) (* a 2))
(define (halve a) (/ a 2))
(define (even? a) (= 0 (remainder a 2)))

(define (mult a b)
  (define (iter a b i)
    (cond
      ((= b 0) i)
      ((even? b) (iter (double a) (halves b) i))
      (else (iter a (- b 1) (+ i a)))))

  (iter a b 0))

(mult 8 7)
```

## [Section 1.2.5 Greatest Common Divisors](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-11.html#%_sec_1.2.5)

### [Exercise 1.21](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-11.html#%_thm_1.21)

```scheme
(define (smallest-divisor n)
  (define (divides? a b)
    (= (remainder b a) 0))

  (define (find-divisor n test-divisor)
    (cond ((> (square test-divisor) n) n)
          ((divides? test-divisor n) test-divisor)
          (else (find-divisor n (+ test-divisor 1)))))

  (find-divisor n 2))
```

I guess this made sense back then, but I don't see any runtime difference.

```scheme

(smallest-divisor 199)
(smallest-divisor 1999)
(smallest-divisor 19999)
```

## [Section 1.2.6 Example: Testing for Primality](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-11.html#%_sec_1.2.6)

### [Exercise 1.22](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-11.html#%_thm_1.22)

```scheme
(define (square x)
  (* x x))

(define (smallest-divisor n)
  (define (divides? a b)
    (= (remainder b a) 0))
    
  (define (find-divisor n test-divisor)
    (cond ((> (square test-divisor) n) n)
          ((divides? test-divisor n) test-divisor)
          (else (find-divisor n (+ test-divisor 1)))))
          
  (find-divisor n 2))

(define (prime? n) 
  (if (= n 1) 
    #f
    (= (smallest-divisor n) n)))

(define (timed-prime-test n)
  (start-prime-test n (runtime)))

(define (start-prime-test n start-time)
  (if (prime? n)
      (report-prime n (- (runtime) start-time))
      #f))

(define (report-prime n elapsed-time)
  (display n)
  (display " is prime (")
  (display elapsed-time)
  (display " seconds)")
  (newline)
  #t)

(define (search-for-primes since until amount)
  (define start-time (runtime))
  
  (define (iter n until missing)
    (cond ((= missing 0) 
           (display "\nTotal time: ")
           (display (- (runtime) start-time))
           (display " seconds\n")
           'done)
          ((= (remainder n 2) 0) 
           (iter (+ n 1) until missing))
          (else 
           (if (timed-prime-test n)
               (iter (+ n 2) until (- missing 1))
               (iter (+ n 2) until missing)))))

  (display "Searching for ")
  (display amount)
  (display " primes between ")
  (display since)
  (display " and ")
  (display until)
  (display ":\n\n")
  
  (iter since until amount))

; Example searches at different ranges
(search-for-primes 100000000000 10000000000000000 3)

; Uncomment to test other ranges:
; (search-for-primes 10000 100000 3)
; (search-for-primes 100000 1000000 3)
; (search-for-primes 1000000 10000000 3)
```

### [Exercise 1.23](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-11.html#%_thm_1.23)

```scheme
(define (square x)
  (* x x))

(define (smallest-divisor n)
  (define (divides? a b)
    (= (remainder b a) 0))

  (define (next test-divisor) 
    (if (= test-divisor 2) 
        3
        (+ test-divisor 2)))

  (define (find-divisor n test-divisor)
    (cond ((> (square test-divisor) n) n)
          ((divides? test-divisor n) test-divisor)
          (else (find-divisor n (next test-divisor)))))

  (find-divisor n 2))

(define (prime? n) 
  (= (smallest-divisor n) n))

(define (timed-prime-test n)
  (start-prime-test n (runtime)))

(define (start-prime-test n start-time)
  (if (prime? n)
      (report-prime n (- (runtime) start-time))
      #f))

(define (report-prime n elapsed-time)
  (display n)
  (display " is prime (")
  (display elapsed-time)
  (display " seconds)")
  (newline)
  #t)

(define (search-for-primes since until amount)
  (define start-time (runtime))
  
  (define (iter n until missing)
    (cond ((= missing 0) 
           (display "\nTotal time: ")
           (display (- (runtime) start-time))
           (display " seconds\n")
           'done)  ; Return a symbol instead of undefined
          ((= (remainder n 2) 0) 
           (iter (+ n 1) until missing))
          (else 
           (if (timed-prime-test n)
               (iter (+ n 2) until (- missing 1))
               (iter (+ n 2) until missing)))))

  (display "Searching for ")
  (display amount)
  (display " primes between ")
  (display since)
  (display " and ")
  (display until)
  (display ":\n\n")
  
  (iter since until amount))

; Example usage:
(search-for-primes 100000000000 10000000000000000 3)
```

## [Section 1.3.1 Procedures as Arguments](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-12.html#%_sec_1.3.1)

### [Exercise 1.29](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-12.html#%_thm_1.29)

```scheme
(define (sum term a next b)
  (if (> a b) 
    0 
    (+ (term a) 
       (sum term (next a) next b))))

;; Simpson's Rule integration implementation
(define (integral f a b n)
  ;; Define h
  (let ((h (/ (- b a) n)))

  ;; Wrap the term in a lambda to pass the h and a values
  (define (fs f h a) 
    (lambda (x) 
      (*
        ;; I guess the diference between this and the other implementation is that
        ;; I don't create the extra function to calculate the coefficient.
        (cond ((or (= x 0) (= x n)) 1)
              ((odd? x) 4)
              (else 2))
        (f (+ a (* x h))))))


  ;; Increment function
  (define (next a) (+ a 1))

  ;; Calculate the integral
  (*
    (/ h 3)
    (sum (fs f h a) 0 next n))))

; Test function
(define (cube x) (* x x x))

; Convert to decimal and compute integral
(exact->inexact (integral cube 0 1 100000))
(integral cube 0 1 100)
```

### [Exercise 1.30](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-12.html#%_thm_1.30)

```scheme
(define (sum term a next b)
  (define (iter a result)
    (if (> a b) 
      result
      (iter (next a) (+ result (term a)))))

  (iter a 0))

;; Simpson's Rule integration implementation
(define (integral f a b n)
  ;; Define h
  (let ((h (/ (- b a) n)))

  ;; Wrap the term in a lambda to pass the h and a values
  (define (fs f h a) 
    (lambda (x) 
      (*
        ;; I guess the diference between this and the other implementation is that
        ;; I don't create the extra function to calculate the coefficient.
        (cond ((or (= x 0) (= x n)) 1)
              ((odd? x) 4)
              (else 2))
        (f (+ a (* x h))))))


  ;; Increment function
  (define (next a) (+ a 1))

  ;; Calculate the integral
  (*
    (/ h 3)
    (sum (fs f h a) 0 next n))))

; Test function
(define (cube x) (* x x x))

; Convert to decimal and compute integral
(exact->inexact (integral cube 0 1 100000))
(integral cube 0 1 100)
```

### [Exercise 1.31](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-12.html#%_thm_1.31)

```scheme
(define (product term a next b)
  (define (iter a result)
    (if (> a b) 
      result
      (iter (next a) (* result (term a)))))

  (iter a 1))

(define (factorial n)
  (define (identity x) x)
  (define (inc x) (+ x 1))
  (product identity 1 inc n))

(display (factorial 5))
```

For the second part of the exercise, we need to define a term that will

calculate the value of the formula given in the exercise. The formula is

$\pi/4=(2*4*4*6*6*8*...)/(3*3*5*5*7*7*...)$.

```scheme

(define (pi-product n)
  (define (inc x) (+ x 1))
  (define (pi-step x)
    (cond 
      ((= x 0) 1)
      ((= x 1) 1)
      ((even? x)
       ;; (display "even ")
       ;; (display x)
       ;; (display " / ")
       ;; (display (+ x 1))
       ;; (newline)
       (/ x (+ x 1)))
      (else
       ;; (display "odd ")
       ;; (display (+ x 1))
       ;; (display " / ")
       ;; (display x)
       ;; (newline)
       ;; (newline)
       (/ (+ x 1) x))))


  (product pi-step 0 inc n))

(display (exact->inexact(pi-product 10000)))
```

### [Exercise 1.32](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-12.html#%_thm_1.32)

```scheme
(define (accumulate combiner null-value term a next b)
  (define (iter a result)
    (if (> a b)
      result
      (iter (next a) (combiner result (term a)))))
  (iter a null-value))

(define (accumulate_rec combiner null-value term a next b)
  (if (> a b)
    null-value
    (combiner (term a) (accumulate_rec combiner null-value term (next a) next b))))


(define (product term a next b)
  (define (prod a b) (* a b))
  (accumulate_rec prod 1 term a next b))

(define (factorial n)
  (define (identity x) x)
  (define (inc x) (+ x 1))
  (product identity 1 inc n))

(display (factorial 5))
```

### [Exercise 1.33](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-12.html#%_thm_1.33)

```scheme
the sum of the squares of the prime numbers in the interval a
 to b
 (assuming that you have a prime? predicate already written)

(load "1.22.scm")

(define (accumulate_filter combiner null-value term a next b filter)
  (define (iter a result)
    (if (> a b)
      result
      (let ((a-value (term a)))
        (if (filter a)
          (iter (next a) (combiner result a-value))
          (iter (next a) result)))))
  (iter a null-value))

(define (prime_sum a b)
  (define (inc x) (+ x 1))
  (define (identity x) x)
  (accumulate_filter + 0 identity a inc b prime?))

(display (prime_sum 0 100))
```

## [Section 1.3.2 Constructing Procedures Using Lambda](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-12.html#%_sec_1.3.2)

### [Exercise 1.34](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-12.html#%_thm_1.34)

```scheme
(define (f g) (g 2))
(define (square x) (* x x))

(f square) ; 4

(f (lambda (z) (* z (+ z 1)))) ; 6

(f f)
```

The interpreter will evaluate the combination `(f f)` by substituting `f`

with its definition.

We'll go: (f f) --> (f 2) --> (2 2) --> error

## [Section 1.3.3 Procedures as General Methods](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-12.html#%_sec_1.3.3)

### [Exercise 1.36](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-12.html#%_thm_1.36)

```scheme
(define tolerance 0.00001)

(define (fixed-point f first-guess)
  (define (close-enough? v1 v2)
    (< (abs (- v1 v2)) tolerance))

  (define (try guess)
    (newline)
    (display guess)
    (let ((next (f guess)))
      (if (close-enough? guess next)
        next
        (try next))))

  (try first-guess))

(define (average x y) (/ (+ x y) 2))
(define (average-damp f)
  (lambda (x) (average x (f x))))

(define (func x) (/ (log 1000) (log x)))

(fixed-point func 2.0) ;; 35 steps

(fixed-point (average-damp func) 2.0) ; 9 steps
```

### [Exercise 1.37](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-12.html#%_thm_1.37)

```scheme
(define (cont-fract n d k)
  (if (= k 0)
      0
      (/ (n k)  (+ (d k) (cont-fract n d (- k 1))))))


(define (n k) 1.0)
```

The desired value is: 0.6180

It can be achieved by calling with k = 11

```scheme
(cont-fract n n 11) ; 
(cont-fract n n 11)
```

### [Exercise 1.38](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-12.html#%_thm_1.38)

```scheme
(define (cont-fract n d k)
  (define (cont-fract-inter n d k i)
    (if (= i k)
        0
        (/ (n i) (+ (d i) (cont-fract-inter n d k (+ i 1))))))
  (cont-fract-inter n d k 1)) ; start counting from 1


(define (n k) 1.0)
(define (d k)
  (if (= (remainder k 3) 2)
      (* 2 (/ (+ k 1) 3))
      1.0))
```

The desired value is $\e - 2$ which is approximately 0.71828

```scheme
(cont-fract n d 1000)
```

### [Exercise 1.39](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-12.html#%_thm_1.39)

```scheme
(define (cont-fract combinator null-value n d k)
  (define (cont-fract-inter n d k i)
    (if (= i k)
        0
        (/ (n i) (combinator (d i) (cont-fract-inter n d k (+ i 1))))))
  (cont-fract-inter n d k null-value))

;; Define the continued fraction for tan(x)
(define (tan-cf x k)
  (define (n k) (square x))      ; Numerator is always x^2
  (define (d k) (+ (* 2 k) 1))   ; Denominator follows 2k + 1 (i.e., 1, 3, 5, ...)
  (/ x
     (- 1
        (cont-fract - 1 n d k))))


;; tan(1) = 1.5574077

(display (exact->inexact (tan-cf 1 100)))
(newline)
(display "done")
```

## [Section 1.3.4: Procedures as General Methods](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-12.html#%_sec_1.3.4)

### [Exercise 1.40](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-12.html#%_thm_1.40)

```scheme
(define (cubic a b c) (lambda (x) (+ (* x x x) (* a x x) (* b x) c)))
```

### [Exercise 1.41](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-12.html#%_thm_1.41)

```scheme
(define (double f)
  (lambda (x) (f (f x))))

(define (inc x) (+ x 1))

(((double (double double)) inc) 5) ;; 21
```

### [Exercise 1.42](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-12.html#%_thm_1.42)

```scheme
(define (compose f g)
  (lambda (x) (f (g x))))

(define (square x) (* x x))
(define (inc x) (+ x 1))

((compose square inc) 6)
```

### [Exercise 1.43](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-12.html#%_thm_1.43)

```scheme
(define (compose f g)
  (lambda (x) (f (g x))))

(define (square x) (* x x))
(define (inc x) (+ x 1))

((compose square inc) 6)

(define (repeated f n)
  (define (iter res m)
    (if (= m 0)
      res
      (iter (f res) (- m 1))))

  (lambda (x) (iter x n)))

((repeated square 2) 5)
(define (compose f g)
  (lambda (x) (f (g x))))

(define (square x) (* x x))
(define (inc x) (+ x 1))

((compose square inc) 6)
```

### [Exercise 1.44](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-12.html#%_thm_1.44)

```scheme
(define dx 0.000001)

(define soothe
  (lambda (f)
    (lambda (x)
      (/ (+ (f (- x dx))
            (f x)
            (f (+ x dx)))
         3))))

;; I won't repeat the repeated procedure here, but it seems very straightforward.
```

