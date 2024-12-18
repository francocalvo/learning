# Chapter 1

## Section 1.1.6 Conditional Expressions and Predicates

### Exercise 1.3: 

Define a procedure that takes three numbers as arguments and returns the sum of the squares of the two larger numbers.

```scm
(define (maxsum a b c)
  (cond
    ((> a b) (if (> b c) (+ a b) (+ a c)))
    ((> a c) (if (> b c) (+ a b) (+ a c)))
    (else (+ b c)))
    )

(display (maxsum 1 2 3))
```

### Exercise 1.8: 

Newton's method for cube roots is based on the fact that if y is an approximation to the cube root of x, then a better approximation is given by the value

```scm
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

### Exercise : 



```scm
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

### Exercise : 



```scm
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

### Exercise : 



```scm
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

### Exercise : 



```scm
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

### Exercise : 



```scm
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

### Exercise : 



```scm
(define (smallest-divisor n)
  (define (divides? a b)
    (= (remainder b a) 0))

  (define (find-divisor n test-divisor)
    (cond ((> (square test-divisor) n) n)
          ((divides? test-divisor n) test-divisor)
          (else (find-divisor n (+ test-divisor 1)))))

  (find-divisor n 2))

;; I guess this made sense back then, but I don't see any runtime difference.
(smallest-divisor 199)
(smallest-divisor 1999)
(smallest-divisor 19999)
```

### Exercise : 



```scm
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

### Exercise : 



```scm
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

### Exercise : 



```scm
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


;; Claude says this is more idiomatic.
;; ; Sum procedure - iteratively applies term from a to b
;; (define (sum term a next b)
;;   (if (> a b)
;;       0
;;       (+ (term a)
;;          (sum term (next a) next b))))
;;
;; ; Simpson's Rule integration implementation
;; (define (integral f a b n)
;;   ; Ensure n is even
;;   (if (odd? n)
;;       (error "n must be even")
;;
;;       (let ((h (/ (- b a) n)))
;;         ; Create coefficient function for Simpson's Rule
;;         (define (simpson-coefficient k)
;;           (cond ((or (= k 0) (= k n)) 1)
;;                 ((odd? k) 4)
;;                 (else 2)))
;;
;;         ; Function to evaluate at each point
;;         (define (term k)
;;           (* (simpson-coefficient k)
;;              (f (+ a (* k h)))))
;;
;;         ; Simple increment function
;;         (define (next x) (+ x 1))
;;
;;         ; Calculate integral using Simpson's Rule
;;         (* (/ h 3)
;;            (sum term 0 next n)))))

; Test function
(define (cube x) (* x x x))

; Convert to decimal and compute integral
(exact->inexact (integral cube 0 1 100000))

(define (cube x) (* x x x))

(integral cube 0 1 100)

```

### Exercise : 



```scm
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


;; Claude says this is more idiomatic.
;; ; Sum procedure - iteratively applies term from a to b
;; (define (sum term a next b)
;;   (if (> a b)
;;       0
;;       (+ (term a)
;;          (sum term (next a) next b))))
;;
;; ; Simpson's Rule integration implementation
;; (define (integral f a b n)
;;   ; Ensure n is even
;;   (if (odd? n)
;;       (error "n must be even")
;;
;;       (let ((h (/ (- b a) n)))
;;         ; Create coefficient function for Simpson's Rule
;;         (define (simpson-coefficient k)
;;           (cond ((or (= k 0) (= k n)) 1)
;;                 ((odd? k) 4)
;;                 (else 2)))
;;
;;         ; Function to evaluate at each point
;;         (define (term k)
;;           (* (simpson-coefficient k)
;;              (f (+ a (* k h)))))
;;
;;         ; Simple increment function
;;         (define (next x) (+ x 1))
;;
;;         ; Calculate integral using Simpson's Rule
;;         (* (/ h 3)
;;            (sum term 0 next n)))))

; Test function
(define (cube x) (* x x x))

; Convert to decimal and compute integral
(exact->inexact (integral cube 0 1 100000))

(define (cube x) (* x x x))

(integral cube 0 1 100)

```

