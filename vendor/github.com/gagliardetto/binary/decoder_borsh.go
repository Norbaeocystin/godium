// Copyright 2021 github.com/gagliardetto
// This file has been modified by github.com/gagliardetto
//
// Copyright 2020 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bin

import (
	"errors"
	"fmt"
	"io"
	"reflect"

	"go.uber.org/zap"
)

func (dec *Decoder) decodeWithOptionBorsh(v interface{}, option *option) (err error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return &InvalidDecoderError{reflect.TypeOf(v)}
	}

	// We decode rv not rv.Elem because the Unmarshaler interface
	// test must be applied at the top level of the value.
	err = dec.decodeBorsh(rv, option)
	if err != nil {
		return err
	}
	return nil
}

func (dec *Decoder) decodeBorsh(rv reflect.Value, opt *option) (err error) {
	if opt == nil {
		opt = newDefaultOption()
	}
	dec.currentFieldOpt = opt

	unmarshaler, rv := indirect(rv, opt.is_Optional() || opt.is_COptional())

	if traceEnabled {
		zlog.Debug("decode: type",
			zap.Stringer("value_kind", rv.Kind()),
			zap.Bool("has_unmarshaler", (unmarshaler != nil)),
			zap.Reflect("options", opt),
		)
	}

	if opt.is_Optional() {
		isPresent, e := dec.ReadOption()
		if e != nil {
			err = fmt.Errorf("decode: %t isPresent, %s", rv.Type(), e)
			return
		}

		if !isPresent {
			if traceEnabled {
				zlog.Debug("decode: skipping optional value", zap.Stringer("type", rv.Kind()))
			}

			rv.Set(reflect.Zero(rv.Type()))
			return
		}

		// we have ptr here we should not go get the element
		unmarshaler, rv = indirect(rv, false)
	}
	if opt.is_COptional() {
		isPresent, e := dec.ReadCOption()
		if e != nil {
			err = fmt.Errorf("decode: %t isPresent, %s", rv.Type(), e)
			return
		}

		if !isPresent {
			if traceEnabled {
				zlog.Debug("decode: skipping optional value", zap.Stringer("type", rv.Kind()))
			}

			rv.Set(reflect.Zero(rv.Type()))
			return
		}

		// we have ptr here we should not go get the element
		unmarshaler, rv = indirect(rv, false)
	}
	// Reset optionality so it won't propagate to child types:
	opt = opt.clone().set_Optional(false).set_COptional(false)

	if unmarshaler != nil {
		if traceEnabled {
			zlog.Debug("decode: using UnmarshalWithDecoder method to decode type")
		}
		return unmarshaler.UnmarshalWithDecoder(dec)
	}

	rt := rv.Type()
	switch rv.Kind() {
	// case reflect.Int:
	// 	// TODO: check if is x32 or x64
	// 	var n int64
	// 	n, err = dec.ReadInt64(LE)
	// 	rv.SetInt(n)
	// 	return
	// case reflect.Uint:
	// 	// TODO: check if is x32 or x64
	// 	var n uint64
	// 	n, err = dec.ReadUint64(LE)
	// 	rv.SetUint(n)
	// 	return
	case reflect.String:
		s, e := dec.ReadString()
		if e != nil {
			err = e
			return
		}
		rv.SetString(s)
		return
	case reflect.Uint8:
		var n byte
		n, err = dec.ReadByte()
		rv.SetUint(uint64(n))
		return
	case reflect.Int8:
		var n int8
		n, err = dec.ReadInt8()
		rv.SetInt(int64(n))
		return
	case reflect.Int16:
		var n int16
		n, err = dec.ReadInt16(LE)
		rv.SetInt(int64(n))
		return
	case reflect.Int32:
		var n int32
		n, err = dec.ReadInt32(LE)
		rv.SetInt(int64(n))
		return
	case reflect.Int64:
		var n int64
		n, err = dec.ReadInt64(LE)
		rv.SetInt(int64(n))
		return
	case reflect.Uint16:
		var n uint16
		n, err = dec.ReadUint16(LE)
		rv.SetUint(uint64(n))
		return
	case reflect.Uint32:
		var n uint32
		n, err = dec.ReadUint32(LE)
		rv.SetUint(uint64(n))
		return
	case reflect.Uint64:
		var n uint64
		n, err = dec.ReadUint64(LE)
		rv.SetUint(n)
		return
	case reflect.Float32:
		var n float32
		n, err = dec.ReadFloat32(LE)
		rv.SetFloat(float64(n))
		return
	case reflect.Float64:
		var n float64
		n, err = dec.ReadFloat64(LE)
		rv.SetFloat(n)
		return
	case reflect.Bool:
		var r bool
		r, err = dec.ReadBool()
		rv.SetBool(r)
		return
	case reflect.Interface:
		// Skip: cannot know the concrete type of the interface.
		// The parent container should implement a custom decoder.
		return nil
		// TODO: handle reflect.Ptr ???
	}
	switch rt.Kind() {
	case reflect.Array:
		l := rt.Len()
		if traceEnabled {
			zlog.Debug("decoding: reading array", zap.Int("length", l))
		}

		switch k := rv.Type().Elem().Kind(); k {
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if err := reflect_readArrayOfUint_(dec, l, k, rv, LE); err != nil {
				return err
			}
		default:
			for i := 0; i < l; i++ {
				if err = dec.decodeBorsh(rv.Index(i), nil); err != nil {
					return
				}
			}
		}
		return
	case reflect.Slice:
		var l int
		if opt.hasSizeOfSlice() {
			l = opt.getSizeOfSlice()
		} else {
			length, err := dec.ReadUint32(LE)
			if err != nil {
				return err
			}
			l = int(length)
		}

		if traceEnabled {
			zlog.Debug("reading slice", zap.Int("len", l), typeField("type", rv))
		}

		if l == 0 {
			// Empty slices are left nil
			return
		}
		if l > dec.Remaining() {
			return io.ErrUnexpectedEOF
		}

		switch k := rv.Type().Elem().Kind(); k {
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if err := reflect_readArrayOfUint_(dec, l, k, rv, LE); err != nil {
				return err
			}
		default:
			rv.Set(reflect.MakeSlice(rt, 0, 0))
			for i := 0; i < l; i++ {
				// create new element of type rt:
				element := reflect.New(rt.Elem())
				// decode into element:
				if err = dec.decodeBorsh(element, nil); err != nil {
					return
				}
				// append to slice:
				rv.Set(reflect.Append(rv, element.Elem()))
			}
		}

	case reflect.Struct:
		if err = dec.decodeStructBorsh(rt, rv); err != nil {
			return
		}

	case reflect.Map:
		l, err := dec.ReadUint32(LE)
		if err != nil {
			return err
		}
		if l == 0 {
			// If the map has no content, keep it nil.
			return nil
		}
		rv.Set(reflect.MakeMap(rt))
		for i := 0; i < int(l); i++ {
			key := reflect.New(rt.Key())
			err := dec.decodeBorsh(key.Elem(), nil)
			if err != nil {
				return err
			}
			val := reflect.New(rt.Elem())
			err = dec.decodeBorsh(val.Elem(), nil)
			if err != nil {
				return err
			}
			rv.SetMapIndex(key.Elem(), val.Elem())
		}
		return nil

	default:
		return fmt.Errorf("decode: unsupported type %q", rt)
	}

	return
}

func (dec *Decoder) deserializeComplexEnum(rv reflect.Value) error {
	rt := rv.Type()
	// read enum identifier
	tmp, err := dec.ReadUint8()
	if err != nil {
		return err
	}
	enum := BorshEnum(tmp)
	rv.Field(0).Set(reflect.ValueOf(enum).Convert(rv.Field(0).Type()))

	// read enum field, if necessary
	if int(enum)+1 >= rt.NumField() {
		return errors.New("complex enum too large")
	}
	field := rv.Field(int(enum) + 1)
	return dec.decodeBorsh(field, nil)
}

var borshEnumType = reflect.TypeOf(BorshEnum(0))

func isTypeBorshEnum(typ reflect.Type) bool {
	return typ.Kind() == reflect.Uint8 && typ == borshEnumType
}

func (dec *Decoder) decodeStructBorsh(rt reflect.Type, rv reflect.Value) (err error) {
	l := rv.NumField()

	if traceEnabled {
		zlog.Debug("decode: struct", zap.Int("fields", l), zap.Stringer("type", rv.Kind()))
	}

	// Handle complex enum:
	if rt.NumField() > 0 {
		// If the first field has type BorshEnum and is flagged with "borsh_enum"
		// we have a complex enum:
		firstField := rt.Field(0)
		if isTypeBorshEnum(firstField.Type) &&
			parseFieldTag(firstField.Tag).IsBorshEnum {
			return dec.deserializeComplexEnum(rv)
		}
	}

	sizeOfMap := map[string]int{}
	seenBinaryExtensionField := false
	for i := 0; i < l; i++ {
		structField := rt.Field(i)
		fieldTag := parseFieldTag(structField.Tag)

		if fieldTag.Skip {
			if traceEnabled {
				zlog.Debug("decode: skipping struct field with skip flag",
					zap.String("struct_field_name", structField.Name),
				)
			}
			continue
		}

		if !fieldTag.BinaryExtension && seenBinaryExtensionField {
			panic(fmt.Sprintf("the `bin:\"binary_extension\"` tags must be packed together at the end of struct fields, problematic field %q", structField.Name))
		}

		if fieldTag.BinaryExtension {
			seenBinaryExtensionField = true
			// FIXME: This works only if what is in `d.data` is the actual full data buffer that
			//        needs to be decoded. If there is for example two structs in the buffer, this
			//        will not work as we would continue into the next struct.
			//
			//        But at the same time, does it make sense otherwise? What would be the inference
			//        rule in the case of extra bytes available? Continue decoding and revert if it's
			//        not working? But how to detect valid errors?
			if len(dec.data[dec.pos:]) <= 0 {
				continue
			}
		}
		v := rv.Field(i)
		if !v.CanSet() {
			// This means that the field cannot be set, to fix this
			// we need to create a pointer to said field
			if !v.CanAddr() {
				// we cannot create a point to field skipping
				if traceEnabled {
					zlog.Debug("skipping struct field that cannot be addressed",
						zap.String("struct_field_name", structField.Name),
						zap.Stringer("struct_value_type", v.Kind()),
					)
				}
				return fmt.Errorf("unable to decode a none setup struc field %q with type %q", structField.Name, v.Kind())
			}
			v = v.Addr()
		}

		if !v.CanSet() {
			if traceEnabled {
				zlog.Debug("skipping struct field that cannot be addressed",
					zap.String("struct_field_name", structField.Name),
					zap.Stringer("struct_value_type", v.Kind()),
				)
			}
			continue
		}

		option := &option{
			is_OptionalField:  fieldTag.Option,
			is_COptionalField: fieldTag.COption,
			Order:             fieldTag.Order,
		}

		if s, ok := sizeOfMap[structField.Name]; ok {
			option.setSizeOfSlice(s)
		}

		if traceEnabled {
			zlog.Debug("decode: struct field",
				zap.Stringer("struct_field_value_type", v.Kind()),
				zap.String("struct_field_name", structField.Name),
				zap.Reflect("struct_field_tags", fieldTag),
				zap.Reflect("struct_field_option", option),
			)
		}

		rt := v.Type()
		ptrImplements := reflect.PtrTo(rt).Implements(unmarshalableType)
		vImplements := rt.Implements(unmarshalableType)
		if ptrImplements || vImplements {
			switch {
			case ptrImplements:
				m := reflect.New(rt)
				val := m.Interface()
				err := val.(BinaryUnmarshaler).UnmarshalWithDecoder(dec)
				if err != nil {
					return err
				}
				v.Set(reflect.ValueOf(val).Elem())
				continue
			case vImplements:
				m := reflect.New(rt.Elem())
				val := m.Interface()
				err := val.(BinaryUnmarshaler).UnmarshalWithDecoder(dec)
				if err != nil {
					return err
				}
				v.Set(reflect.ValueOf(val))
				continue
			}
		}

		if err = dec.decodeBorsh(v, option); err != nil {
			return fmt.Errorf("error while decoding %q field: %w", structField.Name, err)
		}

		if fieldTag.SizeOf != "" {
			size := sizeof(structField.Type, v)
			if traceEnabled {
				zlog.Debug("setting size of field",
					zap.String("field_name", fieldTag.SizeOf),
					zap.Int("size", size),
				)
			}
			sizeOfMap[fieldTag.SizeOf] = size
		}
	}
	return
}

var (
	marshalableType   = reflect.TypeOf((*BinaryMarshaler)(nil)).Elem()
	unmarshalableType = reflect.TypeOf((*BinaryUnmarshaler)(nil)).Elem()
)
