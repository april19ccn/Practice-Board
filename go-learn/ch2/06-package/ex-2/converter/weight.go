package converter

import "fmt"

type Milligrams float64
type Gram float64
type Kilogram float64

func (m Milligrams) String() string { return fmt.Sprintf("%g mg", m) }

func (g Gram) String() string { return fmt.Sprintf("%g g", g) }

func (k Kilogram) String() string { return fmt.Sprintf("%g kg", k) }
