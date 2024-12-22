;; Section 2.1.1
;; Exercise 2.1
(load "../utils.scm") ;; gcd implementation

(define (make-rat n d)
  (let ((abn (abs n)) (abd (abs d)))
    (let ((a (gcd abn abd)))
      (if (> 0 (* n d))
        (cons (* -1 (/ abn a)) (/ abd a))
        (cons (/ abn a) (/ abd a))))))

(define (numer rn) (car rn))
(define (denom rn) (cdr rn))

(define t (make-rat 6 -18))
(numer t)
(denom t)
