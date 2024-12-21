;; Section 1.3.4: Procedures as General Methods
;; Exercise 1.42

(define (compose f g)
  (lambda (x) (f (g x))))

(define (square x) (* x x))
(define (inc x) (+ x 1))

((compose square inc) 6)
