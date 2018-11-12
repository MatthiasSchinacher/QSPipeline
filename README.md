# QSPipeline
One of my first golang- projects, implementing a sort of "pipeline" for quicksort.

I programmed it mainly to check out the go- routines concurrent feature of golang.
There is no deeper or more profound goal I had with this.

# Implementation remarks
As said, this is one of my first "go" programs/ projects. I'm not really familiar with the language, this little
excercise was just to try it out.

The quicksort implemented is not a general one, but sorts only integers (int arrays).

TODO: describe the interface

# Running the "test"
The project has 2 source files, one with the actual quicksort, the other "test"- source a small main program,
which calls/ uses the quicksort. The go- language needs to be installed, then you can run this little gem by
"go run qspipeline.go qspipelinetest.go" on a command- line.

You should actually monitor the system while the script is running, to watch the go-routines using more than
one process/thread and thus benefitting from concurrency.  

Note: if the test is over to quickly, just edit "qspipelinetest.go" and sort a larger array.

Example:

    matthias@pop-os:~/wrk/QSPipeline$ go run qspipeline.go qspipelinetest.go
    Welcome to QSPipline :-) -  2018-11-12 20:25:46.972232694 +0100 CET m=+0.000201879
    sorting ...       -  2018-11-12 20:25:49.075663501 +0100 CET m=+2.103632713
    checking ...      -  2018-11-12 20:26:05.736101631 +0100 CET m=+18.764070829
    shutting down ... -  2018-11-12 20:26:05.790536082 +0100 CET m=+18.818505300
    done.             -  2018-11-12 20:26:05.790570246 +0100 CET m=+18.818539428
