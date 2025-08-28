package product

import "fmt"

func (product *S_Product) validate() error {
	if product.Name == "" {
		return fmt.Errorf("TODO, melhorar essa mensagem de erro")
	}

	if product.Path == "" {
		return fmt.Errorf("TODO, melhorar essa mensagem de erro")
	}

	if product.Endpoints == nil {
		return fmt.Errorf("TODO, melhorar essa mensagem de erro")
	}

	return nil
}
