;; Section 1.2.6 Example: Testing for Primality
;; Exercise 1.23

(define (square x)
  (* x x))

(define (smallest-divisor n)
  (define (divides? a b)
    (= (remainder b a) 0))

  (define (next test-divisor) 
    (if (= test-divisor 2) 
        3
        (+ test-divisor 2)))

  (define (find-divisor n test-divisor)
    (cond ((> (square test-divisor) n) n)
          ((divides? test-divisor n) test-divisor)
          (else (find-divisor n (next test-divisor)))))

  (find-divisor n 2))

(define (prime? n) 
  (= (smallest-divisor n) n))

(define (timed-prime-test n)
  (start-prime-test n (runtime)))

(define (start-prime-test n start-time)
  (if (prime? n)
      (report-prime n (- (runtime) start-time))
      #f))

(define (report-prime n elapsed-time)
  (display n)
  (display " is prime (")
  (display elapsed-time)
  (display " seconds)")
  (newline)
  #t)

(define (search-for-primes since until amount)
  (define start-time (runtime))
  
  (define (iter n until missing)
    (cond ((= missing 0) 
           (display "\nTotal time: ")
           (display (- (runtime) start-time))
           (display " seconds\n")
           'done)  ; Return a symbol instead of undefined
          ((= (remainder n 2) 0) 
           (iter (+ n 1) until missing))
          (else 
           (if (timed-prime-test n)
               (iter (+ n 2) until (- missing 1))
               (iter (+ n 2) until missing)))))

  (display "Searching for ")
  (display amount)
  (display " primes between ")
  (display since)
  (display " and ")
  (display until)
  (display ":\n\n")
  
  (iter since until amount))

; Example usage:
(search-for-primes 100000000000 10000000000000000 3)
