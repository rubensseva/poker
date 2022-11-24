package main

type card string

var (
	colors = []byte{'r', 's', 'h', 'k'}

	values = []byte{'2', '3', '4', '5', '6', '7', '8', '9', 't', 'r', 'd', 'k', 'a'}
	valueOfValue = map[byte]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		't': 10,
		'r': 11,
		'd': 12,
		'k': 13,
		'a': 14,
	}
	numEqualToString = map[int]string{
		2: "pair",
		3: "three of a kind",
		4: "four of a kind",
	}
)

func newStack() []card {
	cards := make([]card, 0, 52)
	for i := 0; i < 4; i++ {
		color := colors[i]
		for j := 0; j < 12; j++ {
			value := values[j]
			cards = append(cards, card([]byte{value, color}))
		}
	}

	return cards
}

func drawFive() []card {
	cards := newStack()
	shuffle(cards)
	return cards[:5]
}

func filterCardsByValue(val byte, hand []card) []card {
	equal := []card{}
	for _, c := range hand {
		if c[0] == val {
			equal = append(equal, c)
		}
	}
	return equal
}
