;; Section 2.2.3: Sequences as Conventional Interfaces
;; Exercise 2.40

(load "../utils.scm") ; flatmap definition

(define (unique-pairs n)
  (flatmap 
    (lambda (i)
      (map 
        (lambda (j) (list i j))
        (enumerate-interval 1 (- i 1))))
    (enumerate-interval 2 n)))

(unique-pairs 5)
