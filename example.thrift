struct flop {
	1: i16 a,
	2: i16 b,
}

exception ExampleException {}

service Ping {
	# Simply returns
	void ping()
}

service Service extends Ping {
	# Returns the number of times this server has seen requests to count
	i16 count()

	# Returns str
	string echo(1: string str)

	# Returns a new flop, with flop.a and flop.b swapped
	flop flip(1: flop orig)

	# Throws an ExampleException
	void fail() throws (1: ExampleException e)
}
