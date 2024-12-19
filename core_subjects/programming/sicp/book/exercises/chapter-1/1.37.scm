;; Section 1.3.3 Procedures as General Methods
;; Exercise 1.37: 

(define (cont-fract n d k)
  (if (= k 0)
      0
      (/ (n k)  (+ (d k) (cont-fract n d (- k 1))))))


(define (n k) 1.0)

;;md The desired value is: 0.6180
;;md It can be achieved by calling with k = 11
(cont-fract n n 11) ; 
(cont-fract n n 11)
