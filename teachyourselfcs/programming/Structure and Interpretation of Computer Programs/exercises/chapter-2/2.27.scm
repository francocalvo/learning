;; Section 2.2.2: Hierarchical Structures
;; Exercise 2.27

(define (deep-reverse p)
  (cond 
    ((null? p) '())
    ((not (pair? p)) p)
    (else 
      (append 
        (deep-reverse (cdr p))
        (list (deep-reverse (car p)))))))

(display (deep-reverse (list 23 72 149 34)))

(define x (list (list 1 2) (list 3 4)))
(display x)
(display (deep-reverse x))
