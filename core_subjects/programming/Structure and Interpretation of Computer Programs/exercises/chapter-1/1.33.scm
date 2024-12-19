;; Section 1.3.1 Procedures as Arguments
;; Exercise 1.33

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
