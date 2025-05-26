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


var contador = 4;

function fmun(paisSeleccionado) {
  var ciudadSelect = document.getElementById("municipio");
  ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
  var ciudad
  if (paisSeleccionado == "666") {
    ciudad = "2";
  } else if (paisSeleccionado == "2968") {
    ciudad = "123";}
    else if (paisSeleccionado == "1128") {
      ciudad = "76";
  }
  // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
          var opcionesCiudades = this.responseText;
          ciudadSelect.innerHTML = opcionesCiudades; // Actualizar la lista de ciudades
      }
  };
  xhr.open("POST", "/obtenerciudades", true);
  xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhr.send("pais=" + ciudad);
}

function fotro(otro){
  var otroSelect =document.getElementById("otro_colonia_div")
  var input2 = document.getElementById("otraColonia")
if (otro=="3"){
  otroSelect.classList.remove('d-none');
  input2.required=true
} else {
  otroSelect.classList.add('d-none');
  input2.required=false
}
}

function fcolonia(municipioSeleccionado) {
  var coloniaSelect = document.getElementById("colonia_div");
  var otroSelect =document.getElementById("otro_colonia_div")
  var sedeSelect = document.getElementById("sede");
  var input = document.getElementById("colonia")
  var input2 = document.getElementById("otraColonia")
  if (municipioSeleccionado == "2159" && sedeSelect.value == "2968") {
    coloniaSelect.classList.remove('d-none');
    otroSelect.classList.add('d-none');
    input.required=true
    input2.required=false
  } else if(sedeSelect.value == "2968") {
    coloniaSelect.classList.add('d-none');
    otroSelect.classList.remove('d-none');
    input.required=false
    input2.required=true
  } else {
    coloniaSelect.classList.add('d-none');
    input.required=false
    input2.required=false
  }
}

function agregarInput() {
  event.preventDefault()
  event.stopPropagation()
    var div = document.getElementById('inputs');

    var input = document.createElement('input');
    input.type = 'text';
    input.name = 'destaca';
    input.id = 'destaca' + contador;
    input.className = 'form-control';
    input.maxLength = '300';

    div.appendChild(input);

    contador++;
}

var contador2 = 4;

function agregarInput2() {
  event.preventDefault()
  event.stopPropagation()
    var div = document.getElementById('inputs2');

    var input = document.createElement('input');
    input.type = 'text';
    input.name = 'positiva';
    input.id = 'positiva' + contador2;
    input.className = 'form-control';
    input.maxLength = '300';

    div.appendChild(input);

    contador2++;
}

var contador3 = 4;

function agregarInput3() {
  event.preventDefault()
  event.stopPropagation()
    var div = document.getElementById('inputs3');

    var input = document.createElement('input');
    input.type = 'text';
    input.name = 'mejorar';
    input.id = 'mejorar' + contador3;
    input.className = 'form-control';
    input.maxLength = '300';

    div.appendChild(input);

    contador3++;
}

var contador4 = 3;

function agregarInput4() {
  event.preventDefault()
  event.stopPropagation()
    var div = document.getElementById('inputs4');

    var input = document.createElement('input');
    input.type = 'text';
    input.name = 'ida';
    input.id = 'ida' + contador4;
    input.className = 'form-control';
    input.maxLength = '50';

    div.appendChild(input);

    contador4++;
}


var contador5 = 3;

function agregarInput5() {
  event.preventDefault()
  event.stopPropagation()
    var div = document.getElementById('inputs5');

    var input = document.createElement('input');
    input.type = 'text';
    input.name = 'idv';
    input.id = 'idv' + contador5;
    input.className = 'form-control';
    input.maxLength = '50';

    div.appendChild(input);

    contador5++;
}
function otropsicologico(value){
if (value =='1'){document.getElementById('otro_psicologico_div').classList.remove('d-none');
}
else{
  document.getElementById('otro_psicologico_div').classList.add('d-none');
}}

function edad(){
  var select = document.getElementById("fechanac");
  var birthDate = new Date(select.value);
var diff_ms = Date.now() -birthDate ;
var age_dt = new Date(diff_ms); 
var age= Math.abs(age_dt.getUTCFullYear() - 1970);
if(age>=18){
  document.getElementById('divmayor').classList.remove('d-none');
  document.getElementById('mayor').required = true;
  document.getElementById('divmenor').classList.add('d-none');
}else
{
  document.getElementById('divmayor').classList.add('d-none');
  document.getElementById('mayor').required = false;
  document.getElementById('divmenor').classList.remove('d-none');
}

}

function permiso() {
  var checkbox = document.getElementById('mayor');
  var divMenor = document.getElementById('divmenor');
  if (checkbox.checked) {
      divMenor.classList.remove('d-none');
  } else {
      divMenor.classList.add('d-none');
  }
}

function Otro() {
  var checkbox = document.getElementById('ocupacionOtro');
  var div = document.getElementById('otro_ocupacion_div');
  if (checkbox.checked) {
    div.classList.remove('d-none');
  } else {
    div.classList.add('d-none');
  }
}

function Otro2() {
  var checkbox = document.getElementById('situacionesotro');
  var div = document.getElementById('otro_situaciones_div');
  if (checkbox.checked) {
    div.classList.remove('d-none');
  } else {
    div.classList.add('d-none');
  }
}

function Otro3() {
  var checkbox = document.getElementById('situacionesninguno');
  var div = document.getElementById('situaciones_ninguno_div');
  if (checkbox.checked) {
    div.classList.add('d-none');
  } else {
    div.classList.remove('d-none');
  }
}

function mostrarhijos(value){
  if (value =='1'){document.getElementById('hijos_div').classList.remove('d-none');
  }
  else{
    document.getElementById('hijos_div').classList.add('d-none');
  }
}

function mostrarhijosludoteca(value){
  if (value =='1'){document.getElementById('hijosludoteca_div').classList.remove('d-none');
  }
  else{
    document.getElementById('hijosludoteca_div').classList.add('d-none');
  }
}

function enfermadaddiv() {
  var checkbox = document.getElementById("enfermedadcheck");
  var div = document.getElementById('enfermedad');
  if (checkbox.checked) {
    div.classList.remove('d-none');
  } else {
    div.classList.add('d-none');
  }
}