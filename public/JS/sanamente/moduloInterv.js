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
              xhr.open("POST", "/obtenermunsedesanamente", true);
              xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
              xhr.send("depsede=" + paisSeleccionado);
            }

            function actualizarsede(paisSeleccionado) {
              var ciudadSelect = document.getElementById("sede2");
             // $('#sede2').selectpicker()[0].selectpicker.destroy(); anterior
              $('#sede2').selectpicker('destroy');
              // $('#sede2').selectpicker('destroy');
              ciudadSelect.innerHTML = ""; // Limpiar la lista de ciudades
              // Enviar petición AJAX para obtener las ciudades correspondientes al país seleccionado
  
              var xhr = new XMLHttpRequest();
              xhr.onreadystatechange = function() {
                  if (this.readyState == 4 && this.status == 200) {
                    var opcionesmunsede= this.responseText;
                      ciudadSelect.innerHTML = opcionesmunsede;
                      $('#sede2').selectpicker(); // Actualizar la lista de ciudades
                  }
              };
              xhr.open("POST", "/obtenersedesanamente", true);
              xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
              xhr.send("munsede=" + paisSeleccionado);
            }           

/* function mostrar(){
  var p9 = document.getElementById("pregunta8.1")
  p9.classList.remove('d-none')
  
  var input= document.getElementById("pregunta8.1");
  var input2= document.getElementById("150");
  input.required = true;
  input2.required = true;
}

function ocultar(){
  var p9 = document.getElementById("pregunta8.1")
  p9.classList.add('d-none')

  var input= document.getElementById("pregunta8.1");
  var input2= document.getElementById("150");
  input.required = false;
  input2.required = false;
} */

function quitar(id){
  var pregunta = document.getElementById(id);
  pregunta.classList.remove('pregunta-vacia');
}
  
function prueba() {
  var miFormulario = document.getElementById("formualrioRegistro");
  if (!miFormulario) return; // Evita errores si el formulario no existe

  for (let i = 1; i <= 6; i++) { // Iterar por pregunta1, pregunta2, ..., pregunta6
      let pregunta = document.getElementById(`pregunta${i}`);
      let respuestaSeleccionada = miFormulario.querySelector(`input[name="p${i}"]:checked`);

      if (!respuestaSeleccionada) {
          pregunta.classList.add('pregunta-vacia');
      } else {
          pregunta.classList.remove('pregunta-vacia'); // Quita la clase si ya respondió
      }
  }
}

document.addEventListener("DOMContentLoaded", function () {
  let hoy = new Date();
  
  // Convertir la fecha a la zona horaria del usuario
  let fechaLocal = new Date(hoy.getFullYear(), hoy.getMonth(), hoy.getDate());

  // Formatear la fecha como YYYY-MM-DD
  let fechaFormateada = fechaLocal.toISOString().split("T")[0];

  document.getElementById("fechaform").value = fechaFormateada;
});



  function hideAlert() {
    var alertContainer = document.getElementById("alertContainer");
    alertContainer.style.display = "none";
  }