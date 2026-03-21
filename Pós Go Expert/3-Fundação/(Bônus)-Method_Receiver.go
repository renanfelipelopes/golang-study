/*
	🔹 O que é esse (c Cliente) antes da função?
	func (c Cliente) Desativar() {}

	Isso se chama:
	👉 Method Receiver (ou simplesmente receiver)

	🧠 Ideia central
	Em Go:
	- Funções viram métodos quando recebem um “dono” (receiver).

	Ou seja:
	- Desativar é uma função
	- (c Cliente) diz:
	👉 “essa função pertence ao tipo Cliente”

	🔧 O que isso permite?
	Graças ao receiver, você pode fazer:
		renan.Desativar()
	Sem o receiver, isso não seria possível.

	❌ Sem receiver → função comum
		func Desativar(c Cliente) {}

	Uso:
		Desativar(renan)
	👉 Isso não é método, é só função.
*/

/*

	🧩 Por que Go faz assim?
	Porque Go não tem classes.
	Então ele separa claramente:
	- Struct → dados
	- Função com receiver → comportamento

	Isso traz:
	- Menos acoplamento
	- Mais clareza
	- Código mais explícito

*/

/*

	⚠️ Valor vs Ponteiro (parte mais importante)
	Receiver por valor
		func (c Cliente) Desativar() {
			c.Ativo = false
		}

	👉 c é uma cópia
	👉 NÃO altera o objeto original

	Receiver por ponteiro (padrão de mercado)
		func (c *Cliente) Desativar() {
			c.Ativo = false
		}

	👉 Agora:
	- Altera o valor original
	- Evita cópia desnecessária
	- É o padrão quando há estado

	🧠 Regra mental simples
	Se o método muda o estado, use *Tipo

	✔ *Cliente → correto
	❌ Cliente → bug silencioso
*/

/*
	🧠 Resumo mental final
	- (c Cliente) é o receiver
	- Transforma função em método
	- Equivalente conceitual ao this
	- Liga comportamento ao tipo
	- Cliente → cópia
	- *Cliente → altera estado (padrão)
*/
