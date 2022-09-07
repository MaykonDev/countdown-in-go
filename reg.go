package main
import "fmt"

func main() {
	fmt.Println("Contagem regressiva em Go")
	fmt.Println("Digite um numero para iniciar: ")
	var inicio int
	fmt.Scanln(&inicio)

	//geradores de delay

	var max int = 950
	var repeticao int = 3000
	var limite int = 750

	for i := inicio; i > 0; i = i - 1 {

		//inicio do gerador de Delay
		for j := 0; j < max; j = j + 1 { //ira repetir ate j ter valor maior de 950
			for k := 0; k < repeticao; k = k + 1 { // ira repetir ate k ter valor maior do que 3000
				for fim := max; fim > limite; fim = fim - 1 { // ira repetir ate fim ter valor menor do que 750
					var delay int = 0 // definindo variavel delay
					if delay < fim {
						delay = delay + 1 // ira repetir ate delay possuir valor maior que fim, ou seja, 950.
					}
				}
			}
		}

		// fim do gerador de delay
		fmt.Println("Tempo: ", i)
	}
}
