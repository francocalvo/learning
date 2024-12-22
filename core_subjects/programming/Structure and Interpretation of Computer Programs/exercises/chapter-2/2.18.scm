;; Section 2.2.1: Representing sequences
;; Exercise 2.18

(define (reverse p)
  (if (null? (cdr p))
    (car p)
    (append (list (reverse (cdr p))) (car p))))

(display (reverse (list 23 72 149 34)))

