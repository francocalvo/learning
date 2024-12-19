;; Section 1.1.6 Conditional Expressions and Predicates
;; Exercise 1.3

(define (maxsum a b c) 
  (cond 
    ((> a b) (if (> b c) (+ a b) (+ a c)))
    ((> a c) (if (> b c) (+ a b) (+ a c)))
    (else (+ b c)))
    )

(display (maxsum 1 2 3))
