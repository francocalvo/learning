;; Section 1.3.1 Procedures as Arguments
;; Exercise 1.29
;; Simpson's Rule is a more accurate method of numerical integration than the 
;; method illustrated above. Using Simpson's Rule, the integral of a function 
;; f between a and b is approximated as
;; $3h[YO +4YI + 2Y2 +4Y3 +2Y4 +...+2Yn-2 +4Yn-1 +Yn]$
;; where $h= (b-a)/n$,for some even integer n,and $Y_k = f(a+kh)$.
;;
;; (Increasing n increases the accuracy of the approximation.) 
;;
;; Define a procedure that takes as arguments f, a, b, and n and returns the 
;; value of the integral, computed using Simpson's Rule. Use your procedure to 
;; integrate cube between 0 and 1 (with n = 100 and n = 1000), and compare the
;; results to those of the integral procedure shown above.

(define (sum term a next b)
  (if (> a b) 
    0 
    (+ (term a) 
       (sum term (next a) next b))))

;; Simpson's Rule integration implementation
(define (integral f a b n)
  ;; Define h
  (let ((h (/ (- b a) n)))

  ;; Wrap the term in a lambda to pass the h and a values
  (define (fs f h a) 
    (lambda (x) 
      (*
        ;; I guess the diference between this and the other implementation is that
        ;; I don't create the extra function to calculate the coefficient.
        (cond ((or (= x 0) (= x n)) 1)
              ((odd? x) 4)
              (else 2))
        (f (+ a (* x h))))))


  ;; Increment function
  (define (next a) (+ a 1))

  ;; Calculate the integral
  (*
    (/ h 3)
    (sum (fs f h a) 0 next n))))

; Test function
(define (cube x) (* x x x))

; Convert to decimal and compute integral
(exact->inexact (integral cube 0 1 100000))
(integral cube 0 1 100)

