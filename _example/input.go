// 基本的な標準入力 型:いろいろ
fmt.Scan

// 一行の標準入力 型:文字列
var sc = bufio.NewScanner(os.Stdin)
sc.Scan()
aString := sc.Text()

// 型確認
fmt.Println(reflect.TypeOf(t))