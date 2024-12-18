(define (smallest-divisor n)
  (define (divides? a b)
    (= (remainder b a) 0))

  (define (find-divisor n test-divisor)
    (cond ((> (square test-divisor) n) n)
          ((divides? test-divisor n) test-divisor)
          (else (find-divisor n (+ test-divisor 1)))))

  (find-divisor n 2))

;; I guess this made sense back then, but I don't see any runtime difference.
(smallest-divisor 199)
(smallest-divisor 1999)
(smallest-divisor 19999)
