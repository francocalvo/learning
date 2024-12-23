;; Section 2.2.1: Representing sequences
;; Exercise 2.23

(define (for-each proc items)
  (cond ((null? items) #t)
        (else (proc (car items))
              (for-each proc (cdr items)))))

(for-each (lambda(x) (newline) (display x))
          (list 57 321 88))
