;; Section 2.2.2: Hierarchical Structures
;; Exercise 2.29

(define (make-mobile left right)
  (list left right))

(define (make-branch length structure)
  (list length structure))

(define (is-mobile? obj)
  (if (and
        (pair? obj)
        (pair? (left-branch obj))
        (pair? (right-branch obj)))
    #t 
    #f))

;;md Part a
(define (left-branch mobile)
  (car mobile))

(define (right-branch mobile)
  (cadr mobile))

(define (branch-length branch)
  (car branch))

(define (branch-structure branch)
  (cadr branch))

;;md Part b
(define (total-weight structure)
  (if (is-mobile? structure)
    (+  
      (total-weight (branch-structure (left-branch structure)))
      (total-weight (branch-structure (right-branch structure))))
    structure))))

;;md Part c
(define (is-balanced? structure)
  (if (is-mobile? structure)
    (if (and
        (is-balanced? (branch-structure (left-branch structure)))
        (is-balanced? (branch-structure (right-branch structure)))
        (= 
          (* 
            (branch-length (left-branch structure))
            (total-weight (branch-structure (left-branch structure))))
          (* 
            (branch-length (right-branch structure))
            (total-weight (branch-structure (right-branch structure))))))
      #t
      #f)
    #t))



;;md Example usage. This should output a weight of 26, and not be balanced.
(define m (make-mobile
  (make-branch 10 (make-mobile (make-branch 5 10) (make-branch 25 2)))
  (make-branch 10 (make-mobile
                    (make-branch 10 (make-mobile (make-branch 5 10) (make-branch 25 2))) 
                    (make-branch 25 2)))))

(display "Example 1")
(newline)
(display (is-balanced? m))
(newline)
(display (total-weight m))
(newline)

;;md This one should output a weight of 40, and be balanced.
(define balanced (make-mobile
  (make-branch 4 (make-mobile
                   (make-branch 6 (make-mobile (make-branch 6 4) (make-branch 4 6))) 
                   (make-branch 6 (make-mobile (make-branch 6 4) (make-branch 4 6)))))
  (make-branch 4 (make-mobile
                   (make-branch 6 (make-mobile (make-branch 6 4) (make-branch 4 6))) 
                   (make-branch 6 (make-mobile (make-branch 6 4) (make-branch 4 6)))))))

(display "Example 2")
(newline)
(display (is-balanced? balanced))
(newline)
(display (total-weight balanced))
(newline)


