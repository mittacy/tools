package validator

import "testing"

type user struct {
	Id int 	`validate:"required"`
	Name string `validate:"min=1,max=10"`
	Age int `validate:"omitempty,min=1,max=150"`
}

func TestValidatorStruct(t *testing.T) {
	u1 := user {
		Id:   12,
		Name: "阿斯达",
		Age:  12,
	}
	if err := ValidatorStruct(&u1); err != nil {
		t.Errorf("%v 应该通过, err: %s\n", u1, err)
	}
	u2 := user {
		Name: "aaaaaaaaaaaaaaaaaaaaaaaa",
	}
	if err := ValidatorStruct(&u2); err == nil {
		t.Errorf("%v 不应该通过\n", u1)
	}
}

func TestValidatorVar(t *testing.T) {
	u := user{Id: 12, Name: "jack"}
	if err := ValidatorVar(u.Name, "name", "min=1,max=10"); err != nil {
		t.Errorf("%v 应该通过, err: %s\n", u, err)
	}
}
