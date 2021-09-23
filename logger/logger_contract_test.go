package logger

import (
	"reflect"
	"testing"
)

func TestFormatValidator(t *testing.T){

	t.Run("check if it is validating correct formats", func(test *testing.T){
		replacersGot, err:=validateFormat("logger : %lda : %utc")
		if err != nil{
			test.Fatal(err)
		}
		replacersWant := []string{"%lda", "%utc"}
		if !reflect.DeepEqual(replacersGot, replacersWant) {
			test.Fatalf("replacers found are not correct")
		}
	})
	t.Run("check if it it gives error if format is empty", func(test *testing.T){
		_, err:=validateFormat("")
		if err == nil || err.Error()!="formater is empty"{
			test.Fatal()
		}
	})
	t.Run("check if gives error for not valid replacer", func(test *testing.T){
		_, err:=validateFormat("logger : %ldt : %utc")
		if err == nil || err.Error()!="invalid replacer %ldt"{
			test.Fatal()
		}
	})

	t.Run("check if gives error for no replacer found", func(test *testing.T){
		_, err:=validateFormat("logger")
		if err == nil || err.Error()!="invalid format, no replacer found"{
			test.Fatal()
		}
	})
}