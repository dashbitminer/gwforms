package Salud

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"formulario/Controller"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func Consentimientopost(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	responsable := r.FormValue("responsable")
	joven := r.FormValue("joven")
	fechaautoriza := r.FormValue("fechaautoriza")
	autoriza := r.FormValue("autoriza")
	fechafirma := r.FormValue("fechafirma")

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into [GWFORMS].[dbo].[ConsentimientoMEX] (responsable, joven, fechaautoriza, autoriza, fechafirma)  OUTPUT INSERTED.ID values (@p1,@p2,@p3,@p4,@p5)")
	if err != nil {
		log.Fatal(err)
	}
	var idGenerado int
	err = stmt.QueryRow(responsable, joven, fechaautoriza, autoriza, fechafirma).Scan(&idGenerado)
	if err != nil {
		log.Fatal(err)
	}

	cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
	if err != nil {

		http.Error(w, "Error al ejecutar: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
	client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
	if err != nil {

		http.Error(w, "Error al ejecutar: "+err.Error(), http.StatusInternalServerError)
		return
	}
	file, fileHandler, err := r.FormFile("pdfConsentimiento")
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Obtiene el nombre del archivo cargado
	fileName := fileHandler.Filename

	//Obtiene la extensión del archivo cargado
	fileExt := filepath.Ext(fileName)

	fileBytes, err := io.ReadAll(file)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tempFile, err := os.CreateTemp("", "upload-*.dat")
	if err != nil {

		http.Error(w, "Error al ejecutar: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	// Escribe los bytes del archivo en el archivo temporal
	if _, err := tempFile.Write(fileBytes); err != nil {

		http.Error(w, "Error al ejecutar: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Upload the file to a block blob
	_, err = client.UploadFile(context.TODO(), "participantes", "salud/consentimiento/mx/"+strconv.Itoa(int(idGenerado))+fileExt, tempFile,
		&azblob.UploadFileOptions{
			BlockSize:   int64(1024),
			Concurrency: uint16(3),
			// If Progress is non-nil, this function is called periodically as bytes are uploaded.
			Progress: func(bytesTransferred int64) {

			},
		})
	if err != nil {

		http.Error(w, "Error al ejecutar: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//Asignando la URL del blob a una variable
	medio := strconv.Itoa(int(idGenerado)) + fileExt

	_, err = db.Exec("UPDATE ConsentimientoMEX SET firma = @p1 WHERE id = @p2", medio, idGenerado)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Fprint(w, "1")
}

func Consentimientopostgt(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	participante := r.FormValue("firma1")
	fechafirma := r.FormValue("fechafirma")
	Codigo := r.FormValue("Codigo")

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into [GWFORMS].[dbo].[ConsentimientoSM] (nomparticipante , fechafirma, codigo)  OUTPUT INSERTED.ID values (@p1,@p2,@p3)")
	if err != nil {
		log.Fatal(err)
	}
	var idGenerado int
	err = stmt.QueryRow(participante, fechafirma, Codigo).Scan(&idGenerado)
	if err != nil {
		log.Fatal(err)
	}

	cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
	if err != nil {

		http.Error(w, "Error al ejecutar: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
	client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
	if err != nil {

		http.Error(w, "Error al ejecutar: "+err.Error(), http.StatusInternalServerError)
		return
	}
	file, fileHandler, err := r.FormFile("pdfConsentimiento")
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Obtiene el nombre del archivo cargado
	fileName := fileHandler.Filename

	//Obtiene la extensión del archivo cargado
	fileExt := filepath.Ext(fileName)

	fileBytes, err := io.ReadAll(file)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tempFile, err := os.CreateTemp("", "upload-*.dat")
	if err != nil {

		http.Error(w, "Error al ejecutar: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	// Escribe los bytes del archivo en el archivo temporal
	if _, err := tempFile.Write(fileBytes); err != nil {

		http.Error(w, "Error al ejecutar: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Upload the file to a block blob
	_, err = client.UploadFile(context.TODO(), "participantes", "salud/consentimiento/gt/"+strconv.Itoa(int(idGenerado))+fileExt, tempFile,
		&azblob.UploadFileOptions{
			BlockSize:   int64(1024),
			Concurrency: uint16(3),
			// If Progress is non-nil, this function is called periodically as bytes are uploaded.
			Progress: func(bytesTransferred int64) {

			},
		})
	if err != nil {

		http.Error(w, "Error al ejecutar: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//Asignando la URL del blob a una variable
	medio := strconv.Itoa(int(idGenerado)) + fileExt

	_, err = db.Exec("UPDATE ConsentimientoMEX SET firma = @p1 WHERE id = @p2", medio, idGenerado)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Fprint(w, "1")
}

/* func EmailSalud(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	// Leer el PDF desde la solicitud y escribirlo en el archivo adjunto
	pdfFile, _, err := r.FormFile("pdfConsentimiento")
	if err != nil {
		log.Fatal(err)
	}
	defer pdfFile.Close()

	pdfData, err := io.ReadAll(pdfFile)
	if err != nil {
		http.Error(w, "Error al leer el archivo PDF", http.StatusBadRequest)
	}
	emailform := r.FormValue("email")
	// Controller.Configuración de correo
	smtpServer := Cfg.SmtpServer
	smtpPort := Cfg.SmtpPort
	smtpUsername := Cfg.SmtpUsername
	smtpPassword := Cfg.SmtpPassword

	// Controller.Configurar el cliente SMTP
	email := gomail.NewMessage()
	email.SetHeader("From", Cfg.SmtpUsername)
	email.SetHeader("To", emailform)
	email.SetHeader("Subject", "Comprobante de Consentimiento")
	email.SetBody("text/html", "Buen día, <br><br> A continuación se adjunta el comprobante del consentimiento que acabas de completar. Cualquier duda, puedes contactar al correo estudio@glasswing.org.  <br><br> ¡Mil gracias!")
	// Adjuntar el archivo PDF
	email.Attach("archivo.pdf", gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write(pdfData)
		return err
	}))

	// Crear el cliente SMTP y enviar el correo
	dialer := gomail.NewDialer(smtpServer, smtpPort, smtpUsername, smtpPassword)
	if err := dialer.DialAndSend(email); err != nil {
		log.Fatal(err)
		http.Error(w, "Error al enviar el correo electrónico", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
} */

func getAccessToken() (string, error) {
	tenantID := os.Getenv("GRAPH_TENANT_ID")
	clientID := os.Getenv("GRAPH_CLIENT_ID")
	clientSecret := os.Getenv("GRAPH_CLIENT_SECRET")

	url := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", tenantID)
	data := "client_id=" + clientID +
		"&scope=https%3A%2F%2Fgraph.microsoft.com%2F.default" +
		"&client_secret=" + clientSecret +
		"&grant_type=client_credentials"

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	if token, ok := result["access_token"].(string); ok {
		return token, nil
	}
	return "", fmt.Errorf("no se pudo obtener el token: %v", result)
}

func EmailSalud(w http.ResponseWriter, r *http.Request) {
	// Leer archivo PDF enviado en el formulario
	pdfFile, _, err := r.FormFile("pdfConsentimiento")
	if err != nil {
		http.Error(w, "Error al leer el archivo PDF: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer pdfFile.Close()

	pdfData, err := io.ReadAll(pdfFile)
	if err != nil {
		http.Error(w, "Error al leer el archivo PDF", http.StatusBadRequest)
		return
	}

	recipient := r.FormValue("email")
	fromUser := os.Getenv("GRAPH_USER_ID")

	// Obtener token de acceso
	token, err := getAccessToken()
	if err != nil {
		http.Error(w, "Error al obtener token de acceso: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Codificar PDF a base64
	encodedPDF := base64.StdEncoding.EncodeToString(pdfData)

	// Construir cuerpo del correo
	emailPayload := map[string]interface{}{
		"message": map[string]interface{}{
			"subject": "Comprobante de Consentimiento",
			"body": map[string]string{
				"contentType": "HTML",
				"content":     "Buen día,<br><br> A continuación se adjunta el comprobante del consentimiento que acabas de completar. Cualquier duda, puedes contactar al correo estudio@glasswing.org.<br><br>¡Mil gracias!",
			},
			"toRecipients": []map[string]interface{}{
				{
					"emailAddress": map[string]string{
						"address": recipient,
					},
				},
			},
			"attachments": []map[string]interface{}{
				{
					"@odata.type":  "#microsoft.graph.fileAttachment",
					"name":         "archivo.pdf",
					"contentBytes": encodedPDF,
					"contentType":  "application/pdf",
					"contentId":    "archivo.pdf",
					"isInline":     false,
				},
			},
		},
		"saveToSentItems": "false",
	}

	body, _ := json.Marshal(emailPayload)

	graphURL := fmt.Sprintf("https://graph.microsoft.com/v1.0/users/%s/sendMail", fromUser)
	req, err := http.NewRequest("POST", graphURL, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "Error al construir la solicitud: "+err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Error al enviar el correo: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Correo enviado exitosamente"))
	} else {
		respBody, _ := io.ReadAll(resp.Body)
		http.Error(w, "Error del Graph API: "+string(respBody), resp.StatusCode)
	}
}

func Postesaludhnnew(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Obtener los valores de los campos del formulario
	idF := 15
	// Autorizacion Menores
	nombrecompleto := r.FormValue("nombreRep")
	identidad := r.FormValue("DUIP")
	parentesco := r.FormValue("parentesco")
	telefono := r.FormValue("numResp")

	// Formulario Registro
	fechaNac := r.FormValue("FechaN")
	nombres := r.FormValue("nombreCompleto")
	apellidos := r.FormValue("apellido")
	sexo := r.FormValue("Sexo")
	pais := 3
	departamento := r.FormValue("pais")
	municipio := r.FormValue("departamento")
	nacionalidad := r.FormValue("nacionalidad")
	identidad2 := r.FormValue("DUIM")
	estudia := r.FormValue("TdC")
	estudioAlcanzado := r.FormValue("UT")
	grado := r.FormValue("GA")
	seccion := r.FormValue("seccion")
	turno := r.FormValue("Turno")
	discapacidad := r.FormValue("dis")
	nuevogw := r.FormValue("gp")
	tipoParticipante := r.FormValue("tipo")
	personalinst := r.FormValue("acc2")
	perfil := r.FormValue("acc3")
	subtipo := r.FormValue("acc4")
	idSede := r.FormValue("sede2")
	autovozimg := r.FormValue("aut2")

	// Datos Centro Educativo

	nombre := r.FormValue("sedehn")
	tipo := r.FormValue("tiposedehn")
	cargo := r.FormValue("cargosedehn")
	departamentoe := r.FormValue("depto")
	municipioe := r.FormValue("munic")
	aldea := r.FormValue("aldeahn")
	caserio := r.FormValue("caseriohn")
	codigo := r.FormValue("codigohn")
	jornada := r.FormValue("jornadahn")
	nivel := r.FormValue("nivelhn")
	ciclo := r.FormValue("ciclohn")
	zona := r.FormValue("zonahn")

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		tx.Rollback()
		http.Error(w, "Error conexion", http.StatusInternalServerError)

	}
	stmt, err := tx.Prepare("INSERT INTO FormularioRegisto(idF,tipoParticipante,pais,idSede,fechaNac,nombres,apellidos,sexo,nacionalidad,identidad,discapacidad,estudia,estudioAlcanzado,grado,turno,seccion,nuevogw,departamento,municipio,autovozimg,personalinst,perfil,subtipo) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14,@P15,@P16,@P17,@P18,@P19,@P20,@P21,@P22,@P23)")
	if err != nil {
		log.Fatal(err)
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
	}

	var lastInsertId int
	err = stmt.QueryRow(idF, tipoParticipante, pais, idSede, fechaNac, nombres, apellidos, sexo, nacionalidad, identidad2, discapacidad, estudia, estudioAlcanzado, grado, turno, seccion, nuevogw, departamento, municipio, autovozimg, personalinst, perfil, subtipo).Scan(&lastInsertId)
	if err != nil {
		log.Fatal(err)
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
	}

	if nombrecompleto != "" {

		stmt, err = tx.Prepare("INSERT INTO AutorizacionMenores(idFR,nombrecompleto, telefono, identidad, parentesco) VALUES(@P1,@P2,@P3,@P4,@P5)")
		if err != nil {
			log.Fatal(err)
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta automenores", http.StatusInternalServerError)
		}

		_, err = stmt.Exec(lastInsertId, nombrecompleto, telefono, identidad, parentesco)
		if err != nil {
			log.Fatal(err)
			tx.Rollback()
			http.Error(w, "Error al ejecutar la consulta automenores", http.StatusInternalServerError)
		}

	}
	//Inserta solo si se cumple lo siguiente
	if personalinst == "14" {
		stmt, err = tx.Prepare("INSERT INTO F_DatosCentroEducativoSalud([idFR],[nombre],[tipo],[cargo],[departamento],[municipio],[aldea],[caserio],[codigo],[jornada],[nivel],[ciclo],[zona]) VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13)")
		if err != nil {
			log.Fatal(err)
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta DatosCentroEducativo", http.StatusInternalServerError)
		}

		_, err = stmt.Exec(lastInsertId, nombre, tipo, cargo, departamentoe, municipioe, aldea, caserio, codigo, jornada, nivel, ciclo, zona)
		if err != nil {
			log.Fatal(err)
			tx.Rollback()
			http.Error(w, "Error al ejecutar la consulta DatosCentroEducativo", http.StatusInternalServerError)
		}

	}

	discapacidadmultiple := r.Form["dis2"]
	if len(discapacidadmultiple) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO MultipleDiscapacidad (idFR,iddiscapacidad) VALUES (@P1,@P2)")
		if err != nil {
			log.Fatal(err)
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta discapacidad", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidadmultiple {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					log.Fatal(err)
					tx.Rollback()
					http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
					return
				}
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles("public/HTML/Respuesta.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Renderizar la plantilla con los datos
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
