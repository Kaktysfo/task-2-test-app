package events

import "errors"

type Priority string

var (
	ErrInvalidPriority = errors.New("Неверный приоритет!\nФормат приоритета: \"Высокий\",\"Средний\",\"Низкий\"")
)

const (
	PriorityHigh   Priority = "Высокий"
	PriorityMedium Priority = "Средний"
	PriorityLow    Priority = "Низкий"
)

func (p Priority) Validate() error {
	switch p {
	case PriorityHigh, PriorityMedium, PriorityLow:
		return nil
	default:
		return ErrInvalidPriority
	}
}
