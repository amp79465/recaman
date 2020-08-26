# recaman
Playing with Recamán's sequence (https://oeis.org/A005132)

These programs explore efficiently computing Recamán's sequence in different languages.
Initially these are Python and Go.
The two python programs that are included are recaman.py and recaman3.py
These are the naive approach and the more sophisticated approach, respectively, the latter of which is described here.

The programs all currently rely on a lookup on the members of the sequence if the candidate is greater than zero, 
and by compressing the members of the sequence for that lookup and writing each member of the sequence as it is found 
to a file performance improved considerably. 
By using a list of lists containing two digits, a lower and an upper bound respectively, we can encode the members of
the sequence quite efficiently while losing the order of the sequence, and any terms that are repeated.
The first 70 terms of the sequence are:
	0, 1, 3, 6, 2, 7, 13, 20, 12, 21, 11, 22, 10, 23, 9, 24, 8, 25, 43, 62, 42, 63, 41, 18, 42, 17, 43, 16, 44, 15, 45, 14, 46, 79, 113, 78, 114, 77, 39, 78, 38, 79, 37, 80, 36, 81, 35, 82, 34, 83, 33, 84, 32, 85, 31, 86, 30, 87, 29, 88, 28, 89, 27, 90, 26, 91, 157, 224, 156, 225, 155

and using this method we can encode them into this structure:
[[0, 3], [6, 18], [20, 39], [41, 46], [62, 63], [77, 91], [113, 114], [155, 157], [224,225]]

This structure is much faster to search in the long term. While the naive approach of appending each addition member of the 
sequence onto a list grows linearly, this packing method grows much more slowly. At 100,000 terms the object has 427 sequences

The naive Go implentation is my clumsy attempt to make recaman.py happen in rnaive.go and it is much, much faster than both
python programs. Once this compressed data structure is implemented each lookup will be able to be run in parellel, and these
two combined will hopefully yield great enough performance to do deep searches of Recamán's sequence.

The OEIS website contains a page which lists the first 100,000 terms of the sequence and this has been used to test the results
of each of these programs.

A note on performance:
While rnaive.go outperformed recaman3.py significantly in smaller terms of the sequence, over 10 million terms recaman3.py
still wins out, completing it in 2.5 hours while rnaive.go took nearly 5 hours. This demonstrates the considerably higher
speed of Go, and correspondingly the utility of this packing function. 
