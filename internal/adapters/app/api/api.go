package api

import (
	"hex/internal/ports"
)

type Adapter struct {
	arith ports.ArithmeticPort
	db    ports.DbPort
}
type Operator struct {
	arith ports.Operator
}

func NewAdapter(db ports.DbPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{
		arith: arith,
		db:    db,
	}
}

func (apia Adapter) GetAddition(a, b int) (int, error) {
	answer, err := apia.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "addition")
	if err != nil {
		return 0, err
	}

	return answer, nil
}
func (apia Adapter) GetMultiplication(a, b int) (int, error) {
	answer, err := apia.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "multiplication")
	if err != nil {
		return 0, err
	}

	return answer, nil
}
func (apia Adapter) GetDivision(a, b int) (int, error) {
	answer, err := apia.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "division")
	if err != nil {
		return 0, err
	}

	return answer, nil
}
func (apia Adapter) GetSubtraction(a, b int) (int, error) {
	answer, err := apia.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "GetSubtraction")
	if err != nil {
		return 0, err
	}

	return answer, nil
}
