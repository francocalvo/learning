;; Section 1.2.2 Tree Recursion;
;; Exercise 1.11

(define (f_r n)
  (cond ((< n 3) n)
        ((>= n 3)
         (+ 
           (f_r (- n 1))
           (* 2 (f_r (- n 2)))
           (* 3 (f_r (- n 3)))
           ))))

(define (f_i n)
  (define (f_i_iter a b c count)
      (cond 
        ((= count 0) (+ a (* 2 b) (* 3 c))) 
        ((> count 0) (f_i_iter (+ a (* 2 b) (* 3 c)) a b (- count 1)))))

  (cond ((< n 3) n)
        ((>= n 3)
          (f_i_iter 2 1 0 (- n 3)))))
                                 

(f_r 2)
(f_r 3)

(f_i 2)
(f_i 3)
