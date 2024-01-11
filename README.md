## Buscar CEP

###Executar

Exemplo de requisição

```bash
# go run cmd/main.go [...ceps]
#exemplo:
go run cmd/main.go 11015002 11015003
```

Resposta:

```
CEP: 11015002
Resposta: {
  "cep": "11015-002",
  "logradouro": "Avenida Conselheiro Nébias",
  "complemento": "de 202 a 376 - lado par",
  "bairro": "Vila Mathias",
  "localidade": "Santos",
  "uf": "SP",
  "ibge": "3548500",
  "gia": "6336",
  "ddd": "13",
  "siafi": "7071"
}

CEP: 11015003
Resposta: {
  "cep": "11015-003",
  "logradouro": "Avenida Conselheiro Nébias",
  "complemento": "de 203 a 381 - lado ímpar",
  "bairro": "Vila Mathias",
  "localidade": "Santos",
  "uf": "SP",
  "ibge": "3548500",
  "gia": "6336",
  "ddd": "13",
  "siafi": "7071"
}
```
