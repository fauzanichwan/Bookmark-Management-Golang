package main

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	sessions "github.com/kataras/go-sessions"
)

type Kategori struct {
	IdKategori int
	IdUser     int
	NmKategori string
	TglBuat    string
}

type KategoriTel struct {
	IdKategori     int
	NmKategori     string
	IdKategoriBook int
	PublishBook    int
}

type JoinKategori struct {
	IdKategoriKat int
	IdUserKat     int
	NmKategoriKat string
	TglBuatKat    string
	IdUserUsers   int
	NmDepanUsers  string
	FotoUsers     string
}

type Berkas struct {
	IdBerkas   int
	IdUser     int
	NamaBerkas string
	Berkas     string
	Tipefile   string
	TglBuat    string
	Status     int
	Publish    int
}

type JoinKirimanBook struct {
	IdKirim         int
	NmDepanPengirim string
	IdPenerima      int
	IdBook          int
	IdFile          int
	IdKategoriKat   int
	IdUserKat       int
	NmKategoriKat   string
	IdUserBook      int
	IdKategoriBook  int
	TglBuatBook     string
	JudulBook       string
	LinkBook        string
	IdUserUsers     int
	NmDepanUsers    string
}

type JoinKirimanFile struct {
	IdKirim         int
	NmDepanPengirim string
	IdPenerima      int
	IdFile          int
	IdFileFile      int
	IdUserFile      int
	NamaFileFile    string
	FileFile        string
	TipefileFile    string
	TglBuatFile     string
	IdUserUser      int
	NmDepanUser     string
}

type JoinTersimpanBook struct {
	IdSimpan         int
	IdUserSimpan     int
	IdUserPengirim   int
	IdBookmarkSimpan int
	IdBookmarkBook   int
	IdUserBook       int
	IdKategoriBook   int
	TglBuatBook      string
	JudulBook        string
	LinkBook         string
	IdUserUser       int
	NmDepanUser      string
	IdKategoriKat    int
	IdUserKat        int
	NmKategoriKat    string
}

type JoinTersimpanFile struct {
	IdSimpan         int
	IdUserSimpan     int
	IdUserPengirim   int
	IdFileSimpan     int
	IdBerkasBerkas   int
	IdUserBerkas     int
	NamaBerkasBerkas string
	BerkasBerkas     string
	TipefileBerkas   string
	TglBuatBerkas    string
	IdUserUser       int
	NmDepanUser      string
}

type JoinBerkas struct {
	IdBerkasBerkas   int
	IdUserBerkas     int
	NamaBerkasBerkas string
	BerkasBerkas     string
	TipefileBerkas   string
	TglBuatBerkas    string
	PublishBerkas    int
	UpdateWaktu      string
	IdUserUser       int
	NmDepanUser      string
	Avatar           string
}

type JoinBook struct {
	IdKategoriKat  int
	NmKategoriKat  string
	IdBookmarkBook int
	IdKategoriBook int
	IdUserBook     int
	JudulBook      string
	LinkBook       string
	TglBuatBook    string
	StatusBook     int
	PublishBook    int
}

type JoinBookTel struct {
	IdKategoriKat  int
	IdUserKat      int
	NmKategoriKat  string
	IdBookmarkBook int
	IdKategoriBook int
	IdUserBook     int
	JudulBook      string
	LinkBook       string
	TglBuatBook    string
	PublishBook    int
	UpdateWaktu    string
	IdUserUser     int
	NmDepanUser    string
	Avatar         string
}

type JoinBookRe struct {
	IdReport       int
	IdUserRe       int
	NmDepanRe      string
	IdBookRe       int
	IndikatorRe    string
	IdBookmarkBook int
	IdUserBook     int
	IdKategoriBook int
	TglBuatBook    string
	JudulBook      string
	LinkBook       string
	IdUserUser     int
	NmDepanUser    string
	IdKategoriKat  int
	IdUserKat      int
	NmKategoriKat  string
}

type JoinUserRe struct {
	IdReport       int
	IdUserRe       int
	NmDepanRe      string
	IdUserReportRe int
	IndikatorRe    string
	IdUser         int
	NmDepan        string
	NmBelakang     string
	Email          string
	Jk             string
	Username       string
	Password       string
	TglGabung      string
	Foto           string
	Role           int
}

type JoinFileRe struct {
	IdReport     int
	IdUserRe     int
	NmDepanRe    string
	IdFileRe     int
	IndikatorRe  string
	IdFileFile   int
	IdUserFile   int
	NamaFileFile string
	FileFile     string
	TipefileFile string
	TglBuatFile  string
	StatusFile   int
	PublishFile  int
	IdUserUser   int
	NmDepanUser  string
}

type JoinTopikRe struct {
	IdReport    int
	IdUserRe    int
	NmDepanRe   string
	IdTopikRe   int
	IndikatorRe string
	IdTopik     int
	IdUser      int
	Topik       string
	TglPost     string
	NmDepan     string
	Avatar      string
	IdUserUser  int
	NmDepanUser string
}

type Bookmark struct {
	IdBookmark int
	IdKategori int
	IdUser     int
	Judul      string
	Link       string
	TglBuat    string
	Status     int
	Publish    int
}

type Users struct {
	IdUser     int
	NmDepan    string
	NmBelakang string
	Email      string
	Jk         string
	Username   string
	Password   string
	TglGabung  string
	Foto       string
	Role       int
}

type Topik struct {
	IdTopik int
	IdUser  int
	Topik   string
	TglPost string
	NmDepan string
	Avatar  string
}

type JoinTopik struct {
	IdTopikTopik    int
	IdUserTopik     int
	TopikTopik      string
	TglPostTopik    string
	IdUserUsers     int
	NmDepanUsers    string
	NmBelakangUsers string
	FotoUsers       string
}

type Komentar struct {
	IdKomentar int
	IdTopik    int
	IdUser     int
	Komentar   string
	TglPost    string
}

type Komen struct {
	IdTopikTopik    int
	IdUserTopik     int
	TopikTopik      string
	TglPostTopik    string
	IdUserUsers     int
	NmDepanUsers    string
	NmBelakangUsers string
	FotoUsers       string
	Komentar        []JoinKomentar
}

type Book struct {
	Kategori []Kategori
	Bookmark []JoinBook
}

type BookTel struct {
	Kategori []KategoriTel
	Bookmark []JoinBookTel
}

type KirimBook struct {
	IdBookmarkBook int
	Users          []Users
}

type KirimFile struct {
	IdFileFile int
	Users      []Users
}

type JoinKomentar struct {
	IdKomentarKomen int
	IdTopikKomen    int
	IdUserKomen     int
	KomentarKomen   string
	TglPostKomen    string
	IdTopikTopik    int
	IdUserTopik     int
	TopikTopik      string
	TglPostTopik    string
	NmDepanTopik    string
	AvatarTopik     string
	IdUserUsers     int
	NmDepanUsers    string
	FotoUsers       string
}

type JoinKomentarRe struct {
	IdReport        int
	IdUserRe        int
	NmDepanRe       string
	IdKomentarRe    int
	IndikatorRe     string
	IdKomentarKomen int
	IdTopikKomen    int
	IdUserKomen     int
	KomentarKomen   string
	TglPostKomen    string
	IdUserUsers     int
	NmDepanUsers    string
	FotoUsers       string
}

func hassdata(s string) string {
	var sha = sha1.New()
	sha.Write([]byte(s))
	var encrypted = sha.Sum(nil)
	var encryptedstring = fmt.Sprintf("%x", encrypted)
	return encryptedstring
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbName := "manajemen"
	db, err := sql.Open(dbDriver, dbUser+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("views/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	var data = make(map[string]string)
	data["err"] = r.URL.Query().Get("err")
	tmpl.ExecuteTemplate(w, "index", data)
}

func lihatkatAdmin(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	selDB, err := db.Query("SELECT kategori.idKategori, kategori.idUser, kategori.nmKategori, kategori.waktuPost, users.idUser, users.nmDepan, users.Foto FROM kategori, users WHERE kategori.idUser = users.idUser ORDER BY kategori.idKategori DESC")
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinKategori
	for selDB.Next() {
		var item = JoinKategori{}
		err = selDB.Scan(&item.IdKategoriKat, &item.IdUserKat, &item.NmKategoriKat, &item.TglBuatKat, &item.IdUserUsers, &item.NmDepanUsers, &item.FotoUsers)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "lihatkatAdmin", result)
	defer db.Close()
}

func tambahberkas(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	tmpl.ExecuteTemplate(w, "tambahberkas", nil)
}

func lihatberkas(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nId := suserid
	selDB, err := db.Query("SELECT idFile, idUser, namaFile, file, tipefile, waktuPost, status, publish FROM file WHERE idUser=? ORDER BY idFile DESC", nId)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []Berkas
	for selDB.Next() {
		var item = Berkas{}
		err = selDB.Scan(&item.IdBerkas, &item.IdUser, &item.NamaBerkas, &item.Berkas, &item.Tipefile, &item.TglBuat, &item.Status, &item.Publish)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "lihatberkas", result)
	defer db.Close()
}

func sortfile(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nId := suserid
	id := r.FormValue("tipefile")
	selDB, err := db.Query("SELECT idFile, idUser, namaFile, file, tipefile, waktuPost, status, publish FROM file WHERE idUser=? AND tipefile=? ORDER BY idFile DESC", nId, id)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []Berkas
	for selDB.Next() {
		var item = Berkas{}
		err = selDB.Scan(&item.IdBerkas, &item.IdUser, &item.NamaBerkas, &item.Berkas, &item.Tipefile, &item.TglBuat, &item.Status, &item.Publish)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "sortfile", result)
	defer db.Close()
}

func daftaruser(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	role0 := 0
	role2 := 2
	selDB, err := db.Query("SELECT idUser, nmDepan, nmBelakang, tglGabung, foto, role FROM users WHERE role=? OR role=? ORDER BY idUser DESC", role0, role2)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []Users
	for selDB.Next() {
		var item = Users{}
		err = selDB.Scan(&item.IdUser, &item.NmDepan, &item.NmBelakang, &item.TglGabung, &item.Foto, &item.Role)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "daftaruser", result)
	defer db.Close()
}

func daftaruserAdmin(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	role0 := 0
	selDB, err := db.Query("SELECT idUser, nmDepan, nmBelakang, tglGabung, foto, role FROM users WHERE role=? ORDER BY idUser DESC", role0)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []Users
	for selDB.Next() {
		var item = Users{}
		err = selDB.Scan(&item.IdUser, &item.NmDepan, &item.NmBelakang, &item.TglGabung, &item.Foto, &item.Role)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "daftaruserAdmin", result)
	defer db.Close()
}

func daftaruserNonAktif(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	role0 := 2
	selDB, err := db.Query("SELECT idUser, nmDepan, nmBelakang, tglGabung, foto, role FROM users WHERE role=? ORDER BY idUser DESC", role0)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []Users
	for selDB.Next() {
		var item = Users{}
		err = selDB.Scan(&item.IdUser, &item.NmDepan, &item.NmBelakang, &item.TglGabung, &item.Foto, &item.Role)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "daftaruserNonAktif", result)
	defer db.Close()
}

func kirimbook(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nid := r.URL.Query().Get("IdBook")
	role0 := 0
	role2 := 2
	selDB, err := db.Query("SELECT idUser, nmDepan, nmBelakang, tglGabung, foto, role FROM users WHERE role=? OR role=? ORDER BY idUser DESC", role0, role2)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []Users
	for selDB.Next() {
		var item = Users{}
		err = selDB.Scan(&item.IdUser, &item.NmDepan, &item.NmBelakang, &item.TglGabung, &item.Foto, &item.Role)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	selDB, err = db.Query("SELECT idBookmark FROM bookmark where idBookmark=?", nid)
	if err != nil {
		panic(err.Error())
	}
	tpk := Bookmark{}
	for selDB.Next() {
		var idBookmark int
		err = selDB.Scan(&idBookmark)
		if err != nil {
			panic(err.Error())
		}
		tpk.IdBookmark = idBookmark
	}
	var kirimBook = KirimBook{}
	kirimBook.IdBookmarkBook = tpk.IdBookmark
	kirimBook.Users = result
	tmpl.ExecuteTemplate(w, "kirimbook", kirimBook)
	defer db.Close()
}

func kirimfile(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nid := r.URL.Query().Get("IdFile")
	role0 := 0
	role2 := 2
	selDB, err := db.Query("SELECT idUser, nmDepan, nmBelakang, tglGabung, foto, role FROM users WHERE role=? OR role=? ORDER BY idUser DESC", role0, role2)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []Users
	for selDB.Next() {
		var item = Users{}
		err = selDB.Scan(&item.IdUser, &item.NmDepan, &item.NmBelakang, &item.TglGabung, &item.Foto, &item.Role)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	selDB, err = db.Query("SELECT idFile FROM file where idFile=?", nid)
	if err != nil {
		panic(err.Error())
	}
	tpk := Berkas{}
	for selDB.Next() {
		var idFile int
		err = selDB.Scan(&idFile)
		if err != nil {
			panic(err.Error())
		}
		tpk.IdBerkas = idFile
	}
	var kirimFile = KirimFile{}
	kirimFile.IdFileFile = tpk.IdBerkas
	kirimFile.Users = result
	tmpl.ExecuteTemplate(w, "kirimfile", kirimFile)
	defer db.Close()
}

func tambahkat(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	tmpl.ExecuteTemplate(w, "tambahkat", nil)
}

func tambahbook(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nId := suserid
	selDB, err := db.Query("SELECT idKategori, idUser, nmKategori FROM kategori WHERE idUser=?", nId)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []Kategori
	for selDB.Next() {
		var item = Kategori{}
		err = selDB.Scan(&item.IdKategori, &item.IdUser, &item.NmKategori)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "tambahbook", result)
	defer db.Close()
}

func authlogin(w http.ResponseWriter, r *http.Request) {
	var db = dbConn()
	defer db.Close()
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		hasspass := hassdata(password)
		sql := "SELECT idUser, nmDepan, role FROM users WHERE username =? AND password =?"
		var data = Users{}
		var err = db.QueryRow(sql, username, hasspass).Scan(&data.Username, &data.NmDepan, &data.Role)
		if err != nil {
			http.Redirect(w, r, "index?err=Maaf, Email dan Password anda Tidak Valid", 301)
		} else {
			session := sessions.Start(w, r)
			session.Set("suserid", data.Username)
			session.Set("nmDepan", data.NmDepan)
			session.Set("role", data.Role)
			http.Redirect(w, r, "home", 301)
		}
	}
}

func proccessreg(w http.ResponseWriter, r *http.Request) {
	var db = dbConn()
	defer db.Close()
	nmdepan := r.FormValue("nmDepan")
	nmbelakang := r.FormValue("nmBelakang")
	email := r.FormValue("email")
	jk := r.FormValue("jk")
	username := r.FormValue("username")
	password := r.FormValue("password")
	hasspass := hassdata(password)
	ava := "default.png"
	if r.Method == "POST" {
		if nmdepan == "" {
			fmt.Fprintln(w, "Isi nama depan")
		} else if nmbelakang == "" {
			fmt.Fprintln(w, "Isi nama belakang")
		} else if email == "" {
			fmt.Fprintln(w, "Isi email")
		} else if jk == "" {
			fmt.Fprintln(w, "Pilih jenis kelamin")
		} else if username == "" {
			fmt.Fprintln(w, "Isi username")
		} else {
			insForm, err := db.Prepare("INSERT INTO users(nmDepan, nmBelakang, email, jk, username, password, foto) VALUES(?,?,?,?,?,?,?)")
			if err != nil {
				panic(err.Error())
			}
			insForm.Exec(nmdepan, nmbelakang, email, jk, username, hasspass, ava)
		}
		defer db.Close()
		http.Redirect(w, r, "index?err=Terima Kasih Sudah Mendaftar, silhkan login", 301)
	}
}

func addkategori(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	idUser := suserid
	nmKategori := r.FormValue("nmkategori")
	stmt, err := db.Prepare("INSERT kategori SET idUser=?, nmKategori=?")
	if err == nil {
		_, err := stmt.Exec(&idUser, &nmKategori)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/lihatkat", http.StatusSeeOther)
		return
	}
}

func proseskirimbook(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var nmDepan = session.GetString("nmDepan")
	var role = session.GetString("role")
	var data = make(map[string]string)
	data["suserid"] = suserid
	data["nmDepan"] = nmDepan
	data["role"] = role
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	IdBook := r.FormValue("IdBook")
	IdPenerima := r.FormValue("IdPenerima")
	stmt, err := db.Prepare("INSERT kiriman SET nmDepanPengirim=?, idPenerima=?, idBook=?")
	if err == nil {
		_, err := stmt.Exec(&nmDepan, &IdPenerima, &IdBook)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/telusuribook", http.StatusSeeOther)
		return
	}
}

func proseskirimfile(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var nmDepan = session.GetString("nmDepan")
	var role = session.GetString("role")
	var data = make(map[string]string)
	data["suserid"] = suserid
	data["nmDepan"] = nmDepan
	data["role"] = role
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	IdFile := r.FormValue("IdFile")
	IdPenerima := r.FormValue("IdPenerima")
	stmt, err := db.Prepare("INSERT kiriman SET nmDepanPengirim=?, idPenerima=?, idFile=?")
	if err == nil {
		_, err := stmt.Exec(&nmDepan, &IdPenerima, &IdFile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/telusurifile", http.StatusSeeOther)
		return
	}
}

func addkomen(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nid := r.FormValue("nid")
	idUser := suserid
	kom := r.FormValue("komentar")
	stmt, err := db.Prepare("INSERT komentar SET idTopik=?, idUser=?, komentar=?")
	if err == nil {
		_, err := stmt.Exec(&nid, &idUser, &kom)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/lihatkomentar?IdTopik="+nid, http.StatusSeeOther)
		return
	}
}

func addkomenAdmin(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nid := r.FormValue("nid")
	idUser := suserid
	kom := r.FormValue("komentar")
	stmt, err := db.Prepare("INSERT komentar SET idTopik=?, idUser=?, komentar=?")
	if err == nil {
		_, err := stmt.Exec(&nid, &idUser, &kom)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/lihatkomentarAdmin?IdTopik="+nid, http.StatusSeeOther)
		return
	}
}

func addtopik(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	idUser := suserid
	tpk := r.FormValue("topik")
	nm := r.FormValue("nmDepan")
	ava := r.FormValue("avatar")
	stmt, err := db.Prepare("INSERT topik SET idUser=?, topik=?, nmDepan=?, avatar=?")
	if err == nil {
		_, err := stmt.Exec(&idUser, &tpk, &nm, &ava)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/lihattopik", http.StatusSeeOther)
		return
	}
}

func simpanbook(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	idUser := suserid
	idUP := r.FormValue("IdUP")
	idBook := r.FormValue("IdBook")
	stmt, err := db.Prepare("INSERT tersimpan SET idUser=?, idUserPengirim=?, idBookmark=?")
	if err == nil {
		_, err := stmt.Exec(&idUser, &idUP, &idBook)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/tersimpanbook", http.StatusSeeOther)
		return
	}
}

func simpanfile(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	idUser := suserid
	idUP := r.FormValue("IdUP")
	idFile := r.FormValue("IdFile")
	stmt, err := db.Prepare("INSERT tersimpan SET idUser=?, idUserPengirim=?, idFile=?")
	if err == nil {
		_, err := stmt.Exec(&idUser, &idUP, &idFile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/tersimpanfile", http.StatusSeeOther)
		return
	}
}

func laporkantopik(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var nmDepan = session.GetString("nmDepan")
	var role = session.GetString("role")
	var data = make(map[string]string)
	data["suserid"] = suserid
	data["nmDepan"] = nmDepan
	data["role"] = role
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	idUser := suserid
	nm := nmDepan
	idTopik := r.FormValue("IdTopik")
	indikator := r.FormValue("indikator")
	stmt, err := db.Prepare("INSERT report SET idUser=?, nmDepan=?, idTopik=?, alasanPelaporan=?")
	if err == nil {
		_, err := stmt.Exec(&idUser, &nm, &idTopik, &indikator)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/lihattopik", http.StatusSeeOther)
		return
	}
}

func laporkanuser(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var nmDepan = session.GetString("nmDepan")
	var role = session.GetString("role")
	var data = make(map[string]string)
	data["suserid"] = suserid
	data["nmDepan"] = nmDepan
	data["role"] = role
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	idUser := suserid
	nm := nmDepan
	id := r.FormValue("IdUser")
	indikator := r.FormValue("indikator")
	stmt, err := db.Prepare("INSERT report SET idUser=?, nmDepan=?, idUserReport=?, alasanPelaporan=?")
	if err == nil {
		_, err := stmt.Exec(&idUser, &nm, &id, &indikator)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/daftaruser", http.StatusSeeOther)
		return
	}
}

func laporkankomen(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var nmDepan = session.GetString("nmDepan")
	var role = session.GetString("role")
	var data = make(map[string]string)
	data["suserid"] = suserid
	data["nmDepan"] = nmDepan
	data["role"] = role
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	idUser := suserid
	nm := nmDepan
	idKomen := r.FormValue("IdKomen")
	idTopik := r.FormValue("IdTopik")
	indikator := r.FormValue("indikator")
	stmt, err := db.Prepare("INSERT report SET idUser=?, nmDepan=?, idKomentar=?, alasanPelaporan=?")
	if err == nil {
		_, err := stmt.Exec(&idUser, &nm, &idKomen, &indikator)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/lihatkomentar?IdTopik="+idTopik, http.StatusSeeOther)
		return
	}
}

func laporkanbook(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var nmDepan = session.GetString("nmDepan")
	var role = session.GetString("role")
	var data = make(map[string]string)
	data["suserid"] = suserid
	data["nmDepan"] = nmDepan
	data["role"] = role
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	idUser := suserid
	nm := nmDepan
	idBook := r.FormValue("IdBook")
	indikator := r.FormValue("indikator")
	stmt, err := db.Prepare("INSERT report SET idUser=?, nmDepan=?, idBook=?, alasanPelaporan=?")
	if err == nil {
		_, err := stmt.Exec(&idUser, &nm, &idBook, &indikator)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/telusuribook", http.StatusSeeOther)
		return
	}
}

func laporkanfile(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var nmDepan = session.GetString("nmDepan")
	var role = session.GetString("role")
	var data = make(map[string]string)
	data["suserid"] = suserid
	data["nmDepan"] = nmDepan
	data["role"] = role
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	idUser := suserid
	nm := nmDepan
	idFile := r.FormValue("IdFile")
	indikator := r.FormValue("indikator")
	stmt, err := db.Prepare("INSERT report SET idUser=?, nmDepan=?, idFile=?, alasanPelaporan=?")
	if err == nil {
		_, err := stmt.Exec(&idUser, &nm, &idFile, &indikator)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/telusurifile", http.StatusSeeOther)
		return
	}
}

func addbookmark(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	IdKategori := r.FormValue("kategori")
	IdUser := suserid
	Judul := r.FormValue("judul")
	Link := r.FormValue("link")
	stmt, err := db.Prepare("INSERT bookmark SET idKategori=?, idUser=?, judul=?, link=?")
	if err == nil {
		_, err := stmt.Exec(&IdKategori, &IdUser, &Judul, &Link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/lihatbook", http.StatusSeeOther)
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var nmDepan = session.GetString("nmDepan")
	var role = session.GetString("role")
	var data = make(map[string]string)
	data["suserid"] = suserid
	data["nmDepan"] = nmDepan
	data["role"] = role
	data["err"] = r.URL.Query().Get("err")
	if suserid != "" {
		tmpl.ExecuteTemplate(w, "home", data)
	} else {
		http.Redirect(w, r, "index?err=Harap login terlebih dahulu", 301)
	}
}

func lihatkat(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nId := suserid
	selDB, err := db.Query("SELECT idKategori, idUser, nmKategori, waktuPost FROM kategori WHERE idUser=? ORDER BY idKategori DESC", nId)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []Kategori
	for selDB.Next() {
		var item = Kategori{}
		err = selDB.Scan(&item.IdKategori, &item.IdUser, &item.NmKategori, &item.TglBuat)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "lihatkat", result)
	defer db.Close()
}

func lihatbook(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nId := suserid
	selDB, err := db.Query("SELECT kategori.idKategori, kategori.nmKategori, bookmark.idBookmark, bookmark.idKategori, bookmark.idUser, bookmark.judul, bookmark.link, bookmark.waktuPost, bookmark.status, bookmark.publish FROM kategori, bookmark WHERE kategori.idKategori = bookmark.idKategori AND bookmark.idUser=? ORDER BY idBookmark DESC", nId)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinBook
	for selDB.Next() {
		var item = JoinBook{}
		err = selDB.Scan(&item.IdKategoriKat, &item.NmKategoriKat, &item.IdBookmarkBook, &item.IdKategoriKat, &item.IdUserBook, &item.JudulBook, &item.LinkBook, &item.TglBuatBook, &item.StatusBook, &item.PublishBook)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	selDB, err = db.Query("SELECT * FROM kategori where idUser=?", suserid)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var out []Kategori
	for selDB.Next() {
		var item = Kategori{}
		err = selDB.Scan(&item.IdKategori, &item.IdUser, &item.NmKategori, &item.TglBuat)
		if err != nil {
			panic(err.Error())
		}
		out = append(out, item)
	}
	var hasil = Book{}
	hasil.Bookmark = result
	hasil.Kategori = out
	tmpl.ExecuteTemplate(w, "lihatbook", hasil)
	defer db.Close()
}

func profil(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nId := suserid
	selDB, err := db.Query("SELECT idUser, nmDepan, nmBelakang, email, jk, username, tglGabung, foto FROM users WHERE idUser=?", nId)
	if err != nil {
		panic(err.Error())
	}
	usr := Users{}
	for selDB.Next() {
		var idUser int
		var nmDepan, nmBelakang, email, jk, username, tglGabung, foto string
		err = selDB.Scan(&idUser, &nmDepan, &nmBelakang, &email, &jk, &username, &tglGabung, &foto)
		if err != nil {
			panic(err.Error())
		}
		usr.IdUser = idUser
		usr.NmDepan = nmDepan
		usr.NmBelakang = nmBelakang
		usr.Email = email
		usr.Jk = jk
		usr.Username = username
		usr.TglGabung = tglGabung
		usr.Foto = foto
	}
	tmpl.ExecuteTemplate(w, "profil", usr)
	defer db.Close()
}

func tambahtopik(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nId := suserid
	selDB, err := db.Query("SELECT idUser, nmDepan, foto FROM users WHERE idUser=?", nId)
	if err != nil {
		panic(err.Error())
	}
	usr := Users{}
	for selDB.Next() {
		var idUser int
		var nmDepan, foto string
		err = selDB.Scan(&idUser, &nmDepan, &foto)
		if err != nil {
			panic(err.Error())
		}
		usr.IdUser = idUser
		usr.NmDepan = nmDepan
		usr.Foto = foto
	}
	tmpl.ExecuteTemplate(w, "tambahtopik", usr)
	defer db.Close()
}

func lihatuser(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nId := r.URL.Query().Get("IdUser")
	selDB, err := db.Query("SELECT idUser, nmDepan, nmBelakang, email, jk, username, tglGabung, foto FROM users WHERE idUser=?", nId)
	if err != nil {
		panic(err.Error())
	}
	usr := Users{}
	for selDB.Next() {
		var idUser int
		var nmDepan, nmBelakang, email, jk, username, tglGabung, foto string
		err = selDB.Scan(&idUser, &nmDepan, &nmBelakang, &email, &jk, &username, &tglGabung, &foto)
		if err != nil {
			panic(err.Error())
		}
		usr.IdUser = idUser
		usr.NmDepan = nmDepan
		usr.NmBelakang = nmBelakang
		usr.Email = email
		usr.Jk = jk
		usr.Username = username
		usr.TglGabung = tglGabung
		usr.Foto = foto
	}
	tmpl.ExecuteTemplate(w, "lihatuser", usr)
	defer db.Close()
}

func lihatuserAdmin(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nId := r.URL.Query().Get("IdUser")
	selDB, err := db.Query("SELECT idUser, nmDepan, nmBelakang, email, jk, username, tglGabung, foto FROM users WHERE idUser=?", nId)
	if err != nil {
		panic(err.Error())
	}
	usr := Users{}
	for selDB.Next() {
		var idUser int
		var nmDepan, nmBelakang, email, jk, username, tglGabung, foto string
		err = selDB.Scan(&idUser, &nmDepan, &nmBelakang, &email, &jk, &username, &tglGabung, &foto)
		if err != nil {
			panic(err.Error())
		}
		usr.IdUser = idUser
		usr.NmDepan = nmDepan
		usr.NmBelakang = nmBelakang
		usr.Email = email
		usr.Jk = jk
		usr.Username = username
		usr.TglGabung = tglGabung
		usr.Foto = foto
	}
	tmpl.ExecuteTemplate(w, "lihatuserAdmin", usr)
	defer db.Close()
}

func lihatpub(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nId := suserid
	idSelector := 1
	selDB, err := db.Query("SELECT kategori.idKategori, kategori.nmKategori, bookmark.idBookmark, bookmark.idKategori, bookmark.idUser, bookmark.judul, bookmark.link, bookmark.waktuPost, bookmark.status, bookmark.publish FROM kategori, bookmark WHERE kategori.idKategori = bookmark.idKategori AND bookmark.idUser=? AND bookmark.publish=? ORDER BY idBookmark DESC", nId, idSelector)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinBook
	for selDB.Next() {
		var item = JoinBook{}
		err = selDB.Scan(&item.IdKategoriKat, &item.NmKategoriKat, &item.IdBookmarkBook, &item.IdKategoriKat, &item.IdUserBook, &item.JudulBook, &item.LinkBook, &item.TglBuatBook, &item.StatusBook, &item.PublishBook)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "lihatpub", result)
	defer db.Close()
}

func lihatpubfile(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nId := suserid
	idSelector := 1
	selDB, err := db.Query("SELECT idFile, idUser, namaFile, file, tipefile, waktuPost, status, publish FROM file WHERE idUser=? AND publish=? ORDER BY idFile DESC", nId, idSelector)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []Berkas
	for selDB.Next() {
		var item = Berkas{}
		err = selDB.Scan(&item.IdBerkas, &item.IdUser, &item.NamaBerkas, &item.Berkas, &item.Tipefile, &item.TglBuat, &item.Status, &item.Publish)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "lihatpubfile", result)
	defer db.Close()
}

func addpublish(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		nid := r.FormValue("selector")
		idbook := r.FormValue("idbook")
		insForm, err := db.Prepare("UPDATE bookmark SET publish=? WHERE idbookmark=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(nid, idbook)
	}

	defer db.Close()
	http.Redirect(w, r, "/lihatpub", 301)
}

func addpublishfile(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		nid := r.FormValue("selector")
		idfile := r.FormValue("idfile")
		insForm, err := db.Prepare("UPDATE file SET publish=? WHERE idFile=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(nid, idfile)
	}

	defer db.Close()
	http.Redirect(w, r, "/lihatpubfile", 301)
}

func opublish(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		nid := r.FormValue("selector")
		idbook := r.FormValue("idbook")
		insForm, err := db.Prepare("UPDATE bookmark SET publish=? WHERE idbookmark=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(nid, idbook)
	}

	defer db.Close()
	http.Redirect(w, r, "/lihatpub", 301)
}

func opublishfile(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		nid := r.FormValue("selector")
		idfile := r.FormValue("idfile")
		insForm, err := db.Prepare("UPDATE file SET publish=? WHERE idFile=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(nid, idfile)
	}

	defer db.Close()
	http.Redirect(w, r, "/lihatpubfile", 301)
}

func lihatfav(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nId := suserid
	idSelector := 1
	selDB, err := db.Query("SELECT kategori.idKategori, kategori.nmKategori, bookmark.idBookmark, bookmark.idKategori, bookmark.idUser, bookmark.judul, bookmark.link, bookmark.waktuPost, bookmark.status, bookmark.publish FROM kategori, bookmark WHERE kategori.idKategori = bookmark.idKategori AND bookmark.idUser=? AND bookmark.status=? ORDER BY idBookmark DESC", nId, idSelector)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinBook
	for selDB.Next() {
		var item = JoinBook{}
		err = selDB.Scan(&item.IdKategoriKat, &item.NmKategoriKat, &item.IdBookmarkBook, &item.IdKategoriKat, &item.IdUserBook, &item.JudulBook, &item.LinkBook, &item.TglBuatBook, &item.StatusBook, &item.PublishBook)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "lihatfav", result)
	defer db.Close()
}

func lihatfavfile(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nId := suserid
	idSelector := 1
	selDB, err := db.Query("SELECT idFile, idUser, namaFile, file, tipefile, waktuPost, status, publish FROM file WHERE idUser=? AND status=? ORDER BY idFile DESC", nId, idSelector)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []Berkas
	for selDB.Next() {
		var item = Berkas{}
		err = selDB.Scan(&item.IdBerkas, &item.IdUser, &item.NamaBerkas, &item.Berkas, &item.Tipefile, &item.TglBuat, &item.Status, &item.Publish)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "lihatfavfile", result)
	defer db.Close()
}

func lihattopik(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	selDB, err := db.Query("SELECT topik.idTopik, topik.idUser, topik.topik, topik.tglPost, users.idUser, users.nmDepan, users.nmBelakang, users.foto FROM topik, users WHERE topik.idUser = users.idUser ORDER BY topik.idTopik DESC")
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinTopik
	for selDB.Next() {
		var item = JoinTopik{}
		err = selDB.Scan(&item.IdTopikTopik, &item.IdUserTopik, &item.TopikTopik, &item.TglPostTopik, &item.IdUserUsers, &item.NmDepanUsers, &item.NmBelakangUsers, &item.FotoUsers)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "lihattopik", result)
	defer db.Close()
}

func laporanTopik(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	selDB, err := db.Query("SELECT report.idReport, report.idUser, report.nmDepan, report.idBook, report.alasanPelaporan, topik.idTopik, topik.idUser, topik.topik, topik.tglPost, users.idUser, users.nmDepan, users.foto FROM report, topik, users WHERE report.idTopik = topik.idTopik AND topik.idUser = users.idUser ORDER BY report.idReport DESC")
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinTopikRe
	for selDB.Next() {
		var item = JoinTopikRe{}
		err = selDB.Scan(&item.IdReport, &item.IdUserRe, &item.NmDepanRe, &item.IdTopikRe, &item.IndikatorRe, &item.IdTopik, &item.IdUser, &item.Topik, &item.TglPost, &item.IdUserUser, &item.NmDepanUser, &item.Avatar)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "laporanTopik", result)
	defer db.Close()
}

func laporanKomentar(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	selDB, err := db.Query("SELECT report.idReport, report.idUser, report.nmDepan, report.idBook, report.alasanPelaporan, komentar.idKomentar, komentar.idTopik, komentar.IdUser, komentar.komentar, komentar.tglPost, users.idUser, users.nmDepan, users.foto FROM report, komentar, users WHERE report.idKomentar = komentar.idKomentar AND komentar.idUser = users.idUser ORDER BY report.idReport DESC")
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinKomentarRe
	for selDB.Next() {
		var item = JoinKomentarRe{}
		err = selDB.Scan(&item.IdReport, &item.IdUserRe, &item.NmDepanRe, &item.IdKomentarRe, &item.IndikatorRe, &item.IdKomentarKomen, &item.IdTopikKomen, &item.IdUserKomen, &item.KomentarKomen, &item.TglPostKomen, &item.IdUserUsers, &item.NmDepanUsers, &item.FotoUsers)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "laporanKomentar", result)
	defer db.Close()
}

func lihattopikAdmin(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	selDB, err := db.Query("SELECT topik.idTopik, topik.idUser, topik.topik, topik.tglPost, users.idUser, users.nmDepan, users.nmBelakang, users.foto FROM topik, users WHERE topik.idUser = users.idUser ORDER BY topik.idTopik DESC")
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinTopik
	for selDB.Next() {
		var item = JoinTopik{}
		err = selDB.Scan(&item.IdTopikTopik, &item.IdUserTopik, &item.TopikTopik, &item.TglPostTopik, &item.IdUserUsers, &item.NmDepanUsers, &item.NmBelakangUsers, &item.FotoUsers)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "lihattopikAdmin", result)
	defer db.Close()
}

func daftartopik(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	nid := suserid
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM topik WHERE idUser=? ORDER BY idTopik DESC", nid)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinTopik
	for selDB.Next() {
		var item = JoinTopik{}
		err = selDB.Scan(&item.IdTopikTopik, &item.IdUserTopik, &item.TopikTopik, &item.TglPostTopik, &item.NmDepanUsers, &item.FotoUsers)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "daftartopik", result)
	defer db.Close()
}

func lihatkomentar(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nid := r.URL.Query().Get("IdTopik")
	selDB, err := db.Query("SELECT komentar.idKomentar, komentar.idTopik, komentar.IdUser, komentar.komentar, komentar.tglPost, topik.idTopik, topik.idUser, topik.topik, topik.tglPost, topik.nmDepan, topik.avatar, users.idUser, users.nmDepan, users.foto FROM komentar, topik, users WHERE komentar.idTopik = topik.idTopik AND komentar.idUser = users.idUser AND topik.idTopik=? ORDER BY komentar.idKomentar ASC", nid)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinKomentar
	for selDB.Next() {
		var item = JoinKomentar{}
		err = selDB.Scan(&item.IdKomentarKomen, &item.IdTopikKomen, &item.IdUserKomen, &item.KomentarKomen, &item.TglPostKomen, &item.IdTopikTopik, &item.IdUserTopik, &item.TopikTopik, &item.TglPostTopik, &item.NmDepanTopik, &item.AvatarTopik, &item.IdUserUsers, &item.NmDepanUsers, &item.FotoUsers)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	selDB, err = db.Query("SELECT * FROM topik where idTopik=?", nid)
	if err != nil {
		panic(err.Error())
	}
	tpk := JoinTopik{}
	for selDB.Next() {
		var idTopik, idUser int
		var topik, tglPost, nmDepan, avatar string
		err = selDB.Scan(&idTopik, &idUser, &topik, &tglPost, &nmDepan, &avatar)
		if err != nil {
			panic(err.Error())
		}
		tpk.IdTopikTopik = idTopik
		tpk.IdUserTopik = idUser
		tpk.TopikTopik = topik
		tpk.TglPostTopik = tglPost
		tpk.NmDepanUsers = nmDepan
		tpk.FotoUsers = avatar
	}
	var komen = Komen{}
	komen.TopikTopik = tpk.TopikTopik
	komen.IdTopikTopik = tpk.IdTopikTopik
	komen.NmDepanUsers = tpk.NmDepanUsers
	komen.TglPostTopik = tpk.TglPostTopik
	komen.FotoUsers = tpk.FotoUsers
	komen.IdUserTopik = tpk.IdUserTopik
	komen.Komentar = result
	fmt.Println()
	fmt.Println(komen.Komentar)
	tmpl.ExecuteTemplate(w, "lihatkomentar", komen)
	defer db.Close()
}

func lihatkomentarAdmin(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nid := r.URL.Query().Get("IdTopik")
	selDB, err := db.Query("SELECT komentar.idKomentar, komentar.idTopik, komentar.IdUser, komentar.komentar, komentar.tglPost, topik.idTopik, topik.idUser, topik.topik, topik.tglPost, topik.nmDepan, topik.avatar, users.idUser, users.nmDepan, users.foto FROM komentar, topik, users WHERE komentar.idTopik = topik.idTopik AND komentar.idUser = users.idUser AND topik.idTopik=? ORDER BY komentar.idKomentar ASC", nid)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinKomentar
	for selDB.Next() {
		var item = JoinKomentar{}
		err = selDB.Scan(&item.IdKomentarKomen, &item.IdTopikKomen, &item.IdUserKomen, &item.KomentarKomen, &item.TglPostKomen, &item.IdTopikTopik, &item.IdUserTopik, &item.TopikTopik, &item.TglPostTopik, &item.NmDepanTopik, &item.AvatarTopik, &item.IdUserUsers, &item.NmDepanUsers, &item.FotoUsers)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	selDB, err = db.Query("SELECT * FROM topik where idTopik=?", nid)
	if err != nil {
		panic(err.Error())
	}
	tpk := JoinTopik{}
	for selDB.Next() {
		var idTopik, idUser int
		var topik, tglPost, nmDepan, avatar string
		err = selDB.Scan(&idTopik, &idUser, &topik, &tglPost, &nmDepan, &avatar)
		if err != nil {
			panic(err.Error())
		}
		tpk.IdTopikTopik = idTopik
		tpk.IdUserTopik = idUser
		tpk.TopikTopik = topik
		tpk.TglPostTopik = tglPost
		tpk.NmDepanUsers = nmDepan
		tpk.FotoUsers = avatar
	}
	var komen = Komen{}
	komen.TopikTopik = tpk.TopikTopik
	komen.IdTopikTopik = tpk.IdTopikTopik
	komen.NmDepanUsers = tpk.NmDepanUsers
	komen.TglPostTopik = tpk.TglPostTopik
	komen.FotoUsers = tpk.FotoUsers
	komen.IdUserTopik = tpk.IdUserTopik
	komen.Komentar = result
	fmt.Println()
	fmt.Println(komen.Komentar)
	tmpl.ExecuteTemplate(w, "lihatkomentarAdmin", komen)
	defer db.Close()
}

func afav(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		nid := r.FormValue("selector")
		idbook := r.FormValue("idbook")
		insForm, err := db.Prepare("UPDATE bookmark SET status=? WHERE idbookmark=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(nid, idbook)
	}

	defer db.Close()
	http.Redirect(w, r, "/lihatfav", 301)
}

func reportuserAdmin(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		sel := r.FormValue("sel")
		id := r.FormValue("IdUser")
		insForm, err := db.Prepare("UPDATE users SET role=? WHERE iduser=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(sel, id)
	}

	defer db.Close()
	http.Redirect(w, r, "/daftaruserAdmin", 301)
}

func pulihkanAkun(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		sel := r.FormValue("sel")
		id := r.FormValue("IdUser")
		insForm, err := db.Prepare("UPDATE users SET role=? WHERE iduser=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(sel, id)
	}

	defer db.Close()
	http.Redirect(w, r, "/daftaruserNonAktif", 301)
}

func afavfile(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		nid := r.FormValue("selector")
		idfile := r.FormValue("idfile")
		insForm, err := db.Prepare("UPDATE file SET status=? WHERE idFile=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(nid, idfile)
	}

	defer db.Close()
	http.Redirect(w, r, "/lihatfavfile", 301)
}

func ofav(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		nid := r.FormValue("selector")
		idbook := r.FormValue("idbook")
		insForm, err := db.Prepare("UPDATE bookmark SET status=? WHERE idbookmark=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(nid, idbook)
	}

	defer db.Close()
	http.Redirect(w, r, "/lihatfav", 301)
}

func ofavfile(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		nid := r.FormValue("selector")
		idfile := r.FormValue("idfile")
		insForm, err := db.Prepare("UPDATE file SET status=? WHERE idFile=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(nid, idfile)
	}

	defer db.Close()
	http.Redirect(w, r, "/lihatfavfile", 301)
}

func ubahprofil(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nId := suserid
	selDB, err := db.Query("SELECT idUser, nmDepan, nmBelakang, email, username FROM users WHERE idUser=?", nId)
	if err != nil {
		panic(err.Error())
	}
	usr := Users{}
	for selDB.Next() {
		var idUser int
		var nmDepan, nmBelakang, email, username string
		err = selDB.Scan(&idUser, &nmDepan, &nmBelakang, &email, &username)
		if err != nil {
			panic(err.Error())
		}
		usr.IdUser = idUser
		usr.NmDepan = nmDepan
		usr.NmBelakang = nmBelakang
		usr.Email = email
		usr.Username = username
	}
	tmpl.ExecuteTemplate(w, "ubahprofil", usr)
	defer db.Close()
}

func doprofil(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	if r.Method == "POST" {
		nmdepan := r.FormValue("nmdepan")
		nmbelakang := r.FormValue("nmbelakang")
		email := r.FormValue("email")
		username := r.FormValue("username")
		nid := suserid
		insForm, err := db.Prepare("UPDATE users SET nmDepan=?, nmBelakang=?, email=?, username=? WHERE idUser=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(nmdepan, nmbelakang, email, username, nid)
	}

	defer db.Close()
	http.Redirect(w, r, "/profil", 301)
}

func dokat(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	if r.Method == "POST" {
		nmKategori := r.FormValue("nmkategori")
		nid := r.FormValue("nid")
		insForm, err := db.Prepare("UPDATE kategori SET nmKategori=? WHERE idKategori=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(nmKategori, nid)
	}
	defer db.Close()
	http.Redirect(w, r, "/lihatkat", 301)
}

func dobook(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		judul := r.FormValue("judul")
		link := r.FormValue("link")
		nid := r.FormValue("IdBook")
		insForm, err := db.Prepare("UPDATE bookmark SET judul=?, link=? WHERE idbookmark=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(judul, link, nid)
	}

	defer db.Close()
	http.Redirect(w, r, "/lihatbook", 301)
}

func ubahfoto(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	if r.Method == "POST" {
		file, header, _ := r.FormFile("avatar")
		nid := suserid
		var statement, err = db.Prepare("UPDATE users SET foto=? WHERE idUser=?")
		CheckError(err)
		statement.Exec(header.Filename, strings.ToLower(nid))
		out, _ := os.Create("./avatar/" + header.Filename)
		_, _ = io.Copy(out, file)
		file.Close()
		out.Close()
	}
	defer db.Close()
	http.Redirect(w, r, "/profil", 301)
}

func ubahfile(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	if r.Method == "POST" {
		namafile := r.FormValue("namafile")
		file, header, _ := r.FormFile("file")
		tipefile := r.FormValue("tipefile")
		id := r.FormValue("IdBerkas")
		var statement, err = db.Prepare("UPDATE file SET namaFile=?, file=?, tipefile=? WHERE idFile=?")
		CheckError(err)
		statement.Exec(namafile, header.Filename, tipefile, strings.ToLower(id))
		out, _ := os.Create("./upload_file/" + header.Filename)
		_, _ = io.Copy(out, file)
		file.Close()
		out.Close()
	}
	defer db.Close()
	http.Redirect(w, r, "/lihatberkas", 301)
}

func listbook(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nId := suserid
	NmKat := r.FormValue("kategori")
	selDB, err := db.Query("SELECT kategori.idKategori, kategori.nmKategori, bookmark.idBookmark, bookmark.idKategori, bookmark.idUser, bookmark.judul, bookmark.link, bookmark.waktuPost, bookmark.status, bookmark.publish FROM kategori, bookmark WHERE kategori.idKategori = bookmark.idKategori AND bookmark.idUser=? AND bookmark.idKategori=? ORDER BY idBookmark DESC", nId, NmKat)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinBook
	for selDB.Next() {
		var item = JoinBook{}
		err = selDB.Scan(&item.IdKategoriKat, &item.NmKategoriKat, &item.IdBookmarkBook, &item.IdKategoriKat, &item.IdUserBook, &item.JudulBook, &item.LinkBook, &item.TglBuatBook, &item.StatusBook, &item.PublishBook)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	selDB, err = db.Query("SELECT * FROM kategori where idUser=?", suserid)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var out []Kategori
	for selDB.Next() {
		var item = Kategori{}
		err = selDB.Scan(&item.IdKategori, &item.IdUser, &item.NmKategori, &item.TglBuat)
		if err != nil {
			panic(err.Error())
		}
		out = append(out, item)
	}
	var hasil = Book{}
	hasil.Bookmark = result
	hasil.Kategori = out
	tmpl.ExecuteTemplate(w, "listbook", hasil)
}

func sortkategori(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nId := suserid
	NmKat := r.FormValue("kategori")
	selDB, err := db.Query("SELECT kategori.idKategori, kategori.nmKategori, bookmark.idBookmark, bookmark.idKategori, bookmark.idUser, bookmark.judul, bookmark.link, bookmark.waktuPost, bookmark.status, bookmark.publish FROM kategori, bookmark WHERE kategori.idKategori = bookmark.idKategori AND bookmark.idUser=? AND bookmark.idKategori=? ORDER BY idBookmark DESC", nId, NmKat)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinBook
	for selDB.Next() {
		var item = JoinBook{}
		err = selDB.Scan(&item.IdKategoriKat, &item.NmKategoriKat, &item.IdBookmarkBook, &item.IdKategoriKat, &item.IdUserBook, &item.JudulBook, &item.LinkBook, &item.TglBuatBook, &item.StatusBook, &item.PublishBook)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	selDB, err = db.Query("SELECT * FROM kategori where idUser=?", suserid)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var out []Kategori
	for selDB.Next() {
		var item = Kategori{}
		err = selDB.Scan(&item.IdKategori, &item.IdUser, &item.NmKategori, &item.TglBuat)
		if err != nil {
			panic(err.Error())
		}
		out = append(out, item)
	}
	var hasil = Book{}
	hasil.Bookmark = result
	hasil.Kategori = out
	tmpl.ExecuteTemplate(w, "sortkategori", hasil)
}

func listbookAdmin(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	NmKat := r.URL.Query().Get("NmKat")
	selDB, err := db.Query("SELECT kategori.idKategori, kategori.idUser, kategori.nmKategori, bookmark.idBookmark, bookmark.idKategori, bookmark.idUser, bookmark.judul, bookmark.link, bookmark.waktuPost, bookmark.publish, bookmark.updatePublish, users.idUser, users.nmDepan , users.Foto FROM kategori, bookmark, users WHERE kategori.idKategori = bookmark.idKategori AND bookmark.idUser = users.idUser AND bookmark.idKategori=? ORDER BY bookmark.updatePublish DESC", NmKat)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinBookTel
	for selDB.Next() {
		var item = JoinBookTel{}
		err = selDB.Scan(&item.IdKategoriKat, &item.IdUserKat, &item.NmKategoriKat, &item.IdBookmarkBook, &item.IdKategoriBook, &item.IdUserBook, &item.JudulBook, &item.LinkBook, &item.TglBuatBook, &item.PublishBook, &item.UpdateWaktu, &item.IdUserUser, &item.NmDepanUser, &item.Avatar)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "listbookAdmin", result)
}

func ubahbook(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nId := r.URL.Query().Get("IdBook")
	selDB, err := db.Query("SELECT idBookmark, idKategori, idUser, judul, link, waktuPost, status, publish FROM bookmark WHERE idBookmark=?", nId)
	if err != nil {
		panic(err.Error())
	}
	bk := Bookmark{}
	for selDB.Next() {
		var idBookmark, idKategori, idUser, status, publish int
		var judul, link, tglBuat string
		err = selDB.Scan(&idBookmark, &idKategori, &idUser, &judul, &link, &tglBuat, &status, &publish)
		if err != nil {
			panic(err.Error())
		}
		bk.IdBookmark = idBookmark
		bk.IdKategori = idKategori
		bk.IdUser = idUser
		bk.Judul = judul
		bk.Link = link
		bk.TglBuat = tglBuat
		bk.Status = status
		bk.Publish = publish
	}
	tmpl.ExecuteTemplate(w, "ubahbook", bk)
	defer db.Close()
}

func ubahkat(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nId := r.URL.Query().Get("IdKat")
	selDB, err := db.Query("SELECT idKategori, idUser, nmKategori FROM kategori WHERE idKategori=?", nId)
	if err != nil {
		panic(err.Error())
	}
	kt := Kategori{}
	for selDB.Next() {
		var idKategori, idUser int
		var nmKategori string
		err = selDB.Scan(&idKategori, &idUser, &nmKategori)
		if err != nil {
			panic(err.Error())
		}
		kt.IdKategori = idKategori
		kt.IdUser = idUser
		kt.NmKategori = nmKategori
	}
	tmpl.ExecuteTemplate(w, "ubahkat", kt)
	defer db.Close()
}

func hapusbook(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	id := r.URL.Query().Get("IdBook")
	delForm, err := db.Prepare("DELETE FROM bookmark WHERE idBookmark=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	defer db.Close()
	http.Redirect(w, r, "/lihatbook", 301)
}

func hapususerAdmin(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	id := r.URL.Query().Get("IdUser")
	delForm, err := db.Prepare("DELETE FROM users WHERE idUser=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	defer db.Close()
	http.Redirect(w, r, "/daftaruserAdmin", 301)
}

func hapusbookAdmin(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	id := r.URL.Query().Get("IdBook")
	delForm, err := db.Prepare("DELETE FROM bookmark WHERE idBookmark=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	defer db.Close()
	http.Redirect(w, r, "/lihatbookAdmin", 301)
}

func hapustopik(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	id := r.URL.Query().Get("IdTopik")
	delForm, err := db.Prepare("DELETE FROM topik WHERE idTopik=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	defer db.Close()
	http.Redirect(w, r, "/daftartopik", 301)
}

func hapuskomentarAdmin(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	idKomen := r.FormValue("IdKomen")
	idTopik := r.FormValue("IdTopik")
	delForm, err := db.Prepare("DELETE FROM komentar WHERE idKomentar=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(idKomen)
	defer db.Close()
	http.Redirect(w, r, "/lihatkomentarAdmin?IdTopik="+idTopik, 301)
}

func hapustopikAdmin(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	id := r.URL.Query().Get("IdTopik")
	delForm, err := db.Prepare("DELETE FROM topik WHERE idTopik=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	defer db.Close()
	http.Redirect(w, r, "/lihattopikAdmin", 301)
}

func hapussimpanbook(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		id := r.FormValue("IdSimpan")
		delForm, err := db.Prepare("DELETE FROM tersimpan WHERE idSimpan=?")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(id)
		defer db.Close()
	}
	http.Redirect(w, r, "/tersimpanbook", 301)
}

func hapusreportbook(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	rep := r.FormValue("IdReport")
	if r.Method == "POST" {
		id := r.FormValue("IdBook")
		delForm, err := db.Prepare("DELETE FROM bookmark WHERE idBookmark=?")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(id)
		defer db.Close()
	}
	http.Redirect(w, r, "/hapusreportbookdirect?IdReport="+rep, 301)
}

func hapusreportkomentar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	rep := r.FormValue("IdReport")
	if r.Method == "POST" {
		id := r.FormValue("IdKomen")
		delForm, err := db.Prepare("DELETE FROM komentar WHERE idKomentar=?")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(id)
		defer db.Close()
	}
	http.Redirect(w, r, "/hapusreportkomentardirect?IdReport="+rep, 301)
}

func hapusreporttopik(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	rep := r.FormValue("IdReport")
	if r.Method == "POST" {
		id := r.FormValue("IdTopik")
		delForm, err := db.Prepare("DELETE FROM topik WHERE idTopik=?")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(id)
		defer db.Close()
	}
	http.Redirect(w, r, "/hapusreporttopikdirect?IdReport="+rep, 301)
}

func hapusreportakun(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	rep := r.FormValue("IdReport")
	if r.Method == "POST" {
		id := r.FormValue("IdUser")
		delForm, err := db.Prepare("DELETE FROM users WHERE idUser=?")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(id)
		defer db.Close()
	}
	http.Redirect(w, r, "/hapusreportakundirect?IdReport="+rep, 301)
}

func hapusreportfile(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	rep := r.FormValue("IdReport")
	if r.Method == "POST" {
		id := r.FormValue("IdFile")
		delForm, err := db.Prepare("DELETE FROM file WHERE idFile=?")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(id)
		defer db.Close()
	}
	http.Redirect(w, r, "/hapusreportfiledirect?IdReport="+rep, 301)
}

func hapuskirimanbook(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		id := r.FormValue("IdKirim")
		delForm, err := db.Prepare("DELETE FROM kiriman WHERE idKirim=?")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(id)
		defer db.Close()
	}
	http.Redirect(w, r, "/kirimanbook", 301)
}

func hapuskirimanfile(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		id := r.FormValue("IdKirim")
		delForm, err := db.Prepare("DELETE FROM kiriman WHERE idKirim=?")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(id)
		defer db.Close()
	}
	http.Redirect(w, r, "/kirimanfile", 301)
}

func hapussimpanfile(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		id := r.FormValue("IdSimpan")
		delForm, err := db.Prepare("DELETE FROM tersimpan WHERE idSimpan=?")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(id)
		defer db.Close()
	}
	http.Redirect(w, r, "/tersimpanfile", 301)
}

func hapusfile(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	id := r.URL.Query().Get("IdBerkas")
	delForm, err := db.Prepare("DELETE FROM file WHERE idFile=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	defer db.Close()
	http.Redirect(w, r, "/lihatberkas", 301)
}

func hapusreportbookdirect(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	id := r.URL.Query().Get("IdReport")
	delForm, err := db.Prepare("DELETE FROM report WHERE idReport=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	defer db.Close()
	http.Redirect(w, r, "/laporanBookmark", 301)
}

func hapusreportkomentardirect(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	id := r.URL.Query().Get("IdReport")
	delForm, err := db.Prepare("DELETE FROM report WHERE idReport=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	defer db.Close()
	http.Redirect(w, r, "/laporanKomentar", 301)
}

func hapusreporttopikdirect(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	id := r.URL.Query().Get("IdReport")
	delForm, err := db.Prepare("DELETE FROM report WHERE idReport=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	defer db.Close()
	http.Redirect(w, r, "/laporanTopik", 301)
}

func hapusreportakundirect(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	id := r.URL.Query().Get("IdReport")
	delForm, err := db.Prepare("DELETE FROM report WHERE idReport=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	defer db.Close()
	http.Redirect(w, r, "/laporanAkun", 301)
}

func hapusreportfiledirect(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	id := r.URL.Query().Get("IdReport")
	delForm, err := db.Prepare("DELETE FROM report WHERE idReport=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	defer db.Close()
	http.Redirect(w, r, "/laporanFile", 301)
}

func hapusfileAdmin(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	id := r.URL.Query().Get("IdFile")
	delForm, err := db.Prepare("DELETE FROM file WHERE idFile=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	defer db.Close()
	http.Redirect(w, r, "/lihatfileAdmin", 301)
}

func hapuskat(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	id := r.URL.Query().Get("IdKat")
	delForm, err := db.Prepare("DELETE FROM kategori WHERE idKategori=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	defer db.Close()
	http.Redirect(w, r, "/lihatkat", 301)
}

func hapuskatAdmin(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	id := r.URL.Query().Get("IdKat")
	delForm, err := db.Prepare("DELETE FROM kategori WHERE idKategori=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	defer db.Close()
	http.Redirect(w, r, "/lihatkatAdmin", 301)
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func addberkas(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	if r.Method == "POST" {
		namafile := r.FormValue("namafile")
		file, header, _ := r.FormFile("file")
		tipefile := r.FormValue("tipefile")
		var statement, err = db.Prepare("INSERT INTO file(idUser,namaFile,file,tipefile) VALUES(?,?,?,?)")
		CheckError(err)
		statement.Exec(strings.ToLower(suserid), namafile, header.Filename, tipefile)
		out, _ := os.Create("./upload_file/" + header.Filename)
		_, _ = io.Copy(out, file)
		file.Close()
		out.Close()
	}
	http.Redirect(w, r, "/lihatberkas", 301)
}

func telusuribook(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nid := 1
	selDB, err := db.Query("SELECT kategori.idKategori, kategori.idUser, kategori.nmKategori, bookmark.idBookmark, bookmark.idKategori, bookmark.idUser, bookmark.judul, bookmark.link, bookmark.waktuPost, bookmark.publish, bookmark.updatePublish, users.idUser, users.nmDepan , users.Foto FROM kategori, bookmark, users WHERE kategori.idKategori = bookmark.idKategori AND bookmark.idUser = users.idUser AND bookmark.publish=? ORDER BY bookmark.updatePublish DESC", nid)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinBookTel
	for selDB.Next() {
		var item = JoinBookTel{}
		err = selDB.Scan(&item.IdKategoriKat, &item.IdUserKat, &item.NmKategoriKat, &item.IdBookmarkBook, &item.IdKategoriBook, &item.IdUserBook, &item.JudulBook, &item.LinkBook, &item.TglBuatBook, &item.PublishBook, &item.UpdateWaktu, &item.IdUserUser, &item.NmDepanUser, &item.Avatar)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	role := 1
	selDB, err = db.Query("SELECT kategori.idKategori, kategori.nmKategori, bookmark.idKategori, bookmark.publish FROM kategori, bookmark WHERE kategori.idKategori = bookmark.idKategori AND bookmark.publish=?", role)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var out []KategoriTel
	for selDB.Next() {
		var item = KategoriTel{}
		err = selDB.Scan(&item.IdKategori, &item.NmKategori, &item.IdKategoriBook, &item.PublishBook)
		if err != nil {
			panic(err.Error())
		}
		out = append(out, item)
	}
	var hasil = BookTel{}
	hasil.Bookmark = result
	hasil.Kategori = out
	tmpl.ExecuteTemplate(w, "telusuribook", hasil)
	defer db.Close()
}

func telusuribooksort(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nid := 1
	Nm := r.FormValue("kategori")
	selDB, err := db.Query("SELECT kategori.idKategori, kategori.idUser, kategori.nmKategori, bookmark.idBookmark, bookmark.idKategori, bookmark.idUser, bookmark.judul, bookmark.link, bookmark.waktuPost, bookmark.publish, bookmark.updatePublish, users.idUser, users.nmDepan , users.Foto FROM kategori, bookmark, users WHERE kategori.idKategori = bookmark.idKategori AND bookmark.idUser = users.idUser AND bookmark.publish=? AND bookmark.idKategori=? ORDER BY bookmark.updatePublish DESC", nid, Nm)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinBookTel
	for selDB.Next() {
		var item = JoinBookTel{}
		err = selDB.Scan(&item.IdKategoriKat, &item.IdUserKat, &item.NmKategoriKat, &item.IdBookmarkBook, &item.IdKategoriBook, &item.IdUserBook, &item.JudulBook, &item.LinkBook, &item.TglBuatBook, &item.PublishBook, &item.UpdateWaktu, &item.IdUserUser, &item.NmDepanUser, &item.Avatar)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	role := 1
	selDB, err = db.Query("SELECT kategori.idKategori, kategori.nmKategori, bookmark.idKategori, bookmark.publish FROM kategori, bookmark WHERE kategori.idKategori = bookmark.idKategori AND bookmark.publish=?", role)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var out []KategoriTel
	for selDB.Next() {
		var item = KategoriTel{}
		err = selDB.Scan(&item.IdKategori, &item.NmKategori, &item.IdKategoriBook, &item.PublishBook)
		if err != nil {
			panic(err.Error())
		}
		out = append(out, item)
	}
	var hasil = BookTel{}
	hasil.Bookmark = result
	hasil.Kategori = out
	tmpl.ExecuteTemplate(w, "telusuribooksort", hasil)
	defer db.Close()
}

func lihatbookAdmin(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	selDB, err := db.Query("SELECT kategori.idKategori, kategori.idUser, kategori.nmKategori, bookmark.idBookmark, bookmark.idKategori, bookmark.idUser, bookmark.judul, bookmark.link, bookmark.waktuPost, bookmark.publish, bookmark.updatePublish, users.idUser, users.nmDepan , users.Foto FROM kategori, bookmark, users WHERE kategori.idKategori = bookmark.idKategori AND bookmark.idUser = users.idUser ORDER BY bookmark.updatePublish DESC")
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinBookTel
	for selDB.Next() {
		var item = JoinBookTel{}
		err = selDB.Scan(&item.IdKategoriKat, &item.IdUserKat, &item.NmKategoriKat, &item.IdBookmarkBook, &item.IdKategoriBook, &item.IdUserBook, &item.JudulBook, &item.LinkBook, &item.TglBuatBook, &item.PublishBook, &item.UpdateWaktu, &item.IdUserUser, &item.NmDepanUser, &item.Avatar)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "lihatbookAdmin", result)
	defer db.Close()
}

func tersimpanbook(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	selDB, err := db.Query("SELECT tersimpan.idSimpan, tersimpan.idUser, tersimpan.idUserPengirim, tersimpan.idBookmark, kategori.idKategori, kategori.nmKategori, bookmark.idBookmark, bookmark.idKategori, bookmark.idUser, bookmark.judul, bookmark.link, bookmark.waktuPost, users.idUser, users.nmDepan FROM tersimpan, kategori, bookmark, users WHERE tersimpan.idUserPengirim = users.idUser AND tersimpan.idBookmark = bookmark.idBookmark AND kategori.idKategori = bookmark.idKategori AND tersimpan.idUser=? ORDER BY tersimpan.idSimpan DESC", suserid)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinTersimpanBook
	for selDB.Next() {
		var item = JoinTersimpanBook{}
		err = selDB.Scan(&item.IdSimpan, &item.IdUserSimpan, &item.IdUserPengirim, &item.IdBookmarkSimpan, &item.IdKategoriKat, &item.NmKategoriKat, &item.IdBookmarkBook, &item.IdKategoriBook, &item.IdUserBook, &item.JudulBook, &item.LinkBook, &item.TglBuatBook, &item.IdUserUser, &item.NmDepanUser)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "tersimpanbook", result)
	defer db.Close()
}

func laporanBookmark(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	selDB, err := db.Query("SELECT report.idReport, report.idUser, report.nmDepan, report.idBook, report.alasanPelaporan, kategori.idKategori, kategori.nmKategori, bookmark.idBookmark, bookmark.idKategori, bookmark.idUser, bookmark.judul, bookmark.link, bookmark.waktuPost, users.idUser, users.nmDepan FROM report, kategori, bookmark, users WHERE bookmark.idUser = users.idUser AND report.idBook = bookmark.idBookmark AND kategori.idKategori = bookmark.idKategori ORDER BY report.idReport DESC")
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinBookRe
	for selDB.Next() {
		var item = JoinBookRe{}
		err = selDB.Scan(&item.IdReport, &item.IdUserRe, &item.NmDepanRe, &item.IdBookRe, &item.IndikatorRe, &item.IdKategoriKat, &item.NmKategoriKat, &item.IdBookmarkBook, &item.IdKategoriBook, &item.IdUserBook, &item.JudulBook, &item.LinkBook, &item.TglBuatBook, &item.IdUserUser, &item.NmDepanUser)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "laporanBookmark", result)
	defer db.Close()
}

func laporanAkun(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	role0 := 0
	selDB, err := db.Query("SELECT report.idReport, report.idUser, report.nmDepan, report.idUserReport, report.alasanPelaporan, users.idUser, users.nmDepan, users.nmBelakang, users.tglGabung, users.foto, users.role FROM report, users WHERE report.idUserReport = users.idUser AND users.role=? ORDER BY report.idReport DESC", role0)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinUserRe
	for selDB.Next() {
		var item = JoinUserRe{}
		err = selDB.Scan(&item.IdReport, &item.IdUserRe, &item.NmDepanRe, &item.IdUserReportRe, &item.IndikatorRe, &item.IdUser, &item.NmDepan, &item.NmBelakang, &item.TglGabung, &item.Foto, &item.Role)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "laporanAkun", result)
	defer db.Close()
}

func laporanFile(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	selDB, err := db.Query("SELECT report.idReport, report.idUser, report.nmDepan, report.idFile, report.alasanPelaporan, file.idFile, file.idUser, file.namaFile, file.file, file.tipefile, file.waktuPost, file.status, file.publish, users.idUser, users.nmDepan FROM report, file, users WHERE file.idUser = users.idUser AND report.idfile = file.idFile ORDER BY report.idReport DESC")
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinFileRe
	for selDB.Next() {
		var item = JoinFileRe{}
		err = selDB.Scan(&item.IdReport, &item.IdUserRe, &item.NmDepanRe, &item.IdFileRe, &item.IndikatorRe, &item.IdFileFile, &item.IdUserFile, &item.NamaFileFile, &item.FileFile, &item.TipefileFile, &item.TglBuatFile, &item.StatusFile, &item.PublishFile, &item.IdUserUser, &item.NmDepanUser)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "laporanFile", result)
	defer db.Close()
}

func kirimanbook(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	selDB, err := db.Query("SELECT kiriman.idKirim, kiriman.nmDepanPengirim, kiriman.idPenerima, kiriman.idBook, kategori.idKategori, kategori.nmKategori, bookmark.idBookmark, bookmark.idKategori, bookmark.idUser, bookmark.judul, bookmark.link, bookmark.waktuPost, users.idUser, users.nmDepan FROM kiriman, kategori, bookmark, users WHERE bookmark.idUser = users.idUser AND kiriman.idBook = bookmark.idBookmark AND kategori.idKategori = bookmark.idKategori AND kiriman.idPenerima=? ORDER BY kiriman.idKirim DESC", suserid)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinKirimanBook
	for selDB.Next() {
		var item = JoinKirimanBook{}
		err = selDB.Scan(&item.IdKirim, &item.NmDepanPengirim, &item.IdPenerima, &item.IdBook, &item.IdKategoriKat, &item.NmKategoriKat, &item.IdBook, &item.IdKategoriBook, &item.IdUserBook, &item.JudulBook, &item.LinkBook, &item.TglBuatBook, &item.IdUserUsers, &item.NmDepanUsers)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "kirimanbook", result)
	defer db.Close()
}

func kirimanfile(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	selDB, err := db.Query("SELECT kiriman.idKirim, kiriman.nmDepanPengirim, kiriman.idPenerima, kiriman.idFile, file.idFile, file.idUser, file.namaFile, file.file, file.tipeFile, file.waktuPost, users.idUser, users.nmDepan FROM kiriman, file, users WHERE file.idUser = users.idUser AND kiriman.idFile = file.idFile AND kiriman.idPenerima=? ORDER BY kiriman.idKirim DESC", suserid)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinKirimanFile
	for selDB.Next() {
		var item = JoinKirimanFile{}
		err = selDB.Scan(&item.IdKirim, &item.NmDepanPengirim, &item.IdPenerima, &item.IdFile, &item.IdFileFile, &item.IdUserFile, &item.NamaFileFile, &item.FileFile, &item.TipefileFile, &item.TglBuatFile, &item.IdUserUser, &item.NmDepanUser)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "kirimanfile", result)
	defer db.Close()
}

func tersimpanfile(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	selDB, err := db.Query("SELECT tersimpan.idSimpan, tersimpan.idUser, tersimpan.idUserPengirim, tersimpan.idFile, file.idFile, file.idUser, file.namaFile, file.file, file.tipefile, file.waktuPost, users.idUser, users.nmDepan FROM tersimpan, file, users WHERE tersimpan.idUserPengirim = users.idUser AND tersimpan.idFile = file.idFile AND tersimpan.idUser=? ORDER BY tersimpan.idSimpan DESC", suserid)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinTersimpanFile
	for selDB.Next() {
		var item = JoinTersimpanFile{}
		err = selDB.Scan(&item.IdSimpan, &item.IdUserSimpan, &item.IdUserPengirim, &item.IdFileSimpan, &item.IdBerkasBerkas, &item.IdUserBerkas, &item.NamaBerkasBerkas, &item.BerkasBerkas, &item.TipefileBerkas, &item.TglBuatBerkas, &item.IdUserUser, &item.NmDepanUser)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "tersimpanfile", result)
	defer db.Close()
}

func telusurifile(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nid := 1
	selDB, err := db.Query("SELECT file.idFile, file.idUser, file.namaFile, file.file, file.tipefile, file.waktuPost, file.publish, file.updatePublish ,users.idUser, users.nmDepan, users.foto FROM file, users WHERE file.idUser = users.idUser AND file.publish=? ORDER BY file.updatePublish DESC", nid)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinBerkas
	for selDB.Next() {
		var item = JoinBerkas{}
		err = selDB.Scan(&item.IdBerkasBerkas, &item.IdUserBerkas, &item.NamaBerkasBerkas, &item.BerkasBerkas, &item.TipefileBerkas, &item.TglBuatBerkas, &item.PublishBerkas, &item.UpdateWaktu, &item.IdUserUser, &item.NmDepanUser, &item.Avatar)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "telusurifile", result)
	defer db.Close()
}

func telusurifilesort(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	nid := 1
	tp := r.FormValue("tipefile")
	selDB, err := db.Query("SELECT file.idFile, file.idUser, file.namaFile, file.file, file.tipefile, file.waktuPost, file.publish, file.updatePublish ,users.idUser, users.nmDepan, users.foto FROM file, users WHERE file.idUser = users.idUser AND file.publish=? AND file.tipefile=? ORDER BY file.updatePublish DESC", nid, tp)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinBerkas
	for selDB.Next() {
		var item = JoinBerkas{}
		err = selDB.Scan(&item.IdBerkasBerkas, &item.IdUserBerkas, &item.NamaBerkasBerkas, &item.BerkasBerkas, &item.TipefileBerkas, &item.TglBuatBerkas, &item.PublishBerkas, &item.UpdateWaktu, &item.IdUserUser, &item.NmDepanUser, &item.Avatar)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "telusurifilesort", result)
	defer db.Close()
}

func lihatfileAdmin(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	var suserid = session.GetString("suserid")
	var data = make(map[string]string)
	if suserid == "" {
		fmt.Println("login dulu")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["suserid"] = suserid
	data["err"] = r.URL.Query().Get("err")
	db := dbConn()
	selDB, err := db.Query("SELECT file.idFile, file.idUser, file.namaFile, file.file, file.tipefile, file.waktuPost, file.publish, file.updatePublish ,users.idUser, users.nmDepan, users.foto FROM file, users WHERE file.idUser = users.idUser ORDER BY file.updatePublish DESC")
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()
	var result []JoinBerkas
	for selDB.Next() {
		var item = JoinBerkas{}
		err = selDB.Scan(&item.IdBerkasBerkas, &item.IdUserBerkas, &item.NamaBerkasBerkas, &item.BerkasBerkas, &item.TipefileBerkas, &item.TglBuatBerkas, &item.PublishBerkas, &item.UpdateWaktu, &item.IdUserUser, &item.NmDepanUser, &item.Avatar)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, item)
	}
	tmpl.ExecuteTemplate(w, "lihatfileAdmin", result)
	defer db.Close()
}

func logout(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	session.Clear()
	sessions.Destroy(w, r)
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Accel-Expires", "0")
	http.Redirect(w, r, "/", 302)
}

func main() {
	//handler halaman
	http.HandleFunc("/", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/profil", profil)
	http.HandleFunc("/lihatkat", lihatkat)
	http.HandleFunc("/tambahkat", tambahkat)
	http.HandleFunc("/tambahbook", tambahbook)
	http.HandleFunc("/addkategori", addkategori)
	http.HandleFunc("/lihatbook", lihatbook)
	http.HandleFunc("/lihatbookAdmin", lihatbookAdmin)
	http.HandleFunc("/addbookmark", addbookmark)
	http.HandleFunc("/addkomen", addkomen)
	http.HandleFunc("/addkomenAdmin", addkomenAdmin)
	http.HandleFunc("/addtopik", addtopik)
	http.HandleFunc("/afav", afav)
	http.HandleFunc("/afavfile", afavfile)
	http.HandleFunc("/lihatfav", lihatfav)
	http.HandleFunc("/lihatfavfile", lihatfavfile)
	http.HandleFunc("/ofav", ofav)
	http.HandleFunc("/ofavfile", ofavfile)
	http.HandleFunc("/ubahprofil", ubahprofil)
	http.HandleFunc("/ubahkat", ubahkat)
	http.HandleFunc("/ubahfoto", ubahfoto)
	http.HandleFunc("/doprofil", doprofil)
	http.HandleFunc("/dokat", dokat)
	http.HandleFunc("/listbook", listbook)
	http.HandleFunc("/sortkategori", sortkategori)
	http.HandleFunc("/sortfile", sortfile)
	http.HandleFunc("/listbookAdmin", listbookAdmin)
	http.HandleFunc("/lihatpub", lihatpub)
	http.HandleFunc("/lihatpubfile", lihatpubfile)
	http.HandleFunc("/lihattopik", lihattopik)
	http.HandleFunc("/lihattopikAdmin", lihattopikAdmin)
	http.HandleFunc("/lihatkomentar", lihatkomentar)
	http.HandleFunc("/lihatkomentarAdmin", lihatkomentarAdmin)
	http.HandleFunc("/lihatuser", lihatuser)
	http.HandleFunc("/lihatuserAdmin", lihatuserAdmin)
	http.HandleFunc("/opublish", opublish)
	http.HandleFunc("/opublishfile", opublishfile)
	http.HandleFunc("/addpublish", addpublish)
	http.HandleFunc("/addpublishfile", addpublishfile)
	http.HandleFunc("/simpanbook", simpanbook)
	http.HandleFunc("/simpanfile", simpanfile)
	http.HandleFunc("/ubahbook", ubahbook)
	http.HandleFunc("/dobook", dobook)
	http.HandleFunc("/ubahfile", ubahfile)
	http.HandleFunc("/hapusbook", hapusbook)
	http.HandleFunc("/hapusbookAdmin", hapusbookAdmin)
	http.HandleFunc("/hapuskirimanbook", hapuskirimanbook)
	http.HandleFunc("/hapuskirimanfile", hapuskirimanfile)
	http.HandleFunc("/lihatberkas", lihatberkas)
	http.HandleFunc("/lihatfileAdmin", lihatfileAdmin)
	http.HandleFunc("/daftartopik", daftartopik)
	http.HandleFunc("/tambahberkas", tambahberkas)
	http.HandleFunc("/hapuskat", hapuskat)
	http.HandleFunc("/hapuskatAdmin", hapuskatAdmin)
	http.HandleFunc("/hapususerAdmin", hapususerAdmin)
	http.HandleFunc("/reportuserAdmin", reportuserAdmin)
	http.HandleFunc("/pulihkanAkun", pulihkanAkun)
	http.HandleFunc("/hapusfile", hapusfile)
	http.HandleFunc("/hapusfileAdmin", hapusfileAdmin)
	http.HandleFunc("/hapustopik", hapustopik)
	http.HandleFunc("/hapustopikAdmin", hapustopikAdmin)
	http.HandleFunc("/hapuskomentarAdmin", hapuskomentarAdmin)
	http.HandleFunc("/kirimanbook", kirimanbook)
	http.HandleFunc("/kirimanfile", kirimanfile)
	http.HandleFunc("/hapussimpanbook", hapussimpanbook)
	http.HandleFunc("/hapussimpanfile", hapussimpanfile)
	http.HandleFunc("/hapusreportbook", hapusreportbook)
	http.HandleFunc("/hapusreportbookdirect", hapusreportbookdirect)
	http.HandleFunc("/hapusreportfile", hapusreportfile)
	http.HandleFunc("/hapusreportfiledirect", hapusreportfiledirect)
	http.HandleFunc("/hapusreportakun", hapusreportakun)
	http.HandleFunc("/hapusreportakundirect", hapusreportakundirect)
	http.HandleFunc("/hapusreporttopik", hapusreporttopik)
	http.HandleFunc("/hapusreporttopikdirect", hapusreporttopikdirect)
	http.HandleFunc("/hapusreportkomentar", hapusreportkomentar)
	http.HandleFunc("/hapusreportkomentardirect", hapusreportkomentardirect)
	http.HandleFunc("/addberkas", addberkas)
	http.HandleFunc("/telusuribook", telusuribook)
	http.HandleFunc("/telusuribooksort", telusuribooksort)
	http.HandleFunc("/telusurifile", telusurifile)
	http.HandleFunc("/telusurifilesort", telusurifilesort)
	http.HandleFunc("/tersimpanbook", tersimpanbook)
	http.HandleFunc("/tersimpanfile", tersimpanfile)
	http.HandleFunc("/daftaruser", daftaruser)
	http.HandleFunc("/daftaruserAdmin", daftaruserAdmin)
	http.HandleFunc("/daftaruserNonAktif", daftaruserNonAktif)
	http.HandleFunc("/tambahtopik", tambahtopik)
	http.HandleFunc("/kirimbook", kirimbook)
	http.HandleFunc("/kirimfile", kirimfile)
	http.HandleFunc("/proseskirimbook", proseskirimbook)
	http.HandleFunc("/proseskirimfile", proseskirimfile)
	http.HandleFunc("/laporkantopik", laporkantopik)
	http.HandleFunc("/laporkankomen", laporkankomen)
	http.HandleFunc("/laporkanbook", laporkanbook)
	http.HandleFunc("/laporkanfile", laporkanfile)
	http.HandleFunc("/laporkanuser", laporkanuser)
	http.HandleFunc("/lihatkatAdmin", lihatkatAdmin)
	http.HandleFunc("/laporanBookmark", laporanBookmark)
	http.HandleFunc("/laporanFile", laporanFile)
	http.HandleFunc("/laporanAkun", laporanAkun)
	http.HandleFunc("/laporanTopik", laporanTopik)
	http.HandleFunc("/laporanKomentar", laporanKomentar)

	//handler proses
	http.HandleFunc("/authlogin", authlogin)
	http.HandleFunc("/proccessreg", proccessreg)
	// http.HandleFunc("/proccesskat", proccesskat)
	// http.HandleFunc("/proccessbook", proccessbook)
	http.HandleFunc("/logout", logout)
	http.Handle("/upload_file/", http.StripPrefix("/upload_file/", http.FileServer(http.Dir("upload_file"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/icon/", http.StripPrefix("/icon/", http.FileServer(http.Dir("icon"))))
	http.Handle("/avatar/", http.StripPrefix("/avatar/", http.FileServer(http.Dir("avatar"))))
	//jalankan server
	http.ListenAndServe(":8000", nil)
}
