package main

import (
	"fmt"
	"sort"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("hello world")
	
	fmt.Println(arraySign([]int{2, 1}))//1	
	fmt.Println(arraySign([]int{-2, 1}) )// -1
	fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1}) )// 1

	
	fmt.Println(isAnagram("anak", "kana") )// true
	fmt.Println(isAnagram("anak", "mana"))// false
	fmt.Println(isAnagram("anagram", "managra") )// true | ini harusnya false 

	
	fmt.Printf("%c\n", findTheDifference("abcd", "abcde") )// 'e'
	fmt.Printf("%c\n", findTheDifference("abcd", "abced") )// 'e'
	fmt.Printf("%c\n", 	findTheDifference("", "y"))/// 'y'


	fmt.Println(	canMakeArithmeticProgression([]int{1, 5, 3}) )// true; 1, 3, 5 adalah baris aritmatik +2
	 
	fmt.Println( canMakeArithmeticProgression([]int{5, 1, 9}) )// true; 9, 5, 1 adalah baris aritmatik -4

	fmt.Println(canMakeArithmeticProgression([]int{1, 2, 4, 8}))// false; 1, 2, 4, 8 bukan baris aritmatik, melainkan geometrik x2

	tesDeck()
}

// https://leetcode.com/problems/sign-of-the-product-of-an-array
func arraySign(nums []int) int {
	// write code here
	sum := 1
	for _, num := range nums {
		sum *= num
	}
	if sum > 0 {
		return 1
	}
	return -1 
	// return 1 if positive
	// return -1 if negative
}

// https://leetcode.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	// write code here
	for i := range s {
		if s[i] != t[len(t) - i - 1] {
			return false
		}
	}
	return true
}

// https://leetcode.com/problems/find-the-difference
func findTheDifference(s string, t string) byte {
	var angka_s, angka_t int32

	for _, angka := range s {
		angka_s += angka
	}
	for _, angka := range t {
		angka_t += angka
	}

	return byte(angka_t - angka_s)
}


// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence
func canMakeArithmeticProgression(arr []int) bool {
	// write code here
	sort.Ints(arr)

	diff := arr[1] - arr[0]


	for i := 2; i < len(arr); i++ {
		if arr[i] - arr[i-1] != diff {
			return false
		}
	}

	return true
}

// Deck represent "standard" deck consist of 52 cards
type Deck struct {
	cards []Card
}

// Card represent a card in "standard" deck
type Card struct {
	symbol int // 0: spade, 1: heart, 2: club, 3: diamond
	number int // Ace: 1, Jack: 11, Queen: 12, King: 13
}

// New insert 52 cards into deck d, sorted by symbol & then number.
// [A Spade, 2 Spade,  ..., A Heart, 2 Heart, ..., J Diamond, Q Diamond, K Diamond ]
// assume Ace-Spade on top of deck.
func (d *Deck) New() {
	// write code here
	d.cards = make([]Card, 0, 52)
	for s := 0; s < 4; s++ {
		for n := 1; n <= 13; n++ {
			d.cards = append(d.cards, Card{symbol: s, number: n})
		}
	}
}

// PeekTop return n cards from the top
func (d Deck) PeekTop(n int) []Card {
	// write code here
	if n > len(d.cards) {
		n = len(d.cards)
	}
	return d.cards[:n]
}

// PeekTop return n cards from the bottom
func (d Deck) PeekBottom(n int) []Card {
	// write code here
	if n > len(d.cards){
		n = len(d.cards)
	}
	return d.cards[len(d.cards) - n:]
}

// PeekCardAtIndex return a card at specified index
func (d Deck) PeekCardAtIndex(idx int) Card {
	return d.cards[idx]
}

// Shuffle randomly shuffle the deck
func (d *Deck) Shuffle() {
	// write code here
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

// Cut perform single "Cut" technique. Move n top cards to bottom
// e.g. Deck: [1, 2, 3, 4, 5]. Cut(3) resulting Deck: [4, 5, 1, 2, 3]
func (d *Deck) Cut(n int) {
	// write code here
	d.cards = append(d.cards[n:], d.cards[:n]...)
}

func (c Card) ToString() string {
	textNum := ""
	switch c.number {
	case 1:
		textNum = "Ace"
	case 11:
		textNum = "Jack"
	case 12:
		textNum = "Queen"
	case 13:
		textNum = "King"
	default:
		textNum = fmt.Sprintf("%d", c.number)
	}
	texts := []string{"Spade", "Heart", "Club", "Diamond"}
	return fmt.Sprintf("%s %s", textNum, texts[c.symbol])
}

func tesDeck() {
	deck := Deck{}
	deck.New()

	top5Cards := deck.PeekTop(3)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}
	fmt.Println("---\n")

	fmt.Println(deck.PeekCardAtIndex(12).ToString()) // Queen Spade
	fmt.Println(deck.PeekCardAtIndex(13).ToString()) // King Spade
	fmt.Println(deck.PeekCardAtIndex(14).ToString()) // Ace Heart
	fmt.Println(deck.PeekCardAtIndex(15).ToString()) // 2 Heart
	fmt.Println("---\n")

	deck.Shuffle()
	top5Cards = deck.PeekTop(10)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}

	fmt.Println("---\n")
	deck.New()
	deck.Cut(5)
	bottomCards := deck.PeekBottom(10)
	for _, c := range bottomCards {
		fmt.Println(c.ToString())
	}
}
