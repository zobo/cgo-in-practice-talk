func (r *Record) GetBytesField(d *Db, number uint16) ([]byte, error) {
	if r.GetFieldType(d, number) != STRTYPE {
		return nil, WDBError("Not an string valued field")
	}

	enc := C.wg_get_field(d.db, r.rec, C.wg_int(number))
	slen := int(C.wg_decode_str_len(d.db, enc))
	sval := C.wg_decode_str(d.db, enc)

	var goSlice []byte // HL
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&goSlice))) // HL
	sliceHeader.Cap = slen // HL
	sliceHeader.Len = slen // HL
	sliceHeader.Data = uintptr(unsafe.Pointer(sval)) // HL
	return goSlice, nil
}