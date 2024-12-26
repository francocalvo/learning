;; Section 2.2.3: Sequences as Conventional Interfaces
;; Exercise 2.41

(load "../utils.scm")

(define (find-triple-s n s)
  (filter 
    (lambda (x) (= s (accumulate + 0 x)))
    (flatmap
      (lambda (x)
        (flatmap 
          (lambda (y) 
            (map 
              (lambda (z) (list x y z)) 
              (enumerate-interval 1 (- y 1))))  ; z goes from 1 to y-1
          (enumerate-interval 2 (- x 1))))      ; y goes from 2 to x-1
      (enumerate-interval 3 n))))               ; x goes from 3 to n

(find-triple-s 10 10)
