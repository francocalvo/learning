;; Section 1.3.4: Procedures as General Methods
;; Exercise 1.43

(define (compose f g)
  (lambda (x) (f (g x))))

(define (square x) (* x x))
(define (inc x) (+ x 1))

((compose square inc) 6)

(define (repeated f n)
  (define (iter res m)
    (if (= m 0)
      res
      (iter (f res) (- m 1))))

  (lambda (x) (iter x n)))

((repeated square 2) 5)
(define (compose f g)
  (lambda (x) (f (g x))))

(define (square x) (* x x))
(define (inc x) (+ x 1))

((compose square inc) 6)
