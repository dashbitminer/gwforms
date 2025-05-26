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
  
  function actualizardepto(paisSeleccionado) {
    var ciudadSelect = document.getElementById("pais2");
    ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
    // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var opcionesmunsede= this.responseText;
            ciudadSelect.innerHTML = opcionesmunsede; // Actualizar la lista de ciudades
          }
    };
    xhr.open("POST", "/obtenerdepsedesalud", true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.send("paissede=" + paisSeleccionado);
    ciudadSelect.required = true;
  }
  
function actualizarmun(depSeleccionado) {
    var ciudadSelect = document.getElementById("departamento2");
    ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
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
    xhr.send("depsede=" + depSeleccionado);
}

function actualizarsede(munSeleccionado) {
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
    xhr.send("munsede=" + munSeleccionado);
}


