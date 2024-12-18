(define (sum term a next b)
  (if (> a b) 
    0 
    (+ (term a) 
       (sum term (next a) next b))))

;; Simpson's Rule integration implementation
(define (integral f a b n)
  ;; Define h
  (let ((h (/ (- b a) n)))

  ;; Wrap the term in a lambda to pass the h and a values
  (define (fs f h a) 
    (lambda (x) 
      (*
        ;; I guess the diference between this and the other implementation is that
        ;; I don't create the extra function to calculate the coefficient.
        (cond ((or (= x 0) (= x n)) 1)
              ((odd? x) 4)
              (else 2))
        (f (+ a (* x h))))))


  ;; Increment function
  (define (next a) (+ a 1))

  ;; Calculate the integral
  (*
    (/ h 3)
    (sum (fs f h a) 0 next n))))


;; Claude says this is more idiomatic.
;; ; Sum procedure - iteratively applies term from a to b
;; (define (sum term a next b)
;;   (if (> a b)
;;       0
;;       (+ (term a)
;;          (sum term (next a) next b))))
;;
;; ; Simpson's Rule integration implementation
;; (define (integral f a b n)
;;   ; Ensure n is even
;;   (if (odd? n)
;;       (error "n must be even")
;;       
;;       (let ((h (/ (- b a) n)))
;;         ; Create coefficient function for Simpson's Rule
;;         (define (simpson-coefficient k)
;;           (cond ((or (= k 0) (= k n)) 1)
;;                 ((odd? k) 4)
;;                 (else 2)))
;;         
;;         ; Function to evaluate at each point
;;         (define (term k)
;;           (* (simpson-coefficient k)
;;              (f (+ a (* k h)))))
;;         
;;         ; Simple increment function
;;         (define (next x) (+ x 1))
;;         
;;         ; Calculate integral using Simpson's Rule
;;         (* (/ h 3)
;;            (sum term 0 next n)))))

; Test function
(define (cube x) (* x x x))

; Convert to decimal and compute integral
(exact->inexact (integral cube 0 1 100000))

(define (cube x) (* x x x))

(integral cube 0 1 100)

