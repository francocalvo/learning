(define (square x)
  (* x x))

(define (smallest-divisor n)
  (define (divides? a b)
    (= (remainder b a) 0))
    
  (define (find-divisor n test-divisor)
    (cond ((> (square test-divisor) n) n)
          ((divides? test-divisor n) test-divisor)
          (else (find-divisor n (+ test-divisor 1)))))
          
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
           'done)
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

; Example searches at different ranges
(search-for-primes 100000000000 10000000000000000 3)

; Uncomment to test other ranges:
; (search-for-primes 10000 100000 3)
; (search-for-primes 100000 1000000 3)
; (search-for-primes 1000000 10000000 3)
