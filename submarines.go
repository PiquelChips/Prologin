// EXERCISE LINK: https://prologin.org/train/2025/qualification/le_juste_semblable
package main

import "bufio"
import "os"
import "strconv"
import "strings"
import "fmt"

// n: le nombre de sous-marins
// m: le nombre de critères
// criteres: les critères de chaque sous-marin
func minimumExclus(n int, m int, criteres [][]int) {
    singles := [][]int{}
    for _, critere := range criteres {
        add_sub(critere, &singles)
    }
    fmt.Printf("%d", len(singles))
}

func add_sub(critere []int, singles *[][]int) {
    for i, single := range *singles {
        if critere_equal(critere, single) {
            (*singles)[i] = (*singles)[len(*singles) - 1]
            *singles = (*singles)[:len(*singles) - 1]
            return
        }
    }
    *singles = append(*singles, critere)
}

func critere_equal(critere []int, critere_ []int) bool {
    for index, i := range critere {
        if critere_[index] != i {
            return false
        }
    }
    return true
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var n int
    scanner.Scan()
    n, _ = strconv.Atoi(scanner.Text())
    var m int
    scanner.Scan()
    m, _ = strconv.Atoi(scanner.Text())
    criteres := make([][]int, n)
    for i := range criteres {
        criteres[i] = make([]int, m)
        scanner.Scan()
        for j, jValue := range strings.SplitN(scanner.Text(), " ", m) {
            criteres[i][j], _ = strconv.Atoi(jValue)
        }
    }
    minimumExclus(n, m, criteres);
}
