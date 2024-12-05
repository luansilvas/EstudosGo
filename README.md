# Gerenciador de Temperatura e Autorizador de Débito
- Autorizador de Débito (Questão 1): Gerencia transações de clientes respeitando limites financeiros e temporais.
- Gerenciador de Temperatura (Questão 2): Coordena processos concorrentes (observador e tratadores) usando notificações.

## Estrutura do Repositório
````
├── gerenciador-de-temperatura/       # Código da Questão 2
│   ├── cmd/gerenciador/
│   │   └── main.go                   # Ponto de entrada do Gerenciador de Temperatura
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go             # Absorve dados do arquivo de variáveis de ambiente
│   │   ├── generator/
│   │   │   └── observador.go         # Lógica do Observador
│   │   ├── logger/
│   │   │   └── logger.go             # Logger configurado
│   │   ├── processor/
│   │   │   └── tratador.go           # Lógica dos Tratadores
│   ├── go.mod
│   ├── go.sum
│   ├── README.md
│
├── transaction-five-minutes/         # Código da Questão 1
│   ├── cmd/main/
│   │   └── main.go                   # Ponto de entrada do Autorizador de Débito
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go             # Absorve dados do arquivo de variáveis de ambiente
│   │   ├── transactions/
│   │   │   ├── manager.go            # Lógica do Gerenciamento de Transações
│   │   │   └── transaction.go        # Modelo de Transação
│   │   ├── ui/
│   │   │   └── ui.go                 # Interface para interações do usuário
│   ├── go.mod
│   ├── go.sum
│   ├── .gitignore
│   ├── README.md
````

## Questão 1: Autorizador de Débito
### Descrição
O Autorizador de Débito gerencia transações de clientes com as seguintes regras:

- Cada cliente possui um histórico de transações.
- Não é permitido que um cliente transacione mais de R$ 1000 em um período de 5 minutos.

Dados da Transação:
- idCliente
- dataHora
- valor

### Armazenamento
**sync.Map**: Segundo a [documentação ](https://pkg.go.dev/sync#Map) permite trabalhar de maneira segura com os acessos de leitura e escrita no map, sobretudo quando trata-se de keys separadas.
Para saber mais sobre a estrutrua do sync.map clique [aqui](https://reliasoftware.com/blog/go-sync-map).

### Estrutura de Arquivos
   - manager.go: Gerencia a validação e armazenamento de transações.
   - transaction.go: Define a estrutura de uma transação.
   - ui.go: Gerencia a interação com o usuário.

## Questão 2: Gerenciador de Temperatura
### Descrição

O Gerenciador de Temperatura possui três processos concorrentes:

#### Observador:
- Gera valores de temperatura entre 0 e 100.
  - Notifica os Tratadores com base nos critérios:
    - Temperatura ≤ 40: Notifica Tratador 1.
    - Temperatura > 50: Notifica Tratador 2.

### Tratadores 1 e 2:
- Cada tratador processa notificações:
  - Após 10 notificações, imprime os valores processados.
  - Sincroniza com o Observador usando [sync.Cond](https://pkg.go.dev/sync#Cond).

### Semáforo
Para implementação de um semafóro, cada goroutine possui seu Cond, que partilham do mesmo Mutex. Tais goroutines ficam com o lock `liberado`, à espera de um sinal é emitido pelo observador que também detém os Conds

Para mais informações, clique [aqui](https://hackernoon.com/lang/pt/entendendo-o-synccond-em-um-guia-para-iniciantes).

### Observador:
Usa Signal para notificar o Tratador correspondente.
### Tratadores:
Usam Wait para aguardar notificações do Observador.

### Estrutura de Arquivos
- **observador.go:** Gera as temperaturas e decide qual tratador notificar.
- **tratador.go:** Processa as notificações.