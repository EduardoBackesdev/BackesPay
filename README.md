# BackesPay - Sistema de Pagamentos com Go

BackesPay é um sistema de pagamentos desenvolvido em Go que demonstra o uso avançado de goroutines e WebSocket para processamento assíncrono de transações e comunicação em tempo real.

## 🚀 Características Principais

### Goroutines para Processamento de Transações
O sistema utiliza goroutines para processar pagamentos BED (TED) de forma assíncrona, permitindo:
- Execução de transações em horário comercial (8h às 15h, dias úteis)
- Processamento não-bloqueante de transações
- Verificação contínua do horário para execução
- Notificação em tempo real do status da transação

```go
func Routine_bed(data repositories.BedRequest) {
    for {
        // Verifica horário comercial e processa a transação
        // Notifica o cliente via WebSocket
    }
}
```

### WebSocket para Comunicação em Tempo Real
Implementação de WebSocket para:
- Conexão persistente com o cliente
- Notificações instantâneas sobre o status das transações
- Gerenciamento de múltiplas conexões com mutex para thread safety
- Pool de conexões para controle eficiente de recursos

```go
var Map_clients = make(map[int]*websocket.Conn)
var Mutex = &sync.Mutex{}
```

## 🛠 Tecnologias Utilizadas

- **Go** - Linguagem principal
- **Gin** - Framework web
- **MySQL** - Banco de dados
- **Gorilla WebSocket** - Implementação de WebSocket
- **JWT** - Autenticação
- **Decimal** - Precisão em operações monetárias

## 📊 Estrutura do Projeto

- `cmd/app`: Ponto de entrada da aplicação
- `internal/api`: Handlers e serviços da API
- `internal/repositories`: Camada de acesso ao banco de dados
- `internal/ws`: Gerenciamento de conexões WebSocket
- `lib/goroutines`: Implementação de rotinas assíncronas
- `routes`: Configuração de rotas da API

## 🔐 Funcionalidades

- Criação e autenticação de contas
- Transferências via BIX (instantâneas)
- Transferências via BED (agendadas/horário comercial)
- Consulta de saldo
- Verificação de destinatário
- Notificações em tempo real de transações

## 🔄 Sistema de Transações

### BIX (Transferência Instantânea)
- Processamento síncrono
- Validação imediata
- Atualização instantânea de saldo

### BED (Transferência Agendada)
- Processamento assíncrono via goroutines
- Execução apenas em horário comercial
- Notificação em tempo real via WebSocket
- Thread-safe com uso de mutex

## 🚧 Próximas Implementações

- **Histórico de Transações**: Sistema para rastreamento e consulta de todas as operações realizadas
  - Integração de uma fila de notificações com RabbitMQ

## 💡 Objetivo do Projeto

Este projeto foi desenvolvido com o propósito de mostrar minhas habilidades em:
- Programação concorrente com goroutines
- Comunicação em tempo real com WebSocket

---

⌨️ Desenvolvido por EduardoBackesdev
