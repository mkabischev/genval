//This file was automatically generated by the genval generator
//Please don't modify it manually. Edit your entity tags and then
//run go generate

package overriding

import (
	"fmt"
)

type Validatable interface {
	Validate() error
}

func callValidateIfValidatable(i interface{}) error {
	if v, ok := i.(Validatable); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (r Age1) validate() error {
	if r.Value < 3 {
		return fmt.Errorf("field Value is less than 3 ")
	}
	if r.Value > 64 {
		return fmt.Errorf("field Value is more than 64 ")
	}
	return nil
}

func (r Age1) Validate() error {
	return r.validate()
}

func (r Age2) validate() error {
	if r.Value < 3 {
		return fmt.Errorf("field Value is less than 3 ")
	}
	if r.Value > 64 {
		return fmt.Errorf("field Value is more than 64 ")
	}
	return nil
}

func (r Age2) Validate() error {
	return r.validate()
}

func (r Age3) validate() error {
	if r.Value < 3 {
		return fmt.Errorf("field Value is less than 3 ")
	}
	if r.Value > 64 {
		return fmt.Errorf("field Value is more than 64 ")
	}
	return nil
}

func (r Age3) Validate() error {
	return r.validate()
}

func (r Age4) validate() error {
	if r.Value < 3 {
		return fmt.Errorf("field Value is less than 3 ")
	}
	if r.Value > 64 {
		return fmt.Errorf("field Value is more than 64 ")
	}
	return nil
}

func (r Age4) Validate() error {
	return r.validate()
}

func (r Age5) validate() error {
	if r.Value < 3 {
		return fmt.Errorf("field Value is less than 3 ")
	}
	if r.Value > 64 {
		return fmt.Errorf("field Value is more than 64 ")
	}
	return nil
}

func (r Age5) Validate() error {
	return r.validate()
}

func (r Request1) validate() error {
	if err := r.Age.Validate(); err != nil {
		return err
	}
	if r.Some < 3 {
		return fmt.Errorf("field Some is less than 3 ")
	}
	if r.Some > 64 {
		return fmt.Errorf("field Some is more than 64 ")
	}
	return nil
}

func (r Request2) validate() error {
	if err := r.Age.ValidateMin10(); err != nil {
		return err
	}
	if r.Some < 3 {
		return fmt.Errorf("field Some is less than 3 ")
	}
	if r.Some > 64 {
		return fmt.Errorf("field Some is more than 64 ")
	}
	return nil
}

func (r Request2) Validate() error {
	return r.validate()
}

func (r Request3) validate() error {
	if err := validateMin10(r.Age); err != nil {
		return err
	}
	if r.Some < 3 {
		return fmt.Errorf("field Some is less than 3 ")
	}
	if r.Some > 64 {
		return fmt.Errorf("field Some is more than 64 ")
	}
	return nil
}

func (r Request3) Validate() error {
	return r.validate()
}

func (r Request4) validate() error {
	if err := r.Age.ValidateMin10(); err != nil {
		return err
	}
	if err := validateMax128(r.Age); err != nil {
		return err
	}
	if r.Some < 3 {
		return fmt.Errorf("field Some is less than 3 ")
	}
	if r.Some > 64 {
		return fmt.Errorf("field Some is more than 64 ")
	}
	return nil
}

func (r Request4) Validate() error {
	return r.validate()
}

func (r Request5) validate() error {
	if err := r.Age.Validate(); err != nil {
		return err
	}
	if r.Some < 3 {
		return fmt.Errorf("field Some is less than 3 ")
	}
	if r.Some > 64 {
		return fmt.Errorf("field Some is more than 64 ")
	}
	return nil
}
