;; Section 2.1.4: Extended Exercise: Interval Arithmetic
;; Exercise 2.07

;;md Alyssa P. Hacker code:
(define (make-interval a b) (cons a b))

(define add-interval
  (lambda (x y)
    (make-interval (+ (lower-bound x) (lower-bound y))
                   (+ (upper-bound x) (upper-bound y)))))

(define mul-interval
  (lambda (x y)
    (let ((p1 (* (lower-bound x) (lower-bound y)))
          (p2 (* (lower-bound x) (upper-bound y)))
          (p3 (* (upper-bound x) (lower-bound y)))
          (p4 (* (upper-bound x) (upper-bound y))))

      (make-interval (min p1 p2 p3 p4) (max p1 p2 p3 p4)))))

(define (div-interval x y)
  (mul-interval 
    x
    (make-interval 
      (/ 1.0 (upper-bound y))
      (/ 1.0 (lower-bound y)))))

;;md My code..
(define upper-bound
  (lambda (x)
    (cdr x)))

(define lower-bound
  (lambda (x)
    (car x)))

