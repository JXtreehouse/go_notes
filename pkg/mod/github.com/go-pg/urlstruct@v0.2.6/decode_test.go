package urlstruct_test

import (
	"database/sql"
	"encoding"
	"net/url"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/go-pg/urlstruct"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "urlstruct")
}

type CustomField struct {
	s string
}

var _ encoding.TextUnmarshaler = (*CustomField)(nil)

func (f *CustomField) UnmarshalText(text []byte) error {
	f.s = string(text)
	return nil
}

type SubFilter struct {
	Count int
}

var _ urlstruct.Unmarshaler = (*SubFilter)(nil)

func (f *SubFilter) UnmarshalValues(values url.Values) error {
	f.Count++
	return nil
}

type Filter struct {
	SubFilter
	Sub   SubFilter
	Count int

	Field    string
	FieldNEQ string
	FieldLT  int8
	FieldLTE int16
	FieldGT  int32
	FieldGTE int64

	Multi    []string
	MultiNEQ []int

	Time         time.Time
	StartTimeGTE time.Time

	NullBool    sql.NullBool
	NullInt64   sql.NullInt64
	NullFloat64 sql.NullFloat64
	NullString  sql.NullString

	Map    map[string]string
	Custom CustomField

	Omit []byte `pg:"-"`
}

var _ urlstruct.Unmarshaler = (*Filter)(nil)

func (f *Filter) UnmarshalValues(values url.Values) error {
	f.Count++
	return nil
}

var _ = Describe("Decode", func() {
	It("decodes struct from Values", func() {
		f := new(Filter)
		err := urlstruct.Unmarshal(url.Values{
			"field":      {"one"},
			"field__neq": {"two"},
			"field__lt":  {"1"},
			"field__lte": {"2"},
			"field__gt":  {"3"},
			"field__gte": {"4"},

			"multi":      {"one", "two"},
			"multi__neq": {"3", "4"},

			"time":            {"1970-01-01T00:00:00Z"},
			"start_time__gte": {"1970-01-01T00:00:00Z"},

			"null_bool":    {"t"},
			"null_int64":   {"1234"},
			"null_float64": {"1.234"},
			"null_string":  {"string"},

			"map[foo]":   {`bar`},
			"map[hello]": {`world`},
			"map[]":      {"invalid"},
			"map][":      {"invalid"},

			"custom": {"custom"},
		}, f)
		Expect(err).NotTo(HaveOccurred())

		Expect(f.Field).To(Equal("one"))
		Expect(f.FieldNEQ).To(Equal("two"))
		Expect(f.FieldLT).To(Equal(int8(1)))
		Expect(f.FieldLTE).To(Equal(int16(2)))
		Expect(f.FieldGT).To(Equal(int32(3)))
		Expect(f.FieldGTE).To(Equal(int64(4)))

		Expect(f.Multi).To(Equal([]string{"one", "two"}))
		Expect(f.MultiNEQ).To(Equal([]int{3, 4}))

		Expect(f.Time).To(BeTemporally("==", time.Unix(0, 0)))
		Expect(f.StartTimeGTE).To(BeTemporally("==", time.Unix(0, 0)))

		Expect(f.NullBool.Valid).To(BeTrue())
		Expect(f.NullBool.Bool).To(BeTrue())

		Expect(f.NullInt64.Valid).To(BeTrue())
		Expect(f.NullInt64.Int64).To(Equal(int64(1234)))

		Expect(f.NullFloat64.Valid).To(BeTrue())
		Expect(f.NullFloat64.Float64).To(Equal(float64(1.234)))

		Expect(f.NullString.Valid).To(BeTrue())
		Expect(f.NullString.String).To(Equal("string"))

		Expect(f.Map).To(Equal(map[string]string{"foo": "bar", "hello": "world"}))
		Expect(f.Custom.s).To(Equal("custom"))
	})

	It("supports names with suffix `[]`", func() {
		f := new(Filter)
		err := urlstruct.Unmarshal(url.Values{
			"field[]": {"one"},
		}, f)
		Expect(err).NotTo(HaveOccurred())
		Expect(f.Field).To(Equal("one"))
	})

	It("supports names with prefix `:`", func() {
		f := new(Filter)
		err := urlstruct.Unmarshal(url.Values{
			":field": {"one"},
		}, f)
		Expect(err).NotTo(HaveOccurred())
		Expect(f.Field).To(Equal("one"))
	})

	It("decodes sql.Null*", func() {
		f := new(Filter)
		err := urlstruct.Unmarshal(url.Values{
			"null_bool":    {""},
			"null_int64":   {""},
			"null_float64": {""},
			"null_string":  {""},
		}, f)
		Expect(err).NotTo(HaveOccurred())

		Expect(f.NullBool.Valid).To(BeTrue())
		Expect(f.NullBool.Bool).To(BeZero())

		Expect(f.NullInt64.Valid).To(BeTrue())
		Expect(f.NullInt64.Int64).To(BeZero())

		Expect(f.NullFloat64.Valid).To(BeTrue())
		Expect(f.NullFloat64.Float64).To(BeZero())

		Expect(f.NullString.Valid).To(BeTrue())
		Expect(f.NullString.String).To(BeZero())
	})

	It("calls UnmarshalValues", func() {
		f := new(Filter)
		err := urlstruct.Unmarshal(url.Values{}, f)
		Expect(err).NotTo(HaveOccurred())
		Expect(f.Count).To(Equal(1))
		Expect(f.SubFilter.Count).To(Equal(1))
		Expect(f.Sub.Count).To(Equal(1))
	})
})
