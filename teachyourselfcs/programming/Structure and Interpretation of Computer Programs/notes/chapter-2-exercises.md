# Chapter 2

## [Section 2.1.1](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-14.html#%_sec_2.1.1)

### [Exercise 2.1](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-14.html#%_thm_2.1)

```scheme
(load "../utils.scm") ;; gcd implementation

(define (make-rat n d)
  (let ((abn (abs n)) (abd (abs d)))
    (let ((a (gcd abn abd)))
      (if (> 0 (* n d))
        (cons (* -1 (/ abn a)) (/ abd a))
        (cons (/ abn a) (/ abd a))))))

(define (numer rn) (car rn))
(define (denom rn) (cdr rn))

(define t (make-rat 6 -18))
(numer t)
(denom t)
```

## [Section 2.1.2](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-14.html#%_sec_2.1.2)

### [Exercise 2.2](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-14.html#%_thm_2.2)

```scheme
(load "../utils.scm") ; average implementation

;; Point definition
(define (make-point x y) (cons x y))
(define (x-point p) (car p))
(define (y-point p) (cdr p))

;; Segment definition
(define (make-segment a b) (cons a b))
(define (segment-start s) (car s))
(define (segment-end s)  (cdr s))

(define (print-point p)
  (newline)
  (display "(")
  (display (x-point p))
  (display ",")
  (display (y-point p))
  (display ")"))

(define (midpoint l)
  (make-point 
    (average 
      (x-point (segment-start l))
      (x-point (segment-end l)))
    (average 
      (y-point (segment-start l))
      (y-point (segment-end l)))))


(define seg (make-segment (make-point 0 0) (make-point 10 10)))

(print-point (midpoint seg))
```

## [Section 2.1.3: What is meant by data](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-14.html#%_sec_2.1.3)

### [Exercise 2.04](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-14.html#%_thm_2.04)

```scheme
(define (cons x y)
  (lambda (m) (m x y)))

(define (cons s)
  (s (lambda (x y) x)))

(define (car z)
  (z (lambda (x y) y)))


(display (car (cons 1 2)))
(newline)
(display (cdr (cons 1 2)))
```

### [Exercise 2.05](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-14.html#%_thm_2.05)

```scheme
(define zero (lambda (f) (lambda (x) x)))

(define add-1 (lambda (n) (lambda (f) (lambda (x) (f ((n f) x))))))

(define one (lambda (f) (lambda (x) (f x))))

(define two (lambda (f) (lambda (x) (f (f x)))))
    
(define (sum a b)
  (lambda (f)
    (lambda (x)
      ((a f) ((b f) x)))))
```

## [Section 2.1.4: Extended Exercise: Interval Arithmetic](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-14.html#%_sec_2.1.4)

### [Exercise 2.07](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-14.html#%_thm_2.07)

```scheme
(define (make-interval a b) (cons a b))

(define add-interval
  (lambda (x y)
    (make-interval (+ (lower-bound x) (lower-bound y))
                   (+ (upper-bound x) (upper-bound y)))))

(define mul-interval
  (lambda (x y)
    (let ((p1 (* (lower-bound x) (lower-bound y)))
          (p2 (* (lower-bound x) (upper-bound y)))
          (p3 (* (upper-bound x) (lower-bound y)))
          (p4 (* (upper-bound x) (upper-bound y))))

      (make-interval (min p1 p2 p3 p4) (max p1 p2 p3 p4)))))

(define (div-interval x y)
  (mul-interval 
    x
    (make-interval 
      (/ 1.0 (upper-bound y))
      (/ 1.0 (lower-bound y)))))
```

My code..

```scheme
(define upper-bound
  (lambda (x)
    (cdr x)))

(define lower-bound
  (lambda (x)
    (car x)))
```

### [Exercise 2.08](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-14.html#%_thm_2.08)

```scheme
(load "./2.07.scm")

(define sub-interval
  (lambda (x y)
    (add-interval 
      x 
      (make-interval
        (* -1 (lower-bound y))
        (* -1 (upper-bound y))))))
```

### [Exercise 2.09](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-14.html#%_thm_2.09)

```scheme
(load "2.08.scm")

(define (with interval)
  (/ (- (upper-bound interval) (lower-bound interval))
     2.0))

(define a (make-interval 5 10))
(define b (make-interval 20 30))


(display (with a)) ;; 2.5
(display (with b)) ;; 5

(display (with (add-interval a b))) ;; 7.5
(display (with (sub-interval b a))) ;; 2.5

(display (with (mul-interval a b))) ;; 100
(display (with (div-interval b a))) ;; 2
```

## [Section 2.2.1: Representing sequences](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-15.html#%_sec_2.2.1)

### [Exercise 2.17](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-15.html#%_thm_2.17)

```scheme
(define (last-pair p)
  (if (null? (cdr p))
    (car p)
    (last-pair (cdr p))))

(last-pair (list 23 72 149 34))
```

### [Exercise 2.18](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-15.html#%_thm_2.18)

```scheme
(define (reverse p)
  (cond 
    ((null? p) nil)
    ((null? (cdr p)) (list (car p)))
    (else (append (reverse (cdr p)) 
                  (list (car p))))))

(display (reverse (list 23 72 149 34)))
```

### [Exercise 2.19](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-15.html#%_thm_2.19)

```scheme
(define (cc amount coins)
  (cond ((= amount 0) 1)
        ((or (< amount 0) (= (length coins) 0)) 0)
        (else (+ (cc amount
                     (cdr coins))
                 (cc (- amount
                       (car coins))
                     coins)))))

(define us-coins (list 1 5 10 25 50))
(define uk-coins (list 1 2 5 10 20 50 100))

(cc 100 us-coins)  ;;=> 292
```

### [Exercise 2.20](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-15.html#%_thm_2.20)

```scheme
(define (same-parity . l)
  (define (iter-sp res l rem)
    (if (null? l)
      res 
      (if 
        (= 
          (remainder (car l) 2)
          rem)
        (iter-sp 
          (append res (list (car l)))
          (cdr l)
          rem)
        (iter-sp 
          res 
          (cdr l) 
          rem))))

  (iter-sp (list (car l)) (cdr l) (remainder (car l) 2)))

(same-parity 1 2 3 4 5 6 7)
```

### [Exercise 2.21](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-15.html#%_thm_2.21)

```scheme
(load "../utils.scm") ; square definition

(define (square-list items)
  (map square items))

(square-list (list 1 2 3 4 5))
```

### [Exercise 2.23](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-15.html#%_thm_2.23)

```scheme
(define (for-each proc items)
  (cond ((null? items) #t)
        (else (proc (car items))
              (for-each proc (cdr items)))))

(for-each (lambda(x) (newline) (display x))
          (list 57 321 88))
```

## [Section 2.2.2: Hierarchical Structures](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-15.html#%_sec_2.2.2)

### [Exercise 2.25](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-15.html#%_thm_2.25)

```scheme
(define a (list 1 3 (list 5 7) 9))
(define b (list (list 7)))
(define c (list 1 (list 2 (list 3 (list 4 (list 5 (list 6 7)))))))


(car (cdaddr a))
(caar b)
(cadadr (cadadr (cadadr c))) 

(define x (list 1 2 3))
(define y (list 4 5 6))
```

### [Exercise 2.27](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-15.html#%_thm_2.27)

```scheme
(define (deep-reverse p)
  (cond 
    ((null? p) '())
    ((not (pair? p)) p)
    (else 
      (append 
        (deep-reverse (cdr p))
        (list (deep-reverse (car p)))))))

(display (deep-reverse (list 23 72 149 34)))

(define x (list (list 1 2) (list 3 4)))
(display x)
(display (deep-reverse x))
```

### [Exercise 2.28](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-15.html#%_thm_2.28)

```scheme
(define (fringe lst)
  (display "Fringe ")
  (display lst)
  (newline)
  (cond 
    ((not (pair? lst)) (list lst))
    ((null? (cdr lst)) (fringe (car lst)))
    (else
      (append
        (fringe (car lst))
        (fringe (cdr lst))))))

(define x (list (list 1 2) (list 3 4)))
(newline)
(display (fringe x))
```

### [Exercise 2.29](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-15.html#%_thm_2.29)

```scheme
(define (make-mobile left right)
  (list left right))

(define (make-branch length structure)
  (list length structure))

(define (is-mobile? obj)
  (if (and
        (pair? obj)
        (pair? (left-branch obj))
        (pair? (right-branch obj)))
    #t 
    #f))
```

Part a

```scheme
(define (left-branch mobile)
  (car mobile))

(define (right-branch mobile)
  (cadr mobile))

(define (branch-length branch)
  (car branch))

(define (branch-structure branch)
  (cadr branch))
```

Part b

```scheme
(define (total-weight structure)
  (if (is-mobile? structure)
    (+  
      (total-weight (branch-structure (left-branch structure)))
      (total-weight (branch-structure (right-branch structure))))
    structure))))
```

Part c

```scheme
(define (is-balanced? structure)
  (if (is-mobile? structure)
    (if (and
        (is-balanced? (branch-structure (left-branch structure)))
        (is-balanced? (branch-structure (right-branch structure)))
        (= 
          (* 
            (branch-length (left-branch structure))
            (total-weight (branch-structure (left-branch structure))))
          (* 
            (branch-length (right-branch structure))
            (total-weight (branch-structure (right-branch structure))))))
      #t
      #f)
    #t))
```

Example usage. This should output a weight of 26, and not be balanced.

```scheme
(define m (make-mobile
  (make-branch 10 (make-mobile (make-branch 5 10) (make-branch 25 2)))
  (make-branch 10 (make-mobile
                    (make-branch 10 (make-mobile (make-branch 5 10) (make-branch 25 2))) 
                    (make-branch 25 2)))))

(display "Example 1")
(newline)
(display (is-balanced? m))
(newline)
(display (total-weight m))
(newline)
```

This one should output a weight of 40, and be balanced.

```scheme
(define balanced (make-mobile
  (make-branch 4 (make-mobile
                   (make-branch 6 (make-mobile (make-branch 6 4) (make-branch 4 6))) 
                   (make-branch 6 (make-mobile (make-branch 6 4) (make-branch 4 6)))))
  (make-branch 4 (make-mobile
                   (make-branch 6 (make-mobile (make-branch 6 4) (make-branch 4 6))) 
                   (make-branch 6 (make-mobile (make-branch 6 4) (make-branch 4 6)))))))

(display "Example 2")
(newline)
(display (is-balanced? balanced))
(newline)
(display (total-weight balanced))
(newline)
```

### [Exercise 2.30](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-15.html#%_thm_2.30)

```scheme
(load "../utils.scm") ;; square, map definitions

(define (square-tree tree)
  (map (lambda (sub-tree)
         (if (pair? sub-tree)
           (square-tree sub-tree)
           (square sub-tree)))
       tree))

(square-tree (list 1
                   (list 2 (list 3 4) 5)
                   (list 6 7)))
```

### [Exercise 2.31](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-15.html#%_thm_2.31)

```scheme
(load "../utils.scm") ;; square, map definitions

(define (tree-map tree proc)
  (map (lambda (sub-tree)
         (if (pair? sub-tree)
           (tree-map sub-tree proc)
           (proc sub-tree)))
       tree))

(define (square-tree tree) (tree-map tree square))

(square-tree (list 1
                   (list 2 (list 3 4) 5)
                   (list 6 7)))
```

## [Section 2.2.3: Sequences as Conventional Interfaces](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-15.html#%_sec_2.2.3)

### [Exercise 2.33](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-15.html#%_thm_2.33)

```scheme
(load "../utils.scm") ;; acumulate definition
(define (ac_map p sequence)
  (accumulate 
    (lambda (x y) (cons (p x) y)) 
    '()
    sequence))

(define (ac_append seq1 seq2)
  (accumulate cons seq2 seq1))

(define (ac_length sequence)
  (accumulate (lambda (x y) (+ x 1)) 0 sequence))


(ac_map square (list 1 2 3))
(ac_append (list 1 2 3) (list 4 5 7))
(ac_length (list 1 2 3))
```

### [Exercise 2.34](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-15.html#%_thm_2.34)

```scheme
(load "../utils.scm")

(define (horner-eval x coefficient-sequence)
  (accumulate 
    (lambda (this-coeff higher-terms)
      (+ 
        this-coeff 
        (* higher-terms x)))
    0
    coefficient-sequence))


(horner-eval 2 (list 1 3 0 5 0 1)) ;; 79

(let ((x 2))
  (+ 1 (* 3 x) (* 5 (expt x 3)) (expt x 5))) ;; 79
```

### [Exercise 2.36](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-15.html#%_thm_2.36)

```scheme
(load "../utils.scm") ;; accumulate definition
```

Ok. So first we'll see the first solution I had.It's way too verbose but

it works perfectly.

```scheme

(define (select-n n seq)
  (if (= n 0)
    (car seq)
    (select-n (- n 1) (cdr seq))))

(define (enumerate-n n seq)
  (if (null? seq)
    '()
    (cons (select-n n (car seq)) (enumerate-n n (cdr seq)))))

(define (pop n seq)
  (cond 
    ((= n 0) seq)
    ((null? seq) '())
    (else (pop (- n 1) (cdr seq)))))

(define (pop-n seqs)
  (if (null? seqs) 
    '()
    (cons 
      (pop 1 (car seqs))
      (pop-n (cdr seqs)))))


(define (accumulate-n op init seqs)
  (if (null? (car seqs))
    '() 
    (cons (accumulate op init (enumerate-n 0 seqs))
          (accumulate-n op init (pop-n seqs)))))
```

Later, I came across this solution... It does the exact same thing, but

without creating a complex pair of functions to express the idea.

```scheme
(define (accumulate-n op init seqs)
  (if (null? (car seqs))
    '() 
    (cons (accumulate op init (map car seqs))
          (accumulate-n op init (map cdr seqs)))))

(accumulate-n + 0 (list (list 1 2 3) (list 4 5 6) (list 7 8 9) (list 10 11 12)))
```

### [Exercise 2.40](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-15.html#%_thm_2.40)

```scheme
(load "../utils.scm") ; flatmap definition

(define (unique-pairs n)
  (flatmap 
    (lambda (i)
      (map 
        (lambda (j) (list i j))
        (enumerate-interval 1 (- i 1))))
    (enumerate-interval 2 n)))

(unique-pairs 5)
```

### [Exercise 2.41](https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-15.html#%_thm_2.41)

```scheme
(load "../utils.scm")

(define (find-triple-s n s)
  (filter 
    (lambda (x) (= s (accumulate + 0 x)))
    (flatmap
      (lambda (x)
        (flatmap 
          (lambda (y) 
            (map 
              (lambda (z) (list x y z)) 
              (enumerate-interval 1 (- y 1))))  ; z goes from 1 to y-1
          (enumerate-interval 2 (- x 1))))      ; y goes from 2 to x-1
      (enumerate-interval 3 n))))               ; x goes from 3 to n

(find-triple-s 10 10)
```

