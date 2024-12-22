;; Section 2.1.4: Extended Exercise: Interval Arithmetic
;; Exercise 2.08

(load "./2.07.scm")

(define sub-interval
  (lambda (x y)
    (add-interval 
      x 
      (make-interval
        (* -1 (lower-bound y))
        (* -1 (upper-bound y))))))
