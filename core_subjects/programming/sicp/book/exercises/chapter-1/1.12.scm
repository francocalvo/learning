;; Section 1.2.2 Tree Recursion
;; Exercise 1.12

(define (pascal col row)
  (cond ((= row 1) 1)
        ((= col 1) 1)
        ((= col row) 1)
        (else (
               + 
               (pascal (- col 1) (- row 1))
               (pascal col (- row 1))
               ))))

(pascal 3 5)

(pascal 3 5)
