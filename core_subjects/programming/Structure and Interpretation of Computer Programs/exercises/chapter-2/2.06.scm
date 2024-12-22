;; Section 2.1.3: What is meant by data
;; Exercise 2.05

;;md ChatGPT made me understand what a Church number is. 
;;md I wrote these, but don't ask WTF it means. Dark magic.
(define zero (lambda (f) (lambda (x) x)))

(define add-1 (lambda (n) (lambda (f) (lambda (x) (f ((n f) x))))))

(define one (lambda (f) (lambda (x) (f x))))

(define two (lambda (f) (lambda (x) (f (f x)))))
    
(define (sum a b)
  (lambda (f)
    (lambda (x)
      ((a f) ((b f) x)))))

