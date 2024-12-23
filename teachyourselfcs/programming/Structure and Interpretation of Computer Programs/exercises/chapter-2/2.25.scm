;; Section 2.2.2: Hierarchical Structures
;; Exercise 2.25

(define a (list 1 3 (list 5 7) 9))
(define b (list (list 7)))
(define c (list 1 (list 2 (list 3 (list 4 (list 5 (list 6 7)))))))


(car (cdaddr a))
(caar b)
(cadadr (cadadr (cadadr c))) 

(define x (list 1 2 3))
(define y (list 4 5 6))
