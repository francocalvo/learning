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
  (if (null? (cdr p))
    (car p)
    (append (list (reverse (cdr p))) (car p))))

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

