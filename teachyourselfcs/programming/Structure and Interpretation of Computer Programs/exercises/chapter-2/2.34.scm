;; Section 2.2.3: Sequences as Conventional Interfaces
;; Exercise 2.34

(load "../utils.scm")

(define (horner-eval x coefficient-sequence)
  (accumulate 
    (lambda (this-coeff higher-terms)
      (+ 
        this-coeff 
        (* higher-terms x)))
    0
    coefficient-sequence))


(horner-eval 2 (list 1 3 0 5 0 1)) ;; 79

(let ((x 2))
  (+ 1 (* 3 x) (* 5 (expt x 3)) (expt x 5))) ;; 79
