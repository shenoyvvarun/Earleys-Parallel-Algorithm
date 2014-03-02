Earleys-Parallel-Algorithm
==========================
Parallelized version of earley's algorithm.

The algorithm exploits the cheap way to create goroutines to 
parallelize the earleys n^3 algorithm.

In this current release:
* The algorithm is slower because of the garbage collection kicks in every 60 second.(This will be fixed in further versions)
* This algorithm starts performing better for bigger and bigger Context sensitive / tree attributed grammar, which is a reasonable assumption to make.

