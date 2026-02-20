package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	/// escrita
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo..."))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)
	f.Close()

	/// leitura
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	/// leitura em lotes de arquivo grande
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}

/*
Aula de Manipulação de Arquivos em Go

Este código demonstra operações básicas e intermediárias de manipulação de arquivos utilizando a biblioteca padrão do Go.

Fluxo do programa:

1. Criação de arquivo
   - Usa os.Create para criar (ou sobrescrever) o arquivo "arquivo.txt".
   - Retorna um ponteiro para o arquivo e um erro, que deve sempre ser tratado.

2. Escrita no arquivo
   - Usa Write para gravar bytes dentro do arquivo.
   - Converte string para []byte pois Write trabalha com bytes.
   - Retorna a quantidade de bytes escritos.

3. Fechamento do arquivo
   - f.Close() libera recursos do sistema operacional.
   - Boa prática: sempre fechar arquivos após uso.

4. Leitura completa do arquivo
   - Usa os.ReadFile para ler todo o conteúdo de uma vez.
   - Ideal para arquivos pequenos.

5. Leitura em partes (buffer)
   - Usa os.Open + bufio.NewReader.
   - Lê o arquivo em pequenos blocos (buffer de 3 bytes neste exemplo).
   - Ideal para arquivos grandes, pois consome menos memória.

6. Remoção do arquivo
   - Usa os.Remove para deletar o arquivo do sistema.

Conceitos importantes:

- Sempre tratar erros retornados pelas funções.
- Arquivos devem ser fechados após uso.
- Para arquivos grandes, prefira leitura com buffer.
- Go trabalha fortemente com bytes em operações de I/O.

Este exemplo cobre o ciclo completo:
Criar -> Escrever -> Ler -> Ler em buffer -> Deletar
*/
