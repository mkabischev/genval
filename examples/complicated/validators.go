//This file was automatically generated by the genval generator
//Please don't modify it manually. Edit your entity tags and then
//run go generate

package complicated

import (
	"fmt"

	"unicode/utf8"
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

func (r AliasArray) validate() error {
	for _, x := range r {
		_ = x
	}
	return nil
}

func (r AliasArray) Validate() error {
	return r.validate()
}

func (r AliasChan) validate() error {
	return nil
}

func (r AliasChan) Validate() error {
	return r.validate()
}

func (r AliasFunc) validate() error {
	return nil
}

func (r AliasFunc) Validate() error {
	return r.validate()
}

func (r AliasOnDogsMapAlias) validate() error {
	if err := DogsMapAlias(r).Validate(); err != nil {
		return err
	}
	return nil
}

func (r AliasOnDogsMapAlias) Validate() error {
	return r.validate()
}

func (r AliasString) validate() error {
	return nil
}

func (r AliasString) Validate() error {
	return r.validate()
}

func (r Dog) validate() error {
	if utf8.RuneCountInString(r.Name) < 1 {
		return fmt.Errorf("field Name is shorter than 1 chars")
	}
	if utf8.RuneCountInString(r.Name) > 64 {
		return fmt.Errorf("field Name is longer than 64 chars")
	}
	return nil
}

func (r Dog) Validate() error {
	return r.validate()
}

func (r DogsMapAlias) validate() error {
	for k, v := range r {
		_ = k
		_ = v
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (r DogsMapAlias) Validate() error {
	return r.validate()
}

func (r State) validate() error {
	switch r {
	case StateOk:
	case StateError:
	default:
		return fmt.Errorf("invalid value for enum State: %v", r)
	}
	return nil
}

func (r Status) validate() error {
	switch r {
	case StatusCreated:
	case StatusPending:
	case StatusActive:
	case StatusFailed:
	default:
		return fmt.Errorf("invalid value for enum Status: %v", r)
	}
	return nil
}

func (r Status) Validate() error {
	return r.validate()
}

func (r User) validate() error {
	if utf8.RuneCountInString(r.Name) < 3 {
		return fmt.Errorf("field Name is shorter than 3 chars")
	}
	if utf8.RuneCountInString(r.Name) > 64 {
		return fmt.Errorf("field Name is longer than 64 chars")
	}
	if r.LastName != nil {
		if utf8.RuneCountInString(*r.LastName) < 1 {
			return fmt.Errorf("field LastName is shorter than 1 chars")
		}
		if utf8.RuneCountInString(*r.LastName) > 5 {
			return fmt.Errorf("field LastName is longer than 5 chars")
		}
	}
	if r.Age < 18 {
		return fmt.Errorf("field Age is less than 18 ")
	}
	if r.Age > 105 {
		return fmt.Errorf("field Age is more than 105 ")
	}
	if r.ChildrenCount == nil {
		return fmt.Errorf("field ChildrenCount is required, should not be nil")
	}
	if *r.ChildrenCount < 0 {
		return fmt.Errorf("field ChildrenCount is less than 0 ")
	}
	if *r.ChildrenCount > 15 {
		return fmt.Errorf("field ChildrenCount is more than 15 ")
	}
	if r.Float < -4.22 {
		return fmt.Errorf("field Float is less than -4.22 ")
	}
	if r.Float > 42.55 {
		return fmt.Errorf("field Float is more than 42.55 ")
	}
	if err := r.Dog.Validate(); err != nil {
		return err
	}
	if r.DogPointer != nil {
		if err := r.DogPointer.Validate(); err != nil {
			return err
		}
	}
	if err := r.DogOptional.ValidateOptional(); err != nil {
		return err
	}
	if len(r.Urls) < 1 {
		return fmt.Errorf("array Urls has less items than 1 ")
	}
	for _, x := range r.Urls {
		_ = x
		if utf8.RuneCountInString(x) > 256 {
			return fmt.Errorf("field x is longer than 256 chars")
		}
	}
	if len(r.Dogs) < 1 {
		return fmt.Errorf("array Dogs has less items than 1 ")
	}
	for _, x := range r.Dogs {
		_ = x
		if x != nil {
			if err := x.Validate(); err != nil {
				return err
			}
		}
	}
	if r.Test != nil {
		if len(*r.Test) < 1 {
			return fmt.Errorf("array Test has less items than 1 ")
		}
		for _, x := range *r.Test {
			_ = x
			if x < 4 {
				return fmt.Errorf("field x is less than 4 ")
			}
		}
	}
	if err := validateSome(r.Some); err != nil {
		return err
	}
	if len(r.SomeArray) < 1 {
		return fmt.Errorf("array SomeArray has less items than 1 ")
	}
	for _, x := range r.SomeArray {
		_ = x
		if err := validateSome(x); err != nil {
			return err
		}
	}
	if len(r.Dict) < 2 {
		return fmt.Errorf("map Dict has less items than 2 ")
	}
	for k, v := range r.Dict {
		_ = k
		_ = v
		if utf8.RuneCountInString(k) > 64 {
			return fmt.Errorf("field k is longer than 64 chars")
		}
		if v < -35 {
			return fmt.Errorf("field v is less than -35 ")
		}
		if v > 34 {
			return fmt.Errorf("field v is more than 34 ")
		}
	}
	for k, v := range r.DictDogs {
		_ = k
		_ = v
		if err := v.ValidateOptional(); err != nil {
			return err
		}
		if err := validateMaxDogName(v); err != nil {
			return err
		}
	}
	if err := r.Alias.Validate(); err != nil {
		return err
	}
	if err := r.AliasOnAlias.Validate(); err != nil {
		return err
	}
	if err := r.AliasOnAliasWithCustomValidate.ValidateAlias(); err != nil {
		return err
	}
	for k, v := range r.MapOfMap {
		_ = k
		_ = v
		if len(v) < 1 {
			return fmt.Errorf("map v has less items than 1 ")
		}
		for k, v := range v {
			_ = k
			_ = v
			if utf8.RuneCountInString(v) < 3 {
				return fmt.Errorf("field v is shorter than 3 chars")
			}
		}
	}
	return nil
}

func (r User) Validate() error {
	return r.validate()
}
