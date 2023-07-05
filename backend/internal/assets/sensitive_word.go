package assets

import (
	"embed"
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/strutil"
	"github.com/importcjj/sensitive"
	"github.com/ohanakogo/exceptiongo"
	"io/fs"
)

//go:embed tencent-sensitive-words
var tencentSensitiveWords embed.FS

func readTencentSensitiveWords() []string {
	file, err := tencentSensitiveWords.Open("tencent-sensitive-words/sensitive_words_lines.txt")
	exceptiongo.QuickThrow[types.IOException](err)

	defer func(file fs.File) {
		err := file.Close()
		exceptiongo.QuickThrow[types.IOException](err)
	}(file)

	readString := fsutil.ReadString(file)
	return strutil.Split(readString, fmt.Sprintf("\n"))
}

var SensitiveWordsFilter = func() *sensitive.Filter {
	filter := sensitive.New()
	for _, word := range readTencentSensitiveWords() {
		filter.AddWord(word)
	}
	return filter
}()
