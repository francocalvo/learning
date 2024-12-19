;; Section 1.3.3 Procedures as General Methods
;; Exercise 1.39: 

;;md I create a more abstracted version of the cont-frac function.
(define (cont-fract combinator null-value n d k)
  (define (cont-fract-inter n d k i)
    (if (= i k)
        0
        (/ (n i) (combinator (d i) (cont-fract-inter n d k (+ i 1))))))
  (cont-fract-inter n d k null-value))

;; Define the continued fraction for tan(x)
(define (tan-cf x k)
  (define (n k) (square x))      ; Numerator is always x^2
  (define (d k) (+ (* 2 k) 1))   ; Denominator follows 2k + 1 (i.e., 1, 3, 5, ...)
  (/ x
     (- 1
        (cont-fract - 1 n d k))))


;; tan(1) = 1.5574077

(display (exact->inexact (tan-cf 1 100)))
(newline)
(display "done")

