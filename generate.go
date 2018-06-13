package collection

//go:generate gotemplate -outfmt "collection_%v" "github.com/imiskolee/collection/template"  SliceInt(int)
//go:generate gotemplate -outfmt "collection_%v" "github.com/imiskolee/collection/template"  SliceInt8(int8)
//go:generate gotemplate -outfmt "collection_%v" "github.com/imiskolee/collection/template"  SliceInt16(int16)
//go:generate gotemplate -outfmt "collection_%v" "github.com/imiskolee/collection/template"  SliceInt32(int32)
//go:generate gotemplate -outfmt "collection_%v" "github.com/imiskolee/collection/template"  SliceInt64(int64)
//go:generate gotemplate -outfmt "collection_%v" "github.com/imiskolee/collection/template"  SliceUint(uint)
//go:generate gotemplate -outfmt "collection_%v" "github.com/imiskolee/collection/template"  SliceUint(uint8)
//go:generate gotemplate -outfmt "collection_%v" "github.com/imiskolee/collection/template"  SliceUint(uint16)
//go:generate gotemplate -outfmt "collection_%v" "github.com/imiskolee/collection/template"  SliceUint(uint32)
//go:generate gotemplate -outfmt "collection_%v" "github.com/imiskolee/collection/template"  SliceUint64(uint64)
//go:generate gotemplate -outfmt "collection_%v" "github.com/imiskolee/collection/template"  SliceFloat32(float32)
//go:generate gotemplate -outfmt "collection_%v" "github.com/imiskolee/collection/template"  SliceFloat64(float64)
//go:generate gotemplate -outfmt "collection_%v" "github.com/imiskolee/collection/template"  SliceByte(byte)
//go:generate gotemplate -outfmt "collection_%v" "github.com/imiskolee/collection/template"  SliceString(string)
//go:generate gotemplate -outfmt "collection_%v" "github.com/imiskolee/collection/template"  SliceBoolean(bool)