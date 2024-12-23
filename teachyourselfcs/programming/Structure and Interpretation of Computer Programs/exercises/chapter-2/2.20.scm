;; Section 2.2.1: Representing sequences
;; Exercise 2.20

(define (same-parity . l)
  (define (iter-sp res l rem)
    (if (null? l)
      res 
      (if 
        (= 
          (remainder (car l) 2)
          rem)
        (iter-sp 
          (append res (list (car l)))
          (cdr l)
          rem)
        (iter-sp 
          res 
          (cdr l) 
          rem))))

  (iter-sp (list (car l)) (cdr l) (remainder (car l) 2)))

(same-parity 1 2 3 4 5 6 7)
