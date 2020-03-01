package digits

import "fmt"

// Digit type
type Digit [5]string

// Zero Digit
var Zero = Digit{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}

// One Digit
var One = Digit{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}

// Two Digit
var Two = Digit{
	"███",
	"  █",
	"███",
	"█  ",
	"███",
}

// Three Digit
var Three = Digit{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}

// Four Digit
var Four = Digit{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}

// Five Digit
var Five = Digit{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}

// Six Digit
var Six = Digit{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}

// Seven Digit
var Seven = Digit{
	"███",
	"  █",
	"  █",
	"  █",
	"  █",
}

// Eight Digit
var Eight = Digit{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}

// Nine Digit
var Nine = Digit{
	"███",
	"█ █",
	"███",
	"  █",
	"███",
}

// Colon Digit
var Colon = Digit{
	"   ",
	" ░ ",
	"   ",
	" ░ ",
	"   ",
}

// Digits has 1 to 9 and colon
var Digits = [...]Digit{
	Zero, One, Two, Three, Four, Five, Six, Seven, Eight, Nine,
}

//DisplayDigit prints Digit
func DisplayDigit(value Digit) {
	for _, val := range value {
		fmt.Printf("%s\n", val)
	}
}
