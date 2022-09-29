// Overlay complex Hello world script.
((hello, world, space = " ") => {
	const message = [
		// First Word
		hello,
		// Second Word
		world
	].join(space);
	console.log(message);
})("Hello", "World!");
