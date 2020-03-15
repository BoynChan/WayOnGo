package JsonParse

import (
	"JsonParse/Json"
	"fmt"
	"testing"
)

// Author:Boyn
// Date:2020/3/15

func TestLexer(t *testing.T) {
	jS := GetString()
	buf = []rune(jS)
	ptr = 0
	length = len(buf)
	tokens := make([]Json.Token, 0)
	for {
		token, err := start()
		if err != nil && err == fmt.Errorf("EOF") {
			break
		} else if err != nil {
			break
		}
		tokens = append(tokens, token)
	}
	tokens = append(tokens, Json.Token{
		Value:     nil,
		TokenType: Json.EndDoc,
	})
	for _, v := range tokens {
		fmt.Println(fmt.Sprintf("%s ## %s", Json.Type2String(v.TokenType), v.Value))
	}
}
