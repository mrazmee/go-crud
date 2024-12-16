package libraries

import (
	"reflect"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Validation struct {
	validate *validator.Validate
	trans    ut.Translator
}

func NewValidation() *Validation {
	translator := en.New() // Membuat objek translator untuk bahasa Inggris
	uni := ut.New(translator, translator) // Membuat objek universal translator dengan bahasa Inggris

	trans, _ := uni.GetTranslator("en") // Mendapatkan translator untuk bahasa Inggris

	validate := validator.New() // Membuat objek validasi baru dari package validator/v10

	// Mendaftarkan terjemahan default untuk validasi bahasa Inggris
	en_translations.RegisterDefaultTranslations(validate, trans)

	// register tag label
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := field.Tag.Get("label")
		return name
	})

	//custom error ke bahasa indonesia
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} harus diisi", true)
		/*
		ut.Add("required", "{0} harus diisi", true):
		- "required": Ini adalah aturan validasi yang akan diterjemahkan.
		- "{0} harus diisi": Ini adalah format pesan yang diterjemahkan. {0} akan digantikan dengan nama field yang gagal divalidasi. Contohnya, jika field "Nama" gagal divalidasi, pesan error yang ditampilkan akan menjadi "Nama harus diisi".
		- true: Menandakan bahwa pesan ini wajib tersedia di translator.
		*/
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
		/*
		Fungsi kedua (untuk menerjemahkan pesan kesalahan):
		Fungsi ini digunakan untuk menerjemahkan dan menghasilkan pesan kesalahan yang sesuai untuk sebuah field yang gagal validasi.
		ut.T("required", fe.Field()):
		- ut.T("required", fe.Field()) mengambil terjemahan pesan kesalahan untuk aturan required dan mengganti {0} dengan nama field yang gagal divalidasi.
		- fe.Field() adalah nama field yang gagal validasi (misalnya "Nama" atau "Email").
		return t: Mengembalikan pesan yang telah diterjemahkan.
		*/
	})

	return &Validation{
		validate: validate, // Mengisi field `validate` dengan objek validator yang telah dibuat
		trans:    trans, // Mengisi field `trans` dengan translator bahasa Inggris
	}
}

func (v *Validation) Struct(s interface{}) interface{} {
	// 1. Membuat map untuk menyimpan error validasi
	errors := make(map[string]string)
	// 2. Melakukan validasi terhadap struct s menggunakan method Validate.Struct
    // v.validate adalah objek validator.Validate yang sudah disiapkan sebelumnya
    // s adalah parameter yang berisi data yang ingin divalidasi
	err := v.validate.Struct(s)

	// 3. Jika terjadi error (validasi gagal), proses error tersebut
	if err != nil {
		// 4. Iterasi atas setiap error validasi yang terjadi
		for _, e := range err.(validator.ValidationErrors) {
			// 5. Mengambil nama field yang gagal validasi menggunakan err.Field()
            // dan menerjemahkan pesan kesalahan menggunakan err.Translate(v.trans)
            // v.trans adalah translator yang digunakan untuk menerjemahkan pesan error ke bahasa yang diinginkan
			errors[e.StructField()] = e.Translate(v.trans)
		}
	}

	// 6. Jika ada error yang ditemukan (peta `errors` tidak kosong), kembalikan peta kesalahan
	if len(errors) > 0 {
		return errors
	}

	// 7. Jika tidak ada error validasi, kembalikan `nil`
	return nil
}