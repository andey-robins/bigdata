#!/bin/bash

# setup the execution environment
rm main
mv time0.out time0.out.bak
mv time1.out time1.out.bak
go build main.go

# time and calculate all of the 0 distance stuff 
echo "Beginning distance 0 calculations"
time ./main -in sentence_files/tiny.txt -size 100 >> time0.out
time ./main -in sentence_files/small.txt -size 100 >> time0.out
time ./main -in sentence_files/100.txt -size 100 >> time0.out
time ./main -in sentence_files/1K.txt -size 1000 >> time0.out
time ./main -in sentence_files/10K.txt -size 10000 >> time0.out
time ./main -in sentence_files/100K.txt -size 100000 >> time0.out
time ./main -in sentence_files/1M.txt -size 1000000 >> time0.out
time ./main -in sentence_files/5M.txt -size 5000000 >> time0.out
time ./main -in sentence_files/25M.txt -size 25000000 >> time0.out

echo "Done with distance 0 calculations, beginning distance 1 calculations"

# time and calculate all of the distance 1 stuff
time ./main -in sentence_files/tiny.txt -size 100 -k 1 >> time1.out
echo "Done with tiny.txt"
time ./main -in sentence_files/small.txt -size 100 -k 1 >> time1.out
echo "Done with small.txt"
time ./main -in sentence_files/100.txt -size 100 -k 1 >> time1.out
echo "Done with 100.txt"
time ./main -in sentence_files/1K.txt -size 1000 -k 1 >> time1.out
echo "Done with 1K.txt"
time ./main -in sentence_files/10K.txt -size 10000 -k 1 >> time1.out
echo "Done with 10K.txt"
time ./main -in sentence_files/100K.txt -size 100000 -k 1 >> time1.out
echo "Done with 100K.txt"
time ./main -in sentence_files/1M.txt -size 1000000 -k 1 >> time1.out
echo "Done with 1M.txt"
time ./main -in sentence_files/5M.txt -size 5000000 -k 1 >> time1.out
echo "Done with 5M.txt"
time ./main -in sentence_files/25M.txt -size 25000000 -k 1 >> time1.out
echo "Done with 25M.txt"