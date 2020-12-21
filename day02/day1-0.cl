#!/usr/bin/cl -Q

; Return number valid.
; data looks like:
;  ```1-3 a: abcde
;     1-3 b: cdefg
;     2-9 c: ccccccccc
; ```
; 2 are valid, n-m c is n to m occurrences of c must be in string for it to be valid

(ql:quickload :cl-ppcre)

(defpackage :day01)

(defun valid (check_line)
    (let 
        ((match 
            (ppcre:scan "123" "abcd123456")))))

;(defun valid (check_line)
;    (let 
;        ((match 
;            (ppcre:scan "(\\d+)-(\\d+)\\s+(\\w):\\s+(\\w+)" check_line))
;
;        (progn
;            (print match)
;            (loop 
;                for m in (regexp:match
;                    "\\([0-9]\\+\\)-\\([0-9]\\+\\)\\s\\+\\(\\w\\):\\s*\\(\\w\\+\\)$" ; \\(\\w\\+\\)\\s*\\$" ; [ \t]\\+\\(.\\): \\(.+\\)[ \t]*$"
;                    check_line)
;                do (print m)
;            )
;            (print check_line)
;            ))))


(print
    (length
        (with-open-file (stream "data")
            (loop for check_line = (read-line stream nil) 
                ; todo check if while line makes it so that it ends prematurely with an
                ; empty line
                while check_line  ; while line gets me no error on empty line
                if (valid check_line)
                collect T))))
