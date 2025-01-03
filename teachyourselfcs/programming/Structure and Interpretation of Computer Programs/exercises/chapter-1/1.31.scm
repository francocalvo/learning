;; Section 1.3.1 Procedures as Arguments
;; Exercise 1.31

(define (product term a next b)
  (define (iter a result)
    (if (> a b) 
      result
      (iter (next a) (* result (term a)))))

  (iter a 1))

(define (factorial n)
  (define (identity x) x)
  (define (inc x) (+ x 1))
  (product identity 1 inc n))

(display (factorial 5))

;;md For the second part of the exercise, we need to define a term that will
;;md calculate the value of the formula given in the exercise. The formula is
;;md $\pi/4=(2*4*4*6*6*8*...)/(3*3*5*5*7*7*...)$. 

(define (pi-product n)
  (define (inc x) (+ x 1))
  (define (pi-step x)
    (cond 
      ((= x 0) 1)
      ((= x 1) 1)
      ((even? x)
       ;; (display "even ")
       ;; (display x)
       ;; (display " / ")
       ;; (display (+ x 1))
       ;; (newline)
       (/ x (+ x 1)))
      (else
       ;; (display "odd ")
       ;; (display (+ x 1))
       ;; (display " / ")
       ;; (display x)
       ;; (newline)
       ;; (newline)
       (/ (+ x 1) x))))


  (product pi-step 0 inc n))

(display (exact->inexact(pi-product 10000)))
