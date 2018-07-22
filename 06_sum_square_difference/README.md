<h3>Sum square difference</h3>
<h4>Problem 6</h4>
<p>The sum of the squares of the first ten natural numbers is,</br>
1^2 + 2^2 + ... + 10^2 = 385</br>
The square of the sum of the first ten natural numbers is,</br>
(1 + 2 + ... + 10)^2 = 552 = 3025</br>
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.</br>
Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.</p>

<h4>v.1.0</h4>
<p>Simple synchronous algorithm to calculate both and find the difference. Whole code was run 3 times with adjusting the amount of numbers.</p>
<pre>
100 numbers            - 9.91µs</br>
1,000,000,000 numbers  - 605.095761ms</br>
10,000,000,000 numbers - 5.957863311s</br>
</pre>

<h4>v.2.0</h4>
<p>The same algorithm but with <strong>go routines</strong> using waiting groups. Whole code was run 3 times with adjusting the amount of numbers.</p>
<pre>
100 numbers            - 12.549µs      - 1.26x <strong>slower</strong> than synchronous version</br>
1,000,000,000 numbers  - 1.625799847s  - 2.69x <strong>slower</strong> than synchronous version</br>
10,000,000,000 numbers - 16.325798972s - 2.74x <strong>slower</strong> than synchronous version</br>
</pre>

<h4>v.1.1</h4>
<p>Still synchronous but all calculations done in one run one after the other. First three - 100 numbers (each run takes less time although it's the same code), 4th - 1,000,000,000 numbers, 5th - 10,000,000,000 numbers.</p>
<pre>
100            - 8.948µs</br>
100.2          - 1.577µs</br>
100.3          - 1.356µs</br>
1,000,000,000  - 3.034843293s</br>
10,000,000,000 - 29.239634605s</br>
</pre>

<h4>v.2.1</h4>
<p>Still same go routines but all calculations done in one run one after the other. First three - 100 numbers (each run takes less time although it's the same code), 4th - 1,000,000,000 numbers, 5th - 10,000,000,000 numbers.</p>
<pre>
100            - 13.772µs      - 1.53x <strong>slower</strong> than synchronous version</br>
100.2          - 3.682µs       - 2.33x <strong>slower</strong> than synchronous version</br>
100.3          - 2.634µs       - 1.94x <strong>slower</strong> than synchronous version</br>
1,000,000,000  - 1.607745276s  - 1.88x <strong>faster</strong> than synchronous version</br>
10,000,000,000 - 16.006218614s - 1.82x <strong>faster</strong> than synchronous version</br>
</pre>