// Example starter JavaScript for disabling form submissions if there are invalid fields
(() => {
  'use strict'

  // Fetch all the forms we want to apply custom Bootstrap validation styles to
  const forms = document.querySelectorAll('.needs-validation')

  // Loop over them and prevent submission
  Array.from(forms).forEach(form => {
    form.addEventListener('submit', event => {
      if(verificarSeleccion() === false){
        event.preventDefault();
      }
      if (!form.checkValidity()) {
        event.preventDefault()
        event.stopPropagation()
      }
       else if (document.getElementById("alertContainer").style.display === "block") {
    event.preventDefault();
    alert("No se puede enviar el formulario mientras este incorrecto el dui");
  }
    }, false)
  })
})()
  
function activaselect2(){
var select = document.getElementById("TdC");
var hiddenDiv = document.getElementById("hidden-question-div");
var hiddenDiv2 = document.getElementById("hidden-question-div2");

var input = document.getElementById("UT");
var input2 = document.getElementById("GA");

  if(select.value==="1"){    
      hiddenDiv.classList.add("d-none");
      input.required = false;
      hiddenDiv2.classList.remove("d-none");
      input2.required = true;        
  }
  
   else {
      hiddenDiv.classList.remove("d-none");
      input.required = true;
      hiddenDiv2.classList.add("d-none");
      input2.required = false;
  }

  }

  function activaselect7(){
      var select = document.getElementById("acc3");
      var hiddenDiv = document.getElementById("acc4d");
      var input4 = document.getElementById("acc4");
      var labelOutput = document.getElementById("acc4l");
      if(select.value==="1"){    
        labelOutput.innerHTML = "Selecciona tu rango/categoría:*"
        hiddenDiv.classList.remove("d-none");
        input4.required = true;
        hiddenDiv.classList.remove("d-none");
        input4.required = true;
        input4.innerHTML = ""; // Limpiar la lista de ciudades
    
        // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
        var xhr = new XMLHttpRequest();
        xhr.onreadystatechange = function() {
            if (this.readyState == 4 && this.status == 200) {
                var opcionesCiudades = this.responseText;
                input4.innerHTML = opcionesCiudades; // Actualizar la lista de ciudades
            }
        };
        xhr.open("POST", "/obtenersubperfil", true);
        xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        xhr.send("pais=" + select.value);
                 }     else {
        hiddenDiv.classList.add("d-none");
        input4.required = false;
          }
        }

function activaselect51(){
var select = document.getElementById("nacionalidad").value;
var input = document.getElementById("DUIM");
var label = document.getElementById("DUIL");
if(select==="1"){   
  label.innerHTML = "Escribe los últimos 4 dígitos de tu Documento de Identidad (DUI):*";
  input.placeholder = "0081";
}else{
  label.innerHTML = "Escribe los últimos 4 dígitos de tu carnet de residencia, documento extranjero o pasaporte:*";
  input.placeholder = "1081";
  }  
}
  document.addEventListener('DOMContentLoaded', (event) => {
      FSector();
});



function FSector() {
  var ciudadSelect = document.getElementById("pais2");
  ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
  var pais = document.getElementById("idp");
  var ciudadSelect2 = document.getElementById("departamento2");
  ciudadSelect2.innerHTML = ""; // Limpiar la lista de ciudades
  $('#sede2').selectize()[0].selectize.destroy(); 
  var sede = document.getElementById("sede2");
  sede.innerHTML = "";

  // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
          var opcionesCiudades = this.responseText;
          ciudadSelect.innerHTML = opcionesCiudades; // Actualizar la lista de ciudades
      }
  };
  xhr.open("POST", "/obtenerdepsedesanamente2", true);
  xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhr.send("area=" + 4 +"&pais="+pais.value);
}

function actualizarmun(paisSeleccionado) {
  var ciudadSelect = document.getElementById("departamento2");

  ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
  $('#sede2').selectize()[0].selectize.destroy(); 
  var sede = document.getElementById("sede2");
  sede.innerHTML = "";
  // Limpiar la lista de ciudades
  // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
          var pcionesmunsede= this.responseText;
          ciudadSelect.innerHTML = pcionesmunsede; // Actualizar la lista de ciudades
      }
  };
  xhr.open("POST", "/obtenermunsedesanamente2", true);
  xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhr.send("depsede=" + paisSeleccionado +"&area="+4);
}

function actualizarsede(paisSeleccionado) {
  var ciudadSelect = document.getElementById("sede2");

  ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
  $('#sede2').selectize()[0].selectize.destroy(); 
  // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
          var opcionesmunsede= this.responseText;
          ciudadSelect.innerHTML = opcionesmunsede; // Actualizar la lista de ciudades
          $('#sede2').selectize({
            sortField: 'text'
        });
        }
  };
  xhr.open("POST", "/obtenersedesanamente2", true);
  xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhr.send("munsede=" + paisSeleccionado +"&area="+4);
}
 
function verificarSeleccion() {
  var select = document.getElementById('sede2');
  if (select.value === '') {
      alert('Por favor, selecciona una sede.');
      return false;
  }
  return true;
}

function activa1(){
  var input = document.getElementById("nombreCompleto");
  var input2 = document.getElementById("DUIM");
  var input3 = document.getElementById("edad");
  var input4 = document.getElementById("TM");

  var output = document.getElementById("nombrep2");
  var output2 = document.getElementById("identidad");
  var output3 = document.getElementById("edadl");
  var output4 = document.getElementById("telefono")


    output.innerHTML = input.value;
    output2.innerHTML = input2.value;
    output3.innerHTML = input3.value;
    output4.innerHTML = input4.value;

  }

  function validateFormat() {
    var formatInput = document.getElementById("DUIM").value;
    var formatRegex = /^\d+$/;
      
       if (!formatRegex.test(formatInput)) {
       var alertContainer = document.getElementById("alertContainer");
       alertContainer.innerHTML = '<div class="alert alert-danger" role="alert">El formato del dui no es válido.</div>';
       alertContainer.style.display = "block";
       return false;
 }
 return true;
}

  function hideAlert() {
    var alertContainer = document.getElementById("alertContainer");
    alertContainer.style.display = "none";
  }