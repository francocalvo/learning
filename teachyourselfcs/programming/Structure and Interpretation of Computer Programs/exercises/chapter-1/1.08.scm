;; Section 1.1.7 Example: Square Roots by Newton's Method
;; Exercise 1.8

( define (cubert-iter guess x)
  ;; (display guess) (newline)
  (if (cube-good-enough? guess x)
    guess
    (cubert-iter (improve-cube-guess guess x) x)))

(define (cube-good-enough? guess x)
  (< 
    (abs (/ (- (* guess guess guess) x) x))
    0.1
  )
)

(define (improve-cube-guess guess x) 
  (/ 
    (+ (/ x (square guess)) (* 2 guess))
    3
  ))

(define (cuberoot x) (cubert-iter 1 x))

(exact->inexact (cuberoot 27))
(exact->inexact (cuberoot 125))

