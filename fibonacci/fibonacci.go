package fibonacci

import (
	"math/big"
)

func CalculateMany(indexFibonacci int) (int, map[int]string) {
	var list = map[int]string{}
	var counterCalculateOne int
	var counterCalculateMany int
	for index := 1; index <= indexFibonacci; index++ {
		counterCalculateOne, list[index] = CalculateOne(index)
		counterCalculateMany++
	}
	return counterCalculateMany * counterCalculateOne, list
}

func CalculateOne(indexFibonacci int) (int, string) {
	fibPrevious, fibNext := big.NewInt(1), big.NewInt(1)
	counter := 0
	for index := 1; index <= indexFibonacci-2; index++ {
		fibPrevious, fibNext = fibNext, fibPrevious
		fibNext.Add(fibNext, fibPrevious)
		counter++
	}
	return counter, fibNext.String()
}

/*
Esta função (abaixo) em Go (Golang) calcula uma série de números de Fibonacci até um determinado índice usando canais (channels) para armazenar e manipular os valores da sequência. Vou explicar passo a passo o que está acontecendo no código:

func CalculateManyWithChannel(indexFibonacci int) map[int]string {: Esta é a declaração de uma função chamada CalculateManyWithChannel que recebe um argumento indexFibonacci do tipo inteiro e retorna um mapa onde as chaves são inteiros e os valores são strings.

channel := make(chan *big.Int, 2): Aqui, um canal chamado channel é criado para armazenar ponteiros para números grandes (*big.Int). Ele tem uma capacidade de 2, o que significa que pode armazenar até 2 valores sem bloquear a execução.

channel <- big.NewInt(0) e channel <- big.NewInt(1): Os primeiros dois valores da sequência de Fibonacci (0 e 1) são enviados para o canal.

var list = map[int]string{}: Um mapa vazio é criado para armazenar os valores calculados da sequência de Fibonacci, onde a chave é o índice e o valor é o número correspondente convertido para uma string.

for i := 1; i <= indexFibonacci; i++ {: Inicia um loop para calcular os números de Fibonacci até o índice fornecido.

fibPrevious := <-channel e fibNext := <-channel: Os dois últimos valores da sequência de Fibonacci (armazenados no canal) são retirados do canal e atribuídos às variáveis fibPrevious e fibNext.

fibNext.Add(fibNext, fibPrevious): O próximo número de Fibonacci é calculado somando os dois valores retirados anteriormente.

list[i] = fibNext.String(): O valor calculado é convertido para uma string e armazenado no mapa list com a chave i (o índice atual na sequência).

channel <- fibNext e channel <- fibPrevious: Os valores atualizados são enviados de volta para o canal para serem usados na próxima iteração do loop.

Finalmente, após o loop, o mapa list contendo os números de Fibonacci até o índice fornecido é retornado.

Esta função utiliza canais para gerenciar os valores da sequência de Fibonacci, calculando os números necessários para o índice especificado e os armazenando em um mapa para posterior utilização.
*/
func CalculateManyWithChannel(indexFibonacci int) (int, map[int]string) {
	channel := make(chan *big.Int, 2)
	channel <- big.NewInt(0)
	channel <- big.NewInt(1)
	var list = map[int]string{}
	counter := 0
	for i := 1; i <= indexFibonacci; i++ {

		fibPrevious := <-channel
		fibNext := <-channel

		fibNext.Add(fibNext, fibPrevious)

		list[i] = fibNext.String()

		channel <- fibNext
		channel <- fibPrevious

		counter++
	}
	return counter, list
}

/*
A diferença principal entre as duas implementações está na forma como os cálculos dos números de Fibonacci são realizados.

Na primeira implementação (utilizando canais), os cálculos são feitos de maneira mais "paralela". Ela utiliza canais para calcular dois valores de Fibonacci ao mesmo tempo, aproveitando o paralelismo oferecido pela manipulação concorrente dos canais. Isso significa que os cálculos são intercalados, com os valores sendo lidos e escritos nos canais ao longo do processo.

Por outro lado, na segunda implementação, os cálculos são sequenciais e são feitos de forma mais tradicional, usando uma função recursiva para calcular cada número de Fibonacci individualmente, começando do zero e indo até o índice desejado.

A primeira implementação pode ser mais rápida devido ao paralelismo introduzido pelo uso de canais. Enquanto um valor está sendo usado para calcular o próximo, a outra rotina já está preparando o valor seguinte. Isso pode reduzir o tempo de execução, especialmente para índices maiores, pois há uma menor dependência de cálculos sequenciais.

No entanto, é importante notar que o uso de canais e concorrência pode trazer complexidade adicional ao código, e o ganho de desempenho pode depender do contexto e do hardware em que o código é executado. Além disso, o custo de criação e gerenciamento dos canais pode afetar o desempenho em determinadas situações.

Em resumo, a primeira implementação aproveita o paralelismo oferecido pelo uso de canais, o que pode resultar em uma execução mais rápida, especialmente para índices grandes, enquanto a segunda implementação segue uma abordagem mais sequencial e tradicional para calcular os números de Fibonacci.
*/
