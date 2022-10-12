package clid

import (
	"errors"
	"reflect"

	"github.com/urfave/cli/v2"
)

var (
	ErrNilData  = errors.New("data is nil")
	ErrNilValue = errors.New("value is nil")
)

func Decode(data *cli.Context, v any) error {
	if data == nil {
		return errors.New("data is nil")
	}
	if v == nil {
		return errors.New("value is nil")
	}

	objValue := reflect.ValueOf(v)

	for objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	for i := 0; i < objValue.NumField(); i++ {
		tag := objValue.Type().Field(i).Tag.Get("cli")
		fieldValue := objValue.Field(i)

		kind := fieldValue.Kind()

		if tag == "" && kind != reflect.Struct && kind != reflect.Ptr {
			continue
		}

		if !fieldValue.CanSet() {
			continue
		}

		switch kind {
		case reflect.String:
			fieldValue.SetString(data.String(tag))
		case reflect.Bool:
			fieldValue.SetBool(data.Bool(tag))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fieldValue.SetInt(int64(data.Int(tag)))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			fieldValue.SetUint(uint64(data.Uint(tag)))
		case reflect.Float32, reflect.Float64:
			fieldValue.SetFloat(data.Float64(tag))
		case reflect.Struct:
			err := Decode(data, fieldValue.Addr().Interface())
			if err != nil {
				return err
			}
		case reflect.Ptr:
			err := Decode(data, fieldValue.Interface())
			if err != nil {
				return err
			}
		default:
			panic("type " + kind.String() + " not implemented")
		}
	}

	return nil
}
