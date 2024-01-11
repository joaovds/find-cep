package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
  for _, cep := range os.Args[1:] {
    err := validation(cep)
    if err != nil {
      fmt.Println(err)
      continue
    }

    err = getCepInfo(cep)
    if err != nil {
      fmt.Println(err)
    }
  }
}

func validation(cep string) error {
  if len(cep) != 8 {
    return fmt.Errorf("CEP inválido: %s", cep)
  }

  return nil
}

func getCepInfo(cep string) (error) {
  req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
  if err != nil {
    return fmt.Errorf("Erro ao buscar CEP: %s", err)
  }
  defer req.Body.Close()

  res, err := io.ReadAll(req.Body)
  if err != nil {
    return fmt.Errorf("Erro ao ler resposta: %s", err)
  }

  fmt.Println("CEP:", cep)
  fmt.Println("Resposta:", string(res))
  fmt.Println("")

  return nil
}
