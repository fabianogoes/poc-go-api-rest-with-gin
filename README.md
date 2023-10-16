<img src="./assets/golang.png" height="90" width="300">

# PoC API Rest com Go e Gin

- [PoC API Rest com Go e Gin](#poc-api-rest-com-go-e-gin)
  - [Stack](#stack)
  - [Go](#go)
    - [Instalação do Go](#instalação-do-go)
    - [Get started with Go](#get-started-with-go)
    - [Referencias para estudo de Go](#referencias-para-estudo-de-go)
  - [Gin](#gin)
    - [Download and Install](#download-and-install)
    - [Documentação oficila do Gin](#documentação-oficila-do-gin)


**Objetivo:** Praticar como criar uma API Rest usando a linguagem [Go](https://go.dev/) com framework [Gin](https://github.com/gin-gonic/gin)

## Stack
- [Go](https://go.dev/)
- [Gin](https://gin-gonic.com/)
- [UUID](https://pkg.go.dev/github.com/google/uuid#section-readme)

## Go
> Crie sistemas simples, seguros e escaláveis ​​com Go
> - All releases: https://go.dev/dl/
> - Download and install: https://go.dev/doc/install

### Instalação do Go
1. Download
2. Remova qualquer instalação anterior do Go excluindo a pasta `/usr/local/go` (se existir) e extraia o arquivo que você acabou de baixar em /usr/local, criando uma nova árvore Go em /usr/local/go:
```shell
 rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz
```
3. Adicione `/usr/local/go/bin` à variável de ambiente PATH.
> Nota: As alterações feitas em um profile podem não ser aplicadas até a próxima vez que você fizer login no computador.    
> Para aplicar as alterações imediatamente, basta executar os comandos shell diretamente ou executá-los a partir do perfil usando um comando como `source $HOME/.profile`. 
4. Verifique se você instalou o **Go** abrindo um prompt de comando e digitando o seguinte comando:
```shell
 go version
```

### Get started with Go
- Visite o [tutorial de primeiros passos](https://go.dev/doc/tutorial/getting-started) para escrever um código Go simples. Demora cerca de 10 minutos para ser concluído.
- Veja também um tutorial sobre [Desenvolvendo uma API RESTful com Go e Gin](https://go.dev/doc/tutorial/web-service-gin)

### Referencias para estudo de Go
- [Golang by Example](https://golangbyexample.com/)
- [Go by example](https://gobyexample.com/)
- [The Go Programming Language](https://github.com/golang/go)
- [Aprenda Golang](https://aprendagolang.com.br/)
- [Golang by geeksforgeeks](https://www.geeksforgeeks.org/golang/)
- [Roadmap.sh Go Developer](https://roadmap.sh/golang)

## Gin

> O que é o **Gin**?
> **Gin** é um framework web escrito em **Golang**. 

**Rápido**    
- Roteamento baseado em árvore Radix
- pequeno espaço de memória. 
- Sem reflexão. 
- Desempenho previsível da API.

**Suporte de middleware**   
Uma solicitação HTTP recebida pode ser tratada por uma cadeia de middleware e pela ação final.   
Por exemplo: Logger, Autorização, GZIP e por fim postar uma mensagem no banco de dados.

**Livre de falhas**
**Gin** pode detectar um pânico ocorrido durante uma solicitação *HTTP* e recuperá-lo.   
Dessa forma, seu servidor estará sempre disponível.   
Também é possível reportar esse pânico ao Sentry, por exemplo!

**Validação JSON**
**Gin** pode analisar e validar o **JSON** de uma solicitação, verificando, por exemplo, a existência de valores obrigatórios.

**Agrupamento de rotas**
Organize melhor suas **rotas**.   
Autorização necessária versus não obrigatória, diferentes versões de API.   
Além disso, os grupos podem ser aninhados infinitamente sem prejudicar o desempenho.

**Gerenciamento de erros**
**Gin** fornece uma maneira conveniente de coletar todos os erros ocorridos durante uma solicitação *HTTP*.   Eventualmente, o *middleware* pode gravá-los em um arquivo de log, em um banco de dados e enviá-los pela rede.

**Renderização integrada**
**Gin** fornece uma API fácil de usar para renderização **JSON**, **XML** e **HTML**.

**Extensível**
Criar um novo middleware é muito fácil, basta verificar o código de exemplo.

### Download and Install
**Download and install**
```shell
go get -u github.com/gin\-gonic/gin
```

**Import**
```shell
import "github.com/gin-gonic/gin"
```

**(Opcional)** 
Importe net/http. Isto é necessário, por exemplo, se estiver usando constantes como http.StatusOK.
```shell
import "net/http"
```

**(opcional)**
`import github.com/google/uuid` or `go get github.com/google/uuid`


### [Documentação oficila do Gin](https://gin-gonic.com/docs/)
- [Introduction](https://gin-gonic.com/docs/introduction/)
- [Quickstart](https://gin-gonic.com/docs/quickstart/)
- [Deployment](https://gin-gonic.com/docs/deployment/)
- [Examples](https://gin-gonic.com/docs/examples/)
- [Testing](https://gin-gonic.com/docs/testing/)

Veja [livros sobre Golang](./Livros.md)

