import "strconv"

const (
	bin = 2
	dec = 10
	hex = 16
)

func convertBase(s string, fromBase, toBase int) (string, error) {
	n, err := strconv.ParseInt(s, fromBase, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(n, toBase), nil
}

func main() {
	var n int64 = 42
	fmt.Println(strconv.FormatInt(n, bin)) // 101010
	fmt.Println(strconv.FormatInt(n, hex)) // 2a
}