# wc
Word, line or byte counter.

flags
  no flags = it will count words
  -l = it will count lines
  -b = it will count bytes


Example:

echo "This is an example command line tool for counting, in various ways, text bytes, lines or words" | ./wc 
17 words 

echo "This is an example command line tool for counting, in various ways, text bytes, lines or words" | ./wc -l
1 lines

echo "This is an example command line tool for counting, in various ways, text bytes, lines or words" | ./wc -b
95 bytes

