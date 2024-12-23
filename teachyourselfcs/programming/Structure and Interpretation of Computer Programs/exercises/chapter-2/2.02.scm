;; Section 2.1.2
;; Exercise 2.2
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
