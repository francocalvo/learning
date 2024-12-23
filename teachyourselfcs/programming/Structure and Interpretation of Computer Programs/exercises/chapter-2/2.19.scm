;; Section 2.2.1: Representing sequences
;; Exercise 2.19

;;md I finished this before finishing reading the exercise.
;;md I won't be using the other way is it's pretty much the same buth with
;;md last-pair.
;;md To answer the last questino, yes, the order of the list matters.

(define (cc amount coins)
  (cond ((= amount 0) 1)
        ((or (< amount 0) (= (length coins) 0)) 0)
        (else (+ (cc amount
                     (cdr coins))
                 (cc (- amount
                       (car coins))
                     coins)))))

(define us-coins (list 1 5 10 25 50))
(define uk-coins (list 1 2 5 10 20 50 100))

(cc 100 us-coins)  ;;=> 292
