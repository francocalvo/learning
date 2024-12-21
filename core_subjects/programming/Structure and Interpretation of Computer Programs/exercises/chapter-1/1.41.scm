;; Section 1.3.4: Procedures as General Methods
;; Exercise 1.41

(define (double f)
  (lambda (x) (f (f x))))

(define (inc x) (+ x 1))

(((double (double double)) inc) 5) ;; 21
