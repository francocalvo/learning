;; Section 2.2.2: Hierarchical Structures
;; Exercise 2.28

(define (fringe lst)
  (display "Fringe ")
  (display lst)
  (newline)
  (cond 
    ((not (pair? lst)) (list lst))
    ((null? (cdr lst)) (fringe (car lst)))
    (else
      (append
        (fringe (car lst))
        (fringe (cdr lst))))))

(define x (list (list 1 2) (list 3 4)))
(newline)
(display (fringe x)) 
