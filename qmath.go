package qmath

import (
	"time"
	"math/rand"
)

type Quote struct {
	Name, Summary, Author string
}

var Quotes []Quote = []Quote{
	{
		Name: "Natural numbers",
		Summary: "Natural numbers are the simple counting numbers (0, 1, 2, 3, ...). The skill of counting is intimatelty linked to the development of complex societies through trade, technology and documentation. Counting requires more than numbers, though. It involves addition, and hence subtraction too.",
		Author: "Paul Glendinning",
	},
	{
		Name: "One",
		Summary: "Together with zero, the number one is at the heart of all arithmetic.",
		Author: "Paul Glendinning",
	},
	{
		Name: "Zero",
		Summary: "Zero is complex idea, and for a long time there was considerable philosophical reluctance to recognize and put a name to it.",
		Author: "Paul Glendinning",
	},
	{
		Name: "Infinity",
		Summary: "Infinity (represented mathimatically as ∞) is simply the concept of endlessness: an infinite object is one that is unbounded.",
		Author: "Paul Glendinning",
	},
	{
		Name: "Number systems",
		Summary: "A number system is a way writing down numbers. In our everyday decimal system, we represent number in the from 434.15, for example. Digits within the number indicate units, tens, hundreds, tenths, hundredths, thousandths and so on, and are called coefficients.",
		Author: "Paul Glendinning",
	},
	{
		Name: "Rational numbers",
		Summary: "Rational numbers are numbers that can be expressed by dividing one integer by another non-zero integer. Thus all rational numbers take the form of fractions or quotients. These are written as one number, the numerator, divided by a second, the denominator.",
		Author: "Paul Glendinning",
	},
	{
		Name: "Squares",
		Summary: "The square of any number x is the product of the number times itself, denoted x².",
		Author: "Paul Glendinning",
	},
	{
		Name: "Prime numbers",
		Summary: "Prime numbers are positive integers that are divisible only by themselves and 1. The first eleven are 2, 3, 5, 7, 11, 13, 17, 19, 23, 29 and 31, but there are infinitely many. By convention, 1 is not considered prime, while 2 is the only even prime. A number that is neither 1 nor a prime is called a composite number.",
		Author: "Paul Glendinning",
	},
	{
		Name: "Divisors and remainders",
		Summary: "A number is a divisor of another number if it divides into that number exactly, with no remainder. So 4 is a divisor of 12, because it can be divided into 12 exactly three times. In this kind of operation, the number being divided, 12, is known as dividend.",
		Author: "Paul Glendinning",
	},
	{
		Name: "Irrational numbers",
		Summary: "Irrational numbers are numbers that cannot be expressed by dividing one natural number by another. Unlike rational numbers, they cannot be expressed as a ratio between to integers, or in a decimal form that either comes to an end or lapses into a regular pattern of repeating digits. Instead, the decimal expansions of irrational numbers carry on forever without periodic repetition.",
		Author: "Paul Glendinning",
	},
}

func Rand() Quote {
	rand.Seed(time.Now().UnixNano())
	return Quotes[rand.Intn(len(Quotes))]
}
