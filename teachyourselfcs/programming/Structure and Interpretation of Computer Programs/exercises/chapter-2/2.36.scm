;; Section 2.2.3: Sequences as Conventional Interfaces
;; Exercise 2.36
(load "../utils.scm") ;; accumulate definition


;;md Ok. So first we'll see the first solution I had.It's way too verbose but
;;md it works perfectly.

(define (select-n n seq)
  (if (= n 0)
    (car seq)
    (select-n (- n 1) (cdr seq))))

(define (enumerate-n n seq)
  (if (null? seq)
    '()
    (cons (select-n n (car seq)) (enumerate-n n (cdr seq)))))

(define (pop n seq)
  (cond 
    ((= n 0) seq)
    ((null? seq) '())
    (else (pop (- n 1) (cdr seq)))))

(define (pop-n seqs)
  (if (null? seqs) 
    '()
    (cons 
      (pop 1 (car seqs))
      (pop-n (cdr seqs)))))


(define (accumulate-n op init seqs)
  (if (null? (car seqs))
    '() 
    (cons (accumulate op init (enumerate-n 0 seqs))
          (accumulate-n op init (pop-n seqs)))))

;;md Later, I came across this solution... It does the exact same thing, but 
;;md without creating a complex pair of functions to express the idea.
(define (accumulate-n op init seqs)
  (if (null? (car seqs))
    '() 
    (cons (accumulate op init (map car seqs))
          (accumulate-n op init (map cdr seqs)))))

(accumulate-n + 0 (list (list 1 2 3) (list 4 5 6) (list 7 8 9) (list 10 11 12)))
