# Utify

![Go Version](https://img.shields.io/github/go-mod/go-version/jonatas-sas/utify)
![License](https://img.shields.io/github/license/jonatas-sas/utify)
![Tests](https://github.com/jonatas-sas/utify/actions/workflows/ci.yml/badge.svg)
![Code Coverage](https://img.shields.io/codecov/c/github/jonatas-sas/utify)
![Stars](https://img.shields.io/github/stars/jonatas-sas/utify?style=social)

Utify é uma biblioteca Go para exibir mensagens estilizadas no terminal com suporte a cores, ícones e formatação avançada.

## Instalação

Para instalar o `utify`, utilize o comando:

```sh
go get github.com/jonatas-sas/utify
```

## Uso

Exemplo básico de uso:

```go
package main

import (
	"github.com/jonatas-sas/utify"
)

func main() {
	utify.Success("Operação concluída com sucesso!", utify.Options{})
	utify.Error("Ocorreu um erro inesperado.", utify.Options{})
	utify.Warning("Isso pode causar problemas.", utify.Options{})
	utify.Info("Informação útil.", utify.Options{})
	utify.Debug("Depuração ativada.", utify.Options{})
	utify.Critical("Erro crítico!", utify.Options{})
}
```

## Opções de Estilização

O `utify` permite configurar mensagens com opções avançadas:

| Opção     | Descrição                       |
| --------- | ------------------------------- |
| `Bold`    | Exibe o texto em negrito        |
| `Italic`  | Exibe o texto em itálico        |
| `NoColor` | Remove cores da saída           |
| `NoIcon`  | Remove ícones das mensagens     |
| `Exit`    | Finaliza a execução após exibir |

Exemplo de uso com opções:

```go
utify.Success("Mensagem em negrito", utify.Options{Bold: true})
utify.Error("Mensagem sem cor", utify.Options{NoColor: true})
```

## Métodos Disponíveis

O `utify` fornece os seguintes métodos para exibir mensagens estilizadas:

### **Mensagens de Status Geral**

- `Success(text string, opts Options)` → Mensagem de sucesso
- `Error(text string, opts Options)` → Mensagem de erro
- `Warning(text string, opts Options)` → Aviso ao usuário
- `Info(text string, opts Options)` → Informação útil
- `Debug(text string, opts Options)` → Mensagem de depuração
- `Critical(text string, opts Options)` → Erro crítico

### **Ações Comuns**

- `Delete(text string, opts Options)` → Remoção de itens
- `Update(text string, opts Options)` → Atualização de dados
- `Install(text string, opts Options)` → Instalação de pacotes
- `Upgrade(text string, opts Options)` → Atualização de versões
- `Edit(text string, opts Options)` → Modificação de itens
- `New(text string, opts Options)` → Criação de novos itens

### **Operações Específicas**

- `Download(text string, opts Options)` → Indica um processo de download
- `Upload(text string, opts Options)` → Indica um processo de upload
- `Sync(text string, opts Options)` → Indica sincronização de dados
- `Search(text string, opts Options)` → Indica busca ou pesquisa

## Uso com `Echo`

Todas as funções acima utilizam internamente o método `Echo`. Caso precise personalizar a exibição, pode chamá-lo diretamente:

```go
utify.Echo(utify.MessageSuccess, "Mensagem customizada", utify.Options{Bold: true, NoIcon: true})
```

## Testes

Para rodar os testes, utilize o comando:

```sh
go test -v
```

Os testes verificam a saída das mensagens e opções, garantindo que a formatação funcione corretamente.

## Contribuição

Contribuições são bem-vindas! Para sugerir melhorias, abra uma issue ou envie um pull request.

## Licença

Este projeto está licenciado sob a licença MIT.
