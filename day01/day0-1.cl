#!/usr/bin/env clisp

; find two entries which sum to 2020, then multiply

(defun range (min max)
    (when (<= min max)
        (cons min (range (+ min 1) max))))

(set 'values (with-open-file (stream "data")
(loop for line = (read-line stream nil) 
    while line
    collect (parse-integer line))))

(loop 
    while values
    do  
        (let ((h (pop values)))
            (loop for i in values
                do
                (loop 
                    for j in values
                    do
                    
                    (if (= 2020 (+ h i j))
                        (progn
                            (print (* h i j)) 
                            (exit)))
                )

            )
        )
    )