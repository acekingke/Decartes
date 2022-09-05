package interleave

/* Algorithm

(define (interleave-seq left right prefix )
  (cond ((null? left) (display (append prefix right)))
        ((null? right) (display (append prefix left)))
        (else (begin (interleave-seq (cdr left) right (append prefix (list (car left))))
                     (interleave-seq left (cdr right) (append prefix (list (car right)))))) ))


(interleave-seq '(x y z) '(a  b c) '())

*/
func Interleave(left []string, right []string, f func([]string)) {
	interleaveSeq(left, right, []string{}, f)
}

func interleaveSeq(left []string, right []string, prefix []string, f func([]string)) {
	if len(left) == 0 {
		f(append(prefix, right...))
	} else if len(right) == 0 {
		f(append(prefix, left...))
	} else {
		interleaveSeq(left[1:], right, append(prefix, left[0]), f)
		interleaveSeq(left, right[1:], append(prefix, right[0]), f)
	}
}
