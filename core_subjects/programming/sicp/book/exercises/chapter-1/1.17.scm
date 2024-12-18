(define (double a) (* a 2))
(define (halve a) (/ a 2))
(define (even? a) (= 0 (remainder a 2)))

(define (mult a b)
  (cond
    ((= b 1) a)
    ((even? b) (mult (double a) (halve b)))
    (else (+ a (mult a (- b 1))))))


(mult 8 7)
