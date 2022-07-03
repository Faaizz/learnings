# Programming Foundations

- [Programming Foundations](#programming-foundations)
  - [Algorithmic Design](#algorithmic-design)
    - [Basic Steps](#basic-steps)
  - [References](#references)

## Algorithmic Design
### Basic Steps
1. **Work out multiple instance yourself**: Work out multiple instances of the problem by hand. 
Be sure to make each step as clear & concise as possible. 
Some steps might be complex and might require being solved with an algorithm themselves.
If you are having problems with this step, you might lack domain knowledge, or the problem might be ill-specified.

2. **Write down your steps from 1**: Write down the steps that we used to solve the problem from 1 above.
For example, to compute *x* to the power of *y*, for the "instance" where *x = 3* and *y = 4*, we might list the following steps:
```
Multiply 3 by 3
    You get 9
Multiply 3 by 9
    You get 27
Multiply 3 by 27
    You get 81
81 is your answer.
```

3. **Generalize your steps from 2**: Compare and contrast the solutions from 2 and try to decern them into a generalized solution.
Extract values into variables and repetitions into loops.
Following the same example of *3* to the power pf *4* from Step 2 above, we can have:
```
Multiply x by 3
    You get 9
Multiply x by 9
    You get 27
Multiply x by 27
    You get 81
81 is your answer.
```

Then:
```
Start with n = 1
n = Multiply x by n
n = Multiply x by n
n = Multiply x by n
n = Multiply x by n
n is your answer
```

Then:
```
Start with n = 1
Count up from 1 to y (inclusive), for each count,
    n =  Multiply x by n
n is your answer
```

4. **Test your generalized solution**: Test your generalized solution, first, on all the cases for which Step 1 was carried out.
Also test corner cases for the variables contained in the algorithm.
If your generalized solution fails for a test case, perform steps 1 & 2 on the particular failed test and incorporate these into your generalization in step 3.


## References
- [Programming Fundamentals - Coursera](https://www.coursera.org/learn/programming-fundamentals)
