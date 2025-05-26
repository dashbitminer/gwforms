// Ejemplo de JavaScript inicial para deshabilitar el envío de formularios si hay campos no válidos
(function () {
    'use strict'
  
    // Obtener todos los formularios a los que queremos aplicar estilos de validación de Bootstrap personalizados
    var forms = document.querySelectorAll('.needs-validation')
  
    // Bucle sobre ellos y evitar el envío
    Array.prototype.slice.call(forms)
      .forEach(function (form) {
        form.addEventListener('submit', function (event) {
          if (!form.checkValidity()) {
            event.preventDefault()
            event.stopPropagation()
          }
  
          form.classList.add('was-validated')
        }, false)
      })
  })()

  function btipo(seguimiento) {
    var ciudadSelect = document.getElementById("subtipodiv");
    var participante = document.getElementById("participante");
    var sede = document.getElementById("sede");
    ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
    // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var opcionesmunsede= this.responseText;
            ciudadSelect.innerHTML = opcionesmunsede; // Actualizar la lista de ciudades
        }
    };
    xhr.open("POST", "/obtenerseguimiento", true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.send("seguimiento=" + seguimiento + "&participante=" + participante.value + "&sede=" + sede.value);
  }

  function listado() {
    var sede= document.getElementById("sede").value;
    var ano= document.getElementById("ano").value;
    var cohorte= document.getElementById("cohorte").value;
    var ciudadSelect = document.getElementById("participante");
    ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
    // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
   if(sede==""||ano==""||cohorte==""){}
    else{
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var opcionesmunsede= this.responseText;
            ciudadSelect.innerHTML = opcionesmunsede; // Actualizar la lista de ciudades
        }
    };
    xhr.open("POST", "/obtenerlistadomd", true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.send("sede=" + sede + "&ano=" + ano + "&cohorte=" + cohorte);
  }
  }

  function fsubtipo(subtipo){
    var seguimiento = document.getElementById("tipo").value;
    var subtipoSelect = document.getElementById("subtipodiv2");
    
    if(subtipo=="1"||subtipo=="2"||subtipo=="3"||subtipo=="5"||subtipo=="6"||subtipo=="7"){
        // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
          subtipoSelect.innerHTML = ""; // Limpiar la lista de ciudades
          var opcionesmunsede= this.responseText;
            subtipoSelect.innerHTML = opcionesmunsede; // Actualizar la lista de ciudades
          if(subtipo=="1"||subtipo=="2"||subtipo=="6"){createSignaturePads();}
          }
    };
    xhr.open("POST", "/obtenerseguimientosubtipo", true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.send("seguimiento=" + seguimiento + "&subtipo=" + subtipo);
  }
  else{ subtipoSelect.innerHTML = ""; }// Limpiar la lista de ciudades
}

  function fseguimientos(subtipo){

    var subtipoSelect = document.getElementById("seguimientosdivnuevo");
    subtipoSelect.innerHTML = ""; // Limpiar la lista de ciudades
    // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var opcionesmunsede= this.responseText;
            subtipoSelect.innerHTML = opcionesmunsede; // Actualizar la lista de ciudades
        }
    };
    xhr.open("POST", "/mdv/seguimientoopciones", true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.send("subtipo=" + subtipo);
  }

  function fenlace(subtipo){

    var subtipoSelect = document.getElementById("enlacediv");
    subtipoSelect.innerHTML = ""; // Limpiar la lista de ciudades
    // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var opcionesmunsede= this.responseText;
            subtipoSelect.innerHTML = opcionesmunsede; // Actualizar la lista de ciudades
        }
    };
    xhr.open("POST", "/obtenerenlace", true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.send("enlace=" + subtipo);
  }

  function fenlace2(subtipo){

    var subtipoSelect = document.getElementById("e3div");
    subtipoSelect.innerHTML = ""; // Limpiar la lista de ciudades
    // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var opcionesmunsede= this.responseText;
            subtipoSelect.innerHTML = opcionesmunsede; // Actualizar la lista de ciudades
        }
    };
    xhr.open("POST", "/obtenerenlace2", true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.send("enlace=" + subtipo);
  }



  function s3(subtipo){

    var subtipoSelect = document.getElementById("s3div");
    subtipoSelect.innerHTML = ""; // Limpiar la lista de ciudades
    // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var opcionesmunsede= this.responseText;
            subtipoSelect.innerHTML = opcionesmunsede; // Actualizar la lista de ciudades
        }
    };
    xhr.open("POST", "/obtenerseguimiento2", true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.send("enlace=" + subtipo);
  }

  function fcomentario(fc){
    if(fc=="1"){
      document.getElementById("comentariosdiv").classList.remove("d-none");
      document.getElementById("comentarios").required = true;
    } else {
      document.getElementById("comentariosdiv").classList.add("d-none");
      document.getElementById("comentarios").required = false;
    }

  }

  window.onload = function() {
    var fechaFicha = document.getElementById('fechasesion');
    var fechaActual = new Date();
    fechaFicha.max = fechaActual.toISOString().split('T')[0];
  }

var signaturePad;
var signaturePad2;

function createSignaturePads(){
    var input = document.getElementById("Signature1draw");
    if (typeof signaturePad !== "undefined" && signaturePad) {
        signaturePad.clear(); // Limpiar la instancia existente
    }
    signaturePad = new SignaturePad(input); // Crear una nueva instancia

    var input2 = document.getElementById("Signature2draw");
    if (typeof signaturePad2 !== "undefined" && signaturePad2) {
        signaturePad2.clear(); // Limpiar la instancia existente
    }
    signaturePad2 = new SignaturePad(input2); // Crear una nueva instancia
}

function clearSignature(){
    if (typeof signaturePad !== "undefined" && signaturePad) {
        signaturePad.clear();
    }
}

function clearSignature2(){
    if (typeof signaturePad2 !== "undefined" && signaturePad2) {
        signaturePad2.clear();
    }
}
function agregarhabito(event) {
    event.preventDefault();

    var bloqueHtml = "<label for='habito' class='form-label' style='font-weight: bold;'>Escriba el hábito que se propuso la o él participante</label>";
    bloqueHtml += " <input type='text' name='habito' class='form-control' maxlength='500' required>";
    bloqueHtml += "<label for='habito' class='form-label' style='font-weight: bold;'>Que debo hacer o dejar de hacer (algo observable)</label>";
    bloqueHtml += " <input type='text' name='compromiso' class='form-control' maxlength='500' required>";
    bloqueHtml += " <input type='text' name='compromiso' class='form-control' maxlength='500'>";
    bloqueHtml += " <input type='text' name='compromiso' class='form-control' maxlength='500'>";
   
    var div = document.getElementById('subtipodiv2');
    div.insertAdjacentHTML('beforeend', bloqueHtml);
}

function agregarpuntos(event) {
  event.preventDefault();
  var bloqueHtml = "<input type='text' class='form-control' name='abordados' required>";
  var div = document.getElementById('agregarpuntos');
  div.insertAdjacentHTML('beforeend', bloqueHtml);
}

function agregaracuerdos(event) {
  event.preventDefault();

  var bloqueHtml = "<input type='text' class='form-control' name='acuerdos' required>";
  var div = document.getElementById('agregaracuerdo');
  div.insertAdjacentHTML('beforeend', bloqueHtml);
}

function agregarapoyos(event) {
  event.preventDefault();

  var bloqueHtml = "<input type='text' class='form-control' name='grupoapoyo' required>";
  var div = document.getElementById('agregarapoyo');
  div.insertAdjacentHTML('beforeend', bloqueHtml);
}

function fab(subtipo){
 var subtipoSelect = document.getElementById("ext");
subtipoSelect.innerHTML = ""; // Limpiar y dejar en blanco

if (subtipo=="2"){
    var bloqueHtml = "<br><label for='ext' class='form-label' style='font-weight: bold;'>¿Qué tipo de seguimiento externo se brindó?</label>";
    bloqueHtml += "  <select class='form-select' id='ext' name='ext' required>"
    bloqueHtml += " <option value='1'>Seguimiento a fondo de crisis</option>";
    bloqueHtml += " <option value='2'>Seguimiento a sistema de referencias (SanaMente)</option>";
    bloqueHtml += "</select>";
    subtipoSelect.insertAdjacentHTML('beforeend', bloqueHtml);
  }
  if (subtipo=="5"){
    var bloqueHtml = "<br><label for='ext' class='form-label' style='font-weight: bold;'>Especifique:</label>";
    bloqueHtml += "<input type='text' class='form-control' name='otrarazon' maxlength='500' required>";
    subtipoSelect.insertAdjacentHTML('beforeend', bloqueHtml);
  }
}

function guardar(event){
event.preventDefault(); // Evitar el envío del formulario
var form = document.getElementById("form");
var formData = new FormData(form);
// verificar los campos requeridos
if (form.checkValidity() === false) {
    form.reportValidity();
    event.stopPropagation();
    return;
}if (typeof signaturePad !== "undefined" && signaturePad && !signaturePad.isEmpty()) {
    var firma1 = signaturePad.toDataURL();
    formData.append("firma1", firma1);
} 

if (typeof signaturePad2 !== "undefined" && signaturePad2 && !signaturePad2.isEmpty()) {
    var firma2 = signaturePad2.toDataURL();
    formData.append("firma2", firma2);
} 
// Enviar petición AJAX
var xhr = new XMLHttpRequest();
xhr.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) {
        var respuesta = this.responseText;
        if(respuesta=="1"){
            alert("Se guardó correctamente");
            // redireccionar a la página de inicio
            window.location.href = "/juventud/gestion/sv";
        } else {
            alert("Hubo un error al guardar, intente de nuevo");
        }
    }
};
xhr.open("POST", "/juventud/gestion/sv/post", true);
xhr.send(formData);
}