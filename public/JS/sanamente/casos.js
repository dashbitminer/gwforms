// Example starter JavaScript for disabling form submissions if there are invalid fields
(() => {
  'use strict'

  // Fetch all the forms we want to apply custom Bootstrap validation styles to
  const forms = document.querySelectorAll('.needs-validation')

  // Loop over them and prevent submission
  Array.from(forms).forEach(form => {
    form.addEventListener('submit', event => {
      validateFormat()
      if (!form.checkValidity()) {
        event.preventDefault()
        event.stopPropagation()
      }
    }, false)
  })
})()

function activaselect2(){
  var select = document.getElementById("aut");
var hiddenDiv = document.getElementById("hidden-question-div");
var input = document.getElementById("p30");
var input2 = document.getElementById("p23");
var input3 = document.getElementById("p24");
var input4 = document.getElementById("p27");
var input5 = document.getElementById("p28");
var input6 = document.getElementById("p29");
var input6 = document.getElementById("p31");
var input7 = document.getElementById("p138");
  if(select.value==="46"){    
      hiddenDiv.classList.remove("d-none");
      input.required = true;
      input2.required = true;
      input3.required = true;
      input4.required = true;
      input5.required = true;
      input6.required = true;
      input7.required = true;
  }
  
   else {
  hiddenDiv.classList.add("d-none");
  input.required = false;
  input2.required = false;
  input3.required = false;
  input4.required = false;
  input5.required = false;
  input6.required = false;
  input7.required = false;
  }

  }

function activaselect3(){
var hiddenDiv2 = document.getElementById("hidden-question-div2");
var input = document.getElementById("p24");
var input2 = document.getElementById("p25");
var input3 = document.getElementById("p26");
  if(input.value<18){ 
    hiddenDiv2.classList.remove("d-none");
    input2.required = true;
    input3.required = true;
     }
  else{
  hiddenDiv2.classList.add("d-none");
  input2.required = false;
  input3.required = false;

  }
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
    $('#p29').selectize()[0].selectize.destroy(); 
    var ciudadSelect = document.getElementById("p29");
    ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
    // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var opcionesmunsede= this.responseText;
            ciudadSelect.innerHTML = opcionesmunsede; // Actualizar la lista de ciudades
            $('#p29').selectize({
              sortField: 'text'
          });
          }
    };
    xhr.open("POST", "/obtenerciudades", true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.send("pais=" + paisSeleccionado);
}

