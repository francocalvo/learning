;; Section 2.2.1: Representing sequences
;; Exercise 2.21

(load "../utils.scm") ; square definition

(define (square-list items)
  (map square items))

(square-list (list 1 2 3 4 5))

