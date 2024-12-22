;; Section 2.1.3: What is meant by data
;; Exercise 2.04

(define (cons x y)
  (lambda (m) (m x y)))

(define (cons s)
  (s (lambda (x y) x)))

(define (car z)
  (z (lambda (x y) y)))


(display (car (cons 1 2)))
(newline)
(display (cdr (cons 1 2)))


