# Alien Invasion
This is a program that reads in a world map, creates N aliens, and then put them in random cities. All the aliens can move randomly. If two aliens stay in the same city, they'll fight, kill each other and destroy that city. The program will run until all the aliens have been destroyed, or each alien has moved at least 10,000 times.

## How to Run

Type `./invasion` or `go run invasion.go` in command line to run the program

Then enter the path to the map you want to use and the alien numbers as required

```
$ ./invasion
What's the path for the map you want to use? maps/example.txt
How many aliens do you want to put in the world? 5
Qu-ux has been destroyed by alien-2 and alien-1
Bar has been destroyed by alien-4 and alien-0
The left cities and their connections are:
Foo west: Baz
Baz east: Foo
Bee
```

To run the test, just use `go test`
```
$ go test
city1 has been destroyed by alien-2 and alien-0
PASS
ok      _/Users/cy/alien-invasion/invasion      0.005s
```