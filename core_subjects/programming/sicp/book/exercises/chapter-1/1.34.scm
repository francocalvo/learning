;; Section 1.3.2 Constructing Procedures Using Lambda
;; Exercise 1.34: 
;; Suppose we define the procedure
;; ```scheme
;; (define (f g) (g 2))
;; ```
;; Then we have
;; ```scheme
;; (f square)
;; 4
;; (f (lambda (z) (* z (+ z 1)))) 6
;; (f (lambda (z) (* z (+ z 1))))
;;
;; What happens if we (perversely) ask the interpreter to evaluate the 
;; combination `(f f )`? Explain.


(define (f g) (g 2))
(define (square x) (* x x))

(f square) ; 4

(f (lambda (z) (* z (+ z 1)))) ; 6

(f f)

;;md The interpreter will evaluate the combination `(f f)` by substituting `f` 
;;md with its definition. 
;;md We'll go: (f f) --> (f 2) --> (2 2) --> error
