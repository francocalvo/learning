;; Section 1.3.1 Procedures as Arguments
;; Exercise 1.32

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
