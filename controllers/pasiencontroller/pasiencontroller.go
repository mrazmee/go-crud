package pasiencontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/mrazmee/go-crud/entities"
	"github.com/mrazmee/go-crud/libraries"
	"github.com/mrazmee/go-crud/models"
)

var validation = libraries.NewValidation()
var pasienModel = models.NewPasienModel()

func Index(response http.ResponseWriter, request *http.Request) {

	pasien, _ := pasienModel.FindAll()

	data := map[string]interface{}{
		"data_pasien" : pasien,
	}

	temp, err := template.ParseFiles("views/pasien/index.html")
	if err != nil {
		panic(err)
	}
	
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet{
		temp, err := template.ParseFiles("views/pasien/add.html")
		if err != nil {
			panic(err)
		}
		
		temp.Execute(response, nil)
	}else if request.Method == http.MethodPost{

		request.ParseForm()

		var pasien entities.Pasien

		//menerima inputan sesuai nama pada tag input di add.html

		pasien.NamaLengkap = request.Form.Get("nama_lengkap")
		pasien.NIK = request.Form.Get("nik")
		pasien.JenisKelamin = request.Form.Get("jenis_kelamin")
		pasien.TempatLahir = request.Form.Get("tempat_lahir")
		pasien.TanggalLahir = request.Form.Get("tanggal_lahir")
		pasien.Alamat = request.Form.Get("alamat")
		pasien.NoHp = request.Form.Get("no_hp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(pasien)

		if vErrors != nil{
			data["pasien"] = pasien
			data["validation"] = vErrors
		}else{
			pasienModel.Create(pasien)
			data["pesan"] = "Data berhasil disimpan"
		}

		temp, err := template.ParseFiles("views/pasien/add.html")
		if err != nil {
			panic(err)
		}
		
		temp.Execute(response, data)


	}

	
	
	
}

func Edit(response http.ResponseWriter, request *http.Request) {
	
	if request.Method == http.MethodGet{

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var pasien entities.Pasien
		pasienModel.Find(id, &pasien) // kenapa alamat? karena method structnya itu pointer maka nilai isiannya harus alamat (&)

		data := map[string]interface{}{
			"pasien" : pasien,
		}

		temp, err := template.ParseFiles("views/pasien/edit.html")
		if err != nil {
			panic(err)
		}
		
		temp.Execute(response, data)
	}else if request.Method == http.MethodPost{

		request.ParseForm()

		var pasien entities.Pasien

		//menerima inputan sesuai nama pada tag input di add.html

		pasien.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		pasien.NamaLengkap = request.Form.Get("nama_lengkap")
		pasien.NIK = request.Form.Get("nik")
		pasien.JenisKelamin = request.Form.Get("jenis_kelamin")
		pasien.TempatLahir = request.Form.Get("tempat_lahir")
		pasien.TanggalLahir = request.Form.Get("tanggal_lahir")
		pasien.Alamat = request.Form.Get("alamat")
		pasien.NoHp = request.Form.Get("no_hp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(pasien)

		if vErrors != nil{
			data["pasien"] = pasien
			data["validation"] = vErrors
		}else{
			pasienModel.Update(pasien)
			data["pesan"] = "Data berhasil diubah"
		}

		temp, err := template.ParseFiles("views/pasien/edit.html")
		if err != nil {
			panic(err)
		}
		
		temp.Execute(response, data)


	}

}

func Delete(response http.ResponseWriter, request *http.Request) {
	
	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	pasienModel.Delete(id)

	http.Redirect(response, request, "/pasien", http.StatusSeeOther)

}