package game

// Could use pointers but let's keep the uint8 thing in level defs
var Messages = [][]string{
	// 0 - Empty default
	{},
	//  "                                        ",
	// 1
	{
		"Work day completed.",
	},
	// 2
	{
		"Ouch! Let's try that again!",
	},
	// 3
	{
		"Careful! You're on thin ice!",
	},
	// 4
	{
		"You're gonna get ME fired!",
	},
	// 5
	{
		"THAT'S IT! GET OUTTA HERE!",
	},
	// 6
	{
		"Day 0 - Training",
		"Press SPACE anytime to Skip",
	},
	// 7
	{
		"This is a demolition site.",
		"Look out for falling objects!",
	},
	// 8
	{
		"You're the Foreman, and you gotta tell",
		"the workers where to move to stay safe.",
	},
	// 9
	{
		"Keys A and D tell workers to move",
		"'LEFT' and 'RIGHT'",
	},
	// 10
	{
		"Keys J, K, L tell single workers to",
		"'STAY PUT!'",
	},
	// 11
	{
		"Workers on the same beam can't pass",
		"each other",
	},
	// 12
	{
		"Sandwiches are a tasty treat worth $50.",
		"Grab them as they fall.",
	},
	// 13
	{
		"I know this is dangerous work.",
		"Cash wads hold $100 if you're brave.",
		"I hope it's worth the risk!",
	},
	// 14
	{
		"Buckets will give em' a nasty bump.",
		"They won't be able to move for a moment",
		"after being hit.",
	},
	// 15
	{
		"If a bucket hits someone while stunned,",
		"it's lights out for that one!",
	},
	// 16
	{
		"Beams are the most dangerous,",
		"all it takes is a single hit and KAPUT",
	},
	// 17
	{
		"If your guys are in trouble",
		"press SPACE to use a bomb!",
		"Bombs will clear all falling objects.",
	},
	// 18
	{
		"Grab one when you see them, and save it",
		"for the next sticky situation.",
	},
	// 19
	{
		"I think you're about ready for your",
		"first shift!",
	},
	// 20
	{
		"Good Luck!",
	},
	// 21
	{
		"Day 1 - Monday",
		"Let's start with one worker",
	},
	// 22
	{
		"A Game by Justin Horton",
		"Made for EbitenJam 2025",
		"Music by Jake Schofield",
		"   Press J to start!   ",
	},
	// 23
	{
		"Woah! That was rough.",
		"You're doing great, kid.",
	},
	// 24
	{
		"Look, I know it's hard.",
		"It's an arcade game!",
	},
	// 25
	{
		"Day 2 - Tuesday",
		"You're ready for a pair now.",
	},
	// 26
	{
		"Day 3 - Wednesday",
		"It's a lot to keep straight.",
	},
	// 27
	{
		"Maybe there is a better way",
		"to do this?",
	},
	// 28
	{
		"Day 4 - Thursday",
		"I can't wait for this week to",
		"be over!",
	},
	// 29
	{
		"Day 5 - Friday",
		"The longest day.",
	},
	// 30
	{
		"I have good news and",
		"I have bad news...",
	},
	// 31
	{
		"Good news is:",
		"One of the crew called in sick tomorrow.",
		"Easier to manage!",
	},
	// 32
	{
		"The bad news is:",
		"Ya'll gotta work over-time.",
	},
	// 33
	{
		"Actually, ain't no telling how much...",
	},
	// 34
	{
		"Well, I've got half day planned.",
		"Good luck, and lock up on your",
		"way out!",
	},
	// 35
	{
		"What the hell?!",
		"How'd you make it this far?",
		"Really, congrats!",
		"Unfortunately this can't go",
		"on forever...",
	},
}
