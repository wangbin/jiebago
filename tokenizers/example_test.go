package tokenizers_test

import (
	"fmt"

	"github.com/wangbin/jiebago/tokenizers"
)

func Example() {
	sentence := []byte("永和服装饰品有限公司")

	// default mode
	tokenizer, _ := tokenizers.NewJiebaTokenizer("../dict.txt", true, false)
	fmt.Println("Default Mode:")
	for _, token := range tokenizer.Tokenize(sentence) {
		fmt.Printf(
			"Term: %s Start: %d End: %d Position: %d Type: %d\n",
			token.Term, token.Start, token.End, token.Position, token.Type)
	}

	//search mode
	tokenizer, _ = tokenizers.NewJiebaTokenizer("../dict.txt", true, true)
	fmt.Println("Search Mode:")
	for _, token := range tokenizer.Tokenize(sentence) {
		fmt.Printf(
			"Term: %s Start: %d End: %d Position: %d Type: %d\n",
			token.Term, token.Start, token.End, token.Position, token.Type)
	}
	// Output:
	// Default Mode:
	// Term: 永和 Start: 0 End: 6 Position: 1 Type: 1
	// Term: 服装 Start: 6 End: 12 Position: 2 Type: 1
	// Term: 饰品 Start: 12 End: 18 Position: 3 Type: 1
	// Term: 有限公司 Start: 18 End: 30 Position: 4 Type: 1
	// Search Mode:
	// Term: 永和 Start: 0 End: 6 Position: 1 Type: 1
	// Term: 服装 Start: 6 End: 12 Position: 2 Type: 1
	// Term: 饰品 Start: 12 End: 18 Position: 3 Type: 1
	// Term: 有限 Start: 18 End: 24 Position: 4 Type: 1
	// Term: 公司 Start: 24 End: 30 Position: 5 Type: 1
	// Term: 有限公司 Start: 18 End: 30 Position: 6 Type: 1
}
