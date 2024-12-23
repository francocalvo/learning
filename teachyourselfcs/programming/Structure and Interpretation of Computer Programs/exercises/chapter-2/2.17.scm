;; Section 2.2.1: Representing sequences
;; Exercise 2.17

(define (last-pair p)
  (if (null? (cdr p))
    (car p)
    (last-pair (cdr p))))

(last-pair (list 23 72 149 34))
