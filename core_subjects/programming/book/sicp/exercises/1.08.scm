(display "Exercise 1.8\n")
(display "================\n")

(display "Newton's method for cube root\n")
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

(display "Trying cuberoot of 27\n")
(display (exact->inexact (cuberoot 27)))

