(define (double a) (* a 2))
(define (halve a) (/ a 2))
(define (even? a) (= 0 (remainder a 2)))

(define (mult a b)
  (define (iter a b i)
    (cond
      ((= b 0) i)
      ((even? b) (iter (double a) (halves b) i))
      (else (iter a (- b 1) (+ i a)))))

  (iter a b 0))

(mult 8 7)
