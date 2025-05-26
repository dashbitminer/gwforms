// Example starter JavaScript for disabling form submissions if there are invalid fields
(() => {
  'use strict'

  // Fetch all the forms we want to apply custom Bootstrap validation styles to
  const forms = document.querySelectorAll('.needs-validation')

  // Loop over them and prevent submission
  Array.from(forms).forEach(form => {
    form.addEventListener('submit', event => {
      if (!form.checkValidity()) {
        event.preventDefault()
        event.stopPropagation()
      }

      form.classList.add('was-validated')
    }, false)
  })
})()



function activaselect2(){
  var select = document.getElementById("Eactual");
var hiddenDiv = document.getElementById("hidden-question-div");
var hiddenDiv2 = document.getElementById("hidden-question-div2");
var hiddenDiv3 = document.getElementById("hidden-question-div3");
var hiddenDiv4 = document.getElementById("hidden-question-div5");
var input = document.getElementById("UT");
var input2 = document.getElementById("GA");
var input3 = document.getElementById("Turno");
var input4 = document.getElementById("seccion");
  if(select.value==="1"){    
      hiddenDiv.classList.add("d-none");
      input.required = false;
      hiddenDiv2.classList.remove("d-none");
      input2.required = true;
      
      activaselect3();
  }
  
   else {
  hiddenDiv.classList.remove("d-none");
  input.required = true;
    hiddenDiv2.classList.add("d-none");
  input2.required = false;
    hiddenDiv3.classList.add("d-none");
  input3.required = false;
      hiddenDiv4.classList.add("d-none");
  input4.required = false;
  }

  }

function activaselect3(){
var hiddenDiv3 = document.getElementById("hidden-question-div3");
var hiddenDiv4 = document.getElementById("hidden-question-div5");
var input2 = document.getElementById("GA");
var input3 = document.getElementById("Turno");
var input4 = document.getElementById("seccion");
  if(input2.value==="38"||input2.value==="39"||input2.value==="40"){ 
  hiddenDiv3.classList.add("d-none");
  input3.required = false;
      hiddenDiv4.classList.add("d-none");
  input4.required = false;}
  else{
  hiddenDiv3.classList.remove("d-none");
  input3.required = false;
    hiddenDiv4.classList.remove("d-none");
  input4.required = false;
  }
  }
  
function activaselect10(){
  var select = document.getElementById("dis");
var hiddenDiv = document.getElementById("hidden-question-div10");
  if(select.value==="1"){    
  hiddenDiv.classList.remove("d-none");
  } else {
  hiddenDiv.classList.add("d-none");
    }
  }


function activaselect51(){
  var select = document.getElementById("nacionalidad");
  var labelOutput = document.getElementById("DUIL");
  if(select.value==="1"){    
    labelOutput.innerHTML = "Número de CURP:*"
    document.getElementById("DUIM").placeholder = "abcd000000abcdef00";
  } else {labelOutput.innerHTML = "Número de tu Documento de Identidad"
    document.getElementById("DUIM").placeholder = ""
  } 
} 

document.querySelector('#formualrioRegistro').addEventListener('submit', function(event) {
var interests = document.querySelectorAll('input[name="dis2"]:checked');
  var select = document.getElementById("dis");
if(select.value==="1"){
if (interests.length === 0) {
  alert('Debes seleccionar al menos una opción de discapacidad en caso de poseer alguna');
  event.preventDefault();
}
}
});

function activaselect52(){
var select = document.getElementById("aut");
var hiddenDiv = document.getElementById("brindo");
var hiddenDiv2 = document.getElementById("brindo2");
var hiddenDiv3 = document.getElementById("brindo3");

var input9= document.getElementById("nombreCompleto");
var input10= document.getElementById("apellido");
var input11= document.getElementById("Sexo");
var input12= document.getElementById("pais");
var input13= document.getElementById("departamento");
var input22= document.getElementById("tipo");
var input23= document.getElementById("TdC");
var input28= document.getElementById("dis");
var input30= document.getElementById("gp");
var input45= document.getElementById("docF");
var input46= document.getElementById("docA");
var input50= document.getElementById("nacionalidad");


if(select.value==="1"){
  hiddenDiv.classList.remove("d-none") ;
  hiddenDiv2.classList.remove("d-none") ;
  hiddenDiv3.classList.add("d-none") 
 
  input9.required = true;
input10.required = true;
input11.required = true;
input12.required = true;
input13.required = true;
input22.required = true;
input23.required = true;
input28.required = true;
input30.required = true;
input45.required = true;
input46.required = true;
input50.required = true;

}
else {
   hiddenDiv.classList.add("d-none") ;
   hiddenDiv2.classList.add("d-none") ;
   hiddenDiv3.classList.remove("d-none") ;
   
   input9.required = false;
   input10.required = false;
   input11.required = false;
   input12.required = false;
   input13.required = false;
   input22.required = false;
   input23.required = false;
   input28.required = false;
   input30.required = false;
   input45.required = false;
  input46.required = false;  
  input50.required = false;
}
}

function activa1(){
  var input = document.getElementById("nombreCompleto");
  var input2 = document.getElementById("apellido");
  var output = document.getElementById("nombrep");
  var output2 = document.getElementById("nombrep2");
    output.innerHTML = input.value +" "+input2.value;
    output2.innerHTML = input.value +" "+input2.value;
  
  }
  
  function validateFormat() {
    var formatInput = document.getElementById("DUIM").value;
       var formatRegex = /^\d{13}$/;
                 if (!formatRegex.test(formatInput)) {
          var alertContainer = document.getElementById("alertContainer");
          alertContainer.innerHTML = '<div class="alert alert-danger" role="alert">El formato del DNI no es válido.</div>';
          alertContainer.style.display = "block";
          return false;
    }
    return true;
  
  }

      function hideAlert() {
    var alertContainer = document.getElementById("alertContainer");
    alertContainer.style.display = "none";
  }


  function actualizarCiudades(paisSeleccionado) {
    var ciudadSelect = document.getElementById("departamento");
    ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades

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
    xhr.send("pais=" + paisSeleccionado);
}

function actualizarmun(paisSeleccionado) {
  var ciudadSelect = document.getElementById("departamento2");
  ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
  var sede = document.getElementById("sede2");
  sede.innerHTML = ""; // Limpiar la lista de ciudades
  // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
          var pcionesmunsede= this.responseText;
          ciudadSelect.innerHTML = pcionesmunsede; // Actualizar la lista de ciudades
      }
  };
  xhr.open("POST", "/obtenermunsedesalud", true);
  xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhr.send("depsede=" + paisSeleccionado);
}

function actualizarsede(paisSeleccionado) {
  var ciudadSelect = document.getElementById("sede2");
  ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
  // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
          var opcionesmunsede= this.responseText;
          ciudadSelect.innerHTML = opcionesmunsede; // Actualizar la lista de ciudades
      }
  };
  xhr.open("POST", "/obtenersedesalud", true);
  xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhr.send("munsede=" + paisSeleccionado);
}

function activaselect5(programa) {
  var ciudadSelect = document.getElementById("mensaje");
  ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
  // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
          var opcionesmunsede= this.responseText;
          ciudadSelect.innerHTML = opcionesmunsede;
          document.getElementById("autorizaciones").classList.remove("d-none"); // Actualizar la lista de ciudades
      }
  };
  xhr.open("POST", "/obtenermensaje", true);
  xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhr.send("programa=" + programa);
}

function ver1(){
  var div= document.getElementById("datospersonales");
  var div1 = document.getElementById("finllenadoRes");
  div.classList.remove("d-none");
  div1.classList.add("d-none");
}

function ocultar1(){
  var div= document.getElementById("datospersonales");
  var div1 = document.getElementById("finllenadoRes");
  div.classList.add("d-none");
  div1.classList.remove("d-none");
}

function ver2(){
  var div= document.getElementById("datospersonales2");
  var div1 = document.getElementById("finllenadoRes1");
  div.classList.remove("d-none");
  div1.classList.add("d-none");
}

function ocultar2(){
  var div= document.getElementById("datospersonales2");
  var div1 = document.getElementById("finllenadoRes1");
  div.classList.add("d-none");
  div1.classList.remove("d-none");
}

function ver3(){
  var div= document.getElementById("datospersonalesm");
  var div1 = document.getElementById("finllenadoMenor");
  div.classList.remove("d-none");
  div1.classList.add("d-none");
}

function ocultar3(){
  var div= document.getElementById("datospersonalesm");
  var div1 = document.getElementById("finllenadoMenor");
  div.classList.add("d-none");
  div1.classList.remove("d-none");
}

function ver4(){
  var div= document.getElementById("datospersonalesm2");
  var div1 = document.getElementById("finllenadoMenor1");
  div.classList.remove("d-none");
  div1.classList.add("d-none");
}

function ocultar4(){
  var div= document.getElementById("datospersonalesm2");
  var div1 = document.getElementById("finllenadoMenor1");
  div.classList.add("d-none");
  div1.classList.remove("d-none");
}

window.onload = function() {
  var fechaFicha = document.getElementById('fechaficha');
  var fechaActual = new Date();
  fechaFicha.max = fechaActual.toISOString().split('T')[0];
}