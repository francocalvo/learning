;; Section 1.3.3 Procedures as General Methods
;; Exercise 1.38: 

(define (cont-fract n d k)
  (define (cont-fract-inter n d k i)
    (if (= i k)
        0
        (/ (n i) (+ (d i) (cont-fract-inter n d k (+ i 1))))))
  (cont-fract-inter n d k 1)) ; start counting from 1


(define (n k) 1.0)
(define (d k)
  (if (= (remainder k 3) 2)
      (* 2 (/ (+ k 1) 3))
      1.0))

;;md The desired value is $\e - 2$ which is approximately 0.71828
(cont-fract n d 1000)
