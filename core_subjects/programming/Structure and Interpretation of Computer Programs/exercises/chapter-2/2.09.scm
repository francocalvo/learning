;; Section 2.1.4: Extended Exercise: Interval Arithmetic
;; Exercise 2.09

(load "2.08.scm")

(define (with interval)
  (/ (- (upper-bound interval) (lower-bound interval))
     2.0))

(define a (make-interval 5 10))
(define b (make-interval 20 30))


(display (with a)) ;; 2.5
(display (with b)) ;; 5

(display (with (add-interval a b))) ;; 7.5
(display (with (sub-interval b a))) ;; 2.5

(display (with (mul-interval a b))) ;; 100
(display (with (div-interval b a))) ;; 2
