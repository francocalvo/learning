;; Section 1.3.1 Procedures as Arguments
;; Exercise 1.30
;; The sum procedure above generates a linear recursion. 
;; The procedure can be rewritten so that the sum is performed iteratively

(define (sum term a next b)
  (define (iter a result)
    (if (> a b) 
      result
      (iter (next a) (+ result (term a)))))

  (iter a 0))

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
