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
 
  function actualizarCiudades(paisSeleccionado) {
    var ciudadSelect = document.getElementById("Municipio");
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

function Otro(){
  document.getElementById("otrodiv").classList.remove("d-none");
}

function Otro2(){
  document.getElementById("otrodiv").classList.add("d-none");
}