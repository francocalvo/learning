;; Section 2.2.1: Representing sequences
;; Exercise 2.18

(define (reverse p)
  (cond 
    ((null? p) nil)
    ((null? (cdr p)) (list (car p)))
    (else (append (reverse (cdr p)) 
                  (list (car p))))))

(display (reverse (list 23 72 149 34)))
