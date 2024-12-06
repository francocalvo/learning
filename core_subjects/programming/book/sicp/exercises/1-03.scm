(define (maxsum a b c) 
  (cond 
    ((> a b) (if (> b c) (+ a b) (+ a c)))
    ((> a c) (if (> b c) (+ a b) (+ a c)))
    (else (+ b c)))
    )

(maxsum 1 2 3)
