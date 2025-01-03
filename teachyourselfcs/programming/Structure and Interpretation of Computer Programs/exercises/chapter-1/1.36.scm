;; Section 1.3.3 Procedures as General Methods
;; Exercise 1.36

(define tolerance 0.00001)

(define (fixed-point f first-guess)
  (define (close-enough? v1 v2)
    (< (abs (- v1 v2)) tolerance))

  (define (try guess)
    (newline)
    (display guess)
    (let ((next (f guess)))
      (if (close-enough? guess next)
        next
        (try next))))

  (try first-guess))

(define (average x y) (/ (+ x y) 2))
(define (average-damp f)
  (lambda (x) (average x (f x))))

(define (func x) (/ (log 1000) (log x)))

(fixed-point func 2.0) ;; 35 steps

(fixed-point (average-damp func) 2.0) ; 9 steps


