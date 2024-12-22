;;;; utils.scm
;;;; SICP Utility Functions
;;;; A collection of fundamental procedures used throughout SICP

;;; Basic Mathematical Operations
;;; These simple functions appear throughout the book as building blocks

(define (square x)
  (* x x))

(define (cube x)
  (* x x x))

(define (average x y)
  (/ (+ x y) 2.0))

;;; Predicate Functions
;;; Used for testing various properties of numbers

(define (positive? x)
  (> x 0))

(define (negative? x)
  (< x 0))

(define (even? n)
  (= (remainder n 2) 0))

(define (odd? n)
  (not (even? n)))

;;; Higher-Order Functions
;;; These functions take other functions as arguments or return functions
;;; They form the backbone of functional programming patterns in SICP

(define (compose f g)
  (lambda (x)
    (f (g x))))

(define (repeated f n)
  (if (= n 1)
      f
      (compose f (repeated f (- n 1)))))

;;; Fixed Point Operations
;;; Used extensively in Chapter 1 for finding solutions to equations

(define tolerance 0.00001)

(define (fixed-point f first-guess)
  (define (close-enough? v1 v2)
    (< (abs (- v1 v2)) tolerance))
  (define (try guess)
    (let ((next (f guess)))
      (if (close-enough? guess next)
          next
          (try next))))
  (try first-guess))

(define (average-damp f)
  (lambda (x)
    (average x (f x))))

;;; Newton's Method and Related Functions
;;; These procedures implement Newton's method for finding roots

(define (deriv g)
  (let ((dx 0.00001))
    (lambda (x)
      (/ (- (g (+ x dx)) (g x))
         dx))))

(define (newton-transform g)
  (lambda (x)
    (- x (/ (g x)
            ((deriv g) x)))))

(define (newton-method g guess)
  (fixed-point (newton-transform g) guess))

;;; Accumulation Functions
;;; These are used for summing series and working with sequences

(define (sum term a next b)
  (if (> a b)
      0
      (+ (term a)
         (sum term (next a) next b))))

(define (product term a next b)
  (if (> a b)
      1
      (* (term a)
         (product term (next a) next b))))

(define (accumulate combiner null-value term a next b)
  (if (> a b)
      null-value
      (combiner (term a)
                (accumulate combiner null-value term (next a) next b))))

(define (filtered-accumulate combiner null-value term a next b predicate)
  (if (> a b)
      null-value
      (combiner (if (predicate a) (term a) null-value)
                (filtered-accumulate combiner null-value term (next a) next b predicate))))

;;; Common Building Blocks
;;; Utility functions that are frequently used in examples and exercises

(define (inc n) (+ n 1))
(define (dec n) (- n 1))
(define (identity x) x)

;;; Prime Numbers
;;; Functions for working with prime numbers, used in various exercises

(define (smallest-divisor n)
  (find-divisor n 2))

(define (find-divisor n test-divisor)
  (cond ((> (square test-divisor) n) n)
        ((divides? test-divisor n) test-divisor)
        (else (find-divisor n (+ test-divisor 1)))))

(define (divides? a b)
  (= (remainder b a) 0))

(define (prime? n)
  (= n (smallest-divisor n)))

;;; GCD
;;; Greatest Common Divisor implementation using Euclidean algorithm

(define (gcd a b)
  (if (= b 0)
      a
      (gcd b (remainder a b))))

;;; Rational numbers implementation
(define (make-rat n d)
  (let ((abn (abs n)) (abd (abs d)))
    (let ((a (gcd abn abd)))
      (if (> 0 (* n d))
        (cons (* -1 (/ abn a)) (/ abd a))
        (cons (/ abn a) (/ abd a))))))

(define (numer rn) (car rn))
(define (denom rn) (cdr rn))
