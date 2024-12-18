(define (even a) (= (remainder a 2) 0))

(define (exp b n)
  (define (iter a b n) 
    (cond 
      ((= n 0) a)
      ((even n) (iter a (* b b) (/ n 2)))   
      (else (iter (* a b) b (- n 1)))))

  (iter 1 b n))


(exp 3 3)
