;; Section 1.3.4: Procedures as General Methods
;; Exercise 1.44

(define dx 0.000001)

(define soothe
  (lambda (f)
    (lambda (x)
      (/ (+ (f (- x dx))
            (f x)
            (f (+ x dx)))
         3))))

;; I won't repeat the repeated procedure here, but it seems very straightforward.
