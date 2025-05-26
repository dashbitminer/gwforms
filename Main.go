package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"formulario/Controller"
	"formulario/Routes"
	"html/template"
	"image"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/microsoft/go-mssqldb"

	"encoding/json"
)

type Modulo struct {
	ID     int
	Nombre string
	Dui    string
	Sede   string
	Fecha  string
	Module int
	Sedeid int
}

type Respuestas struct {
	Value      string
	Label      string
	IdPregunta string
}

type FormData struct {
	Parentesco   []Controller.Option
	Sexo         []Controller.Option
	Camisa       []Controller.Option
	Departamento []Controller.Option
	Respuestas   []Controller.Option
	Respuestas2  []Controller.Option
	Respuestas3  []Controller.Option
	Respuestas4  []Controller.Option
	Ultgrad      []Controller.Option
	Grado        []Controller.Option
	Perfil       []Controller.Option
	Relacioncom  []Controller.Option
	Institucion  []Controller.Option
	Copdonor     []Controller.Option
	Programa     []Controller.Option
	Seccion      []Controller.Option
	Turno        []Controller.Option
	Transporte   []Controller.Option
	Estipendio   []Controller.Option
	Estado_civ   []Controller.Option
	Actividades  []Controller.Option
	Depsede      []Controller.Option
	Grupo        []Controller.Option
	Pueblo       []Controller.Option
	Idioma       []Controller.Option
	ActLaboral   []Controller.Option
	Paises       []Controller.Option
	Paisesn      []Controller.Option
	Empresas     []Controller.Option
	Afectacion   []Controller.Option
	Actitud      []Controller.Option
	Estrategia   []Controller.Option
	Practica     []Controller.Option
	Estado       []Controller.Option
	Afirmacion   []Controller.Option
	Lugar        []Controller.Option
	Clima        []Controller.Option
	Conoce       []Controller.Option
	Consciente   []Controller.Option
	Idpais       int
}

type FormDataNew struct {
	Preguntas     []Respuestas
	Solucion      []Respuestas
	Solucionedu   []Respuestas
	Solucionsalud []Respuestas
	Solucionsa    []Respuestas
}

func main() {
	Cfg, err := Controller.LoadConfig()
	if err != nil {
		log.Fatal("Error cargando la Controller.Configuración:", err)
		return
	}
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	fs2 := http.FileServer(http.Dir("node_modules"))
	http.Handle("/static/", http.StripPrefix("/static/", fs2))
	Routes.ConfigurarRutasraices(Cfg)
	Routes.ConfigurarRutassql(Cfg)
	Routes.ConfigurarRutassalud(Cfg)
	Routes.ConfigurarRutasSanamente(Cfg)
	Routes.ConfigurarRutasjuventudes(Cfg)
	Routes.ConfigurarRutasvoluntariado(Cfg)
	http.HandleFunc("/capacitaciones/empresas", func(w http.ResponseWriter, r *http.Request) { capacitaciones(w) })
	http.HandleFunc("/samsung/inscripcion", func(w http.ResponseWriter, r *http.Request) { formulariosamsung(w, Cfg) })
	http.HandleFunc("/samsung/post", func(w http.ResponseWriter, r *http.Request) { postsamsung(w, r, Cfg) })
	http.HandleFunc("/voluntariado/post/sv", func(w http.ResponseWriter, r *http.Request) { postvolsv(w, r, Cfg) })
	http.HandleFunc("/voluntariado/post/pn", func(w http.ResponseWriter, r *http.Request) { postvolpn(w, r, Cfg) })
	http.HandleFunc("/voluntariado/post/col", func(w http.ResponseWriter, r *http.Request) { postvolcol(w, r, Cfg) })
	http.HandleFunc("/voluntariado/post/cr", func(w http.ResponseWriter, r *http.Request) { postvolcr(w, r, Cfg) })
	http.HandleFunc("/voluntariado/post/hn", func(w http.ResponseWriter, r *http.Request) { postvolhn(w, r, Cfg) })
	http.HandleFunc("/voluntariado/post/gt", func(w http.ResponseWriter, r *http.Request) { postvolgt(w, r, Cfg) })
	http.HandleFunc("/voluntariado/post/mx", func(w http.ResponseWriter, r *http.Request) { postvolmx(w, r, Cfg) })
	http.HandleFunc("/voluntariado/post/rd", func(w http.ResponseWriter, r *http.Request) { postvolrd(w, r, Cfg) })
	http.HandleFunc("/voluntariado/sv", func(w http.ResponseWriter, r *http.Request) { voluntariadosv(w, Cfg) })
	http.HandleFunc("/voluntariado/pn", func(w http.ResponseWriter, r *http.Request) { voluntariadopn(w, Cfg) })
	http.HandleFunc("/voluntariado/rd", func(w http.ResponseWriter, r *http.Request) { voluntariadord(w, Cfg) })
	http.HandleFunc("/voluntariado/col", func(w http.ResponseWriter, r *http.Request) { voluntariadocol(w, Cfg) })
	http.HandleFunc("/voluntariado/gt", func(w http.ResponseWriter, r *http.Request) { voluntariadogt(w, Cfg) })
	http.HandleFunc("/voluntariado/cr", func(w http.ResponseWriter, r *http.Request) { voluntariadocr(w, Cfg) })
	http.HandleFunc("/voluntariado/hn", func(w http.ResponseWriter, r *http.Request) { voluntariadohn(w, Cfg) })
	http.HandleFunc("/voluntariado/mx", func(w http.ResponseWriter, r *http.Request) { voluntariadomx(w, Cfg) })
	http.HandleFunc("/educacion/hn", func(w http.ResponseWriter, r *http.Request) { Eduhn(w, r, Cfg) })
	http.HandleFunc("/educacion/hn/post", func(w http.ResponseWriter, r *http.Request) { posteduhn(w, r, Cfg) })
	http.HandleFunc("/educacion/sv", func(w http.ResponseWriter, r *http.Request) { Edusv(w, r, Cfg) })
	http.HandleFunc("/glasswing/edu/sv", func(w http.ResponseWriter, r *http.Request) { prueba(w, Cfg) })
	http.HandleFunc("/glasswing/mentees/menores", func(w http.ResponseWriter, r *http.Request) { formulariomentoriasmenores(w, Cfg) })
	http.HandleFunc("/glasswing/mentees/post", func(w http.ResponseWriter, r *http.Request) { postmemteesmenores(w, r, Cfg) })
	http.HandleFunc("/educacion/sv/post", func(w http.ResponseWriter, r *http.Request) { postedusv(w, r, Cfg) })
	http.HandleFunc("/educacion/gt", func(w http.ResponseWriter, r *http.Request) { Edugt(w, r, Cfg) })
	http.HandleFunc("/educacion/gt/post", func(w http.ResponseWriter, r *http.Request) { postedugt(w, r, Cfg) })
	http.HandleFunc("/educacion/cr", func(w http.ResponseWriter, r *http.Request) { Educr(w, r, Cfg) })
	http.HandleFunc("/educacion/cr/post", func(w http.ResponseWriter, r *http.Request) { posteducr(w, r, Cfg) })
	http.HandleFunc("/educacion/col", func(w http.ResponseWriter, r *http.Request) { Educol(w, r, Cfg) })
	http.HandleFunc("/educacion/col/post", func(w http.ResponseWriter, r *http.Request) { posteducol(w, r, Cfg) })
	http.HandleFunc("/educacion/pn", func(w http.ResponseWriter, r *http.Request) { Edupn(w, r, Cfg) })
	http.HandleFunc("/educacion/pn/post", func(w http.ResponseWriter, r *http.Request) { postedupn(w, r, Cfg) })
	http.HandleFunc("/educacion/rd", func(w http.ResponseWriter, r *http.Request) { Edurd(w, r, Cfg) })
	http.HandleFunc("/educacion/rd/post", func(w http.ResponseWriter, r *http.Request) { postedurd(w, r, Cfg) })
	http.HandleFunc("/educacion/mx", func(w http.ResponseWriter, r *http.Request) { Edumx(w, r, Cfg) })
	http.HandleFunc("/educacion/mx/post", func(w http.ResponseWriter, r *http.Request) { postedumx(w, r, Cfg) })
	http.HandleFunc("/educacion/ny/eng", func(w http.ResponseWriter, r *http.Request) { eduNYeng(w, Cfg) })
	http.HandleFunc("/educacion/ny/esp", func(w http.ResponseWriter, r *http.Request) { eduNYesp(w, Cfg) })
	http.HandleFunc("/educacion/ny/fr", func(w http.ResponseWriter, r *http.Request) { eduNYfr(w, Cfg) })
	http.HandleFunc("/educacion/ny/post", func(w http.ResponseWriter, r *http.Request) { PostEduNY(w, r, Cfg) })
	http.HandleFunc("/educacion/nna/col", func(w http.ResponseWriter, r *http.Request) { eduNNAcol(w, Cfg) })
	http.HandleFunc("/educacion/nna/pan", func(w http.ResponseWriter, r *http.Request) { eduNNApan(w, Cfg) })
	http.HandleFunc("/educacion/nna/post", func(w http.ResponseWriter, r *http.Request) { PostNnaEdu(w, r, Cfg) })
	http.HandleFunc("/educacion/inscripcion/nna/rd", func(w http.ResponseWriter, r *http.Request) { eduNNArd(w, Cfg) })
	http.HandleFunc("/educacion/inscripcion/nna/mx", func(w http.ResponseWriter, r *http.Request) { eduNNAmx(w, Cfg) })
	http.HandleFunc("/educacion/inscripcion/nna/post", func(w http.ResponseWriter, r *http.Request) { PostNNAEducacion(w, r, Cfg) })
	http.HandleFunc("/salud/sv", func(w http.ResponseWriter, r *http.Request) { saludsv(w, Cfg) })
	http.HandleFunc("/salud/nna/sv", func(w http.ResponseWriter, r *http.Request) { Nnasv(w, r, Cfg) })
	http.HandleFunc("/salud/nna/gt", func(w http.ResponseWriter, r *http.Request) { Nnagt(w, r, Cfg) })
	http.HandleFunc("/salud/nna/hn", func(w http.ResponseWriter, r *http.Request) { Nnahn(w, r, Cfg) })
	http.HandleFunc("/salud/nna/pn", func(w http.ResponseWriter, r *http.Request) { Nnapn(w, r, Cfg) })
	http.HandleFunc("/salud/nna/cr", func(w http.ResponseWriter, r *http.Request) { Nnacr(w, r, Cfg) })
	http.HandleFunc("/salud/nna/col", func(w http.ResponseWriter, r *http.Request) { Nnacol(w, r, Cfg) })
	http.HandleFunc("/salud/nna/mx", func(w http.ResponseWriter, r *http.Request) { Nnamx(w, r, Cfg) })
	http.HandleFunc("/salud/nna/post", func(w http.ResponseWriter, r *http.Request) { PostNnasalud(w, r, Cfg) })
	http.HandleFunc("/salud/nna/cisf/sv", func(w http.ResponseWriter, r *http.Request) { NnaCISFsv(w, r, Cfg) })
	http.HandleFunc("/salud/nna/cisf/post", func(w http.ResponseWriter, r *http.Request) { PostNnaCISF(w, r, Cfg) })
	http.HandleFunc("/salud/sv/post", func(w http.ResponseWriter, r *http.Request) { postesaludsv(w, r, Cfg) })
	http.HandleFunc("/salud/gt", func(w http.ResponseWriter, r *http.Request) { saludgt(w, Cfg) })
	http.HandleFunc("/salud/gt/post", func(w http.ResponseWriter, r *http.Request) { postesaludgt(w, r, Cfg) })
	http.HandleFunc("/salud/hn", func(w http.ResponseWriter, r *http.Request) { saludhn(w, Cfg) })
	http.HandleFunc("/salud/hn/post", func(w http.ResponseWriter, r *http.Request) { postesaludhn(w, r, Cfg) })
	http.HandleFunc("/salud/cr", func(w http.ResponseWriter, r *http.Request) { saludcr(w, Cfg) })
	http.HandleFunc("/salud/cr/post", func(w http.ResponseWriter, r *http.Request) { postesaludcr(w, r, Cfg) })
	http.HandleFunc("/salud/pn", func(w http.ResponseWriter, r *http.Request) { saludpn(w, Cfg) })
	http.HandleFunc("/salud/pn/post", func(w http.ResponseWriter, r *http.Request) { postesaludpn(w, r, Cfg) })
	http.HandleFunc("/salud/mx", func(w http.ResponseWriter, r *http.Request) { saludmx(w, Cfg) })
	http.HandleFunc("/salud/mx/post", func(w http.ResponseWriter, r *http.Request) { postesaludmx(w, r, Cfg) })
	http.HandleFunc("/salud/col", func(w http.ResponseWriter, r *http.Request) { saludcol(w, Cfg) })
	http.HandleFunc("/salud/col/post", func(w http.ResponseWriter, r *http.Request) { postesaludcol(w, r, Cfg) })
	http.HandleFunc("/salud/rd", func(w http.ResponseWriter, r *http.Request) { saludrd(w, Cfg) })
	http.HandleFunc("/salud/rd/post", func(w http.ResponseWriter, r *http.Request) { postesaludrd(w, r, Cfg) })
	http.HandleFunc("/sanamente/equip", func(w http.ResponseWriter, r *http.Request) { equipGeneral(w, r, Cfg) })
	http.HandleFunc("/sanamente/equip/pdf", func(w http.ResponseWriter, r *http.Request) { equipGeneralGuardado(w, r, Cfg) })
	http.HandleFunc("/sanamente/equip/sv", func(w http.ResponseWriter, r *http.Request) { equipsv(w, Cfg) })
	http.HandleFunc("/sanamente/equip/hn", func(w http.ResponseWriter, r *http.Request) { equiphn(w, Cfg) })
	http.HandleFunc("/sanamente/equip/gt", func(w http.ResponseWriter, r *http.Request) { equipgt(w, Cfg) })
	http.HandleFunc("/sanamente/equip/mx", func(w http.ResponseWriter, r *http.Request) { equipmx(w, Cfg) })
	//	http.HandleFunc("/sanamente/equip/email", func(w http.ResponseWriter, r *http.Request) { emailequip(w, r, Cfg) })
	http.HandleFunc("/sanamente/equip/email", func(w http.ResponseWriter, r *http.Request) { emailequip2(w, r) })
	http.HandleFunc("/sanamente/intervencion", func(w http.ResponseWriter, r *http.Request) { intervencionistas(w, r, Cfg) })
	http.HandleFunc("/sanamente/intervencion/post", func(w http.ResponseWriter, r *http.Request) { postintervencionistas(w, r, Cfg) })
	http.HandleFunc("/sanamente/inscripcion/post", func(w http.ResponseWriter, r *http.Request) { postsanamenteinscripcion(w, r, Cfg) })
	http.HandleFunc("/sanamente/inscripcion/col", func(w http.ResponseWriter, r *http.Request) { sanamentecol(w, r) })
	http.HandleFunc("/sanamente/inscripcion/mx", func(w http.ResponseWriter, r *http.Request) { sanamentemx(w, r) })
	http.HandleFunc("/sanamente/inscripcion/gt", func(w http.ResponseWriter, r *http.Request) { sanamentegt(w, r) })
	http.HandleFunc("/sanamente/inscripcion/cr", func(w http.ResponseWriter, r *http.Request) { sanamentecr(w, r) })
	http.HandleFunc("/sanamente/inscripcion/sv", func(w http.ResponseWriter, r *http.Request) { sanamentesv(w, r) })
	http.HandleFunc("/sanamente/inscripcion/sv2", func(w http.ResponseWriter, r *http.Request) { sanamentesvcopy(w) })
	http.HandleFunc("/sanamente/inscripcion/hn", func(w http.ResponseWriter, r *http.Request) { sanamentehn(w, r) })
	http.HandleFunc("/sanamente/inscripcion/pn", func(w http.ResponseWriter, r *http.Request) { sanamentepn(w, r) })
	http.HandleFunc("/sanamente/moduloa/sv", func(w http.ResponseWriter, r *http.Request) { moduloasvHandler(w, Cfg) })
	http.HandleFunc("/sanamente/modulo/post", func(w http.ResponseWriter, r *http.Request) { postmodulo(w, r, Cfg) })
	http.HandleFunc("/sanamente/modulo/col/post", func(w http.ResponseWriter, r *http.Request) { postmodulocol(w, r, Cfg) })
	http.HandleFunc("/sanamente/modulo/intervencionista/post", func(w http.ResponseWriter, r *http.Request) { postmoduloInterv(w, r, Cfg) })
	http.HandleFunc("/sanamente/equip/post", func(w http.ResponseWriter, r *http.Request) { postequip(w, r, Cfg) })
	http.HandleFunc("/sanamente/modulob/sv", func(w http.ResponseWriter, r *http.Request) { modulobsvHandler(w, Cfg) })
	http.HandleFunc("/sanamente/moduloc/sv", func(w http.ResponseWriter, r *http.Request) { modulocsvHandler(w, Cfg) })
	http.HandleFunc("/sanamente/moduloa/gt", func(w http.ResponseWriter, r *http.Request) { moduloagtHandler(w, Cfg) })
	http.HandleFunc("/sanamente/modulob/gt", func(w http.ResponseWriter, r *http.Request) { modulobgtHandler(w, Cfg) })
	http.HandleFunc("/sanamente/moduloc/gt", func(w http.ResponseWriter, r *http.Request) { modulocgtHandler(w, Cfg) })
	http.HandleFunc("/sanamente/moduloa/cr", func(w http.ResponseWriter, r *http.Request) { moduloacrHandler(w, Cfg) })
	http.HandleFunc("/sanamente/modulob/cr", func(w http.ResponseWriter, r *http.Request) { modulobcrHandler(w, Cfg) })
	http.HandleFunc("/sanamente/moduloc/cr", func(w http.ResponseWriter, r *http.Request) { moduloccrHandler(w, Cfg) })
	http.HandleFunc("/sanamente/moduloa/mx", func(w http.ResponseWriter, r *http.Request) { moduloamxHandler(w, Cfg) })
	http.HandleFunc("/sanamente/modulob/mx", func(w http.ResponseWriter, r *http.Request) { modulobmxHandler(w, Cfg) })
	http.HandleFunc("/sanamente/moduloc/mx", func(w http.ResponseWriter, r *http.Request) { modulocmxHandler(w, Cfg) })
	http.HandleFunc("/sanamente/moduloa/hn", func(w http.ResponseWriter, r *http.Request) { moduloahnHandler(w, Cfg) })
	http.HandleFunc("/sanamente/modulob/hn", func(w http.ResponseWriter, r *http.Request) { modulobhnHandler(w, Cfg) })
	http.HandleFunc("/sanamente/moduloc/hn", func(w http.ResponseWriter, r *http.Request) { modulochnHandler(w, Cfg) })
	http.HandleFunc("/sanamente/moduloa/col", func(w http.ResponseWriter, r *http.Request) { moduloacolHandler(w, Cfg) })
	http.HandleFunc("/sanamente/modulob/col", func(w http.ResponseWriter, r *http.Request) { modulobcolHandler(w, Cfg) })
	http.HandleFunc("/sanamente/moduloc/col", func(w http.ResponseWriter, r *http.Request) { moduloccolHandler(w, Cfg) })
	http.HandleFunc("/sanamente/modulo/col", func(w http.ResponseWriter, r *http.Request) { moduloColHandler(w, Cfg) })
	http.HandleFunc("/sanamente/modulo/institucional/mx", func(w http.ResponseWriter, r *http.Request) { moduloMXFBEInstHandler(w, Cfg) })
	http.HandleFunc("/sanamente/modulo/comunitaria", func(w http.ResponseWriter, r *http.Request) { moduloFBEComHandler(w, r, Cfg) })
	http.HandleFunc("/sanamente/moduloa/pn", func(w http.ResponseWriter, r *http.Request) { moduloapnHandler(w, Cfg) })
	http.HandleFunc("/sanamente/modulob/pn", func(w http.ResponseWriter, r *http.Request) { modulobpnHandler(w, Cfg) })
	http.HandleFunc("/sanamente/moduloc/pn", func(w http.ResponseWriter, r *http.Request) { modulocpnHandler(w, Cfg) })
	http.HandleFunc("/sanamente/modulo/intervencionista", func(w http.ResponseWriter, r *http.Request) { moduloInterv(w, r, Cfg) })
	http.HandleFunc("/sanamente/encuesta", func(w http.ResponseWriter, r *http.Request) { EncuestaSM(w, r, Cfg) })
	http.HandleFunc("/sanamente/encuesta2", func(w http.ResponseWriter, r *http.Request) { EncuestaSM2(w, r, Cfg) })
	http.HandleFunc("/sanamente/encuesta3", func(w http.ResponseWriter, r *http.Request) { EncuestaSM3(w, r, Cfg) })
	http.HandleFunc("/sanamente/encuesta2/post", func(w http.ResponseWriter, r *http.Request) { postEncuestaSM2(w, r, Cfg) })
	http.HandleFunc("/sanamente/encuesta3/post", func(w http.ResponseWriter, r *http.Request) { postEncuestaSM3(w, r, Cfg) })
	http.HandleFunc("/sanamente/encuesta/sv", func(w http.ResponseWriter, r *http.Request) { EncuestaSMsv(w, r, Cfg) })
	http.HandleFunc("/sanamente/encuesta/post", func(w http.ResponseWriter, r *http.Request) { postEncuestaSM(w, r, Cfg) })
	http.HandleFunc("/juventud/sv/inscripcion", func(w http.ResponseWriter, r *http.Request) { formulariomdvHandler(w, Cfg) })
	http.HandleFunc("/juventud/sv/psicologico", func(w http.ResponseWriter, r *http.Request) { formulariomdvpsicologico(w, Cfg) })
	http.HandleFunc("/juventud/sv/psicologico/santa_ana", func(w http.ResponseWriter, r *http.Request) { formulariomdvpsicologicoSantaana(w, Cfg) })
	http.HandleFunc("/mdv/sv/psicologico", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/juventud/sv/psicologico", http.StatusMovedPermanently)
	})
	http.HandleFunc("/mdv/sv/psicologico/post", func(w http.ResponseWriter, r *http.Request) { postmdvpsicologico(w, r, Cfg) })
	http.HandleFunc("/mdv/sv/psicologico/santaana/post", func(w http.ResponseWriter, r *http.Request) { postmdvpsicologicoSantaana(w, r, Cfg) })
	http.HandleFunc("/mdv/sv/jc", func(w http.ResponseWriter, r *http.Request) { formulariojcHandler(w, Cfg) })
	http.HandleFunc("/juventud/sv/red", func(w http.ResponseWriter, r *http.Request) { formulariomdvred(w, Cfg) })
	http.HandleFunc("/mdv/red", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/juventud/sv/red", http.StatusMovedPermanently)
	})
	http.HandleFunc("/mdv/fablab", func(w http.ResponseWriter, r *http.Request) { formulariomdvfablab(w, Cfg) })
	http.HandleFunc("/mdv/percepcion", func(w http.ResponseWriter, r *http.Request) { formulariopercepcionmdv(w, Cfg) })
	http.HandleFunc("/mdv/sv/post", func(w http.ResponseWriter, r *http.Request) { postmdv(w, r, Cfg) })
	http.HandleFunc("/mdv/red/post", func(w http.ResponseWriter, r *http.Request) { postredmdv(w, r, Cfg) })
	http.HandleFunc("/participamas/sv", func(w http.ResponseWriter, r *http.Request) { formularioParticipamas(w, Cfg) })
	http.HandleFunc("/participamas/sv/post", func(w http.ResponseWriter, r *http.Request) { postParticipamas(w, r, Cfg) })
	http.HandleFunc("/mentorias/mentees", func(w http.ResponseWriter, r *http.Request) { formulariomentorias(w, Cfg) })
	http.HandleFunc("/mentorias/mentores", func(w http.ResponseWriter, r *http.Request) { formularioMentores(w, Cfg) })
	http.HandleFunc("/mentorias/mentores/millicom", func(w http.ResponseWriter, r *http.Request) { formularioMentoresMillicom(w, Cfg) })
	http.HandleFunc("/mentorias/mentees/post", func(w http.ResponseWriter, r *http.Request) { postmemtees(w, r, Cfg) })
	http.HandleFunc("/mentores/mentores/post", func(w http.ResponseWriter, r *http.Request) { postmentores(w, r, Cfg) })
	http.HandleFunc("/mentores/millicom/post", func(w http.ResponseWriter, r *http.Request) { postMentoresMillicom(w, r, Cfg) })
	http.HandleFunc("/voluntariado/corporativo", func(w http.ResponseWriter, r *http.Request) { formularioCorporativo(w, Cfg) })
	http.HandleFunc("/voluntariado/corporativo/en", func(w http.ResponseWriter, r *http.Request) { formularioCorporativoEnglish(w, Cfg) })
	http.HandleFunc("/voluntariado/corporativo/post", func(w http.ResponseWriter, r *http.Request) { postCorporativo(w, r, Cfg) })
	http.HandleFunc("/jli/participante", func(w http.ResponseWriter, r *http.Request) { participanteJLI(w, Cfg) })
	http.HandleFunc("/jli/participante/cohorte3", func(w http.ResponseWriter, r *http.Request) { participanteJLI2(w, Cfg) })
	http.HandleFunc("/jli/participante/cohorte7", func(w http.ResponseWriter, r *http.Request) { participanteJLIcho5(w, Cfg) })
	http.HandleFunc("/jli/participante/cohorte6", func(w http.ResponseWriter, r *http.Request) { participanteJLIcho6(w, Cfg) })
	http.HandleFunc("/jli/participante/post", func(w http.ResponseWriter, r *http.Request) { postJLI(w, r, Cfg) })
	http.HandleFunc("/mdv/percepcion/post", func(w http.ResponseWriter, r *http.Request) { postmdvpercepcion(w, r, Cfg) })
	http.HandleFunc("/obtenerperfil", func(w http.ResponseWriter, r *http.Request) { obtenerperfil(w, r, Cfg) })
	http.HandleFunc("/obtenersubperfil", func(w http.ResponseWriter, r *http.Request) { obtenersubperfil(w, r, Cfg) })
	http.HandleFunc("/obtenerciudades", func(w http.ResponseWriter, r *http.Request) { obtenerCiudadesHandler(w, r, Cfg) })
	http.HandleFunc("/obtenerciudades2", func(w http.ResponseWriter, r *http.Request) { obtenerCiudadesHandler2(w, r, Cfg) })
	http.HandleFunc("/voluntariadocheck", func(w http.ResponseWriter, r *http.Request) { volcheck(w, r, Cfg) })
	http.HandleFunc("/obtenerciudades3", func(w http.ResponseWriter, r *http.Request) { obtenerCiudadesHandler3(w, r, Cfg) })
	http.HandleFunc("/obteneractividad", func(w http.ResponseWriter, r *http.Request) { obteneractividadHandler(w, r, Cfg) })
	http.HandleFunc("/obtenerdepartamento2", func(w http.ResponseWriter, r *http.Request) { obtenerdep2Handler(w, r, Cfg) })
	http.HandleFunc("/obtenerdepartamentog", func(w http.ResponseWriter, r *http.Request) { obtenerdepGeneral(w, r, Cfg) })
	http.HandleFunc("/obtenermunicipiog", func(w http.ResponseWriter, r *http.Request) { obtenermunGeneral(w, r, Cfg) })
	http.HandleFunc("/obtenersede2", func(w http.ResponseWriter, r *http.Request) { obtenersede2Handler(w, r, Cfg) })
	http.HandleFunc("/obtenersubactividad", func(w http.ResponseWriter, r *http.Request) { obtenersubactividadHandler(w, r, Cfg) })
	http.HandleFunc("/obtenermunsede", func(w http.ResponseWriter, r *http.Request) { obtenermunsedeHandler(w, r, Cfg) })
	http.HandleFunc("/obtenersede", func(w http.ResponseWriter, r *http.Request) { obtenersedeHandler(w, r, Cfg) })
	http.HandleFunc("/obtenermunsedeedu", func(w http.ResponseWriter, r *http.Request) { obtenermunsedeHandleredu(w, r, Cfg) })
	http.HandleFunc("/obtenersedeedu", func(w http.ResponseWriter, r *http.Request) { obtenersedeHandleredu(w, r, Cfg) })
	http.HandleFunc("/obtenermunsedesalud", func(w http.ResponseWriter, r *http.Request) { obtenermunsedeHandlersalud(w, r, Cfg) })
	http.HandleFunc("/obtenersedesalud", func(w http.ResponseWriter, r *http.Request) { obtenersedeHandlersalud(w, r, Cfg) })
	http.HandleFunc("/obtenermunsedesanamente", func(w http.ResponseWriter, r *http.Request) { obtenermunsedeHandlersanamente(w, r, Cfg) })
	http.HandleFunc("/obtenersedesanamente", func(w http.ResponseWriter, r *http.Request) { obtenersedeHandlersanamente(w, r, Cfg) })
	http.HandleFunc("/obtenereventos", func(w http.ResponseWriter, r *http.Request) { obtenerevento(w, r, Cfg) })
	http.HandleFunc("/voluntariado/encuesta/satisfaccion", func(w http.ResponseWriter, r *http.Request) { satisfaccionVoluntariado(w, Cfg) })
	http.HandleFunc("/voluntariado/encuesta/inicial", func(w http.ResponseWriter, r *http.Request) { encuestainicialvoluntariado(w, Cfg) })
	http.HandleFunc("/voluntariado/encuesta/final", func(w http.ResponseWriter, r *http.Request) { encuestafinalvoluntariado(w, Cfg) })
	http.HandleFunc("/voluntariado/encuesta/satisfaccion/post", func(w http.ResponseWriter, r *http.Request) { postencuestasatisfaccionvol(w, r, Cfg) })
	http.HandleFunc("/voluntariado/encuesta/inicial/post", func(w http.ResponseWriter, r *http.Request) { postencuestainicialvol(w, r, Cfg) })
	http.HandleFunc("/voluntariado/encuesta/final/post", func(w http.ResponseWriter, r *http.Request) { postencuestafinalvol(w, r, Cfg) })
	http.HandleFunc("/refcasos", func(w http.ResponseWriter, r *http.Request) { casos(w, Cfg) })
	http.HandleFunc("/refcasos/post", func(w http.ResponseWriter, r *http.Request) { postcasos(w, r, Cfg) })
	http.HandleFunc("/encuesta/gwdata", func(w http.ResponseWriter, r *http.Request) { encuestagwdata(w, Cfg) })
	http.HandleFunc("/encuesta/gwdata/post", func(w http.ResponseWriter, r *http.Request) { postencuestagwdata(w, r, Cfg) })
	http.HandleFunc("/rrhh/feriasalud", func(w http.ResponseWriter, r *http.Request) { rrhhferia(w, Cfg) })
	http.HandleFunc("/rrhh/feriasalud/post", func(w http.ResponseWriter, r *http.Request) { postrrhhferiasalud(w, r, Cfg) })
	http.HandleFunc("/rrhh/transporte", func(w http.ResponseWriter, r *http.Request) { rrhhtransporte(w, Cfg) })
	http.HandleFunc("/rrhh/transporte/post", func(w http.ResponseWriter, r *http.Request) { postrrhhtransporte(w, r, Cfg) })
	http.HandleFunc("/voluntariado/visitacampo", func(w http.ResponseWriter, r *http.Request) { encuestavisitacampo(w, Cfg) })
	http.HandleFunc("/voluntariado/visitacampo/post", func(w http.ResponseWriter, r *http.Request) { postencuestavisitacampo(w, r, Cfg) })
	http.HandleFunc("/jli/encuestapaliativaJLI", func(w http.ResponseWriter, r *http.Request) { encuestapaliativaJLI(w, Cfg) })
	http.HandleFunc("/jli/encuestapaliativaJLI/post", func(w http.ResponseWriter, r *http.Request) { postencuestapaliativaJLI(w, r, Cfg) })
	http.HandleFunc("/obtenersedecampo", func(w http.ResponseWriter, r *http.Request) { obtenersedecampoHandler(w, r, Cfg) })
	http.HandleFunc("/encuesta/dell", func(w http.ResponseWriter, r *http.Request) { encuestadell(w, Cfg) })
	http.HandleFunc("/encuesta/dell/post", func(w http.ResponseWriter, r *http.Request) { postencuestadell(w, r, Cfg) })
	http.HandleFunc("/dell/inscripcion", func(w http.ResponseWriter, r *http.Request) { dell(w, Cfg) })
	http.HandleFunc("/dell/post", func(w http.ResponseWriter, r *http.Request) { postdell(w, r, Cfg) })
	http.HandleFunc("/finanzas/proveedores", func(w http.ResponseWriter, r *http.Request) { Finanzas_datos(w, r, Cfg) })
	http.HandleFunc("/finanzas/clientes", func(w http.ResponseWriter, r *http.Request) { Finanzas_datos_clientes(w, r, Cfg) })
	http.HandleFunc("/finanzas/datos/post", func(w http.ResponseWriter, r *http.Request) { PostFinanzas_datos(w, r, Cfg) })
	http.HandleFunc("/obtenermensaje", func(w http.ResponseWriter, r *http.Request) { obtenermensaje(w, r) })
	http.HandleFunc("/obtenerseguimiento", func(w http.ResponseWriter, r *http.Request) { obtenerseguimiento(w, r, Cfg) })
	http.HandleFunc("/obtenerseguimientosubtipo", func(w http.ResponseWriter, r *http.Request) { obtenerseguimientosubtipo(w, r) })
	http.HandleFunc("/juventud/gestion/sv", func(w http.ResponseWriter, r *http.Request) { gestionmdv(w) })
	http.HandleFunc("/juventud/gestion/sv/post", func(w http.ResponseWriter, r *http.Request) { postseguimientomdv(w, r, Cfg) })
	http.HandleFunc("/obtenerlistadomd", func(w http.ResponseWriter, r *http.Request) { obtenerlistadomdv(w, r, Cfg) })
	http.HandleFunc("/sanamente/chemonics/mun", func(w http.ResponseWriter, r *http.Request) { Obtenermunchemonics(w, r, Cfg) })
	http.HandleFunc("/sanamente/chemonics/sede", func(w http.ResponseWriter, r *http.Request) { Obtenersedechemonics(w, r, Cfg) })
	http.HandleFunc("/sanamente/chemonics/post", func(w http.ResponseWriter, r *http.Request) { postEncuestaSMCHemonics(w, r, Cfg) })
	http.HandleFunc("/sanamente/chemonics", func(w http.ResponseWriter, r *http.Request) { EncuestaSMChemonics(w, r, Cfg) })
	http.HandleFunc("/mdv/seguimientoopciones", func(w http.ResponseWriter, r *http.Request) { obtenerseguimientotipo(w, r) })
	http.HandleFunc("/obtenerdepsedesanamente2", func(w http.ResponseWriter, r *http.Request) { obtenerdepsedeHandlersanamente2(w, r, Cfg) })
	http.HandleFunc("/obtenermunsedesanamente2", func(w http.ResponseWriter, r *http.Request) { obtenermunsedeHandlersanamente2(w, r, Cfg) })
	http.HandleFunc("/obtenersedesanamente2", func(w http.ResponseWriter, r *http.Request) { obtenersedeHandlersanamente2(w, r, Cfg) })
	http.HandleFunc("/sanamente/encuesta/inicial/gt", func(w http.ResponseWriter, r *http.Request) { EncuestaSMInicial(w, r, Cfg) })
	http.HandleFunc("/sanamente/encuesta/final/gt", func(w http.ResponseWriter, r *http.Request) { EncuestaSMFinal(w, r, Cfg) })
	http.HandleFunc("/sanamente/encuesta/inicial/post", func(w http.ResponseWriter, r *http.Request) { postEncuestaSMInicial(w, r, Cfg) })
	http.HandleFunc("/sanamente/encuesta/final/post", func(w http.ResponseWriter, r *http.Request) { postEncuestaSMFinal(w, r, Cfg) })
	http.HandleFunc("/juventud/liderazgo", func(w http.ResponseWriter, r *http.Request) { LiderazgoRedJuvenil(w, r, Cfg) })
	http.HandleFunc("/juventud/liderazgo/post", func(w http.ResponseWriter, r *http.Request) { PostLiderazgoRedJuvenil(w, r, Cfg) })
	http.HandleFunc("/obtenersedes", func(w http.ResponseWriter, r *http.Request) { ObtenersedeGeneral(w, r, Cfg) })
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getOpcionesPreguntaNew(db *sql.DB, Query string) ([]Respuestas, error) {
	rows, err := db.Query(Query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var opciones []Respuestas
	for rows.Next() {
		var value, label, idPregunta string
		err := rows.Scan(&value, &label, &idPregunta)
		if err != nil {
			return nil, err
		}
		opciones = append(opciones, Respuestas{Value: value, Label: label, IdPregunta: idPregunta})
	}

	return opciones, nil
}

func satisfaccionVoluntariado(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db2 *sql.DB
	connString2 := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)

	db2, err := sql.Open("sqlserver", connString2)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx2 := context.Background()
	err = db2.PingContext(ctx2)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db2.Close()

	// Obtener opciones de la base de datos de SQL
	preguntas, err := getOpcionesPreguntaNew(db2, "SELECT distinct [id], [Pregunta], [idformulario] FROM [Voluntariado].[dbo].[Preguntas_formularios] where [idformulario]=1 and [version]=1 and [deleted] is null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	solucion, err := getOpcionesPreguntaNew(db2, "SELECT distinct [id], [Respuesta], [idPregunta] FROM [Voluntariado].[dbo].[respuesta_vol] where [idformulario]=1 and [version]=1 and [deleted] is null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormDataNew{
		Solucion:  solucion,
		Preguntas: preguntas,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Voluntarios/SatisfaccionVoluntariado.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func encuestainicialvoluntariado(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db2 *sql.DB
	connString2 := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)

	db2, err := sql.Open("sqlserver", connString2)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx2 := context.Background()
	err = db2.PingContext(ctx2)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db2.Close()

	// Obtener opciones de la base de datos de SQL
	preguntas, err := getOpcionesPreguntaNew(db2, "SELECT distinct [id], [Pregunta], [idformulario] FROM [Voluntariado].[dbo].[Preguntas_formularios] where [idformulario]=1 and [version]=1 and [deleted] is null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	solucion, err := getOpcionesPreguntaNew(db2, "SELECT distinct [id], [Respuesta], [idPregunta] FROM [Voluntariado].[dbo].[respuesta_vol] where [idformulario]=1 and [version]=1 and [deleted] is null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	solucionedu, err := getOpcionesPreguntaNew(db2, "SELECT * FROM (SELECT DISTINCT id, Respuesta, idPregunta FROM Voluntariado.dbo.respuesta_vol WHERE idformulario = 1 AND [version] = 1 AND idPregunta = 6 AND id IN (27, 136, 29, 41, 42) AND deleted IS NULL) AS sub ORDER BY CASE id WHEN 27 THEN 1 WHEN 136 THEN 2 WHEN 29 THEN 3 WHEN 41 THEN 4 WHEN 42 THEN 5 END;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	solucionsalud, err := getOpcionesPreguntaNew(db2, "SELECT * FROM (SELECT DISTINCT id, Respuesta, idPregunta FROM Voluntariado.dbo.respuesta_vol WHERE idformulario = 1 AND [version] = 1 AND idPregunta = 6 AND id IN (35,38,42) AND deleted IS NULL) AS sub ORDER BY CASE id WHEN 35 THEN 1 WHEN 38 THEN 2 WHEN 42 THEN 3 END;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	solucionsa, err := getOpcionesPreguntaNew(db2, "SELECT * FROM (SELECT DISTINCT id, Respuesta, idPregunta FROM Voluntariado.dbo.respuesta_vol WHERE idformulario = 1 AND [version] = 1 AND idPregunta = 6 AND id IN (137, 138, 42) AND deleted IS NULL) AS sub ORDER BY CASE id WHEN 137 THEN 1 WHEN 138 THEN 2 WHEN 42 THEN 3 END;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormDataNew{
		Solucion:      solucion,
		Preguntas:     preguntas,
		Solucionedu:   solucionedu,
		Solucionsalud: solucionsalud,
		Solucionsa:    solucionsa,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Voluntarios/EncuestaInicialVol.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func encuestafinalvoluntariado(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db2 *sql.DB
	connString2 := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)

	db2, err := sql.Open("sqlserver", connString2)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx2 := context.Background()
	err = db2.PingContext(ctx2)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db2.Close()

	// Obtener opciones de la base de datos de SQL
	preguntas, err := getOpcionesPreguntaNew(db2, "SELECT distinct [id], [Pregunta], [idformulario] FROM [Voluntariado].[dbo].[Preguntas_formularios] where [idformulario]=1 and [version]=1 and [deleted] is null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	solucion, err := getOpcionesPreguntaNew(db2, "SELECT distinct [id], [Respuesta], [idPregunta] FROM [Voluntariado].[dbo].[respuesta_vol] where [idformulario]=1 and [version]=1 and [deleted] is null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	solucionedu, err := getOpcionesPreguntaNew(db2, "SELECT * FROM (SELECT DISTINCT id, Respuesta, idPregunta FROM Voluntariado.dbo.respuesta_vol WHERE idformulario = 1 AND [version] = 1 AND idPregunta = 6 AND id IN (27, 136, 29, 41, 42) AND deleted IS NULL) AS sub ORDER BY CASE id WHEN 27 THEN 1 WHEN 136 THEN 2 WHEN 29 THEN 3 WHEN 41 THEN 4 WHEN 42 THEN 5 END;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	solucionsalud, err := getOpcionesPreguntaNew(db2, "SELECT * FROM (SELECT DISTINCT id, Respuesta, idPregunta FROM Voluntariado.dbo.respuesta_vol WHERE idformulario = 1 AND [version] = 1 AND idPregunta = 6 AND id IN (35,38,42) AND deleted IS NULL) AS sub ORDER BY CASE id WHEN 35 THEN 1 WHEN 38 THEN 2 WHEN 42 THEN 3 END;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	solucionsa, err := getOpcionesPreguntaNew(db2, "SELECT * FROM (SELECT DISTINCT id, Respuesta, idPregunta FROM Voluntariado.dbo.respuesta_vol WHERE idformulario = 1 AND [version] = 1 AND idPregunta = 6 AND id IN (137, 138, 42) AND deleted IS NULL) AS sub ORDER BY CASE id WHEN 137 THEN 1 WHEN 138 THEN 2 WHEN 42 THEN 3 END;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormDataNew{
		Solucion:      solucion,
		Preguntas:     preguntas,
		Solucionedu:   solucionedu,
		Solucionsalud: solucionsalud,
		Solucionsa:    solucionsa,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Voluntarios/EncuestaFinalVol.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func postencuestasatisfaccionvol(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)

	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	// Obtener los valores de los campos del formulario
	pais := cambio(r.FormValue("p1"))
	lugar := cambio(r.FormValue("p2"))
	otrolugar := r.FormValue("otro1")
	otroprograma := r.FormValue("otro2")
	edad := cambio(r.FormValue("p4"))
	sexo := cambio(r.FormValue("p5"))
	tipovoluntario := cambio(r.FormValue("p6"))
	otraactividad := r.FormValue("otro3")
	mediointeres := cambio(r.FormValue("p8"))
	otromedio := r.FormValue("otro4")
	evaluaproceso := cambio(r.FormValue("p9"))
	razonproceso := r.FormValue("xq5")
	tiemporespuesta := cambio(r.FormValue("p37"))
	razontiempo := r.FormValue("xq8")
	evaluacomunicacion := cambio(r.FormValue("p38"))
	razonevaluacioncom := r.FormValue("xq9")
	evaluacontacto := cambio(r.FormValue("p39"))
	razonevaluacontacto := r.FormValue("xq10")
	tiempoinformacion := cambio(r.FormValue("p12"))
	tiempoorientacion := cambio(r.FormValue("p15"))
	evaluacapOGySPI := cambio(r.FormValue("p40"))
	razoncapOGySPI := r.FormValue("xq11")
	evaluacapTecnica := cambio(r.FormValue("p41"))
	razonevaluacapTec := r.FormValue("xq12")
	kitinduccion := cambio(r.FormValue("p19"))
	frasesemana := cambio(r.FormValue("p20"))
	seguimiento := cambio(r.FormValue("p21"))
	areasapoyo := r.FormValue("apoyo")
	evaluaseguimiento := cambio(r.FormValue("p23"))
	razonevaluaseg := r.FormValue("xq6")
	evaluaguia := cambio(r.FormValue("p24"))
	razonevaluaguia := r.FormValue("xq7")
	materiales := cambio(r.FormValue("p25"))
	frasehorario := cambio(r.FormValue("p26"))
	herramientas := cambio(r.FormValue("p27"))
	razonherramientas := r.FormValue("xq1")
	otroapoyoarea := r.FormValue("otro5")
	satisfaccion := cambio(r.FormValue("p29"))
	razonsatisfaccion := r.FormValue("xq2")
	contribucion := r.FormValue("contribucion")
	fraseopinion := cambio(r.FormValue("p31"))
	aprendizaje := r.FormValue("aprendizaje")
	repetiriavol := cambio(r.FormValue("p32"))
	razonrepetiriavol := r.FormValue("xq3")
	recomendacion := cambio(r.FormValue("p33"))
	razonrecomendacion := r.FormValue("xq4")
	comentarios := r.FormValue("coment")

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente

	stmt, err := tx.Prepare("INSERT INTO [Voluntariado].[dbo].[F_encuestaSatisfaccionVol]([pais],[lugar],[otrolugar],[otroprograma] ,[edad] ,[sexo] ,[tipovoluntario] ,[otraactividad] ,[mediointeres] ,[otromedio] ,[evaluaproceso] ,[razonproceso] ,[tiemporespuesta], [razontiempo] ,[evaluacomunicacion], [razonevaluacioncom] ,[evaluacontacto], [razonevaluacontacto] ,[tiempoinformacion] ,[tiempoorientacion] ,[evaluacapOGySPI], [razoncapOGySPI] ,[evaluacapTecnica], [razonevaluacapTec] ,[kitinduccion] ,[frasesemana] ,[seguimiento] ,[areasapoyo] ,[evaluaseguimiento], [razonevaluaseg] ,[evaluaguia], [razonevaluaguia] ,[materiales] ,[frasehorario] ,[herramientas] ,[razonherramientas] ,[otroapoyoarea] ,[satisfaccion] ,[razonsatisfaccion] ,[contribucion] ,[fraseopinion] ,[aprendizaje] ,[repetiriavol] ,[razonrepetiriavol] ,[recomendacion] ,[razonrecomendacion] ,[comentarios]) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14,@P15,@P16,@P17,@P18,@P19,@P20,@P21,@P22,@P23,@P24,@P25,@P26,@P27,@P28,@P29,@P30,@P31,@P32,@P33,@P34,@P35,@P36,@P37,@P38,@P39,@P40,@P41,@P42,@P43,@P44,@P45,@P46,@P47)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	var lastInsertId int
	err = stmt.QueryRow(pais, lugar, otrolugar, otroprograma, edad, sexo, tipovoluntario, otraactividad, mediointeres, otromedio, evaluaproceso, razonproceso, tiemporespuesta, razontiempo, evaluacomunicacion, razonevaluacioncom, evaluacontacto, razonevaluacontacto, tiempoinformacion, tiempoorientacion, evaluacapOGySPI, razoncapOGySPI, evaluacapTecnica, razonevaluacapTec, kitinduccion, frasesemana, seguimiento, areasapoyo, evaluaseguimiento, razonevaluaseg, evaluaguia, razonevaluaguia, materiales, frasehorario, herramientas, razonherramientas, otroapoyoarea, satisfaccion, razonsatisfaccion, contribucion, fraseopinion, aprendizaje, repetiriavol, razonrepetiriavol, recomendacion, razonrecomendacion, comentarios).Scan(&lastInsertId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	areaprograma := r.Form["p3"]
	if len(areaprograma) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO [Voluntariado].[dbo].[F_multipleVol]([idPregunta],[idFR],[idRespuesta],[idFormulario]) VALUES (@P1, @P2, @P3, @P4)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta programa del voluntario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range areaprograma {
			_, err = stmt2.Exec(4, lastInsertId, valor, 1)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta programa del voluntario", http.StatusInternalServerError)
				return
			}
		}
	}

	actividad := r.Form["p7"]
	if len(actividad) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO [Voluntariado].[dbo].[F_multipleVol]([idPregunta],[idFR],[idRespuesta],[idFormulario]) VALUES (@P1, @P2, @P3, @P4)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta actividades del voluntario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range actividad {
			_, err = stmt2.Exec(9, lastInsertId, valor, 1)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta actividades del voluntario", http.StatusInternalServerError)
				return
			}
		}
	}

	capacitaciones := r.Form["c15"]
	if len(capacitaciones) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO [Voluntariado].[dbo].[F_multipleVol]([idPregunta],[idFR],[idRespuesta],[idFormulario]) VALUES (@P1, @P2, @P3, @P4)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta capacitaciones del voluntario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range capacitaciones {
			_, err = stmt2.Exec(18, lastInsertId, valor, 1)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta capacitaciones del voluntario", http.StatusInternalServerError)
				return
			}
		}
	}

	areaformacion := r.Form["p28"]
	if len(areaformacion) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO [Voluntariado].[dbo].[F_multipleVol]([idPregunta],[idFR],[idRespuesta],[idFormulario]) VALUES (@P1, @P2, @P3, @P4)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta formaciones del voluntario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range areaformacion {
			_, err = stmt2.Exec(32, lastInsertId, valor, 1)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta formaciones del voluntario", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	t, err := template.ParseFiles("public/HTML/RespuestaEncuestaSatisVol.html")
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

func postencuestainicialvol(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)

	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	// Obtener los valores de los campos del formulario
	pais := cambio(r.FormValue("p61"))
	lugar := cambio(r.FormValue("p2"))
	otrolugar := r.FormValue("otro1")
	otroprograma := r.FormValue("otro2")
	edad := cambio(r.FormValue("p4"))
	sexo := cambio(r.FormValue("p5"))
	tipovoluntario := cambio(r.FormValue("p6"))
	otraactividad := r.FormValue("otro3")
	medioinscripcion := cambio(r.FormValue("p8"))
	otromedio := r.FormValue("otro4")
	primeravez := cambio(r.FormValue("p54"))
	tiempoparticipacion := r.FormValue("p55")
	evaluaproceso := cambio(r.FormValue("p56"))
	razonproceso := r.FormValue("xq5")
	tiemporespuesta := cambio(r.FormValue("p37"))
	razontiempo := r.FormValue("xq8")
	evaluacomunicacion := cambio(r.FormValue("p38"))
	razonevaluacioncom := r.FormValue("xq9")
	evaluacontacto := cambio(r.FormValue("p39"))
	razonevaluacontacto := r.FormValue("xq10")
	tiempoinformacion := cambio(r.FormValue("p12"))
	tiempoorientacion := cambio(r.FormValue("p15"))
	evaluacapOGySPI := cambio(r.FormValue("p40"))
	razoncapOGySPI := r.FormValue("xq11")
	evaluacapTecnica := cambio(r.FormValue("p41"))
	razonevaluacapTec := r.FormValue("xq12")
	kitinduccion := cambio(r.FormValue("p19"))
	frasesemana := cambio(r.FormValue("p20"))
	seguimiento := cambio(r.FormValue("p21"))
	areasapoyo := r.FormValue("apoyo")
	// evaluaseguimiento := cambio(r.FormValue("p23"))
	// razonevaluaseg := r.FormValue("xq6")
	// evaluaguia := cambio(r.FormValue("p24"))
	// razonevaluaguia := r.FormValue("xq7")
	// materiales := cambio(r.FormValue("p25"))
	frasehorario := cambio(r.FormValue("p58"))
	herramientas := cambio(r.FormValue("p59"))
	razonherramientas := r.FormValue("xq1")
	comentarios := r.FormValue("p36")
	experienciainicial := r.FormValue("p63")

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente

	stmt, err := tx.Prepare("INSERT INTO [Voluntariado].[dbo].[F_encuestaInicialVol]([pais],[lugar],[otrolugar],[otroprograma],[edad],[sexo],[tipovoluntario],[otraactividad],[medioinscripcion],[otromedio],[primeravez],[tiempoparticipacion],[evaluaproceso],[razonproceso],[tiemporespuesta],[razontiempo],[evaluacomunicacion],[razonevaluacioncom],[evaluacontacto],[razonevaluacontacto],[tiempoinformacion],[tiempoorientacion],[evaluacapOGySPI],[razoncapOGySPI],[evaluacapTecnica],[razonevaluacapTec],[kitinduccion],[frasesemana],[seguimiento],[areasapoyo],[frasehorario],[herramientas],[razonherramientas],[comentarios],experienciainicial) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14,@P15,@P16,@P17,@P18,@P19,@P20,@P21,@P22,@P23,@P24,@P25,@P26,@P27,@P28,@P29,@P30,@P31,@P32,@P33,@P34,@P35)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	var lastInsertId int
	err = stmt.QueryRow(pais, lugar, otrolugar, otroprograma, edad, sexo, tipovoluntario, otraactividad, medioinscripcion, otromedio, primeravez, tiempoparticipacion, evaluaproceso, razonproceso, tiemporespuesta, razontiempo, evaluacomunicacion, razonevaluacioncom, evaluacontacto, razonevaluacontacto, tiempoinformacion, tiempoorientacion, evaluacapOGySPI, razoncapOGySPI, evaluacapTecnica, razonevaluacapTec, kitinduccion, frasesemana, seguimiento, areasapoyo, frasehorario, herramientas, razonherramientas, comentarios, experienciainicial).Scan(&lastInsertId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	areaprograma := r.Form["p3"]
	if len(areaprograma) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO [Voluntariado].[dbo].[F_multipleVol]([idPregunta],[idFR],[idRespuesta],[idFormulario]) VALUES (@P1, @P2, @P3, @P4)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta programa del voluntario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range areaprograma {
			_, err = stmt2.Exec(4, lastInsertId, valor, 3)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta programa del voluntario", http.StatusInternalServerError)
				return
			}
		}
	}

	actividad := r.Form["p62"]
	if len(actividad) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO [Voluntariado].[dbo].[F_multipleVol]([idPregunta],[idFR],[idRespuesta],[idFormulario]) VALUES (@P1, @P2, @P3, @P4)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta actividades del voluntario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range actividad {
			_, err = stmt2.Exec(9, lastInsertId, valor, 3)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta actividades del voluntario", http.StatusInternalServerError)
				return
			}
		}
	}

	capacitaciones := r.Form["c15"]
	if len(capacitaciones) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO [Voluntariado].[dbo].[F_multipleVol]([idPregunta],[idFR],[idRespuesta],[idFormulario]) VALUES (@P1, @P2, @P3, @P4)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta capacitaciones del voluntario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range capacitaciones {
			_, err = stmt2.Exec(24, lastInsertId, valor, 3)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta capacitaciones del voluntario", http.StatusInternalServerError)
				return
			}
		}
	}

	areaformacion := r.Form["p28"]
	if len(areaformacion) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO [Voluntariado].[dbo].[F_multipleVol]([idPregunta],[idFR],[idRespuesta],[idFormulario]) VALUES (@P1, @P2, @P3, @P4)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta formaciones del voluntario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range areaformacion {
			_, err = stmt2.Exec(42, lastInsertId, valor, 3)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta formaciones del voluntario", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	t, err := template.ParseFiles("public/HTML/RespuestaEncuestaInicialVol.html")
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

func postencuestafinalvol(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)

	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	// Obtener los valores de los campos del formulario
	pais := cambio(r.FormValue("p1"))
	lugar := cambio(r.FormValue("p2"))
	otrolugar := r.FormValue("otro1")
	otroprograma := r.FormValue("otro2")
	edad := cambio(r.FormValue("p4"))
	sexo := cambio(r.FormValue("p5"))
	tipovoluntario := cambio(r.FormValue("p6"))
	otraactividad := r.FormValue("otro3")
	satisfaccion := cambio(r.FormValue("p29"))
	razonsatisfaccion := r.FormValue("xq5")
	contribucion := r.FormValue("contribucion")
	fraseopinion := cambio(r.FormValue("p31"))
	aprendizaje := r.FormValue("aprendizaje")
	repetiriavol := cambio(r.FormValue("p32"))
	razonrepetiriavol := r.FormValue("xq3")
	recomendacion := cambio(r.FormValue("p33"))
	razonrecomendacion := r.FormValue("xq4")
	comentarios := r.FormValue("coment")

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente

	stmt, err := tx.Prepare("INSERT INTO [Voluntariado].[dbo].[F_encuestaFinalVol]([pais],[lugar],[otrolugar],[otroprograma],[edad],[sexo],[tipovoluntario],[otraactividad],[satisfaccion],[razonsatisfaccion],[contribucion],[fraseopinion],[aprendizaje],[repetiriavol],[razonrepetiriavol],[recomendacion],[razonrecomendacion],[comentarios]) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14,@P15,@P16,@P17,@P18)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return

	}

	var lastInsertId int
	err = stmt.QueryRow(pais, lugar, otrolugar, otroprograma, edad, sexo, tipovoluntario, otraactividad, satisfaccion, razonsatisfaccion, contribucion, fraseopinion, aprendizaje, repetiriavol, razonrepetiriavol, recomendacion, razonrecomendacion, comentarios).Scan(&lastInsertId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	areaprograma := r.Form["p3"]
	if len(areaprograma) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO [Voluntariado].[dbo].[F_multipleVol]([idPregunta],[idFR],[idRespuesta],[idFormulario]) VALUES (@P1, @P2, @P3, @P4)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta programa del voluntario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range areaprograma {
			_, err = stmt2.Exec(4, lastInsertId, valor, 4)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta programa del voluntario", http.StatusInternalServerError)
				return
			}
		}
	}

	actividad := r.Form["p7"]
	if len(actividad) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO [Voluntariado].[dbo].[F_multipleVol]([idPregunta],[idFR],[idRespuesta],[idFormulario]) VALUES (@P1, @P2, @P3, @P4)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta actividades del voluntario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range actividad {
			_, err = stmt2.Exec(9, lastInsertId, valor, 4)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta actividades del voluntario", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	t, err := template.ParseFiles("public/HTML/RespuestaEncuestaFinalVol.html")
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

func encuestavisitacampo(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db2 *sql.DB
	connString2 := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)

	db2, err := sql.Open("sqlserver", connString2)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx2 := context.Background()
	err = db2.PingContext(ctx2)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db2.Close()
	// Obtener opciones de la base de datos de SQL
	solucion, err := getOpcionesPreguntaNew(db2, "SELECT distinct id, Respuesta, idPregunta FROM [Voluntariado].[dbo].[vwRespuestasVisitaCampo]")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	preguntas, err := getOpcionesPreguntaNew(db2, "SELECT distinct id, Pregunta, idformulario FROM [Rescate].[admin_main_gwdata].[Preguntas_formularios] where idFormulario=5")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormDataNew{
		Solucion:  solucion,
		Preguntas: preguntas,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Voluntarios/EncuestaVisitaCampo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postencuestavisitacampo(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)

	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	// Obtener los valores de los campos del formulario
	fechavisita := r.FormValue("visita")
	pais := cambio(r.FormValue("p160"))
	programa := cambio(r.FormValue("p158"))
	actividad := cambio(r.FormValue("p86"))
	sede := cambio(r.FormValue("p84"))
	nombrevisitante := r.FormValue("encargado")
	cargovisitante := r.FormValue("cargo")
	facilitadores := r.FormValue("facilitadores")
	//sesion := r.FormValue("sesion")
	//participantesgwdata := cambio(r.FormValue("pgwdata"))
	//participantesobs := cambio(r.FormValue("pobs"))
	//voluntariosgwdata := cambio(r.FormValue("vgwdata"))
	//voluntariosobs := cambio(r.FormValue("vobs"))
	coberturaasp1 := cambio(r.FormValue("p88"))
	//coberturaasp2 := cambio(r.FormValue("p89"))
	//coberturaasp3 := cambio(r.FormValue("p90"))
	coberturaasp4 := cambio(r.FormValue("p91"))
	coberturaasp5 := cambio(r.FormValue("p92"))
	coberturaobs1 := r.FormValue("coobs1")
	coberturaobs2 := r.FormValue("coobs2")
	//coberturaobs3 := r.FormValue("coobs3")
	coberturaobs4 := r.FormValue("coobs4")
	coberturaobs5 := r.FormValue("coobs5")
	calidadasp1 := cambio(r.FormValue("p93"))
	calidadasp2 := cambio(r.FormValue("p94"))
	calidadasp3 := cambio(r.FormValue("p95"))
	calidadasp4 := cambio(r.FormValue("p96"))
	//calidadasp5 := cambio(r.FormValue("p97"))
	calidadasp6 := cambio(r.FormValue("p98"))
	//calidadasp7 := cambio(r.FormValue("p99"))
	calidadasp8 := cambio(r.FormValue("p100"))
	calidadasp9 := cambio(r.FormValue("p101"))
	calidadasp10 := cambio(r.FormValue("p102"))
	calidadasp11 := cambio(r.FormValue("p103"))
	calidadasp12 := cambio(r.FormValue("p104"))
	calidadasp13 := cambio(r.FormValue("p105"))
	//calidadasp14 := cambio(r.FormValue("p106"))
	calidadasp15 := cambio(r.FormValue("p107"))
	calidadasp16 := cambio(r.FormValue("p108"))
	calidadasp17 := cambio(r.FormValue("p109"))
	calidadasp18 := cambio(r.FormValue("p110"))
	calidadasp19 := cambio(r.FormValue("p111"))
	calidadasp20 := cambio(r.FormValue("p112"))
	calidadasp21 := cambio(r.FormValue("p113"))
	calidadasp22 := cambio(r.FormValue("p114"))
	calidadasp23 := cambio(r.FormValue("p115"))
	calidadasp24 := cambio(r.FormValue("p116"))
	calidadasp25 := cambio(r.FormValue("p117"))
	calidadasp26 := cambio(r.FormValue("p118"))
	calidadasp27 := cambio(r.FormValue("p119"))
	calidadasp28 := cambio(r.FormValue("p121"))
	calidadasp29 := cambio(r.FormValue("p122"))
	calidadasp30 := cambio(r.FormValue("p123"))
	calidadobs1 := r.FormValue("caobs1")
	//calidadobs2 := r.FormValue("caobs2")
	calidadobs3 := r.FormValue("caobs3")
	calidadobs4 := r.FormValue("caobs4")
	//calidadobs5 := r.FormValue("caobs5")
	calidadobs6 := r.FormValue("caobs6")
	//calidadobs7 := r.FormValue("caobs7")
	calidadobs8 := r.FormValue("caobs8")
	calidadobs9 := r.FormValue("caobs9")
	calidadobs10 := r.FormValue("caobs10")
	calidadobs11 := r.FormValue("caobs11")
	calidadobs12 := r.FormValue("caobs12")
	calidadobs13 := r.FormValue("caobs13")
	//calidadobs14 := r.FormValue("caobs14")
	calidadobs15 := r.FormValue("caobs15")
	calidadobs16 := r.FormValue("caobs16")
	calidadobs17 := r.FormValue("caobs17")
	calidadobs18 := r.FormValue("caobs18")
	calidadobs19 := r.FormValue("caobs19")
	calidadobs20 := r.FormValue("caobs20")
	calidadobs21 := r.FormValue("caobs21")
	calidadobs22 := r.FormValue("caobs22")
	calidadobs23 := r.FormValue("caobs23")
	calidadobs24 := r.FormValue("caobs24")
	calidadobs25 := r.FormValue("caobs25")
	calidadobs26 := r.FormValue("caobs26")
	calidadobs27 := r.FormValue("caobs27")
	calidadobs28 := r.FormValue("caobs28")
	calidadobs29 := r.FormValue("caobs29")
	calidadobs30 := r.FormValue("caobs30")
	eseguridadasp1 := cambio(r.FormValue("p124"))
	eseguridadasp2 := cambio(r.FormValue("p188")) //departamento
	eseguridadasp3 := cambio(r.FormValue("actividad"))
	eseguridadasp4 := cambio(r.FormValue("p127"))
	//eseguridadasp5 := cambio(r.FormValue("p128"))
	eseguridadobs1 := r.FormValue("esobs1")
	//eseguridadobs2 := r.FormValue("esobs2")
	//eseguridadobs3 := r.FormValue("esobs3")
	eseguridadobs4 := r.FormValue("esobs4")
	//eseguridadobs5 := r.FormValue("esobs5")
	proteccionasp1 := cambio(r.FormValue("p129"))
	proteccionasp2 := cambio(r.FormValue("p130"))
	proteccionasp3 := cambio(r.FormValue("p131"))
	proteccionasp4 := cambio(r.FormValue("p132"))
	proteccionasp5 := cambio(r.FormValue("p133"))
	proteccionasp6 := cambio(r.FormValue("p134"))
	//proteccionasp7 := cambio(r.FormValue("p135"))
	proteccionobs1 := r.FormValue("piobs1")
	proteccionobs2 := r.FormValue("piobs2")
	proteccionobs3 := r.FormValue("piobs3")
	proteccionobs4 := r.FormValue("piobs4")
	proteccionobs5 := r.FormValue("piobs5")
	proteccionobs6 := r.FormValue("piobs6")
	//proteccionobs7 := r.FormValue("piobs7")
	solicitudes := r.FormValue("comentarios")
	retroalimentacion := r.FormValue("recomendaciones")

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO [Voluntariado].[dbo].[F_visitacampo] (fechavisita,pais,programa,actividad,sede,nombrevisitante,cargovisitante,facilitadores,coberturaasp1,coberturaasp4,coberturaasp5,coberturaobs1,coberturaobs2,coberturaobs4,coberturaobs5,calidadasp1,calidadasp2,calidadasp3,calidadasp4,calidadasp6,calidadasp8,calidadasp9,calidadasp10,calidadasp11,calidadasp12,calidadasp13,calidadasp15,calidadasp16,calidadasp17,calidadasp18,calidadasp19,calidadasp20,calidadasp21,calidadasp22,calidadasp23,calidadasp24,calidadasp25,calidadasp26,calidadasp27,calidadasp28,calidadasp29,calidadasp30,calidadobs1,calidadobs3,calidadobs4,calidadobs6,calidadobs8,calidadobs9,calidadobs10,calidadobs11,calidadobs12,calidadobs13,calidadobs15,calidadobs16,calidadobs17,calidadobs18,calidadobs19,calidadobs20,calidadobs21,calidadobs22,calidadobs23,calidadobs24,calidadobs25,calidadobs26,calidadobs27,calidadobs28,calidadobs29,calidadobs30,eseguridadasp1,eseguridadasp2,eseguridadasp3,eseguridadasp4,eseguridadobs1,eseguridadobs4,proteccionasp1,proteccionasp2,proteccionasp3,proteccionasp4,proteccionasp5,proteccionasp6,proteccionobs1,proteccionobs2,proteccionobs3,proteccionobs4,proteccionobs5,proteccionobs6,solicitudes,retroalimentacion) OUTPUT INSERTED.ID VALUES (@p1,@p2,@p3,@p4,@p5,@p6,@p7,@p8,@p9,@p10,@p11,@p12,@p13,@p14,@p15,@p16,@p17,@p18,@p19,@p20,@p21,@p22,@p23,@p24,@p25,@p26,@p27,@p28,@p29,@p30,@p31,@p32,@p33,@p34,@p35,@p36,@p37,@p38,@p39,@p40,@p41,@p42,@p43,@p44,@p45,@p46,@p47,@p48,@p49,@p50,@p51,@p52,@p53,@p54,@p55,@p56,@p57,@p58,@p59,@p60,@p61,@p62,@p63,@p64,@p65,@p66,@p67,@p68,@p69,@p70,@p71,@p72,@p73,@p74,@p75,@p76,@p77,@p78,@p79,@p80,@p81,@p82,@p83,@p84,@p85,@p86,@p87,@p88)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(fechavisita, pais, programa, actividad, sede, nombrevisitante, cargovisitante, facilitadores, coberturaasp1, coberturaasp4, coberturaasp5, coberturaobs1, coberturaobs2, coberturaobs4, coberturaobs5, calidadasp1, calidadasp2, calidadasp3, calidadasp4, calidadasp6, calidadasp8, calidadasp9, calidadasp10, calidadasp11, calidadasp12, calidadasp13, calidadasp15, calidadasp16, calidadasp17, calidadasp18, calidadasp19, calidadasp20, calidadasp21, calidadasp22, calidadasp23, calidadasp24, calidadasp25, calidadasp26, calidadasp27, calidadasp28, calidadasp29, calidadasp30, calidadobs1, calidadobs3, calidadobs4, calidadobs6, calidadobs8, calidadobs9, calidadobs10, calidadobs11, calidadobs12, calidadobs13, calidadobs15, calidadobs16, calidadobs17, calidadobs18, calidadobs19, calidadobs20, calidadobs21, calidadobs22, calidadobs23, calidadobs24, calidadobs25, calidadobs26, calidadobs27, calidadobs28, calidadobs29, calidadobs30, eseguridadasp1, eseguridadasp2, eseguridadasp3, eseguridadasp4, eseguridadobs1, eseguridadobs4, proteccionasp1, proteccionasp2, proteccionasp3, proteccionasp4, proteccionasp5, proteccionasp6, proteccionobs1, proteccionobs2, proteccionobs3, proteccionobs4, proteccionobs5, proteccionobs6, solicitudes, retroalimentacion)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	r.ParseForm()

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func encuestadell(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}

	// Obtener opciones de la base de datos de MYSQL
	defer db.Close()

	solucion, err := getOpcionesPreguntaNew(db, "SELECT distinct id, Respuesta, idPregunta FROM admin_main_gwdata.respuestasencuestadell")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	preguntas, err := getOpcionesPreguntaNew(db, "SELECT distinct id, Pregunta, idformulario FROM admin_main_gwdata.Preguntas_formularios where idFormulario=8")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormDataNew{
		Solucion:  solucion,
		Preguntas: preguntas,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Otros/encuestaDell.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postencuestadell(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	nombrecompleto := r.FormValue("nombrecompleto")
	fechaNac := r.FormValue("fechaNac")
	sexo := cambio(r.FormValue("grupoRadiio"))
	funcional := cambio(r.FormValue("funcional"))
	confianza := cambio(r.FormValue("confianza"))
	afirmacion := cambio(r.FormValue("gruporadio"))
	otraafirmacion := r.FormValue("otrar")
	mejoras := r.FormValue("mejoras")
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_encuestaDell (nombrecompleto, fechaNac, sexo, funcional, confianza, afirmacion, otraafirmacion, mejoras, fechaRegistro) VALUES (?,?,?,?,?,?,?,?,?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(nombrecompleto, fechaNac, sexo, funcional, confianza, afirmacion, otraafirmacion, mejoras, currentTimeGMT)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	r.ParseForm()

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func encuestapaliativaJLI(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}

	// Obtener opciones de la base de datos de MYSQL

	defer db.Close()

	solucion, err := getOpcionesPreguntaNew(db, "SELECT distinct id, Respuesta, idPregunta FROM admin_main_gwdata.respuesta_vol")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	preguntas, err := getOpcionesPreguntaNew(db, "SELECT distinct id, Pregunta, idformulario FROM admin_main_gwdata.Preguntas_formularios where idFormulario=1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormDataNew{
		Solucion:  solucion,
		Preguntas: preguntas,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/JLI/EncuestaPaliativosJLI.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postencuestapaliativaJLI(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	lugar := cambio(r.FormValue("p1"))
	otrolugar := r.FormValue("otro")
	edad := cambio(r.FormValue("p2"))
	sexo := cambio(r.FormValue("p3"))
	tipovoluntario := cambio(r.FormValue("p4"))
	areaprograma := cambio(r.FormValue("p5"))
	actividad := cambio(r.FormValue("p6"))
	pais := cambio(r.FormValue("p7"))
	medioinscripcion := cambio(r.FormValue("p8"))
	evaluaproceso := cambio(r.FormValue("p9"))
	tiemporespuesta := cambio(r.FormValue("p10"))
	evaluacomunicacion := cambio(r.FormValue("p11"))
	evaluacontacto := cambio(r.FormValue("p12"))
	tiempoinscripcion := cambio(r.FormValue("p13"))
	tiempoorientacion := cambio(r.FormValue("p14"))
	evaluaorientacion := cambio(r.FormValue("p16"))
	evaluaseguridad := cambio(r.FormValue("p17"))
	evaluacapacitacion := cambio(r.FormValue("p18"))
	kitinduccion := cambio(r.FormValue("p19"))
	frasesemana := cambio(r.FormValue("p20"))
	seguimiento := cambio(r.FormValue("p21"))
	tiposeguimiento := r.FormValue("p22")
	evaluaseguimiento := cambio(r.FormValue("p31"))
	evaluaguia := cambio(r.FormValue("p32"))
	materiales := cambio(r.FormValue("p33"))
	frasehorario := cambio(r.FormValue("p34"))
	herramientas := cambio(r.FormValue("p35"))
	razonherramientas := r.FormValue("xq1")
	areaformacion := cambio(r.FormValue("p36"))
	razonformacion := r.FormValue("otro2")
	satisfaccion := cambio(r.FormValue("p37"))
	percepcion := cambio(r.FormValue("p38"))
	razonpercepcion := r.FormValue("xq2")
	repetiriavol := cambio(r.FormValue("p39"))
	razonrepetiriavol := r.FormValue("xq3")
	recomendacion := cambio(r.FormValue("p40"))
	razonrecomendacion := r.FormValue("xq4")
	comentarios := r.FormValue("p41")
	fraseopinion := cambio(r.FormValue("p43"))
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_encuestaVol (lugar, otrolugar, edad, sexo, tipovoluntario, areaprograma, actividad, pais, medioinscripcion, evaluaproceso, tiemporespuesta, evaluacomunicacion, evaluacontacto, tiempoinscripcion, tiempoorientacion, evaluaorientacion, evaluaseguridad, evaluacapacitacion, kitinduccion, frasesemana, seguimiento, tiposeguimiento, evaluaseguimiento, evaluaguia, materiales, frasehorario, herramientas, razonherramientas, areaformacion, razonformacion, satisfaccion, percepcion, razonpercepcion, repetiriavol, razonrepetiriavol, recomendacion, razonrecomendacion, fraseopinion, comentarios, fechaRegistro) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(lugar, otrolugar, edad, sexo, tipovoluntario, areaprograma, actividad, pais, medioinscripcion, evaluaproceso, tiemporespuesta, evaluacomunicacion, evaluacontacto, tiempoinscripcion, tiempoorientacion, evaluaorientacion, evaluaseguridad, evaluacapacitacion, kitinduccion, frasesemana, seguimiento, tiposeguimiento, evaluaseguimiento, evaluaguia, materiales, frasehorario, herramientas, razonherramientas, areaformacion, razonformacion, satisfaccion, percepcion, razonpercepcion, repetiriavol, razonrepetiriavol, recomendacion, razonrecomendacion, fraseopinion, comentarios, currentTimeGMT)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	capacitacion2 := r.Form["p15"]
	if len(capacitacion2) == 0 {
	} else if len(capacitacion2) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO F_capacitacionvol (idp, capacitaciones) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta capacitaciones del voluntario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range capacitacion2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta capacitaciones del voluntario", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func encuestagwdata(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}

	// Obtener opciones de la base de datos de MYSQL
	defer db.Close()

	solucion, err := getOpcionesPreguntaNew(db, "SELECT distinct id, Respuesta, idPregunta FROM admin_main_gwdata.respuesta_gwdata")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	preguntas, err := getOpcionesPreguntaNew(db, "SELECT distinct id, Pregunta, idformulario FROM admin_main_gwdata.Preguntas_formularios where idFormulario=3")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormDataNew{
		Solucion:  solucion,
		Preguntas: preguntas,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Otros/encuestagwdata.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postencuestagwdata(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	pais := cambio(r.FormValue("p78"))
	area := cambio(r.FormValue("p44"))
	organizacion := cambio(r.FormValue("p45"))
	rol := cambio(r.FormValue("p46"))
	otrorol := r.FormValue("otrol")
	tieneusuario := cambio(r.FormValue("p47"))
	quiereusuario := cambio(r.FormValue("p50"))
	necesidad := r.FormValue("necesidad")
	razonnousuario := r.FormValue("otrazon")
	proceso1 := cambio(r.FormValue("p59"))
	sugerencia1 := r.FormValue("sg1")
	proceso2 := cambio(r.FormValue("p60"))
	sugerencia2 := r.FormValue("sg2")
	proceso3 := cambio(r.FormValue("p61"))
	sugerencia3 := r.FormValue("sg3")
	proceso4 := cambio(r.FormValue("p62"))
	sugerencia4 := r.FormValue("sg4")
	proceso5 := cambio(r.FormValue("p63"))
	sugerencia5 := r.FormValue("sg5")
	proceso6 := cambio(r.FormValue("p64"))
	sugerencia6 := r.FormValue("sg6")
	proceso7 := cambio(r.FormValue("p65"))
	sugerencia7 := r.FormValue("sg7")
	proceso8 := cambio(r.FormValue("p66"))
	sugerencia8 := r.FormValue("sg8")
	proceso9 := cambio(r.FormValue("p67"))
	sugerencia9 := r.FormValue("sg9")
	proceso10 := cambio(r.FormValue("p68"))
	sugerencia10 := r.FormValue("sg10")
	proceso11 := cambio(r.FormValue("p69"))
	sugerencia11 := r.FormValue("sg11")
	proceso12 := cambio(r.FormValue("p70"))
	sugerencia12 := r.FormValue("sg12")
	proceso13 := cambio(r.FormValue("p71"))
	sugerencia13 := r.FormValue("sg13")
	proceso14 := cambio(r.FormValue("p72"))
	sugerencia14 := r.FormValue("sg14")
	proceso15 := cambio(r.FormValue("p73"))
	sugerencia15 := r.FormValue("sg15")
	recomendaciona := r.FormValue("fn1")
	recomendacionb := r.FormValue("fn2")
	recomendacionc := r.FormValue("fn3")
	dominio := r.FormValue("fn4")
	capacitacion := cambio(r.FormValue("p58"))
	nota := cambio(r.FormValue("nota"))
	evaluaproceso := r.FormValue("xq")
	otrarecomendacion := r.FormValue("recomendacion")
	entrevista := cambio(r.FormValue("p57"))
	correo := r.FormValue("email")
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_encuestagwdata (pais,area,organizacion,rol,otrorol,tieneusuario,quiereusuario,necesidad,razonnousuario,proceso1,sugerencia1,proceso2,sugerencia2,proceso3,sugerencia3,proceso4,sugerencia4,proceso5,sugerencia5,proceso6,sugerencia6,proceso7,sugerencia7,proceso8,sugerencia8,proceso9,sugerencia9,proceso10,sugerencia10,proceso11,sugerencia11,proceso12,sugerencia12,proceso13,sugerencia13,proceso14,sugerencia14,proceso15,sugerencia15,recomendaciona,recomendacionb,recomendacionc,dominio,capacitacion,nota,evaluaproceso,otrarecomendacion,entrevista,correo,fechaRegistro) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(pais, area, organizacion, rol, otrorol, tieneusuario, quiereusuario, necesidad, razonnousuario, proceso1, sugerencia1, proceso2, sugerencia2, proceso3, sugerencia3, proceso4, sugerencia4, proceso5, sugerencia5, proceso6, sugerencia6, proceso7, sugerencia7, proceso8, sugerencia8, proceso9, sugerencia9, proceso10, sugerencia10, proceso11, sugerencia11, proceso12, sugerencia12, proceso13, sugerencia13, proceso14, sugerencia14, proceso15, sugerencia15, recomendaciona, recomendacionb, recomendacionc, dominio, capacitacion, nota, evaluaproceso, otrarecomendacion, entrevista, correo, currentTimeGMT)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	razonesno := r.Form["p52"]
	if len(razonesno) == 0 {
	} else if len(razonesno) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO F_razonesno (idp, razon) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta razones de no querer ser usuario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range razonesno {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta razones de no querer ser usuario", http.StatusInternalServerError)
				return
			}
		}
	}

	modocapacitacion := r.Form["p140"]
	print(modocapacitacion)
	if len(modocapacitacion) == 0 {
	} else if len(modocapacitacion) != 0 {
		stmt3, err := tx.Prepare("INSERT INTO F_modalcapacitacion (idp, capacitaciones) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta modalidades de capacitacion preferida", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range modocapacitacion {
			_, err = stmt3.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta modalidades de capacitacion preferida", http.StatusInternalServerError)
				return
			}
		}
	}

	dispositivos := r.Form["p48"]
	print(dispositivos)
	if len(dispositivos) == 0 {
	} else if len(dispositivos) != 0 {
		stmt4, err := tx.Prepare("INSERT INTO F_dipositivosgw (idp, dispositivos) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dispositivos que utiliza para GWData", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range dispositivos {
			_, err = stmt4.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta dispositivos que utiliza para GWData", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func casos(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}

	// Obtener opciones de la base de datos de MYSQL
	defer db.Close()

	solucion, err := getOpcionesPreguntaNew(db, "SELECT * FROM admin_main_gwdata.F_casos;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	preguntas, err := getOpcionesPreguntaNew(db, "SELECT distinct id, Pregunta, idformulario FROM admin_main_gwdata.Preguntas_formularios where idFormulario=2")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormDataNew{
		Solucion:  solucion,
		Preguntas: preguntas,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/sv/casos.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func formulariomentorias(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}

	defer db.Close()

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	transporte, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_transporte where activo=1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estipendio, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_estipendio")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estado_civ, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_estadociv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actividad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_ofertavol")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.Vista_Sedes where Idpais=7 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grupo, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_grupoSocial")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pueblo, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_puebloComunidad")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	idioma, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_idioma")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actLaboral, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_actividadLaboral")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Transporte:   transporte,
		Estipendio:   estipendio,
		Estado_civ:   estado_civ,
		Actividades:  actividad,
		Depsede:      depsede,
		Grupo:        grupo,
		Idioma:       idioma,
		ActLaboral:   actLaboral,
		Pueblo:       pueblo,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Mentorias/mentorias.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func capacitaciones(w http.ResponseWriter) {

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Otros/capacitaciones.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func voluntariadosv(w http.ResponseWriter, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()
	// Obtener opciones de la base de datos de  MySQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	camisa, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM volunteering_shirt_sizes where id not in(6,8)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paises, err := Controller.GetOpcionesPregunta(db, "SELECT distinct id, name FROM admin_main_gwdata.countries where id not in (7) order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, conceptoPerfil FROM type_beneficiarios where id in (5,6,9)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	relacioncom, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM type_beneficiarios where tipoVoluntario=3")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	institucion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM voluntary_institutions where country_id=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	copdonor, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_donantes where idp=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	programa, err := Controller.GetOpcionesPregunta(db, "SELECT distinct Ida, voluntariado FROM sede_areas where IdPais=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Crear estructura de datos con opciones para cada pregunta select
	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Camisa:       camisa,
		Paises:       paises,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
		Relacioncom:  relacioncom,
		Institucion:  institucion,
		Copdonor:     copdonor,
		Programa:     programa,
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Voluntarios/sv.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func voluntariadocol(w http.ResponseWriter, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()
	// Obtener opciones de la base de datos de  MySQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	camisa, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM volunteering_shirt_sizes where id not in(6,8)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paises, err := Controller.GetOpcionesPregunta(db, "SELECT distinct id, name FROM admin_main_gwdata.countries where id not in (16) order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=16")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, conceptoPerfil FROM type_beneficiarios where id in (5,6,9)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	relacioncom, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM type_beneficiarios where tipoVoluntario=3")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	copdonor, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_donantes where idp=16")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	programa, err := Controller.GetOpcionesPregunta(db, "SELECT distinct Ida, voluntariado FROM sede_areas where IdPais=7 and Ida not in (8)") //mientras se crean las sedes/escuelas para colombia
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Crear estructura de datos con opciones para cada pregunta select
	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Camisa:       camisa,
		Paises:       paises,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
		Relacioncom:  relacioncom,
		Copdonor:     copdonor,
		Programa:     programa,
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Voluntarios/col.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func voluntariadopn(w http.ResponseWriter, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	// Obtener opciones de la base de datos de  MySQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	camisa, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM volunteering_shirt_sizes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paises, err := Controller.GetOpcionesPregunta(db, "SELECT distinct id, name FROM admin_main_gwdata.countries where id not in (1) order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, conceptoPerfil FROM type_beneficiarios where id in (5,6,9)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	relacioncom, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM type_beneficiarios where tipoVoluntario=3")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	copdonor, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_donantes where idp=1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	programa, err := Controller.GetOpcionesPregunta(db, "SELECT distinct Ida, voluntariado FROM sede_areas where IdPais=7 and Ida not in (8)") //mientras crean las escuelas/sedes para Panamá
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Crear estructura de datos con opciones para cada pregunta select
	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Camisa:       camisa,
		Paises:       paises,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
		Relacioncom:  relacioncom,
		Copdonor:     copdonor,
		Programa:     programa,
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Voluntarios/pn.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func voluntariadord(w http.ResponseWriter, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	// Obtener opciones de la base de datos de  MySQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	camisa, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM volunteering_shirt_sizes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paises, err := Controller.GetOpcionesPregunta(db, "SELECT distinct id, name FROM admin_main_gwdata.countries where id not in (5) order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=5")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, conceptoPerfil FROM type_beneficiarios where id in (5,6,9)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	relacioncom, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM type_beneficiarios where tipoVoluntario=3")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	copdonor, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_donantes where idp=5")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	programa, err := Controller.GetOpcionesPregunta(db, "SELECT distinct Ida, voluntariado FROM sede_areas where IdPais=7 and Ida not in (8)") //mientras crean las escuelas/sedes para Panamá
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Crear estructura de datos con opciones para cada pregunta select
	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Camisa:       camisa,
		Paises:       paises,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
		Relacioncom:  relacioncom,
		Copdonor:     copdonor,
		Programa:     programa,
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Voluntarios/rd.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func voluntariadogt(w http.ResponseWriter, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	// Obtener opciones de la base de datos de  MySQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	camisa, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM volunteering_shirt_sizes where id not in(6,8)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paises, err := Controller.GetOpcionesPregunta(db, "SELECT distinct id, name FROM admin_main_gwdata.countries where id not in(4) order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=4")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, conceptoPerfil FROM type_beneficiarios where id in (5,6,9)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	relacioncom, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM type_beneficiarios where tipoVoluntario=3")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	institucion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM voluntary_institutions where country_id=4")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	copdonor, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_donantes where idp=4")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	programa, err := Controller.GetOpcionesPregunta(db, "SELECT distinct Ida, voluntariado FROM sede_areas where IdPais=4 and Ida not in (8,10)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Crear estructura de datos con opciones para cada pregunta select
	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Camisa:       camisa,
		Paises:       paises,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
		Relacioncom:  relacioncom,
		Institucion:  institucion,
		Copdonor:     copdonor,
		Programa:     programa,
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Voluntarios/gt.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func formulariomdvHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}

	defer db.Close()

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	transporte, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_transporte where activo=1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estipendio, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_estipendio")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estado_civ, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_estadociv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actividad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_ofertavol")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.Vista_Sedes where IdSede in(666,1128)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Transporte:   transporte,
		Estipendio:   estipendio,
		Estado_civ:   estado_civ,
		Actividades:  actividad,
		Depsede:      depsede,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/mdv/mdv.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func formulariojcHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}

	defer db.Close()

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	transporte, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_transporte where activo=1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estipendio, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_estipendio")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estado_civ, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_estadociv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actividad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_ofertavol")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.Vista_Sedes where IdSede in(666,1128)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Transporte:   transporte,
		Estipendio:   estipendio,
		Estado_civ:   estado_civ,
		Actividades:  actividad,
		Depsede:      depsede,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/mdv/jc.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func formularioParticipamas(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}

	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where typePerson=1 and active_at is not null and id not in(10)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/participamas/participamas.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func obtenerCiudadesHandler2(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	pais := r.FormValue("pais")
	var opcionesCiudades string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=? ", pais)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionesCiudades += ("<option value=''></option>")
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesCiudades += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionesCiudades)
}

func obtenerCiudadesHandler3(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	pais := r.FormValue("pais")
	var opcionesCiudades string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT distinct IdSede,Sede FROM vwSedesFormMentoriasGTE_HND where IdPais=? ", pais)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionesCiudades += ("<option value=''></option>")
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesCiudades += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}

	fmt.Fprint(w, opcionesCiudades)
}

func volcheck(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	pais := r.FormValue("pais")
	var opcionesCiudades string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT distinct numIdentVoluntario FROM F_Voluntariado where numIdentVoluntario=? and creacion>'2024-01-01'", pais)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var numIdentVoluntario string
		err := rows.Scan(&numIdentVoluntario)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesCiudades += numIdentVoluntario
	}

	fmt.Fprint(w, opcionesCiudades)
}

func obtenerCiudadesHandler(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	pais := r.FormValue("pais")
	var opcionesCiudades string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT distinct idm,municipio FROM admin_main_gwdata.geografia where idd=? ", pais)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionesCiudades += ("<option value=''></option>")
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesCiudades += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}

	fmt.Fprint(w, opcionesCiudades)
}

func obteneractividadHandler(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	area := r.FormValue("area")
	pais := r.FormValue("pais")
	var opcionesactividad string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT distinct idac, Actividad FROM admin_main_gwdata.actividadesvol WHERE ida=? and idPais=? and Actividad is not null order by Actividad ", area, pais)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionesactividad += "<option value=''></option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesactividad += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionesactividad)
}

func obtenerdep2Handler(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	area := r.FormValue("area")
	pais := r.FormValue("pais")
	var opcionesdep2 string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT Distinct iddep, Departamento FROM admin_main_gwdata.actividadesvol WHERE ida=? and idPais=? order by Departamento", area, pais)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionesdep2 += "<option value=''></option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesdep2 += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionesdep2)
}

func obtenersede2Handler(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	actividad := r.FormValue("actividad")
	area := r.FormValue("area")
	var opcionessubactividad string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT distinct IdSede, Sede FROM admin_main_gwdata.actividadesvol WHERE Iddep=? and ida=? order by Sede", actividad, area)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionessubactividad += "<option value=''></option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionessubactividad += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionessubactividad)
}

func obtenersubactividadHandler(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	actividad := r.FormValue("actividad")
	pais := r.FormValue("pais")
	var opcionessubactividad string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		opcionessubactividad = ""
	}
	defer db.Close()

	rows, err := db.Query("SELECT distinct idsuba,Subactividad FROM admin_main_gwdata.actividadesvol where idac=? and idPais=? order by Subactividad", actividad, pais)
	if err != nil {
		opcionessubactividad = ""
	}
	defer rows.Close()
	opcionessubactividad += "<option value=''></option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			opcionessubactividad = ""
		}
		opcionessubactividad += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionessubactividad)
}

func obtenermunsedeHandler(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	depsede := r.FormValue("depsede")
	var opcionesmunsede string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT Idmun,Municipio FROM admin_main_gwdata.Vista_Sedes where Iddep=? and IdSede in(666,1128) order by Municipio", depsede)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionesmunsede += "<option value=''></option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesmunsede += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionesmunsede)
}

func obtenersedeHandler(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	sede := r.FormValue("munsede")
	var opcionessede string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT IdSede,Sede FROM Vista_Sedes where Idmun=? and Idsede in(666,1128) order by Sede", sede)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionessede += "<option value=''></option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionessede += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionessede)
}

func obtenersedecampoHandler(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	sede := r.FormValue("pais")
	var opcionessede string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT IdSede,Sede FROM Vista_Sedes where Idpais=? order by Sede", sede)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionessede += "<option value=''></option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionessede += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionessede)
}

func formularioCorporativo(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}

	defer db.Close()

	var db2 *sql.DB
	connString2 := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)

	db2, err = sql.Open("sqlserver", connString2)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx2 := context.Background()
	err = db2.PingContext(ctx2)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db2.Close()

	// Obtener opciones de la base de datos de MYSQL

	paises, err := Controller.GetOpcionesPregunta(db2, "SELECT distinct id, pais_es FROM CatPaises where deleted is null order by pais_es")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Paises: paises,
		Sexo:   sexo,
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Voluntarios/Corporativo/corporativo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func formularioCorporativoEnglish(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	// connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	// db, err = sql.Open("mysql", connString)
	// if err != nil {
	// 	log.Fatal("Error creating connection pool: 1", err.Error())
	// }
	// ctx := context.Background()
	// err = db.PingContext(ctx)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// var db2 *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)

	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()
	// Obtener opciones de la base de datos de SQL

	paises, err := Controller.GetOpcionesPregunta(db, "SELECT distinct id, pais_en FROM CatPaises where deleted is null order by pais_en")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Obtener opciones de la base de datos de MYSQL
	/* 	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	   	if err != nil {
	   			http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	   	}
	*/
	data := FormData{
		Paises: paises,
		// Sexo:   sexo,
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Voluntarios/Corporativo/corporativoenglish.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func voluntariadocr(w http.ResponseWriter, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()
	// Obtener opciones de la base de datos de  MySQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	camisa, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM volunteering_shirt_sizes where id not in(6,8)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paises, err := Controller.GetOpcionesPregunta(db, "SELECT distinct id, name FROM admin_main_gwdata.countries where id not in (6) order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=6")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, conceptoPerfil FROM type_beneficiarios where id in (5,6,9)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	relacioncom, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM type_beneficiarios where tipoVoluntario=3")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	copdonor, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_donantes where idp=6")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	programa, err := Controller.GetOpcionesPregunta(db, "SELECT distinct Ida, voluntariado FROM sede_areas where IdPais=7 and Ida not in (8)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Crear estructura de datos con opciones para cada pregunta select
	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Camisa:       camisa,
		Paises:       paises,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
		Relacioncom:  relacioncom,
		Copdonor:     copdonor,
		Programa:     programa,
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Voluntarios/cr.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func voluntariadomx(w http.ResponseWriter, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()
	// Obtener opciones de la base de datos de  MySQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	camisa, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM volunteering_shirt_sizes where id not in(6,8)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paises, err := Controller.GetOpcionesPregunta(db, "SELECT distinct id, name FROM admin_main_gwdata.countries where id not in (17) order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=17")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, conceptoPerfil FROM type_beneficiarios where id in (5,6,9)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	relacioncom, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM type_beneficiarios where tipoVoluntario=3")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	copdonor, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_donantes where idp=17")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	programa, err := Controller.GetOpcionesPregunta(db, "SELECT distinct Ida, voluntariado FROM sede_areas where IdPais=17 and Ida not in(8)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Crear estructura de datos con opciones para cada pregunta select
	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Camisa:       camisa,
		Paises:       paises,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
		Relacioncom:  relacioncom,
		Copdonor:     copdonor,
		Programa:     programa,
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Voluntarios/mx.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func voluntariadohn(w http.ResponseWriter, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()
	// Obtener opciones de la base de datos de  MySQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	camisa, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM volunteering_shirt_sizes where id not in(6,8)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paises, err := Controller.GetOpcionesPregunta(db, "SELECT distinct id, name FROM admin_main_gwdata.countries where id not in(3) order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=3")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, conceptoPerfil FROM type_beneficiarios where id in (5,6,9)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	relacioncom, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM type_beneficiarios where tipoVoluntario=3")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	institucion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM voluntary_institutions where country_id=3")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	copdonor, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_donantes where idp=3")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	programa, err := Controller.GetOpcionesPregunta(db, "SELECT distinct Ida, voluntariado FROM sede_areas where IdPais=7 and Ida not in (8)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Crear estructura de datos con opciones para cada pregunta select
	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Camisa:       camisa,
		Paises:       paises,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
		Relacioncom:  relacioncom,
		Institucion:  institucion,
		Copdonor:     copdonor,
		Programa:     programa,
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Voluntarios/hn.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func postCorporativo(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	pais := r.FormValue("pais")
	nombreCompleto := r.FormValue("nombreCompleto")
	FechaN := r.FormValue("FechaN")
	sexo := r.FormValue("sexo")
	DUIM := r.FormValue("DUIM")
	empresa := r.FormValue("empresa")
	FechaAct := r.FormValue("FechaAct")
	autR1 := r.FormValue("autR1")
	aut3 := r.FormValue("aut3")

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		tx.Rollback()
		return
	}

	stmt, err := tx.Prepare("INSERT INTO FormVoluntarioCorporativo(pais, nombreCompleto, fechaNac, sexo, identidad, empresa, fechaAct, AutCeroT, AutMedios) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	var lastInsertId int
	err = stmt.QueryRow(pais, nombreCompleto, FechaN, sexo, DUIM, empresa, FechaAct, autR1, aut3).Scan(&lastInsertId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type Datos struct {
		Nombre string
		Id     int
		Fecha  string
	}

	t2, err := time.Parse("2006-01-02", FechaAct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	// Formatear la fecha al formato deseado
	FechaAct = t2.Format("02 de Enero de 2006")

	datos := Datos{
		Nombre: nombreCompleto,
		Id:     lastInsertId,
		Fecha:  FechaAct,
	}

	t, err := template.ParseFiles("public/HTML/RespuestaPersonalizadaCorp.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar la plantilla con los datos
	err = t.Execute(w, datos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

/* func emailequip(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	// Leer el PDF desde la solicitud y escribirlo en el archivo adjunto
	pdfFile, _, err := r.FormFile("pdf")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
	email.SetHeader("Subject", "Comprobante de Evaluación EQUIP")
	email.SetBody("text/html", "Buen día, <br><br> A continuación se adjunta el comprobante de la evaluación EQUIP que acabas de completar. Cualquier duda, puedes contactar a la coordinación nacional de M&E.  <br><br> ¡Mil gracias!")
	// Adjuntar el archivo PDF
	email.Attach("archivo.pdf", gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write(pdfData)
		return err
	}))

	// Crear el cliente SMTP y enviar el correo
	dialer := gomail.NewDialer(smtpServer, smtpPort, smtpUsername, smtpPassword)
	if err := dialer.DialAndSend(email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

func emailequip2(w http.ResponseWriter, r *http.Request) {
	pdfFile, _, err := r.FormFile("pdf")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	token, err := getAccessToken()
	if err != nil {
		http.Error(w, "Error al obtener token de acceso: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Codificar el PDF como Base64 para adjuntarlo
	encodedPDF := base64.StdEncoding.EncodeToString(pdfData)

	emailPayload := map[string]interface{}{
		"message": map[string]interface{}{
			"subject": "Comprobante de Evaluación EQUIP",
			"body": map[string]string{
				"contentType": "HTML",
				"content":     "Buen día,<br><br> A continuación se adjunta el comprobante de la evaluación EQUIP que acabas de completar. Cualquier duda, puedes contactar a la coordinación nacional de M&E.<br><br>¡Mil gracias!",
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

func postvolgt(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	fechaNac := r.FormValue("FechaN")
	nacionalidad := cambio(r.FormValue("gi"))
	nombreRepresentante := r.FormValue("nombreRep")
	numIdentResponsable := r.FormValue("DUIP")
	parentescoId := cambio(r.FormValue("parentesco"))
	nombreVoluntarioMenor := r.FormValue("nombreAdolescente")
	edadVoluntario := cambio(r.FormValue("edad"))
	tipoDocIndentRespExtId := cambio(r.FormValue("elec"))
	numIdentVoluntario := r.FormValue("DUI")
	nombre := r.FormValue("nombre")
	apellido := r.FormValue("apellido")
	sexoId := cambio(r.FormValue("sexo"))
	numContacto := r.FormValue("TM")
	correo := r.FormValue("correo")
	tallaCamisaId := cambio(r.FormValue("camisa"))
	tiempoVoluntariadoId := cambio(r.FormValue("voluntario"))
	paisNac := cambio(r.FormValue("paisn"))
	deptoId := cambio(r.FormValue("pais"))
	municipioId := cambio(r.FormValue("departamento"))
	direccion := r.FormValue("direccion")
	estudiaId := cambio(r.FormValue("TdC"))
	ultimoEstudioId := cambio(r.FormValue("UT"))
	gradoEstudioId := cambio(r.FormValue("GA"))
	universidad := r.FormValue("U")
	carreraUniv := r.FormValue("carrera")
	horasSocId := cambio(r.FormValue("HS"))
	refPersonal := r.FormValue("nombrer")
	numRefPersonal := r.FormValue("telr")
	refPersonalb := r.FormValue("nombrer2")
	numRefPersonalb := r.FormValue("telr2")
	perfilVoluntariadoId := cambio(r.FormValue("acc"))
	voluntarioComunitarioId := cambio(r.FormValue("acc2"))
	institucionId := cambio(r.FormValue("accinst"))
	donante := cambio(r.FormValue("acc3"))
	areaDepartamento := r.FormValue("puesto")
	fechaInduccion := r.FormValue("FechaI")
	nombreDesarrollador := r.FormValue("nombrei")
	areaId := cambio(r.FormValue("pais2"))
	actividadId := cambio(r.FormValue("actividad"))
	subActividadId := cambio(r.FormValue("subact"))
	horarioActividad := r.FormValue("OHT")
	deptoActividadId := cambio(r.FormValue("departamento2"))
	sedeActividadId := cambio(r.FormValue("sede2"))
	paisId := 4
	otroClub := r.FormValue("aT")
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)
	fechaVencimientoR := r.FormValue("fechaV")

	if len(fechaVencimientoR) == 0 {
		fechaVencimientoR = "2000-01-01"
	}

	urldocF := ""
	urldocA := ""
	urldocR := ""
	urlpass := ""

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_Voluntariado (fechaNac,nacionalidad,nombreRepresentante,numIdentResponsable,parentescoId,nombreVoluntarioMenor,edadVoluntario,tipoDocIndentRespExtId,numIdentVoluntario,nombre,apellido,sexoId,numContacto,correo,tallaCamisaId,tiempoVoluntariadoId,paisNac,deptoId,municipioId,direccion,estudiaId,ultimoEstudioId,gradoEstudioId,universidad,carreraUniv,horasSocId,refPersonal,numRefPersonal,refPersonalb,numRefPersonalb,perfilVoluntariadoId,voluntarioComunitarioId,institucionId,donante,areaDepartamento,fechaInduccion,nombreDesarrollador,areaId,actividadId,subActividadId,horarioActividad,deptoActividadId,sedeActividadId,otroClub,fechaRegistro,fechaVencimientoR,urlDuiFrontal,urlDuiDorso,urlRenas,paisId,docExtranjero) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(fechaNac, nacionalidad, nombreRepresentante, numIdentResponsable, parentescoId, nombreVoluntarioMenor, edadVoluntario, tipoDocIndentRespExtId, numIdentVoluntario, nombre, apellido, sexoId, numContacto, correo, tallaCamisaId, tiempoVoluntariadoId, paisNac, deptoId, municipioId, direccion, estudiaId, ultimoEstudioId, gradoEstudioId, universidad, carreraUniv, horasSocId, refPersonal, numRefPersonal, refPersonalb, numRefPersonalb, perfilVoluntariadoId, voluntarioComunitarioId, institucionId, donante, areaDepartamento, fechaInduccion, nombreDesarrollador, areaId, actividadId, subActividadId, horarioActividad, deptoActividadId, sedeActividadId, otroClub, currentTimeGMT, fechaVencimientoR, urldocF, urldocA, urldocR, paisId, urlpass)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO f_dias (voluntarioId, diasId) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dias", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta dias", http.StatusInternalServerError)
				return
			}
		}
	}

	discapacidad4 := r.Form["dis4"]
	if len(discapacidad4) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO f_horarios (voluntarioId, horarioId) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta horario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad4 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta horario", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
	handleError(err)
	// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
	client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
	handleError(err)

	file, fileHandler, err := r.FormFile("docF")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename
		fileExt := filepath.Ext(fileName)
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "gt/duif/"+strconv.Itoa(int(idGenerado))+fileExt, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urldocF = strconv.Itoa(int(idGenerado)) + fileExt
	}
	file2, handler2, err := r.FormFile("docA")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename
		fileExt2 := filepath.Ext(fileName2)

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "gt/duia/"+strconv.Itoa(int(idGenerado))+fileExt2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		urldocA = strconv.Itoa(int(idGenerado)) + fileExt2
	}

	file3, fileHandler3, err := r.FormFile("pas")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file3.Close()

		// Obtiene el nombre del archivo cargado
		fileName3 := fileHandler3.Filename
		fileExt3 := filepath.Ext(fileName3)

		fileBytes3, err := io.ReadAll(file3)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile3, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile3.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile3.Write(fileBytes3); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "gt/pasaporte/"+strconv.Itoa(int(idGenerado))+fileExt3, tempFile3,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urlpass = strconv.Itoa(int(idGenerado)) + fileExt3
	}

	file4, fileHandler4, err := r.FormFile("docR")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file4.Close()

		// Obtiene el nombre del archivo cargado
		fileName4 := fileHandler4.Filename
		fileExt4 := filepath.Ext(fileName4)
		fileBytes4, err := io.ReadAll(file4)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile4, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile4.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile4.Write(fileBytes4); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "gt/renas/"+strconv.Itoa(int(idGenerado))+fileExt4, tempFile4,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urldocR = strconv.Itoa(int(idGenerado)) + fileExt4
	}

	tx2, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt4, err := tx2.Prepare("UPDATE F_Voluntariado SET urlDuiFrontal = ?, urlDuiDorso = ? ,docExtranjero= ?, urlRenas=? WHERE id = ?")
	if err != nil {
		tx2.Rollback()
		http.Error(w, "Error al preparar update", http.StatusInternalServerError)
		return
	}
	defer stmt4.Close()

	_, err = stmt4.Exec(urldocF, urldocA, urlpass, urldocR, idGenerado)
	if err != nil {
		tx2.Rollback()
		http.Error(w, "Error al ejecutar ejecutar update", http.StatusInternalServerError)
		return
	}

	r.ParseForm()

	err = tx2.Commit()
	if err != nil {
		tx2.Rollback()
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

func postmdv(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	// Obtener los valores de los campos del formulario
	idF := 1
	pais := 7
	idSede := cambio(r.FormValue("sede"))
	fechaNac := r.FormValue("FechaN")
	nombres := r.FormValue("nombre")
	apellidos := r.FormValue("apellido")
	sexo := cambio(r.FormValue("Sexo"))
	nacionalidad := cambio(r.FormValue("nacionalidad"))
	identidad := r.FormValue("DUIM")
	discapacidad := cambio(r.FormValue("dis"))
	estudia := cambio(r.FormValue("Eactual"))
	estudioAlcanzado := cambio(r.FormValue("UT"))
	grado := cambio(r.FormValue("GA"))
	turno := cambio(r.FormValue("Turno"))
	seccion := cambio(r.FormValue("seccion"))
	telefono := r.FormValue("TM")
	nuevogw := cambio(r.FormValue("anterior"))
	autovozimg := cambio(r.FormValue("aut2"))
	tipoParticipante := 1

	estadocivil := cambio(r.FormValue("EstadoC"))
	idExtranjero := cambio(r.FormValue("extranjero"))
	concatows := cambio(r.FormValue("WS"))
	whatsapp := r.FormValue("TM2")
	medioTransporte := cambio(r.FormValue("Transp"))
	gastoDiario := r.FormValue("gasto")
	contactoEmg := r.FormValue("contacto")
	telEmg := r.FormValue("Telcontacto")
	parentescoEmg := cambio(r.FormValue("relacion"))
	diplomado := cambio(r.FormValue("diplomado"))
	procesoNivelacion := cambio(r.FormValue("nivelacion"))
	direccion := r.FormValue("direccion")
	comptel := r.FormValue("comptel")

	if gastoDiario == "" {
		gastoDiario = "0"
	}
	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	stmt, err := db.Prepare("IF EXISTS(SELECT * FROM FormularioRegisto WHERE nombres = @P1 and apellidos = @P2 and identidad = @P3 and pais = @P4 and fechaNAc = @P5 and idF = @P6 and deleted is null)Begin select 1 end else Begin select 0 end")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	row := stmt.QueryRow(nombres, apellidos, identidad, pais, fechaNac, idF)

	var result int
	_ = row.Scan(&result)

	type Datos struct {
		Nombre   string
		Apellido string
		Id       int
		Programa string
	}

	if result == 1 {

		datos := Datos{
			Nombre:   nombres,
			Apellido: apellidos,
		}

		t, err := template.ParseFiles("Public/Html/RespuestaDuplicada.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Renderizar la plantilla con los datos
		err = t.Execute(w, datos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		tx, err := db.Begin()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return
		}

		stmt, err := tx.Prepare("INSERT INTO FormularioRegisto(idF,tipoParticipante,pais,idSede,fechaNac,nombres,apellidos,sexo,nacionalidad,identidad,discapacidad,estudia,estudioAlcanzado,grado,turno,seccion,telefono,nuevogw,autovozimg) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14,@P15,@P16,@P17,@P18,@P19)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return
		}

		var lastInsertId int
		err = stmt.QueryRow(idF, tipoParticipante, pais, idSede, fechaNac, nombres, apellidos, sexo, nacionalidad, identidad, discapacidad, estudia, estudioAlcanzado, grado, turno, seccion, telefono, nuevogw, autovozimg).Scan(&lastInsertId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return
		}

		stmt, err = tx.Prepare("INSERT INTO RegistroJuventud(idFR,estadocivil ,idExtranjero ,concatows ,whatsapp ,medioTransporte ,gastoDiario ,contactoEmg ,telEmg ,parentescoEmg ,diplomado ,procesoNivelacion ,direccion,compania) VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return
		}

		_, err = stmt.Exec(lastInsertId, estadocivil, idExtranjero, concatows, whatsapp, medioTransporte, gastoDiario, contactoEmg, telEmg, parentescoEmg, diplomado, procesoNivelacion, direccion, comptel)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return
		}

		discapacidadmultiple := r.Form["dis2"]
		if len(discapacidadmultiple) == 0 {
		} else {
			stmt, err := tx.Prepare("INSERT INTO MultipleDiscapacidad (idFR,iddiscapacidad) VALUES (@P1,@P2)")
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al preparar la consulta ocupacion", http.StatusInternalServerError)
				return
			}
			defer stmt.Close()

			for _, valor := range discapacidadmultiple {
				if valor != "" {
					_, err = stmt.Exec(lastInsertId, valor)
					if err != nil {

						tx.Rollback()
						http.Error(w, "Error al ejecutar la consulta situaciones", http.StatusInternalServerError)
						return
					}
				}
			}
		}

		vives := r.Form["vives"]
		if len(vives) == 0 {
		} else {
			stmt, err := tx.Prepare("INSERT INTO MultipleJuventud (idpregunta,idFR,idRespuesta) VALUES (1,@P1,@P2)")
			if err != nil {

				tx.Rollback()
				http.Error(w, "Error al preparar la consulta vives", http.StatusInternalServerError)
				return
			}
			defer stmt.Close()

			for _, valor := range vives {
				if valor != "" {
					_, err = stmt.Exec(lastInsertId, valor)
					if err != nil {

						tx.Rollback()
						http.Error(w, "Error al ejecutar la consulta vives", http.StatusInternalServerError)
						return
					}
				}
			}
		}

		dispositivo := r.Form["dispositivo"]
		if len(dispositivo) == 0 {
		} else {
			stmt, err := tx.Prepare("INSERT INTO MultipleJuventud (idpregunta,idFR,idRespuesta) VALUES (2,@P1,@P2)")
			if err != nil {

				tx.Rollback()
				http.Error(w, "Error al preparar la consulta dispositivo", http.StatusInternalServerError)
				return
			}
			defer stmt.Close()

			for _, valor := range dispositivo {
				if valor != "" {
					_, err = stmt.Exec(lastInsertId, valor)
					if err != nil {

						tx.Rollback()
						http.Error(w, "Error al ejecutar la consulta dispositivo", http.StatusInternalServerError)
						return
					}
				}
			}
		}

		err = tx.Commit()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var programa string
		switch idSede {
		case 666:
			programa = "Creando Profesionales - San Salvador"
		case 1128:
			programa = "Jóvenes Constructores - Santa Ana"

		case 2968:
			programa = "Creando Profesionales - San Miguel"
		}
		datos := Datos{
			Nombre:   nombres,
			Apellido: apellidos,
			Id:       lastInsertId,
			Programa: programa,
		}

		t, err := template.ParseFiles("Public/Html/RespuestaPersonalizada.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Renderizar la plantilla con los datos
		err = t.Execute(w, datos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func cambio(valor string) int {
	if valor == "" {
		return 0 // Valor predeterminado si la cadena está vacía
	}
	resultado, _ := strconv.Atoi(valor) // Convertir la cadena a int
	return resultado
}

func Edusv(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}

	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.sede_areas where Ida=5 and IdPais=7 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where typePerson=1 and active_at is not null and id not in(10)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Educacion/edusv.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func obtenermunsedeHandleredu(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	depsede := r.FormValue("depsede")
	var opcionesmunsede string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT Distinct Idmun,Municipio FROM admin_main_gwdata.sede_areas where Iddep=? and Ida=5 order by Municipio", depsede)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionesmunsede += "<option value=''></option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesmunsede += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}

	fmt.Fprint(w, opcionesmunsede)
}

func obtenersedeHandleredu(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	sede := r.FormValue("munsede")
	var opcionessede string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT IdSede,Sede FROM admin_main_gwdata.vwSedesFormMentoriasESA where Idmun=? order by Sede", sede)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionessede += "<option value=''></option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionessede += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}

	fmt.Fprint(w, opcionessede)
}

func postedusv(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	// Obtener los valores de los campos del formulario
	Nombrerepresentante := r.FormValue("nombreRep")
	Duirepresentante := r.FormValue("DUIP")
	parentesco := r.FormValue("parentesco")
	parentescop := cambio(parentesco)
	edad := r.FormValue("edad")
	edadp := cambio(edad)
	nombreAdolescente := r.FormValue("nombreAdolescente")
	telefonoresp := r.FormValue("numResp")
	nombre := r.FormValue("nombreCompleto")
	apellido := r.FormValue("apellido")
	sexo := r.FormValue("Sexo")
	sexop := cambio(sexo)
	pais := 7
	departamento := r.FormValue("pais")
	departamentop := cambio(departamento)
	municipio := r.FormValue("departamento")
	municipiop := cambio(municipio)
	sede := r.FormValue("sede2")
	sedep := cambio(sede)
	tipo := cambio(r.FormValue("tipo"))
	Estudia := r.FormValue("TdC")
	Estudiap := cambio(Estudia)
	ultimoestudio := r.FormValue("UT")
	ultimoestudiop := cambio(ultimoestudio)
	grado := r.FormValue("GA")
	gradop := cambio(grado)
	seccion := r.FormValue("seccion")
	seccionp := cambio(seccion)
	turno := r.FormValue("Turno")
	turnop := cambio(turno)
	discapacidad := r.FormValue("dis")
	dispacacidadp := cambio(discapacidad)
	nuevogw := r.FormValue("gp")
	nuevogwp := cambio(nuevogw)
	voz_img := r.FormValue("aut2")
	voz_imgp := cambio(voz_img)
	DUI := r.FormValue("DUIM")
	urldocF := ""
	urldocA := ""
	Formulario := 2
	FechaNac := r.FormValue("FechaN")
	nacionalidad := r.FormValue("nacionalidad")
	nacionalidadp := cambio(nacionalidad)

	if r.FormValue("aut") == "10" {
		cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
		handleError(err)
		// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
		client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
		handleError(err)
		file, fileHandler, err := r.FormFile("docF")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "educacion/sv/duif/"+fileName, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		file2, handler2, err := r.FormFile("docA")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "educacion/sv/duia/"+fileName2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)

		urldocF = "https://gwforms.blob.core.windows.net/participantes/educacion/sv/duif/" + fileName
		urldocA = "https://gwforms.blob.core.windows.net/participantes/educacion/sv/duia/" + fileName2
	}

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_participantes (Nombrerepresentante,Duirepresentante,parentesco,edad,nombreAdolescente,telefonoresp,nombre,apellido,sexo,pais,departamento,municipio,sede,tipo,Estudia,ultimoestudio,grado,seccion,turno,discapacidad,nuevogw,voz_img,DUI,urldui,urlduia, Formulario,FechaNac,nacionalidad,fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,? )")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(Nombrerepresentante, Duirepresentante, parentescop, edadp, nombreAdolescente, telefonoresp, nombre, apellido, sexop, pais, departamentop, municipiop, sedep, tipo, Estudiap, ultimoestudiop, gradop, seccionp, turnop, dispacacidadp, nuevogwp, voz_imgp, DUI, urldocF, urldocA, Formulario, FechaNac, nacionalidadp, currentTimeGMT)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_discapacidad (idp, discapacidad) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dis", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func Eduhn(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}

	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=3")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.sede_areas where Ida=5 and IdPais=3 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where typePerson=1 and active_at is not null and id not in(10)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Educacion/eduhn.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func posteduhn(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	// Obtener los valores de los campos del formulario
	Nombrerepresentante := r.FormValue("nombreRep")
	Duirepresentante := r.FormValue("DUIP")
	parentesco := r.FormValue("parentesco")
	parentescop := cambio(parentesco)
	edad := r.FormValue("edad")
	edadp := cambio(edad)
	nombreAdolescente := r.FormValue("nombreAdolescente")
	telefonoresp := r.FormValue("numResp")
	nombre := r.FormValue("nombreCompleto")
	apellido := r.FormValue("apellido")
	sexo := r.FormValue("Sexo")
	sexop := cambio(sexo)
	pais := 3
	departamento := r.FormValue("pais")
	departamentop := cambio(departamento)
	municipio := r.FormValue("departamento")
	municipiop := cambio(municipio)
	sede := r.FormValue("sede2")
	sedep := cambio(sede)
	tipo := cambio(r.FormValue("tipo"))
	Estudia := r.FormValue("TdC")
	Estudiap := cambio(Estudia)
	ultimoestudio := r.FormValue("UT")
	ultimoestudiop := cambio(ultimoestudio)
	grado := r.FormValue("GA")
	gradop := cambio(grado)
	seccion := r.FormValue("seccion")
	seccionp := cambio(seccion)
	turno := r.FormValue("Turno")
	turnop := cambio(turno)
	discapacidad := r.FormValue("dis")
	dispacacidadp := cambio(discapacidad)
	nuevogw := r.FormValue("gp")
	nuevogwp := cambio(nuevogw)
	voz_img := r.FormValue("aut2")
	voz_imgp := cambio(voz_img)
	DUI := r.FormValue("DUIM")
	urldocF := ""
	urldocA := ""
	Formulario := 2
	FechaNac := r.FormValue("FechaN")
	nacionalidad := r.FormValue("nacionalidad")
	nacionalidadp := cambio(nacionalidad)

	if r.FormValue("aut") == "10" {
		cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
		handleError(err)
		// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
		client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
		handleError(err)
		file, fileHandler, err := r.FormFile("docF")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "educacion/hn/duif/"+fileName, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		file2, handler2, err := r.FormFile("docA")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "educacion/hn/duia/"+fileName2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)

		urldocF = "https://gwforms.blob.core.windows.net/participantes/educacion/hn/duif/" + fileName
		urldocA = "https://gwforms.blob.core.windows.net/participantes/educacion/hn/duia/" + fileName2
	}

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_participantes (Nombrerepresentante,Duirepresentante,parentesco,edad,nombreAdolescente,telefonoresp,nombre,apellido,sexo,pais,departamento,municipio,sede,tipo,Estudia,ultimoestudio,grado,seccion,turno,discapacidad,nuevogw,voz_img,DUI,urldui,urlduia, Formulario,FechaNac,nacionalidad,fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,? )")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(Nombrerepresentante, Duirepresentante, parentescop, edadp, nombreAdolescente, telefonoresp, nombre, apellido, sexop, pais, departamentop, municipiop, sedep, tipo, Estudiap, ultimoestudiop, gradop, seccionp, turnop, dispacacidadp, nuevogwp, voz_imgp, DUI, urldocF, urldocA, Formulario, FechaNac, nacionalidadp, currentTimeGMT)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_discapacidad (idp, discapacidad) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dis", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func Edugt(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}

	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=4")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.sede_areas where Ida=5 and IdPais=4 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where typePerson=1 and active_at is not null and id not in(10)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Educacion/edugt.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postedugt(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	Nombrerepresentante := r.FormValue("nombreRep")
	Duirepresentante := r.FormValue("DUIP")
	parentesco := r.FormValue("parentesco")
	parentescop := cambio(parentesco)
	edad := r.FormValue("edad")
	edadp := cambio(edad)
	nombreAdolescente := r.FormValue("nombreAdolescente")
	telefonoresp := r.FormValue("numResp")
	nombre := r.FormValue("nombreCompleto")
	apellido := r.FormValue("apellido")
	sexo := r.FormValue("Sexo")
	sexop := cambio(sexo)
	pais := 4
	departamento := r.FormValue("pais")
	departamentop := cambio(departamento)
	municipio := r.FormValue("departamento")
	municipiop := cambio(municipio)
	sede := r.FormValue("sede2")
	sedep := cambio(sede)
	tipo := cambio(r.FormValue("tipo"))
	Estudia := r.FormValue("TdC")
	Estudiap := cambio(Estudia)
	ultimoestudio := r.FormValue("UT")
	ultimoestudiop := cambio(ultimoestudio)
	grado := r.FormValue("GA")
	gradop := cambio(grado)
	seccion := r.FormValue("seccion")
	seccionp := cambio(seccion)
	turno := r.FormValue("Turno")
	turnop := cambio(turno)
	discapacidad := r.FormValue("dis")
	dispacacidadp := cambio(discapacidad)
	nuevogw := r.FormValue("gp")
	nuevogwp := cambio(nuevogw)
	voz_img := r.FormValue("aut2")
	voz_imgp := cambio(voz_img)
	DUI := r.FormValue("DUIM")
	urldocF := ""
	urldocA := ""
	Formulario := 2
	FechaNac := r.FormValue("FechaN")
	nacionalidad := r.FormValue("nacionalidad")
	nacionalidadp := cambio(nacionalidad)
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	if r.FormValue("aut") == "10" {
		cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
		handleError(err)
		// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
		client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
		handleError(err)
		file, fileHandler, err := r.FormFile("docF")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "educacion/gt/duif/"+fileName, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		file2, handler2, err := r.FormFile("docA")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "educacion/gt/duia/"+fileName2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)

		urldocF = "https://gwforms.blob.core.windows.net/participantes/educacion/gt/duif/" + fileName
		urldocA = "https://gwforms.blob.core.windows.net/participantes/educacion/gt/duia/" + fileName2
	}

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_participantes (Nombrerepresentante,Duirepresentante,parentesco,edad,nombreAdolescente,telefonoresp,nombre,apellido,sexo,pais,departamento,municipio,sede,tipo,Estudia,ultimoestudio,grado,seccion,turno,discapacidad,nuevogw,voz_img,DUI,urldui,urlduia, Formulario,FechaNac,nacionalidad,fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,? )")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(Nombrerepresentante, Duirepresentante, parentescop, edadp, nombreAdolescente, telefonoresp, nombre, apellido, sexop, pais, departamentop, municipiop, sedep, tipo, Estudiap, ultimoestudiop, gradop, seccionp, turnop, dispacacidadp, nuevogwp, voz_imgp, DUI, urldocF, urldocA, Formulario, FechaNac, nacionalidadp, currentTimeGMT)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_discapacidad (idp, discapacidad) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dis", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func saludsv(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}

	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM VSedesSalud where IdPais=7 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where typePerson=1 and active_at is not null and id not in(10)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/saludsv.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postesaludsv(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	Nombrerepresentante := r.FormValue("nombreRep")
	Duirepresentante := r.FormValue("DUIP")
	parentesco := r.FormValue("parentesco")
	parentescop := cambio(parentesco)
	edad := r.FormValue("edad")
	edadp := cambio(edad)
	nombreAdolescente := r.FormValue("nombreAdolescente")
	telefonoresp := r.FormValue("numResp")
	nombre := r.FormValue("nombreCompleto")
	apellido := r.FormValue("apellido")
	sexo := r.FormValue("Sexo")
	sexop := cambio(sexo)
	pais := 7
	departamento := r.FormValue("pais")
	departamentop := cambio(departamento)
	municipio := r.FormValue("departamento")
	municipiop := cambio(municipio)
	sede := r.FormValue("sede2")
	sedep := cambio(sede)
	tipo := cambio(r.FormValue("tipo"))
	Estudia := r.FormValue("TdC")
	Estudiap := cambio(Estudia)
	ultimoestudio := r.FormValue("UT")
	ultimoestudiop := cambio(ultimoestudio)
	grado := r.FormValue("GA")
	gradop := cambio(grado)
	seccion := r.FormValue("seccion")
	seccionp := cambio(seccion)
	turno := r.FormValue("Turno")
	turnop := cambio(turno)
	discapacidad := r.FormValue("dis")
	dispacacidadp := cambio(discapacidad)
	nuevogw := r.FormValue("gp")
	nuevogwp := cambio(nuevogw)
	voz_img := r.FormValue("aut2")
	voz_imgp := cambio(voz_img)
	DUI := r.FormValue("DUIM")
	urldocF := ""
	urldocA := ""
	Formulario := 3
	FechaNac := r.FormValue("FechaN")
	nacionalidad := r.FormValue("nacionalidad")
	nacionalidadp := cambio(nacionalidad)
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	if r.FormValue("aut") == "10" {
		cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
		handleError(err)
		// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
		client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
		handleError(err)
		file, fileHandler, err := r.FormFile("docF")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "salud/sv/duif/"+fileName, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		file2, handler2, err := r.FormFile("docA")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "salud/sv/duia/"+fileName2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)

		urldocF = "https://gwforms.blob.core.windows.net/participantes/salud/sv/duif/" + fileName
		urldocA = "https://gwforms.blob.core.windows.net/participantes/salud/sv/duia/" + fileName2
	}

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_participantes (Nombrerepresentante,Duirepresentante,parentesco,edad,nombreAdolescente,telefonoresp,nombre,apellido,sexo,pais,departamento,municipio,sede,tipo,Estudia,ultimoestudio,grado,seccion,turno,discapacidad,nuevogw,voz_img,DUI,urldui,urlduia, Formulario,FechaNac,nacionalidad,fechaRegistro) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(Nombrerepresentante, Duirepresentante, parentescop, edadp, nombreAdolescente, telefonoresp, nombre, apellido, sexop, pais, departamentop, municipiop, sedep, tipo, Estudiap, ultimoestudiop, gradop, seccionp, turnop, dispacacidadp, nuevogwp, voz_imgp, DUI, urldocF, urldocA, Formulario, FechaNac, nacionalidadp, currentTimeGMT)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_discapacidad (idp, discapacidad) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dis", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func obtenermunsedeHandlersalud(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	depsede := r.FormValue("depsede")
	var opcionesmunsede string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT Distinct Idmun,Municipio FROM VSedesSalud where Iddep=? order by Municipio", depsede)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionesmunsede += "<option value=''></option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesmunsede += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionesmunsede)
}

func obtenersedeHandlersalud(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	sede := r.FormValue("munsede")
	var opcionessede string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT IdSede,Sede FROM VSedesSalud where Idmun=? order by Sede", sede)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionessede += "<option value=''></option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionessede += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}

	fmt.Fprint(w, opcionessede)
}

func Edupn(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.sede_areas where Ida=5 and IdPais=1 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where typePerson=1 and active_at is not null and id not in(10)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Educacion/edupn.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postedupn(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	Nombrerepresentante := r.FormValue("nombreRep")
	Duirepresentante := r.FormValue("DUIP")
	parentesco := r.FormValue("parentesco")
	parentescop := cambio(parentesco)
	edad := r.FormValue("edad")
	edadp := cambio(edad)
	nombreAdolescente := r.FormValue("nombreAdolescente")
	telefonoresp := r.FormValue("numResp")
	nombre := r.FormValue("nombreCompleto")
	apellido := r.FormValue("apellido")
	sexo := r.FormValue("Sexo")
	sexop := cambio(sexo)
	pais := 1
	departamento := r.FormValue("pais")
	departamentop := cambio(departamento)
	municipio := r.FormValue("departamento")
	municipiop := cambio(municipio)
	sede := r.FormValue("sede2")
	sedep := cambio(sede)
	tipo := cambio(r.FormValue("tipo"))
	Estudia := r.FormValue("TdC")
	Estudiap := cambio(Estudia)
	ultimoestudio := r.FormValue("UT")
	ultimoestudiop := cambio(ultimoestudio)
	grado := r.FormValue("GA")
	gradop := cambio(grado)
	seccion := r.FormValue("seccion")
	seccionp := cambio(seccion)
	turno := r.FormValue("Turno")
	turnop := cambio(turno)
	discapacidad := r.FormValue("dis")
	dispacacidadp := cambio(discapacidad)
	nuevogw := r.FormValue("gp")
	nuevogwp := cambio(nuevogw)
	voz_img := r.FormValue("aut2")
	voz_imgp := cambio(voz_img)
	DUI := r.FormValue("DUIM")
	urldocF := ""
	urldocA := ""
	Formulario := 2
	FechaNac := r.FormValue("FechaN")
	nacionalidad := r.FormValue("nacionalidad")
	nacionalidadp := cambio(nacionalidad)

	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	if r.FormValue("aut") == "10" {
		cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
		handleError(err)
		// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
		client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
		handleError(err)
		file, fileHandler, err := r.FormFile("docF")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "educacion/pn/duif/"+fileName, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		file2, handler2, err := r.FormFile("docA")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "educacion/pn/duia/"+fileName2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)

		urldocF = "https://gwforms.blob.core.windows.net/participantes/educacion/pn/duif/" + fileName
		urldocA = "https://gwforms.blob.core.windows.net/participantes/educacion/pn/duia/" + fileName2
	}

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_participantes (Nombrerepresentante,Duirepresentante,parentesco,edad,nombreAdolescente,telefonoresp,nombre,apellido,sexo,pais,departamento,municipio,sede,tipo,Estudia,ultimoestudio,grado,seccion,turno,discapacidad,nuevogw,voz_img,DUI,urldui,urlduia, Formulario,FechaNac,nacionalidad,fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,? )")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(Nombrerepresentante, Duirepresentante, parentescop, edadp, nombreAdolescente, telefonoresp, nombre, apellido, sexop, pais, departamentop, municipiop, sedep, tipo, Estudiap, ultimoestudiop, gradop, seccionp, turnop, dispacacidadp, nuevogwp, voz_imgp, DUI, urldocF, urldocA, Formulario, FechaNac, nacionalidadp, currentTimeGMT)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_discapacidad (idp, discapacidad) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dis", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func Edurd(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=5")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.sede_areas where Ida=5 and IdPais=5 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where typePerson=1 and active_at is not null and id not in(10)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Educacion/edurd.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postedurd(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	Nombrerepresentante := r.FormValue("nombreRep")
	Duirepresentante := r.FormValue("DUIP")
	parentesco := r.FormValue("parentesco")
	parentescop := cambio(parentesco)
	edad := r.FormValue("edad")
	edadp := cambio(edad)
	nombreAdolescente := r.FormValue("nombreAdolescente")
	telefonoresp := r.FormValue("numResp")
	nombre := r.FormValue("nombreCompleto")
	apellido := r.FormValue("apellido")
	sexo := r.FormValue("Sexo")
	sexop := cambio(sexo)
	pais := 5
	departamento := r.FormValue("pais")
	departamentop := cambio(departamento)
	municipio := r.FormValue("departamento")
	municipiop := cambio(municipio)
	sede := r.FormValue("sede2")
	sedep := cambio(sede)
	tipo := cambio(r.FormValue("tipo"))
	Estudia := r.FormValue("TdC")
	Estudiap := cambio(Estudia)
	ultimoestudio := r.FormValue("UT")
	ultimoestudiop := cambio(ultimoestudio)
	grado := r.FormValue("GA")
	gradop := cambio(grado)
	seccion := r.FormValue("seccion")
	seccionp := cambio(seccion)
	turno := r.FormValue("Turno")
	turnop := cambio(turno)
	discapacidad := r.FormValue("dis")
	dispacacidadp := cambio(discapacidad)
	nuevogw := r.FormValue("gp")
	nuevogwp := cambio(nuevogw)
	voz_img := r.FormValue("aut2")
	voz_imgp := cambio(voz_img)
	DUI := r.FormValue("DUIM")
	urldocF := ""
	urldocA := ""
	Formulario := 2
	FechaNac := r.FormValue("FechaN")
	nacionalidad := r.FormValue("nacionalidad")
	nacionalidadp := cambio(nacionalidad)

	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	if r.FormValue("aut") == "10" {
		cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
		handleError(err)
		// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
		client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
		handleError(err)
		file, fileHandler, err := r.FormFile("docF")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "educacion/rd/duif/"+fileName, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		file2, handler2, err := r.FormFile("docA")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "educacion/rd/duia/"+fileName2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)

		urldocF = "https://gwforms.blob.core.windows.net/participantes/educacion/rd/duif/" + fileName
		urldocA = "https://gwforms.blob.core.windows.net/participantes/educacion/rd/duia/" + fileName2
	}

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_participantes (Nombrerepresentante,Duirepresentante,parentesco,edad,nombreAdolescente,telefonoresp,nombre,apellido,sexo,pais,departamento,municipio,sede,tipo,Estudia,ultimoestudio,grado,seccion,turno,discapacidad,nuevogw,voz_img,DUI,urldui,urlduia, Formulario,FechaNac,nacionalidad,fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,? )")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(Nombrerepresentante, Duirepresentante, parentescop, edadp, nombreAdolescente, telefonoresp, nombre, apellido, sexop, pais, departamentop, municipiop, sedep, tipo, Estudiap, ultimoestudiop, gradop, seccionp, turnop, dispacacidadp, nuevogwp, voz_imgp, DUI, urldocF, urldocA, Formulario, FechaNac, nacionalidadp, currentTimeGMT)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_discapacidad (idp, discapacidad) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dis", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func Edumx(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=17")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.sede_areas where Ida=5 and IdPais=17 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where typePerson=1 and active_at is not null and id not in(10)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Educacion/edumx.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postedumx(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	Nombrerepresentante := r.FormValue("nombreRep")
	Duirepresentante := r.FormValue("DUIP")
	parentesco := r.FormValue("parentesco")
	parentescop := cambio(parentesco)
	edad := r.FormValue("edad")
	edadp := cambio(edad)
	nombreAdolescente := r.FormValue("nombreAdolescente")
	telefonoresp := r.FormValue("numResp")
	nombre := r.FormValue("nombreCompleto")
	apellido := r.FormValue("apellido")
	sexo := r.FormValue("Sexo")
	sexop := cambio(sexo)
	pais := 17
	departamento := r.FormValue("pais")
	departamentop := cambio(departamento)
	municipio := r.FormValue("departamento")
	municipiop := cambio(municipio)
	sede := r.FormValue("sede2")
	sedep := cambio(sede)
	tipo := cambio(r.FormValue("tipo"))
	Estudia := r.FormValue("TdC")
	Estudiap := cambio(Estudia)
	ultimoestudio := r.FormValue("UT")
	ultimoestudiop := cambio(ultimoestudio)
	grado := r.FormValue("GA")
	gradop := cambio(grado)
	seccion := r.FormValue("seccion")
	seccionp := cambio(seccion)
	turno := r.FormValue("Turno")
	turnop := cambio(turno)
	discapacidad := r.FormValue("dis")
	dispacacidadp := cambio(discapacidad)
	nuevogw := r.FormValue("gp")
	nuevogwp := cambio(nuevogw)
	voz_img := r.FormValue("aut2")
	voz_imgp := cambio(voz_img)
	DUI := r.FormValue("DUIM")
	urldocF := ""
	urldocA := ""
	Formulario := 2
	FechaNac := r.FormValue("FechaN")
	nacionalidad := r.FormValue("nacionalidad")
	nacionalidadp := cambio(nacionalidad)

	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	if r.FormValue("aut") == "10" {
		cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
		handleError(err)
		// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
		client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
		handleError(err)
		file, fileHandler, err := r.FormFile("docF")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "educacion/mx/duif/"+fileName, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		file2, handler2, err := r.FormFile("docA")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "educacion/mx/duia/"+fileName2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)

		urldocF = "https://gwforms.blob.core.windows.net/participantes/educacion/mx/duif/" + fileName
		urldocA = "https://gwforms.blob.core.windows.net/participantes/educacion/mx/duia/" + fileName2
	}

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_participantes (Nombrerepresentante,Duirepresentante,parentesco,edad,nombreAdolescente,telefonoresp,nombre,apellido,sexo,pais,departamento,municipio,sede,tipo,Estudia,ultimoestudio,grado,seccion,turno,discapacidad,nuevogw,voz_img,DUI,urldui,urlduia, Formulario,FechaNac,nacionalidad,fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(Nombrerepresentante, Duirepresentante, parentescop, edadp, nombreAdolescente, telefonoresp, nombre, apellido, sexop, pais, departamentop, municipiop, sedep, tipo, Estudiap, ultimoestudiop, gradop, seccionp, turnop, dispacacidadp, nuevogwp, voz_imgp, DUI, urldocF, urldocA, Formulario, FechaNac, nacionalidadp, currentTimeGMT)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_discapacidad (idp, discapacidad) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dis", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func Educol(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=16")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.sede_areas where Ida=5 and IdPais=16 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where typePerson=1 and active_at is not null and id not in(10)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Educacion/educol.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func posteducol(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	Nombrerepresentante := r.FormValue("nombreRep")
	Duirepresentante := r.FormValue("DUIP")
	parentesco := r.FormValue("parentesco")
	parentescop := cambio(parentesco)
	edad := r.FormValue("edad")
	edadp := cambio(edad)
	nombreAdolescente := r.FormValue("nombreAdolescente")
	telefonoresp := r.FormValue("numResp")
	nombre := r.FormValue("nombreCompleto")
	apellido := r.FormValue("apellido")
	sexo := r.FormValue("Sexo")
	sexop := cambio(sexo)
	pais := 16
	departamento := r.FormValue("pais")
	departamentop := cambio(departamento)
	municipio := r.FormValue("departamento")
	municipiop := cambio(municipio)
	sede := r.FormValue("sede2")
	sedep := cambio(sede)
	tipo := cambio(r.FormValue("tipo"))
	Estudia := r.FormValue("TdC")
	Estudiap := cambio(Estudia)
	ultimoestudio := r.FormValue("UT")
	ultimoestudiop := cambio(ultimoestudio)
	grado := r.FormValue("GA")
	gradop := cambio(grado)
	seccion := r.FormValue("seccion")
	seccionp := cambio(seccion)
	turno := r.FormValue("Turno")
	turnop := cambio(turno)
	discapacidad := r.FormValue("dis")
	dispacacidadp := cambio(discapacidad)
	nuevogw := r.FormValue("gp")
	nuevogwp := cambio(nuevogw)
	voz_img := r.FormValue("aut2")
	voz_imgp := cambio(voz_img)
	DUI := r.FormValue("DUIM")
	urldocF := ""
	urldocA := ""
	Formulario := 2
	FechaNac := r.FormValue("FechaN")
	nacionalidad := r.FormValue("nacionalidad")
	nacionalidadp := cambio(nacionalidad)

	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	if r.FormValue("aut") == "10" {
		cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
		handleError(err)
		// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
		client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
		handleError(err)
		file, fileHandler, err := r.FormFile("docF")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "educacion/col/duif/"+fileName, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		file2, handler2, err := r.FormFile("docA")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "educacion/col/duia/"+fileName2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)

		urldocF = "https://gwforms.blob.core.windows.net/participantes/educacion/col/duif/" + fileName
		urldocA = "https://gwforms.blob.core.windows.net/participantes/educacion/col/duia/" + fileName2
	}

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_participantes (Nombrerepresentante,Duirepresentante,parentesco,edad,nombreAdolescente,telefonoresp,nombre,apellido,sexo,pais,departamento,municipio,sede,tipo,Estudia,ultimoestudio,grado,seccion,turno,discapacidad,nuevogw,voz_img,DUI,urldui,urlduia, Formulario,FechaNac,nacionalidad,fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,? )")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(Nombrerepresentante, Duirepresentante, parentescop, edadp, nombreAdolescente, telefonoresp, nombre, apellido, sexop, pais, departamentop, municipiop, sedep, tipo, Estudiap, ultimoestudiop, gradop, seccionp, turnop, dispacacidadp, nuevogwp, voz_imgp, DUI, urldocF, urldocA, Formulario, FechaNac, nacionalidadp, currentTimeGMT)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_discapacidad (idp, discapacidad) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dis", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func Educr(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=6")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.sede_areas where Ida=5 and IdPais=6 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where typePerson=1 and active_at is not null and id not in(10)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Educacion/educr.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func posteducr(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	Nombrerepresentante := r.FormValue("nombreRep")
	Duirepresentante := r.FormValue("DUIP")
	parentesco := r.FormValue("parentesco")
	parentescop := cambio(parentesco)
	edad := r.FormValue("edad")
	edadp := cambio(edad)
	nombreAdolescente := r.FormValue("nombreAdolescente")
	telefonoresp := r.FormValue("numResp")
	nombre := r.FormValue("nombreCompleto")
	apellido := r.FormValue("apellido")
	sexo := r.FormValue("Sexo")
	sexop := cambio(sexo)
	pais := 6
	departamento := r.FormValue("pais")
	departamentop := cambio(departamento)
	municipio := r.FormValue("departamento")
	municipiop := cambio(municipio)
	sede := r.FormValue("sede2")
	sedep := cambio(sede)
	tipo := cambio(r.FormValue("tipo"))
	Estudia := r.FormValue("TdC")
	Estudiap := cambio(Estudia)
	ultimoestudio := r.FormValue("UT")
	ultimoestudiop := cambio(ultimoestudio)
	grado := r.FormValue("GA")
	gradop := cambio(grado)
	seccion := r.FormValue("seccion")
	seccionp := cambio(seccion)
	turno := r.FormValue("Turno")
	turnop := cambio(turno)
	discapacidad := r.FormValue("dis")
	dispacacidadp := cambio(discapacidad)
	nuevogw := r.FormValue("gp")
	nuevogwp := cambio(nuevogw)
	voz_img := r.FormValue("aut2")
	voz_imgp := cambio(voz_img)
	DUI := r.FormValue("DUIM")
	urldocF := ""
	urldocA := ""
	Formulario := 2
	FechaNac := r.FormValue("FechaN")
	nacionalidad := r.FormValue("nacionalidad")
	nacionalidadp := cambio(nacionalidad)

	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	if r.FormValue("aut") == "10" {
		cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
		handleError(err)
		// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
		client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
		handleError(err)
		file, fileHandler, err := r.FormFile("docF")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "educacion/cr/duif/"+fileName, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		file2, handler2, err := r.FormFile("docA")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "educacion/cr/duia/"+fileName2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)

		urldocF = "https://gwforms.blob.core.windows.net/participantes/educacion/cr/duif/" + fileName
		urldocA = "https://gwforms.blob.core.windows.net/participantes/educacion/cr/duia/" + fileName2
	}

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_participantes (Nombrerepresentante,Duirepresentante,parentesco,edad,nombreAdolescente,telefonoresp,nombre,apellido,sexo,pais,departamento,municipio,sede,tipo,Estudia,ultimoestudio,grado,seccion,turno,discapacidad,nuevogw,voz_img,DUI,urldui,urlduia, Formulario,FechaNac,nacionalidad,fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(Nombrerepresentante, Duirepresentante, parentescop, edadp, nombreAdolescente, telefonoresp, nombre, apellido, sexop, pais, departamentop, municipiop, sedep, tipo, Estudiap, ultimoestudiop, gradop, seccionp, turnop, dispacacidadp, nuevogwp, voz_imgp, DUI, urldocF, urldocA, Formulario, FechaNac, nacionalidadp, currentTimeGMT)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_discapacidad (idp, discapacidad) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dis", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func eduNYeng(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	var db2 *sql.DB
	connString2 := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)

	db2, err = sql.Open("sqlserver", connString2)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db2.Close()
	ctx2 := context.Background()
	err = db2.PingContext(ctx2)
	if err != nil {
		log.Fatal(err.Error())
	}

	paises, err := Controller.GetOpcionesPregunta(db2, "SELECT distinct id, pais_en FROM CatPaises where deleted is null order by pais_en")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT IdSede,Sede FROM admin_main_gwdata.sede_areas where IdPais=8 and Ida=5 order by Sede")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Sexo:         sexo,
		Parentesco:   parentesco,
		Departamento: departamento,
		Paises:       paises,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Educacion/FormulariosNY/edunyenglish.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func eduNYesp(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	var db2 *sql.DB
	connString2 := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)

	db2, err = sql.Open("sqlserver", connString2)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db2.Close()
	ctx2 := context.Background()
	err = db2.PingContext(ctx2)
	if err != nil {
		log.Fatal(err.Error())
	}

	paises, err := Controller.GetOpcionesPregunta(db2, "SELECT distinct id, pais_es FROM CatPaises where deleted is null order by pais_es")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT IdSede,Sede FROM admin_main_gwdata.sede_areas where IdPais=8 and Ida=5 order by Sede")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Sexo:         sexo,
		Parentesco:   parentesco,
		Departamento: departamento,
		Paises:       paises,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Educacion/FormulariosNY/edunyspanish.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func eduNYfr(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	var db2 *sql.DB
	connString2 := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)

	db2, err = sql.Open("sqlserver", connString2)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db2.Close()
	ctx2 := context.Background()
	err = db2.PingContext(ctx2)
	if err != nil {
		log.Fatal(err.Error())
	}

	paises, err := Controller.GetOpcionesPregunta(db2, "SELECT distinct id, pais_fr FROM CatPaises where deleted is null order by pais_fr")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT IdSede,Sede FROM admin_main_gwdata.sede_areas where IdPais=8 and Ida=5 order by Sede")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Sexo:         sexo,
		Parentesco:   parentesco,
		Departamento: departamento,
		Paises:       paises,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Educacion/FormulariosNY/edunyfrench.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostEduNY(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	var idF int = 7
	// Obtener los valores de los campos del formulario

	fechaNac := r.FormValue("FechaN")
	tipoParticipante := cambio(r.FormValue("tipoPerfil"))
	if tipoParticipante == 0 {
		tipoParticipante = 1
	}
	nombres := r.FormValue("Firstmiddle")
	apellidos := r.FormValue("Surnames")
	idSede := cambio(r.FormValue("sede2"))
	pais := r.FormValue("StudentNationality")
	nuevogw := cambio(r.FormValue("autR8"))
	sexo := cambio(r.FormValue("Sexo"))
	grado := cambio(r.FormValue("GA"))
	municipio := cambio(r.FormValue("municipio"))
	autovozimg := cambio(r.FormValue("autR6")) // medios
	turno := cambio(r.FormValue("idioma"))
	fechaMedI := r.FormValue("FechaSt2")
	fechaMedF := r.FormValue("FechaEnd2")
	nacionalidad := cambio(r.FormValue("autR7")) // libera responsabilidad
	estudia := cambio(r.FormValue("autR10"))     // encuesta

	vozimg := cambio(r.FormValue("autR3")) // medios
	nombrecompleto := r.FormValue("nombreRes")
	telefono := r.FormValue("Telcontacto")
	parentesco := cambio(r.FormValue("relacion"))
	fechaMedIR := r.FormValue("FechaSt1")
	fechaMedFR := r.FormValue("FechaEnd1")
	confirmo1 := cambio(r.FormValue("autR11")) //encuesta

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	stmt, err := db.Prepare("IF EXISTS(SELECT * FROM FormularioRegisto WHERE nombres = @P1 and apellidos = @P2 and fechaNAc = @P3 and idF = @P4 and deleted is null)Begin select 1 end else Begin select 0 end")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	row := stmt.QueryRow(nombres, apellidos, fechaNac, idF)

	var result int
	_ = row.Scan(&result)

	var respuesta string

	if result == 1 {

		respuesta = "-1"
		fmt.Fprint(w, respuesta)

	} else {
		tx, err := db.Begin()
		if err != nil {

			tx.Rollback()
			http.Error(w, "Error conexion", http.StatusInternalServerError)

		}

		stmt, err := tx.Prepare("INSERT INTO FormularioRegisto(idF,tipoParticipante,pais,idSede,fechaNac,nombres,apellidos,sexo,turno,grado,nuevogw,autovozimg,municipio,fechaMedI,fechaMedF,nacionalidad, estudia) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14,@P15,@P16,@P17)")
		if err != nil {

			tx.Rollback()
			http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)

		}

		var lastInsertId int
		err = stmt.QueryRow(idF, tipoParticipante, pais, idSede, fechaNac, nombres, apellidos, sexo, turno, grado, nuevogw, autovozimg, municipio, fechaMedI, fechaMedF, nacionalidad, estudia).Scan(&lastInsertId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		stmt, err = tx.Prepare("INSERT INTO AutorizacionMenores(idFR,nombrecompleto, telefono, parentesco, vozimg, fechaMedIR, fechaMedFR, confirmo1) VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}

		_, err = stmt.Exec(lastInsertId, nombrecompleto, telefono, parentesco, vozimg, fechaMedIR, fechaMedFR, confirmo1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		vives := r.Form["club"]
		if len(vives) == 0 {
		} else {
			stmt, err := tx.Prepare("INSERT INTO MultipleClubes (idFR,idRespuesta) VALUES (@P1,@P2)")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				tx.Rollback()
				return

			}
			defer stmt.Close()

			for _, valor := range vives {
				if valor != "" {
					_, err = stmt.Exec(lastInsertId, valor)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						tx.Rollback()
						return

					}
				}
			}
		}

		err = tx.Commit()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if lastInsertId != 0 {
			//AQUI empieza lo del documento

			cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
			handleError(err)
			// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
			client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
			handleError(err)

			// Obtiene la cadena base64 del formulario
			base64String := r.FormValue("Signature2draw")
			if base64String != "" {

				// Divide la cadena en la coma (,) y toma la segunda parte para decodificarla
				base64Data := strings.Split(base64String, ",")[1]

				// Decodifica el archivo base64
				decodedFile, err := base64.StdEncoding.DecodeString(base64Data)
				if err != nil {
					log.Println("Error al decodificar el archivo base64: ", err)
					http.Error(w, "Error al decodificar el archivo base64", http.StatusBadRequest)
					return
				}

				// Convierte los bytes decodificados en una imagen PNG
				img, _, err := image.Decode(bytes.NewReader(decodedFile))
				if err != nil {
					log.Println("Error al decodificar el archivo PNG: ", err)
					http.Error(w, "Error al decodificar el archivo PNG", http.StatusBadRequest)
					return
				}

				// Crea un archivo temporal para guardar la imagen
				tempFile, err := os.CreateTemp("", "upload-*.png")
				if err != nil {
					log.Println("Error al crear el archivo temporal: ", err)
					http.Error(w, "Error al crear el archivo temporal", http.StatusBadRequest)
					return
				}

				defer tempFile.Close()

				// Escribe la imagen en el archivo temporal
				err = png.Encode(tempFile, img)
				if err != nil {
					log.Println("Error al codificar la imagen PNG: ", err)
					http.Error(w, "Error al codificar la imagen PNG", http.StatusBadRequest)
					return
				}

				// Sube el archivo a un blob de bloque
				_, err = client.UploadFile(context.TODO(), "participantes", "educacion/us/firmarparticipante/"+strconv.Itoa(lastInsertId)+".png", tempFile,
					&azblob.UploadFileOptions{
						BlockSize:   int64(1024),
						Concurrency: uint16(3),
						// If Progress is non-nil, this function is called periodically as bytes are uploaded.
						Progress: func(bytesTransferred int64) {
						},
					})
				if err != nil {
					fmt.Printf("Error al subir el archivo: %+v\n", err)
					http.Error(w, "Error al subir el archivo", http.StatusBadRequest)
					return
				}
			}

			// Obtiene la cadena base64 del formulario
			base64String2 := r.FormValue("Signature1draw")
			if base64String2 != "" {

				// Divide la cadena en la coma (,) y toma la segunda parte para decodificarla
				base64Data2 := strings.Split(base64String2, ",")[1]

				// Decodifica el archivo base64
				decodedFile2, err := base64.StdEncoding.DecodeString(base64Data2)
				if err != nil {
					log.Println("Error al decodificar el archivo base64: ", err)
					http.Error(w, "Error al decodificar el archivo base64", http.StatusBadRequest)
					return
				}

				// Convierte los bytes decodificados en una imagen PNG
				img2, _, err := image.Decode(bytes.NewReader(decodedFile2))
				if err != nil {
					log.Println("Error al decodificar el archivo PNG: ", err)
					http.Error(w, "Error al decodificar el archivo PNG", http.StatusBadRequest)
					return
				}

				// Crea un archivo temporal para guardar la imagen
				tempFile2, err := os.CreateTemp("", "upload-*.png")
				if err != nil {
					log.Println("Error al crear el archivo temporal: ", err)
					http.Error(w, "Error al crear el archivo temporal", http.StatusBadRequest)
					return
				}

				defer tempFile2.Close()

				// Escribe la imagen en el archivo temporal
				err = png.Encode(tempFile2, img2)
				if err != nil {
					log.Println("Error al codificar la imagen PNG: ", err)
					http.Error(w, "Error al codificar la imagen PNG", http.StatusBadRequest)
					return
				}

				// Sube el archivo a un blob de bloque
				_, err = client.UploadFile(context.TODO(), "participantes", "educacion/us/firmaresponsable/"+strconv.Itoa(lastInsertId)+".png", tempFile2,
					&azblob.UploadFileOptions{
						BlockSize:   int64(1024),
						Concurrency: uint16(3),
						// If Progress is non-nil, this function is called periodically as bytes are uploaded.
						Progress: func(bytesTransferred int64) {

						},
					})
				if err != nil {
					fmt.Printf("Error al subir el archivo: %+v\n", err)
					http.Error(w, "Error al subir el archivo", http.StatusBadRequest)
					return
				}

			}
		}
		respuesta = "1"
		fmt.Fprint(w, respuesta)
	}
}

func eduNNAcol(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=16")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.sede_areas where Ida=5 and IdPais=16 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Sexo:         sexo,
		Parentesco:   parentesco,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Educacion/eduNNAcol.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func eduNNApan(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.sede_areas where Ida=5 and IdPais=1 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Sexo:         sexo,
		Parentesco:   parentesco,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Educacion/eduNNApan.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostNnaEdu(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	// Obtener los valores de los campos del formulario
	idF := 6
	pais := r.FormValue("idpais")
	idSede := cambio(r.FormValue("sede2"))
	fechaNac := r.FormValue("FechaN")
	nombres := r.FormValue("nombre")
	apellidos := r.FormValue("apellido")
	sexo := cambio(r.FormValue("Sexo"))
	nacionalidad := cambio(r.FormValue("nacionalidad"))
	departamento := cambio(r.FormValue("pais"))
	municipio := cambio(r.FormValue("departamento"))
	discapacidad := cambio(r.FormValue("dis"))
	estudia := cambio(r.FormValue("Eactual"))
	estudioAlcanzado := cambio(r.FormValue("UT"))
	grado := cambio(r.FormValue("GA"))
	turno := cambio(r.FormValue("Turno"))
	seccion := cambio(r.FormValue("seccion"))
	nuevogw := cambio(r.FormValue("anterior"))
	tipoParticipante := 1

	nombrecompleto := r.FormValue("contacto")
	telefono := r.FormValue("Telcontacto")
	parentesco := cambio(r.FormValue("relacion"))
	identidad := r.FormValue("DUIM")
	confirmo1 := cambio(r.FormValue("autR1"))
	infoAdicional1 := r.FormValue("paisNac")

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	stmt, err := db.Prepare("IF EXISTS(SELECT * FROM FormularioRegisto WHERE nombres = @P1 and apellidos = @P2 and identidad = @P3 and pais = @P4 and fechaNAc = @P5 and idF = @P6 and deleted is null)Begin select 1 end else Begin select 0 end")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	row := stmt.QueryRow(nombres, apellidos, identidad, pais, fechaNac, idF)

	var result int
	_ = row.Scan(&result)

	type Datos struct {
		Nombre   string
		Apellido string
		Id       int
		Programa string
	}

	if result == 1 {

		datos := Datos{
			Nombre:   nombres,
			Apellido: apellidos,
		}

		t, err := template.ParseFiles("Public/Html/RespuestaDuplicada.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Renderizar la plantilla con los datos
		err = t.Execute(w, datos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		tx, err := db.Begin()
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		stmt, err := tx.Prepare("INSERT INTO FormularioRegisto(idF,tipoParticipante,pais,idSede,fechaNac,nombres,apellidos,sexo,nacionalidad,discapacidad,estudia,estudioAlcanzado,grado,turno,seccion,nuevogw,departamento,municipio) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14,@P15,@P16,@P17,@P18)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		var lastInsertId int
		err = stmt.QueryRow(idF, tipoParticipante, pais, idSede, fechaNac, nombres, apellidos, sexo, nacionalidad, discapacidad, estudia, estudioAlcanzado, grado, turno, seccion, nuevogw, departamento, municipio).Scan(&lastInsertId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		stmt, err = tx.Prepare("INSERT INTO AutorizacionMenores(idFR,nombrecompleto, telefono, parentesco, identidad, confirmo1, infoAdicional1) VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		_, err = stmt.Exec(lastInsertId, nombrecompleto, telefono, parentesco, identidad, confirmo1, infoAdicional1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		discapacidadmultiple := r.Form["dis2"]
		if len(discapacidadmultiple) == 0 {
		} else {
			stmt, err := tx.Prepare("INSERT INTO MultipleDiscapacidad (idFR,iddiscapacidad) VALUES (@P1,@P2)")
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al preparar la consulta ocupacion", http.StatusInternalServerError)
				return
			}
			defer stmt.Close()

			for _, valor := range discapacidadmultiple {
				if valor != "" {
					_, err = stmt.Exec(lastInsertId, valor)
					if err != nil {
						tx.Rollback()
						http.Error(w, "Error al ejecutar la consulta situaciones", http.StatusInternalServerError)
						return
					}
				}
			}
		}

		err = tx.Commit()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var programa string = "Educación"

		datos := Datos{
			Nombre:   nombres,
			Apellido: apellidos,
			Id:       lastInsertId,
			Programa: programa,
		}

		t, err := template.ParseFiles("Public/Html/RespuestaPersonalizada.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Renderizar la plantilla con los datos
		err = t.Execute(w, datos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func eduNNArd(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=5")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.sede_areas where Ida=5 and IdPais=5 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Sexo:         sexo,
		Parentesco:   parentesco,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Educacion/eduNNArd.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func eduNNAmx(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=17")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.sede_areas where Ida=5 and IdPais=17 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Sexo:         sexo,
		Parentesco:   parentesco,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Educacion/eduNNAmx.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostNNAEducacion(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	// Obtener los valores de los campos del formulario
	idF := 6
	pais := r.FormValue("idpais")
	idSede := cambio(r.FormValue("sede2"))
	fechaNac := r.FormValue("FechaN")
	nombres := r.FormValue("nombre")
	apellidos := r.FormValue("apellido")
	sexo := cambio(r.FormValue("Sexo"))
	nacionalidad := cambio(r.FormValue("nacionalidad"))
	departamento := cambio(r.FormValue("pais"))
	municipio := cambio(r.FormValue("departamento"))
	discapacidad := cambio(r.FormValue("dis"))
	estudia := cambio(r.FormValue("Eactual"))
	estudioAlcanzado := cambio(r.FormValue("UT"))
	grado := cambio(r.FormValue("GA"))
	turno := cambio(r.FormValue("Turno"))
	seccion := cambio(r.FormValue("seccion"))
	nuevogw := cambio(r.FormValue("anterior"))
	club := r.FormValue("club")
	tipoParticipante := 1

	nombrecompleto := r.FormValue("contacto")
	FechaNacRes := r.FormValue("FechaNR")
	telefono := r.FormValue("Telcontacto")
	parentesco := cambio(r.FormValue("relacion"))
	identidad := r.FormValue("DUIM")

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	stmt, err := db.Prepare("IF EXISTS(SELECT * FROM FormularioRegisto WHERE nombres = @P1 and apellidos = @P2 and pais = @P3 and fechaNAc = @P4 and idF = @P5 and deleted is null)Begin select 1 end else Begin select 0 end")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	row := stmt.QueryRow(nombres, apellidos, pais, fechaNac, idF)

	var result int
	_ = row.Scan(&result)

	type Datos struct {
		Nombre   string
		Apellido string
		Id       int
		Programa string
	}

	if result == 1 {

		datos := Datos{
			Nombre:   nombres,
			Apellido: apellidos,
		}

		t, err := template.ParseFiles("Public/Html/RespuestaDuplicada.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Renderizar la plantilla con los datos
		err = t.Execute(w, datos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		tx, err := db.Begin()
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error conexion", http.StatusInternalServerError)

		}

		stmt, err := tx.Prepare("INSERT INTO FormularioRegisto(idF,tipoParticipante,pais,idSede,fechaNac,nombres,apellidos,sexo,nacionalidad,discapacidad,estudia,estudioAlcanzado,grado,turno,seccion,nuevogw,departamento,municipio) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14,@P15,@P16,@P17,@P18)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		var lastInsertId int
		err = stmt.QueryRow(idF, tipoParticipante, pais, idSede, fechaNac, nombres, apellidos, sexo, nacionalidad, discapacidad, estudia, estudioAlcanzado, grado, turno, seccion, nuevogw, departamento, municipio).Scan(&lastInsertId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		stmt, err = tx.Prepare("INSERT INTO AutorizacionMenores(idFR, nombrecompleto, FechaNacRes, telefono, parentesco, identidad) VALUES(@P1,@P2,@P3,@P4,@P5,@P6)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		_, err = stmt.Exec(lastInsertId, nombrecompleto, FechaNacRes, telefono, parentesco, identidad)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		stmt, err = tx.Prepare("INSERT INTO FormEduNNA(idFR,club) VALUES(@P1,@P2)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		_, err = stmt.Exec(lastInsertId, club)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		discapacidadmultiple := r.Form["dis2"]
		if len(discapacidadmultiple) == 0 {
		} else {
			stmt, err := tx.Prepare("INSERT INTO MultipleDiscapacidad (idFR,iddiscapacidad) VALUES (@P1,@P2)")
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al preparar la consulta ocupacion", http.StatusInternalServerError)
				return
			}
			defer stmt.Close()

			for _, valor := range discapacidadmultiple {
				if valor != "" {
					_, err = stmt.Exec(lastInsertId, valor)
					if err != nil {
						tx.Rollback()
						http.Error(w, "Error al ejecutar la consulta situaciones", http.StatusInternalServerError)
						return
					}
				}
			}
		}

		autorizacionmultiple := r.Form["auto"]
		if len(autorizacionmultiple) == 0 {
		} else {
			stmt, err := tx.Prepare("INSERT INTO MultipleAutorizaciones (idFR,autorizacion,idF) VALUES (@P1,@P2,6)")
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al preparar la consulta ocupacion", http.StatusInternalServerError)
				return
			}
			defer stmt.Close()

			for _, valor := range autorizacionmultiple {
				if valor != "" {
					_, err = stmt.Exec(lastInsertId, valor)
					if err != nil {
						tx.Rollback()
						http.Error(w, "Error al ejecutar la consulta situaciones", http.StatusInternalServerError)
						return
					}
				}
			}
		}

		err = tx.Commit()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var programa string = "Educación"

		datos := Datos{
			Nombre:   nombres,
			Apellido: apellidos,
			Id:       lastInsertId,
			Programa: programa,
		}

		t, err := template.ParseFiles("Public/Html/RespuestaPersonalizada.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Renderizar la plantilla con los datos
		err = t.Execute(w, datos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func saludgt(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=4")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM VSedesSalud where IdPais=4 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where typePerson=1 and active_at is not null and id not in(10)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/saludgt.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postesaludgt(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	Nombrerepresentante := r.FormValue("nombreRep")
	Duirepresentante := r.FormValue("DUIP")
	parentesco := r.FormValue("parentesco")
	parentescop := cambio(parentesco)
	edad := r.FormValue("edad")
	edadp := cambio(edad)
	nombreAdolescente := r.FormValue("nombreAdolescente")
	telefonoresp := r.FormValue("numResp")
	nombre := r.FormValue("nombreCompleto")
	apellido := r.FormValue("apellido")
	sexo := r.FormValue("Sexo")
	sexop := cambio(sexo)
	pais := 4
	departamento := r.FormValue("pais")
	departamentop := cambio(departamento)
	municipio := r.FormValue("departamento")
	municipiop := cambio(municipio)
	sede := r.FormValue("sede2")
	sedep := cambio(sede)
	tipo := cambio(r.FormValue("tipo"))
	Estudia := r.FormValue("TdC")
	Estudiap := cambio(Estudia)
	ultimoestudio := r.FormValue("UT")
	ultimoestudiop := cambio(ultimoestudio)
	grado := r.FormValue("GA")
	gradop := cambio(grado)
	seccion := r.FormValue("seccion")
	seccionp := cambio(seccion)
	turno := r.FormValue("Turno")
	turnop := cambio(turno)
	discapacidad := r.FormValue("dis")
	dispacacidadp := cambio(discapacidad)
	nuevogw := r.FormValue("gp")
	nuevogwp := cambio(nuevogw)
	voz_img := r.FormValue("aut2")
	voz_imgp := cambio(voz_img)
	DUI := r.FormValue("DUIM")
	urldocF := ""
	urldocA := ""
	Formulario := 3
	FechaNac := r.FormValue("FechaN")
	nacionalidad := r.FormValue("nacionalidad")
	nacionalidadp := cambio(nacionalidad)
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	if r.FormValue("aut") == "10" {
		cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
		handleError(err)
		// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
		client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
		handleError(err)
		file, fileHandler, err := r.FormFile("docF")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "salud/gt/duif/"+fileName, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		file2, handler2, err := r.FormFile("docA")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "salud/gt/duia/"+fileName2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)

		urldocF = "https://gwforms.blob.core.windows.net/participantes/salud/gt/duif/" + fileName
		urldocA = "https://gwforms.blob.core.windows.net/participantes/salud/gt/duia/" + fileName2
	}

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_participantes (Nombrerepresentante,Duirepresentante,parentesco,edad,nombreAdolescente,telefonoresp,nombre,apellido,sexo,pais,departamento,municipio,sede,tipo,Estudia,ultimoestudio,grado,seccion,turno,discapacidad,nuevogw,voz_img,DUI,urldui,urlduia, Formulario,FechaNac,nacionalidad,fechaRegistro) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(Nombrerepresentante, Duirepresentante, parentescop, edadp, nombreAdolescente, telefonoresp, nombre, apellido, sexop, pais, departamentop, municipiop, sedep, tipo, Estudiap, ultimoestudiop, gradop, seccionp, turnop, dispacacidadp, nuevogwp, voz_imgp, DUI, urldocF, urldocA, Formulario, FechaNac, nacionalidadp, currentTimeGMT)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_discapacidad (idp, discapacidad) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dis", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func saludcol(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=16")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM VSedesSalud where IdPais=16 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where typePerson=1 and active_at is not null and id not in(10)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/saludcol.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postesaludcol(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	Nombrerepresentante := r.FormValue("nombreRep")
	Duirepresentante := r.FormValue("DUIP")
	parentesco := r.FormValue("parentesco")
	parentescop := cambio(parentesco)
	edad := r.FormValue("edad")
	edadp := cambio(edad)
	nombreAdolescente := r.FormValue("nombreAdolescente")
	telefonoresp := r.FormValue("numResp")
	nombre := r.FormValue("nombreCompleto")
	apellido := r.FormValue("apellido")
	sexo := r.FormValue("Sexo")
	sexop := cambio(sexo)
	pais := 16
	departamento := r.FormValue("pais")
	departamentop := cambio(departamento)
	municipio := r.FormValue("departamento")
	municipiop := cambio(municipio)
	sede := r.FormValue("sede2")
	sedep := cambio(sede)
	tipo := cambio(r.FormValue("tipo"))
	Estudia := r.FormValue("TdC")
	Estudiap := cambio(Estudia)
	ultimoestudio := r.FormValue("UT")
	ultimoestudiop := cambio(ultimoestudio)
	grado := r.FormValue("GA")
	gradop := cambio(grado)
	seccion := r.FormValue("seccion")
	seccionp := cambio(seccion)
	turno := r.FormValue("Turno")
	turnop := cambio(turno)
	discapacidad := r.FormValue("dis")
	dispacacidadp := cambio(discapacidad)
	nuevogw := r.FormValue("gp")
	nuevogwp := cambio(nuevogw)
	voz_img := r.FormValue("aut2")
	voz_imgp := cambio(voz_img)
	DUI := r.FormValue("DUIM")
	urldocF := ""
	urldocA := ""
	Formulario := 3
	FechaNac := r.FormValue("FechaN")
	nacionalidad := r.FormValue("nacionalidad")
	nacionalidadp := cambio(nacionalidad)
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	if r.FormValue("aut") == "10" {
		cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
		handleError(err)
		// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
		client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
		handleError(err)
		file, fileHandler, err := r.FormFile("docF")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "salud/col/duif/"+fileName, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		file2, handler2, err := r.FormFile("docA")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "salud/col/duia/"+fileName2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)

		urldocF = "https://gwforms.blob.core.windows.net/participantes/salud/col/duif/" + fileName
		urldocA = "https://gwforms.blob.core.windows.net/participantes/salud/col/duia/" + fileName2
	}

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_participantes (Nombrerepresentante,Duirepresentante,parentesco,edad,nombreAdolescente,telefonoresp,nombre,apellido,sexo,pais,departamento,municipio,sede,tipo,Estudia,ultimoestudio,grado,seccion,turno,discapacidad,nuevogw,voz_img,DUI,urldui,urlduia, Formulario,FechaNac,nacionalidad,fechaRegistro) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(Nombrerepresentante, Duirepresentante, parentescop, edadp, nombreAdolescente, telefonoresp, nombre, apellido, sexop, pais, departamentop, municipiop, sedep, tipo, Estudiap, ultimoestudiop, gradop, seccionp, turnop, dispacacidadp, nuevogwp, voz_imgp, DUI, urldocF, urldocA, Formulario, FechaNac, nacionalidadp, currentTimeGMT)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_discapacidad (idp, discapacidad) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dis", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func saludcr(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=6")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM VSedesSalud where IdPais=6 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where typePerson=1 and active_at is not null and id not in(10)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/saludcr.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postesaludcr(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	Nombrerepresentante := r.FormValue("nombreRep")
	Duirepresentante := r.FormValue("DUIP")
	parentesco := r.FormValue("parentesco")
	parentescop := cambio(parentesco)
	edad := r.FormValue("edad")
	edadp := cambio(edad)
	nombreAdolescente := r.FormValue("nombreAdolescente")
	telefonoresp := r.FormValue("numResp")
	nombre := r.FormValue("nombreCompleto")
	apellido := r.FormValue("apellido")
	sexo := r.FormValue("Sexo")
	sexop := cambio(sexo)
	pais := 6
	departamento := r.FormValue("pais")
	departamentop := cambio(departamento)
	municipio := r.FormValue("departamento")
	municipiop := cambio(municipio)
	sede := r.FormValue("sede2")
	sedep := cambio(sede)
	tipo := cambio(r.FormValue("tipo"))
	Estudia := r.FormValue("TdC")
	Estudiap := cambio(Estudia)
	ultimoestudio := r.FormValue("UT")
	ultimoestudiop := cambio(ultimoestudio)
	grado := r.FormValue("GA")
	gradop := cambio(grado)
	seccion := r.FormValue("seccion")
	seccionp := cambio(seccion)
	turno := r.FormValue("Turno")
	turnop := cambio(turno)
	discapacidad := r.FormValue("dis")
	dispacacidadp := cambio(discapacidad)
	nuevogw := r.FormValue("gp")
	nuevogwp := cambio(nuevogw)
	voz_img := r.FormValue("aut2")
	voz_imgp := cambio(voz_img)
	DUI := r.FormValue("DUIM")
	urldocF := ""
	urldocA := ""
	Formulario := 3
	FechaNac := r.FormValue("FechaN")
	nacionalidad := r.FormValue("nacionalidad")
	nacionalidadp := cambio(nacionalidad)
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	if r.FormValue("aut") == "10" {
		cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
		handleError(err)
		// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
		client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
		handleError(err)
		file, fileHandler, err := r.FormFile("docF")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "salud/cr/duif/"+fileName, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		file2, handler2, err := r.FormFile("docA")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "salud/cr/duia/"+fileName2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)

		urldocF = "https://gwforms.blob.core.windows.net/participantes/salud/cr/duif/" + fileName
		urldocA = "https://gwforms.blob.core.windows.net/participantes/salud/cr/duia/" + fileName2
	}

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_participantes (Nombrerepresentante,Duirepresentante,parentesco,edad,nombreAdolescente,telefonoresp,nombre,apellido,sexo,pais,departamento,municipio,sede,tipo,Estudia,ultimoestudio,grado,seccion,turno,discapacidad,nuevogw,voz_img,DUI,urldui,urlduia, Formulario,FechaNac,nacionalidad,fechaRegistro) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(Nombrerepresentante, Duirepresentante, parentescop, edadp, nombreAdolescente, telefonoresp, nombre, apellido, sexop, pais, departamentop, municipiop, sedep, tipo, Estudiap, ultimoestudiop, gradop, seccionp, turnop, dispacacidadp, nuevogwp, voz_imgp, DUI, urldocF, urldocA, Formulario, FechaNac, nacionalidadp, currentTimeGMT)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_discapacidad (idp, discapacidad) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dis", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func saludhn(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=3")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM VSedesSalud where IdPais=3 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where typePerson=1 and active_at is not null and id not in(10)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/saludhn.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postesaludhn(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	Nombrerepresentante := r.FormValue("nombreRep")
	Duirepresentante := r.FormValue("DUIP")
	parentesco := r.FormValue("parentesco")
	parentescop := cambio(parentesco)
	edad := r.FormValue("edad")
	edadp := cambio(edad)
	nombreAdolescente := r.FormValue("nombreAdolescente")
	telefonoresp := r.FormValue("numResp")
	nombre := r.FormValue("nombreCompleto")
	apellido := r.FormValue("apellido")
	sexo := r.FormValue("Sexo")
	sexop := cambio(sexo)
	pais := 3
	departamento := r.FormValue("pais")
	departamentop := cambio(departamento)
	municipio := r.FormValue("departamento")
	municipiop := cambio(municipio)
	sede := r.FormValue("sede2")
	sedep := cambio(sede)
	tipo := cambio(r.FormValue("tipo"))
	Estudia := r.FormValue("TdC")
	Estudiap := cambio(Estudia)
	ultimoestudio := r.FormValue("UT")
	ultimoestudiop := cambio(ultimoestudio)
	grado := r.FormValue("GA")
	gradop := cambio(grado)
	seccion := r.FormValue("seccion")
	seccionp := cambio(seccion)
	turno := r.FormValue("Turno")
	turnop := cambio(turno)
	discapacidad := r.FormValue("dis")
	dispacacidadp := cambio(discapacidad)
	nuevogw := r.FormValue("gp")
	nuevogwp := cambio(nuevogw)
	voz_img := r.FormValue("aut2")
	voz_imgp := cambio(voz_img)
	DUI := r.FormValue("DUIM")
	urldocF := ""
	urldocA := ""
	Formulario := 3
	FechaNac := r.FormValue("FechaN")
	nacionalidad := r.FormValue("nacionalidad")
	nacionalidadp := cambio(nacionalidad)
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	if r.FormValue("aut") == "10" {
		cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
		handleError(err)
		// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
		client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
		handleError(err)
		file, fileHandler, err := r.FormFile("docF")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "salud/hn/duif/"+fileName, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		file2, handler2, err := r.FormFile("docA")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "salud/hn/duia/"+fileName2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)

		urldocF = "https://gwforms.blob.core.windows.net/participantes/salud/hn/duif/" + fileName
		urldocA = "https://gwforms.blob.core.windows.net/participantes/salud/hn/duia/" + fileName2
	}

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_participantes (Nombrerepresentante,Duirepresentante,parentesco,edad,nombreAdolescente,telefonoresp,nombre,apellido,sexo,pais,departamento,municipio,sede,tipo,Estudia,ultimoestudio,grado,seccion,turno,discapacidad,nuevogw,voz_img,DUI,urldui,urlduia, Formulario,FechaNac,nacionalidad,fechaRegistro) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(Nombrerepresentante, Duirepresentante, parentescop, edadp, nombreAdolescente, telefonoresp, nombre, apellido, sexop, pais, departamentop, municipiop, sedep, tipo, Estudiap, ultimoestudiop, gradop, seccionp, turnop, dispacacidadp, nuevogwp, voz_imgp, DUI, urldocF, urldocA, Formulario, FechaNac, nacionalidadp, currentTimeGMT)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_discapacidad (idp, discapacidad) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dis", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func sanamentemx(w http.ResponseWriter, r *http.Request) { //se quitó , Cfg *Controller.Config
	// Build connection string
	/* 	var db *sql.DB

	   	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	   	var err error
	   	// Create connection pool
	   	db, err = sql.Open("mysql", connString)
	   	if err != nil {
	   		log.Fatal("Error creating connection pool: 1", err.Error())
	   	}
	   	defer db.Close()
	   	// Obtener opciones de la base de datos de MYSQL
	   	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	   	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=17")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	respuestas2, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM admin_main_gwdata.institutional_people where active_at is not null and id not in(13) order by name")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	   	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	   	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=17 order by Departamento")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where typePerson=1 and active_at is not null and id not in(1,3)")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	data := FormData{
	   		Parentesco:   parentesco,
	   		Sexo:         sexo,
	   		Departamento: departamento,
	   		Respuestas:   respuestas,
	   		Ultgrad:      ultgrad,
	   		Grado:        grado,
	   		Turno:        turno,
	   		Seccion:      seccion,
	   		Depsede:      depsede,
	   		Perfil:       perfil,
	   		Respuestas2:  respuestas2,
	   	}
	   	// Cargar plantilla HTML desde archivo
	   	tmpl, err := template.ParseFiles("public/HTML/sanamente/mx/inscripcionmx.html")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	// Renderizar plantilla HTML con opciones de preguntas select
	   	err = tmpl.Execute(w, data)
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	} */

	// Redirigir a la página HTML por encuesta cerrada
	http.ServeFile(w, r, "public/HTML/sanamente/RespuestaSanamenteRecursoCerrado.html")
}

func sanamentecol(w http.ResponseWriter, r *http.Request) { //se quitó , Cfg *Controller.Config
	// Build connection string
	/* 	var db *sql.DB

	   	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	   	var err error
	   	// Create connection pool
	   	db, err = sql.Open("mysql", connString)
	   	if err != nil {
	   		log.Fatal("Error creating connection pool: 1", err.Error())
	   	}
	   	defer db.Close()
	   	// Obtener opciones de la base de datos de MYSQL
	   	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	   	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=16")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	respuestas2, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM admin_main_gwdata.institutional_people where active_at is not null and id not in(13) order by name ")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	   	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	   	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=16 order by Departamento")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where typePerson=1 and active_at is not null and id not in(1,3)")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	data := FormData{
	   		Parentesco:   parentesco,
	   		Sexo:         sexo,
	   		Departamento: departamento,
	   		Respuestas:   respuestas,
	   		Ultgrad:      ultgrad,
	   		Grado:        grado,
	   		Turno:        turno,
	   		Seccion:      seccion,
	   		Depsede:      depsede,
	   		Perfil:       perfil,
	   		Respuestas2:  respuestas2,
	   	}
	   	// Cargar plantilla HTML desde archivo
	   	tmpl, err := template.ParseFiles("public/HTML/sanamente/col/inscripcioncol.html")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	// Renderizar plantilla HTML con opciones de preguntas select
	   	err = tmpl.Execute(w, data)
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	} */

	// Redirigir a la página HTML por encuesta cerrada
	http.ServeFile(w, r, "public/HTML/sanamente/RespuestaSanamenteRecursoCerrado.html")

}

func sanamentehn(w http.ResponseWriter, r *http.Request) { //se quitó , Cfg *Controller.Config
	// Build connection string
	/* 	var db *sql.DB

	   	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	   	var err error
	   	// Create connection pool
	   	db, err = sql.Open("mysql", connString)
	   	if err != nil {
	   		log.Fatal("Error creating connection pool: 1", err.Error())
	   	}
	   	defer db.Close()
	   	// Obtener opciones de la base de datos de MYSQL
	   	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	   	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=3")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	respuestas2, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM admin_main_gwdata.institutional_people where active_at is not null and id not in(13) order by name")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	   	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	   	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=3 order by Departamento")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where id in (10)")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	data := FormData{
	   		Parentesco:   parentesco,
	   		Sexo:         sexo,
	   		Departamento: departamento,
	   		Respuestas:   respuestas,
	   		Ultgrad:      ultgrad,
	   		Grado:        grado,
	   		Turno:        turno,
	   		Seccion:      seccion,
	   		Depsede:      depsede,
	   		Perfil:       perfil,
	   		Respuestas2:  respuestas2,
	   	}
	   	// Cargar plantilla HTML desde archivo
	   	tmpl, err := template.ParseFiles("public/HTML/sanamente/hn/inscripcionhn.html")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	// Renderizar plantilla HTML con opciones de preguntas select
	   	err = tmpl.Execute(w, data)
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	} */

	// Redirigir a la página HTML por encuesta cerrada
	http.ServeFile(w, r, "public/HTML/sanamente/RespuestaSanamenteRecursoCerrado.html")
}

func sanamentegt(w http.ResponseWriter, r *http.Request) { //se quitó , Cfg *Controller.Config
	// Build connection string
	/* 	var db *sql.DB

	   	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	   	var err error
	   	// Create connection pool
	   	db, err = sql.Open("mysql", connString)
	   	if err != nil {
	   		log.Fatal("Error creating connection pool: 1", err.Error())
	   	}
	   	defer db.Close()
	   	// Obtener opciones de la base de datos de MYSQL
	   	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	   	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=4")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	respuestas2, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM admin_main_gwdata.institutional_people where active_at is not null and id not in(13) order by name")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	   	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	   	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=4 order by Departamento")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where id in (10)")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	data := FormData{
	   		Parentesco:   parentesco,
	   		Sexo:         sexo,
	   		Departamento: departamento,
	   		Respuestas:   respuestas,
	   		Ultgrad:      ultgrad,
	   		Grado:        grado,
	   		Turno:        turno,
	   		Seccion:      seccion,
	   		Depsede:      depsede,
	   		Perfil:       perfil,
	   		Respuestas2:  respuestas2,
	   	}
	   	// Cargar plantilla HTML desde archivo
	   	tmpl, err := template.ParseFiles("public/HTML/sanamente/gt/inscripciongt.html")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	// Renderizar plantilla HTML con opciones de preguntas select
	   	err = tmpl.Execute(w, data)
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	} */

	// Redirigir a la página HTML por encuesta cerrada
	http.ServeFile(w, r, "public/HTML/sanamente/RespuestaSanamenteRecursoCerrado.html")
}

func sanamentecr(w http.ResponseWriter, r *http.Request) { //se quitó , Cfg *Controller.Config
	// Build connection string
	/* 	var db *sql.DB

	   	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	   	var err error
	   	// Create connection pool
	   	db, err = sql.Open("mysql", connString)
	   	if err != nil {
	   		log.Fatal("Error creating connection pool: 1", err.Error())
	   	}
	   	defer db.Close()
	   	// Obtener opciones de la base de datos de MYSQL
	   	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	   	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=6")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	respuestas2, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM admin_main_gwdata.institutional_people where active_at is not null and id not in(13) order by name")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	   	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	   	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=6 order by Departamento")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where id in (10)")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	data := FormData{
	   		Parentesco:   parentesco,
	   		Sexo:         sexo,
	   		Departamento: departamento,
	   		Respuestas:   respuestas,
	   		Ultgrad:      ultgrad,
	   		Grado:        grado,
	   		Turno:        turno,
	   		Seccion:      seccion,
	   		Depsede:      depsede,
	   		Perfil:       perfil,
	   		Respuestas2:  respuestas2,
	   	}
	   	// Cargar plantilla HTML desde archivo
	   	tmpl, err := template.ParseFiles("public/HTML/sanamente/cr/inscripcioncr.html")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	// Renderizar plantilla HTML con opciones de preguntas select
	   	err = tmpl.Execute(w, data)
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	} */

	// Redirigir a la página HTML por encuesta cerrada
	http.ServeFile(w, r, "public/HTML/sanamente/RespuestaSanamenteRecursoCerrado.html")
}

func sanamentesv(w http.ResponseWriter, r *http.Request) { //se quitó , Cfg *Controller.Config
	//se agregó la r para inhabilitar el link
	// Build connection string
	/* 	var db *sql.DB

	   	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	   	var err error
	   	// Create connection pool
	   	db, err = sql.Open("mysql", connString)
	   	if err != nil {
	   		log.Fatal("Error creating connection pool: 1", err.Error())
	   	}
	   	defer db.Close()
	   	// Obtener opciones de la base de datos de MYSQL
	   	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	   	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	respuestas2, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM admin_main_gwdata.institutional_people where active_at is not null or id in (14,15) order by name")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	   	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where id in (10)")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=7")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	data := FormData{
	   		Parentesco:   parentesco,
	   		Sexo:         sexo,
	   		Respuestas:   respuestas,
	   		Ultgrad:      ultgrad,
	   		Grado:        grado,
	   		Turno:        turno,
	   		Seccion:      seccion,
	   		Departamento: departamento,
	   		Perfil:       perfil,
	   		Respuestas2:  respuestas2,
	   	}

	   	// Cargar plantilla HTML desde archivo
	   	tmpl, err := template.ParseFiles("public/HTML/sanamente/sv/inscripcionsv.html")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	// Renderizar plantilla HTML con opciones de preguntas select
	   	err = tmpl.Execute(w, data)
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	} */
	// Redirigir a la página HTML por encuesta cerrada
	http.ServeFile(w, r, "public/HTML/sanamente/RespuestaSanamenteRecursoCerrado.html")
}

func sanamentepn(w http.ResponseWriter, r *http.Request) { //se quitó , Cfg *Controller.Config
	// Build connection string
	/* 	var db *sql.DB

	   	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	   	var err error
	   	// Create connection pool
	   	db, err = sql.Open("mysql", connString)
	   	if err != nil {
	   		log.Fatal("Error creating connection pool: 1", err.Error())
	   	}
	   	defer db.Close()
	   	// Obtener opciones de la base de datos de MYSQL
	   	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	   	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=1")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	respuestas2, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM admin_main_gwdata.institutional_people where active_at is not null order by name")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	   	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	   	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=1 order by Departamento")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where id in (10)")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	data := FormData{
	   		Parentesco:   parentesco,
	   		Sexo:         sexo,
	   		Departamento: departamento,
	   		Respuestas:   respuestas,
	   		Ultgrad:      ultgrad,
	   		Grado:        grado,
	   		Turno:        turno,
	   		Seccion:      seccion,
	   		Depsede:      depsede,
	   		Perfil:       perfil,
	   		Respuestas2:  respuestas2,
	   	}
	   	// Cargar plantilla HTML desde archivo
	   	tmpl, err := template.ParseFiles("public/HTML/sanamente/pn/inscripcionpn.html")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	// Renderizar plantilla HTML con opciones de preguntas select
	   	err = tmpl.Execute(w, data)
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	} */

	// Redirigir a la página HTML por encuesta cerrada
	http.ServeFile(w, r, "public/HTML/sanamente/RespuestaSanamenteRecursoCerrado.html")
}

func saludmx(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=17")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM VSedesSalud where IdPais=17 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where typePerson=1 and active_at is not null and id not in(10)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/saludmx.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postesaludmx(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	Nombrerepresentante := r.FormValue("nombreRep")
	Duirepresentante := r.FormValue("DUIP")
	parentesco := r.FormValue("parentesco")
	parentescop := cambio(parentesco)
	edad := r.FormValue("edad")
	edadp := cambio(edad)
	nombreAdolescente := r.FormValue("nombreAdolescente")
	telefonoresp := r.FormValue("numResp")
	nombre := r.FormValue("nombreCompleto")
	apellido := r.FormValue("apellido")
	sexo := r.FormValue("Sexo")
	sexop := cambio(sexo)
	pais := 17
	departamento := r.FormValue("pais")
	departamentop := cambio(departamento)
	municipio := r.FormValue("departamento")
	municipiop := cambio(municipio)
	sede := r.FormValue("sede2")
	sedep := cambio(sede)
	tipo := cambio(r.FormValue("tipo"))
	Estudia := r.FormValue("TdC")
	Estudiap := cambio(Estudia)
	ultimoestudio := r.FormValue("UT")
	ultimoestudiop := cambio(ultimoestudio)
	grado := r.FormValue("GA")
	gradop := cambio(grado)
	seccion := r.FormValue("seccion")
	seccionp := cambio(seccion)
	turno := r.FormValue("Turno")
	turnop := cambio(turno)
	discapacidad := r.FormValue("dis")
	dispacacidadp := cambio(discapacidad)
	nuevogw := r.FormValue("gp")
	nuevogwp := cambio(nuevogw)
	voz_img := r.FormValue("aut2")
	voz_imgp := cambio(voz_img)
	DUI := r.FormValue("DUIM")
	urldocF := ""
	urldocA := ""
	Formulario := 3
	FechaNac := r.FormValue("FechaN")
	nacionalidad := r.FormValue("nacionalidad")
	nacionalidadp := cambio(nacionalidad)
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	if r.FormValue("aut") == "10" {
		cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
		handleError(err)
		// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
		client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
		handleError(err)
		file, fileHandler, err := r.FormFile("docF")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "salud/mx/duif/"+fileName, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		file2, handler2, err := r.FormFile("docA")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "salud/mx/duia/"+fileName2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)

		urldocF = "https://gwforms.blob.core.windows.net/participantes/salud/mx/duif/" + fileName
		urldocA = "https://gwforms.blob.core.windows.net/participantes/salud/mx/duia/" + fileName2
	}

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_participantes (Nombrerepresentante,Duirepresentante,parentesco,edad,nombreAdolescente,telefonoresp,nombre,apellido,sexo,pais,departamento,municipio,sede,tipo,Estudia,ultimoestudio,grado,seccion,turno,discapacidad,nuevogw,voz_img,DUI,urldui,urlduia, Formulario,FechaNac,nacionalidad,fechaRegistro) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(Nombrerepresentante, Duirepresentante, parentescop, edadp, nombreAdolescente, telefonoresp, nombre, apellido, sexop, pais, departamentop, municipiop, sedep, tipo, Estudiap, ultimoestudiop, gradop, seccionp, turnop, dispacacidadp, nuevogwp, voz_imgp, DUI, urldocF, urldocA, Formulario, FechaNac, nacionalidadp, currentTimeGMT)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_discapacidad (idp, discapacidad) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dis", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func saludpn(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM VSedesSalud where IdPais=1 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where typePerson=1 and active_at is not null and id not in(10)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/saludpn.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postesaludpn(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	Nombrerepresentante := r.FormValue("nombreRep")
	Duirepresentante := r.FormValue("DUIP")
	parentesco := r.FormValue("parentesco")
	parentescop := cambio(parentesco)
	edad := r.FormValue("edad")
	edadp := cambio(edad)
	nombreAdolescente := r.FormValue("nombreAdolescente")
	telefonoresp := r.FormValue("numResp")
	nombre := r.FormValue("nombreCompleto")
	apellido := r.FormValue("apellido")
	sexo := r.FormValue("Sexo")
	sexop := cambio(sexo)
	pais := 1
	departamento := r.FormValue("pais")
	departamentop := cambio(departamento)
	municipio := r.FormValue("departamento")
	municipiop := cambio(municipio)
	sede := r.FormValue("sede2")
	sedep := cambio(sede)
	tipo := cambio(r.FormValue("tipo"))
	Estudia := r.FormValue("TdC")
	Estudiap := cambio(Estudia)
	ultimoestudio := r.FormValue("UT")
	ultimoestudiop := cambio(ultimoestudio)
	grado := r.FormValue("GA")
	gradop := cambio(grado)
	seccion := r.FormValue("seccion")
	seccionp := cambio(seccion)
	turno := r.FormValue("Turno")
	turnop := cambio(turno)
	discapacidad := r.FormValue("dis")
	dispacacidadp := cambio(discapacidad)
	nuevogw := r.FormValue("gp")
	nuevogwp := cambio(nuevogw)
	voz_img := r.FormValue("aut2")
	voz_imgp := cambio(voz_img)
	DUI := r.FormValue("DUIM")
	urldocF := ""
	urldocA := ""
	Formulario := 3
	FechaNac := r.FormValue("FechaN")
	nacionalidad := r.FormValue("nacionalidad")
	nacionalidadp := cambio(nacionalidad)
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	if r.FormValue("aut") == "10" {
		cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
		handleError(err)
		// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
		client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
		handleError(err)
		file, fileHandler, err := r.FormFile("docF")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "salud/pn/duif/"+fileName, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		file2, handler2, err := r.FormFile("docA")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "salud/pn/duia/"+fileName2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)

		urldocF = "https://gwforms.blob.core.windows.net/participantes/salud/pn/duif/" + fileName
		urldocA = "https://gwforms.blob.core.windows.net/participantes/salud/pn/duia/" + fileName2
	}

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_participantes (Nombrerepresentante,Duirepresentante,parentesco,edad,nombreAdolescente,telefonoresp,nombre,apellido,sexo,pais,departamento,municipio,sede,tipo,Estudia,ultimoestudio,grado,seccion,turno,discapacidad,nuevogw,voz_img,DUI,urldui,urlduia, Formulario,FechaNac,nacionalidad,fechaRegistro) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(Nombrerepresentante, Duirepresentante, parentescop, edadp, nombreAdolescente, telefonoresp, nombre, apellido, sexop, pais, departamentop, municipiop, sedep, tipo, Estudiap, ultimoestudiop, gradop, seccionp, turnop, dispacacidadp, nuevogwp, voz_imgp, DUI, urldocF, urldocA, Formulario, FechaNac, nacionalidadp, currentTimeGMT)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_discapacidad (idp, discapacidad) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dis", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func saludrd(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=5")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM VSedesSalud where IdPais=5 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where typePerson=1 and active_at is not null and id not in(10)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/saludrd.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postesaludrd(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	Nombrerepresentante := r.FormValue("nombreRep")
	Duirepresentante := r.FormValue("DUIP")
	parentesco := r.FormValue("parentesco")
	parentescop := cambio(parentesco)
	edad := r.FormValue("edad")
	edadp := cambio(edad)
	nombreAdolescente := r.FormValue("nombreAdolescente")
	telefonoresp := r.FormValue("numResp")
	nombre := r.FormValue("nombreCompleto")
	apellido := r.FormValue("apellido")
	sexo := r.FormValue("Sexo")
	sexop := cambio(sexo)
	pais := 5
	departamento := r.FormValue("pais")
	departamentop := cambio(departamento)
	municipio := r.FormValue("departamento")
	municipiop := cambio(municipio)
	sede := r.FormValue("sede2")
	sedep := cambio(sede)
	tipo := cambio(r.FormValue("tipo"))
	Estudia := r.FormValue("TdC")
	Estudiap := cambio(Estudia)
	ultimoestudio := r.FormValue("UT")
	ultimoestudiop := cambio(ultimoestudio)
	grado := r.FormValue("GA")
	gradop := cambio(grado)
	seccion := r.FormValue("seccion")
	seccionp := cambio(seccion)
	turno := r.FormValue("Turno")
	turnop := cambio(turno)
	discapacidad := r.FormValue("dis")
	dispacacidadp := cambio(discapacidad)
	nuevogw := r.FormValue("gp")
	nuevogwp := cambio(nuevogw)
	voz_img := r.FormValue("aut2")
	voz_imgp := cambio(voz_img)
	DUI := r.FormValue("DUIM")
	urldocF := ""
	urldocA := ""
	Formulario := 3
	FechaNac := r.FormValue("FechaN")
	nacionalidad := r.FormValue("nacionalidad")
	nacionalidadp := cambio(nacionalidad)
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	if r.FormValue("aut") == "10" {
		cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
		handleError(err)
		// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
		client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
		handleError(err)
		file, fileHandler, err := r.FormFile("docF")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "salud/rd/duif/"+fileName, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		file2, handler2, err := r.FormFile("docA")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "participantes", "salud/rd/duia/"+fileName2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)

		urldocF = "https://gwforms.blob.core.windows.net/participantes/salud/rd/duif/" + fileName
		urldocA = "https://gwforms.blob.core.windows.net/participantes/salud/rd/duia/" + fileName2
	}

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_participantes (Nombrerepresentante,Duirepresentante,parentesco,edad,nombreAdolescente,telefonoresp,nombre,apellido,sexo,pais,departamento,municipio,sede,tipo,Estudia,ultimoestudio,grado,seccion,turno,discapacidad,nuevogw,voz_img,DUI,urldui,urlduia, Formulario,FechaNac,nacionalidad,fechaRegistro) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(Nombrerepresentante, Duirepresentante, parentescop, edadp, nombreAdolescente, telefonoresp, nombre, apellido, sexop, pais, departamentop, municipiop, sedep, tipo, Estudiap, ultimoestudiop, gradop, seccionp, turnop, dispacacidadp, nuevogwp, voz_imgp, DUI, urldocF, urldocA, Formulario, FechaNac, nacionalidadp, currentTimeGMT)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_discapacidad (idp, discapacidad) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dis", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func formularioMentores(w http.ResponseWriter, Cfg *Controller.Config) {

	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de  MySQL
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paisesn, err := Controller.GetOpcionesPregunta(db, "SELECT distinct id, name FROM admin_main_gwdata.countries order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paises, err := Controller.GetOpcionesPregunta(db, "SELECT distinct id, name FROM admin_main_gwdata.countries where id in(1,3,4,6,7,8) order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd, departamento FROM admin_main_gwdata.geografia where idp=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	empresas, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.F_catEmpresas where activo = 1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Crear estructura de datos con opciones para cada pregunta select
	data := FormData{
		Sexo:         sexo,
		Departamento: departamento,
		Paises:       paises,
		Paisesn:      paisesn,
		Empresas:     empresas,
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Voluntarios/Mentores/mentores.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func formulariosamsung(w http.ResponseWriter, Cfg *Controller.Config) {

	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de  MySQL
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paises, err := Controller.GetOpcionesPregunta(db, "SELECT distinct id, name FROM admin_main_gwdata.countries where id in(1,24,3,4,5,6,7,15,2) order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Crear estructura de datos con opciones para cada pregunta select
	data := FormData{
		Sexo:   sexo,
		Paises: paises,
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Voluntarios/samsung.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func formularioMentoresMillicom(w http.ResponseWriter, Cfg *Controller.Config) {

	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de  MySQL
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paisesn, err := Controller.GetOpcionesPregunta(db, "SELECT distinct id, name FROM admin_main_gwdata.countries order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paises, err := Controller.GetOpcionesPregunta(db, "SELECT distinct id, name FROM admin_main_gwdata.countries where id in(1,3,4,6,7,8) order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd, departamento FROM admin_main_gwdata.geografia where idp=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Crear estructura de datos con opciones para cada pregunta select
	data := FormData{
		Sexo:         sexo,
		Departamento: departamento,
		Paises:       paises,
		Paisesn:      paisesn,
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Voluntarios/Mentores/mentoresmillicom.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func participanteJLI(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	transporte, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_transporte where activo=1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estipendio, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_estipendio")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estado_civ, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_estadociv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actividad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_ofertavol")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.Vista_Sedes where IdSede in(666,1128)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grupo, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_grupoSocial")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pueblo, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_puebloComunidad")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	idioma, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_idioma")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actLaboral, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_actividadLaboral")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Transporte:   transporte,
		Estipendio:   estipendio,
		Estado_civ:   estado_civ,
		Actividades:  actividad,
		Depsede:      depsede,
		Grupo:        grupo,
		Idioma:       idioma,
		ActLaboral:   actLaboral,
		Pueblo:       pueblo,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/JLI/jli.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func participanteJLI2(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	transporte, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_transporte where activo=1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estipendio, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_estipendio")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estado_civ, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_estadociv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actividad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_ofertavol")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.Vista_Sedes where IdSede in(666,1128)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grupo, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_grupoSocial")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pueblo, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_puebloComunidad")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	idioma, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_idioma")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actLaboral, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_actividadLaboral")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Transporte:   transporte,
		Estipendio:   estipendio,
		Estado_civ:   estado_civ,
		Actividades:  actividad,
		Depsede:      depsede,
		Grupo:        grupo,
		Idioma:       idioma,
		ActLaboral:   actLaboral,
		Pueblo:       pueblo,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/JLI/jli2.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func participanteJLIcho5(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	transporte, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_transporte where activo=1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estipendio, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_estipendio")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estado_civ, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_estadociv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actividad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_ofertavol")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.Vista_Sedes where IdSede in(666,1128)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grupo, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_grupoSocial")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pueblo, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_puebloComunidad")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	idioma, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_idioma")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actLaboral, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_actividadLaboral")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Transporte:   transporte,
		Estipendio:   estipendio,
		Estado_civ:   estado_civ,
		Actividades:  actividad,
		Depsede:      depsede,
		Grupo:        grupo,
		Idioma:       idioma,
		ActLaboral:   actLaboral,
		Pueblo:       pueblo,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/JLI/jlicho5.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func participanteJLIcho6(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	transporte, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_transporte where activo=1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estipendio, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_estipendio")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estado_civ, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_estadociv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actividad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_ofertavol")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.Vista_Sedes where IdSede in(666,1128)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grupo, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_grupoSocial")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pueblo, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_puebloComunidad")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	idioma, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_idioma")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actLaboral, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_actividadLaboral")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Transporte:   transporte,
		Estipendio:   estipendio,
		Estado_civ:   estado_civ,
		Actividades:  actividad,
		Depsede:      depsede,
		Grupo:        grupo,
		Idioma:       idioma,
		ActLaboral:   actLaboral,
		Pueblo:       pueblo,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/JLI/jlicho6.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func obtenerevento(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	pais := r.FormValue("pais")
	var opcionesCiudades string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT id,nombreEvento FROM admin_main_gwdata.F_eventos where paisId=? ", pais)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionesCiudades += ("<option value=''></option>")
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesCiudades += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionesCiudades)
}

func postmemtees(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	FechaNac := r.FormValue("FechaN")
	nombre := r.FormValue("nombreCompleto")
	apellido := r.FormValue("apellido")
	sexo := cambio(r.FormValue("Sexo"))
	pais := cambio(r.FormValue("pais"))
	departamento := cambio(r.FormValue("departamento"))
	municipio := cambio(r.FormValue("municipio"))
	direccion := r.FormValue("direccion")
	DUI := r.FormValue("DUIM")
	telefono := r.FormValue("TM")
	contacto_what := cambio(r.FormValue("WS"))
	telefono_ws := r.FormValue("TM2")
	correo := r.FormValue("correo")
	grupoSocial := cambio(r.FormValue("grupo"))
	puebloComunidad := cambio(r.FormValue("indigena"))
	idioma := cambio(r.FormValue("idioma"))
	facilidadHablar := cambio(r.FormValue("hablar"))
	facilidadEscribir := cambio(r.FormValue("escribir"))
	facilidadComprension := cambio(r.FormValue("entender"))
	Estudia := cambio(r.FormValue("TdC"))
	ultimoestudio := cambio(r.FormValue("UT"))
	grado := cambio(r.FormValue("GA"))
	seccion := cambio(r.FormValue("seccion"))
	turno := cambio(r.FormValue("Turno"))
	trabaja := cambio(r.FormValue("trabajas"))
	actividadLaboral := cambio(r.FormValue("enque"))
	experienciaLaboral := cambio(r.FormValue("hijos2"))
	areaLaboral := cambio(r.FormValue("area"))
	discapacidad := cambio(r.FormValue("dis"))
	nuevogw := cambio(r.FormValue("gp"))
	actividadparticipo := cambio(r.FormValue("ActG"))
	nombreE := r.FormValue("contacto")
	contactoE := r.FormValue("contacto2")
	parentescoE := cambio(r.FormValue("relacion"))
	estadoCivil := cambio(r.FormValue("EstadoC"))
	transporte := cambio(r.FormValue("Transp"))
	otro_gasto := r.FormValue("cantidad")
	hijos := cambio(r.FormValue("hijos"))
	sede := cambio(r.FormValue("sede2"))
	mentoriaSabado := cambio(r.FormValue("sabados"))
	voz_img := cambio(r.FormValue("aut2"))
	Formulario := 4
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_participantes (FechaNac,nombre,apellido,sexo,pais,departamento,municipio,direccion,DUI,telefono,contacto_what,telefono_ws,correo,grupoSocial,puebloComunidad,idioma,facilidadHablar,facilidadEscribir,facilidadComprension,Estudia,ultimoestudio,grado,seccion,turno,trabaja,actividadLaboral,experienciaLaboral,areaLaboral,discapacidad,nuevogw,actividadparticipo,nombreE,contactoE,parentescoE,estadoCivil,transporte,otro_gasto,hijos,sede,mentoriaSabado,fechaRegistro,voz_img,Formulario) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(FechaNac, nombre, apellido, sexo, pais, departamento, municipio, direccion, DUI, telefono, contacto_what, telefono_ws, correo, grupoSocial, puebloComunidad, idioma, facilidadHablar, facilidadEscribir, facilidadComprension, Estudia, ultimoestudio, grado, seccion, turno, trabaja, actividadLaboral, experienciaLaboral, areaLaboral, discapacidad, nuevogw, actividadparticipo, nombreE, contactoE, parentescoE, estadoCivil, transporte, otro_gasto, hijos, sede, mentoriaSabado, currentTimeGMT, voz_img, Formulario)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else if len(discapacidad2) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO F_discapacidad (idp, discapacidad) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dis", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
		}
	}

	vive := r.Form["vives"]
	if len(vive) == 0 {
	} else if len(vive) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO F_vive (idp, vive) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta vives", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range vive {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta vives", http.StatusInternalServerError)
				return
			}
		}
	}

	dispositivos := r.Form["dispositivo"]
	if len(dispositivos) == 0 {
	} else if len(dispositivos) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO F_dispositivos (idp, idDispositivo) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dispositivo", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range dispositivos {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta dispositivo", http.StatusInternalServerError)
				return
			}
		}
	}

	horarios := r.Form["horario"]
	if len(horarios) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_horario (idp, idHorario) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta horario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range horarios {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta horario", http.StatusInternalServerError)
				return
			}
		}
	}

	sabados := r.Form["sabados2"]
	if len(sabados) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_sabado (idp, idSabado) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta sabados", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range sabados {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta sabados", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func postmentores(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario

	FechaNac := r.FormValue("FechaN")
	nacionalidad := cambio(r.FormValue("pais5"))
	paisNac := cambio(r.FormValue("pais4"))
	paisId := cambio(r.FormValue("pais3"))
	nombre := r.FormValue("nombre")
	apellido := r.FormValue("apellido")
	sexoId := cambio(r.FormValue("sexo"))
	donante := r.FormValue("empresa")
	titulo := r.FormValue("titulo")
	telefono_ws := r.FormValue("numws")
	voz_img := cambio(r.FormValue("aut3"))
	Formulario := 2
	correo := r.FormValue("correo")
	tiempoLibre := r.FormValue("pasatiempo")
	cohorte := 0
	orgCali := r.FormValue("CaliOrg")
	nombEmpresa := r.FormValue("nombEmpresa")
	identidad := r.FormValue("identidad")

	if donante == "1" {
		cohorte = 2
	} else if donante == "2" {
		cohorte = 1
	}

	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_Voluntariado (FechaNac,nacionalidad,paisNac,paisId,nombre,apellido,sexoId,donante,titulo,telefono_ws,voz_img,Formulario,correo,tiempoLibre,fechaRegistro,cohorte,nombreRepresentante,parentescoId,numIdentVoluntario) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?,?,?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(FechaNac, nacionalidad, paisNac, paisId, nombre, apellido, sexoId, donante, titulo, telefono_ws, voz_img, Formulario, correo, tiempoLibre, currentTimeGMT, cohorte, orgCali, nombEmpresa, identidad)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	r.ParseForm()

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func postsamsung(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario

	FechaNac := r.FormValue("FechaN")
	paisId := cambio(r.FormValue("pais3"))
	nombre := r.FormValue("nombre")
	apellido := r.FormValue("apellido")
	sexoId := cambio(r.FormValue("sexo"))
	donante := r.FormValue("empresa")
	titulo := r.FormValue("titulo")
	telefono_ws := r.FormValue("numws")
	voz_img := cambio(r.FormValue("aut3"))
	Formulario := 9
	correo := r.FormValue("correo")
	numIdentVoluntario := r.FormValue("dui")
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_Voluntariado (FechaNac,paisId,nombre,apellido,sexoId,donante,titulo,telefono_ws,voz_img,Formulario,correo,fechaRegistro,numIdentVoluntario) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(FechaNac, paisId, nombre, apellido, sexoId, donante, titulo, telefono_ws, voz_img, Formulario, correo, currentTimeGMT, numIdentVoluntario)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	r.ParseForm()

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func EncuestaSM(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paises, err := Controller.GetOpcionesPregunta(db, "SELECT distinct id, name FROM admin_main_gwdata.countries where id in(3,4,7) order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	afectacion, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=2")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actitud, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=4")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estrategia, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=5")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	practica, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=6")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	estado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	afirmacion, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=8")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lugar, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=9")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Respuestas: respuestas,
		Paises:     paises,
		Sexo:       sexo,
		Afectacion: afectacion,
		Actitud:    actitud,
		Estrategia: estrategia,
		Practica:   practica,
		Estado:     estado,
		Afirmacion: afirmacion,
		Lugar:      lugar,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/sv/Encuesta.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func EncuestaSM2(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paises, err := Controller.GetOpcionesPregunta(db, "SELECT distinct id, name FROM admin_main_gwdata.countries where id in(3,4,7) order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	afectacion, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=2")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actitud, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=4")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estrategia, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=5")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	practica, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=6")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	estado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	afirmacion, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=8")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lugar, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=9")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	solucion, err := getOpcionesPreguntaNew(db, "SELECT distinct id, Respuesta, idPregunta FROM admin_main_gwdata.Respuestas_formularios where idFormulario=9")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		id = "0"
	}

	type CombinedData struct {
		Data  FormData
		Data2 FormDataNew
		Data3 string
	}

	data := FormData{
		Respuestas: respuestas,
		Paises:     paises,
		Sexo:       sexo,
		Afectacion: afectacion,
		Actitud:    actitud,
		Estrategia: estrategia,
		Practica:   practica,
		Estado:     estado,
		Afirmacion: afirmacion,
		Lugar:      lugar,
	}

	data2 := FormDataNew{
		Solucion: solucion,
	}

	// Crear la estructura combinada
	combinedData := CombinedData{
		Data:  data,
		Data2: data2,
		Data3: id,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/sv/Encuesta2.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, combinedData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func EncuestaSM3(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string

	/* 	var db *sql.DB

	   	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	   	var err error
	   	// Create connection pool
	   	db, err = sql.Open("mysql", connString)
	   	if err != nil {
	   		log.Fatal("Error creating connection pool: 1", err.Error())
	   	}
	   	ctx := context.Background()
	   	err = db.PingContext(ctx)
	   	if err != nil {
	   		log.Fatal(err.Error())
	   	}

	   	// Obtener opciones de la base de datos de MYSQL

	   	paises, err := Controller.GetOpcionesPregunta(db, "SELECT distinct id, name FROM admin_main_gwdata.countries where id in(3,4,7) order by name")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	estrategia, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=4 and version = 3")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	practica, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=5 and version = 3")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	solucion, err := getOpcionesPreguntaNew(db, "SELECT distinct id, Respuesta, idPregunta FROM admin_main_gwdata.Respuestas_formularios where idFormulario=9 and version = 3")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	id := r.URL.Query().Get("id")
	   	if id == "" {
	   		id = "0"
	   	}

	   	type CombinedData struct {
	   		Data  FormData
	   		Data2 FormDataNew
	   		Data3 string
	   	}

	   	data := FormData{
	   		Paises:     paises,
	   		Sexo:       sexo,
	   		Estrategia: estrategia,
	   		Practica:   practica,
	   	}

	   	data2 := FormDataNew{
	   		Solucion: solucion,
	   	}

	   	// Crear la estructura combinada
	   	combinedData := CombinedData{
	   		Data:  data,
	   		Data2: data2,
	   		Data3: id,
	   	}
	   	// Cargar plantilla HTML desde archivo
	   	tmpl, err := template.ParseFiles("public/HTML/sanamente/sv/Encuesta3.html")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	// Renderizar plantilla HTML con opciones de preguntas select
	   	err = tmpl.Execute(w, combinedData)
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	} */

	// Redirigir a la página HTML por encuesta cerrada
	http.ServeFile(w, r, "public/HTML/sanamente/RespuestaSanamenteEncuestaCerrada.html")
}

func EncuestaSMsv(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paises, err := Controller.GetOpcionesPregunta(db, "SELECT distinct id, name FROM admin_main_gwdata.countries where id in(3,4,7) order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	afectacion, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=2")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actitud, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=4")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estrategia, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=5")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	practica, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=6")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	estado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	afirmacion, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=8")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lugar, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=9")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Respuestas: respuestas,
		Paises:     paises,
		Sexo:       sexo,
		Afectacion: afectacion,
		Actitud:    actitud,
		Estrategia: estrategia,
		Practica:   practica,
		Estado:     estado,
		Afirmacion: afirmacion,
		Lugar:      lugar,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/sv/Encuestasv.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postEncuestaSM(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	//NombreCampo en la BD: = r.FormValue("NombreCampo en el html")
	// Obtener los valores de los campos del formulario

	participa := r.FormValue("participar")
	idPais := cambio(r.FormValue("pais"))
	lugarTrabajo := r.FormValue("trabajo")
	edad := cambio(r.FormValue("edad"))
	anioLaboral := cambio(r.FormValue("aniol"))
	sexo := cambio(r.FormValue("sexo"))
	rolTrabajo := r.FormValue("puesto")
	afir1 := cambio(r.FormValue("afirma1"))
	afir2 := cambio(r.FormValue("afirma2"))
	afir3 := cambio(r.FormValue("afirma3"))
	afir4 := cambio(r.FormValue("afirma4"))
	afir5 := cambio(r.FormValue("afirma5"))
	afir6 := cambio(r.FormValue("afirma6"))
	afir7 := cambio(r.FormValue("afirma7"))
	afir8 := cambio(r.FormValue("afirma8"))
	afir9 := cambio(r.FormValue("afirma9"))
	afec1 := cambio(r.FormValue("afecta1"))
	afec2 := cambio(r.FormValue("afecta2"))
	afec3 := cambio(r.FormValue("afecta3"))
	actitud := cambio(r.FormValue("actitud"))
	estrategia := cambio(r.FormValue("estrategia"))
	apoyo := r.FormValue("apoyo")
	cambioVida := r.FormValue("cambiov")
	motivo := r.FormValue("motivo")
	conoceTecnica := cambio(r.FormValue("conocet"))
	practicaTecnica := cambio(r.FormValue("practicat"))
	estEmocional := cambio(r.FormValue("estadoe"))
	nivAngustia := cambio(r.FormValue("angustia"))
	oportunidad := r.FormValue("oportunidad")
	climaLaboral := cambio(r.FormValue("clima"))
	conoceEfectos := cambio(r.FormValue("conocee"))
	conscienteEfectos := cambio(r.FormValue("consciente"))
	beneficio := r.FormValue("beneficio")
	rel1 := cambio(r.FormValue("rel1"))
	rel2 := cambio(r.FormValue("rel2"))
	rel3 := cambio(r.FormValue("rel3"))
	rel4 := cambio(r.FormValue("rel4"))
	telefono := r.FormValue("telefono")

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO Sanamente_ResEncuesta (participa,idPais,lugarTrabajo,edad,anioLaboral,sexo,rolTrabajo,afir1,afir2,afir3,afir4,afir5,afir6,afir7,afir8,afir9,afec1,afec2,afec3,actitud,estrategia,apoyo,cambioVida,motivo,conoceTecnica,practicaTecnica,estEmocional,nivAngustia,oportunidad,climaLaboral,conoceEfectos,conscienteEfectos,beneficio,rel1,rel2,rel3,rel4,fechaRegistro,telefono) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(participa, idPais, lugarTrabajo, edad, anioLaboral, sexo, rolTrabajo, afir1, afir2, afir3, afir4, afir5, afir6, afir7, afir8, afir9, afec1, afec2, afec3, actitud, estrategia, apoyo, cambioVida, motivo, conoceTecnica, practicaTecnica, estEmocional, nivAngustia, oportunidad, climaLaboral, conoceEfectos, conscienteEfectos, beneficio, rel1, rel2, rel3, rel4, currentTimeGMT, telefono)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	detalle1 := r.Form["afecta"]
	if len(detalle1) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_ResEncuestad (idp, idRespuesta) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta afectaciones", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range detalle1 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta afectaciones", http.StatusInternalServerError)
				return
			}
		}
	}

	detalle2 := r.Form["epractica"]
	if len(detalle2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_ResEncuestad (idp, idRespuesta) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta estrategia practicada", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range detalle2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta estrategia practicada", http.StatusInternalServerError)
				return
			}
		}
	}

	detalle3 := r.Form["opractica"]
	if len(detalle3) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_ResEncuestad (idp, idRespuesta) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta oportunidad de practica", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range detalle3 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta oportunidad de practica", http.StatusInternalServerError)
				return
			}
		}
	}

	detalle4 := r.Form["ptrabajo"]
	if len(detalle4) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_ResEncuestad (idp, idRespuesta) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta practica en el trabajo", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range detalle4 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta practica en el trabajo", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	t, err := template.ParseFiles("public/HTML/RespuestaSanamente.html")
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

func postEncuestaSM2(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	//NombreCampo en la BD: = r.FormValue("NombreCampo en el html")
	// Obtener los valores de los campos del formulario

	participa := r.FormValue("participar")
	idPais := cambio(r.FormValue("pais"))
	lugarTrabajo := r.FormValue("trabajo")
	edad := cambio(r.FormValue("edad"))
	anioLaboral := cambio(r.FormValue("aniol"))
	sexo := cambio(r.FormValue("sexo"))
	perfil := cambio(r.FormValue("rol"))
	rolTrabajo := r.FormValue("puesto")
	afir1 := cambio(r.FormValue("afirma1"))
	afir2 := cambio(r.FormValue("afirma2"))
	afir3 := cambio(r.FormValue("afirma3"))
	afir4 := cambio(r.FormValue("afirma4"))
	afir5 := cambio(r.FormValue("afirma5"))
	afir6 := cambio(r.FormValue("afirma6"))
	afir7 := cambio(r.FormValue("afirma7"))
	afir8 := cambio(r.FormValue("afirma8"))
	afir9 := cambio(r.FormValue("afirma9"))
	afir1b := cambio(r.FormValue("afirma1B"))
	afir2b := cambio(r.FormValue("afirma2B"))
	afir3b := cambio(r.FormValue("afirma3B"))
	afir4b := cambio(r.FormValue("afirma4B"))
	afir5b := cambio(r.FormValue("afirma5B"))
	afir6b := cambio(r.FormValue("afirma6B"))
	afir7b := cambio(r.FormValue("afirma7B"))
	afir8b := cambio(r.FormValue("afirma8B"))
	afir9b := cambio(r.FormValue("afirma9B"))
	afec1 := cambio(r.FormValue("afecta1"))
	afec2 := cambio(r.FormValue("afecta2"))
	afec3 := cambio(r.FormValue("afecta3"))
	actitud := cambio(r.FormValue("actitud"))
	estrategia := cambio(r.FormValue("estrategia"))
	apoyo := cambio(r.FormValue("apoyo"))
	cambioVida := r.FormValue("cambiov")
	motivo := r.FormValue("motivo")
	motivo2 := r.FormValue("motivo2")
	conoceTecnica := cambio(r.FormValue("conocet"))
	practicaTecnica := cambio(r.FormValue("practicat1"))
	autocuido := cambio(r.FormValue("practicat"))
	estEmocional := cambio(r.FormValue("estadoe"))
	nivAngustia := cambio(r.FormValue("angustia"))
	oportunidad := cambio(r.FormValue("op7"))
	climaLaboral := r.FormValue("clima")
	conoceEfectos := cambio(r.FormValue("conocee"))
	conscienteEfectos := cambio(r.FormValue("consciente"))
	beneficio := r.FormValue("beneficio")
	rel1 := cambio(r.FormValue("rel1"))
	rel2 := cambio(r.FormValue("rel2"))
	rel3 := cambio(r.FormValue("rel3"))
	rel4 := cambio(r.FormValue("rel4"))
	rel1b := cambio(r.FormValue("rel12"))
	rel2b := cambio(r.FormValue("rel22"))
	rel3b := cambio(r.FormValue("rel32"))
	rel4b := cambio(r.FormValue("rel42"))
	p81 := cambio(r.FormValue("81"))
	p82 := cambio(r.FormValue("82"))
	op9 := cambio(r.FormValue("op9"))
	Pid := cambio(r.FormValue("id"))
	version := 4

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO Sanamente_ResEncuesta (participa,idPais,lugarTrabajo,edad,anioLaboral,sexo,perfil,rolTrabajo,afir1,afir2,afir3,afir4,afir5,afir6,afir7,afir8,afir9,afir1b,afir2b,afir3b,afir4b,afir5b,afir6b,afir7b,afir8b,afir9b,afec1,afec2,afec3,actitud,estrategia,apoyo,cambioVida,motivo,motivo2,conoceTecnica,practicaTecnica,estEmocional,nivAngustia,oportunidad,climaLaboral,conoceEfectos,conscienteEfectos,beneficio,rel1,rel2,rel3,rel4,rel1b,rel2b,rel3b,rel4b,p81,p82,op9,Pid,version_,autocuido) OUTPUT INSERTED.ID VALUES (@p1,@p2,@p3,@p4,@p5,@p6,@p7,@p8,@p9,@p10,@p11,@p12,@p13,@p14,@p15,@p16,@p17,@p18,@p19,@p20,@p21,@p22,@p23,@p24,@p25,@p26,@p27,@p28,@p29,@p30,@p31,@p32,@p33,@p34,@p35,@p36,@p37,@p38,@p39,@p40,@p41,@p42,@p43,@p44,@p45,@p46,@p47,@p48,@p49,@p50,@p51,@p52,@p53,@p54,@p55,@p56,@p57,@p58)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()
	var lastInsertId int
	err = stmt.QueryRow(participa, idPais, lugarTrabajo, edad, anioLaboral, sexo, perfil, rolTrabajo, afir1, afir2, afir3, afir4, afir5, afir6, afir7, afir8, afir9, afir1b, afir2b, afir3b, afir4b, afir5b, afir6b, afir7b, afir8b, afir9b, afec1, afec2, afec3, actitud, estrategia, apoyo, cambioVida, motivo, motivo2, conoceTecnica, practicaTecnica, estEmocional, nivAngustia, oportunidad, climaLaboral, conoceEfectos, conscienteEfectos, beneficio, rel1, rel2, rel3, rel4, rel1b, rel2b, rel3b, rel4b, p81, p82, op9, Pid, version, autocuido).Scan(&lastInsertId)
	if err != nil {
		tx.Rollback()
		http.Error(w, fmt.Sprintf("Error al ejecutar la consulta: %v", err), http.StatusInternalServerError)
		return
	}

	r.ParseForm()

	detalle1 := r.Form["op10"]
	if len(detalle1) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_ResEncuestad (idp, idRespuesta) VALUES (@p1, @p2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta afectaciones", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range detalle1 {
			_, err = stmt2.Exec(lastInsertId, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta afectaciones", http.StatusInternalServerError)
				return
			}
		}
	}

	detalle2 := r.Form["epractica"]
	if len(detalle2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_ResEncuestad (idp, idRespuesta) VALUES (@p1, @p2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta estrategia practicada", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range detalle2 {
			_, err = stmt2.Exec(lastInsertId, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta estrategia practicada", http.StatusInternalServerError)
				return
			}
		}
	}

	detalle3 := r.Form["opractica"]
	if len(detalle3) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_ResEncuestad (idp, idRespuesta) VALUES (@p1, @p2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta oportunidad de practica", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range detalle3 {
			_, err = stmt2.Exec(lastInsertId, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta oportunidad de practica", http.StatusInternalServerError)
				return
			}
		}
	}

	detalle4 := r.Form["op101"]
	if len(detalle4) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_ResEncuestad (idp, idRespuesta) VALUES (@p1, @p2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta practica en el trabajo", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range detalle4 {
			_, err = stmt2.Exec(lastInsertId, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta practica en el trabajo", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	t, err := template.ParseFiles("public/HTML/RespuestaSanamente.html")
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

func postEncuestaSM3(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	//NombreCampo en la BD: = r.FormValue("NombreCampo en el html")
	// Obtener los valores de los campos del formulario
	idPais := cambio(r.FormValue("pais"))
	participa := r.FormValue("participar")
	lugarTrabajo := r.FormValue("trabajo")
	edad := cambio(r.FormValue("edad"))
	anioLaboral := cambio(r.FormValue("aniol"))
	sexo := cambio(r.FormValue("sexo"))
	perfil := cambio(r.FormValue("rol"))
	afec1 := cambio(r.FormValue("afecta1"))
	afec2 := cambio(r.FormValue("afecta2"))
	afec3 := cambio(r.FormValue("afecta3"))
	actitud := cambio(r.FormValue("actitud"))
	estrategia := cambio(r.FormValue("estrategia"))
	aplicadoP := cambio(r.FormValue("aplicadoP"))
	otraEstrategia := r.FormValue("otraEstrategia")
	apoyo := cambio(r.FormValue("apoyo"))
	otroApoyo := r.FormValue("motivo")
	autocuido := cambio(r.FormValue("practicat"))
	personasAt := cambio(r.FormValue("personasAt1"))
	estudiantesAt := cambio(r.FormValue("personasAt2"))
	opinionAtend := cambio(r.FormValue("opinionAtend"))
	frec1a := cambio(r.FormValue("frecuencia1"))
	frec2a := cambio(r.FormValue("frecuencia2"))
	frec3a := cambio(r.FormValue("frecuencia3"))
	frec4a := cambio(r.FormValue("frecuencia4"))
	frec5a := cambio(r.FormValue("frecuencia5"))
	frec6a := cambio(r.FormValue("frecuencia6"))
	frec7a := cambio(r.FormValue("frecuencia7"))
	frec8a := cambio(r.FormValue("frecuencia8"))
	frec9a := cambio(r.FormValue("frecuencia9"))
	frec1b := cambio(r.FormValue("frecuencia11"))
	frec2b := cambio(r.FormValue("frecuencia12"))
	frec3b := cambio(r.FormValue("frecuencia13"))
	frec4b := cambio(r.FormValue("frecuencia14"))
	frec5b := cambio(r.FormValue("frecuencia15"))
	frec6b := cambio(r.FormValue("frecuencia16"))
	frec7b := cambio(r.FormValue("frecuencia17"))
	frec8b := cambio(r.FormValue("frecuencia18"))
	porcentajeAtend := cambio(r.FormValue("porcentaje"))
	referido := cambio(r.FormValue("82"))
	actitudAtend := cambio(r.FormValue("practicat1"))
	beneficios := r.FormValue("motivo2")
	afir1 := cambio(r.FormValue("afirmaciones1"))
	afir2 := cambio(r.FormValue("afirmaciones2"))
	afir3 := cambio(r.FormValue("afirmaciones3"))
	utilidad := cambio(r.FormValue("utilidad"))
	ejemplos := r.FormValue("ejemplos")
	resolucion := cambio(r.FormValue("resolucion"))
	conociaEfectos := cambio(r.FormValue("conocia"))
	conscienteEfectos := cambio(r.FormValue("consciente"))
	conocimiento := r.FormValue("beneficio")
	rel1 := cambio(r.FormValue("rel1"))
	rel2 := cambio(r.FormValue("rel2"))
	rel3 := cambio(r.FormValue("rel3"))
	rel4 := cambio(r.FormValue("rel4"))
	rel1b := cambio(r.FormValue("rel12"))
	rel2b := cambio(r.FormValue("rel22"))
	rel3b := cambio(r.FormValue("rel32"))
	rel4b := cambio(r.FormValue("rel42"))
	Pid := cambio(r.FormValue("id"))
	version_ := 5

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO [dbo].[Sanamente_ResEncuestaFGSM] (idPais,participa,lugarTrabajo,edad,anioLaboral,sexo,perfil,afec1,afec2,afec3,actitud,estrategia,aplicadoP,otraEstrategia,apoyo,otroApoyo,autocuido,personasAt,estudiantesAt,opinionAtend,frec1a,frec2a,frec3a,frec4a,frec5a,frec6a,frec7a,frec8a,frec9a,frec1b,frec2b,frec3b,frec4b,frec5b,frec6b,frec7b,frec8b,porcentajeAtend,referido,actitudAtend,beneficios,afir1,afir2,afir3,utilidad,ejemplos,resolucion,conociaEfectos,conscienteEfectos,conocimiento,rel1,rel2,rel3,rel4,rel1b,rel2b,rel3b,rel4b,Pid,version_) OUTPUT INSERTED.ID VALUES (@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14,@P15,@P16,@P17,@P18,@P19,@P20,@P21,@P22,@P23,@P24,@P25,@P26,@P27,@P28,@P29,@P30,@P31,@P32,@P33,@P34,@P35,@P36,@P37,@P38,@P39,@P40,@P41,@P42,@P43,@P44,@P45,@P46,@P47,@P48,@P49,@P50,@P51,@P52,@P53,@P54,@P55,@P56,@P57,@P58,@P59,@P60)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()
	var lastInsertId int
	err = stmt.QueryRow(idPais, participa, lugarTrabajo, edad, anioLaboral, sexo, perfil, afec1, afec2, afec3, actitud, estrategia, aplicadoP, otraEstrategia, apoyo, otroApoyo, autocuido, personasAt, estudiantesAt, opinionAtend, frec1a, frec2a, frec3a, frec4a, frec5a, frec6a, frec7a, frec8a, frec9a, frec1b, frec2b, frec3b, frec4b, frec5b, frec6b, frec7b, frec8b, porcentajeAtend, referido, actitudAtend, beneficios, afir1, afir2, afir3, utilidad, ejemplos, resolucion, conociaEfectos, conscienteEfectos, conocimiento, rel1, rel2, rel3, rel4, rel1b, rel2b, rel3b, rel4b, Pid, version_).Scan(&lastInsertId)
	if err != nil {
		tx.Rollback()
		http.Error(w, fmt.Sprintf("Error al ejecutar la consulta: %v", err), http.StatusInternalServerError)
		return
	}

	r.ParseForm()

	detalle1 := r.Form["epractica"]
	if len(detalle1) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_ResEncuestadFGSM (idp, idPregunta, idRespuesta, version_) VALUES (@p1, @p2, @p3, @p4)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta afectaciones", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range detalle1 {
			_, err = stmt2.Exec(lastInsertId, 5, valor, 5)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta afectaciones", http.StatusInternalServerError)
				return
			}
		}
	}

	detalle2 := r.Form["op101"]
	if len(detalle2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_ResEncuestadFGSM (idp, idPregunta, idRespuesta, version_) VALUES (@p1, @p2, @p3, @p4)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta estrategia practicada", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range detalle2 {
			_, err = stmt2.Exec(lastInsertId, 16, valor, 5)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta estrategia practicada", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	t, err := template.ParseFiles("public/HTML/RespuestaSanamenteEncuesta.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar la plantilla con los datos
	err = t.Execute(w, lastInsertId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func postMentoresMillicom(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario

	FechaNac := r.FormValue("FechaN")
	nacionalidad := cambio(r.FormValue("pais5"))
	paisNac := cambio(r.FormValue("pais4"))
	paisId := cambio(r.FormValue("pais3"))
	nombre := r.FormValue("nombre")
	apellido := r.FormValue("apellido")
	sexoId := cambio(r.FormValue("sexo"))
	titulo := r.FormValue("titulo")
	telefono_ws := r.FormValue("numws")
	voz_img := cambio(r.FormValue("aut3"))
	Formulario := 8
	correo := r.FormValue("correo")
	tiempoLibre := r.FormValue("pasatiempo")
	cohorte := cambio(r.FormValue("cohorte"))

	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_Voluntariado (FechaNac,nacionalidad,paisNac,paisId,nombre,apellido,sexoId,titulo,telefono_ws,voz_img,Formulario,correo,tiempoLibre,fechaRegistro,cohorte) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(FechaNac, nacionalidad, paisNac, paisId, nombre, apellido, sexoId, titulo, telefono_ws, voz_img, Formulario, correo, tiempoLibre, currentTimeGMT, cohorte)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	r.ParseForm()

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func postParticipamas(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	// Obtener los valores de los campos del formulario
	Nombrerepresentante := r.FormValue("nombreRep")
	Duirepresentante := r.FormValue("DUIP")
	parentesco := r.FormValue("parentesco")
	parentescop := cambio(parentesco)
	edad := r.FormValue("edad")
	edadp := cambio(edad)
	nombreAdolescente := r.FormValue("nombreAdolescente")
	telefonoresp := r.FormValue("numResp")
	nombre := r.FormValue("nombreCompleto")
	apellido := r.FormValue("apellido")
	sexo := r.FormValue("Sexo")
	sexop := cambio(sexo)
	pais := 7
	departamento := r.FormValue("pais")
	departamentop := cambio(departamento)
	municipio := r.FormValue("departamento")
	municipiop := cambio(municipio)
	sede := r.FormValue("sede2")
	tipo := cambio(r.FormValue("tipo"))
	Estudia := r.FormValue("TdC")
	Estudiap := cambio(Estudia)
	ultimoestudio := r.FormValue("UT")
	ultimoestudiop := cambio(ultimoestudio)
	grado := r.FormValue("GA")
	gradop := cambio(grado)
	seccion := r.FormValue("seccion")
	seccionp := cambio(seccion)
	turno := r.FormValue("Turno")
	turnop := cambio(turno)
	discapacidad := r.FormValue("dis")
	dispacacidadp := cambio(discapacidad)
	nuevogw := r.FormValue("gp")
	nuevogwp := cambio(nuevogw)
	voz_img := r.FormValue("aut2")
	voz_imgp := cambio(voz_img)
	DUI := r.FormValue("DUIM")
	descripcion := r.FormValue("descripcion")
	porque := r.FormValue("porque")
	municipiosede := r.FormValue("departamento2")
	departamentosede := r.FormValue("pais2")
	urldocF := ""
	urldocA := ""
	Formulario := 5
	FechaNac := r.FormValue("FechaN")

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_participantes (Nombrerepresentante,Duirepresentante,parentesco,edad,nombreAdolescente,telefonoresp,nombre,apellido,sexo,pais,departamento,municipio,asociacion,tipo,Estudia,ultimoestudio,grado,seccion,turno,discapacidad,nuevogw,voz_img,DUI,urldui,urlduia, Formulario,FechaNac,fechaRegistro,descripcion,motivop,depAsociacion,munAsociacion) VALUES ( ?,?,?,?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,? )")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(Nombrerepresentante, Duirepresentante, parentescop, edadp, nombreAdolescente, telefonoresp, nombre, apellido, sexop, pais, departamentop, municipiop, sede, tipo, Estudiap, ultimoestudiop, gradop, seccionp, turnop, dispacacidadp, nuevogwp, voz_imgp, DUI, urldocF, urldocA, Formulario, FechaNac, currentTimeGMT, descripcion, porque, departamentosede, municipiosede)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_discapacidad (idp, discapacidad) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dis", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func postvolpn(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	fechaNac := r.FormValue("FechaN")
	nacionalidad := cambio(r.FormValue("gi"))
	nombreRepresentante := r.FormValue("nombreRep")
	numIdentResponsable := r.FormValue("DUIP")
	parentescoId := cambio(r.FormValue("parentesco"))
	nombreVoluntarioMenor := r.FormValue("nombreAdolescente")
	edadVoluntario := cambio(r.FormValue("edad"))
	tipoDocIndentRespExtId := cambio(r.FormValue("elec"))
	numIdentVoluntario := r.FormValue("DUI")
	nombre := r.FormValue("nombre")
	apellido := r.FormValue("apellido")
	sexoId := cambio(r.FormValue("sexo"))
	numContacto := r.FormValue("TM")
	correo := r.FormValue("correo")
	tallaCamisaId := cambio(r.FormValue("camisa"))
	tiempoVoluntariadoId := cambio(r.FormValue("voluntario"))
	paisNac := cambio(r.FormValue("paisn"))
	deptoId := cambio(r.FormValue("pais"))
	municipioId := cambio(r.FormValue("departamento"))
	direccion := r.FormValue("direccion")
	estudiaId := cambio(r.FormValue("TdC"))
	ultimoEstudioId := cambio(r.FormValue("UT"))
	gradoEstudioId := cambio(r.FormValue("GA"))
	universidad := r.FormValue("U")
	carreraUniv := r.FormValue("carrera")
	horasSocId := cambio(r.FormValue("HS"))
	refPersonal := r.FormValue("nombrer")
	numRefPersonal := r.FormValue("telr")
	refPersonalb := r.FormValue("nombrer2")
	numRefPersonalb := r.FormValue("telr2")
	perfilVoluntariadoId := cambio(r.FormValue("acc"))
	voluntarioComunitarioId := cambio(r.FormValue("acc2"))
	donante := cambio(r.FormValue("acc3"))
	areaDepartamento := r.FormValue("puesto")
	fechaInduccion := r.FormValue("FechaI")
	nombreDesarrollador := r.FormValue("nombrei")
	areaId := cambio(r.FormValue("pais2"))
	actividadId := cambio(r.FormValue("actividad"))
	subActividadId := cambio(r.FormValue("subact"))
	horarioActividad := r.FormValue("OHT")
	deptoActividadId := cambio(r.FormValue("departamento2"))
	sedeActividadId := cambio(r.FormValue("sede2"))
	paisId := 1
	otroClub := r.FormValue("aT")
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)
	urldocF := ""
	urldocA := ""
	urlpass := ""

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_Voluntariado (fechaNac,nacionalidad,nombreRepresentante,numIdentResponsable,parentescoId,nombreVoluntarioMenor,edadVoluntario,tipoDocIndentRespExtId,numIdentVoluntario,nombre,apellido,sexoId,numContacto,correo,tallaCamisaId,tiempoVoluntariadoId,paisNac,deptoId,municipioId,direccion,estudiaId,ultimoEstudioId,gradoEstudioId,universidad,carreraUniv,horasSocId,refPersonal,numRefPersonal,refPersonalb,numRefPersonalb,perfilVoluntariadoId,voluntarioComunitarioId,donante,areaDepartamento,fechaInduccion,nombreDesarrollador,areaId,actividadId,subActividadId,horarioActividad,deptoActividadId,sedeActividadId,otroClub,fechaRegistro,urlDuiFrontal,urlDuiDorso,paisId,docExtranjero) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(fechaNac, nacionalidad, nombreRepresentante, numIdentResponsable, parentescoId, nombreVoluntarioMenor, edadVoluntario, tipoDocIndentRespExtId, numIdentVoluntario, nombre, apellido, sexoId, numContacto, correo, tallaCamisaId, tiempoVoluntariadoId, paisNac, deptoId, municipioId, direccion, estudiaId, ultimoEstudioId, gradoEstudioId, universidad, carreraUniv, horasSocId, refPersonal, numRefPersonal, refPersonalb, numRefPersonalb, perfilVoluntariadoId, voluntarioComunitarioId, donante, areaDepartamento, fechaInduccion, nombreDesarrollador, areaId, actividadId, subActividadId, horarioActividad, deptoActividadId, sedeActividadId, otroClub, currentTimeGMT, urldocF, urldocA, paisId, urlpass)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO f_dias (voluntarioId, diasId) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dias", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta dias", http.StatusInternalServerError)
				return
			}
		}
	}

	discapacidad4 := r.Form["dis4"]
	if len(discapacidad4) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO f_horarios (voluntarioId, horarioId) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta horario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad4 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta horario", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
	handleError(err)
	// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
	client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
	handleError(err)

	file, fileHandler, err := r.FormFile("docF")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename
		fileExt := filepath.Ext(fileName)

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "pn/duif/"+strconv.Itoa(int(idGenerado))+fileExt, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urldocF = strconv.Itoa(int(idGenerado)) + fileExt
	}
	file2, handler2, err := r.FormFile("docA")
	if err != http.ErrMissingFile {

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename
		fileExt2 := filepath.Ext(fileName2)

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "pn/duia/"+strconv.Itoa(int(idGenerado))+fileExt2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		urldocA = strconv.Itoa(int(idGenerado)) + fileExt2
	}

	file3, fileHandler3, err := r.FormFile("pas")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file3.Close()

		// Obtiene el nombre del archivo cargado
		fileName3 := fileHandler3.Filename
		fileExt3 := filepath.Ext(fileName3)

		fileBytes3, err := io.ReadAll(file3)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile3, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile3.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile3.Write(fileBytes3); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "pn/pasaporte/"+strconv.Itoa(int(idGenerado))+fileExt3, tempFile3,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urlpass = strconv.Itoa(int(idGenerado)) + fileExt3
	}

	tx2, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt4, err := tx2.Prepare("UPDATE F_Voluntariado SET urlDuiFrontal = ?, urlDuiDorso = ? ,docExtranjero= ? WHERE id = ?")
	if err != nil {
		tx2.Rollback()
		http.Error(w, "Error al preparar update", http.StatusInternalServerError)
		return
	}
	defer stmt4.Close()

	_, err = stmt4.Exec(urldocF, urldocA, urlpass, idGenerado)
	if err != nil {
		tx2.Rollback()
		http.Error(w, "Error al ejecutar ejecutar update", http.StatusInternalServerError)
		return
	}

	r.ParseForm()

	err = tx2.Commit()
	if err != nil {
		tx2.Rollback()
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

func postvolrd(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	fechaNac := r.FormValue("FechaN")
	nacionalidad := cambio(r.FormValue("gi"))
	nombreRepresentante := r.FormValue("nombreRep")
	numIdentResponsable := r.FormValue("DUIP")
	parentescoId := cambio(r.FormValue("parentesco"))
	nombreVoluntarioMenor := r.FormValue("nombreAdolescente")
	edadVoluntario := cambio(r.FormValue("edad"))
	tipoDocIndentRespExtId := cambio(r.FormValue("elec"))
	numIdentVoluntario := r.FormValue("DUI")
	nombre := r.FormValue("nombre")
	apellido := r.FormValue("apellido")
	sexoId := cambio(r.FormValue("sexo"))
	numContacto := r.FormValue("TM")
	correo := r.FormValue("correo")
	tallaCamisaId := cambio(r.FormValue("camisa"))
	tiempoVoluntariadoId := cambio(r.FormValue("voluntario"))
	paisNac := cambio(r.FormValue("paisn"))
	deptoId := cambio(r.FormValue("pais"))
	municipioId := cambio(r.FormValue("departamento"))
	direccion := r.FormValue("direccion")
	estudiaId := cambio(r.FormValue("TdC"))
	ultimoEstudioId := cambio(r.FormValue("UT"))
	gradoEstudioId := cambio(r.FormValue("GA"))
	universidad := r.FormValue("U")
	carreraUniv := r.FormValue("carrera")
	horasSocId := cambio(r.FormValue("HS"))
	refPersonal := r.FormValue("nombrer")
	numRefPersonal := r.FormValue("telr")
	refPersonalb := r.FormValue("nombrer2")
	numRefPersonalb := r.FormValue("telr2")
	perfilVoluntariadoId := cambio(r.FormValue("acc"))
	voluntarioComunitarioId := cambio(r.FormValue("acc2"))
	donante := cambio(r.FormValue("acc3"))
	areaDepartamento := r.FormValue("puesto")
	fechaInduccion := r.FormValue("FechaI")
	nombreDesarrollador := r.FormValue("nombrei")
	areaId := cambio(r.FormValue("pais2"))
	actividadId := cambio(r.FormValue("actividad"))
	subActividadId := cambio(r.FormValue("subact"))
	horarioActividad := r.FormValue("OHT")
	deptoActividadId := cambio(r.FormValue("departamento2"))
	sedeActividadId := cambio(r.FormValue("sede2"))
	paisId := 5
	otroClub := r.FormValue("aT")
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)
	urldocF := ""
	urldocA := ""
	urlpass := ""

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_Voluntariado (fechaNac,nacionalidad,nombreRepresentante,numIdentResponsable,parentescoId,nombreVoluntarioMenor,edadVoluntario,tipoDocIndentRespExtId,numIdentVoluntario,nombre,apellido,sexoId,numContacto,correo,tallaCamisaId,tiempoVoluntariadoId,paisNac,deptoId,municipioId,direccion,estudiaId,ultimoEstudioId,gradoEstudioId,universidad,carreraUniv,horasSocId,refPersonal,numRefPersonal,refPersonalb,numRefPersonalb,perfilVoluntariadoId,voluntarioComunitarioId,donante,areaDepartamento,fechaInduccion,nombreDesarrollador,areaId,actividadId,subActividadId,horarioActividad,deptoActividadId,sedeActividadId,otroClub,fechaRegistro,urlDuiFrontal,urlDuiDorso,paisId,docExtranjero) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(fechaNac, nacionalidad, nombreRepresentante, numIdentResponsable, parentescoId, nombreVoluntarioMenor, edadVoluntario, tipoDocIndentRespExtId, numIdentVoluntario, nombre, apellido, sexoId, numContacto, correo, tallaCamisaId, tiempoVoluntariadoId, paisNac, deptoId, municipioId, direccion, estudiaId, ultimoEstudioId, gradoEstudioId, universidad, carreraUniv, horasSocId, refPersonal, numRefPersonal, refPersonalb, numRefPersonalb, perfilVoluntariadoId, voluntarioComunitarioId, donante, areaDepartamento, fechaInduccion, nombreDesarrollador, areaId, actividadId, subActividadId, horarioActividad, deptoActividadId, sedeActividadId, otroClub, currentTimeGMT, urldocF, urldocA, paisId, urlpass)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO f_dias (voluntarioId, diasId) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dias", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta dias", http.StatusInternalServerError)
				return
			}
		}
	}

	discapacidad4 := r.Form["dis4"]
	if len(discapacidad4) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO f_horarios (voluntarioId, horarioId) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta horario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad4 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta horario", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
	handleError(err)
	// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
	client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
	handleError(err)

	file, fileHandler, err := r.FormFile("docF")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename
		fileExt := filepath.Ext(fileName)

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "rd/duif/"+strconv.Itoa(int(idGenerado))+fileExt, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urldocF = strconv.Itoa(int(idGenerado)) + fileExt
	}
	file2, handler2, err := r.FormFile("docA")
	if err != http.ErrMissingFile {

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename
		fileExt2 := filepath.Ext(fileName2)

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "rd/duia/"+strconv.Itoa(int(idGenerado))+fileExt2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		urldocA = strconv.Itoa(int(idGenerado)) + fileExt2
	}

	file3, fileHandler3, err := r.FormFile("pas")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file3.Close()

		// Obtiene el nombre del archivo cargado
		fileName3 := fileHandler3.Filename
		fileExt3 := filepath.Ext(fileName3)

		fileBytes3, err := io.ReadAll(file3)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile3, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile3.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile3.Write(fileBytes3); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "rd/pasaporte/"+strconv.Itoa(int(idGenerado))+fileExt3, tempFile3,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urlpass = strconv.Itoa(int(idGenerado)) + fileExt3
	}

	tx2, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt4, err := tx2.Prepare("UPDATE F_Voluntariado SET urlDuiFrontal = ?, urlDuiDorso = ? ,docExtranjero= ? WHERE id = ?")
	if err != nil {
		tx2.Rollback()
		http.Error(w, "Error al preparar update", http.StatusInternalServerError)
		return
	}
	defer stmt4.Close()

	_, err = stmt4.Exec(urldocF, urldocA, urlpass, idGenerado)
	if err != nil {
		tx2.Rollback()
		http.Error(w, "Error al ejecutar ejecutar update", http.StatusInternalServerError)
		return
	}

	r.ParseForm()

	err = tx2.Commit()
	if err != nil {
		tx2.Rollback()
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

func postvolcol(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	fechaNac := r.FormValue("FechaN")
	nacionalidad := cambio(r.FormValue("gi"))
	nombreRepresentante := r.FormValue("nombreRep")
	numIdentResponsable := r.FormValue("DUIP")
	parentescoId := cambio(r.FormValue("parentesco"))
	nombreVoluntarioMenor := r.FormValue("nombreAdolescente")
	edadVoluntario := cambio(r.FormValue("edad"))
	tipoDocIndentRespExtId := cambio(r.FormValue("elec"))
	numIdentVoluntario := r.FormValue("DUI")
	nombre := r.FormValue("nombre")
	apellido := r.FormValue("apellido")
	sexoId := cambio(r.FormValue("sexo"))
	numContacto := r.FormValue("TM")
	correo := r.FormValue("correo")
	tallaCamisaId := cambio(r.FormValue("camisa"))
	tiempoVoluntariadoId := cambio(r.FormValue("voluntario"))
	paisNac := cambio(r.FormValue("paisn"))
	deptoId := cambio(r.FormValue("pais"))
	municipioId := cambio(r.FormValue("departamento"))
	direccion := r.FormValue("direccion")
	estudiaId := cambio(r.FormValue("TdC"))
	ultimoEstudioId := cambio(r.FormValue("UT"))
	gradoEstudioId := cambio(r.FormValue("GA"))
	universidad := r.FormValue("U")
	carreraUniv := r.FormValue("carrera")
	horasSocId := cambio(r.FormValue("HS"))
	refPersonal := r.FormValue("nombrer")
	numRefPersonal := r.FormValue("telr")
	refPersonalb := r.FormValue("nombrer2")
	numRefPersonalb := r.FormValue("telr2")
	perfilVoluntariadoId := cambio(r.FormValue("acc"))
	voluntarioComunitarioId := cambio(r.FormValue("acc2"))
	donante := cambio(r.FormValue("acc3"))
	areaDepartamento := r.FormValue("puesto")
	fechaInduccion := r.FormValue("FechaI")
	nombreDesarrollador := r.FormValue("nombrei")
	areaId := cambio(r.FormValue("pais2"))
	actividadId := cambio(r.FormValue("actividad"))
	subActividadId := cambio(r.FormValue("subact"))
	horarioActividad := r.FormValue("OHT")
	deptoActividadId := cambio(r.FormValue("departamento2"))
	sedeActividadId := cambio(r.FormValue("sede2"))
	paisId := 16
	otroClub := r.FormValue("aT")
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)
	urldocF := ""
	urldocA := ""
	urlpass := ""

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_Voluntariado (fechaNac,nacionalidad,nombreRepresentante,numIdentResponsable,parentescoId,nombreVoluntarioMenor,edadVoluntario,tipoDocIndentRespExtId,numIdentVoluntario,nombre,apellido,sexoId,numContacto,correo,tallaCamisaId,tiempoVoluntariadoId,paisNac,deptoId,municipioId,direccion,estudiaId,ultimoEstudioId,gradoEstudioId,universidad,carreraUniv,horasSocId,refPersonal,numRefPersonal,refPersonalb,numRefPersonalb,perfilVoluntariadoId,voluntarioComunitarioId,donante,areaDepartamento,fechaInduccion,nombreDesarrollador,areaId,actividadId,subActividadId,horarioActividad,deptoActividadId,sedeActividadId,otroClub,fechaRegistro,urlDuiFrontal,urlDuiDorso,paisId,docExtranjero) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(fechaNac, nacionalidad, nombreRepresentante, numIdentResponsable, parentescoId, nombreVoluntarioMenor, edadVoluntario, tipoDocIndentRespExtId, numIdentVoluntario, nombre, apellido, sexoId, numContacto, correo, tallaCamisaId, tiempoVoluntariadoId, paisNac, deptoId, municipioId, direccion, estudiaId, ultimoEstudioId, gradoEstudioId, universidad, carreraUniv, horasSocId, refPersonal, numRefPersonal, refPersonalb, numRefPersonalb, perfilVoluntariadoId, voluntarioComunitarioId, donante, areaDepartamento, fechaInduccion, nombreDesarrollador, areaId, actividadId, subActividadId, horarioActividad, deptoActividadId, sedeActividadId, otroClub, currentTimeGMT, urldocF, urldocA, paisId, urlpass)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO f_dias (voluntarioId, diasId) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dias", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta dias", http.StatusInternalServerError)
				return
			}
		}
	}

	discapacidad4 := r.Form["dis4"]
	if len(discapacidad4) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO f_horarios (voluntarioId, horarioId) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta horario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad4 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta horario", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
	handleError(err)
	// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
	client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
	handleError(err)

	file, fileHandler, err := r.FormFile("docF")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename
		fileExt := filepath.Ext(fileName)

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "col/duif/"+strconv.Itoa(int(idGenerado))+fileExt, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urldocF = strconv.Itoa(int(idGenerado)) + fileExt
	}
	file2, handler2, err := r.FormFile("docA")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		// Obtiene el nombre y extension del archivo cargado
		fileName2 := handler2.Filename
		fileExt2 := filepath.Ext(fileName2)

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "col/duia/"+strconv.Itoa(int(idGenerado))+fileExt2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		urldocA = strconv.Itoa(int(idGenerado)) + fileExt2
	}

	file3, fileHandler3, err := r.FormFile("pas")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file3.Close()

		// Obtiene el nombre del archivo cargado
		fileName3 := fileHandler3.Filename
		fileExt3 := filepath.Ext(fileName3)

		fileBytes3, err := io.ReadAll(file3)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile3, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile3.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile3.Write(fileBytes3); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "col/pasaporte/"+strconv.Itoa(int(idGenerado))+fileExt3, tempFile3,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urlpass = strconv.Itoa(int(idGenerado)) + fileExt3
	}

	tx2, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt4, err := tx2.Prepare("UPDATE F_Voluntariado SET urlDuiFrontal = ?, urlDuiDorso = ? ,docExtranjero= ? WHERE id = ?")
	if err != nil {
		tx2.Rollback()
		http.Error(w, "Error al preparar update", http.StatusInternalServerError)
		return
	}
	defer stmt4.Close()

	_, err = stmt4.Exec(urldocF, urldocA, urlpass, idGenerado)
	if err != nil {
		tx2.Rollback()
		http.Error(w, "Error al ejecutar ejecutar update", http.StatusInternalServerError)
		return
	}

	r.ParseForm()

	err = tx2.Commit()
	if err != nil {
		tx2.Rollback()
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

func postvolmx(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	fechaNac := r.FormValue("FechaN")
	nacionalidad := cambio(r.FormValue("gi"))
	nombreRepresentante := r.FormValue("nombreRep")
	numIdentResponsable := r.FormValue("DUIP")
	parentescoId := cambio(r.FormValue("parentesco"))
	nombreVoluntarioMenor := r.FormValue("nombreAdolescente")
	edadVoluntario := cambio(r.FormValue("edad"))
	tipoDocIndentRespExtId := cambio(r.FormValue("elec"))
	numIdentVoluntario := r.FormValue("DUI")
	nombre := r.FormValue("nombre")
	apellido := r.FormValue("apellido")
	sexoId := cambio(r.FormValue("sexo"))
	numContacto := r.FormValue("TM")
	correo := r.FormValue("correo")
	tallaCamisaId := cambio(r.FormValue("camisa"))
	tiempoVoluntariadoId := cambio(r.FormValue("voluntario"))
	paisNac := cambio(r.FormValue("paisn"))
	deptoId := cambio(r.FormValue("pais"))
	municipioId := cambio(r.FormValue("departamento"))
	direccion := r.FormValue("direccion")
	estudiaId := cambio(r.FormValue("TdC"))
	ultimoEstudioId := cambio(r.FormValue("UT"))
	gradoEstudioId := cambio(r.FormValue("GA"))
	universidad := r.FormValue("U")
	carreraUniv := r.FormValue("carrera")
	horasSocId := cambio(r.FormValue("HS"))
	refPersonal := r.FormValue("nombrer")
	numRefPersonal := r.FormValue("telr")
	refPersonalb := r.FormValue("nombrer2")
	numRefPersonalb := r.FormValue("telr2")
	perfilVoluntariadoId := cambio(r.FormValue("acc"))
	voluntarioComunitarioId := cambio(r.FormValue("acc2"))
	donante := cambio(r.FormValue("acc3"))
	areaDepartamento := r.FormValue("puesto")
	fechaInduccion := r.FormValue("FechaI")
	nombreDesarrollador := r.FormValue("nombrei")
	areaId := cambio(r.FormValue("pais2"))
	actividadId := cambio(r.FormValue("actividad"))
	subActividadId := cambio(r.FormValue("subact"))
	horarioActividad := r.FormValue("OHT")
	deptoActividadId := cambio(r.FormValue("departamento2"))
	sedeActividadId := cambio(r.FormValue("sede2"))
	paisId := 17
	otroClub := r.FormValue("aT")
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)
	urldocF := ""
	urldocA := ""
	urlpass := ""

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_Voluntariado (fechaNac,nacionalidad,nombreRepresentante,numIdentResponsable,parentescoId,nombreVoluntarioMenor,edadVoluntario,tipoDocIndentRespExtId,numIdentVoluntario,nombre,apellido,sexoId,numContacto,correo,tallaCamisaId,tiempoVoluntariadoId,paisNac,deptoId,municipioId,direccion,estudiaId,ultimoEstudioId,gradoEstudioId,universidad,carreraUniv,horasSocId,refPersonal,numRefPersonal,refPersonalb,numRefPersonalb,perfilVoluntariadoId,voluntarioComunitarioId,donante,areaDepartamento,fechaInduccion,nombreDesarrollador,areaId,actividadId,subActividadId,horarioActividad,deptoActividadId,sedeActividadId,otroClub,fechaRegistro,urlDuiFrontal,urlDuiDorso,paisId,docExtranjero) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(fechaNac, nacionalidad, nombreRepresentante, numIdentResponsable, parentescoId, nombreVoluntarioMenor, edadVoluntario, tipoDocIndentRespExtId, numIdentVoluntario, nombre, apellido, sexoId, numContacto, correo, tallaCamisaId, tiempoVoluntariadoId, paisNac, deptoId, municipioId, direccion, estudiaId, ultimoEstudioId, gradoEstudioId, universidad, carreraUniv, horasSocId, refPersonal, numRefPersonal, refPersonalb, numRefPersonalb, perfilVoluntariadoId, voluntarioComunitarioId, donante, areaDepartamento, fechaInduccion, nombreDesarrollador, areaId, actividadId, subActividadId, horarioActividad, deptoActividadId, sedeActividadId, otroClub, currentTimeGMT, urldocF, urldocA, paisId, urlpass)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO f_dias (voluntarioId, diasId) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dias", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta dias", http.StatusInternalServerError)
				return
			}
		}
	}

	discapacidad4 := r.Form["dis4"]
	if len(discapacidad4) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO f_horarios (voluntarioId, horarioId) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta horario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad4 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta horario", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
	handleError(err)
	// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
	client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
	handleError(err)

	file, fileHandler, err := r.FormFile("docF")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename
		fileExt := filepath.Ext(fileName)

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "mx/duif/"+strconv.Itoa(int(idGenerado))+fileExt, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urldocF = strconv.Itoa(int(idGenerado)) + fileExt
	}
	file2, handler2, err := r.FormFile("docA")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename
		fileExt2 := filepath.Ext(fileName2)

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "mx/duia/"+strconv.Itoa(int(idGenerado))+fileExt2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		urldocA = strconv.Itoa(int(idGenerado)) + fileExt2
	}

	file3, fileHandler3, err := r.FormFile("pas")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file3.Close()

		// Obtiene el nombre del archivo cargado
		fileName3 := fileHandler3.Filename
		fileExt3 := filepath.Ext(fileName3)

		fileBytes3, err := io.ReadAll(file3)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile3, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile3.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile3.Write(fileBytes3); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "mx/pasaporte/"+strconv.Itoa(int(idGenerado))+fileExt3, tempFile3,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urlpass = strconv.Itoa(int(idGenerado)) + fileExt3
	}

	tx2, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt4, err := tx2.Prepare("UPDATE F_Voluntariado SET urlDuiFrontal = ?, urlDuiDorso = ? ,docExtranjero= ? WHERE id = ?")
	if err != nil {
		tx2.Rollback()
		http.Error(w, "Error al preparar update", http.StatusInternalServerError)
		return
	}
	defer stmt4.Close()

	_, err = stmt4.Exec(urldocF, urldocA, urlpass, idGenerado)
	if err != nil {
		tx2.Rollback()
		http.Error(w, "Error al ejecutar ejecutar update", http.StatusInternalServerError)
		return
	}

	r.ParseForm()

	err = tx2.Commit()
	if err != nil {
		tx2.Rollback()
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

func postvolcr(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	fechaNac := r.FormValue("FechaN")
	nacionalidad := cambio(r.FormValue("gi"))
	nombreRepresentante := r.FormValue("nombreRep")
	numIdentResponsable := r.FormValue("DUIP")
	parentescoId := cambio(r.FormValue("parentesco"))
	nombreVoluntarioMenor := r.FormValue("nombreAdolescente")
	edadVoluntario := cambio(r.FormValue("edad"))
	tipoDocIndentRespExtId := cambio(r.FormValue("elec"))
	numIdentVoluntario := r.FormValue("DUI")
	nombre := r.FormValue("nombre")
	apellido := r.FormValue("apellido")
	sexoId := cambio(r.FormValue("sexo"))
	numContacto := r.FormValue("TM")
	correo := r.FormValue("correo")
	tallaCamisaId := cambio(r.FormValue("camisa"))
	tiempoVoluntariadoId := cambio(r.FormValue("voluntario"))
	paisNac := cambio(r.FormValue("paisn"))
	deptoId := cambio(r.FormValue("pais"))
	municipioId := cambio(r.FormValue("departamento"))
	direccion := r.FormValue("direccion")
	estudiaId := cambio(r.FormValue("TdC"))
	ultimoEstudioId := cambio(r.FormValue("UT"))
	gradoEstudioId := cambio(r.FormValue("GA"))
	universidad := r.FormValue("U")
	carreraUniv := r.FormValue("carrera")
	horasSocId := cambio(r.FormValue("HS"))
	refPersonal := r.FormValue("nombrer")
	numRefPersonal := r.FormValue("telr")
	refPersonalb := r.FormValue("nombrer2")
	numRefPersonalb := r.FormValue("telr2")
	perfilVoluntariadoId := cambio(r.FormValue("acc"))
	voluntarioComunitarioId := cambio(r.FormValue("acc2"))
	donante := cambio(r.FormValue("acc3"))
	areaDepartamento := r.FormValue("puesto")
	fechaInduccion := r.FormValue("FechaI")
	nombreDesarrollador := r.FormValue("nombrei")
	areaId := cambio(r.FormValue("pais2"))
	actividadId := cambio(r.FormValue("actividad"))
	subActividadId := cambio(r.FormValue("subact"))
	horarioActividad := r.FormValue("OHT")
	deptoActividadId := cambio(r.FormValue("departamento2"))
	sedeActividadId := cambio(r.FormValue("sede2"))
	paisId := 6
	otroClub := r.FormValue("aT")
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)
	urldocF := ""
	urldocA := ""
	urlpass := ""

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_Voluntariado (fechaNac,nacionalidad,nombreRepresentante,numIdentResponsable,parentescoId,nombreVoluntarioMenor,edadVoluntario,tipoDocIndentRespExtId,numIdentVoluntario,nombre,apellido,sexoId,numContacto,correo,tallaCamisaId,tiempoVoluntariadoId,paisNac,deptoId,municipioId,direccion,estudiaId,ultimoEstudioId,gradoEstudioId,universidad,carreraUniv,horasSocId,refPersonal,numRefPersonal,refPersonalb,numRefPersonalb,perfilVoluntariadoId,voluntarioComunitarioId,donante,areaDepartamento,fechaInduccion,nombreDesarrollador,areaId,actividadId,subActividadId,horarioActividad,deptoActividadId,sedeActividadId,otroClub,fechaRegistro,urlDuiFrontal,urlDuiDorso,paisId,docExtranjero) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(fechaNac, nacionalidad, nombreRepresentante, numIdentResponsable, parentescoId, nombreVoluntarioMenor, edadVoluntario, tipoDocIndentRespExtId, numIdentVoluntario, nombre, apellido, sexoId, numContacto, correo, tallaCamisaId, tiempoVoluntariadoId, paisNac, deptoId, municipioId, direccion, estudiaId, ultimoEstudioId, gradoEstudioId, universidad, carreraUniv, horasSocId, refPersonal, numRefPersonal, refPersonalb, numRefPersonalb, perfilVoluntariadoId, voluntarioComunitarioId, donante, areaDepartamento, fechaInduccion, nombreDesarrollador, areaId, actividadId, subActividadId, horarioActividad, deptoActividadId, sedeActividadId, otroClub, currentTimeGMT, urldocF, urldocA, paisId, urlpass)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO f_dias (voluntarioId, diasId) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dias", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta dias", http.StatusInternalServerError)
				return
			}
		}
	}

	discapacidad4 := r.Form["dis4"]
	if len(discapacidad4) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO f_horarios (voluntarioId, horarioId) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta horario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad4 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta horario", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
	handleError(err)
	// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
	client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
	handleError(err)

	file, fileHandler, err := r.FormFile("docF")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename
		fileExt := filepath.Ext(fileName)

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "cr/duif/"+strconv.Itoa(int(idGenerado))+fileExt, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urldocF = strconv.Itoa(int(idGenerado)) + fileExt
	}
	file2, handler2, err := r.FormFile("docA")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename
		fileExt2 := filepath.Ext(fileName2)

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "cr/duia/"+strconv.Itoa(int(idGenerado))+fileExt2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		urldocA = strconv.Itoa(int(idGenerado)) + fileExt2
	}

	file3, fileHandler3, err := r.FormFile("pas")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file3.Close()

		// Obtiene el nombre del archivo cargado
		fileName3 := fileHandler3.Filename
		fileExt3 := filepath.Ext(fileName3)

		fileBytes3, err := io.ReadAll(file3)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile3, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile3.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile3.Write(fileBytes3); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "cr/pasaporte/"+strconv.Itoa(int(idGenerado))+fileExt3, tempFile3,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urlpass = strconv.Itoa(int(idGenerado)) + fileExt3
	}

	tx2, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt4, err := tx2.Prepare("UPDATE F_Voluntariado SET urlDuiFrontal = ?, urlDuiDorso = ? ,docExtranjero= ? WHERE id = ?")
	if err != nil {
		tx2.Rollback()
		http.Error(w, "Error al preparar update", http.StatusInternalServerError)
		return
	}
	defer stmt4.Close()

	_, err = stmt4.Exec(urldocF, urldocA, urlpass, idGenerado)
	if err != nil {
		tx2.Rollback()
		http.Error(w, "Error al ejecutar ejecutar update", http.StatusInternalServerError)
		return
	}

	r.ParseForm()

	err = tx2.Commit()
	if err != nil {
		tx2.Rollback()
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

func postvolhn(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	//Validacion cice
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	fechaNac := r.FormValue("FechaN")
	nacionalidad := cambio(r.FormValue("gi"))
	nombreRepresentante := r.FormValue("nombreRep")
	numIdentResponsable := r.FormValue("DUIP")
	parentescoId := cambio(r.FormValue("parentesco"))
	nombreVoluntarioMenor := r.FormValue("nombreAdolescente")
	edadVoluntario := cambio(r.FormValue("edad"))
	tipoDocIndentRespExtId := cambio(r.FormValue("elec"))
	numIdentVoluntario := r.FormValue("DUI")
	nombre := r.FormValue("nombre")
	apellido := r.FormValue("apellido")
	sexoId := cambio(r.FormValue("sexo"))
	numContacto := r.FormValue("TM")
	correo := r.FormValue("correo")
	tallaCamisaId := cambio(r.FormValue("camisa"))
	tiempoVoluntariadoId := cambio(r.FormValue("voluntario"))
	paisNac := cambio(r.FormValue("paisn"))
	deptoId := cambio(r.FormValue("pais"))
	municipioId := cambio(r.FormValue("departamento"))
	direccion := r.FormValue("direccion")
	estudiaId := cambio(r.FormValue("TdC"))
	ultimoEstudioId := cambio(r.FormValue("UT"))
	gradoEstudioId := cambio(r.FormValue("GA"))
	universidad := r.FormValue("U")
	carreraUniv := r.FormValue("carrera")
	horasSocId := cambio(r.FormValue("HS"))
	refPersonal := r.FormValue("nombrer")
	numRefPersonal := r.FormValue("telr")
	refPersonalb := r.FormValue("nombrer2")
	numRefPersonalb := r.FormValue("telr2")
	perfilVoluntariadoId := cambio(r.FormValue("acc"))
	voluntarioComunitarioId := cambio(r.FormValue("acc2"))
	institucionId := cambio(r.FormValue("accinst"))
	donante := cambio(r.FormValue("acc3"))
	areaDepartamento := r.FormValue("puesto")
	fechaInduccion := r.FormValue("FechaI")
	nombreDesarrollador := r.FormValue("nombrei")
	areaId := cambio(r.FormValue("pais2"))
	actividadId := cambio(r.FormValue("actividad"))
	subActividadId := cambio(r.FormValue("subact"))
	horarioActividad := r.FormValue("OHT")
	deptoActividadId := cambio(r.FormValue("departamento2"))
	sedeActividadId := cambio(r.FormValue("sede2"))
	paisId := 3
	otroClub := r.FormValue("aT")
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)
	urldocF := ""
	urldocA := ""
	urlpass := ""

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_Voluntariado (fechaNac,nacionalidad,nombreRepresentante,numIdentResponsable,parentescoId,nombreVoluntarioMenor,edadVoluntario,tipoDocIndentRespExtId,numIdentVoluntario,nombre,apellido,sexoId,numContacto,correo,tallaCamisaId,tiempoVoluntariadoId,paisNac,deptoId,municipioId,direccion,estudiaId,ultimoEstudioId,gradoEstudioId,universidad,carreraUniv,horasSocId,refPersonal,numRefPersonal,refPersonalb,numRefPersonalb,perfilVoluntariadoId,voluntarioComunitarioId,institucionId, donante,areaDepartamento,fechaInduccion,nombreDesarrollador,areaId,actividadId,subActividadId,horarioActividad,deptoActividadId,sedeActividadId,otroClub,fechaRegistro,urlDuiFrontal,urlDuiDorso,paisId,docExtranjero) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(fechaNac, nacionalidad, nombreRepresentante, numIdentResponsable, parentescoId, nombreVoluntarioMenor, edadVoluntario, tipoDocIndentRespExtId, numIdentVoluntario, nombre, apellido, sexoId, numContacto, correo, tallaCamisaId, tiempoVoluntariadoId, paisNac, deptoId, municipioId, direccion, estudiaId, ultimoEstudioId, gradoEstudioId, universidad, carreraUniv, horasSocId, refPersonal, numRefPersonal, refPersonalb, numRefPersonalb, perfilVoluntariadoId, voluntarioComunitarioId, institucionId, donante, areaDepartamento, fechaInduccion, nombreDesarrollador, areaId, actividadId, subActividadId, horarioActividad, deptoActividadId, sedeActividadId, otroClub, currentTimeGMT, urldocF, urldocA, paisId, urlpass)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO f_dias (voluntarioId, diasId) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dias", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta dias", http.StatusInternalServerError)
				return
			}
		}
	}

	discapacidad4 := r.Form["dis4"]
	if len(discapacidad4) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO f_horarios (voluntarioId, horarioId) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta horario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad4 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta horario", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
	handleError(err)
	// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
	client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
	handleError(err)

	file, fileHandler, err := r.FormFile("docF")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre y extension del archivo cargado
		fileName := fileHandler.Filename
		fileExt := filepath.Ext(fileName)

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "hn/duif/"+strconv.Itoa(int(idGenerado))+fileExt, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urldocF = strconv.Itoa(int(idGenerado)) + fileExt
	}
	file2, handler2, err := r.FormFile("docA")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre y extension del archivo cargado
		fileName2 := handler2.Filename
		fileExt2 := filepath.Ext(fileName2)

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "hn/duia/"+strconv.Itoa(int(idGenerado))+fileExt2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		urldocA = strconv.Itoa(int(idGenerado)) + fileExt2
	}

	file3, fileHandler3, err := r.FormFile("pas")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file3.Close()

		// Obtiene el nombre y extension del archivo cargado
		fileName3 := fileHandler3.Filename
		fileExt3 := filepath.Ext(fileName3)

		fileBytes3, err := io.ReadAll(file3)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile3, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile3.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile3.Write(fileBytes3); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "hn/pasaporte/"+strconv.Itoa(int(idGenerado))+fileExt3, tempFile3,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urlpass = strconv.Itoa(int(idGenerado)) + fileExt3

	}

	tx2, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt4, err := tx2.Prepare("UPDATE F_Voluntariado SET urlDuiFrontal = ?, urlDuiDorso = ? ,docExtranjero= ? WHERE id = ?")
	if err != nil {
		tx2.Rollback()
		http.Error(w, "Error al preparar update", http.StatusInternalServerError)
		return
	}
	defer stmt4.Close()

	_, err = stmt4.Exec(urldocF, urldocA, urlpass, idGenerado)
	if err != nil {
		tx2.Rollback()
		http.Error(w, "Error al ejecutar ejecutar update", http.StatusInternalServerError)
		return
	}

	r.ParseForm()

	err = tx2.Commit()
	if err != nil {
		tx2.Rollback()
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

func postvolsv(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()
	// Obtener los valores de los campos del formulario
	fechaNac := r.FormValue("FechaN")
	nacionalidad := cambio(r.FormValue("gi"))
	nombreRepresentante := r.FormValue("nombreRep")
	numIdentResponsable := r.FormValue("DUIP")
	parentescoId := cambio(r.FormValue("parentesco"))
	nombreVoluntarioMenor := r.FormValue("nombreAdolescente")
	edadVoluntario := cambio(r.FormValue("edad"))
	tipoDocIndentRespExtId := cambio(r.FormValue("elec"))
	numIdentVoluntario := r.FormValue("DUI")
	nombre := r.FormValue("nombre")
	apellido := r.FormValue("apellido")
	sexoId := cambio(r.FormValue("sexo"))
	numContacto := r.FormValue("TM")
	correo := r.FormValue("correo")
	tallaCamisaId := cambio(r.FormValue("camisa"))
	tiempoVoluntariadoId := cambio(r.FormValue("voluntario"))
	paisNac := cambio(r.FormValue("paisn"))
	deptoId := cambio(r.FormValue("pais"))
	municipioId := cambio(r.FormValue("departamento"))
	direccion := r.FormValue("direccion")
	estudiaId := cambio(r.FormValue("TdC"))
	ultimoEstudioId := cambio(r.FormValue("UT"))
	gradoEstudioId := cambio(r.FormValue("GA"))
	universidad := r.FormValue("U")
	carreraUniv := r.FormValue("carrera")
	horasSocId := cambio(r.FormValue("HS"))
	refPersonal := r.FormValue("nombrer")
	numRefPersonal := r.FormValue("telr")
	refPersonalb := r.FormValue("nombrer2")
	numRefPersonalb := r.FormValue("telr2")
	perfilVoluntariadoId := cambio(r.FormValue("acc"))
	voluntarioComunitarioId := cambio(r.FormValue("acc2"))
	institucionId := cambio(r.FormValue("accinst"))
	donante := cambio(r.FormValue("acc3"))
	areaDepartamento := r.FormValue("puesto")
	fechaInduccion := r.FormValue("FechaI")
	nombreDesarrollador := r.FormValue("nombrei")
	areaId := cambio(r.FormValue("pais2"))
	actividadId := cambio(r.FormValue("actividad"))
	subActividadId := cambio(r.FormValue("subact"))
	horarioActividad := r.FormValue("OHT")
	deptoActividadId := cambio(r.FormValue("departamento2"))
	sedeActividadId := cambio(r.FormValue("sede2"))
	paisId := 7
	otroClub := r.FormValue("aT")
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)
	urldocF := ""
	urldocA := ""
	urlpass := ""

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_Voluntariado (fechaNac,nacionalidad,nombreRepresentante,numIdentResponsable,parentescoId,nombreVoluntarioMenor,edadVoluntario,tipoDocIndentRespExtId,numIdentVoluntario,nombre,apellido,sexoId,numContacto,correo,tallaCamisaId,tiempoVoluntariadoId,paisNac,deptoId,municipioId,direccion,estudiaId,ultimoEstudioId,gradoEstudioId,universidad,carreraUniv,horasSocId,refPersonal,numRefPersonal,refPersonalb,numRefPersonalb,perfilVoluntariadoId,voluntarioComunitarioId,institucionId,donante,areaDepartamento,fechaInduccion,nombreDesarrollador,areaId,actividadId,subActividadId,horarioActividad,deptoActividadId,sedeActividadId,otroClub,fechaRegistro,urlDuiFrontal,urlDuiDorso,paisId,docExtranjero) OUTPUT INSERTED.ID VALUES (@p1,@p2,@p3,@p4,@p5,@p6,@p7,@p8,@p9,@p10,@p11,@p12,@p13,@p14,@p15,@p16,@p17,@p18,@p19,@p20,@p21,@p22,@p23,@p24,@p25,@p26,@p27,@p28,@p29,@p30,@p31,@p32,@p33,@p34,@p35,@p36,@p37,@p38,@p39,@p40,@p41,@p42,@p43,@p44,@p45,@p46,@p47,@p48,@p49)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()
	var idGenerado int64

	err = stmt.QueryRow(fechaNac, nacionalidad, nombreRepresentante, numIdentResponsable, parentescoId, nombreVoluntarioMenor, edadVoluntario, tipoDocIndentRespExtId, numIdentVoluntario, nombre, apellido, sexoId, numContacto, correo, tallaCamisaId, tiempoVoluntariadoId, paisNac, deptoId, municipioId, direccion, estudiaId, ultimoEstudioId, gradoEstudioId, universidad, carreraUniv, horasSocId, refPersonal, numRefPersonal, refPersonalb, numRefPersonalb, perfilVoluntariadoId, voluntarioComunitarioId, institucionId, donante, areaDepartamento, fechaInduccion, nombreDesarrollador, areaId, actividadId, subActividadId, horarioActividad, deptoActividadId, sedeActividadId, otroClub, currentTimeGMT, urldocF, urldocA, paisId, urlpass).Scan(&idGenerado)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO f_dias (voluntarioId, diasId) VALUES (@p1, @p2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, fmt.Sprintf("Error al preparar la consulta dias: %v", err), http.StatusInternalServerError)
			return
		}

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, fmt.Sprintf("Error al ejecutar la consulta dias: %v", err), http.StatusInternalServerError)
				return
			}
		}
	}

	discapacidad4 := r.Form["dis4"]
	if len(discapacidad4) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO f_horarios (voluntarioId, horarioId) VALUES (@p1, @p2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, fmt.Sprintf("Error al preparar la consulta horario: %v", err), http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad4 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, fmt.Sprintf("Error al preparar la consulta horario: %v", err), http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
	handleError(err)
	// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
	client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
	handleError(err)

	file, fileHandler, err := r.FormFile("docF")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename
		fileExt := filepath.Ext(fileName)
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "sv/duif/"+strconv.Itoa(int(idGenerado))+fileExt, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urldocF = strconv.Itoa(int(idGenerado)) + fileExt
	}
	file2, handler2, err := r.FormFile("docA")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename
		fileExt2 := filepath.Ext(fileName2)

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "sv/duia/"+strconv.Itoa(int(idGenerado))+fileExt2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		urldocA = strconv.Itoa(int(idGenerado)) + fileExt2
	}

	file3, fileHandler3, err := r.FormFile("pas")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file3.Close()

		// Obtiene el nombre del archivo cargado
		fileName3 := fileHandler3.Filename
		fileExt3 := filepath.Ext(fileName3)

		fileBytes3, err := io.ReadAll(file3)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile3, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile3.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile3.Write(fileBytes3); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "voluntariado", "sv/pasaporte/"+strconv.Itoa(int(idGenerado))+fileExt3, tempFile3,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urlpass = strconv.Itoa(int(idGenerado)) + fileExt3
	}

	tx2, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt4, err := tx2.Prepare("UPDATE F_Voluntariado SET urlDuiFrontal = @p1, urlDuiDorso = @p2 ,docExtranjero= @p3 WHERE id = @p4")
	if err != nil {
		tx2.Rollback()
		http.Error(w, "Error al preparar update", http.StatusInternalServerError)
		return
	}
	defer stmt4.Close()

	_, err = stmt4.Exec(urldocF, urldocA, urlpass, idGenerado)
	if err != nil {
		tx2.Rollback()
		http.Error(w, "Error al ejecutar ejecutar update", http.StatusInternalServerError)
		return
	}

	err = tx2.Commit()
	if err != nil {
		tx2.Rollback()
	}

	data := struct {
		ID     int
		Nombre string
	}{
		ID:     int(idGenerado),
		Nombre: nombre + " " + apellido,
	}
	t, err := template.ParseFiles("public/HTML/Respuestavol.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar la plantilla con los datos
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func moduloasvHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=7 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=7  and activo=1 order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/sv/ModuloA.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postJLI(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	// Obtener los valores de los campos del formulario
	Formulario := 6
	FechaNac := r.FormValue("FechaN")
	pais := cambio(r.FormValue("pais"))
	psocio := r.FormValue("persona")
	prol := cambio(r.FormValue("rol"))
	sede := cambio(r.FormValue("sede2"))
	nombre := r.FormValue("nombreCompleto")
	apellido := r.FormValue("apellido")
	estadoCivil := cambio(r.FormValue("EstadoC"))
	nacionalidad := cambio(r.FormValue("nacionalidad"))
	departamento := cambio(r.FormValue("departamento"))
	municipio := cambio(r.FormValue("municipio"))
	comunidad := r.FormValue("comunidad")
	sector := r.FormValue("sector")
	calle := r.FormValue("calle")
	bloque := r.FormValue("bloque")
	apartamento := r.FormValue("apartamento")
	casa := r.FormValue("casa")
	preferencia := r.FormValue("preferencia")
	direccion := r.FormValue("direccion")
	sexo := cambio(r.FormValue("Sexo"))
	discapacidad := cambio(r.FormValue("dis"))
	Estudia := cambio(r.FormValue("TdC"))
	ultimoestudio := cambio(r.FormValue("UT"))
	establecimiento := cambio(r.FormValue("privado"))
	grado := cambio(r.FormValue("GA"))
	turno := cambio(r.FormValue("Turno"))
	Nombrerepresentante := r.FormValue("responsable")
	telefonoresp := r.FormValue("contacto2")
	ccomunidad := cambio(r.FormValue("ccomunidad"))
	migrante := cambio(r.FormValue("migrante"))
	fmigrante := r.FormValue("fmigrante")
	empleo := cambio(r.FormValue("empleo"))
	gastos := cambio(r.FormValue("gastos"))
	nuevogw := cambio(r.FormValue("gp"))
	rolp := cambio(r.FormValue("pv"))
	actividadparticipo := cambio(r.FormValue("ActG"))
	DUI := r.FormValue("DUIM")
	correo := r.FormValue("correo")
	telefono := r.FormValue("TM")
	depNac := cambio(r.FormValue("pais2"))
	munNac := cambio(r.FormValue("departamento2"))
	beneficiario := r.FormValue("beneficiario")
	beneficiarioapellido := r.FormValue("beneficiarioapellido")
	parentesco := cambio(r.FormValue("relacion"))
	urldui := ""
	urlduia := ""
	urlconstancia := ""
	urlinscripcion := ""
	autoResponsable1 := cambio(r.FormValue("pbrindaa"))
	autoResponsable2 := cambio(r.FormValue("pbrindab"))
	autoResponsable3 := cambio(r.FormValue("pbrindac"))
	autoResponsable4 := cambio(r.FormValue("pbrindad"))
	autoResponsable5 := cambio(r.FormValue("pbrindae"))
	autoParticipante1 := cambio(r.FormValue("haceptaa"))
	autoParticipante2 := cambio(r.FormValue("haceptab"))
	autoParticipante3 := cambio(r.FormValue("haceptac"))
	autoParticipante4 := cambio(r.FormValue("haceptad"))
	autoParticipante5 := cambio(r.FormValue("haceptae"))
	cohorte := cambio(r.FormValue("cohorte"))
	perfilparticipante := cambio(r.FormValue("perfilparticipante"))

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_participantes (FechaNac,pais,psocio,prol,sede,nombre,apellido,estadoCivil,nacionalidad,departamento,municipio,comunidad,sector,calle,bloque,apartamento,casa,preferencia,direccion,sexo,discapacidad,Estudia,ultimoestudio,establecimiento,grado,turno,Nombrerepresentante,telefonoresp,ccomunidad,migrante,fmigrante,empleo,gastos,nuevogw,rolp,actividadparticipo,DUI,correo,telefono,depNac,munNac,beneficiario,parentesco,urldui,urlduia,urlconstancia,urlinscripcion,fechaRegistro,Formulario,autoResponsable1,autoResponsable2,autoResponsable3,autoResponsable4,autoResponsable5,autoParticipante1,autoParticipante2,autoParticipante3,autoParticipante4,autoParticipante5,beneficiarioapellido,cohorte,perfilparticipante) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?, ?, ?, ?, ?,?, ?, ?, ?, ?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,? )")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(FechaNac, pais, psocio, prol, sede, nombre, apellido, estadoCivil, nacionalidad, departamento, municipio, comunidad, sector, calle, bloque, apartamento, casa, preferencia, direccion, sexo, discapacidad, Estudia, ultimoestudio, establecimiento, grado, turno, Nombrerepresentante, telefonoresp, ccomunidad, migrante, fmigrante, empleo, gastos, nuevogw, rolp, actividadparticipo, DUI, correo, telefono, depNac, munNac, beneficiario, parentesco, urldui, urlduia, urlconstancia, urlinscripcion, currentTimeGMT, Formulario, autoResponsable1, autoResponsable2, autoResponsable3, autoResponsable4, autoResponsable5, autoParticipante1, autoParticipante2, autoParticipante3, autoParticipante4, autoParticipante5, beneficiarioapellido, cohorte, perfilparticipante)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_discapacidad (idp, discapacidad) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dis", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
		}
	}

	situacion := r.Form["situacion"]
	if len(situacion) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_situacion (idp, situacion) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta situacion", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range situacion {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta situacion", http.StatusInternalServerError)
				return
			}
		}
	}

	dependencia := r.Form["dependencia"]
	if len(dependencia) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_dependecia (idp, dependencia) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dependencia", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range dependencia {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta dependencia", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
	handleError(err)
	// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
	client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
	handleError(err)
	file, fileHandler, err := r.FormFile("docF")
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
		handleError(err)
	}
	defer tempFile.Close()

	// Escribe los bytes del archivo en el archivo temporal
	if _, err := tempFile.Write(fileBytes); err != nil {
		handleError(err)
	}

	// Upload the file to a block blob
	_, err = client.UploadFile(context.TODO(), "jli", "identidadF/"+strconv.Itoa(cohorte)+"/"+nombre+" "+apellido+"_"+DUI+"_"+strconv.Itoa(int(idGenerado))+fileExt, tempFile,
		&azblob.UploadFileOptions{
			BlockSize:   int64(1024),
			Concurrency: uint16(3),
			// If Progress is non-nil, this function is called periodically as bytes are uploaded.
			Progress: func(bytesTransferred int64) {

			},
		})
	handleError(err)

	//Asignando la URL del blob a una variable
	urldui = nombre + " " + apellido + "_" + DUI + "_" + strconv.Itoa(int(idGenerado)) + fileExt

	file2, handler2, err := r.FormFile("docA")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Obtiene el nombre del archivo cargado
		fileName2 := handler2.Filename

		//Obtiene la extensión del archivo cargado
		fileExt2 := filepath.Ext(fileName2)

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "jli", "identidadA/"+strconv.Itoa(cohorte)+"/"+nombre+" "+apellido+"_"+DUI+"_"+strconv.Itoa(int(idGenerado))+fileExt2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})
		handleError(err)
		//Asignando la URL del blob a una variable
		urlduia = nombre + " " + apellido + "_" + DUI + "_" + strconv.Itoa(int(idGenerado)) + fileExt2
	}
	tx2, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt4, err := tx2.Prepare("UPDATE F_participantes SET urldui = ?, urlduia = ? WHERE id = ?")
	if err != nil {
		tx2.Rollback()
		http.Error(w, "Error al preparar update", http.StatusInternalServerError)
		return
	}
	defer stmt4.Close()

	_, err = stmt4.Exec(urldui, urlduia, idGenerado)
	if err != nil {
		tx2.Rollback()
		http.Error(w, "Error al ejecutar ejecutar update", http.StatusInternalServerError)
		return
	}

	r.ParseForm()

	err = tx2.Commit()
	if err != nil {
		tx2.Rollback()
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

func obtenermunsedeHandlersanamente(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	depsede := r.FormValue("depsede")
	var opcionesmunsede string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT Distinct Idmun,Municipio FROM VSedessanamente where Iddep=? order by Municipio", depsede)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionesmunsede += "<option value=''>Seleccionar un municipio</option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesmunsede += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionesmunsede)
}

func obtenersedeHandlersanamente(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	sede := r.FormValue("munsede")
	var opcionessede string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT IdSede,Sede FROM VSedessanamente where Idmun=? order by Sede", sede)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionessede += "<option value=''>Seleccionar una sede</option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionessede += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionessede)
}

func postmodulo(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	// Obtener los valores de los campos del formulario

	nombre := r.FormValue("nombreCompleto")
	Formador := cambio(r.FormValue("apellido"))
	nacionalidad := r.FormValue("nacionalidad")
	pais := r.FormValue("idp")
	sede := r.FormValue("sede2")
	dui := r.FormValue("DUIM")
	p1 := cambio(r.FormValue("p1"))
	p2 := cambio(r.FormValue("p2"))
	p3 := cambio(r.FormValue("p3"))
	p4 := cambio(r.FormValue("p4"))
	p5 := cambio(r.FormValue("p5"))
	modulo := r.FormValue("modulo")
	if modulo == "0" || modulo == "" {
		t, err := template.ParseFiles("Public/Html/RespuestaSanamenteError.html")
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

	} else {
		type Datos struct {
			Nombre   string
			Id       int
			Programa string
		}

		tx, err := db.Begin()
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error conexion", http.StatusInternalServerError)

		}

		stmt, err := tx.Prepare("INSERT INTO SanamenteModulos (fechaRegistro, nombre, formador, nacionalidad, pais, sede, dui, p1,p2,p3,p4,p5,modulo) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		var lastInsertId int
		err = stmt.QueryRow(currentTimeGMT, nombre, Formador, nacionalidad, pais, sede, dui, p1, p2, p3, p4, p5, modulo).Scan(&lastInsertId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		err = tx.Commit()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var programa string
		switch modulo {

		case "1":
			programa = "A"

		case "2":
			programa = "B"

		case "3":
			programa = "C"
		}

		datos := Datos{
			Nombre:   nombre,
			Id:       lastInsertId,
			Programa: programa,
		}

		t, err := template.ParseFiles("Public/Html/RespuestaSanamentePersonalizadaquiz.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Renderizar la plantilla con los datos
		err = t.Execute(w, datos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func postmodulocol(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	// Obtener los valores de los campos del formulario

	nombre := r.FormValue("nombreCompleto")
	Formador := cambio(r.FormValue("apellido"))
	nacionalidad := r.FormValue("nacionalidad")
	pais := r.FormValue("idp")
	sede := r.FormValue("sede2")
	dui := r.FormValue("DUIM")
	p1 := cambio(r.FormValue("p1"))
	p2 := cambio(r.FormValue("p2"))
	p3 := cambio(r.FormValue("p3"))
	p4 := cambio(r.FormValue("p4"))
	p5 := cambio(r.FormValue("p5"))
	p6 := cambio(r.FormValue("p6"))
	p7 := cambio(r.FormValue("p7"))
	p8 := cambio(r.FormValue("p8"))

	modulo := r.FormValue("modulo")

	type Datos struct {
		Nombre   string
		Id       int
		Programa string
	}

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error conexion", http.StatusInternalServerError)
		return
	}

	stmt, err := tx.Prepare("INSERT INTO SanamenteModulos (fechaRegistro, nombre, formador, nacionalidad, pais, sede, dui, p1,p2,p3,p4,p5,p6,p7,p8,modulo) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14,@P15,@P16)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)

	}

	var lastInsertId int
	err = stmt.QueryRow(currentTimeGMT, nombre, Formador, nacionalidad, pais, sede, dui, p1, p2, p3, p4, p5, p6, p7, p8, modulo).Scan(&lastInsertId)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	datos := Datos{
		Nombre: nombre,
		Id:     lastInsertId,
	}

	t, err := template.ParseFiles("Public/Html/RespuestaSanamentePersonalizadaquizUnico.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Renderizar la plantilla con los datos
	err = t.Execute(w, datos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postmoduloInterv(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	// Obtener los valores de los campos del formulario

	nombre := r.FormValue("nombreCompleto")
	Formador := cambio(r.FormValue("apellido"))
	nacionalidad := r.FormValue("nacionalidad")
	pais := r.FormValue("idp")
	sede := r.FormValue("sede2")
	dui := r.FormValue("DUIM")
	p1 := cambio(r.FormValue("p1"))
	p2 := cambio(r.FormValue("p2"))
	p3 := cambio(r.FormValue("p3"))
	p4 := cambio(r.FormValue("p4"))
	p5 := cambio(r.FormValue("p5"))
	p6 := cambio(r.FormValue("p6"))
	/*	p7 := cambio(r.FormValue("p7"))
	 	p8 := cambio(r.FormValue("p8"))
		p9 := cambio(r.FormValue("p9")) */
	version := cambio(r.FormValue("version"))
	fechaformacion := r.FormValue("fechaform")

	modulo := r.FormValue("modulo")

	if p1 == 0 {
		t, err := template.ParseFiles("Public/Html/RespuestaSanamenteError.html")
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

	} else {

		type Datos struct {
			Nombre   string
			Id       int
			Programa string
		}

		tx, err := db.Begin()
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error conexion", http.StatusInternalServerError)

		}

		stmt, err := tx.Prepare("INSERT INTO [dbo].[SanamenteModulosIntervencion] (created_at, nombre, formador, nacionalidad, pais, sede, dui, p1,p2,p3,p4,p5,p6,modulo,Version,fechaformacion) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14,@P15,@P16)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		var lastInsertId int
		err = stmt.QueryRow(currentTimeGMT, nombre, Formador, nacionalidad, pais, sede, dui, p1, p2, p3, p4, p5, p6, modulo, version, fechaformacion).Scan(&lastInsertId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		err = tx.Commit()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		datos := Datos{
			Nombre: nombre,
			Id:     lastInsertId,
		}

		t, err := template.ParseFiles("Public/Html/RespuestaSanamentePersonalizadaquizUnico.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Renderizar la plantilla con los datos
		err = t.Execute(w, datos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func postcasos(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	// Obtener los valores de los campos del formulario

	nombre := r.FormValue("p23")
	tipo := cambio(r.FormValue("p30"))
	edad := cambio(r.FormValue("p24"))
	nombre_res := r.FormValue("p25")
	telefono_res := r.FormValue("p26")
	telefono := r.FormValue("p27")
	departamento := cambio(r.FormValue("p28"))
	municipio := cambio(r.FormValue("p29"))
	motivo := r.FormValue("p31")
	formador := cambio(r.FormValue("p138"))

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO ReferenciaCasos (nombre, tipo, edad, telefono,nombre_res, telefono_res,departamento, municipio, motivo,formador) VALUES (@p1, @p2, @p3, @p4, @p5, @p6,@p7, @p8,@p9,@p10)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(nombre, tipo, edad, telefono, nombre_res, telefono_res, departamento, municipio, motivo, formador)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v", err)
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}
	r.ParseForm()

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	t, err := template.ParseFiles("public/HTML/Respuestacasos.html")
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

func modulobsvHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=7 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=7  and activo=1  order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/sv/ModuloB.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func modulocsvHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=7 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=7  and activo=1  order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/sv/ModuloC.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func prueba(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.sede_areas where Ida=5 and IdPais=7 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where typePerson=1 and active_at is not null and id not in(10)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Educacion/formprueba.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func formulariomentoriasmenores(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	transporte, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_transporte where activo=1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estipendio, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_estipendio")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estado_civ, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_estadociv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actividad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_ofertavol")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.Vista_Sedes where Idpais=7 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grupo, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_grupoSocial")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pueblo, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_puebloComunidad")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	idioma, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_idioma")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actLaboral, err := Controller.GetOpcionesPregunta(db, "SELECT Id, name FROM F_actividadLaboral")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paises, err := Controller.GetOpcionesPregunta(db, "SELECT distinct id, name FROM admin_main_gwdata.countries where id in(3,4,7) order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Transporte:   transporte,
		Estipendio:   estipendio,
		Estado_civ:   estado_civ,
		Actividades:  actividad,
		Depsede:      depsede,
		Grupo:        grupo,
		Idioma:       idioma,
		ActLaboral:   actLaboral,
		Pueblo:       pueblo,
		Paises:       paises,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Mentorias/mentorias_menores.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postmemteesmenores(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario

	edad := cambio(r.FormValue("edad"))
	nombre := r.FormValue("nombreCompleto")
	apellido := r.FormValue("apellido")
	sexo := cambio(r.FormValue("Sexo"))
	telefono := r.FormValue("TM")
	correo := r.FormValue("correo")
	trabaja := cambio(r.FormValue("trabajas"))
	actividadLaboral := cambio(r.FormValue("enque"))
	experienciaLaboral := cambio(r.FormValue("hijos2"))
	areaLaboral := cambio(r.FormValue("area"))
	mentoriaSabado := cambio(r.FormValue("sabados"))
	internet := cambio(r.FormValue("internet"))
	tiempoLibre := r.FormValue("tiempo")
	desde := cambio(r.FormValue("desde"))
	participacion := cambio(r.FormValue("aut1"))
	datos := cambio(r.FormValue("aut2"))
	voz_img := cambio(r.FormValue("aut3"))
	contacto := cambio(r.FormValue("aut4"))
	otro := r.FormValue("otro")
	pais := cambio(r.FormValue("pais"))
	sede := cambio(r.FormValue("programa"))

	Formulario := 7
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_participantes (nombre,apellido,sexo,pais,sede,contacto_what,situacion,descripcion,tiempoLibre,telefono,autoParticipante1,autoParticipante2,autoParticipante4,correo,trabaja,actividadLaboral,experienciaLaboral,areaLaboral,mentoriaSabado,fechaRegistro,voz_img,Formulario,edad) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(nombre, apellido, sexo, pais, sede, internet, desde, otro, tiempoLibre, telefono, participacion, datos, contacto, correo, trabaja, actividadLaboral, experienciaLaboral, areaLaboral, mentoriaSabado, currentTimeGMT, voz_img, Formulario, edad)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	vive := r.Form["vives"]
	if len(vive) == 0 {
	} else if len(vive) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO F_vive (idp, vive) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta vives", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range vive {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta vives", http.StatusInternalServerError)
				return
			}
		}
	}

	dispositivos := r.Form["dispositivo"]
	if len(dispositivos) == 0 {
	} else if len(dispositivos) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO F_dispositivos (idp, idDispositivo) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dispositivo", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range dispositivos {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta dispositivo", http.StatusInternalServerError)
				return
			}
		}
	}

	horarios := r.Form["horario"]
	if len(horarios) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_horario (idp, idHorario) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta horario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range horarios {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta horario", http.StatusInternalServerError)
				return
			}
		}
	}

	sabados := r.Form["sabados2"]
	if len(sabados) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_sabado (idp, idSabado) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta sabados", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range sabados {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta sabados", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func moduloagtHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=4 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=4  and activo=1 order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/gt/ModuloAgt.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func modulobgtHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=4 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=4  and activo=1  order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/gt/ModuloBgt.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func modulocgtHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=4 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=4 and activo=1 order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/gt/ModuloCgt.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func moduloacrHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=6 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=6  and activo=1 order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/cr/ModuloAcr.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func modulobcrHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=6 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=6  and activo=1  order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/cr/ModuloBcr.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func moduloccrHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=6 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=6 and activo=1 order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/cr/ModuloCcr.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func moduloamxHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=17 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=17  and activo=1 order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/mx/ModuloAmx.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func modulobmxHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=17 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=17  and activo=1  order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/mx/ModuloBmx.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func modulocmxHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=17 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=17 and activo=1 order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/mx/ModuloCmx.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func moduloahnHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=3 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=3 and activo=1 order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/hn/ModuloAhn.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func modulobhnHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=3 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=3 and activo=1 order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/hn/ModuloBhn.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func modulochnHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=3 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=3 and activo=1 order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/hn/ModuloChn.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func moduloColHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	var db2 *sql.DB
	connString2 := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)

	db2, err = sql.Open("sqlserver", connString2)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx2 := context.Background()
	err = db2.PingContext(ctx2)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db2.Close()

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM [Sanamente_quiz_P] where modulo=4 and idPregunta=1 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM [Sanamente_quiz_P] where modulo=4 and idPregunta=2 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=16 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM [Sanamente_quiz_P] where modulo=4 and idPregunta=3 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM [Sanamente_quiz_P] where modulo=4 and idPregunta=4 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM [Sanamente_quiz_P] where modulo=4 and idPregunta=5 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas2, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM [Sanamente_quiz_P] where modulo=4 and idPregunta=6 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas3, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM [Sanamente_quiz_P] where modulo=4 and idPregunta=7 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas4, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM [Sanamente_quiz_P] where modulo=4 and idPregunta=8 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=16 and activo=1 order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
		Respuestas2:  respuestas2,
		Respuestas3:  respuestas3,
		Respuestas4:  respuestas4,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/col/Modulocol.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func moduloMXFBEInstHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	var db2 *sql.DB
	connString2 := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)

	db2, err = sql.Open("sqlserver", connString2)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx2 := context.Background()
	err = db2.PingContext(ctx2)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db2.Close()

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM [Sanamente_quiz_P] where modulo=4 and idPregunta=1 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM [Sanamente_quiz_P] where modulo=4 and idPregunta=2 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=17 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM [Sanamente_quiz_P] where modulo=4 and idPregunta=3 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM [Sanamente_quiz_P] where modulo=4 and idPregunta=4 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM [Sanamente_quiz_P] where modulo=4 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas2, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM [Sanamente_quiz_P] where modulo=4 and idPregunta=6 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas3, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM [Sanamente_quiz_P] where modulo=4 and idPregunta=7 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas4, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM [Sanamente_quiz_P] where modulo=4 and idPregunta=8 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=17 and activo=1 order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
		Respuestas2:  respuestas2,
		Respuestas3:  respuestas3,
		Respuestas4:  respuestas4,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/mx/ModulomxFBEInst.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func moduloFBEComHandler(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	idpais := 0

	pais := r.URL.Query().Get("p")

	switch pais {
	case "sv":
		idpais = 7
	case "gt":
		idpais = 4
	case "hn":
		idpais = 3
	case "mx":
		idpais = 17
	case "col":
		idpais = 16
	case "pn":
		idpais = 1
	}

	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()

	var db2 *sql.DB
	connString2 := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)

	db2, err = sql.Open("sqlserver", connString2)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db2.Close()
	ctx2 := context.Background()
	err = db2.PingContext(ctx2)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM GWFORMS.[dbo].[Sanamente_quiz_P] where modulo=5 and idPregunta=1 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM GWFORMS.[dbo].[Sanamente_quiz_P] where modulo=5 and idPregunta=2 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	consultadepto := fmt.Sprintf("SELECT Distinct Iddep,Departamento FROM DATAGW.[dbo].[vw_Sedessanamente] where IdPais=%d order by Departamento", idpais)
	departamento, err := Controller.GetOpcionesPregunta(db2, consultadepto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM GWFORMS.[dbo].[Sanamente_quiz_P] where modulo=5 and idPregunta=3 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM GWFORMS.[dbo].[Sanamente_quiz_P] where modulo=5 and idPregunta=4 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM GWFORMS.[dbo].[Sanamente_quiz_P] where modulo=5 and idPregunta=5 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas2, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM GWFORMS.[dbo].[Sanamente_quiz_P] where modulo=5 and idPregunta=6 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas3, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM GWFORMS.[dbo].[Sanamente_quiz_P] where modulo=5 and idPregunta=7 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas4, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Respuesta FROM GWFORMS.[dbo].[Sanamente_quiz_P] where modulo=5 and idPregunta=8 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	consultaperfil := fmt.Sprintf("SELECT id,Formador FROM F_formadorSanamente where idpais=%d and activo=1 order by Formador", idpais)

	perfil, err := Controller.GetOpcionesPregunta(db2, consultaperfil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	consultapaises := fmt.Sprintf("SELECT id,name FROM DATAGW.[admin_main_gwdata].[countries] where id=%d", idpais)

	paises, err := Controller.GetOpcionesPregunta(db2, consultapaises)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	consultapaisesn := fmt.Sprintf("SELECT DISTINCT id,gentilicio FROM [GWFORMS].[dbo].[DetallesPaises] where idpais=%d", idpais)

	paisesn, err := Controller.GetOpcionesPregunta(db2, consultapaisesn)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Idpais:       idpais,
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
		Respuestas2:  respuestas2,
		Respuestas3:  respuestas3,
		Respuestas4:  respuestas4,
		Paises:       paises,
		Paisesn:      paisesn,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/ModuloFBECom.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func moduloapnHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=1 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=1 and activo=1 order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/pn/ModuloApn.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func modulobpnHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=1 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=1 and activo=1 order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/pn/ModuloBpn.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func modulocpnHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=1 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=1 and activo=1 order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/pn/ModuloCpn.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func moduloInterv(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	stmt, err := db.Prepare("EXECUTE [dbo].[spSanamenteInterv] @p1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pais := r.URL.Query().Get("pais")
	idpais := 0
	switch pais {
	case "pa":
		idpais = 1
	case "sv":
		idpais = 7
	case "hn":
		idpais = 3
	case "gt":
		idpais = 4
	case "mx":
		idpais = 17
	case "co":
		idpais = 16
	}

	// ejecutar procedimiento almacenado
	rows, err := stmt.Query(idpais)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close() // Create a slice to hold the data

	type Opciones struct {
		Id           int
		Respuesta    string
		Idpregunta   int
		OrdenInterno string
	}

	var Filtros []Opciones

	for rows.Next() {
		var Opciones Opciones
		err := rows.Scan(&Opciones.Id, &Opciones.Respuesta, &Opciones.Idpregunta, &Opciones.OrdenInterno)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		Filtros = append(Filtros, Opciones)
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/ModuloIntervencionistas.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, Filtros)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func formulariopercepcionmdv(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas2, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_opcRespuestas")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas3, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_opcEntorno where id in (1,2,3)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas4, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_opcEntorno where id in (4,5,6)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Sexo:        sexo,
		Respuestas:  respuestas,
		Respuestas2: respuestas2,
		Respuestas3: respuestas3,
		Respuestas4: respuestas4,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/MdV/percepcion.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func formulariomdvred(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estado_civ, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_estadociv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actividad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_ofertavol")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.Vista_Sedes where IdSede in(666,1128)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Estado_civ:   estado_civ,
		Actividades:  actividad,
		Depsede:      depsede,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/mdv/red.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func formulariomdvfablab(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estado_civ, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_estadociv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actividad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_ofertavol")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.Vista_Sedes where IdSede in(666,1128)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Estado_civ:   estado_civ,
		Actividades:  actividad,
		Depsede:      depsede,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/mdv/fablab.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postmdvpercepcion(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	//NombreCampo en la BD: = r.FormValue("NombreCampo en el html")
	// Obtener los valores de los campos del formulario

	cp1 := r.FormValue("participar")
	cp2 := r.FormValue("actividad")
	cp3 := r.FormValue("anterior")
	edad := cambio(r.FormValue("edad"))
	sexo := cambio(r.FormValue("Sexo"))
	p1 := r.FormValue("p1")
	p2 := cambio(r.FormValue("p2"))
	p2_1 := r.FormValue("p21")
	p3 := cambio(r.FormValue("p3"))
	p3_1 := r.FormValue("p31")
	p4 := cambio(r.FormValue("p4"))
	p4_1 := r.FormValue("p41")
	p5 := cambio(r.FormValue("p5"))
	p5_1 := r.FormValue("p51")
	p6 := cambio(r.FormValue("p6"))
	p6_1 := r.FormValue("p61")
	p7 := cambio(r.FormValue("p7"))
	p8 := cambio(r.FormValue("p8"))
	p9 := cambio(r.FormValue("p9"))
	p9_1 := r.FormValue("p91")
	p10 := cambio(r.FormValue("p10"))
	p10_1 := r.FormValue("p101")
	p11 := r.FormValue("p11")
	p12 := r.FormValue("p12")

	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_EncuestaPercepcion (fechaRegistro,cp1 ,cp2,cp3 ,edad ,sexo ,p1 ,p2 ,p2_1 ,p3 ,p3_1 ,p4 ,p4_1 ,p5 ,p5_1 ,p6 ,p6_1 ,p7 ,p8 ,p9 ,p9_1 ,p10 ,p10_1 ,p11 ,p12) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?, ?, ?, ?,?, ?, ?, ?,?, ?, ?, ? ,?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(currentTimeGMT, cp1, cp2, cp3, edad, sexo, p1, p2, p2_1, p3, p3_1, p4, p4_1, p5, p5_1, p6, p6_1, p7, p8, p9, p9_1, p10, p10_1, p11, p12)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	r.ParseForm()

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func postredmdv(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los valores de los campos del formulario
	Nombrerepresentante := r.FormValue("nombreRep")
	Duirepresentante := r.FormValue("DUIP")
	parentesco := r.FormValue("parentesco")
	parentescop := cambio(parentesco)
	edad := r.FormValue("edad")
	edadp := cambio(edad)
	nombreAdolescente := r.FormValue("nombreAdolescente")
	telefonoresp := r.FormValue("numResp")
	nombre := r.FormValue("nombreCompleto")
	apellido := r.FormValue("apellido")
	sexo := r.FormValue("Sexo")
	sexop := cambio(sexo)
	pais := 7
	departamento := r.FormValue("pais")
	departamentop := cambio(departamento)
	municipio := r.FormValue("departamento")
	municipiop := cambio(municipio)
	sede := r.FormValue("sede2")
	sedep := cambio(sede)
	diplomado := r.FormValue("diplomado")
	diplomadop := cambio(diplomado)
	direccion := r.FormValue("direccion")
	tipo := 1
	telefono := r.FormValue("TM")
	contacto_what := r.FormValue("WS")
	contacto_whatp := cambio(contacto_what)
	telefono_ws := r.FormValue("TM2")
	Estudia := r.FormValue("TdC")
	Estudiap := cambio(Estudia)
	ultimoestudio := r.FormValue("UT")
	ultimoestudiop := cambio(ultimoestudio)
	grado := r.FormValue("GA")
	gradop := cambio(grado)
	seccion := r.FormValue("seccion")
	seccionp := cambio(seccion)
	turno := r.FormValue("Turno")
	turnop := cambio(turno)
	discapacidad := r.FormValue("dis")
	dispacacidadp := cambio(discapacidad)
	nuevogw := r.FormValue("gp")
	nuevogwp := cambio(nuevogw)
	actividadparticipo := r.FormValue("ActG")
	actividadparticipop := cambio(actividadparticipo)
	carnet := r.FormValue("carnet")
	nombreE := r.FormValue("contacto")
	contactoE := r.FormValue("contacto2")
	parentescoE := r.FormValue("relacion")
	parentescoEp := cambio(parentescoE)
	transporte := r.FormValue("Transp")
	transportep := cambio(transporte)
	gasto := r.FormValue("gasto")
	gastop := cambio(gasto)
	otro := r.FormValue("cantidad")
	hijos := r.FormValue("hijos")
	hijosp := cambio(hijos)
	voz_img := r.FormValue("aut2")
	voz_imgp := cambio(voz_img)
	DUI := r.FormValue("DUIM")
	urldocF := ""
	urldocA := ""
	Formulario := r.FormValue("formulario")
	FechaNac := r.FormValue("FechaN")
	nacionalidad := r.FormValue("nacionalidad")
	nacionalidadp := cambio(nacionalidad)
	rjfor := cambio(r.FormValue("rjfor"))
	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO F_participantes (Nombrerepresentante,Duirepresentante,parentesco,edad,nombreAdolescente,telefonoresp,nombre,apellido,sexo,pais,departamento,municipio,sede,diplomado,direccion,tipo,telefono,contacto_what,telefono_ws,Estudia,ultimoestudio,grado,seccion,turno,discapacidad,nuevogw,actividadparticipo,carnet,nombreE,contactoE,parentescoE,transporte,gasto,otro_gasto,hijos,voz_img,DUI,urldui,urlduia, Formulario,FechaNac,nacionalidad,fechaRegistro,actividadLaboral) VALUES (?,?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(Nombrerepresentante, Duirepresentante, parentescop, edadp, nombreAdolescente, telefonoresp, nombre, apellido, sexop, pais, departamentop, municipiop, sedep, diplomadop, direccion, tipo, telefono, contacto_whatp, telefono_ws, Estudiap, ultimoestudiop, gradop, seccionp, turnop, dispacacidadp, nuevogwp, actividadparticipop, carnet, nombreE, contactoE, parentescoEp, transportep, gastop, otro, hijosp, voz_imgp, DUI, urldocF, urldocA, Formulario, FechaNac, nacionalidadp, currentTimeGMT, rjfor)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	discapacidad2 := r.Form["dis2"]
	if len(discapacidad2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO F_discapacidad (idp, discapacidad) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dis", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range discapacidad2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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

func equipsv(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=1 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=2 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct id,Perfil FROM admin_main_gwdata.Sanamente_perfiles where activo=1 and pais=7 order by Perfil")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=3 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=4 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=5 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas2, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=6 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas3, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=7 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=7  and activo=1 order by Formador;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	programa, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM admin_main_gwdata.sub_components WHERE id IN(106,43,42) ORDER BY name ASC;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Respuestas2:  respuestas2,
		Respuestas3:  respuestas3,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
		Programa:     programa,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/sv/Formadores.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func equipgt(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=1 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=2 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct id,Perfil FROM admin_main_gwdata.Sanamente_perfiles where activo=1 and pais=4 order by Perfil")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=3 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=4 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=5 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas2, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=6 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas3, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=7 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=4  and activo=1 order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	programa, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM admin_main_gwdata.sub_components WHERE id IN(106,43,42) ORDER BY name ASC;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Respuestas2:  respuestas2,
		Respuestas3:  respuestas3,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
		Programa:     programa,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/gt/Formadores.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func equipmx(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=1 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=2 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct id,Perfil FROM admin_main_gwdata.Sanamente_perfiles where activo=1 and pais=17 order by Perfil")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=3 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=4 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=5 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas2, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=6 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas3, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=7 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=17  and activo=1 order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	programa, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM admin_main_gwdata.sub_components WHERE id IN(106,43,42) ORDER BY name ASC;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Respuestas2:  respuestas2,
		Respuestas3:  respuestas3,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
		Programa:     programa,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/mx/Formadores.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func equiphn(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=1 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=2 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct id,Perfil FROM admin_main_gwdata.Sanamente_perfiles where activo=1 and pais=3 order by Perfil")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=3 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=4 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=5 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas2, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=6 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas3, err := Controller.GetOpcionesPregunta(db, "SELECT id,comportamiento FROM admin_main_gwdata.Sanamente_quiz_F where idCompetencia=7 and Version=0;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=3  and activo=1 order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	programa, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM admin_main_gwdata.sub_components WHERE id IN(106,43,42) ORDER BY name ASC;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Respuestas2:  respuestas2,
		Respuestas3:  respuestas3,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
		Programa:     programa,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/hn/Formadores.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postequip(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	// Obtener los valores de los campos del formulario
	email := r.FormValue("email")
	perfil := r.FormValue("perfil")
	formador := cambio(r.FormValue("apellido"))
	tipoformacion := cambio(r.FormValue("formaciones"))
	subtipoformacion := cambio(r.FormValue("fentrenamiento"))
	pais := r.FormValue("idp")
	p1 := r.FormValue("p1obs")
	p2 := r.FormValue("p2obs")
	p3 := r.FormValue("p3obs")
	p4 := r.FormValue("p4obs")
	p5 := r.FormValue("p5obs")
	p6 := r.FormValue("p6obs")
	p7 := r.FormValue("p7obs")
	m1 := cambio(r.FormValue("m1"))
	m2 := cambio(r.FormValue("m2"))
	m3 := cambio(r.FormValue("m3"))
	fechaAplicacion := r.FormValue("faplicacion")

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO Sanamente_equip (fechaRegistro,email , perfil, formador,pais, p1,p2,p3,p4,p5,p6,p7,m1,m2,m3,fechaAplicacion,tipoformacion,subtipoformacion) VALUES (?,?,?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?,?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(currentTimeGMT, email, perfil, formador, pais, p1, p2, p3, p4, p5, p6, p7, m1, m2, m3, fechaAplicacion, tipoformacion, subtipoformacion)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}

	idGenerado, err := resultado.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	r.ParseForm()

	p1r := r.Form["p1"]

	if len(p1r) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_equipd (idp, respuesta) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dias", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range p1r {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta p1", http.StatusInternalServerError)
				return
			}
		}
	}

	p2r := r.Form["p2"]

	if len(p1r) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_equipd (idp, respuesta) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dias", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range p2r {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta p2", http.StatusInternalServerError)
				return
			}
		}
	}

	p3r := r.Form["p3"]

	if len(p3r) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_equipd (idp, respuesta) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dias", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range p3r {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta p3", http.StatusInternalServerError)
				return
			}
		}
	}

	p4r := r.Form["p4"]

	if len(p4r) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_equipd (idp, respuesta) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dias", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range p4r {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta p4", http.StatusInternalServerError)
				return
			}
		}
	}

	p5r := r.Form["p5"]

	if len(p1r) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_equipd (idp, respuesta) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dias", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range p5r {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta p5", http.StatusInternalServerError)
				return
			}
		}
	}

	p6r := r.Form["p6"]

	if len(p1r) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_equipd (idp, respuesta) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dias", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range p6r {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta p6", http.StatusInternalServerError)
				return
			}
		}
	}

	p7r := r.Form["p7"]

	if len(p1r) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_equipd (idp, respuesta) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta dias", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range p7r {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta p7", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	t, err := template.ParseFiles("public/HTML/RespuestaSanamente.html")
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

func intervencionistas(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	idpais := 0

	pais := r.URL.Query().Get("p")

	switch pais {
	case "sv":
		idpais = 7
	case "gt":
		idpais = 4
	case "hn":
		idpais = 3
	case "mx":
		idpais = 17
	case "col":
		idpais = 16
	case "pn":
		idpais = 1
	}

	var db2 *sql.DB
	connString2 := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)

	db2, err := sql.Open("sqlserver", connString2)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx2 := context.Background()
	err = db2.PingContext(ctx2)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db2.Close()

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db2, "SELECT id,comportamiento FROM Sanamente_quiz_Form where idCompetencia=1 and versionf=0 and nivel=2;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db2, "SELECT id,comportamiento FROM Sanamente_quiz_Form where idCompetencia=2 and versionf=0 and nivel=2;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/* 	consultadepto := fmt.Sprintf("SELECT Distinct id,Perfil FROM DATAGW.admin_main_gwdata.Sanamente_perfiles where pais=%d and activo=1 order by Perfil", idpais)

	   	departamento, err := Controller.GetOpcionesPregunta(db2, consultadepto)
	   	if err != nil {
	   			http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	   	} */

	departamento, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Perfil FROM Sanamente_Perfiles where pais is null and activo=1 order by Perfil;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db2, "SELECT id,comportamiento FROM Sanamente_quiz_Form where idCompetencia=3 and versionf=0 and nivel=2;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db2, "SELECT id,comportamiento FROM Sanamente_quiz_Form where idCompetencia=4 and versionf=0 and nivel=2;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db2, "SELECT id,comportamiento FROM Sanamente_quiz_Form where idCompetencia=5 and versionf=0 and nivel=2;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas2, err := Controller.GetOpcionesPregunta(db2, "SELECT id,comportamiento FROM Sanamente_quiz_Form where idCompetencia=6 and versionf=0 and nivel=2;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas3, err := Controller.GetOpcionesPregunta(db2, "SELECT id,comportamiento FROM Sanamente_quiz_Form where idCompetencia=7 and versionf=0 and nivel=2;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	consultaperfil := fmt.Sprintf("SELECT id,intervencionista FROM F_IntervencionistaSanamente where idpais=%d and activo=1 order by intervencionista", idpais)

	perfil, err := Controller.GetOpcionesPregunta(db2, consultaperfil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Respuestas2:  respuestas2,
		Respuestas3:  respuestas3,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
		Idpais:       idpais,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/sv/Intervencionistas.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postintervencionistas(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)

	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	// Obtener los valores de los campos del formulario
	email := r.FormValue("email")
	perfil := r.FormValue("perfil")
	formador := cambio(r.FormValue("apellido"))
	pais := r.FormValue("idp")
	p1 := r.FormValue("p1obs")
	p2 := r.FormValue("p2obs")
	p3 := r.FormValue("p3obs")
	p4 := r.FormValue("p4obs")
	p5 := r.FormValue("p5obs")
	p6 := r.FormValue("p6obs")
	p7 := r.FormValue("p7obs")
	// m1 := cambio(r.FormValue("m1"))
	// m2 := cambio(r.FormValue("m2"))
	// m3 := cambio(r.FormValue("m3"))
	fechaAplicacion := r.FormValue("faplicacion")
	idF := 8

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO Sanamente_Equip (idF, email, perfil, formador, pais, p1, p2, p3, p4, p5, p6, p7, fechaAplicacion) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)

	}

	var lastInsertId int
	err = stmt.QueryRow(idF, email, perfil, formador, pais, p1, p2, p3, p4, p5, p6, p7, fechaAplicacion).Scan(&lastInsertId)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
	}

	p1r := r.Form["p1"]

	if len(p1r) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_Equipd (idp, respuesta) VALUES (@P1, @P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta competencia 1", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range p1r {
			_, err = stmt2.Exec(lastInsertId, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta competencia 1", http.StatusInternalServerError)
				return
			}
		}
	}

	p2r := r.Form["p2"]

	if len(p1r) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_Equipd (idp, respuesta) VALUES (@P1, @P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta competencia 2", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range p2r {
			_, err = stmt2.Exec(lastInsertId, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta competencia 2", http.StatusInternalServerError)
				return
			}
		}
	}

	p3r := r.Form["p3"]

	if len(p3r) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_Equipd (idp, respuesta) VALUES (@P1, @P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta competencia 3", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range p3r {
			_, err = stmt2.Exec(lastInsertId, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta competencia 3", http.StatusInternalServerError)
				return
			}
		}
	}

	p4r := r.Form["p4"]

	if len(p4r) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_Equipd (idp, respuesta) VALUES (@P1, @P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta competencia 4", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range p4r {
			_, err = stmt2.Exec(lastInsertId, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta competencia 4", http.StatusInternalServerError)
				return
			}
		}
	}

	p5r := r.Form["p5"]

	if len(p1r) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_Equipd (idp, respuesta) VALUES (@P1, @P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta competencia 5", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range p5r {
			_, err = stmt2.Exec(lastInsertId, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta competencia 5", http.StatusInternalServerError)
				return
			}
		}
	}

	p6r := r.Form["p6"]

	if len(p1r) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_Equipd (idp, respuesta) VALUES (@P1, @P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta competencia 6", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range p6r {
			_, err = stmt2.Exec(lastInsertId, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta competencia 6", http.StatusInternalServerError)
				return
			}
		}
	}

	p7r := r.Form["p7"]

	if len(p1r) != 0 {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_Equipd (idp, respuesta) VALUES (@P1, @P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta competencia 7", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range p7r {
			_, err = stmt2.Exec(lastInsertId, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta competencia 7", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	t, err := template.ParseFiles("public/HTML/RespuestaSanamenteInterv.html")
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

func equipGeneral(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	idpais := 0

	pais := r.URL.Query().Get("p")

	switch pais {
	case "sv":
		idpais = 7
	case "gt":
		idpais = 4
	case "hn":
		idpais = 3
	case "mx":
		idpais = 17
	case "col":
		idpais = 16
	case "pn":
		idpais = 1
	}

	var db2 *sql.DB
	connString2 := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)

	db2, err := sql.Open("sqlserver", connString2)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx2 := context.Background()
	err = db2.PingContext(ctx2)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db2.Close()

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db2, "SELECT id,comportamiento FROM Sanamente_quiz_Form where idCompetencia=1 and versionf=0 and nivel=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	sexo, err := Controller.GetOpcionesPregunta(db2, "SELECT id,comportamiento FROM Sanamente_quiz_Form where idCompetencia=2 and versionf=0 and nivel=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	consultadepto := fmt.Sprintf("SELECT Distinct id,Perfil FROM DATAGW.admin_main_gwdata.Sanamente_perfiles where pais=%d and activo=1 order by Perfil", idpais)

	departamento, err := Controller.GetOpcionesPregunta(db2, consultadepto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// departamento, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Perfil FROM Sanamente_Perfiles where pais is null and activo=1 order by Perfil;")
	// if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	respuestas, err := Controller.GetOpcionesPregunta(db2, "SELECT id,comportamiento FROM Sanamente_quiz_Form where idCompetencia=3 and versionf=0 and nivel=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db2, "SELECT id,comportamiento FROM Sanamente_quiz_Form where idCompetencia=4 and versionf=0 and nivel=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db2, "SELECT id,comportamiento FROM Sanamente_quiz_Form where idCompetencia=5 and versionf=0 and nivel=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas2, err := Controller.GetOpcionesPregunta(db2, "SELECT id,comportamiento FROM Sanamente_quiz_Form where idCompetencia=6 and versionf=0 and nivel=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas3, err := Controller.GetOpcionesPregunta(db2, "SELECT id,comportamiento FROM Sanamente_quiz_Form where idCompetencia=7 and versionf=0 and nivel=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	consultaperfil := fmt.Sprintf("SELECT id,Formador FROM F_formadorSanamente where idpais=%d and activo=1 order by Formador", idpais)

	perfil, err := Controller.GetOpcionesPregunta(db2, consultaperfil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	programa, err := Controller.GetOpcionesPregunta(db2, "SELECT id,name FROM DATAGW.admin_main_gwdata.sub_components WHERE id IN(106,43,42) ORDER BY name ASC;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Respuestas2:  respuestas2,
		Respuestas3:  respuestas3,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
		Idpais:       idpais,
		Programa:     programa,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/FormadoresGeneral.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func equipGeneralGuardado(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	idpais := 0

	pais := r.URL.Query().Get("p")

	switch pais {
	case "sv":
		idpais = 7
	case "gt":
		idpais = 4
	case "hn":
		idpais = 3
	case "mx":
		idpais = 17
	case "col":
		idpais = 16
	case "pn":
		idpais = 1
	}

	var db2 *sql.DB
	connString2 := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)

	db2, err := sql.Open("sqlserver", connString2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ctx2 := context.Background()
	err = db2.PingContext(ctx2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db2.Close()

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db2, "SELECT id,comportamiento FROM Sanamente_quiz_Form where idCompetencia=1 and versionf=0 and nivel=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db2, "SELECT id,comportamiento FROM Sanamente_quiz_Form where idCompetencia=2 and versionf=0 and nivel=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	consultadepto := fmt.Sprintf("SELECT Distinct id,Perfil FROM DATAGW.admin_main_gwdata.Sanamente_perfiles where pais=%d and activo=1 order by Perfil", idpais)

	departamento, err := Controller.GetOpcionesPregunta(db2, consultadepto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// departamento, err := Controller.GetOpcionesPregunta(db2, "SELECT id,Perfil FROM Sanamente_Perfiles where pais is null and activo=1 order by Perfil;")
	// if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	respuestas, err := Controller.GetOpcionesPregunta(db2, "SELECT id,comportamiento FROM Sanamente_quiz_Form where idCompetencia=3 and versionf=0 and nivel=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db2, "SELECT id,comportamiento FROM Sanamente_quiz_Form where idCompetencia=4 and versionf=0 and nivel=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db2, "SELECT id,comportamiento FROM Sanamente_quiz_Form where idCompetencia=5 and versionf=0 and nivel=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas2, err := Controller.GetOpcionesPregunta(db2, "SELECT id,comportamiento FROM Sanamente_quiz_Form where idCompetencia=6 and versionf=0 and nivel=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas3, err := Controller.GetOpcionesPregunta(db2, "SELECT id,comportamiento FROM Sanamente_quiz_Form where idCompetencia=7 and versionf=0 and nivel=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	consultaperfil := fmt.Sprintf("SELECT id,Formador FROM F_formadorSanamente where idpais=%d and activo=1 order by Formador", idpais)

	perfil, err := Controller.GetOpcionesPregunta(db2, consultaperfil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	programa, err := Controller.GetOpcionesPregunta(db2, "SELECT id,name FROM DATAGW.admin_main_gwdata.sub_components WHERE id IN(106,43,42) ORDER BY name ASC;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Respuestas2:  respuestas2,
		Respuestas3:  respuestas3,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
		Idpais:       idpais,
		Programa:     programa,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/FormadoresGeneralGuardado.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func obtenerperfil(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	pais := r.FormValue("pais")
	idpais := r.FormValue("idp")
	var opcionesCiudades string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT id,name FROM admin_main_gwdata.sanamente_subtipos where if(3<>?,id_institucional=? and id not in (12,13),id_institucional=?) order by name", idpais, pais, pais)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionesCiudades += ("<option value=''></option>")
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesCiudades += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionesCiudades)
}

func obtenersubperfil(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	pais := r.FormValue("pais")
	var opcionesCiudades string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT id,name FROM admin_main_gwdata.beneficiaries_subtypes where id not in(17,61,63,64,65,70,71) and institutional_person_id=? order by name", pais)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionesCiudades += ("<option value=''></option>")
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesCiudades += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}

	fmt.Fprint(w, opcionesCiudades)
}

func postsanamenteinscripcion(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	idF := 4
	pais := r.FormValue("idp")
	idSede := cambio(r.FormValue("sede2"))
	fechaNac := r.FormValue("FechaN")
	nombres := Controller.QuitarTildesYMayusculas(r.FormValue("nombreCompleto"))
	apellidos := Controller.QuitarTildesYMayusculas(r.FormValue("apellido"))
	sexo := cambio(r.FormValue("Sexo"))
	nacionalidad := cambio(r.FormValue("nacionalidad"))
	departamento := cambio(r.FormValue("pais"))
	municipio := cambio(r.FormValue("departamento"))
	discapacidad := cambio(r.FormValue("dis"))
	estudia := cambio(r.FormValue("TdC"))
	estudioAlcanzado := cambio(r.FormValue("UT"))
	grado := cambio(r.FormValue("GA"))
	turno := cambio(r.FormValue("Turno"))
	seccion := cambio(r.FormValue("seccion"))
	nuevogw := cambio(r.FormValue("gp"))
	autovozimg := cambio(r.FormValue("aut2"))
	tipoParticipante := r.FormValue("tipo")
	personalinst := cambio(r.FormValue("acc2"))
	perfil := cambio(r.FormValue("acc3"))
	subtipo := cambio(r.FormValue("acc4"))
	correo := r.FormValue("correo")
	telefono := r.FormValue("TM")

	identidad := r.FormValue("DUIM")

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	stmt, err := db.Prepare("IF EXISTS(SELECT * FROM FormularioRegisto WHERE nombres = @P1 and apellidos = @P2 and identidad = @P3 and pais = @P4 and fechaNAc = @P5 and idF = @P6 and deleted is null)Begin select 1 end else Begin select 0 end")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	row := stmt.QueryRow(nombres, apellidos, identidad, pais, fechaNac, idF)

	var result int
	_ = row.Scan(&result)

	type Datos struct {
		Nombre   string
		Apellido string
		Id       int
		Programa string
	}

	if result == 1 {

		datos := Datos{
			Nombre:   nombres,
			Apellido: apellidos,
		}

		t, err := template.ParseFiles("Public/Html/RespuestaDuplicadaSanamente.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Renderizar la plantilla con los datos
		err = t.Execute(w, datos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		tx, err := db.Begin()
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error conexion", http.StatusInternalServerError)

		}

		stmt, err := tx.Prepare("INSERT INTO FormularioRegisto(idF,tipoParticipante,pais,idSede,fechaNac,nombres,apellidos,sexo,nacionalidad,identidad,discapacidad,estudia,estudioAlcanzado,grado,turno,seccion,nuevogw,autovozimg,departamento,municipio, telefono, correo,personalinst,perfil,subtipo) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14,@P15,@P16,@P17,@P18,@P19,@P20,@P21,@P22,@P23,@P24,@P25)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		var lastInsertId int
		err = stmt.QueryRow(idF, tipoParticipante, pais, idSede, fechaNac, nombres, apellidos, sexo, nacionalidad, identidad, discapacidad, estudia, estudioAlcanzado, grado, turno, seccion, nuevogw, autovozimg, departamento, municipio, telefono, correo, personalinst, perfil, subtipo).Scan(&lastInsertId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		discapacidadmultiple := r.Form["dis2"]
		if len(discapacidadmultiple) == 0 {
		} else {
			stmt, err := tx.Prepare("INSERT INTO MultipleDiscapacidad (idFR,iddiscapacidad) VALUES (@P1,@P2)")
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al preparar la consulta ocupacion", http.StatusInternalServerError)
				return
			}
			defer stmt.Close()

			for _, valor := range discapacidadmultiple {
				if valor != "" {
					_, err = stmt.Exec(lastInsertId, valor)
					if err != nil {
						tx.Rollback()
						http.Error(w, "Error al ejecutar la consulta situaciones", http.StatusInternalServerError)
						return
					}
				}
			}
		}

		err = tx.Commit()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		programa := "Sanamente"

		datos := Datos{
			Nombre:   nombres,
			Apellido: apellidos,
			Id:       lastInsertId,
			Programa: programa,
		}

		t, err := template.ParseFiles("Public/Html/RespuestaPersonalizada.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Renderizar la plantilla con los datos
		err = t.Execute(w, datos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func rrhhferia(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	var sangre int
	err = db.QueryRow("SELECT sum(sangre) FROM admin_main_gwdata.feria_salud where eliminado is null").Scan(&sangre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sangre = 40 - sangre

	var citologia int
	err = db.QueryRow("SELECT sum(citologia) FROM admin_main_gwdata.feria_salud where eliminado is null").Scan(&citologia)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	citologia = 50 - citologia

	var yoga int
	err = db.QueryRow("SELECT sum(yoga) FROM admin_main_gwdata.feria_salud where eliminado is null").Scan(&yoga)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	yoga = 15 - yoga

	var baile int
	err = db.QueryRow("SELECT sum(baile) FROM admin_main_gwdata.feria_salud where eliminado is null").Scan(&baile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	baile = 15 - baile

	var saludables int
	err = db.QueryRow("SELECT sum(saludables) FROM admin_main_gwdata.feria_salud where eliminado is null").Scan(&saludables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	saludables = 20 - saludables

	type FormData2 struct {
		Sangre     int
		Citologia  int
		Yoga       int
		Baile      int
		Saludables int
	}
	data := FormData2{
		Sangre:     sangre,
		Citologia:  citologia,
		Yoga:       yoga,
		Baile:      baile,
		Saludables: saludables,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/RRHH/feriasalud.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postrrhhferiasalud(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	currentTimeUTC := time.Now().UTC()
	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		fmt.Println(err)
	}
	currentTimeGMT := currentTimeUTC.In(location)

	// Obtener los valores de los campos del formulario

	nombre := r.FormValue("nombre")
	dui := r.FormValue("dui")
	area := r.FormValue("area")
	puesto := r.FormValue("puesto")
	correo := r.FormValue("correo")
	celular := r.FormValue("celular")
	sangre := cambio(r.FormValue("sangre"))
	citologia := cambio(r.FormValue("citologia"))
	yoga := cambio(r.FormValue("yoga"))
	baile := cambio(r.FormValue("baile"))
	saludable := cambio(r.FormValue("saludable"))
	mensaje := ""
	mensajec := ""
	mensajes := ""
	mensajey := ""
	mensajeb := ""
	formulario := 1
	var sangrec int
	err = db.QueryRow("SELECT sum(sangre) FROM admin_main_gwdata.feria_salud where eliminado is null").Scan(&sangrec)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sangrec = 40 - sangrec - sangre
	if sangrec < 0 && sangre > 0 {
		sangre = 0
		mensaje = "Lamentamos informarte que se han agotado los cupos disponibles para donar sangre. Te invitamos a estar pendiente de otras actividades."
	}
	var citologiac int
	err = db.QueryRow("SELECT sum(citologia) FROM admin_main_gwdata.feria_salud where eliminado is null").Scan(&citologiac)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	citologiac = 50 - citologiac - citologia
	if citologiac < 0 && citologia > 0 {
		citologia = 0
		mensajec = "Lamentamos informarte que se han agotado los cupos disponibles para exámenes de citologías. Te invitamos a estar pendiente de otras actividades."
	}
	var yogac int
	err = db.QueryRow("SELECT sum(yoga) FROM admin_main_gwdata.feria_salud where eliminado is null").Scan(&yogac)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	yogac = 15 - yogac - yoga
	if yogac < 0 && yoga > 0 {
		yoga = 0
		mensajey = "Lamentamos informarte que se han agotado los cupos disponibles para yoga. Te invitamos a estar pendiente de otras actividades."
	}
	var bailec int
	err = db.QueryRow("SELECT sum(baile) FROM admin_main_gwdata.feria_salud where eliminado is null").Scan(&bailec)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bailec = 15 - bailec - baile
	if bailec < 0 && baile > 0 {
		bailec = 0
		mensajeb = "Lamentamos informarte que se han agotado los cupos disponibles para baile corporativo. Te invitamos a estar pendiente de otras actividades."
	}
	var saludablesc int
	err = db.QueryRow("SELECT sum(saludables) FROM admin_main_gwdata.feria_salud where eliminado is null").Scan(&saludablesc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	saludablesc = 20 - saludablesc - saludable
	if saludablesc < 0 && saludable > 0 {
		saludable = 0
		mensajes = "Lamentamos informarte que se han agotado los cupos disponibles para la sesiones en vivo de preparación de menús saludables. Te invitamos a estar pendiente de otras actividades."
	}
	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente
	stmt, err := tx.Prepare("INSERT INTO feria_salud (fechaRegistro, nombre_completo, DUI, area, puesto, email, telefono, sangre, citologia, yoga, baile, saludables,formulario) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? ,?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(currentTimeGMT, nombre, dui, area, puesto, correo, celular, sangre, citologia, yoga, baile, saludable, formulario)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v", err)
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		return
	}
	r.ParseForm()

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	type Mensajes struct {
		Mensajes string
		Mensajec string
		Mensajey string
		Mensajeb string
		Mensaje  string
	}

	data := Mensajes{
		Mensajes: mensajes,
		Mensajec: mensajec,
		Mensajey: mensajey,
		Mensajeb: mensajeb,
		Mensaje:  mensaje,
	}

	t, err := template.ParseFiles("public/HTML/RRHH/Respuesta.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar la plantilla con los datos
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func dell(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	institucion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM schools where id in (3279,4144,4145)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Institucion:  institucion,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Otros/Dell_ins.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postdell(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Obtener los valores de los campos del formulario

	// Autorizacion Menores
	nombrecompleto := r.FormValue("nombreRep")
	identidad := r.FormValue("DUIP")
	parentesco := r.FormValue("parentesco")
	infoAdicional1 := r.FormValue("nombreAdolescente")
	telefono := r.FormValue("numResp")
	telefono2 := r.FormValue("numRespNO")

	// Formulario Registro
	idF := 11
	tipoParticipante := 1
	fechaNac := r.FormValue("FechaN")
	nombres := r.FormValue("nombreCompleto")
	apellidos := r.FormValue("apellido")
	sexo := r.FormValue("Sexo")
	nacionalidad := r.FormValue("nacionalidad")
	identidad2 := r.FormValue("DUIM")
	pais := 1
	departamento := r.FormValue("pais")
	municipio := r.FormValue("departamento")
	idSede := r.FormValue("sedes")
	estudia := r.FormValue("TdC")
	estudioAlcanzado := r.FormValue("UT")
	grado := r.FormValue("GA")
	seccion := r.FormValue("seccion")
	turno := r.FormValue("Turno")
	discapacidad := r.FormValue("dis")
	nuevogw := r.FormValue("gp")
	autovozimg := r.FormValue("aut2")
	infopersonal := r.FormValue("aut")

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error conexion", http.StatusInternalServerError)
	}

	if infopersonal != "2" {
		stmt, err := tx.Prepare("INSERT INTO FormularioRegisto(idF,tipoParticipante,pais,idSede,fechaNac,nombres,apellidos,sexo,nacionalidad,identidad,discapacidad,estudia,estudioAlcanzado,grado,turno,seccion,nuevogw,departamento,municipio,autovozimg) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14,@P15,@P16,@P17,@P18,@P19,@P20)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		}

		var lastInsertId int
		err = stmt.QueryRow(idF, tipoParticipante, pais, idSede, fechaNac, nombres, apellidos, sexo, nacionalidad, identidad2, discapacidad, estudia, estudioAlcanzado, grado, turno, seccion, nuevogw, departamento, municipio, autovozimg).Scan(&lastInsertId)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
		}

		if identidad2 != "" {
		} else {
			stmt, err = tx.Prepare("INSERT INTO AutorizacionMenores(idF,idFR,nombrecompleto, telefono, identidad, parentesco,infoAdicional1) VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7)")
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al preparar la consulta automenores", http.StatusInternalServerError)
			}

			_, err = stmt.Exec(11, lastInsertId, nombrecompleto, telefono, identidad, parentesco, infoAdicional1)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta automenores", http.StatusInternalServerError)
			}
		}

		discapacidadmultiple := r.Form["dis2"]
		if len(discapacidadmultiple) == 0 {
		} else {
			stmt, err := tx.Prepare("INSERT INTO MultipleDiscapacidad (idFR,iddiscapacidad) VALUES (@P1,@P2)")
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al preparar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
			defer stmt.Close()

			for _, valor := range discapacidadmultiple {
				if valor != "" {
					_, err = stmt.Exec(lastInsertId, valor)
					if err != nil {
						tx.Rollback()
						http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
						return
					}
				}
			}
		}
	} else {
		stmt, err := tx.Prepare("INSERT INTO AutorizacionMenores(idF,nombrecompleto, telefono, identidad, parentesco,infoAdicional1) VALUES(@P1,@P2,@P3,@P4,@P5,@P6)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta automenores", http.StatusInternalServerError)
		}

		_, err = stmt.Exec(11, nombrecompleto, telefono2, identidad, parentesco, infoAdicional1)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al ejecutar la consulta automenores", http.StatusInternalServerError)
		}
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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

func rrhhtransporte(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Departamento: departamento,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/RRHH/transporte.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postrrhhtransporte(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	// Create connection pool
	nombre := r.FormValue("nombre")
	area := r.FormValue("areas")
	otro := r.FormValue("areaotro")
	puesto := r.FormValue("puesto")
	sede := cambio(r.FormValue("sede"))
	transporte := cambio(r.FormValue("transporte"))
	proyecto := r.FormValue("proyecto")
	departamento := cambio(r.FormValue("Departamento"))
	municipio := cambio(r.FormValue("Municipio"))
	direccion := r.FormValue("direccion")
	profesion := r.FormValue("profesion")

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database2)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error conexion", http.StatusInternalServerError)

	}

	stmt, err := tx.Prepare("INSERT INTO Respuestas_transporte(nombre, area, otro, puesto, sede, transporte, proyecto, departamento, municipio, direccion,profesion) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)

	}

	var lastInsertId int
	err = stmt.QueryRow(nombre, area, otro, puesto, sede, transporte, proyecto, departamento, municipio, direccion, profesion).Scan(&lastInsertId)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
	}

	horario := r.Form["dias"]
	if len(horario) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO Transporte_dia (IDR,Dias) VALUES (@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta horario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range horario {
			_, err = stmt.Exec(lastInsertId, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta horario", http.StatusInternalServerError)
				return
			}
		}
	}

	hijos := r.Form["hijo"]
	fechas := r.Form["fechanac"]

	if len(hijos) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO Transporte_fechaNac (IDR,Nombre, fechaNac) VALUES (@P1,@P2, @P3)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta hijo,fecha", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for i := range hijos {
			if hijos[i] != "" {
				_, err = stmt.Exec(lastInsertId, hijos[i], fechas[i])
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("Public/Html/Respuesta.html")
	// Renderizar plantilla HTML con opciones de preguntas select
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func formulariomdvpsicologico(w http.ResponseWriter, Cfg *Controller.Config) {
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	municipio, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM municipalities where fkCodeState='CP1101' order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estado_civ, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_estadociv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Respuestas:   respuestas,
		Estado_civ:   estado_civ,
		Departamento: municipio,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/mdv/psicologico.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postmdvpsicologico(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	// Create connection pool
	colonia := cambio(r.FormValue("colonia"))
	otracolonia := r.FormValue("otraColonia")
	nombre := r.FormValue("nombre")
	diplomado := r.FormValue("diplomado")
	fechanac := r.FormValue("fechanac")
	estadoCivil := r.FormValue("estadoCivil")
	motivo := r.FormValue("motivo")
	otraOcupacion := r.FormValue("otraOcupacion")
	convivencia := r.FormValue("convivencia")
	considera := r.FormValue("considera")
	otraSituacion := r.FormValue("otraSituacion")
	situacionesMas := r.FormValue("situacionesmas")
	psicologico := r.FormValue("psicologico")
	otroPsicologico := r.FormValue("otroPsicologico")
	municipio := r.FormValue("municipio")
	idvn := r.FormValue("idvn")
	hijos := cambio(r.FormValue("hijos"))
	ludoteca := cambio(r.FormValue("ludoteca"))
	hijosludoteca := cambio(r.FormValue("hijosludoteca"))
	papa := r.FormValue("papa")
	sede := r.FormValue("sede")

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error conexion", http.StatusInternalServerError)

	}

	stmt, err := tx.Prepare("INSERT INTO RespuestasPsicosocial(sede,nombre, diplomado, fechanac, estadoCivil, motivo, otraOcupacion, convivencia, considera, otraSituacion, situacionesMas, psicologico, otroPsicologico, idvn,municipio,papa,hijos,ludoteca,hijosludoteca,colonia,otracolonia) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14,@P15,@P16,@P17,@P18,@P19,@P20,@P21)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)

	}

	var lastInsertId int
	err = stmt.QueryRow(sede, nombre, diplomado, fechanac, estadoCivil, motivo, otraOcupacion, convivencia, considera, otraSituacion, situacionesMas, psicologico, otroPsicologico, idvn, municipio, papa, hijos, ludoteca, hijosludoteca, colonia, otracolonia).Scan(&lastInsertId)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
	}

	ocupacion := r.Form["ocupacion"]
	if len(ocupacion) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO RespMutiNumeros (idPregunta,idp,idRespuesta) VALUES (1,@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta ocupacion", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range ocupacion {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					tx.Rollback()
					return

				}
			}
		}
	}

	situaciones := r.Form["situaciones"]
	if len(situaciones) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO RespMutiNumeros (idPregunta,idp,idRespuesta) VALUES (2,@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta situaciones", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range situaciones {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					tx.Rollback()
					return

				}
			}
		}
	}

	parentescomas := r.Form["parentescomas"]
	if len(parentescomas) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO RespMutiNumeros (idPregunta,idp,idRespuesta) VALUES (3,@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta situaciones", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range parentescomas {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					tx.Rollback()
					return

				}
			}
		}
	}

	parentescomenos := r.Form["parentescomenos"]
	if len(parentescomenos) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO RespMutiNumeros (idPregunta,idp,idRespuesta) VALUES (4,@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta situaciones", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range parentescomenos {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					tx.Rollback()
					return

				}
			}
		}
	}

	destaca := r.Form["destaca"]
	if len(ocupacion) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO RespMutiLetras (idPregunta,idp,respuesta) VALUES (5,@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta situaciones", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range destaca {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					tx.Rollback()
					return

				}
			}
		}
	}

	positiva := r.Form["positiva"]
	if len(positiva) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO RespMutiLetras (idPregunta,idp,respuesta) VALUES (6,@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta situaciones", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range positiva {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					tx.Rollback()
					return

				}
			}
		}
	}
	mejorar := r.Form["mejorar"]
	if len(mejorar) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO RespMutiLetras (idPregunta,idp,respuesta) VALUES (7,@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta situaciones", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range mejorar {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					tx.Rollback()
					return

				}
			}
		}
	}

	ida := r.Form["ida"]
	if len(ida) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO RespMutiLetras (idPregunta,idp,respuesta) VALUES (8,@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta situaciones", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range ida {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					tx.Rollback()
					return

				}
			}
		}
	}
	idv := r.Form["idv"]
	if len(idv) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO RespMutiLetras (idPregunta,idp,respuesta) VALUES (9,@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta situaciones", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range idv {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					tx.Rollback()
					return

				}
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("Public/Html/Respuesta.html")
	// Renderizar plantilla HTML con opciones de preguntas select
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func Finanzas_datos(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database4)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	pais, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct id, name FROM [DATAGW].[dbo].[vw_statesgw] where country_id=7 order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type ActividadAE struct {
		ID     int
		Nombre string
	}

	type SubcategoriaAE struct {
		ID        int
		Nombre    string
		Actividad []ActividadAE
	}

	type AE struct {
		ID           int
		Nombre       string
		Subcategoria []SubcategoriaAE
	}

	stmt, err := db.Prepare(`
    SELECT c.id, c.nombre, sc.id, sc.nombre, a.id,Concat(a.[codigo],' - ',a.[nombre])
    FROM [FINANZAS].[dbo].CategoriaAE c
    LEFT JOIN [FINANZAS].[dbo].SubCategoriaAE sc ON c.id = sc.idCategoria
    LEFT JOIN [FINANZAS].[dbo].ActEconomica a ON sc.id = a.idSubcategoria
    WHERE c.deleted IS NULL AND sc.deleted IS NULL AND a.deleted IS NULL
`)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()
	// ejecutar procedimiento almacenado
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	// Continuación del código existente
	var categoriaMap = make(map[int]*AE)
	var categoria []AE

	for rows.Next() {
		var idCategoria int
		var nombreCategoria string
		var idSubcategoria int
		var nombreSubcategoria string
		var idActividad int
		var nombreActividad string
		err := rows.Scan(&idCategoria, &nombreCategoria, &idSubcategoria, &nombreSubcategoria, &idActividad, &nombreActividad)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Verificar si la categoría ya existe
		if _, ok := categoriaMap[idCategoria]; !ok {
			categoriaMap[idCategoria] = &AE{
				ID:           idCategoria,
				Nombre:       nombreCategoria,
				Subcategoria: []SubcategoriaAE{},
			}
		}

		// Buscar subcategoría
		var subcategoriaExistente *SubcategoriaAE
		for i, sub := range categoriaMap[idCategoria].Subcategoria {
			if sub.ID == idSubcategoria {
				subcategoriaExistente = &categoriaMap[idCategoria].Subcategoria[i]
				break
			}
		}

		// Si la subcategoría no existe, añadirla
		if subcategoriaExistente == nil {
			subcategoriaExistente = &SubcategoriaAE{
				ID:        idSubcategoria,
				Nombre:    nombreSubcategoria,
				Actividad: []ActividadAE{},
			}
			categoriaMap[idCategoria].Subcategoria = append(categoriaMap[idCategoria].Subcategoria, *subcategoriaExistente)
		}

		// Añadir actividad a la subcategoría
		actividad := ActividadAE{
			ID:     idActividad,
			Nombre: nombreActividad,
		}
		subcategoriaExistente.Actividad = append(subcategoriaExistente.Actividad, actividad)
	}

	// Convertir el mapa a una lista
	for _, cat := range categoriaMap {
		categoria = append(categoria, *cat)
	}

	// Ahora `categoria` contiene todas las categorías, subcategorías y actividades organizadas.
	data := FormData{
		Paises: pais,
	}

	datos := struct {
		Data      FormData
		Categoria []AE
	}{
		Data:      data,
		Categoria: categoria,
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Finanzas/Proveedores.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, datos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Finanzas_datos_clientes(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database4)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	pais, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct id, name FROM [DATAGW].[dbo].[vw_statesgw] where country_id=7 order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type ActividadAE struct {
		ID     int
		Nombre string
	}

	type SubcategoriaAE struct {
		ID        int
		Nombre    string
		Actividad []ActividadAE
	}

	type AE struct {
		ID           int
		Nombre       string
		Subcategoria []SubcategoriaAE
	}

	stmt, err := db.Prepare(`
    SELECT c.id, c.nombre, sc.id, sc.nombre, a.id,Concat(a.[codigo],' - ',a.[nombre])
    FROM [FINANZAS].[dbo].CategoriaAE c
    LEFT JOIN [FINANZAS].[dbo].SubCategoriaAE sc ON c.id = sc.idCategoria
    LEFT JOIN [FINANZAS].[dbo].ActEconomica a ON sc.id = a.idSubcategoria
    WHERE c.deleted IS NULL AND sc.deleted IS NULL AND a.deleted IS NULL
`)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()
	// ejecutar procedimiento almacenado
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	// Continuación del código existente
	var categoriaMap = make(map[int]*AE)
	var categoria []AE

	for rows.Next() {
		var idCategoria int
		var nombreCategoria string
		var idSubcategoria int
		var nombreSubcategoria string
		var idActividad int
		var nombreActividad string
		err := rows.Scan(&idCategoria, &nombreCategoria, &idSubcategoria, &nombreSubcategoria, &idActividad, &nombreActividad)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Verificar si la categoría ya existe
		if _, ok := categoriaMap[idCategoria]; !ok {
			categoriaMap[idCategoria] = &AE{
				ID:           idCategoria,
				Nombre:       nombreCategoria,
				Subcategoria: []SubcategoriaAE{},
			}
		}

		// Buscar subcategoría
		var subcategoriaExistente *SubcategoriaAE
		for i, sub := range categoriaMap[idCategoria].Subcategoria {
			if sub.ID == idSubcategoria {
				subcategoriaExistente = &categoriaMap[idCategoria].Subcategoria[i]
				break
			}
		}

		// Si la subcategoría no existe, añadirla
		if subcategoriaExistente == nil {
			subcategoriaExistente = &SubcategoriaAE{
				ID:        idSubcategoria,
				Nombre:    nombreSubcategoria,
				Actividad: []ActividadAE{},
			}
			categoriaMap[idCategoria].Subcategoria = append(categoriaMap[idCategoria].Subcategoria, *subcategoriaExistente)
		}

		// Añadir actividad a la subcategoría
		actividad := ActividadAE{
			ID:     idActividad,
			Nombre: nombreActividad,
		}
		subcategoriaExistente.Actividad = append(subcategoriaExistente.Actividad, actividad)
	}

	// Convertir el mapa a una lista
	for _, cat := range categoriaMap {
		categoria = append(categoria, *cat)
	}

	// Ahora `categoria` contiene todas las categorías, subcategorías y actividades organizadas.
	data := FormData{
		Paises: pais,
	}

	datos := struct {
		Data      FormData
		Categoria []AE
	}{
		Data:      data,
		Categoria: categoria,
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Finanzas/Clientes.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, datos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func obtenerdepGeneral(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	pais := r.FormValue("pais")
	var opcionesdep2 string
	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database4)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	rows, err := db.Query("SELECT Distinct id, name FROM [DATAGW].[dbo].[vw_statesgw] where country_id=@p1 order by name", pais)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionesdep2 += "<option value=''></option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesdep2 += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionesdep2)
}

func obtenermunGeneral(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	departamento := r.FormValue("departamento")
	var opcionesdep2 string
	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database4)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	rows, err := db.Query("SELECT Distinct id, name FROM [DATAGW].[dbo].[vw_municipalities] where state_id=@p1 order by name", departamento)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionesdep2 += "<option value=''></option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesdep2 += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionesdep2)
}

func PostFinanzas_datos(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	nombres := r.FormValue("nombres")
	apellidos := r.FormValue("apellidos")
	razonSocial := r.FormValue("razonSocial")
	nombreComercial := r.FormValue("nombreComercial")
	actividadEconomica := r.FormValue("actividadEconomica")
	actividadEconomica2 := r.FormValue("actividadEconomica2")
	actividadEconomica3 := r.FormValue("actividadEconomica3")
	tipoPersona := r.FormValue("tipoPersona")
	NIT := r.FormValue("NIT")
	nombreRepresentante := r.FormValue("nombreRepre")
	NRC := r.FormValue("NRC")
	identidad := r.FormValue("identidad")
	tipoContribuyente := r.FormValue("tipoContribuyente")
	domiciliado := r.FormValue("domiciliado")
	direccion := r.FormValue("direccion")
	pais := r.FormValue("pais")
	departamento := r.FormValue("departamento")
	municipio := r.FormValue("municipio")
	telefono1 := r.FormValue("telefono1")
	telefono2 := r.FormValue("telefono2")
	correo := r.FormValue("correo")
	personaContacto := r.FormValue("personaContacto")
	urlNITa := ""
	urlNITb := ""
	urlNRCa := ""
	urlNRCb := ""
	urlDUIa := ""
	urlDUIb := ""

	pago := r.FormValue("pago")
	banco := r.FormValue("banco")
	tipocuenta := r.FormValue("tipocuenta")
	cuenta := r.FormValue("cuenta")
	activo := 1
	tipo := r.FormValue("tipo")
	urlCuenta := ""

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database4)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	var nombrePais string
	err = db.QueryRow("SELECT name FROM [DATAGW].[dbo].[countriesgw] where id=@p1", pais).Scan(&nombrePais)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var db2 *sql.DB
	connString2 := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database5)
	var err2 error
	// Create connection pool
	db2, err2 = sql.Open("sqlserver", connString2)
	if err2 != nil {
		log.Fatal("Error creating connection pool: ", err2.Error())
	}

	ctx2 := context.Background()
	err = db2.PingContext(ctx2)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db2.Close()

	tx, err := db2.Begin()
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error conexion", http.StatusInternalServerError)
		return
	}

	stmt, err := tx.Prepare("INSERT INTO [FINANZAS].[dbo].[Proveedores] (tipo,nombres, apellidos, razonSocial, nombreComercial, actividadEconomica, tipoPersona, NIT, nombreRepresentante, NRC, identidad, tipoContribuyente, domiciliado, direccion, pais, departamento, municipio, telefono1, telefono2, correo, personaContacto,activo,actividadEconomica2,actividadEconomica3,pago,banco,tipocuenta,cuenta) OUTPUT INSERTED.ID VALUES (@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14,@P15,@P16,@P17,@P18,@P19,@P20,@P21,@p22,@p23,@p24,@p25,@p26,@p27,@p28)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}

	var lastInsertId int
	err = stmt.QueryRow(tipo, nombres, apellidos, razonSocial, nombreComercial, actividadEconomica, tipoPersona, NIT, nombreRepresentante, NRC, identidad, tipoContribuyente, domiciliado, direccion, pais, departamento, municipio, telefono1, telefono2, correo, personaContacto, activo, actividadEconomica2, actividadEconomica3, pago, banco, tipocuenta, cuenta).Scan(&lastInsertId)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
	handleError(err)
	// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
	client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
	handleError(err)

	file, fileHandler, err := r.FormFile("identidadF")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename
		fileExt := filepath.Ext(fileName)
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "finanzas", nombrePais+"/identidadf/"+strconv.Itoa(int(lastInsertId))+fileExt, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urlDUIa = strconv.Itoa(int(lastInsertId)) + fileExt
	}

	file2, fileHandler2, err := r.FormFile("identidadA")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file2.Close()

		// Obtiene el nombre del archivo cargado
		fileName2 := fileHandler2.Filename
		fileExt2 := filepath.Ext(fileName2)
		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "finanzas", nombrePais+"/identidadA/"+strconv.Itoa(int(lastInsertId))+fileExt2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urlDUIb = strconv.Itoa(int(lastInsertId)) + fileExt2
	}

	file3, fileHandler3, err := r.FormFile("NRCF")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file3.Close()

		// Obtiene el nombre del archivo cargado
		fileName3 := fileHandler3.Filename
		fileExt3 := filepath.Ext(fileName3)

		fileBytes3, err := io.ReadAll(file3)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile3, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile3.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile3.Write(fileBytes3); err != nil {
			handleError(err)
		}
		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "finanzas", nombrePais+"/NRCF/"+strconv.Itoa(int(lastInsertId))+fileExt3, tempFile3,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urlNRCa = strconv.Itoa(int(lastInsertId)) + fileExt3
	}

	file4, fileHandler4, err := r.FormFile("NRCA")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file4.Close()

		// Obtiene el nombre del archivo cargado
		fileName4 := fileHandler4.Filename
		fileExt4 := filepath.Ext(fileName4)
		fileBytes4, err := io.ReadAll(file4)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile4, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile4.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile4.Write(fileBytes4); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "finanzas", nombrePais+"/NRCA/"+strconv.Itoa(int(lastInsertId))+fileExt4, tempFile4,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urlNRCb = strconv.Itoa(int(lastInsertId)) + fileExt4
	}

	file5, fileHandler5, err := r.FormFile("NITF")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file5.Close()

		// Obtiene el nombre del archivo cargado
		fileName5 := fileHandler5.Filename
		fileExt5 := filepath.Ext(fileName5)

		fileBytes5, err := io.ReadAll(file5)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile5, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile5.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile5.Write(fileBytes5); err != nil {
			handleError(err)
		}
		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "finanzas", nombrePais+"/NITF/"+strconv.Itoa(int(lastInsertId))+fileExt5, tempFile5,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urlNITa = strconv.Itoa(int(lastInsertId)) + fileExt5
	}

	file6, fileHandler6, err := r.FormFile("NITA")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file6.Close()

		// Obtiene el nombre del archivo cargado
		fileName6 := fileHandler6.Filename
		fileExt6 := filepath.Ext(fileName6)

		fileBytes6, err := io.ReadAll(file6)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile6, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile6.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile6.Write(fileBytes6); err != nil {
			handleError(err)
		}
		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "finanzas", nombrePais+"/NITA/"+strconv.Itoa(int(lastInsertId))+fileExt6, tempFile6,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urlNITb = strconv.Itoa(int(lastInsertId)) + fileExt6
	}

	file7, fileHandler7, err := r.FormFile("capturacuenta")
	if err != http.ErrMissingFile {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file7.Close()

		// Obtiene el nombre del archivo cargado
		fileName7 := fileHandler7.Filename
		fileExt7 := filepath.Ext(fileName7)

		fileBytes7, err := io.ReadAll(file7)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile7, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile7.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile7.Write(fileBytes7); err != nil {
			handleError(err)
		}
		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "finanzas", nombrePais+"/Cuenta/"+strconv.Itoa(int(lastInsertId))+fileExt7, tempFile7,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urlCuenta = strconv.Itoa(int(lastInsertId)) + fileExt7
	}
	// Insertar los valores en la tabla correspondiente
	stmt4, err := tx.Prepare("UPDATE Proveedores SET urlNITa = @P1,urlNITb= @P2,urlDUIa= @P3,urlDUIb= @P4, urlNRCa= @P5, urlNRCb= @P6,urlCuenta=@p7 WHERE id = @P8")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar update", http.StatusInternalServerError)
		return
	}
	defer stmt4.Close()

	_, err = stmt4.Exec(urlNITa, urlNITb, urlDUIa, urlDUIb, urlNRCa, urlNRCb, urlCuenta, lastInsertId)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar ejecutar update", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	fmt.Fprint(w, 1)
}

func RRHH_inventario(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database4)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	pais, err := Controller.GetOpcionesPregunta(db, "SELECT id, dispositivos FROM dispositivos order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Paises: pais,
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/RH/inventario.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Nnasv(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM VSedesSalud where IdPais=7 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	perfil := Controller.Option{
		Value: "7",
		Label: "sv",
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Sexo:         sexo,
		Parentesco:   parentesco,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       []Controller.Option{perfil},
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/NNsv.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Nnagt(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=4")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM VSedesSalud where IdPais=4 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	perfil := Controller.Option{
		Value: "7",
		Label: "sv",
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Sexo:         sexo,
		Parentesco:   parentesco,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       []Controller.Option{perfil},
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/NNAgt.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Nnahn(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=3")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM VSedesSalud where IdPais=3 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	perfil := Controller.Option{
		Value: "7",
		Label: "sv",
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Sexo:         sexo,
		Parentesco:   parentesco,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       []Controller.Option{perfil},
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/NNAhn.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Nnapn(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM VSedesSalud where IdPais=1 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	perfil := Controller.Option{
		Value: "7",
		Label: "sv",
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Sexo:         sexo,
		Parentesco:   parentesco,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       []Controller.Option{perfil},
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/NNApn.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Nnacr(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=6")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM VSedesSalud where IdPais=6 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	perfil := Controller.Option{
		Value: "7",
		Label: "sv",
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Sexo:         sexo,
		Parentesco:   parentesco,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       []Controller.Option{perfil},
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/NNAcr.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Nnacol(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=16")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM VSedesSalud where IdPais=16 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	perfil := Controller.Option{
		Value: "7",
		Label: "sv",
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Sexo:         sexo,
		Parentesco:   parentesco,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       []Controller.Option{perfil},
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/NNAcol.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Nnamx(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=17")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.sede_areas where Ida=5 and IdPais=17 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Sexo:         sexo,
		Parentesco:   parentesco,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/NNAmx.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostNnasalud(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var idF int
	// Obtener los valores de los campos del formulario
	programa := cambio(r.FormValue("programa"))
	if programa == 0 {
		idF = 2
	} else {
		idF = 3
	}
	pais := r.FormValue("idpais")
	idSede := cambio(r.FormValue("sede2"))
	fechaNac := r.FormValue("FechaN")
	nombres := r.FormValue("nombre")
	apellidos := r.FormValue("apellido")
	sexo := cambio(r.FormValue("Sexo"))
	nacionalidad := cambio(r.FormValue("nacionalidad"))
	departamento := cambio(r.FormValue("pais"))
	municipio := cambio(r.FormValue("departamento"))
	discapacidad := cambio(r.FormValue("dis"))
	estudia := cambio(r.FormValue("Eactual"))
	estudioAlcanzado := cambio(r.FormValue("UT"))
	grado := cambio(r.FormValue("GA"))
	turno := cambio(r.FormValue("Turno"))
	seccion := cambio(r.FormValue("seccion"))
	nuevogw := cambio(r.FormValue("anterior"))
	autovozimg := cambio(r.FormValue("autm3"))
	tipoParticipante := 1

	nombrecompleto := r.FormValue("contacto")
	telefono := r.FormValue("Telcontacto")
	parentesco := cambio(r.FormValue("relacion"))
	identidad := r.FormValue("DUIM")
	vozimg := cambio(r.FormValue("autR3"))
	confirmo1 := cambio(r.FormValue("autR4"))
	confirmo2 := cambio(r.FormValue("autR5"))
	confirmo3 := cambio(r.FormValue("autR6"))
	confirmo4 := cambio(r.FormValue("autR7"))
	confirmo5 := cambio(r.FormValue("autR8"))
	FechaNacRes := r.FormValue("FechaNR")
	fechaficha := r.FormValue("fechaficha")

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	stmt, err := db.Prepare("IF EXISTS(SELECT * FROM FormularioRegisto WHERE nombres = @P1 and apellidos = @P2 and identidad = @P3 and pais = @P4 and fechaNAc = @P5 and idF = @P6 and deleted is null)Begin select 1 end else Begin select 0 end")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	row := stmt.QueryRow(nombres, apellidos, identidad, pais, fechaNac, idF)

	var result int
	_ = row.Scan(&result)

	type Datos struct {
		Nombre   string
		Apellido string
		Id       int
		Programa string
	}

	if result == 1 {

		datos := Datos{
			Nombre:   nombres,
			Apellido: apellidos,
		}

		t, err := template.ParseFiles("Public/Html/RespuestaDuplicada.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Renderizar la plantilla con los datos
		err = t.Execute(w, datos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		tx, err := db.Begin()
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error conexion", http.StatusInternalServerError)

		}

		stmt, err := tx.Prepare("INSERT INTO FormularioRegisto(idF,tipoParticipante,pais,idSede,fechaNac,nombres,apellidos,sexo,nacionalidad,identidad,discapacidad,estudia,estudioAlcanzado,grado,turno,seccion,nuevogw,autovozimg,departamento,municipio) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14,@P15,@P16,@P17,@P18,@P19,@P20)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		var lastInsertId int
		err = stmt.QueryRow(idF, tipoParticipante, pais, idSede, fechaNac, nombres, apellidos, sexo, nacionalidad, identidad, discapacidad, estudia, estudioAlcanzado, grado, turno, seccion, nuevogw, autovozimg, departamento, municipio).Scan(&lastInsertId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		stmt, err = tx.Prepare("INSERT INTO AutorizacionMenores(idFR,nombrecompleto, telefono, parentesco, identidad, vozimg, confirmo1, confirmo2, confirmo3, confirmo4, confirmo5, fechaficha, FechaNacRes) VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		_, err = stmt.Exec(lastInsertId, nombrecompleto, telefono, parentesco, identidad, vozimg, confirmo1, confirmo2, confirmo3, confirmo4, confirmo5, fechaficha, FechaNacRes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		discapacidadmultiple := r.Form["dis2"]
		if len(discapacidadmultiple) == 0 {
		} else {
			stmt, err := tx.Prepare("INSERT INTO MultipleDiscapacidad (idFR,iddiscapacidad) VALUES (@P1,@P2)")
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al preparar la consulta ocupacion", http.StatusInternalServerError)
				return
			}
			defer stmt.Close()

			for _, valor := range discapacidadmultiple {
				if valor != "" {
					_, err = stmt.Exec(lastInsertId, valor)
					if err != nil {
						tx.Rollback()
						http.Error(w, "Error al ejecutar la consulta situaciones", http.StatusInternalServerError)
						return
					}
				}
			}
		}

		err = tx.Commit()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var programa string
		switch idF {

		case 2:
			programa = "Club de niñas"

		case 3:
			programa = "Red de graduadas"
		}

		datos := Datos{
			Nombre:   nombres,
			Apellido: apellidos,
			Id:       lastInsertId,
			Programa: programa,
		}

		t, err := template.ParseFiles("Public/Html/RespuestaPersonalizada.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Renderizar la plantilla con los datos
		err = t.Execute(w, datos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func NnaCISFsv(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil := Controller.Option{
		Value: "7",
		Label: "sv",
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Sexo:         sexo,
		Parentesco:   parentesco,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Perfil:       []Controller.Option{perfil},
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/CISFsa.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostNnaCISF(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	var idF int = 5
	// Obtener los valores de los campos del formulario
	pais := r.FormValue("idpais")
	idSede := cambio(r.FormValue("sede2"))
	fechaNac := r.FormValue("FechaN")
	nombres := r.FormValue("nombre")
	apellidos := r.FormValue("apellido")
	sexo := cambio(r.FormValue("Sexo"))
	nacionalidad := cambio(r.FormValue("nacionalidad"))
	departamento := cambio(r.FormValue("pais"))
	municipio := cambio(r.FormValue("departamento"))
	discapacidad := cambio(r.FormValue("dis"))
	estudia := cambio(r.FormValue("Eactual"))
	estudioAlcanzado := cambio(r.FormValue("UT"))
	grado := cambio(r.FormValue("GA"))
	turno := cambio(r.FormValue("Turno"))
	seccion := cambio(r.FormValue("seccion"))
	nuevogw := cambio(r.FormValue("anterior"))
	tipoParticipante := 1

	nombrecompleto := r.FormValue("contacto")
	telefono := r.FormValue("Telcontacto")
	infoAdicional1 := r.FormValue("cargo")
	identidad := r.FormValue("DUIM")
	fechaficha := r.FormValue("fechaficha")

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	stmt, err := db.Prepare("IF EXISTS(SELECT * FROM FormularioRegisto WHERE nombres = @P1 and apellidos = @P2 and identidad = @P3 and pais = @P4 and fechaNAc = @P5 and idF = @P6 and deleted is null)Begin select 1 end else Begin select 0 end")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	row := stmt.QueryRow(nombres, apellidos, identidad, pais, fechaNac, idF)

	var result int
	_ = row.Scan(&result)

	type Datos struct {
		Nombre   string
		Apellido string
		Id       int
		Programa string
	}

	if result == 1 {

		datos := Datos{
			Nombre:   nombres,
			Apellido: apellidos,
		}

		t, err := template.ParseFiles("Public/Html/RespuestaDuplicada.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Renderizar la plantilla con los datos
		err = t.Execute(w, datos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		tx, err := db.Begin()
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error conexion", http.StatusInternalServerError)

		}

		stmt, err := tx.Prepare("INSERT INTO FormularioRegisto(idF,tipoParticipante,pais,idSede,fechaNac,nombres,apellidos,sexo,nacionalidad,identidad,discapacidad,estudia,estudioAlcanzado,grado,turno,seccion,nuevogw,departamento,municipio) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14,@P15,@P16,@P17,@P18,@P19)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		var lastInsertId int
		err = stmt.QueryRow(idF, tipoParticipante, pais, idSede, fechaNac, nombres, apellidos, sexo, nacionalidad, identidad, discapacidad, estudia, estudioAlcanzado, grado, turno, seccion, nuevogw, departamento, municipio).Scan(&lastInsertId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return

		}

		stmt, err = tx.Prepare("INSERT INTO AutorizacionMenores(idFR,nombrecompleto, telefono, identidad, fechaficha,infoAdicional1) VALUES(@P1,@P2,@P3,@P4,@P5,@P6)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta automenores", http.StatusInternalServerError)
		}

		_, err = stmt.Exec(lastInsertId, nombrecompleto, telefono, identidad, fechaficha, infoAdicional1)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al ejecutar la consulta automenores", http.StatusInternalServerError)
		}

		discapacidadmultiple := r.Form["dis2"]
		if len(discapacidadmultiple) == 0 {
		} else {
			stmt, err := tx.Prepare("INSERT INTO MultipleDiscapacidad (idFR,iddiscapacidad) VALUES (@P1,@P2)")
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al preparar la consulta discapacidad", http.StatusInternalServerError)
				return
			}
			defer stmt.Close()

			for _, valor := range discapacidadmultiple {
				if valor != "" {
					_, err = stmt.Exec(lastInsertId, valor)
					if err != nil {
						tx.Rollback()
						http.Error(w, "Error al ejecutar la consulta discapacidad", http.StatusInternalServerError)
						return
					}
				}
			}
		}

		err = tx.Commit()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var programa string = "Intervenciones en el Centro de Integración Social Femenino"

		datos := Datos{
			Nombre:   nombres,
			Apellido: apellidos,
			Id:       lastInsertId,
			Programa: programa,
		}

		t, err := template.ParseFiles("Public/Html/RespuestaPersonalizada.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Renderizar la plantilla con los datos
		err = t.Execute(w, datos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func obtenermensaje(w http.ResponseWriter, r *http.Request) {
	programa := r.FormValue("programa")
	var opcionesCiudades string
	switch programa {
	case "0":

		opcionesCiudades += ("<p style='text-align: justify;'>¡Buen día! Reciba un cordial saludo de la Fundación Crisálida Internacional/Glasswing International (en adelante Glasswing). El presente es un formulario necesario con el fin de solicitar su autorización para que la niña o adolescente (en adelante NNA) a su cargo, pueda participar en nuestro programa Club de Niñas que desarrollamos en el centro educativo o sede en el que la NNA a su cargo asiste. Todas las actividades son realizadas dentro del centro educativo o sede correspondiente en los días entre lunes y sábado, en horarios que son acordados con la administración de la sede, por lo que la NNA no requiere movilizarse hacia otro lugar. En caso de que se realicen actividades fuera del centro educativo, solicitaremos una autorización especial para la actividad y traslado de la NNA. La oferta de actividades dentro de nuestro programa, incluye áreas como: bienestar integral y emocional, relación con el entorno, relaciones saludables, toma de decisiones y proyecto de vida, entre otras. Las niñas y jóvenes participan de módulos de formación y actividades de integración comunitaria, para lo cual es necesario el respaldo de su tutelar. A continuación, se detalla la información sobre las autorizaciones solicitadas:</p>")

	case "1":
		opcionesCiudades += ("<p style='text-align: justify;'>¡Buen día! Reciba un cordial saludo de la Fundación Crisálida Internacional/Glasswing International (en adelante Glasswing). El presente es un formulario necesario con el fin de solicitar su autorización para que la niña o adolescente (en adelante NNA) a su cargo, pueda participar en nuestro programa Red de Graduadas que desarrollamos en el centro educativo o sede en la NNA a su cargo asiste. Todas las actividades son realizadas dentro del centro educativo o sede correspondiente en los días entre lunes y sábado, en horarios que son acordados con la administración de la sede, por lo que la NNA no requiere movilizarse hacia otro lugar. En caso de que se realicen actividades fuera del centro educativo, solicitaremos una autorización especial para la actividad y traslado de la NNA. La oferta de actividades dentro de nuestro programa, incluye áreas como: liderazgo comunitario, comunicación y oratoria, trabajo en equipo e identificación de necesidades comunitarias donde se proporcionarán herramientas para elaborar y gestionar una acción comunitaria. Las niñas y jóvenes participan de módulos de formación y actividades de integración comunitaria, para lo cual es necesario el respaldo de su tutelar. A continuación, se detalla la información sobre las autorizaciones solicitadas:</p>")
	}
	fmt.Fprint(w, opcionesCiudades)

}

func obtenerseguimiento(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	seguimiento := r.FormValue("seguimiento")
	participante := r.FormValue("participante")
	sede := r.FormValue("sede")
	var opcionesCiudades string
	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	switch seguimiento {
	case "1":
		opcionesCiudades += "<br><label for='subtipo' class='form-label' style='font-weight: bold;'>Seleccione el seguimiento general realizado:</label>"
		opcionesCiudades += "<select class='form-select' id='subtipo' name='subtipo' onchange='fsubtipo(this.value)' required>"
		opcionesCiudades += "<option value='' selected></option>"
		opcionesCiudades += "<option value='1'>Seguimiento a la participación.</option>"
		opcionesCiudades += "<option value='2'>Seguimiento a la inserción.</option>"
		opcionesCiudades += "</select>"
		opcionesCiudades += "<div id='subtipodiv2'></div>"
	case "2":
		opcionesCiudades += "<br><label for='subtipo' class='form-label' style='font-weight: bold;'>¿Cuál fue el tipo de intervención realizada?</label>"
		opcionesCiudades += "<select class='form-select' id='subtipo' name='subtipo'   onchange='fsubtipo(this.value)' required>"
		opcionesCiudades += "<option value='' selected></option>"
		opcionesCiudades += "<option value='3'>Primeros auxilios psicológicos.</option>"
		opcionesCiudades += "<option value='4'>Seguimiento a salud física de participante.</option>"
		opcionesCiudades += "<option value='5'>Consejería/Escucha activa</option>"
		opcionesCiudades += "</select>"
		opcionesCiudades += "<div id='subtipodiv2'></div>"

	case "3":
		stmp, err := db.Prepare("IF EXISTS (SELECT * FROM [GWFORMS].[dbo].[SeguimientoJuventud] where tipo=3 and idp=@p1 and deleted is null) SELECT 1 AS Resultado;ELSE SELECT 0 AS Resultado;")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		rows, err := stmp.Query(participante)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var resultado int
		for rows.Next() {
			err := rows.Scan(&resultado)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		switch resultado {
		case 0:
			opcionesCiudades += "<br><label for='doc' class='form-label' style='font-weight: bold;'>Adjunte el plan de hábitos:</label>"
			opcionesCiudades += "<br><input type='file' class='form-control-file' id='doc' name='doc' required><br>"
			if sede != "1128" {
				opcionesCiudades += "<br><label for='habito' class='form-label' style='font-weight: bold;'>Escriba el hábito que se propuso la o él participante:</label>"
				opcionesCiudades += " <input type='text' name='habito' class='form-control' maxlength='500' required>"
				opcionesCiudades += "<br><label for='habito' class='form-label' style='font-weight: bold;'>De acuerdo al participante ¿Qué debo hacer o dejar de hacer (algo observable)?</label>"
				opcionesCiudades += " <input type='text' name='compromiso' class='form-control' maxlength='500' required>"
				opcionesCiudades += " <input type='text' name='compromiso' class='form-control' maxlength='500'>"
				opcionesCiudades += " <input type='text' name='compromiso' class='form-control' maxlength='500'>"
				opcionesCiudades += "<div id='subtipodiv2'></div>"
				opcionesCiudades += " <button class='btn btn-primary' onclick='agregarhabito(event)'>Agregar hábito</button><br>"
				opcionesCiudades += "<br><label style='font-weight: bold;' for='grupoapoyo' class='form-label'>Mi grupo de apoyo:</label><br>"
				opcionesCiudades += "<input type='text' class='form-control' name='grupoapoyo' required>"
				opcionesCiudades += "<div id='agregarapoyo'></div>"
				opcionesCiudades += "<button type='button' class='btn btn-primary' onclick='agregarapoyos(event)'>Agregar apoyo</button><br>"

			}
		case 1:
			if sede != "1128" {
				stmp, err := db.Prepare("Exec spVisualHabitosObjetivos @p1")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				rows, err := stmp.Query(participante)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				defer rows.Close()

				opcionesCiudades += "<br><table class='table table-bordered table-striped table-hover'>"
				opcionesCiudades += "<thead>"
				opcionesCiudades += "<tr>"
				opcionesCiudades += "<th scope='col'>Hábito</th>"
				opcionesCiudades += "<th scope='col'>Objetivo</th>"
				opcionesCiudades += "<th scope='col'>Logro</th>"
				opcionesCiudades += "</tr>"

				for rows.Next() {
					var idh, idho int
					var habito, objetivo string
					err := rows.Scan(&idh, &habito, &objetivo, &idho)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					opcionesCiudades += "<tr>"
					opcionesCiudades += "<td>" + habito + "</td>"
					opcionesCiudades += "<td>" + objetivo + "</td>"
					opcionesCiudades += "<td style='text-align: center;'><input type='checkbox' name='logro' value='" + strconv.Itoa(idho) + "' required></td>"
				}
				opcionesCiudades += "</table>"

				opcionesCiudades += "<br><label for='habito' class='form-label' style='font-weight: bold;'>Escriba el hábito que se propuso la o él participante</label>"
				opcionesCiudades += " <input type='text' name='habito' class='form-control' maxlength='500'>"
				opcionesCiudades += "<label for='habito' class='form-label' style='font-weight: bold;'>Que debo hacer o dejar de hacer (algo observable)</label>"
				opcionesCiudades += " <input type='text' name='compromiso' class='form-control' maxlength='500'>"
				opcionesCiudades += " <input type='text' name='compromiso' class='form-control' maxlength='500'>"
				opcionesCiudades += " <input type='text' name='compromiso' class='form-control' maxlength='500'>"
				opcionesCiudades += "<div id='subtipodiv2'></div>"
				opcionesCiudades += " <button class='btn btn-primary' onclick='agregarhabito(event)'>Agregar hábito</button><br>"
				opcionesCiudades += "<br><label style='font-weight: bold;' for='grupoapoyo' class='form-label'>Mi grupo de apoyo:</label><br>"
				opcionesCiudades += "<input type='text' class='form-control' name='grupoapoyo' required>"
				opcionesCiudades += "<div id='agregarapoyo'></div>"
				opcionesCiudades += "<button type='button' class='btn btn-primary' onclick='agregarapoyos(event)'>Agregar apoyo</button><br>"
			}
		}
	case "4":
		opcionesCiudades += "<br><label for='subtipo' class='form-label' style='font-weight: bold;'>¿Se ha cumplido el protocolo de deserción del programa Creando Profesionales?</label>"
		opcionesCiudades += "<select class='form-select' id='subtipo' name='subtipo' onchange='fsubtipo(this.value)' required>"
		opcionesCiudades += "<option value='' selected></option>"
		opcionesCiudades += "<option value='6'>Sí</option>"
		opcionesCiudades += "<option value='7'>No</option>"
		opcionesCiudades += "</select>"
		opcionesCiudades += "<div id='subtipodiv2'></div>"
	}
	fmt.Fprint(w, opcionesCiudades)
}

func obtenerseguimientosubtipo(w http.ResponseWriter, r *http.Request) {

	opcion := r.FormValue("subtipo")

	var opcionesCiudades string
	switch opcion {
	case "1":
		opcionesCiudades += "<br><label style='font-weight: bold;' for='subtipo' class='form-label'>¿Cuáles fueron los puntos abordados?</label><br>"
		opcionesCiudades += "<input type='text' class='form-control' name='abordados' required>"
		opcionesCiudades += "<div id='agregarpuntos'></div>"
		opcionesCiudades += "<button type='button' class='btn btn-primary' onclick='agregarpuntos(event)'>Agregar punto</button><br>"
		opcionesCiudades += "<br><label style='font-weight: bold;' for='subtipo' class='form-label'>¿Cuáles son los acuerdos de la sesión brindados por las y los participantes?</label><br>"
		opcionesCiudades += "<input type='text' class='form-control' name='acuerdos' required>"
		opcionesCiudades += "<div id='agregaracuerdo'></div>"
		opcionesCiudades += "<button type='button' class='btn btn-primary' onclick='agregaracuerdos(event)'>Agregar acuerdo</button><br>"
		opcionesCiudades += "<br><label for='subtipo' class='form-label' style='font-weight: bold;'>¿El seguimiento individual involucra a otra persona participante del programa o referente familiar?</label>"
		opcionesCiudades += "<select class='form-select' id='subtipo2' name='subtipo2' required>"
		opcionesCiudades += "<option value='' selected></option>"
		opcionesCiudades += "<option value='1'>Sí</option>"
		opcionesCiudades += "<option value='2'>No</option>"
		opcionesCiudades += "</select>"
		opcionesCiudades += "<br><label for='firma' class='form-label' style='font-weight: bold;'>Firma de la persona que brindó el seguimiento:</label><br>"
		opcionesCiudades += "<canvas id='Signature1draw' width='750px' height='250px' class='border'></canvas><br>"
		opcionesCiudades += "<button type='button' id='borrar1' class='btn btn-primary' onclick='clearSignature()'>Borrar firma</button><br>"
		opcionesCiudades += "<br><label for='firma2' class='form-label' style='font-weight: bold;'>Firma del participante:</label><br>"
		opcionesCiudades += "<canvas id='Signature2draw' width='750px' height='250px' class='border'></canvas><br>"
		opcionesCiudades += "<button type='button' id='borrar2' class='btn btn-primary' onclick='clearSignature2()'>Borrar firma</button>"
		opcionesCiudades += "<div id='e3div'></div>"
	case "2":
		opcionesCiudades += "<br><label style='font-weight: bold;' for='subtipo' class='form-label'>¿Cuáles fueron los puntos abordados?</label><br>"
		opcionesCiudades += "<input type='text' class='form-control' name='abordados' required>"
		opcionesCiudades += "<div id='agregarpuntos'></div>"
		opcionesCiudades += "<button type='button' class='btn btn-primary' onclick='agregarpuntos(event)'>Agregar punto</button><br>"
		opcionesCiudades += "<br><label style='font-weight: bold;' for='subtipo' class='form-label'>¿Cuáles son los acuerdos de la sesión brindados por las y los participantes?</label><br>"
		opcionesCiudades += "<input type='text' class='form-control' name='acuerdos' required>"
		opcionesCiudades += "<div id='agregaracuerdo'></div>"
		opcionesCiudades += "<button type='button' class='btn btn-primary' onclick='agregaracuerdos(event)'>Agregar acuerdo</button><br>"
		opcionesCiudades += "<br><label for='firma' class='form-label' style='font-weight: bold;'>Firma de la persona que brindó el seguimiento:</label><br>"
		opcionesCiudades += "<canvas id='Signature1draw' width='750px' height='250px' class='border'></canvas><br>"
		opcionesCiudades += "<button type='button' id='borrar1' class='btn btn-primary' onclick='clearSigniture()'>Borrar firma</button><br>"
		opcionesCiudades += "<br><label for='firma2' class='form-label' style='font-weight: bold;'>Firma del participante:</label><br>"
		opcionesCiudades += "<canvas id='Signature2draw' width='750px' height='250px' class='border'></canvas><br>"
		opcionesCiudades += "<button type='button' id='borrar2' class='btn btn-primary' onclick='clearSigniture2()'>Borrar firma</button>"
		opcionesCiudades += "<div id='e3div'></div>"
	case "3":
		opcionesCiudades += "<br><label style='font-weight: bold;' for='subtipo' class='form-label'>¿Qué aspectos individuales se abordaron durante la sesión?</label>"
		opcionesCiudades += "<select class='form-select' id='subtipo2' name='subtipo2' multiple required>"
		opcionesCiudades += "<option value='6'>Situaciones familiares.</option>"
		opcionesCiudades += "<option value='7'>Situaciones económicas.</option>"
		opcionesCiudades += "<option value='8'>Situaciones sentimentales.</option>"
		opcionesCiudades += "<option value='9'>Situaciones de salud física o mental.</option>"
		opcionesCiudades += "<option value='10'>Otra situación.</option>"
		opcionesCiudades += "</select>"
	case "5":
		opcionesCiudades += "<br><label style='font-weight: bold;' for='subtipo' class='form-label'>¿Qué aspectos individuales se abordaron durante la sesión?</label>"
		opcionesCiudades += "<select class='form-select' id='subtipo2' name='subtipo2' multiple required>"
		opcionesCiudades += "<option value='6'>Situaciones familiares.</option>"
		opcionesCiudades += "<option value='7'>Situaciones económicas.</option>"
		opcionesCiudades += "<option value='8'>Situaciones sentimentales.</option>"
		opcionesCiudades += "<option value='9'>Situaciones de salud física o mental.</option>"
		opcionesCiudades += "<option value='10'>Otra situación.</option>"
		opcionesCiudades += "</select>"
	case "6":
		opcionesCiudades += "<br><label for='doc' class='form-label' style='font-weight: bold;'>Adjunte acta de deserción:</label>"
		opcionesCiudades += "<br><input type='file' class='form-control-file' id='doc' name='doc' required><br>"
		opcionesCiudades += "<div id='agregarpuntos'></div>"
		opcionesCiudades += "<br><label style='font-weight: bold;' for='subtipo' class='form-label'>Describa las razones de la deserción</label><br>"
		opcionesCiudades += "<input type='text' class='form-control' name='acuerdos' required maxlength='1000'>"
		opcionesCiudades += "<div id='agregaracuerdo'></div>"
		opcionesCiudades += "<button type='button' class='btn btn-primary' onclick='agregaracuerdos(event)'>Agregar descripción</button><br>"
		opcionesCiudades += "<br><label for='firma' class='form-label' style='font-weight: bold;'>Firma de la persona que brindó el seguimiento:</label><br>"
		opcionesCiudades += "<canvas id='Signature1draw' width='750px' height='250px' class='border'></canvas><br>"
		opcionesCiudades += "<button type='button' id='borrar1' class='btn btn-primary' onclick='clearSigniture()'>Borrar firma</button><br>"
		opcionesCiudades += "<br><label for='firma2' class='form-label' style='font-weight: bold;'>Firma del participante:</label><br>"
		opcionesCiudades += "<canvas id='Signature2draw' width='750px' height='250px' class='border'></canvas><br>"
		opcionesCiudades += "<button type='button' id='borrar2' class='btn btn-primary' onclick='clearSigniture2()'>Borrar firma</button>"
		opcionesCiudades += "<div id='e3div'></div>"
	case "7":
		opcionesCiudades += "<br><label style='font-weight: bold;' for='subtipo' class='form-label'>¿Cuál o cuáles son las razones por las que  no se completó el acta de deserción?</label><br>"
		opcionesCiudades += "<input type='text' class='form-control' name='abordados' required maxlength='1000'>"
		opcionesCiudades += "<div id='agregarpuntos'></div>"
		opcionesCiudades += "<button type='button' class='btn btn-primary' onclick='agregarpuntos(event)'>Agregar razón</button><br>"
		opcionesCiudades += "<br><label style='font-weight: bold;' for='subtipo' class='form-label'>Describa las razones de la deserción</label><br>"
		opcionesCiudades += "<input type='text' class='form-control' name='acuerdos' required maxlength='1000'>"
		opcionesCiudades += "<div id='agregaracuerdo'></div>"
		opcionesCiudades += "<button type='button' class='btn btn-primary' onclick='agregaracuerdos(event)'>Agregar descripción</button><br>"
		opcionesCiudades += "<div id='e3div'></div>"
	}
	fmt.Fprint(w, opcionesCiudades)
}

func gestionmdv(w http.ResponseWriter) {

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/mdv/seguimiento.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func obtenerlistadomdv(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	sede := r.FormValue("sede")
	ano := r.FormValue("ano")
	cohorte := r.FormValue("cohorte")
	var fechainicio string
	var fechafin string
	switch cohorte {
	case "1":
		fechainicio = fmt.Sprintf("%s-01-01", ano)
		fechafin = fmt.Sprintf("%s-06-01", ano)
	case "2":
		fechainicio = fmt.Sprintf("%s-07-01", ano)
		fechafin = fmt.Sprintf("%s-12-01", ano)
	}

	var opcionesCiudades string
	var db *sql.DB

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=Rescate;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	rows, err := db.Query("exec mdv_Seguimiento @p1,@p2,@p3", sede, fechainicio, fechafin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionesCiudades += ("<option value=''></option>")
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesCiudades += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}

	fmt.Fprint(w, opcionesCiudades)

}

func postseguimientomdv(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	participante := r.FormValue("participante")
	fechasesion := r.FormValue("fechasesion")
	duracion := r.FormValue("duracion")
	tipo := r.FormValue("tipo")
	subtipo := r.FormValue("subtipo")
	seguimiento := r.FormValue("pro")
	tiposeguimiento := r.FormValue("ab")
	seguimientoexterno := r.FormValue("ext")
	comentarios := r.FormValue("comentarios")

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error conexion", http.StatusInternalServerError)
	}

	stmt, err := tx.Prepare("INSERT INTO SeguimientoJuventud (idp,fechaSesion,duracion,tipo,subtipo,seguimiento,tipoSeg,motivo,comentario) OUTPUT INSERTED.ID VALUES (@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	var idseguimiento int
	err = stmt.QueryRow(participante, fechasesion, duracion, tipo, subtipo, seguimiento, tiposeguimiento, seguimientoexterno, comentarios).Scan(&idseguimiento)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	abordados := r.Form["abordados"]
	if len(abordados) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO MultiplePuntosAbordados (idsj,puntosabordados) VALUES (@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta horario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range abordados {
			_, err = stmt.Exec(idseguimiento, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta horario", http.StatusInternalServerError)
				return
			}
		}
	}

	subtipo2 := r.Form["subtipo2"]
	if len(subtipo2) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO MultipleAspectos (idsj,[ida]) VALUES (@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta horario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range subtipo2 {
			_, err = stmt.Exec(idseguimiento, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta horario", http.StatusInternalServerError)
				return
			}
		}
	}

	acuerdos := r.Form["acuerdos"]
	if len(acuerdos) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO MultipleAcuerdos (idsj,acuerdo) VALUES (@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta horario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range acuerdos {
			_, err = stmt.Exec(idseguimiento, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta horario", http.StatusInternalServerError)
				return
			}
		}
	}
	counthabito := 0
	habito := r.Form["habito"]
	if len(habito) == 0 {
	} else {
		counthabito++
		stmt, err := tx.Prepare("INSERT INTO Habitos (idsj, habitos) OUTPUT inserted.ID VALUES (@p1, @p2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta horario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		idGeneradoMap := make(map[string]int64)
		for _, valor := range habito {
			var idGenerado int64
			err = stmt.QueryRow(idseguimiento, valor).Scan(&idGenerado)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta horario", http.StatusInternalServerError)
				return
			}
			idGeneradoMap[valor] = idGenerado
		}

		compromiso := r.Form["compromiso"]
		if len(compromiso) == 0 {
		} else {
			stmt, err := tx.Prepare("INSERT INTO HabitosObjetivo (idh, objetivo) VALUES (@p1, @p2)")
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al preparar la consulta horario", http.StatusInternalServerError)
				return
			}
			defer stmt.Close()

			for i, habitoKey := range habito {
				idGenerado := idGeneradoMap[habitoKey]
				for j := 0; j < 3; j++ {
					valor := compromiso[i*3+j]
					if valor != "" {
						_, err = stmt.Exec(idGenerado, valor)
						if err != nil {
							tx.Rollback()
							http.Error(w, "Error al ejecutar la consulta horario", http.StatusInternalServerError)
							return
						}
					}
				}
			}
		}
	}

	logro := r.Form["logro"]
	if len(logro) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO Logros (idsj,idho) VALUES (@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta horario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range logro {
			_, err = stmt.Exec(idseguimiento, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta horario", http.StatusInternalServerError)
				return
			}
		}
	}

	grupoapoyo := r.Form["grupoapoyo"]
	if len(grupoapoyo) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO MultipleApoyo (idsj,apoyo) VALUES (@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta horario", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range grupoapoyo {
			_, err = stmt.Exec(idseguimiento, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta horario", http.StatusInternalServerError)
				return
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
	handleError(err)
	// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
	client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
	handleError(err)

	file, fileHandler, err := r.FormFile("firma1")
	if err != nil && err != http.ErrMissingFile {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err != http.ErrMissingFile {
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename
		fileExt := filepath.Ext(fileName)
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "juventud", "seguimiento/sv/firmapsicologo/"+strconv.Itoa(idseguimiento)+fileExt, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		_, err = db.Exec("UPDATE SeguimientoJuventud SET urla = @p1 WHERE id = @p2", strconv.Itoa(idseguimiento)+fileExt, idseguimiento)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	file2, fileHandler2, err := r.FormFile("firma2")
	if err != nil && err != http.ErrMissingFile {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err != http.ErrMissingFile {
		defer file2.Close()

		// Obtiene el nombre del archivo cargado
		fileName2 := fileHandler2.Filename
		fileExt2 := filepath.Ext(fileName2)
		fileBytes2, err := io.ReadAll(file2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile2, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile2.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile2.Write(fileBytes2); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "juventud", "seguimiento/sv/firmaparticipante/"+strconv.Itoa(idseguimiento)+fileExt2, tempFile2,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		_, err = db.Exec("UPDATE SeguimientoJuventud SET urlb = @p1 WHERE id = @p2", strconv.Itoa(idseguimiento)+fileExt2, idseguimiento)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	file4, fileHandler4, err := r.FormFile("doc")
	if err != nil && err != http.ErrMissingFile {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err != http.ErrMissingFile {
		defer file4.Close()

		// Obtiene el nombre del archivo cargado
		fileName4 := fileHandler4.Filename
		fileExt4 := filepath.Ext(fileName4)
		fileBytes4, err := io.ReadAll(file4)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile4, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile4.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile4.Write(fileBytes4); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "juventud", "seguimiento/sv/documento/"+strconv.Itoa(idseguimiento)+fileExt4, tempFile4,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)

		_, err = db.Exec("UPDATE SeguimientoJuventud SET urldoc = @p1 WHERE id = @p2", strconv.Itoa(idseguimiento)+fileExt4, idseguimiento)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	respuesta := 1
	fmt.Fprint(w, respuesta)
}

func Obtenermunchemonics(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	departamento := r.FormValue("departamento")
	var opcionesmunsede string
	// Build connection string
	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()
	stmt, err := db.Prepare("Select distinct idMunicipio,Municipio from vwSedesChemonics where idDepartamento=@p1 order by Municipio")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := stmt.Query(departamento)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	opcionesmunsede += "<option value=''></option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesmunsede += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionesmunsede)
}

func Obtenersedechemonics(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	departamento := r.FormValue("municipio")
	var opcionesmunsede string
	// Build connection string
	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()
	stmt, err := db.Prepare("Select distinct idSede,Sede from vwSedesChemonics where idMunicipio=@p1 order by Sede")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := stmt.Query(departamento)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionesmunsede += "<option value=''></option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesmunsede += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionesmunsede)
}

func postEncuestaSMCHemonics(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	//NombreCampo en la BD: = r.FormValue("NombreCampo en el html")
	// Obtener los valores de los campos del formulario

	participa := r.FormValue("participar")
	idPais := cambio(r.FormValue("pais"))
	lugarTrabajo := r.FormValue("sede")
	edad := cambio(r.FormValue("edad"))
	anioLaboral := cambio(r.FormValue("aniol"))
	sexo := cambio(r.FormValue("sexo"))
	perfil := cambio(r.FormValue("rol"))
	rolTrabajo := r.FormValue("puesto")
	afir1 := cambio(r.FormValue("afirma1"))
	afir2 := cambio(r.FormValue("afirma2"))
	afir3 := cambio(r.FormValue("afirma3"))
	afir4 := cambio(r.FormValue("afirma4"))
	afir5 := cambio(r.FormValue("afirma5"))
	afir6 := cambio(r.FormValue("afirma6"))
	afir7 := cambio(r.FormValue("afirma7"))
	afir8 := cambio(r.FormValue("afirma8"))
	afir9 := cambio(r.FormValue("afirma9"))
	afir1b := cambio(r.FormValue("afirma1B"))
	afir2b := cambio(r.FormValue("afirma2B"))
	afir3b := cambio(r.FormValue("afirma3B"))
	afir4b := cambio(r.FormValue("afirma4B"))
	afir5b := cambio(r.FormValue("afirma5B"))
	afir6b := cambio(r.FormValue("afirma6B"))
	afir7b := cambio(r.FormValue("afirma7B"))
	afir8b := cambio(r.FormValue("afirma8B"))
	afir9b := cambio(r.FormValue("afirma9B"))
	afec1 := cambio(r.FormValue("afecta1"))
	afec2 := cambio(r.FormValue("afecta2"))
	afec3 := cambio(r.FormValue("afecta3"))
	actitud := cambio(r.FormValue("actitud"))
	estrategia := cambio(r.FormValue("estrategia"))
	apoyo := cambio(r.FormValue("apoyo"))
	cambioVida := r.FormValue("cambiov")
	motivo := r.FormValue("motivo")
	motivo2 := r.FormValue("motivo2")
	conoceTecnica := cambio(r.FormValue("conocet"))
	practicaTecnica := cambio(r.FormValue("practicat1"))
	autocuido := cambio(r.FormValue("practicat"))
	estEmocional := cambio(r.FormValue("estadoe"))
	nivAngustia := cambio(r.FormValue("angustia"))
	oportunidad := cambio(r.FormValue("op7"))
	climaLaboral := r.FormValue("clima")
	conoceEfectos := cambio(r.FormValue("conocee"))
	conscienteEfectos := cambio(r.FormValue("consciente"))
	beneficio := r.FormValue("beneficio")
	rel1 := cambio(r.FormValue("rel1"))
	rel2 := cambio(r.FormValue("rel2"))
	rel3 := cambio(r.FormValue("rel3"))
	rel4 := cambio(r.FormValue("rel4"))
	rel1b := cambio(r.FormValue("rel12"))
	rel2b := cambio(r.FormValue("rel22"))
	rel3b := cambio(r.FormValue("rel32"))
	rel4b := cambio(r.FormValue("rel42"))
	p81 := cambio(r.FormValue("81"))
	p82 := cambio(r.FormValue("82"))
	op9 := cambio(r.FormValue("op9"))
	Pid := cambio(r.FormValue("id"))
	version := 3

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente

	stmt, err := tx.Prepare("INSERT INTO Sanamente_ResEncuesta (participa,idPais,lugarTrabajo,edad,anioLaboral,sexo,perfil,rolTrabajo,afir1,afir2,afir3,afir4,afir5,afir6,afir7,afir8,afir9,afir1b,afir2b,afir3b,afir4b,afir5b,afir6b,afir7b,afir8b,afir9b,afec1,afec2,afec3,actitud,estrategia,apoyo,cambioVida,motivo,motivo2,conoceTecnica,practicaTecnica,estEmocional,nivAngustia,oportunidad,climaLaboral,conoceEfectos,conscienteEfectos,beneficio,rel1,rel2,rel3,rel4,rel1b,rel2b,rel3b,rel4b,p81,p82,op9,Pid,version_,autocuido) OUTPUT INSERTED.ID VALUES (@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14,@P15,@P16,@P17,@P18,@P19,@P20,@P21,@P22,@P23,@P24,@P25,@P26,@P27,@P28,@P29,@P30,@P31,@P32,@P33,@P34,@P35,@P36,@P37,@P38,@P39,@P40,@P41,@P42,@P43,@P44,@P45,@P46,@P47,@P48,@P49,@P50,@P51,@P52,@P53,@P54,@P55,@P56,@P57,@P58)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	var idGenerado int
	err = stmt.QueryRow(participa, idPais, lugarTrabajo, edad, anioLaboral, sexo, perfil, rolTrabajo, afir1, afir2, afir3, afir4, afir5, afir6, afir7, afir8, afir9, afir1b, afir2b, afir3b, afir4b, afir5b, afir6b, afir7b, afir8b, afir9b, afec1, afec2, afec3, actitud, estrategia, apoyo, cambioVida, motivo, motivo2, conoceTecnica, practicaTecnica, estEmocional, nivAngustia, oportunidad, climaLaboral, conoceEfectos, conscienteEfectos, beneficio, rel1, rel2, rel3, rel4, rel1b, rel2b, rel3b, rel4b, p81, p82, op9, Pid, version, autocuido).Scan(&idGenerado)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	r.ParseForm()

	detalle1 := r.Form["op10"]
	if len(detalle1) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_ResEncuestad (idp, idRespuesta) VALUES (@p1,@p2)")
		if err != nil {
			http.Error(w, "Error al preparar la consulta afectaciones"+err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return
		}
		defer stmt.Close()

		for _, valor := range detalle1 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta afectaciones", http.StatusInternalServerError)
				return
			}
		}
	}

	detalle2 := r.Form["epractica"]
	if len(detalle2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_ResEncuestad (idp, idRespuesta) VALUES (@p1,@p2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta estrategia practicada", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range detalle2 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta estrategia practicada", http.StatusInternalServerError)
				return
			}
		}
	}

	detalle3 := r.Form["opractica"]
	if len(detalle3) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_ResEncuestad (idp, idRespuesta) VALUES (@p1,@p2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta oportunidad de practica", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range detalle3 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta oportunidad de practica", http.StatusInternalServerError)
				return
			}
		}
	}

	detalle4 := r.Form["op101"]
	if len(detalle4) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO Sanamente_ResEncuestad (idp, idRespuesta) VALUES (@p1,@p2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta practica en el trabajo", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range detalle4 {
			_, err = stmt2.Exec(idGenerado, valor)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta practica en el trabajo", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	t, err := template.ParseFiles("public/HTML/RespuestaSanamente.html")
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

func EncuestaSMChemonics(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paises, err := Controller.GetOpcionesPregunta(db, "SELECT distinct id, name FROM admin_main_gwdata.states where fkCodeCountry='+503' order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	afectacion, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=2")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actitud, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=4")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estrategia, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=5")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	practica, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=6")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	estado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=7")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	afirmacion, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=8")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lugar, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_Encuesta where idPregunta=9")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	solucion, err := getOpcionesPreguntaNew(db, "SELECT distinct id, Respuesta, idPregunta FROM admin_main_gwdata.Respuestas_formularios where idFormulario=9")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		id = "0"
	}

	type CombinedData struct {
		Data  FormData
		Data2 FormDataNew
		Data3 string
	}

	data := FormData{
		Respuestas: respuestas,
		Paises:     paises,
		Sexo:       sexo,
		Afectacion: afectacion,
		Actitud:    actitud,
		Estrategia: estrategia,
		Practica:   practica,
		Estado:     estado,
		Afirmacion: afirmacion,
		Lugar:      lugar,
	}

	data2 := FormDataNew{
		Solucion: solucion,
	}

	// Crear la estructura combinada
	combinedData := CombinedData{
		Data:  data,
		Data2: data2,
		Data3: id,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/sv/Encuestachemonics.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, combinedData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func obtenerseguimientotipo(w http.ResponseWriter, r *http.Request) {
	seguimiento := r.FormValue("subtipo")
	var opcionesCiudades string
	switch seguimiento {
	case "1":
		opcionesCiudades += "<br><label for='ab' class='form-label' style='font-weight: bold;'>¿Qué tipo de seguimiento se brindará?</label>"
		opcionesCiudades += "<select class='form-select' id='ab' name='ab' onchange='fab(this.value)' required>"
		opcionesCiudades += "<option value='' selected></option>"
		opcionesCiudades += "<option value='1'>Seguimiento interno.</option>"
		opcionesCiudades += "<option value='2'>Seguimiento externo.</option>"
		opcionesCiudades += "</select>"
		opcionesCiudades += "<div id='ext'></div>"
	case "2":
		opcionesCiudades += "<br><label for='ab' class='form-label' style='font-weight: bold;'>¿Por qué?</label>"
		opcionesCiudades += "<select class='form-select' id='ab' name='ab' onchange='fab(this.value)'  required>"
		opcionesCiudades += "<option value='' selected></option>"
		opcionesCiudades += "<option value='3'>El joven no brindó consentimiento al seguimiento</option>"
		opcionesCiudades += "<option value='4'>El joven decidió abandonar el programa</option>"
		opcionesCiudades += "<option value='5'>Otra razón</option>"
		opcionesCiudades += "</select>"
		opcionesCiudades += "<div id='ext'></div>"
	}
	fmt.Fprint(w, opcionesCiudades)
}

func sanamentesvcopy(w http.ResponseWriter) {
	type Datos struct {
		Nombre   string
		Apellido string
		Id       int
		Programa string
	}
	nombre := "prueba"

	apellido := "prueba2"

	datos := Datos{
		Nombre:   nombre,
		Apellido: apellido,
	}

	t, err := template.ParseFiles("Public/Html/RespuestaDuplicadaSanamente.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Renderizar la plantilla con los datos
	err = t.Execute(w, datos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func obtenerdepsedeHandlersanamente2(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	area := r.FormValue("area")
	pais := r.FormValue("pais")
	var opcionesmunsede string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT Distinct Iddep,Departamento FROM VSedessanamente where categorizacion=? and IdPais=? order by Departamento", area, pais)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionesmunsede += "<option value=''></option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesmunsede += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionesmunsede)
}

func obtenermunsedeHandlersanamente2(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	depsede := r.FormValue("depsede")
	area := r.FormValue("area")
	var opcionesmunsede string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT Distinct Idmun,Municipio FROM VSedessanamente where Iddep=? and categorizacion=? order by Municipio", depsede, area)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionesmunsede += "<option value=''></option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesmunsede += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionesmunsede)
}

func obtenersedeHandlersanamente2(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	sede := r.FormValue("munsede")
	area := r.FormValue("area")
	var opcionessede string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT IdSede,Sede FROM VSedessanamente where (Idmun=? and categorizacion=?) or (2=? and ?=4 and IdSede IN(3436,3980)) order by Sede", sede, area, sede, area)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionessede += "<option value=''></option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionessede += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionessede)
}

func EncuestaSMInicial(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	// Build connection string
	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()
	stmt, err := db.Prepare("Select [id],[seccion],[texto] from PreguntasFormularios where [idF]=9")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := stmt.Query()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Preguntas struct {
		Id       int
		Seccion  int
		Pregunta string
	}

	var Datos []Preguntas

	for rows.Next() {

		var P Preguntas

		err := rows.Scan(&P.Id, &P.Seccion, &P.Pregunta)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		Datos = append(Datos, P)
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/gt/EncuestaInicialEFPEM.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, Datos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func EncuestaSMFinal(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	// Build connection string
	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()
	stmt, err := db.Prepare("Select [id],[seccion],[texto] from PreguntasFormularios where [idF]=9")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := stmt.Query()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Preguntas struct {
		Id       int
		Seccion  int
		Pregunta string
	}

	var Datos []Preguntas

	for rows.Next() {

		var P Preguntas

		err := rows.Scan(&P.Id, &P.Seccion, &P.Pregunta)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		Datos = append(Datos, P)
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/gt/EncuestaFinalEFPEM.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, Datos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func LiderazgoRedJuvenil(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/mdv/liderazgo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func PostLiderazgoRedJuvenil(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	nombrecompleto := r.FormValue("nombreCompleto")
	edad := r.FormValue("edad")
	depres := r.FormValue("departamento")
	telefono := r.FormValue("numeroCelular")
	pais := 7
	cohorte := 2
	correo := r.FormValue("correoPersonal")
	contactoemg := r.FormValue("telefonoEmergencia")
	referencia := r.FormValue("nombreEmergencia")
	organizacion := r.FormValue("organizacion")
	otraOrganizacion := r.FormValue("otraOrganizacion")
	urla := ""

	if organizacion == "-1" {
		organizacion = otraOrganizacion
	}

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error conexion", http.StatusInternalServerError)
	}

	stmt, err := tx.Prepare("INSERT INTO RegistroLiderazgo (nombrecompleto,edad,depres,telefono,pais,cohorte,correo,contactoemg,referencia,organizacion) OUTPUT INSERTED.ID VALUES (@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	var idseguimiento int
	err = stmt.QueryRow(nombrecompleto, edad, depres, telefono, pais, cohorte, correo, contactoemg, referencia, organizacion).Scan(&idseguimiento)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cred, err := azblob.NewSharedKeyCredential(Cfg.AccountName, Cfg.AccountKey)
	handleError(err)
	// The service URL for blob endpoints is usually in the form: http(s)://<account>.blob.core.windows.net/
	client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", Cfg.AccountName), cred, nil)
	handleError(err)

	file, fileHandler, err := r.FormFile("archivo1")
	if err != nil && err != http.ErrMissingFile {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err != http.ErrMissingFile {
		defer file.Close()

		// Obtiene el nombre del archivo cargado
		fileName := fileHandler.Filename
		fileExt := filepath.Ext(fileName)
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile, err := os.CreateTemp("", "upload-*.dat")
		if err != nil {
			handleError(err)
		}
		defer tempFile.Close()

		// Escribe los bytes del archivo en el archivo temporal
		if _, err := tempFile.Write(fileBytes); err != nil {
			handleError(err)
		}

		// Upload the file to a block blob
		_, err = client.UploadFile(context.TODO(), "juventud", "liderazgo/cartainteres/"+strconv.Itoa(idseguimiento)+fileExt, tempFile,
			&azblob.UploadFileOptions{
				BlockSize:   int64(1024),
				Concurrency: uint16(3),
				// If Progress is non-nil, this function is called periodically as bytes are uploaded.
				Progress: func(bytesTransferred int64) {

				},
			})

		handleError(err)
		urla = strconv.Itoa(idseguimiento) + fileExt
	}

	_, err = db.Exec("UPDATE RegistroLiderazgo SET urla = @p1 WHERE id = @p2", urla, idseguimiento)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/mdv/liderazgoRespuesta.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func postEncuestaSMInicial(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	//NombreCampo en la BD: = r.FormValue("NombreCampo en el html")
	// Obtener los valores de los campos del formulario

	identidad := r.FormValue("dpi")
	sexo := cambio(r.FormValue("sexo"))
	fechaNac := r.FormValue("fechaN")
	tiempolaboral := cambio(r.FormValue("aniol"))
	rolTrabajo := cambio(r.FormValue("puesto"))
	otrorol := r.FormValue("otrorol")
	tipodocente := cambio(r.FormValue("docente"))
	ntrabajo := cambio(r.FormValue("ntrabaja"))
	catedra := r.FormValue("catedra")
	ntrabajoEst := cambio(r.FormValue("ntrabajaEst"))
	p1_1 := cambio(r.FormValue("practica1"))
	p1_2 := cambio(r.FormValue("practica2"))
	p1_3 := cambio(r.FormValue("practica3"))
	p1_4 := cambio(r.FormValue("practica4"))
	p1_5 := cambio(r.FormValue("practica5"))
	p1_6 := cambio(r.FormValue("practica6"))
	p1_7 := cambio(r.FormValue("practica7"))
	p1_8 := cambio(r.FormValue("practica8"))
	p1_9 := cambio(r.FormValue("practica9"))
	p1_10 := cambio(r.FormValue("practica10"))
	p1_11 := cambio(r.FormValue("practica11"))
	p1_12 := cambio(r.FormValue("practica12"))
	p1_13 := cambio(r.FormValue("practica13"))
	p1_14 := cambio(r.FormValue("practica14"))
	p1_15 := cambio(r.FormValue("practica15"))
	p1_16 := cambio(r.FormValue("practica16"))
	p1_17 := cambio(r.FormValue("practica17"))
	p1_18 := cambio(r.FormValue("practica18"))
	p1_19 := cambio(r.FormValue("practica19"))
	p1_20 := cambio(r.FormValue("practica20"))
	p2_1 := cambio(r.FormValue("afirma21"))
	p2_2 := cambio(r.FormValue("afirma22"))
	p2_3 := cambio(r.FormValue("afirma23"))
	p2_4 := cambio(r.FormValue("afirma24"))
	p2_5 := cambio(r.FormValue("afirma25"))
	p2_6 := cambio(r.FormValue("afirma26"))
	p2_7 := cambio(r.FormValue("afirma27"))
	p3_1 := cambio(r.FormValue("afecta28"))
	p3_2 := cambio(r.FormValue("afecta29"))
	p3_3 := cambio(r.FormValue("afecta30"))
	p4 := cambio(r.FormValue("actitud"))
	p5_1 := cambio(r.FormValue("situacion31"))
	p5_2 := cambio(r.FormValue("situacion32"))
	p5_3 := cambio(r.FormValue("situacion33"))
	p5_4 := cambio(r.FormValue("situacion34"))
	p5_5 := cambio(r.FormValue("situacion35"))
	p5_6 := cambio(r.FormValue("situacion36"))
	p5_7 := cambio(r.FormValue("situacion37"))
	p6 := cambio(r.FormValue("angustia"))
	p7_1 := cambio(r.FormValue("vida38"))
	p7_2 := cambio(r.FormValue("vida39"))
	p7_3 := cambio(r.FormValue("vida40"))
	p7_4 := cambio(r.FormValue("vida41"))
	p7_5 := cambio(r.FormValue("vida42"))
	p7_6 := cambio(r.FormValue("vida43"))
	p7_7 := cambio(r.FormValue("vida44"))
	p7_8 := cambio(r.FormValue("vida45"))
	p7_9 := cambio(r.FormValue("vida46"))
	p8_1 := cambio(r.FormValue("estrategia47"))
	p8_2 := cambio(r.FormValue("estrategia48"))
	p8_3 := cambio(r.FormValue("estrategia49"))
	p8_4 := cambio(r.FormValue("estrategia50"))
	p8_5 := cambio(r.FormValue("estrategia51"))
	p8_6 := cambio(r.FormValue("estrategia52"))
	p8_7 := cambio(r.FormValue("estrategia53"))
	p8_8 := cambio(r.FormValue("estrategia54"))
	p8_9 := cambio(r.FormValue("estrategia55"))
	p8_10 := cambio(r.FormValue("estrategia56"))
	p8_11 := cambio(r.FormValue("estrategia57"))
	p8_12 := cambio(r.FormValue("estrategia58"))
	p8_13 := cambio(r.FormValue("estrategia59"))
	p8_14 := cambio(r.FormValue("estrategia60"))
	p8_15 := cambio(r.FormValue("estrategia61"))
	p8_16 := cambio(r.FormValue("estrategia62"))
	p8_17 := cambio(r.FormValue("estrategia63"))
	p8_18 := cambio(r.FormValue("estrategia64"))
	p8_19 := cambio(r.FormValue("estrategia65"))
	p8_20 := cambio(r.FormValue("estrategia66"))
	p8_21 := cambio(r.FormValue("estrategia67"))
	p9 := cambio(r.FormValue("realidad"))
	p10 := cambio(r.FormValue("efectos"))
	p11 := cambio(r.FormValue("nreferido"))
	p12_1 := cambio(r.FormValue("entorno68"))
	p12_2 := cambio(r.FormValue("entorno69"))
	p12_3 := cambio(r.FormValue("entorno70"))
	p12_4 := cambio(r.FormValue("entorno71"))
	p12_5 := cambio(r.FormValue("entorno72"))
	p12_6 := cambio(r.FormValue("entorno73"))
	p12_7 := cambio(r.FormValue("entorno74"))
	p12_8 := cambio(r.FormValue("entorno75"))
	p12_9 := cambio(r.FormValue("entorno76"))
	p12_10 := cambio(r.FormValue("entorno77"))
	p12_11 := cambio(r.FormValue("entorno78"))
	p12_12 := cambio(r.FormValue("entorno79"))
	p12_13 := cambio(r.FormValue("entorno80"))
	p13_1 := cambio(r.FormValue("rel1"))
	p13_2 := cambio(r.FormValue("rel2"))
	p13_3 := cambio(r.FormValue("rel3"))

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente

	stmt, err := tx.Prepare("INSERT INTO SanamenteEncuesta_EFPEM (identidad,sexo,fechaNac,tiempolaboral,rolTrabajo,otrorol,tipodocente,ntrabajo,catedra,ntrabajoEst,p1_1,p1_2,p1_3,p1_4,p1_5,p1_6,p1_7,p1_8,p1_9,p1_10,p1_11,p1_12,p1_13,p1_14,p1_15,p1_16,p1_17,p1_18,p1_19,p1_20,p2_1,p2_2,p2_3,p2_4,p2_5,p2_6,p2_7,p3_1,p3_2,p3_3,p4,p5_1,p5_2,p5_3,p5_4,p5_5,p5_6,p5_7,p6,p7_1,p7_2,p7_3,p7_4,p7_5,p7_6,p7_7,p7_8,p7_9,p8_1,p8_2,p8_3,p8_4,p8_5,p8_6,p8_7,p8_8,p8_9,p8_10,p8_11,p8_12,p8_13,p8_14,p8_15,p8_16,p8_17,p8_18,p8_19,p8_20,p8_21,p9,p10,p11,p12_1,p12_2,p12_3,p12_4,p12_5,p12_6,p12_7,p12_8,p12_9,p12_10,p12_11,p12_12,p12_13,p13_1,p13_2,p13_3) OUTPUT INSERTED.ID VALUES (@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14,@P15,@P16,@P17,@P18,@P19,@P20,@P21,@P22,@P23,@P24,@P25,@P26,@P27,@P28,@P29,@P30,@P31,@P32,@P33,@P34,@P35,@P36,@P37,@P38,@P39,@P40,@P41,@P42,@P43,@P44,@P45,@P46,@P47,@P48,@P49,@P50,@P51,@P52,@P53,@P54,@P55,@P56,@P57,@P58,@P59,@P60,@P61,@P62,@P63,@P64,@P65,@P66,@P67,@P68,@P69,@P70,@P71,@P72,@P73,@P74,@P75,@P76,@P77,@P78,@P79,@P80,@P81,@P82,@P83,@P84,@P85,@P86,@P87,@P88,@P89,@P90,@P91,@P92,@P93,@P94,@P95,@P96,@P97,@P98)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	var idGenerado int
	err = stmt.QueryRow(identidad, sexo, fechaNac, tiempolaboral, rolTrabajo, otrorol, tipodocente, ntrabajo, catedra, ntrabajoEst, p1_1, p1_2, p1_3, p1_4, p1_5, p1_6, p1_7, p1_8, p1_9, p1_10, p1_11, p1_12, p1_13, p1_14, p1_15, p1_16, p1_17, p1_18, p1_19, p1_20, p2_1, p2_2, p2_3, p2_4, p2_5, p2_6, p2_7, p3_1, p3_2, p3_3, p4, p5_1, p5_2, p5_3, p5_4, p5_5, p5_6, p5_7, p6, p7_1, p7_2, p7_3, p7_4, p7_5, p7_6, p7_7, p7_8, p7_9, p8_1, p8_2, p8_3, p8_4, p8_5, p8_6, p8_7, p8_8, p8_9, p8_10, p8_11, p8_12, p8_13, p8_14, p8_15, p8_16, p8_17, p8_18, p8_19, p8_20, p8_21, p9, p10, p11, p12_1, p12_2, p12_3, p12_4, p12_5, p12_6, p12_7, p12_8, p12_9, p12_10, p12_11, p12_12, p12_13, p13_1, p13_2, p13_3).Scan(&idGenerado)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	r.ParseForm()

	detalle1 := r.Form["ayuda"]
	if len(detalle1) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO MultipleSanamente (idp, respuesta, idF) VALUES (@p1,@p2,@p3)")
		if err != nil {
			http.Error(w, "Error al preparar la consulta apoyo especializado"+err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return
		}
		defer stmt.Close()

		for _, valor := range detalle1 {
			_, err = stmt2.Exec(idGenerado, valor, 9)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta apoyo especializado", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	t, err := template.ParseFiles("public/HTML/RespuestaSanamente.html")
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

func postEncuestaSMFinal(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	//NombreCampo en la BD: = r.FormValue("NombreCampo en el html")
	// Obtener los valores de los campos del formulario

	identidad := r.FormValue("dpi")
	sexo := cambio(r.FormValue("sexo"))
	fechaNac := r.FormValue("fechaN")
	rolTrabajo := cambio(r.FormValue("puesto"))
	otrorol := r.FormValue("otrorol")
	p1_1 := cambio(r.FormValue("practica1"))
	p1_2 := cambio(r.FormValue("practica2"))
	p1_3 := cambio(r.FormValue("practica3"))
	p1_4 := cambio(r.FormValue("practica4"))
	p1_5 := cambio(r.FormValue("practica5"))
	p1_6 := cambio(r.FormValue("practica6"))
	p1_7 := cambio(r.FormValue("practica7"))
	p1_8 := cambio(r.FormValue("practica8"))
	p1_9 := cambio(r.FormValue("practica9"))
	p1_10 := cambio(r.FormValue("practica10"))
	p1_11 := cambio(r.FormValue("practica11"))
	p1_12 := cambio(r.FormValue("practica12"))
	p1_13 := cambio(r.FormValue("practica13"))
	p1_14 := cambio(r.FormValue("practica14"))
	p1_15 := cambio(r.FormValue("practica15"))
	p1_16 := cambio(r.FormValue("practica16"))
	p1_17 := cambio(r.FormValue("practica17"))
	p1_18 := cambio(r.FormValue("practica18"))
	p1_19 := cambio(r.FormValue("practica19"))
	p1_20 := cambio(r.FormValue("practica20"))
	p2 := cambio(r.FormValue("aplicaautoc"))
	p3_1 := cambio(r.FormValue("afirma21"))
	p3_2 := cambio(r.FormValue("afirma22"))
	p3_3 := cambio(r.FormValue("afirma23"))
	p3_4 := cambio(r.FormValue("afirma24"))
	p3_5 := cambio(r.FormValue("afirma25"))
	p3_6 := cambio(r.FormValue("afirma26"))
	p3_7 := cambio(r.FormValue("afirma27"))
	p4_1 := cambio(r.FormValue("afecta28"))
	p4_2 := cambio(r.FormValue("afecta29"))
	p4_3 := cambio(r.FormValue("afecta30"))
	p5 := cambio(r.FormValue("actitud"))
	p6_1 := cambio(r.FormValue("situacion31"))
	p6_2 := cambio(r.FormValue("situacion32"))
	p6_3 := cambio(r.FormValue("situacion33"))
	p6_4 := cambio(r.FormValue("situacion34"))
	p6_5 := cambio(r.FormValue("situacion35"))
	p6_6 := cambio(r.FormValue("situacion36"))
	p6_7 := cambio(r.FormValue("situacion37"))
	p7 := r.FormValue("angustia")
	p8_1 := cambio(r.FormValue("vida38"))
	p8_2 := cambio(r.FormValue("vida39"))
	p8_3 := cambio(r.FormValue("vida40"))
	p8_4 := cambio(r.FormValue("vida41"))
	p8_5 := cambio(r.FormValue("vida42"))
	p8_6 := cambio(r.FormValue("vida43"))
	p8_7 := cambio(r.FormValue("vida44"))
	p8_8 := cambio(r.FormValue("vida45"))
	p8_9 := cambio(r.FormValue("vida46"))
	p9_1 := cambio(r.FormValue("estrategia47"))
	p9_2 := cambio(r.FormValue("estrategia48"))
	p9_3 := cambio(r.FormValue("estrategia49"))
	p9_4 := cambio(r.FormValue("estrategia50"))
	p9_5 := cambio(r.FormValue("estrategia51"))
	p9_6 := cambio(r.FormValue("estrategia52"))
	p9_7 := cambio(r.FormValue("estrategia53"))
	p9_8 := cambio(r.FormValue("estrategia54"))
	p9_9 := cambio(r.FormValue("estrategia55"))
	p9_10 := cambio(r.FormValue("estrategia56"))
	p9_11 := cambio(r.FormValue("estrategia57"))
	p9_12 := cambio(r.FormValue("estrategia58"))
	p9_13 := cambio(r.FormValue("estrategia59"))
	p9_14 := cambio(r.FormValue("estrategia60"))
	p9_15 := cambio(r.FormValue("estrategia61"))
	p9_16 := cambio(r.FormValue("estrategia62"))
	p9_17 := cambio(r.FormValue("estrategia63"))
	p9_18 := cambio(r.FormValue("estrategia64"))
	p9_19 := cambio(r.FormValue("estrategia65"))
	p9_20 := cambio(r.FormValue("estrategia66"))
	p9_21 := cambio(r.FormValue("estrategia67"))
	p10 := cambio(r.FormValue("realidad"))
	p11 := cambio(r.FormValue("objetivoSM"))
	p12 := cambio(r.FormValue("efectosT"))
	conocimientoTr := r.FormValue("conocimientoT")
	p13 := cambio(r.FormValue("actitudesF"))
	nreferido := r.FormValue("nreferido")
	npracticas := r.FormValue("npracticas")
	p16 := cambio(r.FormValue("caso"))
	p19_1 := cambio(r.FormValue("entorno68"))
	p19_2 := cambio(r.FormValue("entorno69"))
	p19_3 := cambio(r.FormValue("entorno70"))
	p19_4 := cambio(r.FormValue("entorno71"))
	p19_5 := cambio(r.FormValue("entorno72"))
	p19_6 := cambio(r.FormValue("entorno73"))
	p19_7 := cambio(r.FormValue("entorno74"))
	p19_8 := cambio(r.FormValue("entorno75"))
	p19_9 := cambio(r.FormValue("entorno76"))
	p19_10 := cambio(r.FormValue("entorno77"))
	p19_11 := cambio(r.FormValue("entorno78"))
	p19_12 := cambio(r.FormValue("entorno79"))
	p19_13 := cambio(r.FormValue("entorno80"))
	p20 := cambio(r.FormValue("clima"))
	climaLaboral := r.FormValue("climaT")
	p21_1 := cambio(r.FormValue("rel1"))
	p21_2 := cambio(r.FormValue("rel2"))
	p21_3 := cambio(r.FormValue("rel3"))

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	// Insertar los valores en la tabla correspondiente

	stmt, err := tx.Prepare("INSERT INTO SanamenteEncuestaFinal_EFPEM (identidad, sexo, fechaNac, rolTrabajo, otrorol, p1_1, p1_2, p1_3, p1_4, p1_5, p1_6, p1_7, p1_8, p1_9, p1_10, p1_11, p1_12, p1_13, p1_14, p1_15, p1_16, p1_17, p1_18, p1_19, p1_20, p2, p3_1, p3_2, p3_3, p3_4, p3_5, p3_6, p3_7, p4_1, p4_2, p4_3, p5, p6_1, p6_2, p6_3, p6_4, p6_5, p6_6, p6_7, p7, p8_1, p8_2, p8_3, p8_4, p8_5, p8_6, p8_7, p8_8, p8_9, p9_1, p9_2, p9_3, p9_4, p9_5, p9_6, p9_7, p9_8, p9_9, p9_10, p9_11, p9_12, p9_13, p9_14, p9_15, p9_16, p9_17, p9_18, p9_19, p9_20, p9_21, p10, p11, p12, conocimientoTr, p13, nreferido, npracticas, p16, p19_1, p19_2, p19_3, p19_4, p19_5, p19_6, p19_7, p19_8, p19_9, p19_10, p19_11, p19_12, p19_13, p20, climaLaboral, p21_1, p21_2, p21_3) OUTPUT INSERTED.ID VALUES (@P1, @P2, @P3, @P4, @P5, @P6, @P7, @P8, @P9, @P10, @P11, @P12, @P13, @P14, @P15, @P16, @P17, @P18, @P19, @P20, @P21, @P22, @P23, @P24, @P25, @P26, @P27, @P28, @P29, @P30, @P31, @P32, @P33, @P34, @P35, @P36, @P37, @P38, @P39, @P40, @P41, @P42, @P43, @P44, @P45, @P46, @P47, @P48, @P49, @P50, @P51, @P52, @P53, @P54, @P55, @P56, @P57, @P58, @P59, @P60, @P61, @P62, @P63, @P64, @P65, @P66, @P67, @P68, @P69, @P70, @P71, @P72, @P73, @P74, @P75, @P76, @P77, @P78, @P79, @P80, @P81, @P82, @P83, @P84, @P85, @P86, @P87, @P88, @P89, @P90, @P91, @P92, @P93, @P94, @P95, @P96, @P97, @P98, @P99, @P100, @P101)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	var idGenerado int
	err = stmt.QueryRow(identidad, sexo, fechaNac, rolTrabajo, otrorol, p1_1, p1_2, p1_3, p1_4, p1_5, p1_6, p1_7, p1_8, p1_9, p1_10, p1_11, p1_12, p1_13, p1_14, p1_15, p1_16, p1_17, p1_18, p1_19, p1_20, p2, p3_1, p3_2, p3_3, p3_4, p3_5, p3_6, p3_7, p4_1, p4_2, p4_3, p5, p6_1, p6_2, p6_3, p6_4, p6_5, p6_6, p6_7, p7, p8_1, p8_2, p8_3, p8_4, p8_5, p8_6, p8_7, p8_8, p8_9, p9_1, p9_2, p9_3, p9_4, p9_5, p9_6, p9_7, p9_8, p9_9, p9_10, p9_11, p9_12, p9_13, p9_14, p9_15, p9_16, p9_17, p9_18, p9_19, p9_20, p9_21, p10, p11, p12, conocimientoTr, p13, nreferido, npracticas, p16, p19_1, p19_2, p19_3, p19_4, p19_5, p19_6, p19_7, p19_8, p19_9, p19_10, p19_11, p19_12, p19_13, p20, climaLaboral, p21_1, p21_2, p21_3).Scan(&idGenerado)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	r.ParseForm()

	detalle1 := r.Form["ayuda"]
	if len(detalle1) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO MultipleSanamente (idp, respuesta, idPregunta, idF) VALUES (@p1,@p2,@p3,@p4)")
		if err != nil {
			http.Error(w, "Error al preparar la consulta apoyo especializado"+err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return
		}
		defer stmt.Close()

		for _, valor := range detalle1 {
			_, err = stmt2.Exec(idGenerado, valor, 1, 13)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta apoyo especializado", http.StatusInternalServerError)
				return
			}
		}
	}

	detalle2 := r.Form["fomentaautoc"]
	if len(detalle2) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO MultipleSanamente (idp, respuesta, idPregunta, idF) VALUES (@p1,@p2,@p3,@p4)")
		if err != nil {
			http.Error(w, "Error al preparar la consulta fomenta autocuidado"+err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return
		}
		defer stmt.Close()

		for _, valor := range detalle2 {
			_, err = stmt2.Exec(idGenerado, valor, 2, 13)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta fomenta autocuidado", http.StatusInternalServerError)
				return
			}
		}
	}

	detalle3 := r.Form["practicasAt"]
	if len(detalle3) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO MultipleSanamente (idp, respuesta, idPregunta, idF) VALUES (@p1,@p2,@p3,@p4)")
		if err != nil {
			http.Error(w, "Error al preparar la consulta practica autocuidado"+err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return
		}
		defer stmt.Close()

		for _, valor := range detalle3 {
			_, err = stmt2.Exec(idGenerado, valor, 14, 13)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta practica autocuidado", http.StatusInternalServerError)
				return
			}
		}
	}

	detalle4 := r.Form["practicasEq"]
	if len(detalle4) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO MultipleSanamente (idp, respuesta, idPregunta, idF) VALUES (@p1,@p2,@p3,@p4)")
		if err != nil {
			http.Error(w, "Error al preparar la consulta practica equipo"+err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return
		}
		defer stmt.Close()

		for _, valor := range detalle4 {
			_, err = stmt2.Exec(idGenerado, valor, 17, 13)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta practica equipo", http.StatusInternalServerError)
				return
			}
		}
	}

	detalle5 := r.Form["practicasTr"]
	if len(detalle5) == 0 {
	} else {
		stmt2, err := tx.Prepare("INSERT INTO MultipleSanamente (idp, respuesta, idPregunta, idF) VALUES (@p1,@p2,@p3,@p4)")
		if err != nil {
			http.Error(w, "Error al preparar la consulta practica trabajo"+err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return
		}
		defer stmt.Close()

		for _, valor := range detalle5 {
			_, err = stmt2.Exec(idGenerado, valor, 18, 13)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta practica trabajo", http.StatusInternalServerError)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	t, err := template.ParseFiles("public/HTML/RespuestaSanamente.html")
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

func ObtenersedeGeneral(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	departamento := r.FormValue("departamento")
	programa := r.FormValue("programa")
	var opcionesmunsede string
	// Build connection string
	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database4)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()
	stmt, err := db.Prepare("Select distinct idSede,Sede from sede_areas where Iddep=@p1 and Ida=@p2 and IdSede not in (135,1636,3177) order by Sede")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := stmt.Query(departamento, programa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionesmunsede += "<option value=''></option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesmunsede += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionesmunsede)
}
func formulariomdvpsicologicoSantaana(w http.ResponseWriter, Cfg *Controller.Config) {
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	municipio, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM municipalities where fkCodeState='CP1101' order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estado_civ, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM F_estadociv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Respuestas:   respuestas,
		Estado_civ:   estado_civ,
		Departamento: municipio,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/mdv/psicologico copy.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postmdvpsicologicoSantaana(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	// Create connection pool
	colonia := cambio(r.FormValue("colonia"))
	otracolonia := r.FormValue("otraColonia")
	nombre := r.FormValue("nombre")
	diplomado := r.FormValue("diplomado")
	fechanac := r.FormValue("fechanac")
	estadoCivil := r.FormValue("estadoCivil")
	motivo := r.FormValue("motivo")
	otraOcupacion := r.FormValue("otraOcupacion")
	convivencia := r.FormValue("convivencia")
	considera := r.FormValue("considera")
	otraSituacion := r.FormValue("otraSituacion")
	situacionesMas := r.FormValue("situacionesmas")
	psicologico := r.FormValue("psicologico")
	otroPsicologico := r.FormValue("otroPsicologico")
	municipio := r.FormValue("municipio")
	idvn := r.FormValue("idvn")
	hijos := cambio(r.FormValue("hijos"))
	ludoteca := cambio(r.FormValue("ludoteca"))
	hijosludoteca := cambio(r.FormValue("hijosludoteca"))
	papa := r.FormValue("papa")
	sede := r.FormValue("sede")
	enfermedad := r.FormValue("enfermedad")

	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error conexion", http.StatusInternalServerError)

	}

	stmt, err := tx.Prepare("INSERT INTO RespuestasPsicosocial(sede,nombre, diplomado, fechanac, estadoCivil, motivo, otraOcupacion, convivencia, considera, otraSituacion, situacionesMas, psicologico, otroPsicologico, idvn,municipio,papa,hijos,ludoteca,hijosludoteca,colonia,otracolonia,enfermedad) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14,@P15,@P16,@P17,@P18,@P19,@P20,@P21,@P22)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)

	}

	var lastInsertId int
	err = stmt.QueryRow(sede, nombre, diplomado, fechanac, estadoCivil, motivo, otraOcupacion, convivencia, considera, otraSituacion, situacionesMas, psicologico, otroPsicologico, idvn, municipio, papa, hijos, ludoteca, hijosludoteca, colonia, otracolonia, enfermedad).Scan(&lastInsertId)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
	}

	ocupacion := r.Form["ocupacion"]
	if len(ocupacion) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO RespMutiNumeros (idPregunta,idp,idRespuesta) VALUES (1,@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta ocupacion", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range ocupacion {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					tx.Rollback()
					return

				}
			}
		}
	}

	situaciones := r.Form["situaciones"]
	if len(situaciones) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO RespMutiNumeros (idPregunta,idp,idRespuesta) VALUES (2,@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta situaciones", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range situaciones {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					tx.Rollback()
					return

				}
			}
		}
	}

	parentescomas := r.Form["parentescomas"]
	if len(parentescomas) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO RespMutiNumeros (idPregunta,idp,idRespuesta) VALUES (3,@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta situaciones", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range parentescomas {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					tx.Rollback()
					return

				}
			}
		}
	}

	parentescomenos := r.Form["parentescomenos"]
	if len(parentescomenos) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO RespMutiNumeros (idPregunta,idp,idRespuesta) VALUES (4,@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta situaciones", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range parentescomenos {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					tx.Rollback()
					return

				}
			}
		}
	}

	destaca := r.Form["destaca"]
	if len(ocupacion) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO RespMutiLetras (idPregunta,idp,respuesta) VALUES (5,@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta situaciones", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range destaca {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					tx.Rollback()
					return

				}
			}
		}
	}

	positiva := r.Form["positiva"]
	if len(positiva) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO RespMutiLetras (idPregunta,idp,respuesta) VALUES (6,@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta situaciones", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range positiva {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					tx.Rollback()
					return

				}
			}
		}
	}
	mejorar := r.Form["mejorar"]
	if len(mejorar) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO RespMutiLetras (idPregunta,idp,respuesta) VALUES (7,@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta situaciones", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range mejorar {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					tx.Rollback()
					return

				}
			}
		}
	}

	ida := r.Form["ida"]
	if len(ida) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO RespMutiLetras (idPregunta,idp,respuesta) VALUES (8,@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta situaciones", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range ida {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					tx.Rollback()
					return

				}
			}
		}
	}
	idv := r.Form["idv"]
	if len(idv) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO RespMutiLetras (idPregunta,idp,respuesta) VALUES (9,@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta situaciones", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range idv {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					tx.Rollback()
					return

				}
			}
		}
	}
	salariodependencia := r.Form["salariodependencia"]
	if len(salariodependencia) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO RespMutiNumeros (idPregunta,idp,idRespuesta) VALUES (10,@P1,@P2)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta situaciones", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range salariodependencia {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					tx.Rollback()
					return

				}
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("Public/Html/Respuesta.html")
	// Renderizar plantilla HTML con opciones de preguntas select
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func moduloacolHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=16 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=16  and activo=1 order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/col/ModuloAcol.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func modulobcolHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=16 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=16  and activo=1  order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/col/ModuloBcol.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func moduloccolHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM admin_main_gwdata.VSedessanamente where IdPais=16 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=16 and activo=1 order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/col/ModuloCcol.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
