;; Section 2.2.2: Hierarchical Structures
;; Exercise 2.31

(load "../utils.scm") ;; square, map definitions

(define (tree-map tree proc)
  (map (lambda (sub-tree)
         (if (pair? sub-tree)
           (tree-map sub-tree proc)
           (proc sub-tree)))
       tree))

(define (square-tree tree) (tree-map tree square))

(square-tree (list 1
                   (list 2 (list 3 4) 5)
                   (list 6 7)))


