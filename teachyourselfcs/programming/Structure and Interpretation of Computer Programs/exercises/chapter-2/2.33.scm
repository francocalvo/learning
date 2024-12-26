;; Section 2.2.3: Sequences as Conventional Interfaces
;; Exercise 2.33

(load "../utils.scm") ;; acumulate definition
(define (ac_map p sequence)
  (accumulate 
    (lambda (x y) (cons (p x) y)) 
    '()
    sequence))

(define (ac_append seq1 seq2)
  (accumulate cons seq2 seq1))

(define (ac_length sequence)
  (accumulate (lambda (x y) (+ x 1)) 0 sequence))


(ac_map square (list 1 2 3))
(ac_append (list 1 2 3) (list 4 5 7))
(ac_length (list 1 2 3))
